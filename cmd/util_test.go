package cmd

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func Test_truncateString(t *testing.T) {
	type args struct {
		s string
		l int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args{"aaaaaaaaa", 10},
			"aaaaaaaaa",
		},
		{
			args{"aaaaaaaaaa", 10},
			"aaaaaaaaaa",
		},
		{
			args{"aaaaaaaaaaa", 10},
			"aaaaaaaaaa...",
		},
		{
			args{"aaaaaaaaa\naaaaa", 10},
			"aaaaaaaaa...",
		},
		{
			args{"aaaaaaaaaa\naaaaa", 10},
			"aaaaaaaaaa...",
		},
		{
			args{"aaaaaaaaaaa\naaaaa", 10},
			"aaaaaaaaaa...",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := truncateString(tt.args.s, tt.args.l)
			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{0, 0},
			0,
		},
		{
			args{1, 0},
			1,
		},
		{
			args{0, 1},
			1,
		},
		{
			args{-1, 0},
			0,
		},
		{
			args{0, -1},
			0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := max(tt.args.x, tt.args.y)
			assert.Equal(t, got, tt.want)
		})
	}
}
