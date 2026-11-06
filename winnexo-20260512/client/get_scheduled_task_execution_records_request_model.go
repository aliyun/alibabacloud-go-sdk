// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskExecutionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetScheduledTaskExecutionRecordsRequest
	GetCollaborationGroupId() *string
	SetInitiatorUserId(v string) *GetScheduledTaskExecutionRecordsRequest
	GetInitiatorUserId() *string
	SetPage(v int32) *GetScheduledTaskExecutionRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *GetScheduledTaskExecutionRecordsRequest
	GetPageSize() *int32
	SetStatus(v string) *GetScheduledTaskExecutionRecordsRequest
	GetStatus() *string
	SetTaskId(v string) *GetScheduledTaskExecutionRecordsRequest
	GetTaskId() *string
	SetTenantId(v string) *GetScheduledTaskExecutionRecordsRequest
	GetTenantId() *string
}

type GetScheduledTaskExecutionRecordsRequest struct {
	// The ID of the collaboration group to which the task belongs, such as cg_101. If you specify this parameter, a group space task is created, and the caller must be a valid group member. If you leave this parameter empty, a personal task is created.
	//
	// example:
	//
	// 1112
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The business status code. A value of 200 indicates success. A failure is indicated by a backend error code, such as ERR.	- or InvalidParameter.*.
	//
	// example:
	//
	// 5
	InitiatorUserId *string `json:"initiatorUserId,omitempty" xml:"initiatorUserId,omitempty"`
	// The page number. Default value: 1. Minimum value: 1. Maximum value: 200.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The initiator filter based on the platform user ID. The scope is the executor of the record. For manual execution, it is the user who triggers the task. For automatic execution, it is the task creator. To view only the tasks that you initiated, specify the current user ID.
	//
	// example:
	//
	// failed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The execution status filter in lowercase. Valid values: pending (queued), running (running), success (successful), failed (failed), timeout (timed out), and cancelled (cancelled). If you do not specify this parameter, no filtering is applied. If you specify this parameter, no future plan items are generated.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The ID of the active tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduledTaskExecutionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskExecutionRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetInitiatorUserId() *string {
	return s.InitiatorUserId
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScheduledTaskExecutionRecordsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetCollaborationGroupId(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetInitiatorUserId(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.InitiatorUserId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetPage(v int32) *GetScheduledTaskExecutionRecordsRequest {
	s.Page = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetPageSize(v int32) *GetScheduledTaskExecutionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetStatus(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.Status = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetTaskId(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.TaskId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) SetTenantId(v string) *GetScheduledTaskExecutionRecordsRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskExecutionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
