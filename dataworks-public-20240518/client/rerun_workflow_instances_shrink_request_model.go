// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v int64) *RerunWorkflowInstancesShrinkRequest
	GetBizdate() *int64
	SetEndTriggerTime(v int64) *RerunWorkflowInstancesShrinkRequest
	GetEndTriggerTime() *int64
	SetEnvType(v string) *RerunWorkflowInstancesShrinkRequest
	GetEnvType() *string
	SetFilterShrink(v string) *RerunWorkflowInstancesShrinkRequest
	GetFilterShrink() *string
	SetIdsShrink(v string) *RerunWorkflowInstancesShrinkRequest
	GetIdsShrink() *string
	SetName(v string) *RerunWorkflowInstancesShrinkRequest
	GetName() *string
	SetProjectId(v int64) *RerunWorkflowInstancesShrinkRequest
	GetProjectId() *int64
	SetStartTriggerTime(v int64) *RerunWorkflowInstancesShrinkRequest
	GetStartTriggerTime() *int64
	SetStatus(v string) *RerunWorkflowInstancesShrinkRequest
	GetStatus() *string
	SetType(v string) *RerunWorkflowInstancesShrinkRequest
	GetType() *string
	SetWorkflowId(v int64) *RerunWorkflowInstancesShrinkRequest
	GetWorkflowId() *int64
}

type RerunWorkflowInstancesShrinkRequest struct {
	// The business date used for matching manual workflow instances.
	//
	// example:
	//
	// 1710239005403
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The end trigger time of the manual workflow instance used for matching. This parameter must be used together with the StartTriggerTime.
	//
	// example:
	//
	// 1710239005403
	EndTriggerTime *int64 `json:"EndTriggerTime,omitempty" xml:"EndTriggerTime,omitempty"`
	// The environment of the workspace. Valid values:
	//
	// Prod Dev
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The match conditions for internal instances of manual workflow instances.
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The instance IDs used for matching manual workflow instances.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The manual workflow name, used for fuzzy matching.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The start trigger time (creation time) of the manual workflow instance used for matching. This parameter must be used together with EndTriggerTime.
	//
	// example:
	//
	// 1710239005403
	StartTriggerTime *int64 `json:"StartTriggerTime,omitempty" xml:"StartTriggerTime,omitempty"`
	// The status used for matching manual workflow instances.
	//
	// Valid values:
	//
	// 	- Success
	//
	// 	- Failure
	//
	// example:
	//
	// Failure
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the workflow instance. Valid values:
	//
	// ManualWorkflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// ManualWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The workflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s RerunWorkflowInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RerunWorkflowInstancesShrinkRequest) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *RerunWorkflowInstancesShrinkRequest) GetEndTriggerTime() *int64 {
	return s.EndTriggerTime
}

func (s *RerunWorkflowInstancesShrinkRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *RerunWorkflowInstancesShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *RerunWorkflowInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *RerunWorkflowInstancesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *RerunWorkflowInstancesShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RerunWorkflowInstancesShrinkRequest) GetStartTriggerTime() *int64 {
	return s.StartTriggerTime
}

func (s *RerunWorkflowInstancesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *RerunWorkflowInstancesShrinkRequest) GetType() *string {
	return s.Type
}

func (s *RerunWorkflowInstancesShrinkRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *RerunWorkflowInstancesShrinkRequest) SetBizdate(v int64) *RerunWorkflowInstancesShrinkRequest {
	s.Bizdate = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetEndTriggerTime(v int64) *RerunWorkflowInstancesShrinkRequest {
	s.EndTriggerTime = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetEnvType(v string) *RerunWorkflowInstancesShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetFilterShrink(v string) *RerunWorkflowInstancesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetIdsShrink(v string) *RerunWorkflowInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetName(v string) *RerunWorkflowInstancesShrinkRequest {
	s.Name = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetProjectId(v int64) *RerunWorkflowInstancesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetStartTriggerTime(v int64) *RerunWorkflowInstancesShrinkRequest {
	s.StartTriggerTime = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetStatus(v string) *RerunWorkflowInstancesShrinkRequest {
	s.Status = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetType(v string) *RerunWorkflowInstancesShrinkRequest {
	s.Type = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) SetWorkflowId(v int64) *RerunWorkflowInstancesShrinkRequest {
	s.WorkflowId = &v
	return s
}

func (s *RerunWorkflowInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
