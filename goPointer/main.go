package main

import (
	"fmt"
)

type Ninja struct {
	Name  string
	Power int
}

func main() {
	// we check the pass by copy across function and its impact on values

	Leaf := &Ninja{"Kakashi", 50000}
	//Leaf.Power = Super(Leaf)

	Super(Leaf)
	fmt.Println(Leaf.Power)
}

/*
	func Super(n Ninja) int {
		n.Power += 10000
		return n.Power
	}
*/
func Super(n *Ninja) {
	n.Power += 10000
}
