#!/usr/bin/make -f

DOCKER_BUILDKIT=1
COSMOS_BUILD_OPTIONS ?= ""
PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIM=github.com/ingenuity-build/quicksilver/test/simulation
PACKAGES_E2E=$(shell go list ./... | grep '/e2e')
VERSION=$(shell git describe --tags | head -n1)
DOCKER_VERSION ?= $(VERSION)
TMVERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
BINDIR ?= $(GOPATH)/bin
QS_BINARY = quicksilverd
QS_DIR = quicksilver
BUILDDIR ?= $(CURDIR)/build
HTTPS_GIT := https://github.com/ingenuity-build/quicksilver.git

DOCKER := $(shell which docker)
DOCKERCOMPOSE := $(shell which docker-compose)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
COMMIT_HASH := $(shell git rev-parse --short=7 HEAD)
DOCKER_TAG := $(COMMIT_HASH)

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

export GO111MODULE = on

# Default target executed when no arguments are given to make.
default_target: all

.PHONY: default_target build

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=quicksilver \
          -X github.com/cosmos/cosmos-sdk/version.AppName=$(QS_BINARY) \
          -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
          -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
          -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TMVERSION)

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  CGO_ENABLED=1
  build_tags += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif
# handle pebbledb
ifeq (pebbledb,$(findstring pebbledb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += pebbledb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb
endif

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
	build_tags += muslc
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))
ldflags += -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# # The below include contains the tools and runsim targets.
# include contrib/devtools/Makefile

###############################################################################
###                                  Build                                  ###
###############################################################################

BUILD_TARGETS := build install

check_version:
ifneq ($(GO_MINOR_VERSION),20)
	@echo "ERROR: Go version 1.20 is required for building Quicksilver. There are consensus breaking changes between binaries compiled with and without Go 1.20."
	exit 1
endif

build: BUILD_ARGS=-o $(BUILDDIR)/

build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): check_version go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./cmd/quicksilverd

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-docker:
	DOCKER_BUILDKIT=1 $(DOCKER) build . -f Dockerfile -t quicksilverzone/quicksilver:$(DOCKER_VERSION) -t quicksilverzone/quicksilver:latest

build-docker-local: build
	DOCKER_BUILDKIT=1 $(DOCKER) build -f Dockerfile.local . -t quicksilverzone/quicksilver:$(DOCKER_VERSION)

build-docker-release: build-docker
	$(DOCKER)  run -v /tmp:/tmp quicksilverzone/quicksilver:$(DOCKER_VERSION) cp /usr/local/bin/quicksilverd /tmp/quicksilverd
	mv /tmp/quicksilverd build/quicksilverd-$(DOCKER_VERSION)-amd64

push-docker: build-docker
	$(DOCKERCOMPOSE) push quicksilver

reload-docker:
	$(DOCKERCOMPOSE) up -d --force-recreate quicksilver

test-docker:
	./scripts/simple-test.sh
test-docker-regen:
	./scripts/simple-test.sh -r
test-docker-multi:
	./scripts/multi-test.sh
test-docker-multi-regen:
	./scripts/multi-test.sh -r
build-docker-all:
	$(DOCKERCOMPOSE) build

push-docker-all:
	$(DOCKERCOMPOSE) push

$(MOCKS_DIR):
	mkdir -p $(MOCKS_DIR)

distclean: clean tools-clean

clean:
	rm -rf \
    $(BUILDDIR)/ \
    artifacts/ \
    tmp-swagger-gen/

all: build vulncheck vet

build-all: tools build lint test

.PHONY: distclean clean build-all

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

TOOLS_DESTDIR  = $(GOPATH)/bin
STATIK         = $(TOOLS_DESTDIR)/statik
RUNSIM         = $(TOOLS_DESTDIR)/runsim

# Install the runsim binary with a temporary workaround of entering an outside
# directory as the "go get" command ignores the -mod option and will polute the
# go.{mod, sum} files.
#
# ref: https://github.com/golang/go/issues/30515
runsim: $(RUNSIM)
$(RUNSIM):
	@echo "Installing runsim..."
	@(cd /tmp && go install github.com/cosmos/tools/cmd/runsim@v1.0.0)


