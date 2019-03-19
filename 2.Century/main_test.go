package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYears(t *testing.T) {
	assert.Equal(t, centuryFromYear(1994), 20)
	assert.Equal(t, centuryFromYear(1700), 17)
	assert.Equal(t, centuryFromYear(1904), 20)
	assert.Equal(t, centuryFromYear(200), 2)
	assert.Equal(t, centuryFromYear(8), 1)
}
