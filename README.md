# WebAssembly in Go - Example

Example code on how to create a WebAssembly code using Go.


## Compile 

To compile, just run the command:

```
GOOS=js GOARCH=wasm go build -o assets/main.wasm cmd/wasm/main.go
```

## Run

To run:

```
go run cmd/server/main.go
```

Open the browser with the address `localhost:8000`.


## Article

This code was done for the article ["How to Write a WebAssembly App in Go"](https://medium.com/vacatronics/how-to-write-a-webassembly-app-in-go-fd769fa2b64b?source=friends_link&sk=3ead83bf637ee79131dc8f7d787c3a3d)
