// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StartTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StartTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartTaskRequest
	GetResourceOwnerId() *int64
	SetStartNow(v bool) *StartTaskRequest
	GetStartNow() *bool
	SetTaskId(v int64) *StartTaskRequest
	GetTaskId() *int64
}

type StartTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Indicates whether to start immediately. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No, which means the job will not be started.
	//
	// example:
	//
	// true
	StartNow *bool `json:"StartNow,omitempty" xml:"StartNow,omitempty"`
	// The unique job ID of the robot calling task. You can view it on the [Task Management](https://aiccs.console.aliyun.com/job/list) page or obtain it by using the [CreateTask](https://help.aliyun.com/document_detail/223556.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartTaskRequest) GetStartNow() *bool {
	return s.StartNow
}

func (s *StartTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *StartTaskRequest) SetOwnerId(v int64) *StartTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTaskRequest) SetResourceOwnerAccount(v string) *StartTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTaskRequest) SetResourceOwnerId(v int64) *StartTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTaskRequest) SetStartNow(v bool) *StartTaskRequest {
	s.StartNow = &v
	return s
}

func (s *StartTaskRequest) SetTaskId(v int64) *StartTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartTaskRequest) Validate() error {
	return dara.Validate(s)
}
