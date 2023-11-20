package semver

import (
	"testing"
)

type testCase struct {
	v1       string
	v2       string
	expected int
}

func TestSemver(t *testing.T) {
	tests := []testCase{
		{"1.0.0", "1.0.0", 0},
		{"1.0.1", "1.0.0", 1},
		{"1.0.2", "1.2.3", -1},
		{"1.0.0-alpha", "1.0.0-alpha", 0},
		{"1.0.0-beta", "1.0.0-alpha", 1},
		{"1.0.0-alpha", "1.0.0-beta", -1},
		{"1.0.0-alpha+001", "1.0.0-beta+001", -1},
		{"1.0.0-beta+001", "1.0.0-alpha+001", 1},
		{"1.0.0+20130313144700", "1.0.0+20130313144701", 0}, // metadata does not affect precedence
		{"1.0.0-alpha+001", "1.0.0-alpha+002", 0},
		{"1.0.0", "1.0.1", -1},
		{"1.0.1", "1.0.0", 1},
		{"1.0.0", "1.0.0", 0},
	}

	for _, test := range tests {
		c, err := Compare(test.v1, test.v2)
		if err != nil {
			t.Error(err)
		}
		if c != test.expected {
			t.Errorf("expected %s and %s to be %d but got %d", test.v1, test.v2, test.expected, c)
		}
	}
}
