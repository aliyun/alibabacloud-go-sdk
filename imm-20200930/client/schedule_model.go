// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchedule interface {
	dara.Model
	String() string
	GoString() string
	SetGamma(v float32) *Schedule
	GetGamma() *float32
	SetLRScheduler(v string) *Schedule
	GetLRScheduler() *string
	SetStepSize(v int64) *Schedule
	GetStepSize() *int64
}

type Schedule struct {
	// example:
	//
	// 0.97
	Gamma *float32 `json:"Gamma,omitempty" xml:"Gamma,omitempty"`
	// example:
	//
	// StepLR
	LRScheduler *string `json:"LRScheduler,omitempty" xml:"LRScheduler,omitempty"`
	// example:
	//
	// 1
	StepSize *int64 `json:"StepSize,omitempty" xml:"StepSize,omitempty"`
}

func (s Schedule) String() string {
	return dara.Prettify(s)
}

func (s Schedule) GoString() string {
	return s.String()
}

func (s *Schedule) GetGamma() *float32 {
	return s.Gamma
}

func (s *Schedule) GetLRScheduler() *string {
	return s.LRScheduler
}

func (s *Schedule) GetStepSize() *int64 {
	return s.StepSize
}

func (s *Schedule) SetGamma(v float32) *Schedule {
	s.Gamma = &v
	return s
}

func (s *Schedule) SetLRScheduler(v string) *Schedule {
	s.LRScheduler = &v
	return s
}

func (s *Schedule) SetStepSize(v int64) *Schedule {
	s.StepSize = &v
	return s
}

func (s *Schedule) Validate() error {
	return dara.Validate(s)
}
