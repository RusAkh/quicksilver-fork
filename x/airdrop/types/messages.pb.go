// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: quicksilver/airdrop/v1/messages.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgClaim struct {
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	Action  int32  `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty" yaml:"action"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Proof   []byte `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty" yaml:"proof"`
}

func (m *MsgClaim) Reset()         { *m = MsgClaim{} }
func (m *MsgClaim) String() string { return proto.CompactTextString(m) }
func (*MsgClaim) ProtoMessage()    {}
func (*MsgClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b0828c7de1949a1, []int{0}
}
func (m *MsgClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaim.Merge(m, src)
}
func (m *MsgClaim) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaim.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaim proto.InternalMessageInfo

type MsgClaimResponse struct {
	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty" yaml:"amount"`
}

func (m *MsgClaimResponse) Reset()         { *m = MsgClaimResponse{} }
func (m *MsgClaimResponse) String() string { return proto.CompactTextString(m) }
func (*MsgClaimResponse) ProtoMessage()    {}
func (*MsgClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b0828c7de1949a1, []int{1}
}
func (m *MsgClaimResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimResponse.Merge(m, src)
}
func (m *MsgClaimResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgClaim)(nil), "quicksilver.airdrop.v1.MsgClaim")
	proto.RegisterType((*MsgClaimResponse)(nil), "quicksilver.airdrop.v1.MsgClaimResponse")
}

func init() {
	proto.RegisterFile("quicksilver/airdrop/v1/messages.proto", fileDescriptor_2b0828c7de1949a1)
}

