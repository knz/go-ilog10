package ilog10

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastUint64Log10(t *testing.T) {
	testData := []uint64{
		1,
		7,
		9,
		10,
		11,
		999,
		1000,
		1001,
		1<<63 - 1,
		1 << 63,
		9999999999999999999,
		10000000000000000000,
		10000000000000000001,
		18446744073709551615, // 2^64-1 (Uint64Max)
	}

	for _, tc := range testData {
		repr := strconv.FormatUint(tc, 10)
		exp := uint(len(repr) - 1)
		assert.Equal(t, exp, FastUint64Log10(tc), "number = %d", tc)
	}

	// Some random numbers for good measure.
	for i := 0; i < 100; {
		// We want a random number of digits, not just a random number.
		magnitude := int(rand.Float32() * float32(64))
		n := rand.Uint64() >> magnitude
		if n == 0 {
			continue
		}

		repr := strconv.FormatUint(n, 10)
		exp := uint(len(repr) - 1)
		assert.Equal(t, exp, FastUint64Log10(n), "number = %d", n)
		i++
	}
}

func TestFastUint32Log10(t *testing.T) {
	testData := []uint32{
		1,
		7,
		9,
		10,
		11,
		999,
		1000,
		1001,
		1<<31 - 1,
		1 << 31,
		999999999,
		1000000000,
		1000000001,
		4294967295, // 2^32-1 (Uint32Max)
	}

	for _, tc := range testData {
		repr := strconv.FormatUint(uint64(tc), 10)
		exp := uint(len(repr) - 1)
		assert.Equal(t, exp, FastUint32Log10(tc), "number = %d", tc)
	}

	// Some random numbers for good measure.
	for i := 0; i < 100; {
		// We want a random number of digits, not just a random number.
		magnitude := int(rand.Float32() * float32(32))
		n := rand.Uint32() >> magnitude
		if n == 0 {
			continue
		}

		repr := strconv.FormatUint(uint64(n), 10)
		exp := uint(len(repr) - 1)
		assert.Equal(t, exp, FastUint32Log10(n), "number = %d", n)
		i++
	}
}
