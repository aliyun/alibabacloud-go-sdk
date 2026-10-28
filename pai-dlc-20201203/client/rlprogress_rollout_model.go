// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressRollout interface {
	dara.Model
	String() string
	GoString() string
	SetFinished(v int32) *RLProgressRollout
	GetFinished() *int32
	SetProcessed(v *RLProgressProcessed) *RLProgressRollout
	GetProcessed() *RLProgressProcessed
	SetRatePerMin(v float64) *RLProgressRollout
	GetRatePerMin() *float64
}

type RLProgressRollout struct {
	// The total number of completed items within the window.
	//
	// example:
	//
	// 500
	Finished *int32 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The rollout processed count.
	//
	// if can be null:
	// true
	Processed *RLProgressProcessed `json:"Processed,omitempty" xml:"Processed,omitempty"`
	// The completion rate (items per minute), estimated from the last 120 completion events.
	//
	// example:
	//
	// 31.2
	RatePerMin *float64 `json:"RatePerMin,omitempty" xml:"RatePerMin,omitempty"`
}

func (s RLProgressRollout) String() string {
	return dara.Prettify(s)
}

func (s RLProgressRollout) GoString() string {
	return s.String()
}

func (s *RLProgressRollout) GetFinished() *int32 {
	return s.Finished
}

func (s *RLProgressRollout) GetProcessed() *RLProgressProcessed {
	return s.Processed
}

func (s *RLProgressRollout) GetRatePerMin() *float64 {
	return s.RatePerMin
}

func (s *RLProgressRollout) SetFinished(v int32) *RLProgressRollout {
	s.Finished = &v
	return s
}

func (s *RLProgressRollout) SetProcessed(v *RLProgressProcessed) *RLProgressRollout {
	s.Processed = v
	return s
}

func (s *RLProgressRollout) SetRatePerMin(v float64) *RLProgressRollout {
	s.RatePerMin = &v
	return s
}

func (s *RLProgressRollout) Validate() error {
	if s.Processed != nil {
		if err := s.Processed.Validate(); err != nil {
			return err
		}
	}
	return nil
}
