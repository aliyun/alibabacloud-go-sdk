// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationMaintainConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeActiveOperationMaintainConfRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationMaintainConfRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeActiveOperationMaintainConfRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeActiveOperationMaintainConfRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationMaintainConfRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationMaintainConfRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeActiveOperationMaintainConfRequest
	GetSecurityToken() *string
}

type DescribeActiveOperationMaintainConfRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-re*********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeActiveOperationMaintainConfRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationMaintainConfRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationMaintainConfRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationMaintainConfRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationMaintainConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActiveOperationMaintainConfRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActiveOperationMaintainConfRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationMaintainConfRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationMaintainConfRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeActiveOperationMaintainConfRequest) SetOwnerAccount(v string) *DescribeActiveOperationMaintainConfRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) SetOwnerId(v int64) *DescribeActiveOperationMaintainConfRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) SetRegionId(v string) *DescribeActiveOperationMaintainConfRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) SetResourceGroupId(v string) *DescribeActiveOperationMaintainConfRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationMaintainConfRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationMaintainConfRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) SetSecurityToken(v string) *DescribeActiveOperationMaintainConfRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationMaintainConfRequest) Validate() error {
	return dara.Validate(s)
}
