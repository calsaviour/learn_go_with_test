GO=go

all: hello

hello: hello.go
	$(GO) run hello.go

test:
	$(GO) test

