// Copyright 2023 Aiden Fox Ivey, Alisya K. All rights reserved.
// Use of this source code is governed by an epic-style
// license that can be found in the LICENSE file.

// Package merkle implements a SHA256-backed Merkle tree.

package merkle

import (
	"crypto/sha256"
)

type MerkleTree struct {
	root       *Node
	merkleRoot []byte
	leaves     []*Node
}

type Node struct {
	left   *Node
	right  *Node
	parent *Node
	value  []byte
}

func (n Node) hash() []byte {
	hashArray := sha256.Sum256(n.value)
	return hashArray[:]
}

// Consume the bytes b of some serialized object, and return a valid Merkle tree.
// func NewMerkleTree(b []byte) (*MerkleTree, error) {
// 	t := &MerkleTree{}
// 	root, leafs, err := buildWithContent(b)
// 	if err != nil {
// 		return nil, err
// 	}
// 	t.root = root
// 	t.leaves = leafs
// 	t.merkleRoot = root.hash()
// 	return t, nil
// }

// func buildWithContent(b []byte) (*Node, []*Node, error) {

// }

// chew on bytes and gives us the chunks
// -\_(*_*)_/-
func splitIntoDataBlocks(b []byte, blockSize int) []*Node {
	output := make([]*Node, 0)

	for i := 0; i < len(b); i += blockSize {
		end := i + blockSize

		if end > len(b) {
			end = len(b)
		}

		b := sha256.Sum256(b[i:end])
		output = append(output, &Node{nil, nil, nil, b[:]})
	}
	return output
}

// spits out the root hash of the merkle tree
func Apply(b []byte) uint32 {
	// dataBlocks := splitIntoDataBlocks(b)
	// rootHash := recursion(dataBlocks)
	// return rootHash

	return 3
}

// OVERFLOW - watch out
func add(a []byte, b []byte) []byte {
	c := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func recursion(left, right Node) Node {
	f := sha256.Sum256(add(left.value, right.value))
	return Node{left: &left, right: &right, value: f[:]}
}
