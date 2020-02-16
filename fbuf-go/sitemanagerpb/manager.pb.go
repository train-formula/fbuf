// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sitemanager/manager.proto

package sitemanagerpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type InitializeSiteRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeSiteRequest) Reset()         { *m = InitializeSiteRequest{} }
func (m *InitializeSiteRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeSiteRequest) ProtoMessage()    {}
func (*InitializeSiteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e393e113b6828dc9, []int{0}
}

func (m *InitializeSiteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeSiteRequest.Unmarshal(m, b)
}
func (m *InitializeSiteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeSiteRequest.Marshal(b, m, deterministic)
}
func (m *InitializeSiteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeSiteRequest.Merge(m, src)
}
func (m *InitializeSiteRequest) XXX_Size() int {
	return xxx_messageInfo_InitializeSiteRequest.Size(m)
}
func (m *InitializeSiteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeSiteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeSiteRequest proto.InternalMessageInfo

type InitializeSiteResponse struct {
	SiteId               string   `protobuf:"bytes,1,opt,name=site_id,json=siteId,proto3" json:"site_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeSiteResponse) Reset()         { *m = InitializeSiteResponse{} }
func (m *InitializeSiteResponse) String() string { return proto.CompactTextString(m) }
func (*InitializeSiteResponse) ProtoMessage()    {}
func (*InitializeSiteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e393e113b6828dc9, []int{1}
}

func (m *InitializeSiteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeSiteResponse.Unmarshal(m, b)
}
func (m *InitializeSiteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeSiteResponse.Marshal(b, m, deterministic)
}
func (m *InitializeSiteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeSiteResponse.Merge(m, src)
}
func (m *InitializeSiteResponse) XXX_Size() int {
	return xxx_messageInfo_InitializeSiteResponse.Size(m)
}
func (m *InitializeSiteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeSiteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeSiteResponse proto.InternalMessageInfo

func (m *InitializeSiteResponse) GetSiteId() string {
	if m != nil {
		return m.SiteId
	}
	return ""
}

func init() {
	proto.RegisterType((*InitializeSiteRequest)(nil), "fbuf.sitemanager.InitializeSiteRequest")
	proto.RegisterType((*InitializeSiteResponse)(nil), "fbuf.sitemanager.InitializeSiteResponse")
}

func init() { proto.RegisterFile("sitemanager/manager.proto", fileDescriptor_e393e113b6828dc9) }

var fileDescriptor_e393e113b6828dc9 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0xce, 0x2c, 0x49,
	0xcd, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x02, 0x69, 0x49, 0xa5, 0x69, 0x7a, 0x48, 0xf2, 0x52, 0xf2, 0xe9, 0xf9, 0xf9, 0xe9, 0x39,
	0xa9, 0xfa, 0x60, 0xf9, 0xa4, 0xd2, 0x34, 0xfd, 0x92, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc,
	0x02, 0x88, 0x16, 0x25, 0x71, 0x2e, 0x51, 0xcf, 0xbc, 0xcc, 0x92, 0xcc, 0xc4, 0x9c, 0xcc, 0xaa,
	0xd4, 0xe0, 0xcc, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x43, 0x2e, 0x31,
	0x74, 0x89, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x71, 0x2e, 0x76, 0x90, 0x15, 0xf1, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x51, 0x11, 0x17,
	0x37, 0x48, 0xa1, 0x2f, 0xc4, 0x6e, 0xa1, 0x64, 0x2e, 0x3e, 0x54, 0x13, 0x84, 0xd4, 0xf5, 0xd0,
	0x1d, 0xa8, 0x87, 0xd5, 0x72, 0x29, 0x0d, 0xc2, 0x0a, 0x21, 0x8e, 0x71, 0x32, 0x88, 0xd2, 0x4b,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0xcc, 0xd3,
	0x4d, 0xcb, 0x2f, 0xca, 0x2d, 0xcd, 0x49, 0xd4, 0x07, 0x99, 0xa1, 0x9b, 0x9e, 0xaf, 0x8f, 0x64,
	0x4c, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0xe3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x18,
	0x08, 0x13, 0x48, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SiteManagerClient is the client API for SiteManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SiteManagerClient interface {
	InitializeSite(ctx context.Context, in *InitializeSiteRequest, opts ...grpc.CallOption) (*InitializeSiteResponse, error)
}

type siteManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteManagerClient(cc grpc.ClientConnInterface) SiteManagerClient {
	return &siteManagerClient{cc}
}

func (c *siteManagerClient) InitializeSite(ctx context.Context, in *InitializeSiteRequest, opts ...grpc.CallOption) (*InitializeSiteResponse, error) {
	out := new(InitializeSiteResponse)
	err := c.cc.Invoke(ctx, "/fbuf.sitemanager.SiteManager/InitializeSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteManagerServer is the server API for SiteManager service.
type SiteManagerServer interface {
	InitializeSite(context.Context, *InitializeSiteRequest) (*InitializeSiteResponse, error)
}

// UnimplementedSiteManagerServer can be embedded to have forward compatible implementations.
type UnimplementedSiteManagerServer struct {
}

func (*UnimplementedSiteManagerServer) InitializeSite(ctx context.Context, req *InitializeSiteRequest) (*InitializeSiteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeSite not implemented")
}

func RegisterSiteManagerServer(s *grpc.Server, srv SiteManagerServer) {
	s.RegisterService(&_SiteManager_serviceDesc, srv)
}

func _SiteManager_InitializeSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteManagerServer).InitializeSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fbuf.sitemanager.SiteManager/InitializeSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteManagerServer).InitializeSite(ctx, req.(*InitializeSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SiteManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fbuf.sitemanager.SiteManager",
	HandlerType: (*SiteManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeSite",
			Handler:    _SiteManager_InitializeSite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sitemanager/manager.proto",
}
