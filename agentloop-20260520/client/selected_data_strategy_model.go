// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectedDataStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *SelectedDataStrategy
	GetEndTime() *int64
	SetEvalIds(v []*string) *SelectedDataStrategy
	GetEvalIds() []*string
	SetStartTime(v int64) *SelectedDataStrategy
	GetStartTime() *int64
}

type SelectedDataStrategy struct {
	// The end time of the statistical window in UNIX seconds.
	//
	// example:
	//
	// 1735660800
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The list of evaluation IDs.
	//
	// example:
	//
	// ["2646a9380bba3c86371f55a0226322941794983634"]
	EvalIds []*string `json:"evalIds,omitempty" xml:"evalIds,omitempty" type:"Repeated"`
	// The estimated start time in milliseconds.
	//
	// > You can modify this parameter only when the course is in the preview state.
	//
	// example:
	//
	// 1735574400
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s SelectedDataStrategy) String() string {
	return dara.Prettify(s)
}

func (s SelectedDataStrategy) GoString() string {
	return s.String()
}

func (s *SelectedDataStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SelectedDataStrategy) GetEvalIds() []*string {
	return s.EvalIds
}

func (s *SelectedDataStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SelectedDataStrategy) SetEndTime(v int64) *SelectedDataStrategy {
	s.EndTime = &v
	return s
}

func (s *SelectedDataStrategy) SetEvalIds(v []*string) *SelectedDataStrategy {
	s.EvalIds = v
	return s
}

func (s *SelectedDataStrategy) SetStartTime(v int64) *SelectedDataStrategy {
	s.StartTime = &v
	return s
}

func (s *SelectedDataStrategy) Validate() error {
	return dara.Validate(s)
}
