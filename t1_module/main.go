package main

import (
	"fmt"

	"github.com/Babsztyl/puppy"
)

func main() {

	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.Bark() + "\t" + puppy.Barks())

}
