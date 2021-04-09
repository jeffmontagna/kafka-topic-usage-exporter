# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=kafka-topic-usage-exporter
BINARY_UNIX=$(BINARY_NAME)-linux-x64
VERSION=0.0.2

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v -ldflags "-X main.version=$(VERSION)"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v -ldflags "-X main.version=$(VERSION)"
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/spf13/viper
	$(GOGET) github.com/Shopify/sarama

# Cross compilation
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v -ldflags "-X main.version=$(VERSION)"
