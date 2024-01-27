// Copyright 2023 Aiden Fox Ivey, Alisya K. All rights reserved.
// Use of this source code is governed by an epic-style
// license that can be found in the LICENSE file.

package merkle

import (
	"crypto/sha256"
	"bytes"
	"testing"
)

func Test_Bytes(t *testing.T) {
	b := []byte("hello")
	nodes := splitIntoDataBlocks(b, 4)

	expectedBlocks := [][]byte{
		[]byte("hell"),
		[]byte("o"),
	}

	if len(nodes) != len(expectedBlocks) {
		t.Fatalf("Expected %d blocks, got %d.\n", len(nodes), len(expectedBlocks))
	}

	for i, node := range nodes {
		shaSum := sha256.Sum256(expectedBlocks[i])
		if !bytes.Equal(node.value, shaSum[:]) {
			t.Fatal("Mismatching bytes.")
		}
	}
}
