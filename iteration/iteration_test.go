package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectString := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("got '%s', expected '%s'", repeated, expected)
		}
	}

	t.Run("repeating a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectString(t, repeated, expected)
	})

	t.Run("repeating a 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertCorrectString(t, repeated, expected)
	})

	t.Run("repeating b 7 times", func(t *testing.T) {
		repeated := Repeat("b", 7)
		expected := "bbbbbbb"

		assertCorrectString(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}