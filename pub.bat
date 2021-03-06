go clean -cache -modcache -i -r
set GOARCH=wasm
set GOOS=js
go build -o urlv4.wasm