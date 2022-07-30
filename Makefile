.PHONY: build check help

help:
	@echo "use [make check]: code review"
	@echo "use [make build]: build project"
build:
	@go build -o product-srv main.go
check:
	@golangci-lint run