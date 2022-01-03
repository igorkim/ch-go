//go:build !(amd64 || arm64) || purego

// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"

	"github.com/go-faster/errors"
)

var _ = binary.LittleEndian // clickHouse uses LittleEndian

// DecodeColumn decodes UInt32 rows from *Reader.
func (c *ColUInt32) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 32 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := *c
	// Move bound check out of loop.
	//
	// See https://github.com/golang/go/issues/30945.
	_ = data[len(data)-size]
	for i := 0; i <= len(data)-size; i += size {
		v = append(v,
			binary.LittleEndian.Uint32(data[i:i+size]),
		)
	}
	*c = v
	return nil
}

// EncodeColumn encodes UInt32 rows to *Buffer.
func (c ColUInt32) EncodeColumn(b *Buffer) {
	const size = 32 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint32(
			b.Buf[offset:offset+size],
			v,
		)
		offset += size
	}
}
