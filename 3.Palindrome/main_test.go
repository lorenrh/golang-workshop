package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindrome(t *testing.T) {
	assert.Equal(t, checkPalindrome("aabaa"), true)
	assert.Equal(t, checkPalindrome("abac"), false)
	assert.Equal(t, checkPalindrome("abacaba"), true)
	assert.Equal(t, checkPalindrome("az"), false)
	assert.Equal(t, checkPalindrome("z"), true)
	assert.Equal(t, checkPalindrome("zzzazzazz"), false)

}
