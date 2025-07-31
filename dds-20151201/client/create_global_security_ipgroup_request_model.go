// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGIpList(v string) *CreateGlobalSecurityIPGroupRequest
	GetGIpList() *string
	SetGlobalIgName(v string) *CreateGlobalSecurityIPGroupRequest
	GetGlobalIgName() *string
	SetOwnerAccount(v string) *CreateGlobalSecurityIPGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateGlobalSecurityIPGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateGlobalSecurityIPGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateGlobalSecurityIPGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateGlobalSecurityIPGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CreateGlobalSecurityIPGroupRequest
	GetSecurityToken() *string
}

type CreateGlobalSecurityIPGroupRequest struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP address whitelists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27.18.86.18,183.92.40.129
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template. The name must meet the following requirements:
	//
	// 	- It can contain lowercase letters, digits, and underscores (_).
	//
	// 	- It must start with a letter and end with a letter or digit.
	//
	// 	- It must be 2 to 120 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// aizy
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s CreateGlobalSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalSecurityIPGroupRequest) GetGIpList() *string {
	return s.GIpList
}

func (s *CreateGlobalSecurityIPGroupRequest) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *CreateGlobalSecurityIPGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateGlobalSecurityIPGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateGlobalSecurityIPGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGlobalSecurityIPGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateGlobalSecurityIPGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateGlobalSecurityIPGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateGlobalSecurityIPGroupRequest) SetGIpList(v string) *CreateGlobalSecurityIPGroupRequest {
	s.GIpList = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetGlobalIgName(v string) *CreateGlobalSecurityIPGroupRequest {
	s.GlobalIgName = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetOwnerAccount(v string) *CreateGlobalSecurityIPGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetOwnerId(v int64) *CreateGlobalSecurityIPGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetRegionId(v string) *CreateGlobalSecurityIPGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetResourceOwnerAccount(v string) *CreateGlobalSecurityIPGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetResourceOwnerId(v int64) *CreateGlobalSecurityIPGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) SetSecurityToken(v string) *CreateGlobalSecurityIPGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGlobalSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
