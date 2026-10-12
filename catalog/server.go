package catalog

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/DICKENSOKOTH/GO-micro/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{service: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostProduct(ctx context.Context, r *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	p, err := s.service.PostProduct(ctx, r.name, r.Description, r.Price)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.PostProductResponse{Product: &pb.Product{
		id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}}, nil
}

func (s *grpcServer) GetProduct(ctx context.Context, r *pb.GetproductRequest) *pb.GetProductresponse {
	p, err := s.service.GetProduct(ctx, r.id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.GetProductresponse{
		Product: &pb.product{
			Id:         p.ID,
			Name:       p.Name,
			Desription: p.Description,
			Price:      p.Price,
		},
	}, nil
}

func (s *grpcServer) GetProducts(ctx context.Context, r *pb.GetproductsRequest) (*pb.GetProductsresponse, error) {
	var res []product
	var err error
	if r.Query != "" {
		res, err = s.service.SearchProducts(ctx, r.Query, r.skip, r.take)
	} else if len(r.Ids) != 0 {
		res, err = s.service.GetProductByIDs(ctx, r.Ids)
	} else {
		res, err = s.service.GetProducts(ctx, r.skip, r.take)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []*pb.Product{}
	for _, p := range res {
		products = append(
			products,
			&pb.Products{
				Id:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			},
		)
	}
	return &pb.GetProductsResponse{Products: products}, nil
}
