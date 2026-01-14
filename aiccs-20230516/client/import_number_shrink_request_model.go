// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomersShrink(v string) *ImportNumberShrinkRequest
	GetCustomersShrink() *string
	SetFailReturn(v int64) *ImportNumberShrinkRequest
	GetFailReturn() *int64
	SetOutId(v string) *ImportNumberShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *ImportNumberShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ImportNumberShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportNumberShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ImportNumberShrinkRequest
	GetTaskId() *int64
}

type ImportNumberShrinkRequest struct {
	// This parameter is required.
	CustomersShrink *string `json:"Customers,omitempty" xml:"Customers,omitempty"`
	// example:
	//
	// 0
	FailReturn *int64 `json:"FailReturn,omitempty" xml:"FailReturn,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportNumberShrinkRequest) GetCustomersShrink() *string {
	return s.CustomersShrink
}

func (s *ImportNumberShrinkRequest) GetFailReturn() *int64 {
	return s.FailReturn
}

func (s *ImportNumberShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *ImportNumberShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportNumberShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportNumberShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportNumberShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ImportNumberShrinkRequest) SetCustomersShrink(v string) *ImportNumberShrinkRequest {
	s.CustomersShrink = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetFailReturn(v int64) *ImportNumberShrinkRequest {
	s.FailReturn = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetOutId(v string) *ImportNumberShrinkRequest {
	s.OutId = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetOwnerId(v int64) *ImportNumberShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetResourceOwnerAccount(v string) *ImportNumberShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetResourceOwnerId(v int64) *ImportNumberShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportNumberShrinkRequest) SetTaskId(v int64) *ImportNumberShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ImportNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
