// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressConfig interface {
	dara.Model
	String() string
	GoString() string
	SetNumMinibatches(v int32) *RLProgressConfig
	GetNumMinibatches() *int32
	SetPpoMiniBatchSize(v int32) *RLProgressConfig
	GetPpoMiniBatchSize() *int32
	SetRolloutN(v int32) *RLProgressConfig
	GetRolloutN() *int32
	SetTotalSteps(v int32) *RLProgressConfig
	GetTotalSteps() *int32
	SetTrainBatchSize(v int32) *RLProgressConfig
	GetTrainBatchSize() *int32
}

type RLProgressConfig struct {
	// The number of mini-batches per step.
	//
	// example:
	//
	// 4
	NumMinibatches *int32 `json:"NumMinibatches,omitempty" xml:"NumMinibatches,omitempty"`
	// The PPO mini-batch size.
	//
	// example:
	//
	// 128
	PpoMiniBatchSize *int32 `json:"PpoMiniBatchSize,omitempty" xml:"PpoMiniBatchSize,omitempty"`
	// The number of rollouts per prompt.
	//
	// example:
	//
	// 8
	RolloutN *int32 `json:"RolloutN,omitempty" xml:"RolloutN,omitempty"`
	// The total number of training steps.
	//
	// example:
	//
	// 3
	TotalSteps *int32 `json:"TotalSteps,omitempty" xml:"TotalSteps,omitempty"`
	// The training batch size.
	//
	// example:
	//
	// 512
	TrainBatchSize *int32 `json:"TrainBatchSize,omitempty" xml:"TrainBatchSize,omitempty"`
}

func (s RLProgressConfig) String() string {
	return dara.Prettify(s)
}

func (s RLProgressConfig) GoString() string {
	return s.String()
}

func (s *RLProgressConfig) GetNumMinibatches() *int32 {
	return s.NumMinibatches
}

func (s *RLProgressConfig) GetPpoMiniBatchSize() *int32 {
	return s.PpoMiniBatchSize
}

func (s *RLProgressConfig) GetRolloutN() *int32 {
	return s.RolloutN
}

func (s *RLProgressConfig) GetTotalSteps() *int32 {
	return s.TotalSteps
}

func (s *RLProgressConfig) GetTrainBatchSize() *int32 {
	return s.TrainBatchSize
}

func (s *RLProgressConfig) SetNumMinibatches(v int32) *RLProgressConfig {
	s.NumMinibatches = &v
	return s
}

func (s *RLProgressConfig) SetPpoMiniBatchSize(v int32) *RLProgressConfig {
	s.PpoMiniBatchSize = &v
	return s
}

func (s *RLProgressConfig) SetRolloutN(v int32) *RLProgressConfig {
	s.RolloutN = &v
	return s
}

func (s *RLProgressConfig) SetTotalSteps(v int32) *RLProgressConfig {
	s.TotalSteps = &v
	return s
}

func (s *RLProgressConfig) SetTrainBatchSize(v int32) *RLProgressConfig {
	s.TrainBatchSize = &v
	return s
}

func (s *RLProgressConfig) Validate() error {
	return dara.Validate(s)
}
