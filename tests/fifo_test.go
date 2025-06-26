package cache_test

import (
	"testing"

	"caching-labwork/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFIFOCache tests the FIFO cache implementation
func TestFIFOCache(t *testing.T) {
	c := cache.NewFIFOCache[string, int](3)

	// Test basic operations
	err := c.Set("a", 1)
	require.NoError(t, err)

	val, err := c.Get("a")
	require.NoError(t, err)
	assert.Equal(t, 1, val)

	// Test FIFO eviction
	err = c.Set("b", 2)
	require.NoError(t, err)
	err = c.Set("c", 3)
	require.NoError(t, err)
	err = c.Set("d", 4) // This should evict "a"
	require.NoError(t, err)

	_, err = c.Get("a")
	assert.Error(t, err)
	assert.Equal(t, cache.ErrKeyNotFound, err)

	val, err = c.Get("b")
	require.NoError(t, err)
	assert.Equal(t, 2, val)

	// Test delete
	err = c.Delete("b")
	require.NoError(t, err)

	_, err = c.Get("b")
	assert.Error(t, err)

	// Test clear
	c.Clear()
	_, err = c.Get("c")
	assert.Error(t, err)
} 