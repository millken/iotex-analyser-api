// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// StakingServiceClient is the client API for StakingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StakingServiceClient interface {
	GetVoteByHeight(ctx context.Context, in *StakingRequest, opts ...grpc.CallOption) (*StakingResponse, error)
	GetCandidateVoteByHeight(ctx context.Context, in *StakingRequest, opts ...grpc.CallOption) (*StakingResponse, error)
}

type stakingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStakingServiceClient(cc grpc.ClientConnInterface) StakingServiceClient {
	return &stakingServiceClient{cc}
}

func (c *stakingServiceClient) GetVoteByHeight(ctx context.Context, in *StakingRequest, opts ...grpc.CallOption) (*StakingResponse, error) {
	out := new(StakingResponse)
	err := c.cc.Invoke(ctx, "/api.StakingService/GetVoteByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stakingServiceClient) GetCandidateVoteByHeight(ctx context.Context, in *StakingRequest, opts ...grpc.CallOption) (*StakingResponse, error) {
	out := new(StakingResponse)
	err := c.cc.Invoke(ctx, "/api.StakingService/GetCandidateVoteByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakingServiceServer is the server API for StakingService service.
// All implementations must embed UnimplementedStakingServiceServer
// for forward compatibility
type StakingServiceServer interface {
	GetVoteByHeight(context.Context, *StakingRequest) (*StakingResponse, error)
	GetCandidateVoteByHeight(context.Context, *StakingRequest) (*StakingResponse, error)
	mustEmbedUnimplementedStakingServiceServer()
}

// UnimplementedStakingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStakingServiceServer struct {
}

func (UnimplementedStakingServiceServer) GetVoteByHeight(context.Context, *StakingRequest) (*StakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoteByHeight not implemented")
}
func (UnimplementedStakingServiceServer) GetCandidateVoteByHeight(context.Context, *StakingRequest) (*StakingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandidateVoteByHeight not implemented")
}
func (UnimplementedStakingServiceServer) mustEmbedUnimplementedStakingServiceServer() {}

// UnsafeStakingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StakingServiceServer will
// result in compilation errors.
type UnsafeStakingServiceServer interface {
	mustEmbedUnimplementedStakingServiceServer()
}

func RegisterStakingServiceServer(s grpc.ServiceRegistrar, srv StakingServiceServer) {
	s.RegisterService(&StakingService_ServiceDesc, srv)
}

func _StakingService_GetVoteByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingServiceServer).GetVoteByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StakingService/GetVoteByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingServiceServer).GetVoteByHeight(ctx, req.(*StakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StakingService_GetCandidateVoteByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StakingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakingServiceServer).GetCandidateVoteByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.StakingService/GetCandidateVoteByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakingServiceServer).GetCandidateVoteByHeight(ctx, req.(*StakingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StakingService_ServiceDesc is the grpc.ServiceDesc for StakingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StakingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StakingService",
	HandlerType: (*StakingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVoteByHeight",
			Handler:    _StakingService_GetVoteByHeight_Handler,
		},
		{
			MethodName: "GetCandidateVoteByHeight",
			Handler:    _StakingService_GetCandidateVoteByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_staking.proto",
}
