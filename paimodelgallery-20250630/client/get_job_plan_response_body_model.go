// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *GetJobPlanResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetJobPlanResponseBody
	GetGmtModifiedTime() *string
	SetJobPlanCurrentStep(v string) *GetJobPlanResponseBody
	GetJobPlanCurrentStep() *string
	SetJobPlanId(v string) *GetJobPlanResponseBody
	GetJobPlanId() *string
	SetJobPlanName(v string) *GetJobPlanResponseBody
	GetJobPlanName() *string
	SetJobPlanSteps(v []*GetJobPlanResponseBodyJobPlanSteps) *GetJobPlanResponseBody
	GetJobPlanSteps() []*GetJobPlanResponseBodyJobPlanSteps
	SetJobPlanType(v string) *GetJobPlanResponseBody
	GetJobPlanType() *string
	SetOwnerId(v string) *GetJobPlanResponseBody
	GetOwnerId() *string
	SetRequestId(v string) *GetJobPlanResponseBody
	GetRequestId() *string
	SetTags(v []*GetJobPlanResponseBodyTags) *GetJobPlanResponseBody
	GetTags() []*GetJobPlanResponseBodyTags
	SetTemplateId(v string) *GetJobPlanResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetJobPlanResponseBody
	GetTemplateName() *string
	SetUserId(v string) *GetJobPlanResponseBody
	GetUserId() *string
	SetWorkspaceId(v string) *GetJobPlanResponseBody
	GetWorkspaceId() *string
}

type GetJobPlanResponseBody struct {
	// The creation time of the job plan.
	//
	// example:
	//
	// 2026-09-18 10:00:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modified time of the job plan.
	//
	// example:
	//
	// 2026-09-18 10:30:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The current step.
	//
	// example:
	//
	// DatasetSynthesis
	JobPlanCurrentStep *string `json:"JobPlanCurrentStep,omitempty" xml:"JobPlanCurrentStep,omitempty"`
	// The job plan ID.
	//
	// example:
	//
	// jp-xxxxxx
	JobPlanId *string `json:"JobPlanId,omitempty" xml:"JobPlanId,omitempty"`
	// The job plan name, which is unique within the workspace.
	//
	// example:
	//
	// ModelGalleryxxx
	JobPlanName *string `json:"JobPlanName,omitempty" xml:"JobPlanName,omitempty"`
	// The job plan steps.
	JobPlanSteps []*GetJobPlanResponseBodyJobPlanSteps `json:"JobPlanSteps,omitempty" xml:"JobPlanSteps,omitempty" type:"Repeated"`
	// The job plan type.
	//
	// example:
	//
	// Distillation
	JobPlanType *string `json:"JobPlanType,omitempty" xml:"JobPlanType,omitempty"`
	// The Alibaba Cloud account ID that owns the job plan.
	//
	// example:
	//
	// 1234567890123456
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82-9624-EC2B1779848E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tags.
	Tags []*GetJobPlanResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The distillation template ID used to create the job plan. An empty value indicates that this is not a scenario-based distillation task.
	//
	// example:
	//
	// advanced_cot_distill
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The display name of the distillation template, localized based on the requested language.
	//
	// example:
	//
	// 思维链推理蒸馏
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The Alibaba Cloud account ID that created the job plan.
	//
	// example:
	//
	// 1234567890123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 62469
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetJobPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobPlanResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobPlanResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetJobPlanResponseBody) GetJobPlanCurrentStep() *string {
	return s.JobPlanCurrentStep
}

func (s *GetJobPlanResponseBody) GetJobPlanId() *string {
	return s.JobPlanId
}

func (s *GetJobPlanResponseBody) GetJobPlanName() *string {
	return s.JobPlanName
}

func (s *GetJobPlanResponseBody) GetJobPlanSteps() []*GetJobPlanResponseBodyJobPlanSteps {
	return s.JobPlanSteps
}

func (s *GetJobPlanResponseBody) GetJobPlanType() *string {
	return s.JobPlanType
}

func (s *GetJobPlanResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetJobPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobPlanResponseBody) GetTags() []*GetJobPlanResponseBodyTags {
	return s.Tags
}

func (s *GetJobPlanResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetJobPlanResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetJobPlanResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetJobPlanResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetJobPlanResponseBody) SetGmtCreateTime(v string) *GetJobPlanResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobPlanResponseBody) SetGmtModifiedTime(v string) *GetJobPlanResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetJobPlanResponseBody) SetJobPlanCurrentStep(v string) *GetJobPlanResponseBody {
	s.JobPlanCurrentStep = &v
	return s
}

