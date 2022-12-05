gen-proto:
	#protoc --proto_path=proto --go_out=out example.proto


	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/example.proto

image:
	docker build -t grpc-example:latest .