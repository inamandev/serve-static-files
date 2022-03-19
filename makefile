all: lint test build

build:
	go build -o dist/static-file-webserver main.go

test:
	go test

lint:
	go vet