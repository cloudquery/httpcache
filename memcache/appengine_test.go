//go:build appengine
// +build appengine

package memcache

import (
	"testing"

	"github.com/cloudquery/httpcache/test"

	"appengine/aetest"
)

func TestAppEngine(t *testing.T) {
	ctx, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	test.Cache(t, New(ctx))
}
