package main

import (
	"fmt"

	"github.com/titoforns/golang/variables"
)

func main() {
	estado, text := variables.ConviertoaTexto(34)
	fmt.Println(estado)
	fmt.Println(text)
}
