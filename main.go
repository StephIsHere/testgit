package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Println(rand.NormFloat64() * 100)
	}
}
