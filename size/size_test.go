package size

import "testing"

func TestSize(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{1}, "small"},
		{"10", args{10}, "big"},
		{"100", args{100}, "huge"},
		{"1000", args{1000}, "enormous"},
		{"-1", args{-1}, "negative"},
		{"0", args{0}, "zero"},
		{"-10", args{-10}, "small"},
		{"-100", args{-100}, "big"},
		{"-1000", args{-1000}, "huge"},
		{"-10000", args{-10000}, "enormous"},
		{"-100000", args{-100000}, "enormous"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Size(tt.args.a); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
