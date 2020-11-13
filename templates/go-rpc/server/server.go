package {{.ServiceName}}

import (
	"context"

	"go.uber.org/zap"
	pb "github.com/frontedxyz/synthwave/services/{{.ServiceName}}/proto"
)

// Server .
type Server struct {
	Logger *zap.Logger
}

// Ping .
func (s *Server) Ping(ctx context.Context, req *pb.PingRequest) (rsp *pb.PingResponse, err error) {
	return &pb.PingResponse{
		Ok: true,
	}, nil
}
