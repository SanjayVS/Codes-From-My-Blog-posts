package main

import (
	"testing"
)

func BenchmarkFibV1(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
