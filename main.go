package main

import "fmt"

type Person struct {
	name string
}

func main() {
	persons := []*Person{{name: "inas"},
		{name: "Fatur"},
		{name: "Bayu"},
		{name: "Mocha"},
		{name: "Frizky"},
		{name: "Hasanudin"},
		{name: "Yusuf"},
		{name: "Thalia"},
		{name: "Rizal"},
		{name: "Barru"}}

	printFriends := func(friends []*Person) {
		// do something
		for _, pr := range friends {
			fmt.Printf("Helo friend, my name is %s \n", pr.name)
		}
	}
	printFriends(persons)
}
