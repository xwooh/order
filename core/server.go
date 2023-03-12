package core

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"order/global/logger"
	pb "order/pb/proto"
	"order/service"
)

type ServiceRegistry struct {
	Register func(grpc.ServiceRegistrar, interface{})
	Service  interface{}
}

func RegisterService(grpcServer *grpc.Server) {
	// 注册 Service
	pb.RegisterGreeterServiceServer(grpcServer, &service.GreeterService{})
	pb.RegisterOrderServiceServer(grpcServer, &service.OrderService{})
}

func Serve(host string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		logger.Error(err)
	}

	grpcServer := grpc.NewServer()

	RegisterService(grpcServer)

	logger.Infof("grpc server is running at %s:%d ...", host, port)

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve:%v", err)
	}
}
