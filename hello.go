package hello

import "fmt"

func Say() string {
	return "Hello!"
}

func SayTo(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}