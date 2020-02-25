// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sitemanager/manager.proto

package sitemanagerpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Route struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Host                 string               `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Bucket               string               `protobuf:"bytes,5,opt,name=bucket,proto3" json:"bucket,omitempty"`
	BucketRegion         string               `protobuf:"bytes,6,opt,name=bucket_region,json=bucketRegion,proto3" json:"bucket_region,omitempty"`
	BasePath             string               `protobuf:"bytes,7,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_e393e113b6828dc9, []int{0}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Route) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Route) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Route) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Route) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Route) GetBucketRegion() string {
	if m != nil {
		return m.BucketRegion
	}
	return ""
}

func (m *Route) GetBasePath() string {
	if m != nil {
		return m.BasePath
	}
	return ""
}

type InitializeSiteRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeSiteRequest) Reset()         { *m = InitializeSiteRequest{} }
func (m *InitializeSiteRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeSiteRequest) ProtoMessage()    {}
func (*InitializeSiteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e393e113b6828dc9, []int{1}
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
	return fileDescriptor_e393e113b6828dc9, []int{2}
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
	proto.RegisterType((*Route)(nil), "fbuf.sitemanager.Route")
	proto.RegisterType((*InitializeSiteRequest)(nil), "fbuf.sitemanager.InitializeSiteRequest")
	proto.RegisterType((*InitializeSiteResponse)(nil), "fbuf.sitemanager.InitializeSiteResponse")
}

func init() { proto.RegisterFile("sitemanager/manager.proto", fileDescriptor_e393e113b6828dc9) }

var fileDescriptor_e393e113b6828dc9 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xbf, 0x4f, 0xfa, 0x40,
	0x14, 0x4f, 0xf9, 0x42, 0xf9, 0xf2, 0x50, 0x62, 0x2e, 0x11, 0x2a, 0x0e, 0x12, 0x1c, 0x64, 0xe1,
	0xaa, 0x38, 0x39, 0xe2, 0xc6, 0x60, 0x62, 0xaa, 0x93, 0x4b, 0x73, 0x6d, 0x5f, 0xdb, 0x8b, 0xb4,
	0x57, 0xdb, 0x77, 0x8b, 0xff, 0xba, 0x8b, 0xe9, 0xb5, 0x18, 0x24, 0x26, 0x4c, 0x7d, 0xef, 0xf3,
	0xab, 0x2f, 0x9f, 0x83, 0x8b, 0x4a, 0x12, 0x66, 0x22, 0x17, 0x09, 0x96, 0x6e, 0xfb, 0xe5, 0x45,
	0xa9, 0x48, 0xb1, 0xb3, 0x38, 0xd0, 0x31, 0xdf, 0xe3, 0xa7, 0x57, 0x89, 0x52, 0xc9, 0x16, 0x5d,
	0xc3, 0x07, 0x3a, 0x76, 0x49, 0x66, 0x58, 0x91, 0xc8, 0x8a, 0xc6, 0x32, 0xff, 0xb2, 0xa0, 0xe7,
	0x29, 0x4d, 0xc8, 0x46, 0xd0, 0x91, 0x91, 0x63, 0xcd, 0xac, 0xc5, 0xc0, 0xeb, 0xc8, 0x88, 0x3d,
	0x00, 0x84, 0x25, 0x0a, 0xc2, 0xc8, 0x17, 0xe4, 0x74, 0x66, 0xd6, 0x62, 0xb8, 0x9a, 0xf2, 0x26,
	0x8f, 0xef, 0xf2, 0xf8, 0xeb, 0x2e, 0xcf, 0x1b, 0xb4, 0xea, 0x35, 0xd5, 0x56, 0x5d, 0x44, 0x3b,
	0xeb, 0xbf, 0xe3, 0xd6, 0x56, 0xbd, 0x26, 0xc6, 0xa0, 0x9b, 0xaa, 0x8a, 0x9c, 0xae, 0xb9, 0xc3,
	0xcc, 0x6c, 0x0c, 0x76, 0xa0, 0xc3, 0x77, 0x24, 0xa7, 0x67, 0xd0, 0x76, 0x63, 0xd7, 0x70, 0xda,
	0x4c, 0x7e, 0x89, 0x89, 0x54, 0xb9, 0x63, 0x1b, 0xfa, 0xa4, 0x01, 0x3d, 0x83, 0xb1, 0x4b, 0x18,
	0x04, 0xa2, 0x42, 0xbf, 0x10, 0x94, 0x3a, 0x7d, 0x23, 0xf8, 0x5f, 0x03, 0xcf, 0x82, 0xd2, 0xf9,
	0x04, 0xce, 0x37, 0xb9, 0x24, 0x29, 0xb6, 0xf2, 0x13, 0x5f, 0x24, 0xa1, 0x87, 0x1f, 0x1a, 0x2b,
	0x9a, 0xdf, 0xc1, 0xf8, 0x90, 0xa8, 0x0a, 0x95, 0x57, 0xc8, 0x26, 0xd0, 0xaf, 0x0b, 0xf6, 0x7f,
	0xba, 0xb2, 0xeb, 0x75, 0x13, 0xad, 0x4a, 0x18, 0xd6, 0xc2, 0xa7, 0xa6, 0x79, 0x16, 0xc2, 0xe8,
	0x77, 0x02, 0xbb, 0xe1, 0x87, 0xcf, 0xc3, 0xff, 0xfc, 0xf9, 0x74, 0x71, 0x5c, 0xd8, 0x1c, 0xf3,
	0x78, 0xfb, 0xc6, 0x13, 0x49, 0xa9, 0x0e, 0x78, 0xa8, 0x32, 0x97, 0x4a, 0x21, 0xf3, 0x65, 0xac,
	0xca, 0x4c, 0x6f, 0x85, 0x5b, 0x67, 0x2c, 0x13, 0xe5, 0xee, 0xc5, 0x14, 0x41, 0x60, 0x9b, 0xfa,
	0xef, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x20, 0x7d, 0x36, 0x20, 0x46, 0x02, 0x00, 0x00,
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
