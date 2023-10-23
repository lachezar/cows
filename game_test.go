package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	var c []candidate = make([]candidate, 0, ((10 - 1) * 9 * 8 * 7))
	c = permutations(make([]uint8, 4), 3, c)

	assert.Equal(t, c[0], candidate{1, 0, 2, 3})
	assert.Equal(t, c[len(c)-1], candidate{9, 8, 7, 6})
}
