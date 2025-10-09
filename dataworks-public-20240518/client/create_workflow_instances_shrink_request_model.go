// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStartEnabled(v bool) *CreateWorkflowInstancesShrinkRequest
	GetAutoStartEnabled() *bool
	SetComment(v string) *CreateWorkflowInstancesShrinkRequest
	GetComment() *string
	SetDefaultRunPropertiesShrink(v string) *CreateWorkflowInstancesShrinkRequest
	GetDefaultRunPropertiesShrink() *string
	SetEnvType(v string) *CreateWorkflowInstancesShrinkRequest
	GetEnvType() *string
	SetName(v string) *CreateWorkflowInstancesShrinkRequest
	GetName() *string
	SetPeriodsShrink(v string) *CreateWorkflowInstancesShrinkRequest
	GetPeriodsShrink() *string
	SetProjectId(v int64) *CreateWorkflowInstancesShrinkRequest
	GetProjectId() *int64
	SetTagCreationPolicy(v string) *CreateWorkflowInstancesShrinkRequest
	GetTagCreationPolicy() *string
	SetTagsShrink(v string) *CreateWorkflowInstancesShrinkRequest
	GetTagsShrink() *string
	SetTaskParameters(v string) *CreateWorkflowInstancesShrinkRequest
	GetTaskParameters() *string
	SetType(v string) *CreateWorkflowInstancesShrinkRequest
	GetType() *string
	SetWorkflowId(v int64) *CreateWorkflowInstancesShrinkRequest
	GetWorkflowId() *int64
	SetWorkflowParameters(v string) *CreateWorkflowInstancesShrinkRequest
	GetWorkflowParameters() *string
}

type CreateWorkflowInstancesShrinkRequest struct {
	// The default value is true.
	//
	// example:
	//
	// true
	AutoStartEnabled *bool `json:"AutoStartEnabled,omitempty" xml:"AutoStartEnabled,omitempty"`
	// The reason for the creation.
	//
	// example:
	//
	// create for test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The runtime configuration.
	DefaultRunPropertiesShrink *string `json:"DefaultRunProperties,omitempty" xml:"DefaultRunProperties,omitempty"`
	// The project environment.
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// WorkflowInstance1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration of the data backfilling period.
	PeriodsShrink *string `json:"Periods,omitempty" xml:"Periods,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tag creation policy. Valid values:
	//
	// 	- Append: New tags are added on top of the existing tags of the manual workflow.
	//
	// 	- Overwrite: Existing tags of the manual workflow are not inherited. New tags are created directly.
	//
	// example:
	//
	// Append
	TagCreationPolicy *string `json:"TagCreationPolicy,omitempty" xml:"TagCreationPolicy,omitempty"`
	// The task tag list.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The task-specific parameters. The value is in the JSON format. The key specifies the task ID. You can call the GetTask operation to obtain the format of the value by querying the script parameters.
	//
	// example:
	//
	// {
	//
	//   "1001": "key1=val2 key2=val2",
	//
	//   "1002": "key1=val2 key2=val2"
	//
	// }
	TaskParameters *string `json:"TaskParameters,omitempty" xml:"TaskParameters,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// 	- SupplementData: Data backfill. The usage of RootTaskIds and IncludeTaskIds varies based on the backfill mode. See the description of the DefaultRunProperties.Mode parameter.
	//
	// 	- ManualWorkflow: Manual workflow. WorkflowId is required for a manual workflow. RootTaskIds is optional. If not specified, the system uses the default root task list of the manual workflow.
	//
	// 	- Manual: Manual task. You only need to specify RootTaskIds. This is the list of manual tasks to run.
	//
	// 	- SmokeTest: Smoke test. You only need to specify RootTaskIds. This is the list of test tasks to run.
	//
	// This parameter is required.
	//
	// example:
	//
	// SupplementData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow to which the instance belongs. This parameter is set to 1 for auto triggered tasks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters. This parameter takes effect only when you set the `WorkflowId` parameter to a value other than 1. If your workflow is an auto triggered workflow, configure this parameter in the key=value format. The parameters that you configure in this parameter have a lower priority than task parameters. If your workflow is a manually triggered workflow, configure this parameter in the JSON format. The parameters that you configure in this parameter have a higher priority than task parameters.
	//
	// example:
	//
	// {
	//
	//   "key1": "value1",
	//
	//   "key2": "value2"
	//
	// }
	WorkflowParameters *string `json:"WorkflowParameters,omitempty" xml:"WorkflowParameters,omitempty"`
}

func (s CreateWorkflowInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstancesShrinkRequest) GetAutoStartEnabled() *bool {
	return s.AutoStartEnabled
}

func (s *CreateWorkflowInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateWorkflowInstancesShrinkRequest) GetDefaultRunPropertiesShrink() *string {
	return s.DefaultRunPropertiesShrink
}

func (s *CreateWorkflowInstancesShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateWorkflowInstancesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkflowInstancesShrinkRequest) GetPeriodsShrink() *string {
	return s.PeriodsShrink
}

func (s *CreateWorkflowInstancesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateWorkflowInstancesShrinkRequest) GetTagCreationPolicy() *string {
	return s.TagCreationPolicy
}

func (s *CreateWorkflowInstancesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateWorkflowInstancesShrinkRequest) GetTaskParameters() *string {
	return s.TaskParameters
}

func (s *CreateWorkflowInstancesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateWorkflowInstancesShrinkRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *CreateWorkflowInstancesShrinkRequest) GetWorkflowParameters() *string {
	return s.WorkflowParameters
}

func (s *CreateWorkflowInstancesShrinkRequest) SetAutoStartEnabled(v bool) *CreateWorkflowInstancesShrinkRequest {
	s.AutoStartEnabled = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetComment(v string) *CreateWorkflowInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetDefaultRunPropertiesShrink(v string) *CreateWorkflowInstancesShrinkRequest {
	s.DefaultRunPropertiesShrink = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetEnvType(v string) *CreateWorkflowInstancesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetName(v string) *CreateWorkflowInstancesShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetPeriodsShrink(v string) *CreateWorkflowInstancesShrinkRequest {
	s.PeriodsShrink = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetProjectId(v int64) *CreateWorkflowInstancesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetTagCreationPolicy(v string) *CreateWorkflowInstancesShrinkRequest {
	s.TagCreationPolicy = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetTagsShrink(v string) *CreateWorkflowInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetTaskParameters(v string) *CreateWorkflowInstancesShrinkRequest {
	s.TaskParameters = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetType(v string) *CreateWorkflowInstancesShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetWorkflowId(v int64) *CreateWorkflowInstancesShrinkRequest {
	s.WorkflowId = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) SetWorkflowParameters(v string) *CreateWorkflowInstancesShrinkRequest {
	s.WorkflowParameters = &v
	return s
}

func (s *CreateWorkflowInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
