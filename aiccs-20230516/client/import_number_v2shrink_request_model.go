// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNumberV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomersShrink(v string) *ImportNumberV2ShrinkRequest
	GetCustomersShrink() *string
	SetFailReturn(v int64) *ImportNumberV2ShrinkRequest
	GetFailReturn() *int64
	SetOutId(v string) *ImportNumberV2ShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *ImportNumberV2ShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ImportNumberV2ShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportNumberV2ShrinkRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *ImportNumberV2ShrinkRequest
	GetTaskId() *int64
}

type ImportNumberV2ShrinkRequest struct {
	CustomersShrink *string `json:"Customers,omitempty" xml:"Customers,omitempty"`
	// example:
	//
	// 1
	FailReturn *int64 `json:"FailReturn,omitempty" xml:"FailReturn,omitempty"`
	// example:
	//
	// 示例值
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 92
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportNumberV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportNumberV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportNumberV2ShrinkRequest) GetCustomersShrink() *string {
	return s.CustomersShrink
}

func (s *ImportNumberV2ShrinkRequest) GetFailReturn() *int64 {
	return s.FailReturn
}

func (s *ImportNumberV2ShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *ImportNumberV2ShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportNumberV2ShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportNumberV2ShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportNumberV2ShrinkRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ImportNumberV2ShrinkRequest) SetCustomersShrink(v string) *ImportNumberV2ShrinkRequest {
	s.CustomersShrink = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) SetFailReturn(v int64) *ImportNumberV2ShrinkRequest {
	s.FailReturn = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) SetOutId(v string) *ImportNumberV2ShrinkRequest {
	s.OutId = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) SetOwnerId(v int64) *ImportNumberV2ShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) SetResourceOwnerAccount(v string) *ImportNumberV2ShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) SetResourceOwnerId(v int64) *ImportNumberV2ShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) SetTaskId(v int64) *ImportNumberV2ShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ImportNumberV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
