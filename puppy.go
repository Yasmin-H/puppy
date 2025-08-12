package puppy

import (
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
	return "I'm from version 1.1.0"
}
