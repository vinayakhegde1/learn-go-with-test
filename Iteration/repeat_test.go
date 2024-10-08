package iteration

import "testing"

const repeatedCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}
	return repeated
}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
