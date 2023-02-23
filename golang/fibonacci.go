package fibonacci

import "math/big"

// classic method
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

// optimized easy understand
func FibonacciEasyOptimized(n int) int {
	if n <= 2 {
		return 1
	}

	prePre, pre, cur := 1, 1, 1
	for i := 3; i <= n; i++ {
		cur = prePre + cur
		prePre = pre
		pre = cur
	}

	return cur
}

// optimized
// but 93 will overflows 9223372036854775807
/*
94:19740274219868223167
   9223372036854775807
89:1779979416004714189
90:2880067194370816120
91:4660046610375530309
92:7540113804746346429
93:12200160415121876738

int8: -128 ~ 127
int16: -32768 ~ 32767
int32: -2147483648 ~ 2147483647
int64: -9223372036854775808 ~ 9223372036854775807
uint8: 0 ~ 255
uint16: 0 ~ 65535
uint32: 0 ~ 4294967295
uint64: 0 ~ 18446744073709551615
*/
func FibonacciOptimized(n int) int {
	if n <= 2 {
		return 1
	}

	pre, cur := 1, 1
	for i := 3; i <= n; i++ {
		pre, cur = cur, pre + cur
	}

	return cur
}

func FibonacciOptimizedBig(n int) *big.Int {
	if n <= 2 {
		return big.NewInt(1)
	}

	pre, cur := big.NewInt(1), big.NewInt(1)
	for i := 3; i <= n; i++ {
		pre, cur = cur, pre.Add(pre, cur)
	}

	return cur
}

func StringToBigint(s string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil
	}

	return n
}