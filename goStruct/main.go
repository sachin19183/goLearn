package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println(" The usage is to provide an integer value as an argument")
		os.Exit(1)
	}
	inPower, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	rank := 4
	//we explore different ways of defining structs in golang
	type Ninja struct {
		Name  string
		Power int
	}
	// first way of declaring a struct variable
	Leaf := Ninja{Name: "Naruto"}
	Leaf.Power = 1000000
	// second way of declaring a struct variable
	Sand := Ninja{
		Name:  "Gaara",
		Power: 100000,
	}
	// third way of declaring a struct variable
	Cloud := Ninja{}
	Cloud.Name = "Octopops"
	Cloud.Power = 150000
	// fourth way of declaring a struct variable
	Rain := Ninja{"Nagako", 2200000}

	fmt.Println(" We now comapre powers of Ninja from the three hidden villages:LEAF, CLOUD & SAND")
	if inPower > Sand.Power {
		rank = 3
		//fmt.Println(" The input power is more than that of Gara from the SAND Village")
		if inPower > Cloud.Power {
			rank = 2
			//fmt.Println(" The input power is more than that of Gara from the SAND Village")
		}
		if inPower > Leaf.Power {
			rank = 1
		}
		if inPower > Rain.Power {
			rank = 0
		}
	}
	switch rank {
	case 4:
		fmt.Println(" The input power is not comaprable to the power of ninja's from the three hidden villages")
	case 3:
		fmt.Println(" The input power is capable of defeating Gara of the sand village only")
	case 2:
		fmt.Println(" The input power is capable of defeating Gara of the sand village & Octopops from the Clud Village (one at a time)")
	case 1:
		fmt.Println("The input power is capable of defeating Gara of the sand village, Octopops form the cloud village and Naruto from the Leaf village. Wo owns this much power???")
	case 0:
		fmt.Println(" You understand Pain")
	default:
		fmt.Println(" The input is not valid")
	}
	/*	if Cloud.Power > Sand.Power {
			if Cloud.Power > Leaf.Power {
				fmt.Printf(" Clound Ninja %s is the most powerful ninja", Cloud.Name)
			} else {
				fmt.Printf(" Leaf Ninja %s is the most powerful ninja", Leaf.Name)
			}
		} else {
			if Sand.Power > Leaf.Power {
				fmt.Printf(" Sand Ninja %s is the most powerful ninja", Sand.Name)
			} else {
				fmt.Printf(" Leaf Ninja %s is the most powerful ninja", Leaf.Name)
			}
		}
	*/

}
