// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalled(v string) *ListDialogRequest
	GetCalled() *string
	SetOwnerId(v int64) *ListDialogRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListDialogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDialogRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ListDialogRequest
	GetTaskId() *int64
}

type ListDialogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 130****0000
	Called               *string `json:"Called,omitempty" xml:"Called,omitempty"`
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

func (s ListDialogRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDialogRequest) GoString() string {
	return s.String()
}

func (s *ListDialogRequest) GetCalled() *string {
	return s.Called
}

func (s *ListDialogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDialogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDialogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDialogRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListDialogRequest) SetCalled(v string) *ListDialogRequest {
	s.Called = &v
	return s
}

func (s *ListDialogRequest) SetOwnerId(v int64) *ListDialogRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDialogRequest) SetResourceOwnerAccount(v string) *ListDialogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDialogRequest) SetResourceOwnerId(v int64) *ListDialogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDialogRequest) SetTaskId(v int64) *ListDialogRequest {
	s.TaskId = &v
	return s
}

func (s *ListDialogRequest) Validate() error {
	return dara.Validate(s)
}
