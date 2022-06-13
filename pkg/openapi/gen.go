// Package api contains the REST interfaces.
package api

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -generate client,types,server,spec -package api -o api.gen.go api.yml
//go:generate gofmt -s -w api.gen.go


// ~/go/bin/oapi-codegen api.yml > api.gen.go