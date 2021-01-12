BINARY=metadata

build:
	go build -v -o bin/${BINARY}

build_64linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/${BINARY}

test:
	go test -v ./... -cover

