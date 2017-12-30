.PHONY: test

bin/server: cmd/server/main.go server/*.go
	go build -ldflags "-X main.appVersion=`git rev-list HEAD -n1`" -o bin/server cmd/server/main.go

test:
	go test ./...
