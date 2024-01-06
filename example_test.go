package maybe_test

import (
	"fmt"

	"github.com/shota3506/maybe"
)

func ExampleOption() {
	opt := maybe.New("Hello World!")
	if v, ok := opt.Value(); ok {
		// do something
		fmt.Printf("value is %s", v)
	}
}
