//go:build amd64 && !nounsafe

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes Date rows from *Reader.
func (c *ColDate) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([]Date, rows)...)
	s := *(*slice)(unsafe.Pointer(c))
	s.Len *= 2
	s.Cap *= 2
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}

// EncodeColumn encodes Date rows to *Buffer.
func (c ColDate) EncodeColumn(b *Buffer) {
	if len(c) == 0 {
		return
	}
	offset := len(b.Buf)
	const size = 16 / 8
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	s := *(*slice)(unsafe.Pointer(&c))
	s.Len *= 2
	s.Cap *= 2
	src := *(*[]byte)(unsafe.Pointer(&s))
	dst := b.Buf[offset:]
	copy(dst, src)
}
