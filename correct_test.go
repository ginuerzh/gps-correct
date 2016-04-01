package correct

import (
	"testing"
)

func TestCorrectGoogle(t *testing.T) {
	lat := 30.123456
	lng := 100.123456
	loc := CorrectGoogle(lat, lng)
	t.Log(loc)
}
