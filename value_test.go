package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestValueEqual(t *testing.T) {
	testCases := map[string]struct {
		f        func(mock testing.TB)
		errCount int
	}{
		"match nil": {
			f: func(mock testing.TB) {
				type s struct{}
				var p *s
				gt.Value(mock, p).Equal(nil)
			},
			errCount: 0,
		},
		"unmatched pointer": {
			f: func(mock testing.TB) {
				type s struct{}
				gt.Value(mock, &s{}).Equal(nil)
			},
			errCount: 1,
		},
		"match struct (ptr)": {
			f: func(mock testing.TB) {
				type s struct {
					a string
					b int
				}

				gt.Value(mock, &s{a: "x", b: 1}).Equal(&s{a: "x", b: 1})
			},
			errCount: 0,
		},
		"not match struct (ptr)": {
			f: func(mock testing.TB) {
				type s struct {
					a string
					b int
				}

				gt.Value(mock, &s{a: "x", b: 2}).Equal(&s{a: "x", b: 1})
			},
			errCount: 1,
		},
		"match number": {
			f: func(mock testing.TB) {
				gt.Value(mock, 1).Equal(1)
			},
			errCount: 0,
		},
		"unmatched number": {
			f: func(mock testing.TB) {
				gt.Value(mock, 1).Equal(2)
			},
			errCount: 1,
		},
		"match string": {
			f: func(mock testing.TB) {
				gt.Value(mock, "abc").Equal("abc")
			},
			errCount: 0,
		},
		"unmatched string": {
			f: func(mock testing.TB) {
				gt.Value(mock, "abc").Equal("xyz")
			},
			errCount: 1,
		},
		"matched []byte": {
			f: func(mock testing.TB) {
				gt.Value(mock, []byte("abc")).Equal([]byte("abc"))
			},
			errCount: 0,
		},
		"unmatched []byte": {
			f: func(mock testing.TB) {
				gt.Value(mock, []byte("abc")).Equal([]byte("xyz"))
			},
			errCount: 1,
		},
		"[]byte length 0 and greater than 0": {
			f: func(mock testing.TB) {
				gt.Value(mock, []byte{}).Equal([]byte("a"))
			},
			errCount: 1,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			cnt := newCounter()
			tc.f(cnt)

			if cnt.errs != tc.errCount {
				t.Errorf("error count is not match: expected %d, but actual %d", tc.errCount, cnt.errs)
			}
			if cnt.errs != tc.errCount {
				t.Errorf("error count is not match: expected %d, but actual %d", tc.errCount, cnt.errs)
			}
		})
	}
}

func TestValueNil(t *testing.T) {
	testCases := map[string]struct {
		value any
		isErr bool
	}{
		"match nil": {
			value: nil,
			isErr: false,
		},
		"string is not Nil": {
			value: "a",
			isErr: true,
		},
		"int is not Nil": {
			value: 0,
			isErr: true,
		},
		"float is not Nil": {
			value: 1.23,
			isErr: true,
		},
		"struct is not Nil": {
			value: struct{}{},
			isErr: true,
		},
		"struct ptr is not Nil": {
			value: &struct{}{},
			isErr: true,
		},
		"function is not Nil": {
			value: func() {},
			isErr: true,
		},
		"chain is not Nil": {
			value: make(chan bool),
			isErr: true,
		},
		"slice is not Nil": {
			value: []int{1, 2, 3},
			isErr: true,
		},
		"empty slice is Nil": {
			value: []int{},
			isErr: false,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			cnt := newCounter()

			gt.Value(cnt, tc.value).Nil()

			if (cnt.errs > 0) != tc.isErr {
				t.Errorf("Expected isErr: %v, but actual: %v (%T)", tc.isErr, cnt.errs, tc.value)
			}
		})
	}
}

func TestValueNotNil(t *testing.T) {
	testCases := map[string]struct {
		value any
		isErr bool
	}{
		"unmatched not nil": {
			value: nil,
			isErr: true,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			cnt := newCounter()

			gt.Value(cnt, tc.value).NotNil()

			if (cnt.errs > 0) != tc.isErr {
				t.Errorf("Expected isErr: %v, but actual: %v (%T)", tc.isErr, cnt.errs, tc.value)
			}
		})
	}
}
