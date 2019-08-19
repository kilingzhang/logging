package logging

import (
	"bytes"
	"reflect"
	"testing"
)

func TestStandardLogger(t *testing.T) {
	tests := []struct {
		name string
		want *Logger
	}{
		{
			"StandardLogger",
			StandardLogger(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StandardLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StandardLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetOutput(t *testing.T) {
	tests := []struct {
		name    string
		wantOut string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			SetOutput(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("SetOutput() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestSetFormatter(t *testing.T) {
	type args struct {
		formatter Formatter
	}
	tests := []struct {
		name string
		args args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetFormatter(tt.args.formatter)
		})
	}
}

func TestSetReportCaller(t *testing.T) {
	type args struct {
		include bool
	}
	tests := []struct {
		name string
		args args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetReportCaller(tt.args.include)
		})
	}
}

func TestSetLevel(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name string
		args args
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLevel(tt.args.level)
		})
	}
}

func TestGetLevel(t *testing.T) {
	tests := []struct {
		name string
		want Level
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLevel(); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLevelEnabled(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLevelEnabled(tt.args.level); got != tt.want {
				t.Errorf("IsLevelEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
