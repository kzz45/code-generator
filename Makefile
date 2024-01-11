GOBIN := ""

client-gen:
	go build -o client-gen ./cmd/client-gen/main.go
	# cp client-gen ${GOBIN} && rm client-gen
