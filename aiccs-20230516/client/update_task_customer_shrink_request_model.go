// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskCustomerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomersShrink(v string) *UpdateTaskCustomerShrinkRequest
	GetCustomersShrink() *string
	SetOwnerId(v int64) *UpdateTaskCustomerShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateTaskCustomerShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTaskCustomerShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *UpdateTaskCustomerShrinkRequest
	GetTaskId() *int64
}

type UpdateTaskCustomerShrinkRequest struct {
	// 外呼客户
	//
	// This parameter is required.
	CustomersShrink      *string `json:"Customers,omitempty" xml:"Customers,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 59
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTaskCustomerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskCustomerShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskCustomerShrinkRequest) GetCustomersShrink() *string {
	return s.CustomersShrink
}

func (s *UpdateTaskCustomerShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTaskCustomerShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTaskCustomerShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTaskCustomerShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateTaskCustomerShrinkRequest) SetCustomersShrink(v string) *UpdateTaskCustomerShrinkRequest {
	s.CustomersShrink = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetOwnerId(v int64) *UpdateTaskCustomerShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetResourceOwnerAccount(v string) *UpdateTaskCustomerShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetResourceOwnerId(v int64) *UpdateTaskCustomerShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) SetTaskId(v int64) *UpdateTaskCustomerShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTaskCustomerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
