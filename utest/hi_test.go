package utest

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// this is like before after
	// only running in one package
	fmt.Println("Before unit test")

	m.Run() //run all test case

	fmt.Println("After unit test")
}

func BenchmarkHiOneSub(b *testing.B) {
	b.Run("Wisma", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hi("Wisma")
		}
	})
	b.Run("Eka", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hi("Eka")
		}
	})

}
func BenchmarkHiOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hi("wisma")
	}
}
func BenchmarkHiTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hi("eka")
	}
}

func BenchmarkTableHi(b *testing.B) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Test Wisma",
			request:  "Wisma",
			expected: "Hello Wisma",
		},
		{
			name:     "Test Eka",
			request:  "Eka",
			expected: "Hello Eka",
		},
		{
			name:     "Test Putu",
			request:  "Putu",
			expected: "Hello Putu",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := Hi(test.request)
				require.Equal(b, test.expected, result)
			}
		})
	}
}
func TestTableHi(t *testing.T) {

	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Test Wisma",
			request:  "Wisma",
			expected: "Hello Wisma",
		},
		{
			name:     "Test Eka",
			request:  "Eka",
			expected: "Hello Eka",
		},
		{
			name:     "Test Putu",
			request:  "Putu",
			expected: "Hello Putu",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hi(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Golang Subtest
func TestSubTest(t *testing.T) {
	t.Run("Wisma", func(t *testing.T) {
		result := Hi("Wisma")
		//expected, result
		assert.Equal(t, "Hello Wismat", result, "Result must be Hello Wisma")
		fmt.Println("this is for ub test 1")
	})
	t.Run("Eka", func(t *testing.T) {
		//expected, result
		result := Hi("Eka")
		assert.Equal(t, "Hello Eka", result, "Result must be Hello Eka")
		fmt.Println("this is for ub test 2")
	})
}

// End of subtest

func TestHiSkip(t *testing.T) {
	// this skip the test not failing the test
	if runtime.GOOS == "darwin" {
		t.Skip("cannot run in mac os")
	}
	result := Hi("Wisma")
	//expected, result
	assert.Equal(t, "Hello Wismat", result, "Result must be Hello Wisma")
	fmt.Println("ASSERT")
}
func TestHiAssert(t *testing.T) {
	result := Hi("Wisma")
	//expected, result
	assert.Equal(t, "Hello Wismat", result, "Result must be Hello Wisma")
	fmt.Println("ASSERT")
}
func TestHiRequire(t *testing.T) {
	result := Hi("Wisma")
	//expected, result //call fail now
	require.Equal(t, "Hello Wismat", result, "Result must be Hello Wisma")
	fmt.Println("Require")
}

func TestHiOne(t *testing.T) {
	result := Hi("Wisma")
	if result != "Hello Wisma2" {
		t.Fail()
	}
	//continue here
	fmt.Println("this is fail")
}
func TestHiTwo(t *testing.T) {
	result := Hi("Wisma")
	if result != "Hello Wisma1" {
		t.FailNow()
		//stopped
	}
	fmt.Println("this is fail now")
}
func TestHiThree(t *testing.T) {
	result := Hi("Wisma")
	if result != "Hello Wisma1" {
		t.Fatal("Result must be Hello Wisma")
		//stopped
	}
	fmt.Println("this is fatal")
}
func TestHiFour(t *testing.T) {
	result := Hi("Wisma")
	if result != "Hello Wisma1" {
		t.Error("Result must be Hello Wisma")
	}
	fmt.Println("this is error")
}
