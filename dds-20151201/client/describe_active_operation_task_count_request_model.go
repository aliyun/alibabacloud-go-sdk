// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeActiveOperationTaskCountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTaskCountRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeActiveOperationTaskCountRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskCountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTaskCountRequest
	GetResourceOwnerId() *int64
}

type DescribeActiveOperationTaskCountRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// sg-bp5689ac****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeActiveOperationTaskCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskCountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationTaskCountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationTaskCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActiveOperationTaskCountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTaskCountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationTaskCountRequest) SetOwnerAccount(v string) *DescribeActiveOperationTaskCountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetOwnerId(v int64) *DescribeActiveOperationTaskCountRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetResourceGroupId(v string) *DescribeActiveOperationTaskCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTaskCountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTaskCountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTaskCountRequest) Validate() error {
	return dara.Validate(s)
}
