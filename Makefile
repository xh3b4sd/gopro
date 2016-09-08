.PHONY: all clean gogenerate goget gopro



GOPATH := ${PWD}/.workspace
PATH := ${PATH}:${PWD}/.workspace/bin:${PWD}/bin
export GOPATH
export PATH



VERSION := $(shell git rev-parse --short HEAD)



all: goget gogenerate gopro

clean:
	@rm -rf .workspace/

gogenerate:
	@go generate ./...

goget:
	@# Setup workspace.
	@mkdir -p ${PWD}/.workspace/src/github.com/xh3b4sd/
	@ln -fs ${PWD} ${PWD}/.workspace/src/github.com/xh3b4sd/
	@# Fetch build dependencies.
	@go get github.com/xh3b4sd/loader
	@# Fetch the rest of the project dependencies.
	@go get -d -v ./...

gorun:
	@go run \
		-ldflags "-X main.version=${VERSION}" \
		main.go



gopro:
	@go build \
		-o .workspace/bin/gopro \
		-ldflags "-X main.version=${VERSION}" \
		github.com/xh3b4sd/gopro
