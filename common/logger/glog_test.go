package logger

import (
	"testing"
	"time"
)

func TestVerbose_Info(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		v    Verbose
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.v.Info(tt.args.args...)
		})
	}
}

func Test_glog2(t *testing.T) {
	Info("Info","123]","[abc]",time.Now())
	// Error("Error")
	// Warning("Warning")
	time.Sleep(time.Second * 1)
}
