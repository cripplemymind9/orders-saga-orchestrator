// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: orders-saga-orchestrator.proto

package api

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
	OrdersSagaOrchestratorService_GetOrderSaga_FullMethodName = "/api.v1.orderssagaorchestrator.OrdersSagaOrchestratorService/GetOrderSaga"
)

// OrdersSagaOrchestratorServiceClient is the client API for OrdersSagaOrchestratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления сагами заказов
type OrdersSagaOrchestratorServiceClient interface {
	// Получение саги по идентификатору заказа
	GetOrderSaga(ctx context.Context, in *GetOrderSagaRequest, opts ...grpc.CallOption) (*GetOrderSagaResponse, error)
}

type ordersSagaOrchestratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersSagaOrchestratorServiceClient(cc grpc.ClientConnInterface) OrdersSagaOrchestratorServiceClient {
	return &ordersSagaOrchestratorServiceClient{cc}
}

func (c *ordersSagaOrchestratorServiceClient) GetOrderSaga(ctx context.Context, in *GetOrderSagaRequest, opts ...grpc.CallOption) (*GetOrderSagaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderSagaResponse)
	err := c.cc.Invoke(ctx, OrdersSagaOrchestratorService_GetOrderSaga_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersSagaOrchestratorServiceServer is the server API for OrdersSagaOrchestratorService service.
// All implementations must embed UnimplementedOrdersSagaOrchestratorServiceServer
// for forward compatibility.
//
// Сервис для управления сагами заказов
type OrdersSagaOrchestratorServiceServer interface {
	// Получение саги по идентификатору заказа
	GetOrderSaga(context.Context, *GetOrderSagaRequest) (*GetOrderSagaResponse, error)
	mustEmbedUnimplementedOrdersSagaOrchestratorServiceServer()
}

// UnimplementedOrdersSagaOrchestratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrdersSagaOrchestratorServiceServer struct{}

func (UnimplementedOrdersSagaOrchestratorServiceServer) GetOrderSaga(context.Context, *GetOrderSagaRequest) (*GetOrderSagaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderSaga not implemented")
}
func (UnimplementedOrdersSagaOrchestratorServiceServer) mustEmbedUnimplementedOrdersSagaOrchestratorServiceServer() {
}
func (UnimplementedOrdersSagaOrchestratorServiceServer) testEmbeddedByValue() {}

// UnsafeOrdersSagaOrchestratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersSagaOrchestratorServiceServer will
// result in compilation errors.
type UnsafeOrdersSagaOrchestratorServiceServer interface {
	mustEmbedUnimplementedOrdersSagaOrchestratorServiceServer()
}

func RegisterOrdersSagaOrchestratorServiceServer(s grpc.ServiceRegistrar, srv OrdersSagaOrchestratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrdersSagaOrchestratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrdersSagaOrchestratorService_ServiceDesc, srv)
}

func _OrdersSagaOrchestratorService_GetOrderSaga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderSagaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersSagaOrchestratorServiceServer).GetOrderSaga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersSagaOrchestratorService_GetOrderSaga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersSagaOrchestratorServiceServer).GetOrderSaga(ctx, req.(*GetOrderSagaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersSagaOrchestratorService_ServiceDesc is the grpc.ServiceDesc for OrdersSagaOrchestratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersSagaOrchestratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.orderssagaorchestrator.OrdersSagaOrchestratorService",
	HandlerType: (*OrdersSagaOrchestratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderSaga",
			Handler:    _OrdersSagaOrchestratorService_GetOrderSaga_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders-saga-orchestrator.proto",
}
