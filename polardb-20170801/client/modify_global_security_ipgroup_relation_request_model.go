// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetDBClusterId() *string
	SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetGlobalSecurityGroupId() *string
	SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyGlobalSecurityIPGroupRelationRequest
	GetSecurityToken() *string
}

type ModifyGlobalSecurityIPGroupRelationRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the IP whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// g-zsldxfiwjmti0kcm****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
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

func (s ModifyGlobalSecurityIPGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetDBClusterId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetRegionId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetResourceGroupId(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) SetSecurityToken(v string) *ModifyGlobalSecurityIPGroupRelationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
