package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for i, arg := range os.Args[0:] {
		s += sep + arg
		fmt.Printf("Index : %d , Value : %s \n", i, arg)
		sep = " "
	}
	fmt.Println(s)
}
