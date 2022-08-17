// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/participationrewards/v1/proposals.proto

package types

import (
	encoding_json "encoding/json"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AddProtocolDataProposal struct {
	Title       string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Protocol    string                   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty" yaml:"protocol"`
	Type        string                   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty" yaml:"type"`
	Key         string                   `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	Data        encoding_json.RawMessage `protobuf:"bytes,6,opt,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty" yaml:"data"`
}

func (m *AddProtocolDataProposal) Reset()      { *m = AddProtocolDataProposal{} }
func (*AddProtocolDataProposal) ProtoMessage() {}
func (*AddProtocolDataProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94433b2236a43ef, []int{0}
}
func (m *AddProtocolDataProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddProtocolDataProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddProtocolDataProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddProtocolDataProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProtocolDataProposal.Merge(m, src)
}
func (m *AddProtocolDataProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddProtocolDataProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProtocolDataProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddProtocolDataProposal proto.InternalMessageInfo

type AddProtocolDataProposalWithDeposit struct {
	Title       string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Protocol    string                   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty" yaml:"protocol"`
	Type        string                   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty" yaml:"type"`
	Key         string                   `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	Data        encoding_json.RawMessage `protobuf:"bytes,6,opt,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty" yaml:"data"`
	Deposit     string                   `protobuf:"bytes,7,opt,name=deposit,proto3" json:"deposit,omitempty" yaml:"deposit"`
}

func (m *AddProtocolDataProposalWithDeposit) Reset()         { *m = AddProtocolDataProposalWithDeposit{} }
func (m *AddProtocolDataProposalWithDeposit) String() string { return proto.CompactTextString(m) }
func (*AddProtocolDataProposalWithDeposit) ProtoMessage()    {}
func (*AddProtocolDataProposalWithDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_d94433b2236a43ef, []int{1}
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddProtocolDataProposalWithDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProtocolDataProposalWithDeposit.Merge(m, src)
}
func (m *AddProtocolDataProposalWithDeposit) XXX_Size() int {
	return m.Size()
}
func (m *AddProtocolDataProposalWithDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProtocolDataProposalWithDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_AddProtocolDataProposalWithDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddProtocolDataProposal)(nil), "quicksilver.interchainstaking.v1.AddProtocolDataProposal")
	proto.RegisterType((*AddProtocolDataProposalWithDeposit)(nil), "quicksilver.interchainstaking.v1.AddProtocolDataProposalWithDeposit")
}

func init() {
	proto.RegisterFile("quicksilver/participationrewards/v1/proposals.proto", fileDescriptor_d94433b2236a43ef)
}

var fileDescriptor_d94433b2236a43ef = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xbf, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0x93, 0xeb, 0xf5, 0x07, 0x6e, 0xd5, 0x4a, 0xa6, 0x12, 0xa1, 0x42, 0xc9, 0xc9, 0x2c,
	0x48, 0x40, 0xac, 0xaa, 0xdb, 0x2d, 0x88, 0xd2, 0x15, 0xa9, 0xca, 0x82, 0x84, 0x84, 0x90, 0x2f,
	0xb1, 0x7c, 0x8f, 0xcb, 0xd9, 0x26, 0xf6, 0x5d, 0xc9, 0x7f, 0xd0, 0x91, 0xb1, 0xe3, 0xfd, 0x39,
	0x2c, 0x48, 0x1d, 0x99, 0x22, 0x74, 0xb7, 0x30, 0x67, 0x64, 0x42, 0x71, 0x72, 0x70, 0x03, 0x2c,
	0x6c, 0xdd, 0xf2, 0xde, 0xe7, 0xf3, 0x5e, 0xec, 0xaf, 0x64, 0x74, 0xf6, 0x71, 0x06, 0xe9, 0xc4,
	0x40, 0x3e, 0xe7, 0x05, 0xd5, 0xac, 0xb0, 0x90, 0x82, 0x66, 0x16, 0x94, 0x2c, 0xf8, 0x15, 0x2b,
	0x32, 0x43, 0xe7, 0xa7, 0x54, 0x17, 0x4a, 0x2b, 0xc3, 0x72, 0x13, 0xeb, 0x42, 0x59, 0x85, 0x07,
	0x1b, 0x43, 0x31, 0x48, 0xcb, 0x8b, 0x74, 0xcc, 0x40, 0x1a, 0xcb, 0x26, 0x20, 0x45, 0x3c, 0x3f,
	0x3d, 0x39, 0x16, 0x4a, 0x28, 0x27, 0xd3, 0xe6, 0xab, 0x9d, 0x3b, 0x79, 0x98, 0x2a, 0x33, 0x55,
	0xe6, 0x7d, 0x0b, 0xda, 0xa2, 0x43, 0x8f, 0x84, 0x52, 0x22, 0xe7, 0x94, 0x69, 0xa0, 0x4c, 0x4a,
	0x65, 0xdd, 0x19, 0x3a, 0x4a, 0x6e, 0x7a, 0xe8, 0xc1, 0xcb, 0x2c, 0xbb, 0x6c, 0x8a, 0x54, 0xe5,
	0x17, 0xcc, 0xb2, 0xcb, 0xee, 0x4c, 0xf8, 0x18, 0x6d, 0x5b, 0xb0, 0x39, 0x0f, 0xfc, 0x81, 0xff,
	0xe4, 0x5e, 0xd2, 0x16, 0x78, 0x80, 0xf6, 0x33, 0x6e, 0xd2, 0x02, 0x74, 0xb3, 0x27, 0xe8, 0x39,
	0xb6, 0xd9, 0xc2, 0x14, 0xed, 0xe9, 0x6e, 0x5f, 0xb0, 0xd5, 0xe0, 0xf3, 0xfb, 0x75, 0x15, 0x1d,
	0x95, 0x6c, 0x9a, 0x0f, 0xc9, 0x9a, 0x90, 0xe4, 0xb7, 0x84, 0x1f, 0xa3, 0xbe, 0x2d, 0x35, 0x0f,
	0xfa, 0x4e, 0x3e, 0xaa, 0xab, 0x68, 0xbf, 0x95, 0x9b, 0x2e, 0x49, 0x1c, 0xc4, 0x03, 0xb4, 0x35,
	0xe1, 0x65, 0xb0, 0xed, 0x9c, 0xc3, 0xba, 0x8a, 0x50, 0xeb, 0x4c, 0x78, 0x49, 0x92, 0x06, 0xe1,
	0x17, 0xa8, 0x9f, 0x31, 0xcb, 0x82, 0x1d, 0xa7, 0x3c, 0xfd, 0xb3, 0xa6, 0xe9, 0x92, 0x9f, 0x55,
	0x14, 0x70, 0x99, 0xaa, 0x0c, 0xa4, 0xa0, 0x1f, 0x8c, 0x92, 0x71, 0xc2, 0xae, 0x5e, 0x73, 0x63,
	0x98, 0xe0, 0x89, 0x1b, 0x1c, 0x1e, 0x5c, 0x2f, 0x22, 0xef, 0x66, 0x11, 0x79, 0x3f, 0x16, 0x91,
	0x47, 0xbe, 0xf6, 0x10, 0xf9, 0x47, 0x34, 0x6f, 0xc0, 0x8e, 0x2f, 0xb8, 0x56, 0x06, 0xec, 0x9d,
	0x4e, 0xe9, 0xe0, 0x3f, 0x52, 0xc2, 0xcf, 0xd0, 0x6e, 0xd6, 0xde, 0x3d, 0xd8, 0x75, 0xbf, 0xc1,
	0x75, 0x15, 0x1d, 0x76, 0x3b, 0x5a, 0x40, 0x92, 0xb5, 0x32, 0xdc, 0xbb, 0xee, 0xf2, 0x3c, 0x7f,
	0xf7, 0x65, 0x19, 0xfa, 0xb7, 0xcb, 0xd0, 0xff, 0xbe, 0x0c, 0xfd, 0xcf, 0xab, 0xd0, 0xbb, 0x5d,
	0x85, 0xde, 0xb7, 0x55, 0xe8, 0xbd, 0x7d, 0x25, 0xc0, 0x8e, 0x67, 0xa3, 0x38, 0x55, 0x53, 0x0a,
	0x52, 0x70, 0x39, 0x03, 0x5b, 0x3e, 0x1f, 0xcd, 0x20, 0xcf, 0xe8, 0xe6, 0x2b, 0xfa, 0xf4, 0xf7,
	0x77, 0xd4, 0x5c, 0xdc, 0x8c, 0x76, 0x5c, 0x50, 0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38,
	0x26, 0x5c, 0xb8, 0x78, 0x03, 0x00, 0x00,
}

func (m *AddProtocolDataProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddProtocolDataProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddProtocolDataProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Protocol) > 0 {
		i -= len(m.Protocol)
		copy(dAtA[i:], m.Protocol)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Protocol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AddProtocolDataProposalWithDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddProtocolDataProposalWithDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddProtocolDataProposalWithDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Protocol) > 0 {
		i -= len(m.Protocol)
		copy(dAtA[i:], m.Protocol)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Protocol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddProtocolDataProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func (m *AddProtocolDataProposalWithDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddProtocolDataProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddProtocolDataProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddProtocolDataProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = encoding_json.RawMessage(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AddProtocolDataProposalWithDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AddProtocolDataProposalWithDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddProtocolDataProposalWithDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)