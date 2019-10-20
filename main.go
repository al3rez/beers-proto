//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate protoc pkg/pb/beers.proto --go_out=plugins=grpc:.
//go:generate protoc pkg/pb/grpc_health_v1/health.proto --go_out=plugins=grpc:.
package main

func main() {}
