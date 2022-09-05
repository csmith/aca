package aca

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGenerator_withDelimiters(t *testing.T) {
	gen1, err := NewGenerator("-", rand.NewSource(0))
	assert.NoErrorf(t, err, "NewGenerator should not error")
	assert.Regexp(t, "^[a-z]+\\-[a-z]+\\-[a-z]+$", gen1.Generate(), "Generator should use configured delimiters")

	gen2, err := NewGenerator(" ~delim~ ", rand.NewSource(0))
	assert.NoErrorf(t, err, "NewGenerator should not error")
	assert.Regexp(t, "^[a-z]+ ~delim~ [a-z]+ ~delim~ [a-z]+$", gen2.Generate(), "Generator should use configured delimiters")
}

func TestGenerator_consistentWithSameSource(t *testing.T) {
	gen1, err := NewGenerator("-", rand.NewSource(0))
	assert.NoErrorf(t, err, "NewGenerator should not error")
	first := gen1.Generate()

	gen2, err := NewGenerator("-", rand.NewSource(0))
	assert.NoErrorf(t, err, "NewGenerator should not error")
	second := gen2.Generate()

	assert.Equal(t, first, second, "Generators with similarly seeded sources should return the same strings")
}

func TestGenerator_consistentWithDifferentSource(t *testing.T) {
	gen1, err := NewGenerator("-", rand.NewSource(0))
	assert.NoErrorf(t, err, "NewGenerator should not error")
	first := gen1.Generate()

	gen2, err := NewGenerator("-", rand.NewSource(1))
	assert.NoErrorf(t, err, "NewGenerator should not error")
	second := gen2.Generate()

	assert.NotEqual(t, first, second, "Generators with differently seeded sources should return the same strings")
}

func Example() {
	gen, err := NewGenerator("-", rand.NewSource(time.Now().UnixNano()))
	if err != nil {
		// Handle this better!
		panic(err)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(gen.Generate())
	}
}
