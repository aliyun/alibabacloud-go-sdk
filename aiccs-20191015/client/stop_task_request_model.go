// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StopTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *StopTaskRequest
	GetTaskId() *int64
}

type StopTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique job ID of the robot calling task. You can view it on the [Task Management](https://aiccs.console.aliyun.com/job/list) page or obtain it by using the [CreateTask](https://help.aliyun.com/document_detail/223556.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StopTaskRequest) SetOwnerId(v int64) *StopTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopTaskRequest) SetResourceOwnerAccount(v string) *StopTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopTaskRequest) SetResourceOwnerId(v int64) *StopTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopTaskRequest) SetTaskId(v int64) *StopTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopTaskRequest) Validate() error {
	return dara.Validate(s)
}
