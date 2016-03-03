// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/steevee/rsc/devweb/slave"

	_ "github.com/steevee/rsc/blog/post"
)

func main() {
	slave.Main()
}
