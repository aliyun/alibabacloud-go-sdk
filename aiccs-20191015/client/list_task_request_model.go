// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ListTaskRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListTaskRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListTaskRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTaskRequest
	GetResourceOwnerId() *int64
	SetRobotName(v string) *ListTaskRequest
	GetRobotName() *string
	SetStatus(v string) *ListTaskRequest
	GetStatus() *string
	SetTaskId(v int64) *ListTaskRequest
	GetTaskId() *int64
	SetTaskName(v string) *ListTaskRequest
	GetTaskName() *string
}

type ListTaskRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RobotName            *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	// example:
	//
	// STOP
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId   *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskRequest) GoString() string {
	return s.String()
}

func (s *ListTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTaskRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTaskRequest) GetRobotName() *string {
	return s.RobotName
}

func (s *ListTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTaskRequest) SetOwnerId(v int64) *ListTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTaskRequest) SetPageNo(v int32) *ListTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListTaskRequest) SetPageSize(v int32) *ListTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskRequest) SetResourceOwnerAccount(v string) *ListTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTaskRequest) SetResourceOwnerId(v int64) *ListTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTaskRequest) SetRobotName(v string) *ListTaskRequest {
	s.RobotName = &v
	return s
}

func (s *ListTaskRequest) SetStatus(v string) *ListTaskRequest {
	s.Status = &v
	return s
}

func (s *ListTaskRequest) SetTaskId(v int64) *ListTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskRequest) SetTaskName(v string) *ListTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListTaskRequest) Validate() error {
	return dara.Validate(s)
}
