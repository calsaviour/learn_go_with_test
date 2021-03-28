GO=go
test:
	$(GO) test ./hello/**
	$(GO) test ./integers/**
	$(GO) test ./iteration/**

