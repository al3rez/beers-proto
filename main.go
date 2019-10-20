//go:generate go get -u github.com/golang/protobuf/protoc-gen-go
//go:generate protoc pkg/pb/beers.proto --go_out=plugins=grpc:.
package main

func main() {}
