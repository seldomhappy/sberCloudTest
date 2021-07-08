package shortString_test

import (
	shortString "sberCloudTest/pkg"
	"testing"
)

var testFlags = []struct {
	in  string
	out string
}{
	{"aaabbccccdeeeff", "3a2b4c1d3e2f"},
	{"jjjsllkkkkklllllll", "3j1s2l5k7l"},
	{"ihhhj", "1i3h1j"},
}

func TestShortString(t *testing.T) {
	for _, val := range testFlags {
		res := shortString.Get(val.in)
		if res != val.out {
			t.Errorf("GetShortString(%v) = %s, expected %s",
				val.in, res, val.out)
		}
	}
}
