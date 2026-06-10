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
	// task1 parameters
	//
	// This parameter is required.
	//
	// example:
	//
	// task1参数
	Task1 *StartAIDiffAnalysisRequestTask1 `json:"task1,omitempty" xml:"task1,omitempty" type:"Struct"`
	// task2 parameters
	//
	// This parameter is required.
	//
	// example:
	//
	// task2参数，目前只支持相同analysisId和pid的对比
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
	if s.Task1 != nil {
		if err := s.Task1.Validate(); err != nil {
			return err
		}
	}
	if s.Task2 != nil {
		if err := s.Task2.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartAIDiffAnalysisRequestTask1 struct {
	// AI analysis ID
	//
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	// PIDs of AI job processes; batch input is supported, separated by commas
	Pids []*string `json:"pids,omitempty" xml:"pids,omitempty" type:"Repeated"`
	// Step end time, computed based on the selected step number
	//
	// example:
	//
	// 4660551334179.955
	StepEnd *float32 `json:"step_end,omitempty" xml:"step_end,omitempty"`
	// Step start time, computed based on the selected step number
	//
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
	// AI analysis ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	// Process IDs (PIDs) of AI jobs. Batch input is supported, with PIDs separated by commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// 452651:python3 ./test.py
	Pids []*string `json:"pids,omitempty" xml:"pids,omitempty" type:"Repeated"`
	// Step end time, computed based on the selected step number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4660551334179.955
	StepEnd *float32 `json:"step_end,omitempty" xml:"step_end,omitempty"`
	// Step start time, computed based on the selected step number.
	//
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
