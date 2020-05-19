# Tutorial

# Theoretical Part
The slides for the presentation can be found [here](./slides.pdf)

# Practical Part
To follow this tutorial, you will have to have go installed on your system and have the `go` executable in your path. A good tutorial on installing go can be found [here](https://golang.org/doc/install)

## Example I: Zoo
The first example is a simple zoo, which showcases polymorphism in go. For this we will first create a type `Animal` in it's own package (in this case `zoo`):

```go
// zoo/animals.go

package zoo

type Animal struct {
    Name string
}

```

As you can see, the animal has a single string field `Name`, because in this example *all* animals will have a name.

Next we will create a type for an actual animal, a dog:

```go
// zoo/animals.go

type Dog struct {
    Animal
}

```

The `Dog` type contains an `Animal` instance as a field without a name. This makes all of `Animal`'s properties and functions accessible on `Dog`, in this case only the `Name`.

The next step is to create a function for the dog to "say it's name" (which a dog can obviously not do):

```go
// zoo/animals.go

import "fmt"

func (dog *Dog) SayYourName() {
    fmt.Println("I cannot say my name!")
}

```

After this is done we can test it from our main file (which we still have to create):

```go
// main.go

package main

import (
	"github.com/lucaschimweg/dhbw-go-portfolio/zoo"
)

func main() {
	dog := zoo.Dog{zoo.Animal{Name: "Bello"}}
	dog.SayYourName()
}


```

When we run this by executing `go run main.go` on your command line, you should see the output `"I cannot say my name!"`.


The next part is to include a second animal type, `Pattot`, and allow it to say its name:

```go
// zoo/animals.go

type Parrot struct {
	Animal
}

func (parrot *Parrot) SayYourName() {
	fmt.Println("Hello, my name is " + parrot.Name)
}

```


Because both `Dog` and `Parrot` can "say" their names, we want to have a common type for referencing them. This can be accomplished with interfaces in go. We will create an interface `Speaker` in the package `zoo`:

```go
// zoo/animals.go

type Speaker interface {
	SayYourName()
}

```

Go types implement interfaces implicitly. That means both `Dog` and `Parrot` automatically implement `Speaker` without any code changes to them! 

Now we are going to use this interface to create an animal-friendly function `ForceToSayName` in our main package:

```go
// main.go

func ForceToSayName(speaker zoo.Speaker) {
	speaker.SayYourName()
}

```

This function will accept any `Speaker` instance and will "force" it to say it's name. We update our main function to use this:

```go
func main() {
	dog := zoo.Dog{zoo.Animal{Name: "Bello"}}
	ForceToSayName(&dog)

	parrot := zoo.Parrot{zoo.Animal{Name: "Carl"}}
	ForceToSayName(&parrot)
}
```

Running this should result in this output:
```
I cannot say my name!
Hello, my name is Carl
```

If you got this output, you have completed the first tutorial! You can find the full code for `zoo/animals.go` [here](./zoo/animals.go)


## Tutorial II: Fibonacci
In this tutorial we will use go's goroutines to multithread the calculation of fibonacci numbers. We will use the much slower recursive approach for this. You would obviously never do this in any time critical use-case but it is great for demonstration purposes, so we are going to use it here.

First of all, we create a package `fibonacci` and add a traditional fibonacci function into it. I'll assume that you know what fibonacci is and how it works:

```go
// fibonacci/fibonacci.go

package fibonacci

func Fib(n int64) int64 {
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    return Fib(n-1) + Fib(n-2)
}

```

So far, so good. To be able to use this function in parallel, we have to ditch return types and use channels instead. That means instead of returning an `int64`, we will accept a `chan<- int64` as a parameter (this is a channel we can write to, accepting `int64`s). Instead of returning numbers we have to write them into the channel. To recursivly call the function we have to create a channel first.

We will also rename the function to `fibChan` to make clear it uses channels:

```go
// fibonacci/fibonacci.go

func fibChan(n int64, res chan<- int64) {
    if n == 0 {
		res <- 0
		return
	}

	if n == 1 {
		res <- 1
		return
	}

	ch := make(chan int64)
	fibChan(n-1, ch)
	fibChan(n-2, ch)

	res <- (<-ch) + (<-ch)
}
```

The last line of this is pretty interesting: It reads two `int64`s from the channel `ch` and write their sum to the channel `res`.

To multithread this approach, all we have to do is to write `go` before the recursive calls to `fibChan`:

```go
// fibonacci/fibonacci.go

    // ...

	ch := make(chan int64)
	go fibChan(n-1, ch)
	go fibChan(n-2, ch)

	res <- (<-ch) + (<-ch)
```

This tells go to run these calls asynchronously.


What we should do now is create a wrapper function to `fibChan` which simplifies calling it (so we don't have to create a channel manually every time we want to know some fibonacci number):

```go
// fibonacci/fibonacci.go

func Fib(n int64) int64 {
	ch := make(chan int64)
	go fibChan(n, ch)
	return <- ch
}

```

We can now call `Fib` from our main.go:


```go
// main.go
import (
	"fmt"
	"github.com/lucaschimweg/dhbw-go-portfolio/fibonacci"
)

func main() {
	fmt.Println(fibonacci.Fib(25))
}

```

To run this, we type `go run main.go` in our terminal. If you receive the output `75025`, you have completed the second tutorial! You can find the full code for fibonacci [here](./fibonacci/fibonacci.go).

