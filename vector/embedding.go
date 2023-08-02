package vector

import "gonum.org/v1/gonum/mat"

type Embedding struct {
	Vector *mat.VecDense // the embedding vector
	Data   string        // the data represted by the embedding
}
