package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	var times int
	flag.IntVar(&times, "count", 1, "How many times to generate a name.")

	flag.Parse()

	for i := 0; i < times; i++ {
		name := genName()
		fmt.Println(name)
	}
}

func genName() string {
	var consonants []rune
	var vowels []rune

	for c := 'a'; c <= 'z'; c++ {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowels = append(vowels, c)
		default:
			consonants = append(consonants, c)
		}
	}

	name := make([]rune, 5)
	name[0] = consonants[rand.Intn(len(consonants))]
	name[1] = vowels[rand.Intn(len(vowels))]
	name[2] = consonants[rand.Intn(len(consonants))]
	name[3] = consonants[rand.Intn(len(consonants))]
	name[4] = vowels[rand.Intn(len(vowels))]

	return string(name)
}