statik: $(STATIK)
$(STATIK):
	@echo "Installing statik..."
	@(cd /tmp && go install github.com/rakyll/statik@v0.1.6)

docs-tools:
ifeq (, $(shell which yarn))
	@echo "Installing yarn..."
	@npm install -g yarn
else
	@echo "yarn already installed; skipping..."
endif

tools: tools-stamp
tools-stamp: contract-tools docs-tools proto-tools statik runsim
	# Create dummy file to satisfy dependency and avoid
	# rebuilding when this Makefile target is hit twice
	# in a row.
	touch $@

tools-clean:
	rm -f $(RUNSIM)
	rm -f tools-stamp

docs-tools-stamp: docs-tools
	# Create dummy file to satisfy dependency and avoid
	# rebuilding when this Makefile target is hit twice
	# in a row.
	touch $@

.PHONY: runsim statik tools contract-tools docs-tools proto-tools  tools-stamp tools-clean docs-tools-stamp

go.sum: go.mod
	echo "Ensure dependencies have not been modified ..." >&2
	go mod verify
	go mod tidy

###############################################################################
###                              Documentation                              ###
###############################################################################

update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
.PHONY: update-swagger-docs

godocs:
	@echo "--> Wait a few seconds and visit http://localhost:6060/pkg/github.com/ingenuity-build/quicksilver/types"
	godoc -http=:6060

# Start docs site at localhost:8080
docs-serve:
	@cd docs && \
	yarn && \
	yarn run serve

# Build the site into docs/.vuepress/dist
build-docs:
	@$(MAKE) docs-tools-stamp && \
	cd docs && \
	yarn && \
	yarn run build

