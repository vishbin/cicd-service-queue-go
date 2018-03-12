package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkConcatString(b *testing.B) {
	var str string

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += "x"
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	var buffer bytes.Buffer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	var builder strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		builder.WriteString("x")
	}
}
