package zoo

import "fmt"

type Speaker interface {
	SayYourName()
}

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}

func (dog *Dog) SayYourName() {
	fmt.Println("I cannot say my name!")
}

type Parrot struct {
	Animal
}

func (parrot *Parrot) SayYourName() {
	fmt.Println("Hello, my name is " + parrot.Name)
}
