// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalIgName(v string) *DeleteGlobalSecurityIPGroupRequest
	GetGlobalIgName() *string
	SetGlobalSecurityGroupId(v string) *DeleteGlobalSecurityIPGroupRequest
	GetGlobalSecurityGroupId() *string
	SetOwnerAccount(v string) *DeleteGlobalSecurityIPGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteGlobalSecurityIPGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteGlobalSecurityIPGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteGlobalSecurityIPGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteGlobalSecurityIPGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteGlobalSecurityIPGroupRequest
	GetSecurityToken() *string
}

type DeleteGlobalSecurityIPGroupRequest struct {
	// The name of the IP whitelist template. The name must meet the following requirements:
	//
	// 	- It can contain lowercase letters, digits, and underscores (_).
	//
	// 	- It must start with a letter and end with a letter or digit.
	//
	// 	- It must be 2 to 120 characters in length.
	//
	// example:
	//
	// corp
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// g-qxieqf40xjst1ngpr3jz
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteGlobalSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteGlobalSecurityIPGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetGlobalIgName(v string) *DeleteGlobalSecurityIPGroupRequest {
	s.GlobalIgName = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetGlobalSecurityGroupId(v string) *DeleteGlobalSecurityIPGroupRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetOwnerAccount(v string) *DeleteGlobalSecurityIPGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetOwnerId(v int64) *DeleteGlobalSecurityIPGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetRegionId(v string) *DeleteGlobalSecurityIPGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetResourceOwnerAccount(v string) *DeleteGlobalSecurityIPGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetResourceOwnerId(v int64) *DeleteGlobalSecurityIPGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) SetSecurityToken(v string) *DeleteGlobalSecurityIPGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteGlobalSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
