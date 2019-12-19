package ilog10

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDigitsU64(t *testing.T) {
	testData := []uint64{
		0,
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
		exp := uint(len(repr))
		assert.Equal(t, exp, NumUint64DecimalDigits(tc), "number = %d", tc)
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
		exp := uint(len(repr))
		assert.Equal(t, exp, NumUint64DecimalDigits(n), "number = %d", n)
		i++
	}
}

func TestNumDigitsS64(t *testing.T) {
	testData := []int64{
		0,
		1, -1,
		7, -7,
		9, -9,
		10, -10,
		11, -11,
		999, -999,
		1000, -1000,
		1001, -1001,
		1<<62 - 1, -(1<<62 - 1),
		1 << 62, -(1 << 62),
		999999999999999999,
		-999999999999999999,
		1000000000000000000,
		-1000000000000000000,
		9223372036854775807,  // Int64Min
		-9223372036854775808, // Int64Max
	}

	for _, tc := range testData {
		repr := strconv.FormatInt(tc, 10)
		exp := uint(len(repr))
		if tc < 0 {
			exp = exp - 1 // Remove the minus sign
		}
		assert.Equal(t, exp, NumInt64DecimalDigits(tc), "number = %d", tc)
	}

	// Some random numbers for good measure.
	for i := 0; i < 100; {
		// We want a random number of digits, not just a random number.
		magnitude := int(rand.Float32() * float32(64))
		u := rand.Uint64() >> magnitude
		if u == 0 {
			continue
		}
		n := int64(u)

		repr := strconv.FormatInt(n, 10)
		exp := uint(len(repr))
		if n < 0 {
			exp = exp - 1 // Remove the minus sign
		}
		assert.Equal(t, exp, NumInt64DecimalDigits(n), "number = %d", n)
		i++
	}
}

func TestNumDigitsU32(t *testing.T) {
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
		exp := uint(len(repr))
		assert.Equal(t, exp, NumUint32DecimalDigits(tc), "number = %d", tc)
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
		exp := uint(len(repr))
		assert.Equal(t, exp, NumUint32DecimalDigits(n), "number = %d", n)
		i++
	}
}

func TestNumDigitsS32(t *testing.T) {
	testData := []int32{
		0,
		1, -1,
		7, -7,
		9, -9,
		10, -10,
		11, -11,
		999, -999,
		1000, -1000,
		1001, -1001,
		1<<30 - 1, -(1<<30 - 1),
		1 << 30, -(1 << 30),
		999999999,
		-999999999,
		1000000000,
		-1000000000,
		2147483647,  // Int32Min
		-2147483648, // Int32Max
	}

	for _, tc := range testData {
		repr := strconv.FormatInt(int64(tc), 10)
		exp := uint(len(repr))
		if tc < 0 {
			exp = exp - 1 // Remove the minus sign
		}
		assert.Equal(t, exp, NumInt32DecimalDigits(tc), "number = %d", tc)
	}

	// Some random numbers for good measure.
	for i := 0; i < 100; {
		// We want a random number of digits, not just a random number.
		magnitude := int(rand.Float32() * float32(32))
		u := rand.Uint32() >> magnitude
		if u == 0 {
			continue
		}
		n := int32(u)

		repr := strconv.FormatInt(int64(n), 10)
		exp := uint(len(repr))
		if n < 0 {
			exp = exp - 1 // Remove the minus sign
		}
		assert.Equal(t, exp, NumInt32DecimalDigits(n), "number = %d", n)
		i++
	}
}
