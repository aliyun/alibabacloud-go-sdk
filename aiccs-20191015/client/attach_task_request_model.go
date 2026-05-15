// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallString(v string) *AttachTaskRequest
	GetCallString() *string
	SetOwnerId(v int64) *AttachTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AttachTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachTaskRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *AttachTaskRequest
	GetTaskId() *int64
}

type AttachTaskRequest struct {
	CallString           *string `json:"CallString,omitempty" xml:"CallString,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AttachTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachTaskRequest) GoString() string {
	return s.String()
}

func (s *AttachTaskRequest) GetCallString() *string {
	return s.CallString
}

func (s *AttachTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AttachTaskRequest) SetCallString(v string) *AttachTaskRequest {
	s.CallString = &v
	return s
}

func (s *AttachTaskRequest) SetOwnerId(v int64) *AttachTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachTaskRequest) SetResourceOwnerAccount(v string) *AttachTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachTaskRequest) SetResourceOwnerId(v int64) *AttachTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachTaskRequest) SetTaskId(v int64) *AttachTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AttachTaskRequest) Validate() error {
	return dara.Validate(s)
}
