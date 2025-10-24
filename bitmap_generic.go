// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build noasm || (!amd64 && !arm64)

package bitmap

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) And(other Bitmap, extra ...Bitmap) {
	max := minlen(*dst, other, extra)
	dst.shrink(max)
	and(*dst, max, other, extra)
}

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) AndNot(other Bitmap, extra ...Bitmap) {
	max := minlen(*dst, other, extra)
	andn(*dst, max, other, extra)
}

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Or(other Bitmap, extra ...Bitmap) {
	max := maxlen(*dst, other, extra)
	dst.grow(max - 1)
	bitmaps := growMany(append([]Bitmap{other}, extra...), max)
	or(*dst, bitmaps[0], bitmaps[1:])
}

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (dst *Bitmap) Xor(other Bitmap, extra ...Bitmap) {
	max := maxlen(*dst, other, extra)
	dst.grow(max - 1)
	bitmaps := growMany(append([]Bitmap{other}, extra...), max)
	xor(*dst, bitmaps[0], bitmaps[1:])
}

// Count returns the number of elements in this bitmap
func (dst Bitmap) Count() int {
	return count(dst)
}
