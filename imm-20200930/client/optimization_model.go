// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimization interface {
	dara.Model
	String() string
	GoString() string
	SetLearningRate(v float32) *Optimization
	GetLearningRate() *float32
	SetOptimizer(v string) *Optimization
	GetOptimizer() *string
}

type Optimization struct {
	// example:
	//
	// 0.01
	LearningRate *float32 `json:"LearningRate,omitempty" xml:"LearningRate,omitempty"`
	// example:
	//
	// SGD
	Optimizer *string `json:"Optimizer,omitempty" xml:"Optimizer,omitempty"`
}

func (s Optimization) String() string {
	return dara.Prettify(s)
}

func (s Optimization) GoString() string {
	return s.String()
}

func (s *Optimization) GetLearningRate() *float32 {
	return s.LearningRate
}

func (s *Optimization) GetOptimizer() *string {
	return s.Optimizer
}

func (s *Optimization) SetLearningRate(v float32) *Optimization {
	s.LearningRate = &v
	return s
}

func (s *Optimization) SetOptimizer(v string) *Optimization {
	s.Optimizer = &v
	return s
}

func (s *Optimization) Validate() error {
	return dara.Validate(s)
}
