// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *ModifyVpcAttributeRequest
	GetCidrBlock() *string
	SetDescription(v string) *ModifyVpcAttributeRequest
	GetDescription() *string
	SetEnableDnsHostname(v bool) *ModifyVpcAttributeRequest
	GetEnableDnsHostname() *bool
	SetEnableIPv6(v bool) *ModifyVpcAttributeRequest
	GetEnableIPv6() *bool
	SetIpv6CidrBlock(v string) *ModifyVpcAttributeRequest
	GetIpv6CidrBlock() *string
	SetIpv6Isp(v string) *ModifyVpcAttributeRequest
	GetIpv6Isp() *string
	SetOwnerAccount(v string) *ModifyVpcAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyVpcAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyVpcAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyVpcAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyVpcAttributeRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *ModifyVpcAttributeRequest
	GetVpcId() *string
	SetVpcName(v string) *ModifyVpcAttributeRequest
	GetVpcName() *string
}

type ModifyVpcAttributeRequest struct {
	// The new IPv4 CIDR block of the VPC.
	//
	// You can specify a larger or smaller IPv4 CIDR block than the IPv4 CIDR block of the VPC. The subnet mask must be 8 to 28 bits in length. If you specify a smaller IPv4 CIDR block and existing IP addresses do not fall within the CIDR block, the modification fails.
	//
	// >  If you modify the CIDR block of a VPC, your existing services are not affected.
	//
	// example:
	//
	// 192.168.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The new description of the VPC.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my VPC.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the DNS hostname feature is enabled. Valid values:
	//
	// 	- **false*	- (default): disabled.
	//
	// 	- **true**: enabled.
	//
	// example:
	//
	// false
	EnableDnsHostname *bool `json:"EnableDnsHostname,omitempty" xml:"EnableDnsHostname,omitempty"`
	// Specifies whether to enable IPv6 CIDR blocks. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	EnableIPv6 *bool `json:"EnableIPv6,omitempty" xml:"EnableIPv6,omitempty"`
	// The IPv6 CIDR block of the VPC.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The type of IPv6 CIDR block. Valid values:
	//
	// 	- **BGP*	- (default)
	//
	// 	- **ChinaMobile**
	//
	// 	- **ChinaUnicom**
	//
	// 	- **ChinaTelecom**
	//
	// >  If your Alibaba Cloud account is allowed to activate single-ISP bandwidth, you can set this parameter to **ChinaTelecom**, **ChinaUnicom**, or **ChinaMobile**.
	//
	// example:
	//
	// BGP
	Ipv6Isp      *string `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VPC that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1qtbach57ywecf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The new name of the VPC.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Vpc-1
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ModifyVpcAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcAttributeRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ModifyVpcAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVpcAttributeRequest) GetEnableDnsHostname() *bool {
	return s.EnableDnsHostname
}

func (s *ModifyVpcAttributeRequest) GetEnableIPv6() *bool {
	return s.EnableIPv6
}

func (s *ModifyVpcAttributeRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *ModifyVpcAttributeRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *ModifyVpcAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyVpcAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVpcAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVpcAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyVpcAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyVpcAttributeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyVpcAttributeRequest) GetVpcName() *string {
	return s.VpcName
}

func (s *ModifyVpcAttributeRequest) SetCidrBlock(v string) *ModifyVpcAttributeRequest {
	s.CidrBlock = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetDescription(v string) *ModifyVpcAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetEnableDnsHostname(v bool) *ModifyVpcAttributeRequest {
	s.EnableDnsHostname = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetEnableIPv6(v bool) *ModifyVpcAttributeRequest {
	s.EnableIPv6 = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetIpv6CidrBlock(v string) *ModifyVpcAttributeRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetIpv6Isp(v string) *ModifyVpcAttributeRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetOwnerAccount(v string) *ModifyVpcAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetOwnerId(v int64) *ModifyVpcAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetRegionId(v string) *ModifyVpcAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetResourceOwnerAccount(v string) *ModifyVpcAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetResourceOwnerId(v int64) *ModifyVpcAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetVpcId(v string) *ModifyVpcAttributeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyVpcAttributeRequest) SetVpcName(v string) *ModifyVpcAttributeRequest {
	s.VpcName = &v
	return s
}

func (s *ModifyVpcAttributeRequest) Validate() error {
	return dara.Validate(s)
}
