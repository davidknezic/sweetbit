// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package sweetrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Test struct {
	One                  int32    `protobuf:"varint,1,opt,name=one,proto3" json:"one,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (dst *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(dst, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

func (m *Test) GetOne() int32 {
	if m != nil {
		return m.One
	}
	return 0
}

type WpaNetwork struct {
	Ssid                 string   `protobuf:"bytes,1,opt,name=ssid,proto3" json:"ssid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WpaNetwork) Reset()         { *m = WpaNetwork{} }
func (m *WpaNetwork) String() string { return proto.CompactTextString(m) }
func (*WpaNetwork) ProtoMessage()    {}
func (*WpaNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}
func (m *WpaNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WpaNetwork.Unmarshal(m, b)
}
func (m *WpaNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WpaNetwork.Marshal(b, m, deterministic)
}
func (dst *WpaNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WpaNetwork.Merge(dst, src)
}
func (m *WpaNetwork) XXX_Size() int {
	return xxx_messageInfo_WpaNetwork.Size(m)
}
func (m *WpaNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_WpaNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_WpaNetwork proto.InternalMessageInfo

func (m *WpaNetwork) GetSsid() string {
	if m != nil {
		return m.Ssid
	}
	return ""
}

type SubscribeWpaNetworksRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeWpaNetworksRequest) Reset()         { *m = SubscribeWpaNetworksRequest{} }
func (m *SubscribeWpaNetworksRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeWpaNetworksRequest) ProtoMessage()    {}
func (*SubscribeWpaNetworksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}
func (m *SubscribeWpaNetworksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeWpaNetworksRequest.Unmarshal(m, b)
}
func (m *SubscribeWpaNetworksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeWpaNetworksRequest.Marshal(b, m, deterministic)
}
func (dst *SubscribeWpaNetworksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeWpaNetworksRequest.Merge(dst, src)
}
func (m *SubscribeWpaNetworksRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeWpaNetworksRequest.Size(m)
}
func (m *SubscribeWpaNetworksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeWpaNetworksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeWpaNetworksRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Test)(nil), "sweetrpc.Test")
	proto.RegisterType((*WpaNetwork)(nil), "sweetrpc.WpaNetwork")
	proto.RegisterType((*SubscribeWpaNetworksRequest)(nil), "sweetrpc.SubscribeWpaNetworksRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SweetClient is the client API for Sweet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SweetClient interface {
	GetFeature(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error)
	SubscribeWpaNetworks(ctx context.Context, in *SubscribeWpaNetworksRequest, opts ...grpc.CallOption) (Sweet_SubscribeWpaNetworksClient, error)
}

type sweetClient struct {
	cc *grpc.ClientConn
}

func NewSweetClient(cc *grpc.ClientConn) SweetClient {
	return &sweetClient{cc}
}

func (c *sweetClient) GetFeature(ctx context.Context, in *Test, opts ...grpc.CallOption) (*Test, error) {
	out := new(Test)
	err := c.cc.Invoke(ctx, "/sweetrpc.Sweet/GetFeature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sweetClient) SubscribeWpaNetworks(ctx context.Context, in *SubscribeWpaNetworksRequest, opts ...grpc.CallOption) (Sweet_SubscribeWpaNetworksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sweet_serviceDesc.Streams[0], "/sweetrpc.Sweet/SubscribeWpaNetworks", opts...)
	if err != nil {
		return nil, err
	}
	x := &sweetSubscribeWpaNetworksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Sweet_SubscribeWpaNetworksClient interface {
	Recv() (*WpaNetwork, error)
	grpc.ClientStream
}

type sweetSubscribeWpaNetworksClient struct {
	grpc.ClientStream
}

func (x *sweetSubscribeWpaNetworksClient) Recv() (*WpaNetwork, error) {
	m := new(WpaNetwork)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SweetServer is the server API for Sweet service.
type SweetServer interface {
	GetFeature(context.Context, *Test) (*Test, error)
	SubscribeWpaNetworks(*SubscribeWpaNetworksRequest, Sweet_SubscribeWpaNetworksServer) error
}

func RegisterSweetServer(s *grpc.Server, srv SweetServer) {
	s.RegisterService(&_Sweet_serviceDesc, srv)
}

func _Sweet_GetFeature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Test)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SweetServer).GetFeature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sweetrpc.Sweet/GetFeature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SweetServer).GetFeature(ctx, req.(*Test))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sweet_SubscribeWpaNetworks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeWpaNetworksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SweetServer).SubscribeWpaNetworks(m, &sweetSubscribeWpaNetworksServer{stream})
}

type Sweet_SubscribeWpaNetworksServer interface {
	Send(*WpaNetwork) error
	grpc.ServerStream
}

type sweetSubscribeWpaNetworksServer struct {
	grpc.ServerStream
}

func (x *sweetSubscribeWpaNetworksServer) Send(m *WpaNetwork) error {
	return x.ServerStream.SendMsg(m)
}

var _Sweet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sweetrpc.Sweet",
	HandlerType: (*SweetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeature",
			Handler:    _Sweet_GetFeature_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeWpaNetworks",
			Handler:       _Sweet_SubscribeWpaNetworks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x0a, 0x82, 0x40,
	0x10, 0x40, 0x5d, 0xd2, 0xc8, 0x39, 0x44, 0x0c, 0x1e, 0xc4, 0x08, 0x64, 0x21, 0xe8, 0xb4, 0x44,
	0xfd, 0x43, 0xdd, 0x3a, 0x68, 0xe0, 0x59, 0x6d, 0x0e, 0x12, 0xb4, 0xdb, 0xee, 0x88, 0xbf, 0xd1,
	0x27, 0x87, 0x42, 0x2c, 0x44, 0x74, 0x7b, 0xc3, 0x3c, 0xe6, 0x31, 0x10, 0x5b, 0xd3, 0x2a, 0x63,
	0x35, 0x6b, 0x5c, 0xb8, 0x81, 0x88, 0xad, 0x69, 0x65, 0x0a, 0xe1, 0x95, 0x1c, 0xe3, 0x0a, 0x66,
	0xfa, 0x41, 0xa9, 0xc8, 0xc5, 0x2e, 0x2a, 0x46, 0x94, 0x39, 0x40, 0x65, 0xea, 0x0b, 0xf1, 0xa0,
	0xed, 0x1d, 0x11, 0x42, 0xe7, 0xba, 0xdb, 0x24, 0xc4, 0xc5, 0xc4, 0x72, 0x03, 0xeb, 0xb2, 0x6f,
	0x5c, 0x6b, 0xbb, 0x86, 0xbc, 0xea, 0x0a, 0x7a, 0xf6, 0xe4, 0xf8, 0xf0, 0x12, 0x10, 0x95, 0x63,
	0x07, 0x15, 0xc0, 0x99, 0xf8, 0x44, 0x35, 0xf7, 0x96, 0x70, 0xa9, 0x3e, 0x75, 0x35, 0xa6, 0xb3,
	0xaf, 0x59, 0x06, 0x58, 0x41, 0xf2, 0xeb, 0x30, 0x6e, 0xbd, 0xf9, 0x27, 0x9c, 0x25, 0x5e, 0xf3,
	0x5b, 0x19, 0xec, 0x45, 0x33, 0x9f, 0xde, 0x3f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xff, 0xcb,
	0x66, 0xca, 0x0b, 0x01, 0x00, 0x00,
}