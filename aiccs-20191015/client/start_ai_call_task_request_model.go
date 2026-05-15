// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAiCallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StartAiCallTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StartAiCallTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartAiCallTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *StartAiCallTaskRequest
	GetTaskId() *string
}

type StartAiCallTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 121212121****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartAiCallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAiCallTaskRequest) GoString() string {
	return s.String()
}

func (s *StartAiCallTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartAiCallTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartAiCallTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartAiCallTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartAiCallTaskRequest) SetOwnerId(v int64) *StartAiCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartAiCallTaskRequest) SetResourceOwnerAccount(v string) *StartAiCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartAiCallTaskRequest) SetResourceOwnerId(v int64) *StartAiCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartAiCallTaskRequest) SetTaskId(v string) *StartAiCallTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartAiCallTaskRequest) Validate() error {
	return dara.Validate(s)
}
