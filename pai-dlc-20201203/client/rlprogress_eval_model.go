// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRLProgressEval interface {
	dara.Model
	String() string
	GoString() string
	SetDone(v bool) *RLProgressEval
	GetDone() *bool
	SetFinished(v int32) *RLProgressEval
	GetFinished() *int32
	SetPct(v float64) *RLProgressEval
	GetPct() *float64
	SetProgress(v int32) *RLProgressEval
	GetProgress() *int32
	SetReady(v int32) *RLProgressEval
	GetReady() *int32
	SetTotal(v int32) *RLProgressEval
	GetTotal() *int32
}

type RLProgressEval struct {
	// Indicates whether Total is greater than 0 and Finished is not less than Total.
	//
	// example:
	//
	// true
	Done *bool `json:"Done,omitempty" xml:"Done,omitempty"`
	// The number of finished samples.
	//
	// example:
	//
	// 500
	Finished *int32 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// The progress percentage, which is the ratio of Progress to Total.
	//
	// example:
	//
	// 100
	Pct *float64 `json:"Pct,omitempty" xml:"Pct,omitempty"`
	// The progress count, which is the greater value of Ready and Finished.
	//
	// example:
	//
	// 500
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The number of ready samples.
	//
	// example:
	//
	// 500
	Ready *int32 `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// The target number of samples.
	//
	// example:
	//
	// 3
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s RLProgressEval) String() string {
	return dara.Prettify(s)
}

func (s RLProgressEval) GoString() string {
	return s.String()
}

func (s *RLProgressEval) GetDone() *bool {
	return s.Done
}

func (s *RLProgressEval) GetFinished() *int32 {
	return s.Finished
}

func (s *RLProgressEval) GetPct() *float64 {
	return s.Pct
}

func (s *RLProgressEval) GetProgress() *int32 {
	return s.Progress
}

func (s *RLProgressEval) GetReady() *int32 {
	return s.Ready
}

func (s *RLProgressEval) GetTotal() *int32 {
	return s.Total
}

func (s *RLProgressEval) SetDone(v bool) *RLProgressEval {
	s.Done = &v
	return s
}

func (s *RLProgressEval) SetFinished(v int32) *RLProgressEval {
	s.Finished = &v
	return s
}

func (s *RLProgressEval) SetPct(v float64) *RLProgressEval {
	s.Pct = &v
	return s
}

func (s *RLProgressEval) SetProgress(v int32) *RLProgressEval {
	s.Progress = &v
	return s
}

func (s *RLProgressEval) SetReady(v int32) *RLProgressEval {
	s.Ready = &v
	return s
}

func (s *RLProgressEval) SetTotal(v int32) *RLProgressEval {
	s.Total = &v
	return s
}

func (s *RLProgressEval) Validate() error {
	return dara.Validate(s)
}
