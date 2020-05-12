package main

import (
	"fmt"
	"github.com/lucaschimweg/dhbw-go-portfolio/fibonacci"
	"github.com/lucaschimweg/dhbw-go-portfolio/zoo"
)

func main() {
	/*dog := zoo.Dog{zoo.Animal{Name: "Bello"}}
	ForceToSayName(&dog)

	parrot := zoo.Parrot{zoo.Animal{Name: "Carl"}}
	ForceToSayName(&parrot)*/

	fmt.Println(fibonacci.Fib(25))
}

func ForceToSayName(speaker zoo.Speaker) {
	speaker.SayYourName()
}
