package main

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/AjaySand/gRPC-calculator/pb"
)

type Server struct {
	pb.UnimplementedCalculatorServer
}

func (s *Server) Add(context context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.X + in.Y,
	}, nil
}

func (s *Server) Subtract(context context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.X - in.Y,
	}, nil
}

func (s *Server) Multiply(context context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	return &pb.CalculatorResponse{
		Result: in.X * in.Y,
	}, nil
}

func (s *Server) Divide(context context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	if in.Y == 0 {
		return nil, errors.New("Cannot divide by zero")
	}

	return &pb.CalculatorResponse{
		Result: in.X / in.Y,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	pb.RegisterCalculatorServer(server, &Server{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
