// Copyright 2022 The age Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import "filippo.io/age/internal/testkit"

func main() {
	f := testkit.NewTestFile()
	f.VersionLine("v1")
	f.X25519(testkit.TestX25519Recipient)
	f.HMAC()
	f.Nonce(f.Rand(12))
	// Marked as header failure because we read the nonce while reading the
	// header, before handing off to the STREAM implementation.
	f.ExpectHeaderFailure()
	f.Generate()
}
