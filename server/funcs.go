package main

import (
	context "context"
	"server/protos"
)

type server struct {
}

func (s *server) Register(ctx context.Context, r *protos.RegisterRequest) (*protos.RegisterResponse, error) {

}

func (s *server) ServiceRegister(context.Context, *protos.ServiceRegisterRequest) (*protos.ServiceRegisterResponse, error) {

}

func (s *server) CheckHealth(context.Context, *protos.HeartBeat) (*protos.HeartBeatResponse, error) {

}

func (s *server) GetService(context.Context, *protos.GetServiceRequest) (*protos.GetServiceResponse, error) {

}
