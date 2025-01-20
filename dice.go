package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Rolling the dice....")
	roll := rand.Intn(6) + 1
	fmt.Printf("You rolled number %d!\n", roll)
}
