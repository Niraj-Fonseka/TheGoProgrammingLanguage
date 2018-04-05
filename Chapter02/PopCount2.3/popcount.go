package main

import (
	"fmt"
)

//pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Println("Running PoPCount")
	count := PopCount(11)
	fmt.Println(count)
}

//PoPCount returns the population count ( number of set bits ) of x
func PopCount(x uint64) int {

	var result byte
	var index uint

	for index = 0; index < 8; index++ {
		result += pc[byte(x>>(index*8))]
	}

	return int(result)
}
