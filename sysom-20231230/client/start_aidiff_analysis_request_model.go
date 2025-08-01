// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIDiffAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTask1(v *StartAIDiffAnalysisRequestTask1) *StartAIDiffAnalysisRequest
	GetTask1() *StartAIDiffAnalysisRequestTask1
	SetTask2(v *StartAIDiffAnalysisRequestTask2) *StartAIDiffAnalysisRequest
	GetTask2() *StartAIDiffAnalysisRequestTask2
}

type StartAIDiffAnalysisRequest struct {
	// This parameter is required.
	Task1 *StartAIDiffAnalysisRequestTask1 `json:"task1,omitempty" xml:"task1,omitempty" type:"Struct"`
	// This parameter is required.
	Task2 *StartAIDiffAnalysisRequestTask2 `json:"task2,omitempty" xml:"task2,omitempty" type:"Struct"`
}

func (s StartAIDiffAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIDiffAnalysisRequest) GoString() string {
	return s.String()
}

func (s *StartAIDiffAnalysisRequest) GetTask1() *StartAIDiffAnalysisRequestTask1 {
	return s.Task1
}

func (s *StartAIDiffAnalysisRequest) GetTask2() *StartAIDiffAnalysisRequestTask2 {
	return s.Task2
}

func (s *StartAIDiffAnalysisRequest) SetTask1(v *StartAIDiffAnalysisRequestTask1) *StartAIDiffAnalysisRequest {
	s.Task1 = v
	return s
}

func (s *StartAIDiffAnalysisRequest) SetTask2(v *StartAIDiffAnalysisRequestTask2) *StartAIDiffAnalysisRequest {
	s.Task2 = v
	return s
}

func (s *StartAIDiffAnalysisRequest) Validate() error {
	return dara.Validate(s)
}

type StartAIDiffAnalysisRequestTask1 struct {
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string   `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	Pids       []*string `json:"pids,omitempty" xml:"pids,omitempty" type:"Repeated"`
	// example:
	//
	// 4660551334179.955
	StepEnd *float32 `json:"step_end,omitempty" xml:"step_end,omitempty"`
	// example:
	//
	// 4660550379415.497
	StepStart *float32 `json:"step_start,omitempty" xml:"step_start,omitempty"`
}

func (s StartAIDiffAnalysisRequestTask1) String() string {
	return dara.Prettify(s)
}

func (s StartAIDiffAnalysisRequestTask1) GoString() string {
	return s.String()
}

func (s *StartAIDiffAnalysisRequestTask1) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *StartAIDiffAnalysisRequestTask1) GetPids() []*string {
	return s.Pids
}

func (s *StartAIDiffAnalysisRequestTask1) GetStepEnd() *float32 {
	return s.StepEnd
}

func (s *StartAIDiffAnalysisRequestTask1) GetStepStart() *float32 {
	return s.StepStart
}

func (s *StartAIDiffAnalysisRequestTask1) SetAnalysisId(v string) *StartAIDiffAnalysisRequestTask1 {
	s.AnalysisId = &v
	return s
}

func (s *StartAIDiffAnalysisRequestTask1) SetPids(v []*string) *StartAIDiffAnalysisRequestTask1 {
	s.Pids = v
	return s
}

func (s *StartAIDiffAnalysisRequestTask1) SetStepEnd(v float32) *StartAIDiffAnalysisRequestTask1 {
	s.StepEnd = &v
	return s
}

func (s *StartAIDiffAnalysisRequestTask1) SetStepStart(v float32) *StartAIDiffAnalysisRequestTask1 {
	s.StepStart = &v
	return s
}

func (s *StartAIDiffAnalysisRequestTask1) Validate() error {
	return dara.Validate(s)
}

type StartAIDiffAnalysisRequestTask2 struct {
	// This parameter is required.
	//
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 452651:python3 ./test.py
	Pids []*string `json:"pids,omitempty" xml:"pids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 4660551334179.955
	StepEnd *float32 `json:"step_end,omitempty" xml:"step_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4660550379415.497
	StepStart *float32 `json:"step_start,omitempty" xml:"step_start,omitempty"`
}

func (s StartAIDiffAnalysisRequestTask2) String() string {
	return dara.Prettify(s)
}

func (s StartAIDiffAnalysisRequestTask2) GoString() string {
	return s.String()
}

func (s *StartAIDiffAnalysisRequestTask2) GetAnalysisId() *string {
	return s.AnalysisId
}

func (s *StartAIDiffAnalysisRequestTask2) GetPids() []*string {
	return s.Pids
}

func (s *StartAIDiffAnalysisRequestTask2) GetStepEnd() *float32 {
	return s.StepEnd
}

func (s *StartAIDiffAnalysisRequestTask2) GetStepStart() *float32 {
	return s.StepStart
}

func (s *StartAIDiffAnalysisRequestTask2) SetAnalysisId(v string) *StartAIDiffAnalysisRequestTask2 {
	s.AnalysisId = &v
	return s
}

func (s *StartAIDiffAnalysisRequestTask2) SetPids(v []*string) *StartAIDiffAnalysisRequestTask2 {
	s.Pids = v
	return s
}

func (s *StartAIDiffAnalysisRequestTask2) SetStepEnd(v float32) *StartAIDiffAnalysisRequestTask2 {
	s.StepEnd = &v
	return s
}

func (s *StartAIDiffAnalysisRequestTask2) SetStepStart(v float32) *StartAIDiffAnalysisRequestTask2 {
	s.StepStart = &v
	return s
}

func (s *StartAIDiffAnalysisRequestTask2) Validate() error {
	return dara.Validate(s)
}
