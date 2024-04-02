package test_test

import (
	"testing"

	"github.com/cloudquery/httpcache"
	"github.com/cloudquery/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
