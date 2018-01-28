package main

import (
	"fmt"
	"os"
	"plugin"
)

// Greeter Greeter interface
type Greeter interface {
	Greet()
}

func panicOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	var mod string
	switch lang {
	case "english":
		mod = "./plugins/eng/eng.so"
	case "chinese":
		mod = "./plugins/chi/chi.so"
	default:
		fmt.Println("don't speak that stupid language")
		os.Exit(1)
	}

	plug, err := plugin.Open(mod)
	panicOnErr(err)

	symGreeter, err := plug.Lookup("Greeter")
	panicOnErr(err)

	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("Unexpected module")
		os.Exit(1)
	}

	greeter.Greet()
}
