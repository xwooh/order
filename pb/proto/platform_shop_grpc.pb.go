// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlatformShopServiceClient is the client API for PlatformShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlatformShopServiceClient interface {
	// 通过 id 获取平台门店详细信息
	GetPlatformShopById(ctx context.Context, in *GetPlatformShopByIdRequest, opts ...grpc.CallOption) (*GetPlatformShopByIdResponse, error)
}

type platformShopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlatformShopServiceClient(cc grpc.ClientConnInterface) PlatformShopServiceClient {
	return &platformShopServiceClient{cc}
}

func (c *platformShopServiceClient) GetPlatformShopById(ctx context.Context, in *GetPlatformShopByIdRequest, opts ...grpc.CallOption) (*GetPlatformShopByIdResponse, error) {
	out := new(GetPlatformShopByIdResponse)
	err := c.cc.Invoke(ctx, "/order.platform_shop.PlatformShopService/GetPlatformShopById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlatformShopServiceServer is the server API for PlatformShopService service.
// All implementations should embed UnimplementedPlatformShopServiceServer
// for forward compatibility
type PlatformShopServiceServer interface {
	// 通过 id 获取平台门店详细信息
	GetPlatformShopById(context.Context, *GetPlatformShopByIdRequest) (*GetPlatformShopByIdResponse, error)
}

// UnimplementedPlatformShopServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPlatformShopServiceServer struct {
}

func (UnimplementedPlatformShopServiceServer) GetPlatformShopById(context.Context, *GetPlatformShopByIdRequest) (*GetPlatformShopByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlatformShopById not implemented")
}

// UnsafePlatformShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlatformShopServiceServer will
// result in compilation errors.
type UnsafePlatformShopServiceServer interface {
	mustEmbedUnimplementedPlatformShopServiceServer()
}

func RegisterPlatformShopServiceServer(s grpc.ServiceRegistrar, srv PlatformShopServiceServer) {
	s.RegisterService(&PlatformShopService_ServiceDesc, srv)
}

func _PlatformShopService_GetPlatformShopById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlatformShopByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformShopServiceServer).GetPlatformShopById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.platform_shop.PlatformShopService/GetPlatformShopById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformShopServiceServer).GetPlatformShopById(ctx, req.(*GetPlatformShopByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlatformShopService_ServiceDesc is the grpc.ServiceDesc for PlatformShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlatformShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.platform_shop.PlatformShopService",
	HandlerType: (*PlatformShopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlatformShopById",
			Handler:    _PlatformShopService_GetPlatformShopById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platform_shop.proto",
}
