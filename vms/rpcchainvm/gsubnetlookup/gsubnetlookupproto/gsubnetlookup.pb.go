// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gsubnetlookup.proto

package gsubnetlookupproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SubnetIDRequest struct {
	ChainID              []byte   `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubnetIDRequest) Reset()         { *m = SubnetIDRequest{} }
func (m *SubnetIDRequest) String() string { return proto.CompactTextString(m) }
func (*SubnetIDRequest) ProtoMessage()    {}
func (*SubnetIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2050f029eebeabc, []int{0}
}

func (m *SubnetIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubnetIDRequest.Unmarshal(m, b)
}
func (m *SubnetIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubnetIDRequest.Marshal(b, m, deterministic)
}
func (m *SubnetIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubnetIDRequest.Merge(m, src)
}
func (m *SubnetIDRequest) XXX_Size() int {
	return xxx_messageInfo_SubnetIDRequest.Size(m)
}
func (m *SubnetIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubnetIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubnetIDRequest proto.InternalMessageInfo

func (m *SubnetIDRequest) GetChainID() []byte {
	if m != nil {
		return m.ChainID
	}
	return nil
}

type SubnetIDResponse struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubnetIDResponse) Reset()         { *m = SubnetIDResponse{} }
func (m *SubnetIDResponse) String() string { return proto.CompactTextString(m) }
func (*SubnetIDResponse) ProtoMessage()    {}
func (*SubnetIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2050f029eebeabc, []int{1}
}

func (m *SubnetIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubnetIDResponse.Unmarshal(m, b)
}
func (m *SubnetIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubnetIDResponse.Marshal(b, m, deterministic)
}
func (m *SubnetIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubnetIDResponse.Merge(m, src)
}
func (m *SubnetIDResponse) XXX_Size() int {
	return xxx_messageInfo_SubnetIDResponse.Size(m)
}
func (m *SubnetIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubnetIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubnetIDResponse proto.InternalMessageInfo

func (m *SubnetIDResponse) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*SubnetIDRequest)(nil), "gsubnetlookupproto.SubnetIDRequest")
	proto.RegisterType((*SubnetIDResponse)(nil), "gsubnetlookupproto.SubnetIDResponse")
}

func init() { proto.RegisterFile("gsubnetlookup.proto", fileDescriptor_a2050f029eebeabc) }

var fileDescriptor_a2050f029eebeabc = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2f, 0x2e, 0x4d,
	0xca, 0x4b, 0x2d, 0xc9, 0xc9, 0xcf, 0xcf, 0x2e, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x42, 0x11, 0x04, 0x8b, 0x29, 0x69, 0x73, 0xf1, 0x07, 0x83, 0x05, 0x3d, 0x5d, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x33, 0x12, 0x33, 0xf3, 0x3c, 0x5d,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x25, 0x25, 0x2e, 0x01, 0x84, 0xe2, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0xa8, 0x42, 0xa6, 0xcc, 0x14,
	0xa3, 0x54, 0x2e, 0x1e, 0x88, 0x1a, 0x1f, 0xb0, 0x2d, 0x42, 0xa1, 0x5c, 0x1c, 0x30, 0x3d, 0x42,
	0xca, 0x7a, 0x98, 0x2e, 0xd0, 0x43, 0xb3, 0x5e, 0x4a, 0x05, 0xbf, 0x22, 0x88, 0xb5, 0x49, 0x6c,
	0x60, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x17, 0xb8, 0x69, 0xe9, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubnetLookupClient is the client API for SubnetLookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubnetLookupClient interface {
	SubnetID(ctx context.Context, in *SubnetIDRequest, opts ...grpc.CallOption) (*SubnetIDResponse, error)
}

type subnetLookupClient struct {
	cc grpc.ClientConnInterface
}

func NewSubnetLookupClient(cc grpc.ClientConnInterface) SubnetLookupClient {
	return &subnetLookupClient{cc}
}

func (c *subnetLookupClient) SubnetID(ctx context.Context, in *SubnetIDRequest, opts ...grpc.CallOption) (*SubnetIDResponse, error) {
	out := new(SubnetIDResponse)
	err := c.cc.Invoke(ctx, "/gsubnetlookupproto.SubnetLookup/SubnetID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubnetLookupServer is the server API for SubnetLookup service.
type SubnetLookupServer interface {
	SubnetID(context.Context, *SubnetIDRequest) (*SubnetIDResponse, error)
}

// UnimplementedSubnetLookupServer can be embedded to have forward compatible implementations.
type UnimplementedSubnetLookupServer struct {
}

func (*UnimplementedSubnetLookupServer) SubnetID(ctx context.Context, req *SubnetIDRequest) (*SubnetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubnetID not implemented")
}

func RegisterSubnetLookupServer(s *grpc.Server, srv SubnetLookupServer) {
	s.RegisterService(&_SubnetLookup_serviceDesc, srv)
}

func _SubnetLookup_SubnetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubnetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubnetLookupServer).SubnetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gsubnetlookupproto.SubnetLookup/SubnetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubnetLookupServer).SubnetID(ctx, req.(*SubnetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubnetLookup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gsubnetlookupproto.SubnetLookup",
	HandlerType: (*SubnetLookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubnetID",
			Handler:    _SubnetLookup_SubnetID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gsubnetlookup.proto",
}
