// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CancelTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CancelTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CancelTaskRequest
	GetTaskId() *int64
}

type CancelTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique job ID of the robot calling task. You can view it in the [Task Management](https://aiccs.console.aliyun.com/job/list) interface or obtain it by using the [CreateTask](https://help.aliyun.com/document_detail/223556.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CancelTaskRequest) SetOwnerId(v int64) *CancelTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerAccount(v string) *CancelTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerId(v int64) *CancelTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetTaskId(v int64) *CancelTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelTaskRequest) Validate() error {
	return dara.Validate(s)
}
