// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *TaskListRequest
	GetCreateTime() *string
	SetLastCallTime(v string) *TaskListRequest
	GetLastCallTime() *string
	SetOwnerId(v int64) *TaskListRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskListRequest
	GetResourceOwnerId() *int64
	SetStatus(v int64) *TaskListRequest
	GetStatus() *int64
	SetTaskId(v int64) *TaskListRequest
	GetTaskId() *int64
}

type TaskListRequest struct {
	// 创建时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 最后外呼时间
	//
	// example:
	//
	// 2023-04-05 12:11:11
	LastCallTime         *string `json:"LastCallTime,omitempty" xml:"LastCallTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务状态。1 未启用，2 启用中，4 已停止
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 2
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskListRequest) GoString() string {
	return s.String()
}

func (s *TaskListRequest) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TaskListRequest) GetLastCallTime() *string {
	return s.LastCallTime
}

func (s *TaskListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskListRequest) GetStatus() *int64 {
	return s.Status
}

func (s *TaskListRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskListRequest) SetCreateTime(v string) *TaskListRequest {
	s.CreateTime = &v
	return s
}

func (s *TaskListRequest) SetLastCallTime(v string) *TaskListRequest {
	s.LastCallTime = &v
	return s
}

func (s *TaskListRequest) SetOwnerId(v int64) *TaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskListRequest) SetResourceOwnerAccount(v string) *TaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskListRequest) SetResourceOwnerId(v int64) *TaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskListRequest) SetStatus(v int64) *TaskListRequest {
	s.Status = &v
	return s
}

func (s *TaskListRequest) SetTaskId(v int64) *TaskListRequest {
	s.TaskId = &v
	return s
}

func (s *TaskListRequest) Validate() error {
	return dara.Validate(s)
}
