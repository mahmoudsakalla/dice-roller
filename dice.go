//Mahmoud Sakalla - Very short simple dice roller

package main

import (
	"fmt"
	"math/rand"
)

// program will roll a single six sided dice
func main() {

	fmt.Println("Rolling the dice....")
	roll := rand.Intn(6) + 1
	fmt.Printf("You rolled number %d\n", roll)
}
