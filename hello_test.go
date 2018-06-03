package hello_test

import (
	"testing"
	"github.com/gavincabbage/hello"
)

func TestSay(t *testing.T) {
	if hello.Say() != "Hello!" {
		t.FailNow()
	}
}

func TestSayTo(t *testing.T) {
	if hello.SayTo("World") != "Hello, World!" {
		t.FailNow()
	}
}
