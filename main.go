package main

import (
	sampleV1 "awesomeProject/sample/gen/sample/v1"
	lr "awesomeProject/utils/logger"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	HOST = "localhost"
	PORT = 5000
)

type Server struct {
	sampleV1.UnimplementedSampleServiceServer
}

var (
	defaultServer *grpc.Server
	logger        *lr.Logger
)

func (s *Server) RegisterService(mainServer *grpc.Server) error {
	sampleV1.RegisterSampleServiceServer(mainServer, &Server{})
	fmt.Println("Register Success")
	fmt.Println(mainServer)

	return nil
}

func (s *Server) GetInfo(ctx context.Context, req *sampleV1.GetInfoInfoRequest) (*sampleV1.GetInfoResponse, error) {
	fmt.Println("===== GetInfo Start =====", ctx)

	requestBody := req.GetSendMessage()
	if requestBody != "" {
		fmt.Println("[Request] GetInfo Request Body:", requestBody)
	}

	responseData := &sampleV1.GetInfoResponse{
		ResponseMessage: HOST,
	}

	return responseData, nil
}

func main() {
	//서버 부팅 시 로그 출력.
	address := HOST + ":" + strconv.Itoa(PORT)

	lis, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}

	/*defaultServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpcMiddleware(),
			zeroLog.NewUnaryServerInterceptorWithLogger(logger.GetHandle()),
		)),
	)
	logger.Printf("server listening at %v", lis.Addr())
	if err := defaultServer.Serve(lis); err != nil {
		logger.Fatal().Msgf("failed to serve: %v", err)
	}*/

	grpcServer := grpc.NewServer()

	sampleService := Server{}
	sampleService.RegisterService(grpcServer)

	fmt.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}

}
