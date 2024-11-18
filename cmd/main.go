package main

import (
	grpc2 "github.com/RVodassa/geo-microservices-geo_service/internal/grpc-server"
	"github.com/RVodassa/geo-microservices-geo_service/internal/service"
	pb "github.com/RVodassa/geo-microservices-geo_service/proto/generated"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

func main() {
	// Загружаем переменные окружения из .env файла
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Println("Ошибка загрузки .env файла")
	}

	ApiKey, SecretKey := os.Getenv("apiKey"), os.Getenv("secretKey")

	geoService := service.NewGeoService(ApiKey, SecretKey)

	lis, err := net.Listen("tcp", ":30303")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGeoServiceServer(grpcServer, grpc2.NewServer(geoService))

	log.Println("gRPC server is running on port ")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
