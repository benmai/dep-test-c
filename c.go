// package c imports package a and package b.
package c

import (
	a "github.com/benmai/dep-test-a"
	b "github.com/benmai/dep-test-b"
)

func C() {
	a.A()
	b.B()
}
