package ludecomp

import "github.com/unixpickle/num-analysis/kahan"

type Vector []float64

func (v Vector) Dot(v1 Vector) float64 {
	values := make([]float64, len(v))
	for i, x := range v {
		values[i] = x * v1[i]
	}
	return kahan.Sum64(values)
}
