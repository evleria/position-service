// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	OpenPosition(ctx context.Context, in *OpenPositionRequest, opts ...grpc.CallOption) (*OpenPositionResponse, error)
	ClosePosition(ctx context.Context, in *ClosePositionRequest, opts ...grpc.CallOption) (*ClosePositionResponse, error)
	GetOpenPosition(ctx context.Context, in *GetOpenPositionRequest, opts ...grpc.CallOption) (*GetOpenPositionResponse, error)
	SetStopLoss(ctx context.Context, in *SetStopLossRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) OpenPosition(ctx context.Context, in *OpenPositionRequest, opts ...grpc.CallOption) (*OpenPositionResponse, error) {
	out := new(OpenPositionResponse)
	err := c.cc.Invoke(ctx, "/pb.PositionService/OpenPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) ClosePosition(ctx context.Context, in *ClosePositionRequest, opts ...grpc.CallOption) (*ClosePositionResponse, error) {
	out := new(ClosePositionResponse)
	err := c.cc.Invoke(ctx, "/pb.PositionService/ClosePosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetOpenPosition(ctx context.Context, in *GetOpenPositionRequest, opts ...grpc.CallOption) (*GetOpenPositionResponse, error) {
	out := new(GetOpenPositionResponse)
	err := c.cc.Invoke(ctx, "/pb.PositionService/GetOpenPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) SetStopLoss(ctx context.Context, in *SetStopLossRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.PositionService/SetStopLoss", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	OpenPosition(context.Context, *OpenPositionRequest) (*OpenPositionResponse, error)
	ClosePosition(context.Context, *ClosePositionRequest) (*ClosePositionResponse, error)
	GetOpenPosition(context.Context, *GetOpenPositionRequest) (*GetOpenPositionResponse, error)
	SetStopLoss(context.Context, *SetStopLossRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) OpenPosition(context.Context, *OpenPositionRequest) (*OpenPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenPosition not implemented")
}
func (UnimplementedPositionServiceServer) ClosePosition(context.Context, *ClosePositionRequest) (*ClosePositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePosition not implemented")
}
func (UnimplementedPositionServiceServer) GetOpenPosition(context.Context, *GetOpenPositionRequest) (*GetOpenPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpenPosition not implemented")
}
func (UnimplementedPositionServiceServer) SetStopLoss(context.Context, *SetStopLossRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStopLoss not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_OpenPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).OpenPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PositionService/OpenPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).OpenPosition(ctx, req.(*OpenPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_ClosePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).ClosePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PositionService/ClosePosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).ClosePosition(ctx, req.(*ClosePositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetOpenPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOpenPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetOpenPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PositionService/GetOpenPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetOpenPosition(ctx, req.(*GetOpenPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_SetStopLoss_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStopLossRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).SetStopLoss(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PositionService/SetStopLoss",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).SetStopLoss(ctx, req.(*SetStopLossRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenPosition",
			Handler:    _PositionService_OpenPosition_Handler,
		},
		{
			MethodName: "ClosePosition",
			Handler:    _PositionService_ClosePosition_Handler,
		},
		{
			MethodName: "GetOpenPosition",
			Handler:    _PositionService_GetOpenPosition_Handler,
		},
		{
			MethodName: "SetStopLoss",
			Handler:    _PositionService_SetStopLoss_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "position_service.proto",
}
