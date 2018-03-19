assert
======

# About

Assert serves as a helper for golang testing.

# Install

```golang
import "github.com/damondouglas/assert"
```

# Usage

```golang
func TestSomeThing(t *testing) {
    a := assert.New(t)
    a.Equals(1,1)       // ok
    a.NotEquals(1,2)    // Test fails here
}
```