var fileDescriptor_2b0828c7de1949a1 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0xed, 0x6e, 0xbb, 0x0e, 0xf5, 0x57, 0x2c, 0xb2, 0x2e, 0x92, 0x84, 0x81, 0xc2,
	0x5a, 0x35, 0xa1, 0x0a, 0x1e, 0xf6, 0xb8, 0x3d, 0x79, 0x28, 0x48, 0x8e, 0x5e, 0xca, 0x6c, 0x32,
	0x4e, 0x07, 0x93, 0x79, 0x31, 0x33, 0x59, 0xba, 0x37, 0xf1, 0xd4, 0xa3, 0xe0, 0x3f, 0xd0, 0x3f,
	0xa7, 0xde, 0x0a, 0x5e, 0x3c, 0x05, 0xd9, 0xf5, 0xe0, 0x79, 0xff, 0x02, 0xc9, 0x4c, 0x22, 0x5d,
	0x10, 0x7a, 0x7b, 0x79, 0xdf, 0x4f, 0xbe, 0xf9, 0xbe, 0xbc, 0x87, 0x0f, 0x3e, 0x55, 0x22, 0xf9,
	0xa8, 0x44, 0x36, 0x67, 0x65, 0x44, 0x45, 0x99, 0x96, 0x50, 0x44, 0xf3, 0xa3, 0x28, 0x67, 0x4a,
	0x51, 0xce, 0x54, 0x58, 0x94, 0xa0, 0xc1, 0x7d, 0x7c, 0x03, 0x0b, 0x5b, 0x2c, 0x9c, 0x1f, 0x8d,
	0xf6, 0x39, 0x70, 0x30, 0x48, 0xd4, 0x54, 0x96, 0x1e, 0x3d, 0x49, 0x40, 0xe5, 0xa0, 0x4e, 0xad,
	0x60, 0x1f, 0x5a, 0xe9, 0x29, 0x07, 0xe0, 0x19, 0x8b, 0x68, 0x21, 0x22, 0x2a, 0x25, 0x68, 0xaa,
	0x05, 0xc8, 0x56, 0x25, 0xdf, 0x11, 0x1e, 0x9c, 0x28, 0x7e, 0x9c, 0x51, 0x91, 0xbb, 0x21, 0x1e,
	0x24, 0x67, 0x54, 0xc8, 0x53, 0x91, 0x0e, 0x51, 0x80, 0xc6, 0x77, 0xa6, 0x8f, 0xd6, 0xb5, 0x7f,
	0x7f, 0x41, 0xf3, 0x6c, 0x42, 0x3a, 0x85, 0xc4, 0xbb, 0xa6, 0x7c, 0x9b, 0xba, 0xcf, 0xf0, 0x0e,
	0x4d, 0x1a, 0xb7, 0xe1, 0x56, 0x80, 0xc6, 0xfd, 0xe9, 0xc3, 0x75, 0xed, 0xdf, 0xb5, 0xb4, 0xed,
	0x93, 0xb8, 0x05, 0xdc, 0x17, 0x78, 0x97, 0xa6, 0x69, 0xc9, 0x94, 0x1a, 0x6e, 0x1b, 0x67, 0x77,
	0x5d, 0xfb, 0xf7, 0x5a, 0xd6, 0x0a, 0x24, 0xee, 0x10, 0xf7, 0x10, 0xf7, 0x8b, 0x12, 0xe0, 0xc3,
	0xb0, 0x17, 0xa0, 0xf1, 0xde, 0x74, 0xff, 0xaa, 0xf6, 0xd1, 0xba, 0xf6, 0xf7, 0x2c, 0x6f, 0x24,
	0x12, 0x5b, 0x64, 0x32, 0xb8, 0xb8, 0xf4, 0x9d, 0x3f, 0x97, 0xbe, 0x43, 0x8e, 0xf1, 0x83, 0x6e,
	0x94, 0x98, 0xa9, 0x02, 0xa4, 0x62, 0x26, 0x62, 0x0e, 0x95, 0xd4, 0x66, 0xa0, 0xde, 0x46, 0x44,
	0xd3, 0x6f, 0x22, 0x9a, 0x62, 0xd2, 0x6b, 0x8c, 0x5e, 0x5d, 0x20, 0xbc, 0x7d, 0xa2, 0xb8, 0xfb,
	0x19, 0xe1, 0xbe, 0xfd, 0x2b, 0x41, 0xf8, 0xff, 0x55, 0x84, 0xdd, 0xc7, 0x46, 0xe3, 0xdb, 0x88,
	0x2e, 0x0e, 0x79, 0xfe, 0xe5, 0xc7, 0xef, 0x6f, 0x5b, 0x07, 0x24, 0x88, 0x6e, 0x5e, 0x81, 0x3e,
	0x6f, 0x0e, 0xa0, 0xbb, 0x85, 0xa4, 0x79, 0x63, 0x82, 0x0e, 0xa7, 0xef, 0xae, 0x96, 0x1e, 0xba,
	0x5e, 0x7a, 0xe8, 0xd7, 0xd2, 0x43, 0x5f, 0x57, 0x9e, 0x73, 0xbd, 0xf2, 0x9c, 0x9f, 0x2b, 0xcf,
	0x79, 0xff, 0x86, 0x0b, 0x7d, 0x56, 0xcd, 0xc2, 0x04, 0xf2, 0x48, 0x48, 0xce, 0x64, 0x25, 0xf4,
	0xe2, 0xe5, 0xac, 0x12, 0x59, 0xba, 0x61, 0x7c, 0xfe, 0xcf, 0x54, 0x2f, 0x0a, 0xa6, 0x66, 0x3b,
	0x66, 0xe9, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x9c, 0x9a, 0x40, 0x84, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Claim(ctx context.Context, in *MsgClaim, opts ...grpc.CallOption) (*MsgClaimResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Claim(ctx context.Context, in *MsgClaim, opts ...grpc.CallOption) (*MsgClaimResponse, error) {
	out := new(MsgClaimResponse)
	err := c.cc.Invoke(ctx, "/quicksilver.airdrop.v1.Msg/Claim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Claim(context.Context, *MsgClaim) (*MsgClaimResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Claim(ctx context.Context, req *MsgClaim) (*MsgClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Claim not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/quicksilver.airdrop.v1.Msg/Claim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Claim(ctx, req.(*MsgClaim))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "quicksilver.airdrop.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Claim",
			Handler:    _Msg_Claim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quicksilver/airdrop/v1/messages.proto",
}

func (m *MsgClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Action != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgClaimResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + sovMessages(uint64(m.Action))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgClaimResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Amount != 0 {
		n += 1 + sovMessages(uint64(m.Amount))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgClaimResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgClaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)