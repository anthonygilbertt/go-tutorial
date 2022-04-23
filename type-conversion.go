package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	//what is the v for?
	//what is the T for?
	fmt.Printf("%v, %T\n", i, i)

	var j string
	// j = float32(i)
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
