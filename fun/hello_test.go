
package fun

import (
  "testing"
  "github.com/hello/fun"
)

func TestHello(t *testing.T) {
  want := fun.Hello();
  got := "Goodbye, moon."
  if got != want {
    t.Errorf("Main() = %q, want %q", got, want)
  }
}
