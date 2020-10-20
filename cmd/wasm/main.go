package main

import (
	"fmt"
	"math"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan bool)

	fmt.Println("Hello from WebAssembly")
	js.Global().Set("square", js.FuncOf(jsSquare))

	<-c
}

func square(input float64) float64 {
	return math.Pow(input, 2)
}

func jsSquare(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid no of args passed"
	}

	input, err := strconv.ParseFloat(args[0].String(), 64)

	jsDoc := js.Global().Get("document")
	outElement := jsDoc.Call("getElementById", "out")

	fmt.Println("Input: ", input)
	if err != nil {
		return err.Error()
	}

	square := square(input)
	outElement.Set("textContent", square)
	return square
}
