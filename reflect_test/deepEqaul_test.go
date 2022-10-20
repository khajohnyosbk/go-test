package reflecttest

import (
	"testing"
)

type TestCase struct {
	x string
	y string
}

type TestCases []TestCase

var testCases = TestCases{
	TestCase{
		x: "ssfasf",
		y: "asfasfacasdf",
	},
	TestCase{
		x: "asfasfaszxcasf",
		y: "asfaslcaspdfoaspf",
	},
	TestCase{
		x: "a",
		y: "b",
	},
}

func BenchmarkReflectDeepCompare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, t := range testCases {
			reflectDeepCompare(&t.x, &t.y)
		}
	}
}

func BenchmarkNormalCompare(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, t := range testCases {
			normalCompare(&t.x, &t.y)
		}
	}
}
