# leb128
[![Travis](https://img.shields.io/travis/Equim-chan/leb128.svg)](https://travis-ci.org/Equim-chan/leb128)
[![Coverage Status](https://img.shields.io/codecov/c/gh/Equim-chan/leb128.svg)](https://codecov.io/gh/Equim-chan/leb128)
[![Release](https://img.shields.io/github/release/Equim-chan/leb128.svg)](https://github.com/Equim-chan/leb128/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Equim-chan/leb128)](https://goreportcard.com/report/github.com/Equim-chan/leb128)
[![License](https://img.shields.io/badge/BSD-3-blue.svg)](https://github.com/Equim-chan/leb128/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/ekyu.moe/leb128)

Go implementation of [Little Endian Base 128 codec](https://en.wikipedia.org/wiki/LEB128).

## Install
```bash
$ go get -u ekyu.moe/base256
# or better
$ dep ensure -add ekyu.moe/base256
```

## Example
```go
package main

import (
    "fmt"

    "ekyu.moe/leb128"
)

func main() {
    // Notes: These specs are taken from https://en.wikipedia.org/wiki/LEB128
    // Encode unsigned LEB128
    fmt.Printf("%x\n", leb128.AppendUleb128(nil, 624485)) //=> e58e26

    // Encode signed LEB128
    fmt.Printf("%x\n", leb128.AppendSleb128(nil, -624485)) //=> 9bf159

    // Decode unsigned LEB128, n is the number of bytes read
    u, n := leb128.DecodeUleb128([]byte{0xe5, 0x8e, 0x26, 'a', 'b', 'c'})
    fmt.Printf("%d %d\n", u, n) //=> 624485 3

    // Decode signed LEB128, n is the number of bytes read
    s, n := leb128.DecodeSleb128([]byte{0x9b, 0xf1, 0x59, 'd', 'e', 'f'})
    fmt.Printf("%d %d\n", s, n) //=> -624485 3
}
```

## LICENSE
[BSD-3-clause](https://github.com/Equim-chan/base256/blob/master/LICENSE)

Notes: The encode part is an edited fork of [/src/cmd/internal/dwarf/dwarf.go](https://golang.org/src/cmd/internal/dwarf/dwarf.go) licensed under [a BSD-style license](https://golang.org/LICENSE).
