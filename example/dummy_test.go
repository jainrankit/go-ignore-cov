package example_test

import (
	"testing"

	"github.com/jainrankit/go-ignore-cov/example"
)

func TestSayHello(t *testing.T) {
	example.MaybeSayHello()
}
