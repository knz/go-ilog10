package ilog10

import (
	"math"
	"testing"
)

func BenchmarkFastLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numDigitsFastLog(v32())
	}
}

func numDigitsFastLog(u uint32) int {
	return 1 + int(FastUint32Log10(u))
}

func BenchmarkMathLog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numDigitsMathLog(v32())
	}
}

func numDigitsMathLog(u uint32) int {
	return int(math.Log10(float64(u))) + 1
}

func BenchmarkMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numDigitsMult(v32())
	}
}

func numDigitsMult(u uint32) int {
	n := 0
	tmp := uint32(1)
	for tmp <= u {
		n++
		tmp *= 10
	}
	return n
}

func BenchmarkManyIfs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numDigitsBranch(v32())
	}
}

// Courtesy of @nvanbenchoten.
func numDigitsBranch(u uint32) int {
	if u < 100000 {
		if u < 100 {
			if u < 10 {
				return 1
			}
			return 2
		}
		if u < 1000 {
			return 3
		}
		if u < 10000 {
			return 4
		}
		return 5
	}
	if u < 10000000 {
		if u < 1000000 {
			return 6
		}
		return 7
	}
	if u < 100000000 {
		return 8
	}
	if u < 1000000000 {
		return 9
	}
	return 10
}

//go:noinline
func v32() uint32 {
	// This fuunction exists to avoid constant propagation in the benchmark.
	return 999999999
}
