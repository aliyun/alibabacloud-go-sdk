// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCancelCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNumbers(v []*string) *TaskCancelCallRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *TaskCancelCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskCancelCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskCancelCallRequest
	GetResourceOwnerId() *int64
	SetTags(v []*string) *TaskCancelCallRequest
	GetTags() []*string
	SetTaskId(v int64) *TaskCancelCallRequest
	GetTaskId() *int64
}

type TaskCancelCallRequest struct {
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tags                 []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCancelCallRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskCancelCallRequest) GoString() string {
	return s.String()
}

func (s *TaskCancelCallRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *TaskCancelCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskCancelCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskCancelCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskCancelCallRequest) GetTags() []*string {
	return s.Tags
}

func (s *TaskCancelCallRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCancelCallRequest) SetNumbers(v []*string) *TaskCancelCallRequest {
	s.Numbers = v
	return s
}

func (s *TaskCancelCallRequest) SetOwnerId(v int64) *TaskCancelCallRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCancelCallRequest) SetResourceOwnerAccount(v string) *TaskCancelCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCancelCallRequest) SetResourceOwnerId(v int64) *TaskCancelCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCancelCallRequest) SetTags(v []*string) *TaskCancelCallRequest {
	s.Tags = v
	return s
}

func (s *TaskCancelCallRequest) SetTaskId(v int64) *TaskCancelCallRequest {
	s.TaskId = &v
	return s
}

func (s *TaskCancelCallRequest) Validate() error {
	return dara.Validate(s)
}
