package main

import "fmt"

func main() {
	fmt.Println("test")
}

// oapi-codegen -package dto -generate types api/openapi.yml > internal/infrastructure/dto/dto-api.gen.go
// oapi-codegen -package dto -generate types,client,server api/openapi.yml > internal/infrastructure/dto/dto-api.gen.go
