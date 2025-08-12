package puppy

import (
	"fmt"

	"github.com/Yasmin-H/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func sayHello() string {
	return "Practising versioning :)"
}

func from11() string {
	fmt.Println("I'm from version 1.1.0")
}
