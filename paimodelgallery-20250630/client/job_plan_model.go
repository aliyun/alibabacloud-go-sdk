// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobPlan interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *JobPlan
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *JobPlan
	GetGmtModifiedTime() *string
	SetJobPlanCurrentStep(v string) *JobPlan
	GetJobPlanCurrentStep() *string
	SetJobPlanId(v string) *JobPlan
	GetJobPlanId() *string
	SetJobPlanName(v string) *JobPlan
	GetJobPlanName() *string
	SetJobPlanSteps(v []*JobPlanJobPlanSteps) *JobPlan
	GetJobPlanSteps() []*JobPlanJobPlanSteps
	SetJobPlanType(v string) *JobPlan
	GetJobPlanType() *string
	SetOwnerId(v string) *JobPlan
	GetOwnerId() *string
	SetTags(v []*JobPlanTags) *JobPlan
	GetTags() []*JobPlanTags
	SetTemplateId(v string) *JobPlan
	GetTemplateId() *string
	SetTemplateName(v string) *JobPlan
	GetTemplateName() *string
	SetUserId(v string) *JobPlan
	GetUserId() *string
	SetWorkspaceId(v string) *JobPlan
	GetWorkspaceId() *string
}

type JobPlan struct {
	GmtCreateTime      *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime    *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	JobPlanCurrentStep *string                `json:"JobPlanCurrentStep,omitempty" xml:"JobPlanCurrentStep,omitempty"`
	JobPlanId          *string                `json:"JobPlanId,omitempty" xml:"JobPlanId,omitempty"`
	JobPlanName        *string                `json:"JobPlanName,omitempty" xml:"JobPlanName,omitempty"`
	JobPlanSteps       []*JobPlanJobPlanSteps `json:"JobPlanSteps,omitempty" xml:"JobPlanSteps,omitempty" type:"Repeated"`
	JobPlanType        *string                `json:"JobPlanType,omitempty" xml:"JobPlanType,omitempty"`
	OwnerId            *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Tags               []*JobPlanTags         `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The distillation template ID used when creating the task plan. An empty value indicates that this is not a scenario-specific distillation task.
	//
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The display name of the distillation template used, localized based on the language specified in the request.
	//
	// example:
	//
	// Chain-of-Thought Reasoning Distillation
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId  *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s JobPlan) String() string {
	return dara.Prettify(s)
}

func (s JobPlan) GoString() string {
	return s.String()
}

func (s *JobPlan) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *JobPlan) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *JobPlan) GetJobPlanCurrentStep() *string {
	return s.JobPlanCurrentStep
}

func (s *JobPlan) GetJobPlanId() *string {
	return s.JobPlanId
}

func (s *JobPlan) GetJobPlanName() *string {
	return s.JobPlanName
}

func (s *JobPlan) GetJobPlanSteps() []*JobPlanJobPlanSteps {
	return s.JobPlanSteps
}

func (s *JobPlan) GetJobPlanType() *string {
	return s.JobPlanType
}

func (s *JobPlan) GetOwnerId() *string {
	return s.OwnerId
}

func (s *JobPlan) GetTags() []*JobPlanTags {
	return s.Tags
}

func (s *JobPlan) GetTemplateId() *string {
	return s.TemplateId
}

func (s *JobPlan) GetTemplateName() *string {
	return s.TemplateName
}

func (s *JobPlan) GetUserId() *string {
	return s.UserId
}

func (s *JobPlan) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *JobPlan) SetGmtCreateTime(v string) *JobPlan {
	s.GmtCreateTime = &v
	return s
}

func (s *JobPlan) SetGmtModifiedTime(v string) *JobPlan {
	s.GmtModifiedTime = &v
	return s
}

func (s *JobPlan) SetJobPlanCurrentStep(v string) *JobPlan {
	s.JobPlanCurrentStep = &v
	return s
}

func (s *JobPlan) SetJobPlanId(v string) *JobPlan {
	s.JobPlanId = &v
	return s
}

func (s *JobPlan) SetJobPlanName(v string) *JobPlan {
	s.JobPlanName = &v
	return s
}

func (s *JobPlan) SetJobPlanSteps(v []*JobPlanJobPlanSteps) *JobPlan {
	s.JobPlanSteps = v
	return s
}

func (s *JobPlan) SetJobPlanType(v string) *JobPlan {
	s.JobPlanType = &v
	return s
}

func (s *JobPlan) SetOwnerId(v string) *JobPlan {
	s.OwnerId = &v
	return s
}

func (s *JobPlan) SetTags(v []*JobPlanTags) *JobPlan {
	s.Tags = v
	return s
}

func (s *JobPlan) SetTemplateId(v string) *JobPlan {
	s.TemplateId = &v
	return s
}

func (s *JobPlan) SetTemplateName(v string) *JobPlan {
	s.TemplateName = &v
	return s
}

func (s *JobPlan) SetUserId(v string) *JobPlan {
	s.UserId = &v
	return s
}

func (s *JobPlan) SetWorkspaceId(v string) *JobPlan {
	s.WorkspaceId = &v
	return s
}

func (s *JobPlan) Validate() error {
	if s.JobPlanSteps != nil {
		for _, item := range s.JobPlanSteps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type JobPlanJobPlanSteps struct {
	JobPlanStepId   *string                `json:"JobPlanStepId,omitempty" xml:"JobPlanStepId,omitempty"`
	JobPlanStepName *string                `json:"JobPlanStepName,omitempty" xml:"JobPlanStepName,omitempty"`
	JobPlanStepSpec map[string]interface{} `json:"JobPlanStepSpec,omitempty" xml:"JobPlanStepSpec,omitempty"`
	JobPlanStepType *string                `json:"JobPlanStepType,omitempty" xml:"JobPlanStepType,omitempty"`
}

func (s JobPlanJobPlanSteps) String() string {
	return dara.Prettify(s)
}

func (s JobPlanJobPlanSteps) GoString() string {
	return s.String()
}

func (s *JobPlanJobPlanSteps) GetJobPlanStepId() *string {
	return s.JobPlanStepId
}

func (s *JobPlanJobPlanSteps) GetJobPlanStepName() *string {
	return s.JobPlanStepName
}

func (s *JobPlanJobPlanSteps) GetJobPlanStepSpec() map[string]interface{} {
	return s.JobPlanStepSpec
}

func (s *JobPlanJobPlanSteps) GetJobPlanStepType() *string {
	return s.JobPlanStepType
}

func (s *JobPlanJobPlanSteps) SetJobPlanStepId(v string) *JobPlanJobPlanSteps {
	s.JobPlanStepId = &v
	return s
}

func (s *JobPlanJobPlanSteps) SetJobPlanStepName(v string) *JobPlanJobPlanSteps {
	s.JobPlanStepName = &v
	return s
}

func (s *JobPlanJobPlanSteps) SetJobPlanStepSpec(v map[string]interface{}) *JobPlanJobPlanSteps {
	s.JobPlanStepSpec = v
	return s
}

func (s *JobPlanJobPlanSteps) SetJobPlanStepType(v string) *JobPlanJobPlanSteps {
	s.JobPlanStepType = &v
	return s
}

func (s *JobPlanJobPlanSteps) Validate() error {
	return dara.Validate(s)
}

type JobPlanTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s JobPlanTags) String() string {
	return dara.Prettify(s)
}

func (s JobPlanTags) GoString() string {
	return s.String()
}

func (s *JobPlanTags) GetKey() *string {
	return s.Key
}

func (s *JobPlanTags) GetValue() *string {
	return s.Value
}

func (s *JobPlanTags) SetKey(v string) *JobPlanTags {
	s.Key = &v
	return s
}

func (s *JobPlanTags) SetValue(v string) *JobPlanTags {
	s.Value = &v
	return s
}

func (s *JobPlanTags) Validate() error {
	return dara.Validate(s)
}
