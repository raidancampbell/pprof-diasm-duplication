package main

import (
	"github.com/pkg/profile"
	"math/rand"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	for i := 0; i < 800000000; i++ {
		closeEnough(int32(rand.Uint32()), int32(rand.Uint32()), 100)
	}
}

func closeEnough(A, B, tolerance int32) bool {
	if A >= 0 && B >= 0 {
		return int32Abs(A-B) < tolerance
	}
	if A >= 0 && B < 0 {
		return int32Abs(A-int32Abs(B)) < tolerance
	}
	if A < 0 && B >= 0 {
		return int32Abs(int32Abs(A)-B) < tolerance
	}
	return int32Abs(int32Abs(A)-int32Abs(B)) < tolerance
}

//go:noinline
func int32Abs(i int32) int32 {
	mask := i >> 31
	return (mask + i) ^ mask
}
