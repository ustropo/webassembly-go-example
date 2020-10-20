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