func (s *GetJobPlanResponseBody) SetJobPlanId(v string) *GetJobPlanResponseBody {
	s.JobPlanId = &v
	return s
}

func (s *GetJobPlanResponseBody) SetJobPlanName(v string) *GetJobPlanResponseBody {
	s.JobPlanName = &v
	return s
}

func (s *GetJobPlanResponseBody) SetJobPlanSteps(v []*GetJobPlanResponseBodyJobPlanSteps) *GetJobPlanResponseBody {
	s.JobPlanSteps = v
	return s
}

func (s *GetJobPlanResponseBody) SetJobPlanType(v string) *GetJobPlanResponseBody {
	s.JobPlanType = &v
	return s
}

func (s *GetJobPlanResponseBody) SetOwnerId(v string) *GetJobPlanResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetJobPlanResponseBody) SetRequestId(v string) *GetJobPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobPlanResponseBody) SetTags(v []*GetJobPlanResponseBodyTags) *GetJobPlanResponseBody {
	s.Tags = v
	return s
}

func (s *GetJobPlanResponseBody) SetTemplateId(v string) *GetJobPlanResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetJobPlanResponseBody) SetTemplateName(v string) *GetJobPlanResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetJobPlanResponseBody) SetUserId(v string) *GetJobPlanResponseBody {
	s.UserId = &v
	return s
}

func (s *GetJobPlanResponseBody) SetWorkspaceId(v string) *GetJobPlanResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobPlanResponseBody) Validate() error {
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

type GetJobPlanResponseBodyJobPlanSteps struct {
	// The generated PAIFlow PipelineRunId.
	//
	// example:
	//
	// pipeline-xxxxx
	JobPlanStepId *string `json:"JobPlanStepId,omitempty" xml:"JobPlanStepId,omitempty"`
	// The job plan step name.
	//
	// example:
	//
	// DistillationDatasetSynthesis
	JobPlanStepName *string `json:"JobPlanStepName,omitempty" xml:"JobPlanStepName,omitempty"`
	// The job plan step configuration.
	JobPlanStepSpec map[string]interface{} `json:"JobPlanStepSpec,omitempty" xml:"JobPlanStepSpec,omitempty"`
	// The job plan step type.
	//
	// example:
	//
	// PAIFlow
	JobPlanStepType *string `json:"JobPlanStepType,omitempty" xml:"JobPlanStepType,omitempty"`
}

func (s GetJobPlanResponseBodyJobPlanSteps) String() string {
	return dara.Prettify(s)
}

func (s GetJobPlanResponseBodyJobPlanSteps) GoString() string {
	return s.String()
}

func (s *GetJobPlanResponseBodyJobPlanSteps) GetJobPlanStepId() *string {
	return s.JobPlanStepId
}

func (s *GetJobPlanResponseBodyJobPlanSteps) GetJobPlanStepName() *string {
	return s.JobPlanStepName
}

func (s *GetJobPlanResponseBodyJobPlanSteps) GetJobPlanStepSpec() map[string]interface{} {
	return s.JobPlanStepSpec
}

func (s *GetJobPlanResponseBodyJobPlanSteps) GetJobPlanStepType() *string {
	return s.JobPlanStepType
}

func (s *GetJobPlanResponseBodyJobPlanSteps) SetJobPlanStepId(v string) *GetJobPlanResponseBodyJobPlanSteps {
	s.JobPlanStepId = &v
	return s
}

func (s *GetJobPlanResponseBodyJobPlanSteps) SetJobPlanStepName(v string) *GetJobPlanResponseBodyJobPlanSteps {
	s.JobPlanStepName = &v
	return s
}

func (s *GetJobPlanResponseBodyJobPlanSteps) SetJobPlanStepSpec(v map[string]interface{}) *GetJobPlanResponseBodyJobPlanSteps {
	s.JobPlanStepSpec = v
	return s
}

func (s *GetJobPlanResponseBodyJobPlanSteps) SetJobPlanStepType(v string) *GetJobPlanResponseBodyJobPlanSteps {
	s.JobPlanStepType = &v
	return s
}

func (s *GetJobPlanResponseBodyJobPlanSteps) Validate() error {
	return dara.Validate(s)
}

type GetJobPlanResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// bar
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetJobPlanResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetJobPlanResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetJobPlanResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetJobPlanResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetJobPlanResponseBodyTags) SetKey(v string) *GetJobPlanResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetJobPlanResponseBodyTags) SetValue(v string) *GetJobPlanResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetJobPlanResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
