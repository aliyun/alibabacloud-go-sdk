// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetGlobalIgName() *string
	SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetGlobalSecurityGroupId() *string
	SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupNameRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyGlobalSecurityIPGroupNameRequest
	GetSecurityToken() *string
}

type ModifyGlobalSecurityIPGroupNameRequest struct {
	// The name of the IP whitelist template. The name must meet the following requirements:
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must start with a letter and end with a letter or digit.
	//
	// 	- The name must be 2 to 120 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
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
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.GlobalIgName = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupNameRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetRegionId(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetResourceGroupId(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) SetSecurityToken(v string) *ModifyGlobalSecurityIPGroupNameRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupNameRequest) Validate() error {
	return dara.Validate(s)
}
