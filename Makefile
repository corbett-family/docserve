all: dist

ui/node_modules:
	cd ui; yarn install

ui/public: ui/node_modules
	cd ui; yarn build

clean:
	rm -rf ui/public
	rm -rf dist

dist: dist/linux-amd64/docserve dist/mac-arm64/docserve

dist/linux-amd64/docserve: ui/public main.go
	mkdir -p dist/linux-amd64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o dist/linux-amd64/docserve

dist/mac-arm64/docserve: ui/public main.go
	mkdir -p dist/mac-arm64
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -a -o dist/mac-arm64/docserve
