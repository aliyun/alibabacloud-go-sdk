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
	// Specifies whether to run the workflow instance immediately after it is created. Default value: true.
	//
	// example:
	//
	// true
	AutoStartEnabled *bool `json:"AutoStartEnabled,omitempty" xml:"AutoStartEnabled,omitempty"`
	// The reason for creation.
	//
	// example:
	//
	// create for test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The runtime configurations.
	DefaultRunPropertiesShrink *string `json:"DefaultRunProperties,omitempty" xml:"DefaultRunProperties,omitempty"`
	// The project environment. Valid values:
	//
	// - Prod: production
	//
	// - Dev: development
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
	// The data backfill period settings.
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
	// - Append: Append mode. Inherits existing tags from the manual workflow and appends new tags.
	//
	// - Overwrite: Overwrite mode. Does not inherit existing tags from the manual workflow and directly creates tags.
	//
	// example:
	//
	// Append
	TagCreationPolicy *string `json:"TagCreationPolicy,omitempty" xml:"TagCreationPolicy,omitempty"`
	// The list of task tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The task parameters, which are used to set parameters for specific tasks. The value is in JSON format. The key is the task ID, and the value format refers to the task script parameters (Task.Script.Parameter in the response of the GetTask operation).
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
	// - SupplementData: data backfill. The methods for specifying RootTaskIds and IncludeTaskIds vary based on the data backfill mode. For more information, see the parameter description of DefaultRunProperties.Mode.
	//
	// - ManualWorkflow: manually triggered workflow. You must set WorkflowId to the ID of the manually triggered workflow. RootTaskIds is optional. If you do not specify RootTaskIds, the default root node list of the manually triggered workflow is used.
	//
	// - Manual: manually triggered task. You only need to specify RootTaskIds, which is the list of manually triggered tasks to be run.
	//
	// - SmokeTest: smoke test. You only need to specify RootTaskIds, which is the list of test tasks to be run.
	//
	// - TriggerWorkflow: event-triggered workflow. You must set WorkflowId to the ID of the event-triggered workflow. IncludeTaskIds is optional. If you do not specify IncludeTaskIds, the entire workflow is run.
	//
	// This parameter is required.
	//
	// example:
	//
	// SupplementData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the workflow. The WorkflowId of a scheduled task is always 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	// The workflow parameters. This parameter takes effect only when a specific workflow is specified (WorkflowId != 1). For a periodic workflow or an event-triggered workflow, the format is key=value, and the priority is lower than that of task parameters. For a manually triggered workflow, the format is JSON, and the priority is higher than that of task parameters.
	//
	// example:
	//
	// "key=value" format:
	//
	// key1=value1 key2=value2
	//
	// JSON format:
	//
	// {"key1":"value1", "key2": "value2"}
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
