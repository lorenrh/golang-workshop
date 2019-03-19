package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	numbers := []int{3, 6, -2, -5, 7, 3}
	res := adjacentElementsProduct(numbers)

	assert.Equal(t, 21, res)

	assert.Equal(t, adjacentElementsProduct([]int{5, 1, 2, 3, 1, 4}), 6)

	assert.Equal(t, adjacentElementsProduct([]int{1, 2, 3, 0}), 6)
	assert.Equal(t, adjacentElementsProduct([]int{9, 5, 10, 2, 24, -1, -48}), 50)
	assert.Equal(t, adjacentElementsProduct([]int{-23, 4, -3, 8, -12}), -12)

}
