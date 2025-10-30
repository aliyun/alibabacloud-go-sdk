// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindBatchAxgShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAxgBindListShrink(v string) *BindBatchAxgShrinkRequest
	GetAxgBindListShrink() *string
	SetOwnerId(v int64) *BindBatchAxgShrinkRequest
	GetOwnerId() *int64
	SetPoolKey(v string) *BindBatchAxgShrinkRequest
	GetPoolKey() *string
	SetResourceOwnerAccount(v string) *BindBatchAxgShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindBatchAxgShrinkRequest
	GetResourceOwnerId() *int64
}

type BindBatchAxgShrinkRequest struct {
	// This parameter is required.
	AxgBindListShrink *string `json:"AxgBindList,omitempty" xml:"AxgBindList,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// FC2235****
	PoolKey              *string `json:"PoolKey,omitempty" xml:"PoolKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindBatchAxgShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindBatchAxgShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindBatchAxgShrinkRequest) GetAxgBindListShrink() *string {
	return s.AxgBindListShrink
}

func (s *BindBatchAxgShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindBatchAxgShrinkRequest) GetPoolKey() *string {
	return s.PoolKey
}

func (s *BindBatchAxgShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindBatchAxgShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindBatchAxgShrinkRequest) SetAxgBindListShrink(v string) *BindBatchAxgShrinkRequest {
	s.AxgBindListShrink = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetOwnerId(v int64) *BindBatchAxgShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetPoolKey(v string) *BindBatchAxgShrinkRequest {
	s.PoolKey = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetResourceOwnerAccount(v string) *BindBatchAxgShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) SetResourceOwnerId(v int64) *BindBatchAxgShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindBatchAxgShrinkRequest) Validate() error {
	return dara.Validate(s)
}
