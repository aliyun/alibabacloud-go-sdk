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
	SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyGlobalSecurityIPGroupRequest struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas (,). You can create up to 1,000 IP addresses or CIDR blocks for all IP address whitelists.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27.16.214.10,111.60.117.181
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// dps
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP whitelist template.
	//
	// This parameter is required.
	//
	// example:
	//
	// g-fwjk62egbsrp4sftxmmr
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

func (s *ModifyGlobalSecurityIPGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyGlobalSecurityIPGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *ModifyGlobalSecurityIPGroupRequest) SetResourceOwnerAccount(v string) *ModifyGlobalSecurityIPGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) SetResourceOwnerId(v int64) *ModifyGlobalSecurityIPGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
