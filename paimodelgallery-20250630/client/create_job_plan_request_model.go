// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobPlanName(v string) *CreateJobPlanRequest
	GetJobPlanName() *string
	SetJobPlanSteps(v []*CreateJobPlanRequestJobPlanSteps) *CreateJobPlanRequest
	GetJobPlanSteps() []*CreateJobPlanRequestJobPlanSteps
	SetJobPlanType(v string) *CreateJobPlanRequest
	GetJobPlanType() *string
	SetTag(v []*CreateJobPlanRequestTag) *CreateJobPlanRequest
	GetTag() []*CreateJobPlanRequestTag
	SetTemplateId(v string) *CreateJobPlanRequest
	GetTemplateId() *string
	SetWorkspaceId(v string) *CreateJobPlanRequest
	GetWorkspaceId() *string
}

type CreateJobPlanRequest struct {
	// The name of the job plan.
	//
	// example:
	//
	// ModelGalleryxxx
	JobPlanName *string `json:"JobPlanName,omitempty" xml:"JobPlanName,omitempty"`
	// The steps of the job plan.
	JobPlanSteps []*CreateJobPlanRequestJobPlanSteps `json:"JobPlanSteps,omitempty" xml:"JobPlanSteps,omitempty" type:"Repeated"`
	// The type of the job plan.
	//
	// example:
	//
	// Distillation
	JobPlanType *string `json:"JobPlanType,omitempty" xml:"JobPlanType,omitempty"`
	// Note: According to the Alibaba Cloud tag system specification, this parameter name is in singular form.
	Tag []*CreateJobPlanRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the scenario-specific distillation template, obtained from ListDistillationTemplates. If this parameter is not specified, a general-purpose job plan is created.
	//
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 62469
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateJobPlanRequest) GetJobPlanName() *string {
	return s.JobPlanName
}

func (s *CreateJobPlanRequest) GetJobPlanSteps() []*CreateJobPlanRequestJobPlanSteps {
	return s.JobPlanSteps
}

func (s *CreateJobPlanRequest) GetJobPlanType() *string {
	return s.JobPlanType
}

func (s *CreateJobPlanRequest) GetTag() []*CreateJobPlanRequestTag {
	return s.Tag
}

func (s *CreateJobPlanRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateJobPlanRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateJobPlanRequest) SetJobPlanName(v string) *CreateJobPlanRequest {
	s.JobPlanName = &v
	return s
}

func (s *CreateJobPlanRequest) SetJobPlanSteps(v []*CreateJobPlanRequestJobPlanSteps) *CreateJobPlanRequest {
	s.JobPlanSteps = v
	return s
}

func (s *CreateJobPlanRequest) SetJobPlanType(v string) *CreateJobPlanRequest {
	s.JobPlanType = &v
	return s
}

func (s *CreateJobPlanRequest) SetTag(v []*CreateJobPlanRequestTag) *CreateJobPlanRequest {
	s.Tag = v
	return s
}

func (s *CreateJobPlanRequest) SetTemplateId(v string) *CreateJobPlanRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateJobPlanRequest) SetWorkspaceId(v string) *CreateJobPlanRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateJobPlanRequest) Validate() error {
	if s.JobPlanSteps != nil {
		for _, item := range s.JobPlanSteps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateJobPlanRequestJobPlanSteps struct {
	// The name of the job plan step.
	//
	// example:
	//
	// DatasetSynthesis
	JobPlanStepName *string `json:"JobPlanStepName,omitempty" xml:"JobPlanStepName,omitempty"`
	// The detailed configuration of the job plan step.
	JobPlanStepSpec map[string]interface{} `json:"JobPlanStepSpec,omitempty" xml:"JobPlanStepSpec,omitempty"`
	// The type of the job plan step.
	//
	// example:
	//
	// PAIFlow
	JobPlanStepType *string `json:"JobPlanStepType,omitempty" xml:"JobPlanStepType,omitempty"`
}

func (s CreateJobPlanRequestJobPlanSteps) String() string {
	return dara.Prettify(s)
}

func (s CreateJobPlanRequestJobPlanSteps) GoString() string {
	return s.String()
}

func (s *CreateJobPlanRequestJobPlanSteps) GetJobPlanStepName() *string {
	return s.JobPlanStepName
}

func (s *CreateJobPlanRequestJobPlanSteps) GetJobPlanStepSpec() map[string]interface{} {
	return s.JobPlanStepSpec
}

func (s *CreateJobPlanRequestJobPlanSteps) GetJobPlanStepType() *string {
	return s.JobPlanStepType
}

func (s *CreateJobPlanRequestJobPlanSteps) SetJobPlanStepName(v string) *CreateJobPlanRequestJobPlanSteps {
	s.JobPlanStepName = &v
	return s
}

func (s *CreateJobPlanRequestJobPlanSteps) SetJobPlanStepSpec(v map[string]interface{}) *CreateJobPlanRequestJobPlanSteps {
	s.JobPlanStepSpec = v
	return s
}

func (s *CreateJobPlanRequestJobPlanSteps) SetJobPlanStepType(v string) *CreateJobPlanRequestJobPlanSteps {
	s.JobPlanStepType = &v
	return s
}

func (s *CreateJobPlanRequestJobPlanSteps) Validate() error {
	return dara.Validate(s)
}

type CreateJobPlanRequestTag struct {
	// **Key**
	//
	// example:
	//
	// foo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// **Value**
	//
	// example:
	//
	// bar
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateJobPlanRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateJobPlanRequestTag) GoString() string {
	return s.String()
}

func (s *CreateJobPlanRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateJobPlanRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateJobPlanRequestTag) SetKey(v string) *CreateJobPlanRequestTag {
	s.Key = &v
	return s
}

func (s *CreateJobPlanRequestTag) SetValue(v string) *CreateJobPlanRequestTag {
	s.Value = &v
	return s
}

func (s *CreateJobPlanRequestTag) Validate() error {
	return dara.Validate(s)
}
