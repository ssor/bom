package bom

import (
	"bufio"
	"io"
)

const (
	bom0 = 0xef
	bom1 = 0xbb
	bom2 = 0xbf
)

// CleanBom returns b with the 3 byte BOM stripped off the front if it is present.
// If the BOM is not present, then b is returned.
func CleanBom(b []byte) []byte {
	if len(b) >= 3 &&
		b[0] == bom0 &&
		b[1] == bom1 &&
		b[2] == bom2 {
		return b[3:]
	}
	return b
}

// NewReaderWithoutBom returns an io.Reader that will skip over initial UTF-8 byte order marks.
func NewReaderWithoutBom(r io.Reader) io.Reader {
	buf := bufio.NewReader(r)
	b, err := buf.Peek(3)
	if err != nil {
		// not enough bytes
		return buf
	}
	if b[0] == bom0 && b[1] == bom1 && b[2] == bom2 {
		discardBytes(buf, 3)
	}
	return buf
}

func discardBytes(buf *bufio.Reader, n int) {
	// the Discard method was introduced in Go 1.5
	buf.Discard(n)
}
