// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGIpList(v string) *ModifyGlobalSecurityIPGroupRequest
	GetGIpList() *string
	SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupRequest
	GetGlobalIgName() *string
	SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRequest
	GetGlobalSecurityGroupId() *string
	SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyGlobalSecurityIPGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyGlobalSecurityIPGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyGlobalSecurityIPGroupRequest
	GetSecurityToken() *string
}

type ModifyGlobalSecurityIPGroupRequest struct {
	// The IP address in the whitelist template.
	//
	// >  Multiple IP addresses are separated by commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP whitelists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
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
	// rg-**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGlobalSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetGIpList(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.GIpList = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetGlobalIgName(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.GlobalIgName = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetGlobalSecurityGroupId(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetOwnerId(v int64) *ModifyGlobalSecurityIPGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetRegionId(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetResourceGroupId(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetSecurityToken(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
