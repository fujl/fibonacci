package fibonacci

import "testing"
import "math/big"

var testExpect = map[int]int {
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	10: 55,
	13: 233,
	20: 6765,
	30: 832040,
	35: 9227465,
}

var testBigExpect = map[int]int {
	50:  12586269025,
	55:  139583862445,
	60:  1548008755920,
	73:  806515533049393,
	90:  2880067194370816120,
}

var testExpectBigint = map[int]*big.Int {
	1: StringToBigint("1"),
	2: StringToBigint("1"),
	3: StringToBigint("2"),
	60: StringToBigint("1548008755920"),
	70: StringToBigint("190392490709135"),
	73: StringToBigint("806515533049393"),
	80: StringToBigint("23416728348467685"),
	90: StringToBigint("2880067194370816120"),
	93: StringToBigint("12200160415121876738"),
	94: StringToBigint("19740274219868223167"),
	95: StringToBigint("31940434634990099905"),
	96: StringToBigint("51680708854858323072"),
	97: StringToBigint("83621143489848422977"),
	98: StringToBigint("135301852344706746049"),
	99: StringToBigint("218922995834555169026"),
	100:StringToBigint("354224848179261915075"),
	200:StringToBigint("280571172992510140037611932413038677189525"),
}

func TestFibonacci(t *testing.T) {
	for target, expect := range testExpect {
		value := Fibonacci(target)
		if value != expect {
			t.Errorf("%d's fibonacci numer is %d, but expect is %d", target, value, expect)
		}
	}
}

func TestFibonacciEasyOptimized(t *testing.T) {
	for target, expect := range testExpect {
		value := FibonacciEasyOptimized(target)
		if value != expect {
			t.Errorf("%d's fibonacci numer is %d, but expect is %d", target, value, expect)
		}
	}
}

func TestFibonacciEasyOptimizedBig(t *testing.T) {
	for target, expect := range testBigExpect {
		value := FibonacciEasyOptimized(target)
		if value != expect {
			t.Errorf("%d's fibonacci numer is %d, but expect is %d", target, value, expect)
		}
	}
}

func TestFibonacciOptimized(t *testing.T) {
	for target, expect := range testExpect {
		value := FibonacciOptimized(target)
		if value != expect {
			t.Errorf("%d's fibonacci numer is %d, but expect is %d", target, value, expect)
		}
	}
}

func TestStringToBigint(t *testing.T) {
	expect := big.NewInt(10)
	sb := StringToBigint("10")
	if expect.Cmp(sb) != 0 {
		t.Error("10 StringToBigint is ", sb)
	}

	sb = StringToBigint("12s0")
	if sb != nil {
		t.Error("12s0 StringToBigint is ", sb)
	}
}

func TestFibonacciOptimizedBig(t *testing.T) {
	for target, expect := range testExpectBigint {
		value := FibonacciOptimizedBig(target)
		if value.Cmp(expect) != 0 {
			t.Error(target, "'s fibonacci numer is ", value, ", but expect is ", expect)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30)
	}
}

func BenchmarkFibonacciEasyOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciEasyOptimized(30)
	}
}

func BenchmarkFibonacciOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciOptimized(30)
	}
}

func BenchmarkFibonacciEasyOptimizedBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciEasyOptimized(92)
	}
}

func BenchmarkFibonacciOptimizedBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciOptimized(92)
	}
}

func BenchmarkFibonacciOptimizedBB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciOptimizedBig(92)
	}
}