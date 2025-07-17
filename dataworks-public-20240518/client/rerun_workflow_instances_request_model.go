// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v int64) *RerunWorkflowInstancesRequest
	GetBizdate() *int64
	SetEndTriggerTime(v int64) *RerunWorkflowInstancesRequest
	GetEndTriggerTime() *int64
	SetEnvType(v string) *RerunWorkflowInstancesRequest
	GetEnvType() *string
	SetFilter(v *RerunWorkflowInstancesRequestFilter) *RerunWorkflowInstancesRequest
	GetFilter() *RerunWorkflowInstancesRequestFilter
	SetIds(v []*int64) *RerunWorkflowInstancesRequest
	GetIds() []*int64
	SetName(v string) *RerunWorkflowInstancesRequest
	GetName() *string
	SetProjectId(v int64) *RerunWorkflowInstancesRequest
	GetProjectId() *int64
	SetStartTriggerTime(v int64) *RerunWorkflowInstancesRequest
	GetStartTriggerTime() *int64
	SetStatus(v string) *RerunWorkflowInstancesRequest
	GetStatus() *string
	SetType(v string) *RerunWorkflowInstancesRequest
	GetType() *string
	SetWorkflowId(v int64) *RerunWorkflowInstancesRequest
	GetWorkflowId() *int64
}

type RerunWorkflowInstancesRequest struct {
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
	Filter *RerunWorkflowInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The instance IDs used for matching manual workflow instances.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
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

func (s RerunWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *RerunWorkflowInstancesRequest) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *RerunWorkflowInstancesRequest) GetEndTriggerTime() *int64 {
	return s.EndTriggerTime
}

func (s *RerunWorkflowInstancesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *RerunWorkflowInstancesRequest) GetFilter() *RerunWorkflowInstancesRequestFilter {
	return s.Filter
}

func (s *RerunWorkflowInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *RerunWorkflowInstancesRequest) GetName() *string {
	return s.Name
}

func (s *RerunWorkflowInstancesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RerunWorkflowInstancesRequest) GetStartTriggerTime() *int64 {
	return s.StartTriggerTime
}

func (s *RerunWorkflowInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *RerunWorkflowInstancesRequest) GetType() *string {
	return s.Type
}

func (s *RerunWorkflowInstancesRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *RerunWorkflowInstancesRequest) SetBizdate(v int64) *RerunWorkflowInstancesRequest {
	s.Bizdate = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetEndTriggerTime(v int64) *RerunWorkflowInstancesRequest {
	s.EndTriggerTime = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetEnvType(v string) *RerunWorkflowInstancesRequest {
	s.EnvType = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetFilter(v *RerunWorkflowInstancesRequestFilter) *RerunWorkflowInstancesRequest {
	s.Filter = v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetIds(v []*int64) *RerunWorkflowInstancesRequest {
	s.Ids = v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetName(v string) *RerunWorkflowInstancesRequest {
	s.Name = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetProjectId(v int64) *RerunWorkflowInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetStartTriggerTime(v int64) *RerunWorkflowInstancesRequest {
	s.StartTriggerTime = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetStatus(v string) *RerunWorkflowInstancesRequest {
	s.Status = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetType(v string) *RerunWorkflowInstancesRequest {
	s.Type = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) SetWorkflowId(v int64) *RerunWorkflowInstancesRequest {
	s.WorkflowId = &v
	return s
}

func (s *RerunWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type RerunWorkflowInstancesRequestFilter struct {
	// Specifies whether to rerun the matched instances and all downstream instances.
	//
	// example:
	//
	// false
	RerunDownstreamEnabled *bool `json:"RerunDownstreamEnabled,omitempty" xml:"RerunDownstreamEnabled,omitempty"`
	// The internal task IDs used for matching manual workflow instances.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The statuses of internal tasks used for matching manual workflow instances.
	TaskInstanceStatuses []*string `json:"TaskInstanceStatuses,omitempty" xml:"TaskInstanceStatuses,omitempty" type:"Repeated"`
	// The internal task name used for matching the manual workflow instance.
	//
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Match internal tasks within the manual workflow by type.
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s RerunWorkflowInstancesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s RerunWorkflowInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *RerunWorkflowInstancesRequestFilter) GetRerunDownstreamEnabled() *bool {
	return s.RerunDownstreamEnabled
}

func (s *RerunWorkflowInstancesRequestFilter) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *RerunWorkflowInstancesRequestFilter) GetTaskInstanceStatuses() []*string {
	return s.TaskInstanceStatuses
}

func (s *RerunWorkflowInstancesRequestFilter) GetTaskName() *string {
	return s.TaskName
}

func (s *RerunWorkflowInstancesRequestFilter) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *RerunWorkflowInstancesRequestFilter) SetRerunDownstreamEnabled(v bool) *RerunWorkflowInstancesRequestFilter {
	s.RerunDownstreamEnabled = &v
	return s
}

func (s *RerunWorkflowInstancesRequestFilter) SetTaskIds(v []*int64) *RerunWorkflowInstancesRequestFilter {
	s.TaskIds = v
	return s
}

func (s *RerunWorkflowInstancesRequestFilter) SetTaskInstanceStatuses(v []*string) *RerunWorkflowInstancesRequestFilter {
	s.TaskInstanceStatuses = v
	return s
}

func (s *RerunWorkflowInstancesRequestFilter) SetTaskName(v string) *RerunWorkflowInstancesRequestFilter {
	s.TaskName = &v
	return s
}

func (s *RerunWorkflowInstancesRequestFilter) SetTaskTypes(v []*string) *RerunWorkflowInstancesRequestFilter {
	s.TaskTypes = v
	return s
}

func (s *RerunWorkflowInstancesRequestFilter) Validate() error {
	return dara.Validate(s)
}
