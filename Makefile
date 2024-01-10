export

.PHONY: build

build:
	@echo "Building plugin..."
	@go build -buildmode=plugin -o build/plug.so -ldflags="-s -w"  plugin/*
	@echo "Building app..."
	@go build -o build/app -ldflags="-s -w"  app/*
