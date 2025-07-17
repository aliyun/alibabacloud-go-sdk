// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *ListTasksRequest
	GetIds() []*int64
	SetName(v string) *ListTasksRequest
	GetName() *string
	SetOwner(v string) *ListTasksRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListTasksRequest
	GetProjectEnv() *string
	SetProjectId(v int64) *ListTasksRequest
	GetProjectId() *int64
	SetRuntimeResource(v string) *ListTasksRequest
	GetRuntimeResource() *string
	SetSortBy(v string) *ListTasksRequest
	GetSortBy() *string
	SetTaskType(v string) *ListTasksRequest
	GetTaskType() *string
	SetTriggerRecurrence(v string) *ListTasksRequest
	GetTriggerRecurrence() *string
	SetTriggerType(v string) *ListTasksRequest
	GetTriggerType() *string
	SetWorkflowId(v int64) *ListTasksRequest
	GetWorkflowId() *int64
}

type ListTasksRequest struct {
	// The ID of the task.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
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
	// The environment of the workspace.
	//
	// Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
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
	// The running mode of the task after it is triggered. This parameter takes effect only if the TriggerType parameter is set to Scheduler.
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
	// 	- Scheduler: scheduling cycle-based trigger
	//
	// 	- Manual: manual trigger
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

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListTasksRequest) GetName() *string {
	return s.Name
}

func (s *ListTasksRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListTasksRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTasksRequest) GetRuntimeResource() *string {
	return s.RuntimeResource
}

func (s *ListTasksRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksRequest) GetTriggerRecurrence() *string {
	return s.TriggerRecurrence
}

func (s *ListTasksRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListTasksRequest) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListTasksRequest) SetIds(v []*int64) *ListTasksRequest {
	s.Ids = v
	return s
}

func (s *ListTasksRequest) SetName(v string) *ListTasksRequest {
	s.Name = &v
	return s
}

func (s *ListTasksRequest) SetOwner(v string) *ListTasksRequest {
	s.Owner = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProjectEnv(v string) *ListTasksRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListTasksRequest) SetProjectId(v int64) *ListTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksRequest) SetRuntimeResource(v string) *ListTasksRequest {
	s.RuntimeResource = &v
	return s
}

func (s *ListTasksRequest) SetSortBy(v string) *ListTasksRequest {
	s.SortBy = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListTasksRequest) SetTriggerRecurrence(v string) *ListTasksRequest {
	s.TriggerRecurrence = &v
	return s
}

func (s *ListTasksRequest) SetTriggerType(v string) *ListTasksRequest {
	s.TriggerType = &v
	return s
}

func (s *ListTasksRequest) SetWorkflowId(v int64) *ListTasksRequest {
	s.WorkflowId = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
