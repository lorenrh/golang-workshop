package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := []string{"1", "2", "3", "4"}

	res, err := sum(numbers)

	if err != nil {
		t.Errorf("sum expected nil but got %s", err)
	}

	if res != 10 {
		t.Errorf("sum expected %d but got %d", 10, res)
	}

	assert.NoError(t, err)
	assert.Equal(t, 10, res)
}

func TestSumNegative(t *testing.T) {
	numbers := []string{"-1", "-2", "-3", "-4"}

	res, err := sum(numbers)

	if err != nil {
		t.Errorf("sum expected nil but got %s", err)
	}

	if res != -10 {
		t.Errorf("sum expected %d but got %d", -10, res)
	}

	assert.NoError(t, err)
	assert.Equal(t, -10, res)
}

func TestError(t *testing.T) {
	numbers := []string{"st"}

	_, err := sum(numbers)

	if err == nil {
		t.Errorf("sum expected nil but got %s", err)
	}
}
