// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalSecurityIPGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeGlobalSecurityIPGroupRelationRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRelationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeGlobalSecurityIPGroupRelationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeGlobalSecurityIPGroupRelationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGlobalSecurityIPGroupRelationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRelationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeGlobalSecurityIPGroupRelationRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeGlobalSecurityIPGroupRelationRequest
	GetSecurityToken() *string
}

type DescribeGlobalSecurityIPGroupRelationRequest struct {
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// rg-**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeGlobalSecurityIPGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalSecurityIPGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetDBClusterId(v string) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetOwnerId(v int64) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetRegionId(v string) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetResourceGroupId(v string) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetResourceOwnerAccount(v string) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetResourceOwnerId(v int64) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) SetSecurityToken(v string) *DescribeGlobalSecurityIPGroupRelationRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalSecurityIPGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
