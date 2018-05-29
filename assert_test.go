package assert

import (
	"errors"
	"testing"
)

func TestEquals(t *testing.T) {
	a := New(t)
	a.Equals(2, 2)
	a.Equals(nil, nil)
	a.Equals("foo", "foo")
	a.Equals(true, true)
}

func TestNotEquals(t *testing.T) {
	a := New(t)
	a.NotEquals(1, 2)
	a.NotEquals(1, nil)
	a.NotEquals("foo", "bar")
	a.NotEquals(true, false)
}

func TestHandleError(t *testing.T) {
	a := New(t)
	err := errors.New("some error")
	a.HandleError(err)
}
