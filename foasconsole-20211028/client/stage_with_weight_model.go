// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStageWithWeight interface {
	dara.Model
	String() string
	GoString() string
	SetStepIndex(v int32) *StageWithWeight
	GetStepIndex() *int32
	SetStepName(v string) *StageWithWeight
	GetStepName() *string
	SetWeight(v int32) *StageWithWeight
	GetWeight() *int32
}

type StageWithWeight struct {
	StepIndex *int32  `json:"StepIndex,omitempty" xml:"StepIndex,omitempty"`
	StepName  *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	Weight    *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s StageWithWeight) String() string {
	return dara.Prettify(s)
}

func (s StageWithWeight) GoString() string {
	return s.String()
}

func (s *StageWithWeight) GetStepIndex() *int32 {
	return s.StepIndex
}

func (s *StageWithWeight) GetStepName() *string {
	return s.StepName
}

func (s *StageWithWeight) GetWeight() *int32 {
	return s.Weight
}

func (s *StageWithWeight) SetStepIndex(v int32) *StageWithWeight {
	s.StepIndex = &v
	return s
}

func (s *StageWithWeight) SetStepName(v string) *StageWithWeight {
	s.StepName = &v
	return s
}

func (s *StageWithWeight) SetWeight(v int32) *StageWithWeight {
	s.Weight = &v
	return s
}

func (s *StageWithWeight) Validate() error {
	return dara.Validate(s)
}
