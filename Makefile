ifneq ("$(wildcard .env)","")
  $(info using .env)
  include .env
  export $(shell sed 's/=.*//' .env)
endif

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@go run main.go

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building...'	
	go build main.go

## clean/apps: clear generated bin files
.PHONY: clean/apps
clean/apps:
	@echo 'Remove builded apps'
	@rm -rf ./bin

## docker/up: start the local stack in background
.PHONY: docker/up
docker/up:
	docker-compose up -d 

## docker/down: shutdown the running containers
.PHONY: docker/down
docker/down:
	docker-compose down	

## audit: tidy dependencies, format and vet all code
.PHONY: audit
audit:
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	golangci-lint run

## tidy: tidy dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	golangci-lint run --fix