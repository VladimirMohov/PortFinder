BIN_NAME=PortFinder
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=darwin go build -o ./app/${BIN_NAME}-darwin portFinder.go
	GOARCH=amd64 GOOS=linux go build -o ./app/${BIN_NAME}-linux portFinder.go
	GOARCH=amd64 GOOS=windows go build -o ./app/${BIN_NAME}-windows portFinder.go

run: build
	./app/${BIN_NAME}-linux

clean:
	go clean
	rm ./app/${BIN_NAME}-darwin
	rm ./app/${BIN_NAME}-linux
	rm ./app/${BIN_NAME}-windows