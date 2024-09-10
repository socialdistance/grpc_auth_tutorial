generate:
		protoc -I protoss/proto protoss/proto/sso/sso.proto --go_out=protoss/gen/go --go_opt=paths=source_relative --go-grpc_out=protoss/gen/go/ --go-grpc_opt=paths=source_relative

run:
		go run sso/cmd/sso/main.go --config=sso/config/local.yaml