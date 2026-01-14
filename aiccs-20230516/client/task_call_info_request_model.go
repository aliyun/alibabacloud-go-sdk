// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskCallInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *TaskCallInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskCallInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskCallInfoRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *TaskCallInfoRequest
	GetTaskId() *int64
}

type TaskCallInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskCallInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskCallInfoRequest) GoString() string {
	return s.String()
}

func (s *TaskCallInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskCallInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskCallInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskCallInfoRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskCallInfoRequest) SetOwnerId(v int64) *TaskCallInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskCallInfoRequest) SetResourceOwnerAccount(v string) *TaskCallInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskCallInfoRequest) SetResourceOwnerId(v int64) *TaskCallInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskCallInfoRequest) SetTaskId(v int64) *TaskCallInfoRequest {
	s.TaskId = &v
	return s
}

func (s *TaskCallInfoRequest) Validate() error {
	return dara.Validate(s)
}
