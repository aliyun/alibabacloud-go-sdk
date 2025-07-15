// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateVpcIpv6CidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolType(v string) *AllocateVpcIpv6CidrRequest
	GetAddressPoolType() *string
	SetClientToken(v string) *AllocateVpcIpv6CidrRequest
	GetClientToken() *string
	SetIpv6CidrBlock(v string) *AllocateVpcIpv6CidrRequest
	GetIpv6CidrBlock() *string
	SetIpv6Isp(v string) *AllocateVpcIpv6CidrRequest
	GetIpv6Isp() *string
	SetOwnerAccount(v string) *AllocateVpcIpv6CidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AllocateVpcIpv6CidrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AllocateVpcIpv6CidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AllocateVpcIpv6CidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AllocateVpcIpv6CidrRequest
	GetResourceOwnerId() *int64
}

type AllocateVpcIpv6CidrRequest struct {
	// The type of the IPv6 address pool. Valid values:
	//
	// - **aliyun*	- (default): IPv6 CIDR block is allocated by the system.
	//
	// - **custom**: custom IPv6 CIDR block.
	//
	// example:
	//
	// custom
	AddressPoolType *string `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IPv6 CIDR block that you want to reserve.
	//
	// example:
	//
	// 2408:XXXX:0:a600::/56
	Ipv6CidrBlock *string `json:"Ipv6CidrBlock,omitempty" xml:"Ipv6CidrBlock,omitempty"`
	// The type of IPv6 CIDR block. Valid values:
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP)
	//
	// 	- **BGP_International**: BGP (Multi-ISP) International
	//
	// 	- **ChinaMobile**: China Mobile (Single-ISP)
	//
	// 	- **ChinaUnicom**: China Unicom (Single-ISP)
	//
	// 	- **ChinaTelecom**: China Telecom (Single-ISP)
	//
	// 	- **ChinaMobile_L2**: China Mobile (Single-ISP)_L2
	//
	// 	- **ChinaUnicom_L2**: China Unicom (Single-ISP)_L2
	//
	// 	- **ChinaTelecom_L2**: China Telecom (Single-ISP)_L2
	//
	// > 	- If your account is included in the whitelist, you can set this parameter to one of the following values: **ChinaTelecom**, **ChinaUnicom**, **ChinaMobile**, **ChinaTelecom_L2**, **ChinaUnicom_L2**, **ChinaMobile_L2**, and **BGP_International**.
	//
	// > 	- You can reserve only one IPv6 CIDR block of each type. You can reserve another IPv6 CIDR block only after the existing one is allocated to a VPC.
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
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AllocateVpcIpv6CidrRequest) String() string {
	return dara.Prettify(s)
}

func (s AllocateVpcIpv6CidrRequest) GoString() string {
	return s.String()
}

func (s *AllocateVpcIpv6CidrRequest) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *AllocateVpcIpv6CidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AllocateVpcIpv6CidrRequest) GetIpv6CidrBlock() *string {
	return s.Ipv6CidrBlock
}

func (s *AllocateVpcIpv6CidrRequest) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *AllocateVpcIpv6CidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AllocateVpcIpv6CidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AllocateVpcIpv6CidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AllocateVpcIpv6CidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AllocateVpcIpv6CidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AllocateVpcIpv6CidrRequest) SetAddressPoolType(v string) *AllocateVpcIpv6CidrRequest {
	s.AddressPoolType = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetClientToken(v string) *AllocateVpcIpv6CidrRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetIpv6CidrBlock(v string) *AllocateVpcIpv6CidrRequest {
	s.Ipv6CidrBlock = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetIpv6Isp(v string) *AllocateVpcIpv6CidrRequest {
	s.Ipv6Isp = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetOwnerAccount(v string) *AllocateVpcIpv6CidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetOwnerId(v int64) *AllocateVpcIpv6CidrRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetRegionId(v string) *AllocateVpcIpv6CidrRequest {
	s.RegionId = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetResourceOwnerAccount(v string) *AllocateVpcIpv6CidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) SetResourceOwnerId(v int64) *AllocateVpcIpv6CidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateVpcIpv6CidrRequest) Validate() error {
	return dara.Validate(s)
}
