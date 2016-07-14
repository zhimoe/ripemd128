package ripemd128

import (
	"fmt"
	"io"
	"testing"
)

type mdTest struct {
	out string
	in  string
}

var vectors = [...]mdTest{
	{"cdf26213a150dc3ecb610f18f6b38b46", ""},
	{"86be7afa339d0fc7cfc785e72f578d33", "a"},
	{"c14a12199c66e4ba84636b0f69144c77", "abc"},
}

func TestVectors(t *testing.T) {
	for i := 0; i < len(vectors); i++ {
		tv := vectors[i]
		md := New()
		for j := 0; j < 3; j++ {
			if j < 2 {
				io.WriteString(md, tv.in)
			} else {
				io.WriteString(md, tv.in[0:len(tv.in)/2])
				md.Sum(nil)
				io.WriteString(md, tv.in[len(tv.in)/2:])
			}
			s := fmt.Sprintf("%x", md.Sum(nil))
			if s != tv.out {
				t.Fatalf("RIPEMD-128[%d](%s) = %s, expected %s", j, tv.in, s, tv.out)
			}
			md.Reset()
		}
	}
}
