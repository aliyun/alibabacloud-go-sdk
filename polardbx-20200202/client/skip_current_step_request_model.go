// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipCurrentStepRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentStep(v string) *SkipCurrentStepRequest
	GetCurrentStep() *string
	SetRegionId(v string) *SkipCurrentStepRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *SkipCurrentStepRequest
	GetSlinkTaskId() *string
}

type SkipCurrentStepRequest struct {
	// example:
	//
	// PRE_CHECK
	CurrentStep *string `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s SkipCurrentStepRequest) String() string {
	return dara.Prettify(s)
}

func (s SkipCurrentStepRequest) GoString() string {
	return s.String()
}

func (s *SkipCurrentStepRequest) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *SkipCurrentStepRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SkipCurrentStepRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *SkipCurrentStepRequest) SetCurrentStep(v string) *SkipCurrentStepRequest {
	s.CurrentStep = &v
	return s
}

func (s *SkipCurrentStepRequest) SetRegionId(v string) *SkipCurrentStepRequest {
	s.RegionId = &v
	return s
}

func (s *SkipCurrentStepRequest) SetSlinkTaskId(v string) *SkipCurrentStepRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *SkipCurrentStepRequest) Validate() error {
	return dara.Validate(s)
}
