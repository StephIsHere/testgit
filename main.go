package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 3)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int()
		fmt.Println(arr[i])
	}

}
