// Copyright 2017 Equim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leb128_test

import (
	"fmt"

	"ekyu.moe/leb128"
)

func ExampleAppendUleb128() {
	fmt.Printf("Encoded: %x\n", leb128.AppendUleb128(nil, 624485))
	// Output:
	// Encoded: e58e26
}

func ExampleAppendSleb128() {
	fmt.Printf("Encoded: %x\n", leb128.AppendSleb128(nil, -624485))
	// Output:
	// Encoded: 9bf159
}

func ExampleDecodeUleb128() {
	u, n := leb128.DecodeUleb128([]byte{0xe5, 0x8e, 0x26, 'a', 'b', 'c'})
	fmt.Printf("Decoded: %d\nRead: %d bytes", u, n)
	// Output:
	// Decoded: 624485
	// Read: 3 bytes
}

func ExampleDecodeSleb128() {
	s, n := leb128.DecodeSleb128([]byte{0x9b, 0xf1, 0x59, 'd', 'e', 'f'})
	fmt.Printf("Decoded: %d\nRead: %d bytes", s, n)
	// Output:
	// Decoded: -624485
	// Read: 3 bytes
}