# This builds a docs site for each branch/tag in `./docs/versions`
# and copies each site to a version prefixed path. The last entry inside
# the `versions` file will be the default root index.html.
build-docs-versioned:
	@$(MAKE) docs-tools-stamp && \
	cd docs && \
	while read -r branch path_prefix; do \
		(git checkout $${branch} && npm install && VUEPRESS_BASE="/$${path_prefix}/" npm run build) ; \
		mkdir -p ~/output/$${path_prefix} ; \
		cp -r .vuepress/dist/* ~/output/$${path_prefix}/ ; \
		cp ~/output/$${path_prefix}/index.html ~/output ; \
	done < versions ;

.PHONY: docs-serve build-docs build-docs-versioned

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test: test-unit
test-all: test-unit test-race
PACKAGES_UNIT=$(shell go list ./...)
TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-cover test-race

# Test runs-specific rules. To add a new test target, just add
# a new rule, customise ARGS or TEST_PACKAGES ad libitum, and
# append the new rule to the TEST_TARGETS list.
test-unit: ARGS=-timeout=10m -race
test-unit: TEST_PACKAGES=$(PACKAGES_UNIT)

test-race: ARGS=-race
test-race: TEST_PACKAGES=$(PACKAGES_NOSIMULATION)
$(TEST_TARGETS): run-tests

test-unit-cover: ARGS=-timeout=10m -race -coverprofile=coverage.txt -covermode=atomic
test-unit-cover: TEST_PACKAGES=$(PACKAGES_UNIT)

run-tests:
ifneq (,$(shell which tparse 2>/dev/null))
	go test -short -mod=readonly -json $(ARGS) $(EXTRA_ARGS) $(TEST_PACKAGES) | tparse
else
	go test -short -mod=readonly $(ARGS)  $(EXTRA_ARGS) $(TEST_PACKAGES)
endif

test-import:
	@go test -short ./tests/importer -v --vet=off --run=TestImportBlocks --datadir tmp \
	--blockchain blockchain
	rm -rf tests/importer/tmp

test-rpc:
	./scripts/integration-test-all.sh -t "rpc" -q 1 -z 1 -s 2 -m "rpc" -r "true"

test-rpc-pending:
	./scripts/integration-test-all.sh -t "pending" -q 1 -z 1 -s 2 -m "pending" -r "true"


vulncheck: $(BUILDDIR)/
	GOBIN=$(BUILDDIR) go install golang.org/x/vuln/cmd/govulncheck@latest
	$(BUILDDIR)/govulncheck ./...

vet:
	@echo "Running vet..."
	@go vet ./...
	@echo "Done!"

RUNNER_BASE_IMAGE_DISTROLESS := gcr.io/distroless/static-debian11
RUNNER_BASE_IMAGE_ALPINE := alpine:3.16
RUNNER_BASE_IMAGE_NONROOT := gcr.io/distroless/static-debian11:nonroot
docker-build-debug:
	@DOCKER_BUILDKIT=1 $(DOCKER) build -t quicksilver:${COMMIT} --build-arg BASE_IMG_TAG=debug --build-arg RUNNER_IMAGE=$(RUNNER_BASE_IMAGE_ALPINE) -f test/e2e/e2e.Dockerfile .
	@DOCKER_BUILDKIT=1 $(DOCKER) tag quicksilver:${COMMIT} quicksilver:debug

docker-build-e2e-init-chain:
	@DOCKER_BUILDKIT=1 docker build -t quicksilver-e2e-init-chain:debug --build-arg E2E_SCRIPT_NAME=chain --platform=linux/x86_64 -f test/e2e/initialization/init.Dockerfile .

docker-build-e2e-init-node:
	@DOCKER_BUILDKIT=1 docker build -t quicksilver-e2e-init-node:debug --build-arg E2E_SCRIPT_NAME=node --platform=linux/x86_64 -f test/e2e/initialization/init.Dockerfile .

build-e2e-script:
	mkdir -p $(BUILDDIR)
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./test/e2e/initialization/$(E2E_SCRIPT_NAME)

# test-e2e runs a full e2e test suite
# deletes any pre-existing QUICKSILVER containers before running.
#
# Deletes Docker resources at the end.
# Utilizes Go cache.
test-e2e: e2e-setup test-e2e-ci e2e-remove-resources

# TODO
# currently skipping upgrades

# test-e2e-ci runs a full e2e test suite
# does not do any validation about the state of the Docker environment
# As a result, avoid using this locally.
test-e2e-ci:
	@VERSION=$(VERSION) QUICKSILVER_E2E=True QUICKSILVER_E2E_DEBUG_LOG=False QUICKSILVER_E2E_SKIP_UPGRADE=True QUICKSILVER_E2E_UPGRADE_VERSION=$(E2E_UPGRADE_VERSION)  go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E)

# test-e2e-debug runs a full e2e test suite but does
# not attempt to delete Docker resources at the end.
test-e2e-debug: e2e-setup
	@VERSION=$(VERSION) QUICKSILVER_E2E=True QUICKSILVER_E2E_DEBUG_LOG=True QUICKSILVER_E2E_SKIP_UPGRADE=True QUICKSILVER_E2E_UPGRADE_VERSION=$(E2E_UPGRADE_VERSION) QUICKSILVER_E2E_SKIP_CLEANUP=True go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E) -count=1

# test-e2e-short runs the e2e test with only short tests.
# Does not delete any of the containers after running.
# Deletes any existing containers before running.
# Does not use Go cache.
test-e2e-short: e2e-setup
	@VERSION=$(VERSION) QUICKSILVER_E2E=True QUICKSILVER_E2E_DEBUG_LOG=True QUICKSILVER_E2E_SKIP_UPGRADE=True QUICKSILVER_E2E_SKIP_IBC=True QUICKSILVER_E2E_SKIP_STATE_SYNC=True QUICKSILVER_E2E_SKIP_CLEANUP=True go test -mod=readonly -timeout=25m -v $(PACKAGES_E2E) -count=1

e2e-setup: e2e-check-image-sha e2e-remove-resources
	@echo Finished e2e environment setup, ready to start the test

e2e-check-image-sha:
	test/e2e/scripts/run/check_image_sha.sh

e2e-remove-resources:
	test/e2e/scripts/run/remove_stale_resources.sh


.PHONY: run-tests test test-all test-import test-rpc $(TEST_TARGETS)

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_CI_NUM_BLOCKS ?= 125
SIM_CI_BLOCK_SIZE ?= 50
SIM_PERIOD ?= 5
SIM_COMMIT ?= true
SIM_TIMEOUT ?= 24h


test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@go test -short -mod=readonly $(PACKAGES_SIM) -run ^TestAppStateDeterminism -Enabled=true \
		-NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -Period=$(SIM_PERIOD) \
		-v -timeout $(SIM_TIMEOUT)

## test-sim-ci: Run lightweight simulation for CI pipeline
test-sim-ci:
	@echo "Running non-determinism test..."
	@go test -short -mod=readonly $(PACKAGES_SIM) -run ^TestAppStateDeterminism -Enabled=true \
		-NumBlocks=$(SIM_CI_NUM_BLOCKS) -BlockSize=$(SIM_CI_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -Period=$(SIM_PERIOD) \
		-v -timeout $(SIM_TIMEOUT)

test-sim-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.$(QS_DIR)/config/genesis.json will be used."
	@go test -short -mod=readonly $(PACKAGES_SIM) -run TestFullAppSimulation -Genesis=${HOME}/.$(QS_DIR)/config/genesis.json \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -Seed=99 \
		-Period=$(SIM_PERIOD) -v -timeout  $(SIM_TIMEOUT)

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(PACKAGES_SIM) -ExitOnFail 50 5 TestAppImportExport

test-sim-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(PACKAGES_SIM) -ExitOnFail 50 5 TestAppSimulationAfterImport

test-sim-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.$(QS_DIR)/config/genesis.json will be used."
	@$(BINDIR)/runsim -Genesis=${HOME}/.$(QS_DIR)/config/genesis.json -SimAppPkg=$(PACKAGES_SIM) -ExitOnFail 400 5 TestFullAppSimulation

test-sim-multi-seed-long: runsim
	@echo "Running long multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(PACKAGES_SIM) -ExitOnFail 500 50 TestFullAppSimulation

test-sim-multi-seed-short: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 10 TestFullAppSimulation

test-sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -short -mod=readonly $(SIMAPP) -benchmem -bench=BenchmarkInvariants -run=^$ \
	-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) \
	-Period=1 -Commit=$(SIM_COMMIT) -Seed=57 -v -timeout  $(SIM_TIMEOUT)

.PHONY: \
test-sim-nondeterminism \
test-sim-ci \
test-sim-custom-genesis-fast \
test-sim-import-export \
test-sim-after-import \
test-sim-custom-genesis-multi-seed \
test-sim-multi-seed-short \
test-sim-multi-seed-long \
test-sim-benchmark-invariants

benchmark:
	@go test -short -mod=readonly -bench=. $(PACKAGES_NOSIMULATION)
.PHONY: benchmark

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --out-format=tab

lint-fix:
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --fix --out-format=tab --issues-exit-code=0

.PHONY: lint lint-fix

format:
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -name '*.pb.go' -not -name '*.gw.go' | xargs go run mvdan.cc/gofumpt -w .
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -name '*.pb.go' -not -name '*.gw.go' | xargs go run github.com/client9/misspell/cmd/misspell -w
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -name '*.pb.go' -not -name '*.gw.go' | xargs go run golang.org/x/tools/cmd/goimports -w -local github.com/ingenuity-build/quicksilver
.PHONY: format

mdlint:
	@echo "--> Running markdown linter"
	@$(DOCKER) run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md"

mdlint-fix:
	@$(DOCKER)  run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md" --fix

###############################################################################
###                                Protobuf                                 ###
###############################################################################

BUF_VERSION=1.15.1

proto-all: proto-gen

proto-gen:
	@echo "🤖 Generating code from protobuf..."
	@$(DOCKER) run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		quicksilver-proto sh ./proto/generate.sh
	@echo "✅ Completed code generation!"

proto-lint:
	@echo "🤖 Running protobuf linter..."
	@$(DOCKER) run --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) lint
	@echo "✅ Completed protobuf linting!"

proto-format:
	@echo "🤖 Running protobuf format..."
	@$(DOCKER) run --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) format -w
	@echo "✅ Completed protobuf format!"

proto-breaking-check:
	@echo "🤖 Running protobuf breaking check against develop branch..."
	@$(DOCKER) run --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) breaking --against '.git#branch=develop'
	@echo "✅ Completed protobuf breaking check!"

proto-setup:
	@echo "🤖 Setting up protobuf environment..."
	@$(DOCKER) build --rm --tag quicksilver-proto:latest --file proto/Dockerfile .
	@echo "✅ Setup protobuf environment!"


