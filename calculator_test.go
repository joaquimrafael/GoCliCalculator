package main

import "testing"

func TestCalculate(t *testing.T) {
	cases := []struct {
		name    string
		a, b    float64
		op      string
		want    float64
		wantErr bool
	}{
		{name: "addition", a: 2, b: 3, op: "+", want: 5},
		{name: "subtraction", a: 5, b: 3, op: "-", want: 2},
		{name: "multiplication", a: 4, b: 3, op: "*", want: 12},
		{name: "division", a: 10, b: 4, op: "/", want: 2.5},
		{name: "negatives", a: -2, b: -3, op: "+", want: -5},
		{name: "division by zero", a: 1, b: 0, op: "/", wantErr: true},
		{name: "invalid operator", a: 1, b: 2, op: "%", wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculate(tc.a, tc.b, tc.op)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("calculate(%v, %v, %q): expected an error, got nil", tc.a, tc.b, tc.op)
				}
				return
			}

			if err != nil {
				t.Fatalf("calculate(%v, %v, %q): unexpected error: %v", tc.a, tc.b, tc.op, err)
			}
			if got != tc.want {
				t.Errorf("calculate(%v, %v, %q) = %v, want %v", tc.a, tc.b, tc.op, got, tc.want)
			}
		})
	}
}
