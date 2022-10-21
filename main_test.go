package main

import "testing"

func Test_getLastNonEmptyPart(t *testing.T) {
	tests := []struct {
		name  string
		parts []string
		want  string
	}{
		{"Empty array", []string{}, ""},
		{"Single element", []string{"foo"}, "foo"},
		{"Single, empty element", []string{""}, ""},
		{"Several trailing, empty elements", []string{"foo", "", "", ""}, "foo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastNonEmptyPart(tt.parts); got != tt.want {
				t.Errorf("getLastNonEmptyPart() = %v, want %v", got, tt.want)
			}
		})
	}
}
