// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *ListTasksShrinkRequest
	GetIdsShrink() *string
	SetName(v string) *ListTasksShrinkRequest
	GetName() *string
	SetOwner(v string) *ListTasksShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListTasksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksShrinkRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListTasksShrinkRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListTasksShrinkRequest
	GetProjectId() *int64
	SetRuntimeResource(v string) *ListTasksShrinkRequest
	GetRuntimeResource() *string
	SetSortBy(v string) *ListTasksShrinkRequest
	GetSortBy() *string
	SetTaskType(v string) *ListTasksShrinkRequest
	GetTaskType() *string
	SetTriggerRecurrence(v string) *ListTasksShrinkRequest
	GetTriggerRecurrence() *string
	SetTriggerType(v string) *ListTasksShrinkRequest
	GetTriggerType() *string
	SetWorkflowId(v int64) *ListTasksShrinkRequest
	GetWorkflowId() *int64
}

type ListTasksShrinkRequest struct {
	// The ID of the task.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The name of the task. Fuzzy match is supported.
	//
	// example:
	//
	// SQL node
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// 1000
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace environment.
	//
	// Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The information about the resource group. Set this parameter to the ID of a resource group for scheduling.
	//
	// example:
	//
	// S_res_group_524258031846018_1684XXXXXXXXX
	RuntimeResource *string `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty"`
	// The field that is used to sort tasks. This parameter is configured in the format of "Sorting field Sorting order". You can set the sorting order to Desc or Asc. If you do not specify the sorting order, Asc is used by default. Valid values:
	//
	// 	- `ModifyTime (Desc/Asc)`
	//
	// 	- `CreateTime (Desc/Asc)`
	//
	// 	- `Id (Desc/Asc)`
	//
	//     Default value: `Id Desc`.
	//
	// example:
	//
	// Id Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- ODPS_SQL
	//
	// 	- SPARK
	//
	// 	- PY_ODPS
	//
	// 	- PY_ODPS3
	//
	// 	- ODPS_SCRIPT
	//
	// 	- ODPS_MR
	//
	// 	- COMPONENT_SQL
	//
	// 	- EMR_HIVE
	//
	// 	- EMR_MR
	//
	// 	- EMR_SPARK_SQL
	//
	// 	- EMR_SPARK
	//
	// 	- EMR_SHELL
	//
	// 	- EMR_PRESTO
	//
	// 	- EMR_IMPALA
	//
	// 	- SPARK_STREAMING
	//
	// 	- EMR_KYUUBI
	//
	// 	- EMR_TRINO
	//
	// 	- HOLOGRES_SQL
	//
	// 	- HOLOGRES_SYNC_DDL
	//
	// 	- HOLOGRES_SYNC_DATA
	//
	// example:
	//
	// ODPS_SQL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The run mode when triggered. Valid only if TriggerType is Scheduler.
	//
	// Valid values:
	//
	// 	- Pause
	//
	// 	- Skip
	//
	// 	- Normal
	//
	// example:
	//
	// Normal
	TriggerRecurrence *string `json:"TriggerRecurrence,omitempty" xml:"TriggerRecurrence,omitempty"`
	// The trigger type.
	//
	// Valid values:
	//
	// 	- Scheduler: Triggered by schedule.
	//
	// 	- Manual: Triggered manually.
	//
	// example:
	//
	// Scheduler
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The ID of the workflow to which the task belongs.
	//
	// example:
	//
	// 1234
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListTasksShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListTasksShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListTasksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksShrinkRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTasksShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTasksShrinkRequest) GetRuntimeResource() *string {
	return s.RuntimeResource
}

func (s *ListTasksShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTasksShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksShrinkRequest) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListTasksShrinkRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListTasksShrinkRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListTasksShrinkRequest) SetIdsShrink(v string) *ListTasksShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetName(v string) *ListTasksShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListTasksShrinkRequest) SetOwner(v string) *ListTasksShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageNumber(v int32) *ListTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageSize(v int32) *ListTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectEnv(v string) *ListTasksShrinkRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectId(v int64) *ListTasksShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetRuntimeResource(v string) *ListTasksShrinkRequest {
	s.RuntimeResource = &v
	return s
}

func (s *ListTasksShrinkRequest) SetSortBy(v string) *ListTasksShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskType(v string) *ListTasksShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTriggerRecurrence(v string) *ListTasksShrinkRequest {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTriggerType(v string) *ListTasksShrinkRequest {
	s.TriggerType = &v
	return s
}

func (s *ListTasksShrinkRequest) SetWorkflowId(v int64) *ListTasksShrinkRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
