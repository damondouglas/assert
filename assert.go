package assert

import (
	"runtime/debug"
	"testing"
)

// Assert serves as a helper for testing
type Assert struct {
	t *testing.T
}

// New instantiates Assert
func New(t *testing.T) *Assert {
	a := new(Assert)
	a.t = t
	return a
}

// Equals evaluates equivalence between two interfaces
func (a *Assert) Equals(k interface{}, m interface{}) {
	if k != m {
		a.t.Errorf("%v should equal %v", k, m)
		buf := debug.Stack()
		a.t.Log(string(buf))
	}
}

// NotEquals evaluates non-equivalence between two interfaces
func (a *Assert) NotEquals(k interface{}, m interface{}) {
	if k == m {
		a.t.Errorf("%v should not equal %v", k, m)
		buf := debug.Stack()
		a.t.Log(string(buf))
	}
}

// HandleError logs error and outputs stacktrace
func (a *Assert) HandleError(err error) {
	if err != nil {
		a.t.Error(err)
		buf := debug.Stack()
		a.t.Log(string(buf))
	}
}
