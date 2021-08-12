// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// BiddingServiceClient is the client API for BiddingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BiddingServiceClient interface {
	AddPosition(ctx context.Context, in *AddPositionRequest, opts ...grpc.CallOption) (*AddPositionResponse, error)
}

type biddingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBiddingServiceClient(cc grpc.ClientConnInterface) BiddingServiceClient {
	return &biddingServiceClient{cc}
}

func (c *biddingServiceClient) AddPosition(ctx context.Context, in *AddPositionRequest, opts ...grpc.CallOption) (*AddPositionResponse, error) {
	out := new(AddPositionResponse)
	err := c.cc.Invoke(ctx, "/BiddingService/AddPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BiddingServiceServer is the server API for BiddingService service.
// All implementations must embed UnimplementedBiddingServiceServer
// for forward compatibility
type BiddingServiceServer interface {
	AddPosition(context.Context, *AddPositionRequest) (*AddPositionResponse, error)
	mustEmbedUnimplementedBiddingServiceServer()
}

// UnimplementedBiddingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBiddingServiceServer struct {
}

func (UnimplementedBiddingServiceServer) AddPosition(context.Context, *AddPositionRequest) (*AddPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPosition not implemented")
}
func (UnimplementedBiddingServiceServer) mustEmbedUnimplementedBiddingServiceServer() {}

// UnsafeBiddingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BiddingServiceServer will
// result in compilation errors.
type UnsafeBiddingServiceServer interface {
	mustEmbedUnimplementedBiddingServiceServer()
}

func RegisterBiddingServiceServer(s grpc.ServiceRegistrar, srv BiddingServiceServer) {
	s.RegisterService(&BiddingService_ServiceDesc, srv)
}

func _BiddingService_AddPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BiddingServiceServer).AddPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BiddingService/AddPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BiddingServiceServer).AddPosition(ctx, req.(*AddPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BiddingService_ServiceDesc is the grpc.ServiceDesc for BiddingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BiddingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BiddingService",
	HandlerType: (*BiddingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPosition",
			Handler:    _BiddingService_AddPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "price_service.proto",
}
