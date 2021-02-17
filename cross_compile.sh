#!/bin/bash

env GOOS=darwin GOARCH=amd64 go build -o build/macos/amd64/serve main.go;
#env GOOS=darwin GOARCH=arm64 go build -o build/macos/arm64/serve main.go;
env GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/serve main.go;
env GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/serve main.go;
env GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/serve.exe main.go;
