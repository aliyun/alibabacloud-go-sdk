// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobPlanCurrentStep(v string) *UpdateJobPlanRequest
	GetJobPlanCurrentStep() *string
	SetTag(v []*UpdateJobPlanRequestTag) *UpdateJobPlanRequest
	GetTag() []*UpdateJobPlanRequestTag
}

type UpdateJobPlanRequest struct {
	JobPlanCurrentStep *string                    `json:"JobPlanCurrentStep,omitempty" xml:"JobPlanCurrentStep,omitempty"`
	Tag                []*UpdateJobPlanRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UpdateJobPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobPlanRequest) GetJobPlanCurrentStep() *string {
	return s.JobPlanCurrentStep
}

func (s *UpdateJobPlanRequest) GetTag() []*UpdateJobPlanRequestTag {
	return s.Tag
}

func (s *UpdateJobPlanRequest) SetJobPlanCurrentStep(v string) *UpdateJobPlanRequest {
	s.JobPlanCurrentStep = &v
	return s
}

func (s *UpdateJobPlanRequest) SetTag(v []*UpdateJobPlanRequestTag) *UpdateJobPlanRequest {
	s.Tag = v
	return s
}

func (s *UpdateJobPlanRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateJobPlanRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateJobPlanRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobPlanRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateJobPlanRequestTag) GetKey() *string {
	return s.Key
}

func (s *UpdateJobPlanRequestTag) GetValue() *string {
	return s.Value
}

func (s *UpdateJobPlanRequestTag) SetKey(v string) *UpdateJobPlanRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateJobPlanRequestTag) SetValue(v string) *UpdateJobPlanRequestTag {
	s.Value = &v
	return s
}

func (s *UpdateJobPlanRequestTag) Validate() error {
	return dara.Validate(s)
}
