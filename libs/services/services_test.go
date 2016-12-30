package services

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestService(t *testing.T) {
	Init()
	spew.Dump(_default_pool)

	t.Log("service ends...")
}
