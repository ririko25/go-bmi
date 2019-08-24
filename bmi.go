package bmi

import "math"

// Calc bmiを作る
func Calc(weight, height int) (float64, error) {
	heightInMeter := float64(height) / 100.0
	bmi := float64(weight) / (heightInMeter * heightInMeter)
	bmi2 := round(bmi)
	return bmi2, nil
}

// 四捨五入
func round(f float64) float64 {
	return math.Floor(f + .5)

}
