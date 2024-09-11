package example

import (
	"fmt"

	"github.com/jainrankit/go-ignore-cov/example/hello"
)

// this package should have 100% code coverage if we remove the ignored statements
func MaybeSayHello() {
	// coverage:ignore
	if err, ok := hello.SayHello(); err != nil && ok {
		// coverage:ignore
		fmt.Println("BOOM")
	}
	// coverage:ignore
	fmt.Println("OK")
}
