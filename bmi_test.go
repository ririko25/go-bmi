package bmi_test

import (
	"testing"

	"github.com/ririko25/go-bmi"
)

func TestCalc(t *testing.T) {

	tests := []struct {
		name    string
		weight  int
		height  int
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "hoge", weight: 45, height: 153, want: 19, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bmi.Calc(tt.weight, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
