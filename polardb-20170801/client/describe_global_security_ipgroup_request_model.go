// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityGroupId(v string) *DescribeGlobalSecurityIPGroupRequest
	GetGlobalSecurityGroupId() *string
	SetOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGlobalSecurityIPGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeGlobalSecurityIPGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGlobalSecurityIPGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGlobalSecurityIPGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeGlobalSecurityIPGroupRequest
	GetSecurityToken() *string
}

type DescribeGlobalSecurityIPGroupRequest struct {
	// The ID of the IP whitelist template.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IP whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGlobalSecurityIPGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetGlobalSecurityGroupId(v string) *DescribeGlobalSecurityIPGroupRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetOwnerId(v int64) *DescribeGlobalSecurityIPGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetRegionId(v string) *DescribeGlobalSecurityIPGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetResourceGroupId(v string) *DescribeGlobalSecurityIPGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetResourceOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetResourceOwnerId(v int64) *DescribeGlobalSecurityIPGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) SetSecurityToken(v string) *DescribeGlobalSecurityIPGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
