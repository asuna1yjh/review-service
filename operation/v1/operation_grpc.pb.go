// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.1
// source: operation/v1/operation.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Operation_AuditReview_FullMethodName = "/api.operation.v1.Operation/AuditReview"
	Operation_AuditAppeal_FullMethodName = "/api.operation.v1.Operation/AuditAppeal"
)

// OperationClient is the client API for Operation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationClient interface {
	// 审核评论
	AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...grpc.CallOption) (*AuditReviewReply, error)
	// 审核申诉
	AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...grpc.CallOption) (*AuditAppealReply, error)
}

type operationClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationClient(cc grpc.ClientConnInterface) OperationClient {
	return &operationClient{cc}
}

func (c *operationClient) AuditReview(ctx context.Context, in *AuditReviewRequest, opts ...grpc.CallOption) (*AuditReviewReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuditReviewReply)
	err := c.cc.Invoke(ctx, Operation_AuditReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationClient) AuditAppeal(ctx context.Context, in *AuditAppealRequest, opts ...grpc.CallOption) (*AuditAppealReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuditAppealReply)
	err := c.cc.Invoke(ctx, Operation_AuditAppeal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServer is the server API for Operation service.
// All implementations must embed UnimplementedOperationServer
// for forward compatibility.
type OperationServer interface {
	// 审核评论
	AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error)
	// 审核申诉
	AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error)
	mustEmbedUnimplementedOperationServer()
}

// UnimplementedOperationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOperationServer struct{}

func (UnimplementedOperationServer) AuditReview(context.Context, *AuditReviewRequest) (*AuditReviewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditReview not implemented")
}
func (UnimplementedOperationServer) AuditAppeal(context.Context, *AuditAppealRequest) (*AuditAppealReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuditAppeal not implemented")
}
func (UnimplementedOperationServer) mustEmbedUnimplementedOperationServer() {}
func (UnimplementedOperationServer) testEmbeddedByValue()                   {}

// UnsafeOperationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationServer will
// result in compilation errors.
type UnsafeOperationServer interface {
	mustEmbedUnimplementedOperationServer()
}

func RegisterOperationServer(s grpc.ServiceRegistrar, srv OperationServer) {
	// If the following call pancis, it indicates UnimplementedOperationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Operation_ServiceDesc, srv)
}

func _Operation_AuditReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).AuditReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operation_AuditReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).AuditReview(ctx, req.(*AuditReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operation_AuditAppeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuditAppealRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServer).AuditAppeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Operation_AuditAppeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServer).AuditAppeal(ctx, req.(*AuditAppealRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Operation_ServiceDesc is the grpc.ServiceDesc for Operation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Operation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.operation.v1.Operation",
	HandlerType: (*OperationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuditReview",
			Handler:    _Operation_AuditReview_Handler,
		},
		{
			MethodName: "AuditAppeal",
			Handler:    _Operation_AuditAppeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operation/v1/operation.proto",
}
