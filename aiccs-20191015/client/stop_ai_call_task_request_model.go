// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAiCallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *StopAiCallTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopAiCallTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopAiCallTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *StopAiCallTaskRequest
	GetTaskId() *string
}

type StopAiCallTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1187**************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopAiCallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAiCallTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAiCallTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopAiCallTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopAiCallTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopAiCallTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopAiCallTaskRequest) SetOwnerId(v int64) *StopAiCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopAiCallTaskRequest) SetResourceOwnerAccount(v string) *StopAiCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopAiCallTaskRequest) SetResourceOwnerId(v int64) *StopAiCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopAiCallTaskRequest) SetTaskId(v string) *StopAiCallTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopAiCallTaskRequest) Validate() error {
	return dara.Validate(s)
}
