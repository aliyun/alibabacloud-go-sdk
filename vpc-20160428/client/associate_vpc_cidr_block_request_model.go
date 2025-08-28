// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateVpcCidrBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIPv6CidrBlock(v string) *AssociateVpcCidrBlockRequest
	GetIPv6CidrBlock() *string
	SetIpVersion(v string) *AssociateVpcCidrBlockRequest
	GetIpVersion() *string
	SetIpamPoolId(v string) *AssociateVpcCidrBlockRequest
	GetIpamPoolId() *string
	SetIpv6CidrMask(v int32) *AssociateVpcCidrBlockRequest
	GetIpv6CidrMask() *int32
	SetIpv6Isp(v string) *AssociateVpcCidrBlockRequest
	GetIpv6Isp() *string
	SetOwnerAccount(v string) *AssociateVpcCidrBlockRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateVpcCidrBlockRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateVpcCidrBlockRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateVpcCidrBlockRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateVpcCidrBlockRequest
	GetResourceOwnerId() *int64
	SetSecondaryCidrBlock(v string) *AssociateVpcCidrBlockRequest
	GetSecondaryCidrBlock() *string
	SetSecondaryCidrMask(v int32) *AssociateVpcCidrBlockRequest
	GetSecondaryCidrMask() *int32
	SetVpcId(v string) *AssociateVpcCidrBlockRequest
	GetVpcId() *string
}

type AssociateVpcCidrBlockRequest struct {
	// The IPv6 CIDR block that you want to add to the VPC.
	//
	// >  You can specify only one of **SecondaryCidrBlock*	- and **Ipv6CidrBlock**.
	//
	// example:
	//
	// 2408:XXXX:0:6a::/56
	IPv6CidrBlock *string `json:"IPv6CidrBlock,omitempty" xml:"IPv6CidrBlock,omitempty"`
	// The version of the IP address. Valid values:
	//
	// 	- **IPV4**: the IPv4 address.
	//
	// 	- **IPV6**: the IPv6 address. If you set **IpVersion*	- to **IPV6*	- and do not specify **SecondaryCidrBlock**, you can add a secondary IPv6 CIDR block to the VPC.
	//
	// example:
	//
	// IPV4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-sycmt3p2a9v63i****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// Add an IPv6 CIDR block from the IPAM pool to the VPC by entering a mask.
	//
	// >  To add an IPv6 CIDR block to a VPC, specify at least one of the IPv6CidrBlock and Ipv6CidrMask parameters.
	//
	// example:
	//
	// 56
	Ipv6CidrMask *int32 `json:"Ipv6CidrMask,omitempty" xml:"Ipv6CidrMask,omitempty"`
	// The type of the IPv6 CIDR block. Valid values:
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
	// The region ID of the VPC to which you want to add a secondary CIDR block.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// ch-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IPv4 CIDR block to be added. Take note of the following requirements:
	//
	// 	- You can specify one of the following standard IPv4 CIDR blocks or their subnets as the secondary IPv4 CIDR block of the VPC: 192.168.0.0/16, 172.16.0.0/12, and 10.0.0.0/8.
	//
	// 	- You can also use a custom CIDR block other than 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, 169.254.0.0/16, or their subnets as the secondary IPv4 CIDR block of the VPC.
	//
	// The CIDR block must meet the following requirements:
	//
	// 	- The CIDR block cannot start with 0. The subnet mask must be 8 to 28 bits in length.
	//
	// 	- The CIDR block cannot overlap with the primary CIDR block or an existing secondary CIDR block of the VPC.
	//
	// >  You must and can specify only one of **SecondaryCidrBlock*	- and **Ipv6CidrBlock**.
	//
	// example:
	//
	// 192.168.0.0/16
	SecondaryCidrBlock *string `json:"SecondaryCidrBlock,omitempty" xml:"SecondaryCidrBlock,omitempty"`
	// Add an IPv4 CIDR block from the IPAM pool to the VPC by specifying a mask.
	//
	// >  If you use an IPAM pool, you must specify at least one of SecondaryCidrBlock and SecondaryCidrMask.
	//
	// example:
	//
	// 16
	SecondaryCidrMask *int32 `json:"SecondaryCidrMask,omitempty" xml:"SecondaryCidrMask,omitempty"`
	// The ID of the VPC to which you want to add a secondary CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-o6wrloqsdqc9io3mg****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s AssociateVpcCidrBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateVpcCidrBlockRequest) GoString() string {
	return s.String()
}

func (s *AssociateVpcCidrBlockRequest) GetIPv6CidrBlock() *string {
	return s.IPv6CidrBlock
}

func (s *AssociateVpcCidrBlockRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *AssociateVpcCidrBlockRequest) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *AssociateVpcCidrBlockRequest) GetIpv6CidrMask() *int32 {
	return s.Ipv6CidrMask
}

func (s *AssociateVpcCidrBlockRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *AssociateVpcCidrBlockRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateVpcCidrBlockRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateVpcCidrBlockRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateVpcCidrBlockRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateVpcCidrBlockRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateVpcCidrBlockRequest) GetSecondaryCidrBlock() *string {
	return s.SecondaryCidrBlock
}

func (s *AssociateVpcCidrBlockRequest) GetSecondaryCidrMask() *int32 {
	return s.SecondaryCidrMask
}

func (s *AssociateVpcCidrBlockRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *AssociateVpcCidrBlockRequest) SetIPv6CidrBlock(v string) *AssociateVpcCidrBlockRequest {
	s.IPv6CidrBlock = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetIpVersion(v string) *AssociateVpcCidrBlockRequest {
	s.IpVersion = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetIpamPoolId(v string) *AssociateVpcCidrBlockRequest {
	s.IpamPoolId = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetIpv6CidrMask(v int32) *AssociateVpcCidrBlockRequest {
	s.Ipv6CidrMask = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetIpv6Isp(v string) *AssociateVpcCidrBlockRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetOwnerAccount(v string) *AssociateVpcCidrBlockRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetOwnerId(v int64) *AssociateVpcCidrBlockRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetRegionId(v string) *AssociateVpcCidrBlockRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetResourceOwnerAccount(v string) *AssociateVpcCidrBlockRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetResourceOwnerId(v int64) *AssociateVpcCidrBlockRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetSecondaryCidrBlock(v string) *AssociateVpcCidrBlockRequest {
	s.SecondaryCidrBlock = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetSecondaryCidrMask(v int32) *AssociateVpcCidrBlockRequest {
	s.SecondaryCidrMask = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) SetVpcId(v string) *AssociateVpcCidrBlockRequest {
	s.VpcId = &v
	return s
}

func (s *AssociateVpcCidrBlockRequest) Validate() error {
	return dara.Validate(s)
}
