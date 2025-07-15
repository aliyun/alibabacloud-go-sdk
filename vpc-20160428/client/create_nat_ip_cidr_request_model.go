// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNatIpCidrRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateNatIpCidrRequest
	GetDryRun() *bool
	SetNatGatewayId(v string) *CreateNatIpCidrRequest
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *CreateNatIpCidrRequest
	GetNatIpCidr() *string
	SetNatIpCidrDescription(v string) *CreateNatIpCidrRequest
	GetNatIpCidrDescription() *string
	SetNatIpCidrName(v string) *CreateNatIpCidrRequest
	GetNatIpCidrName() *string
	SetOwnerAccount(v string) *CreateNatIpCidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNatIpCidrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNatIpCidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNatIpCidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNatIpCidrRequest
	GetResourceOwnerId() *int64
}

type CreateNatIpCidrRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**(default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) NAT gateway with which you want to associate the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT CIDR block that you want to associate with the NAT gateway.
	//
	// The new CIDR block must meet the following conditions:
	//
	// 	- The NAT CIDR block must fall within 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or their subnets.
	//
	// 	- The subnet mask must be 16 to 32 bits in length.
	//
	// 	- The NAT CIDR block cannot overlap with the private CIDR block of the VPC to which the NAT gateway belongs. If you want to use other IP addresses from the private CIDR block of the VPC to provide NAT services, create a vSwitch and attach the vSwitch to another VPC NAT gateway.
	//
	// 	- If you want to use public IP addresses to provide NAT services, make sure that the public IP addresses fall within a customer CIDR block of the VPC to which the VPC NAT gateway belongs. For more information, see [What is customer CIDR block?](https://help.aliyun.com/document_detail/185311.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the NAT CIDR block.
	//
	// The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// mycidr
	NatIpCidrDescription *string `json:"NatIpCidrDescription,omitempty" xml:"NatIpCidrDescription,omitempty"`
	// The name of the CIDR block.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newcidr
	NatIpCidrName *string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway with which you want to associate the CIDR block.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateNatIpCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpCidrRequest) GoString() string {
	return s.String()
}

func (s *CreateNatIpCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNatIpCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateNatIpCidrRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatIpCidrRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *CreateNatIpCidrRequest) GetNatIpCidrDescription() *string {
	return s.NatIpCidrDescription
}

func (s *CreateNatIpCidrRequest) GetNatIpCidrName() *string {
	return s.NatIpCidrName
}

func (s *CreateNatIpCidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNatIpCidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNatIpCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNatIpCidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNatIpCidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNatIpCidrRequest) SetClientToken(v string) *CreateNatIpCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetDryRun(v bool) *CreateNatIpCidrRequest {
	s.DryRun = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetNatGatewayId(v string) *CreateNatIpCidrRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetNatIpCidr(v string) *CreateNatIpCidrRequest {
	s.NatIpCidr = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetNatIpCidrDescription(v string) *CreateNatIpCidrRequest {
	s.NatIpCidrDescription = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetNatIpCidrName(v string) *CreateNatIpCidrRequest {
	s.NatIpCidrName = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetOwnerAccount(v string) *CreateNatIpCidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetOwnerId(v int64) *CreateNatIpCidrRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetRegionId(v string) *CreateNatIpCidrRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetResourceOwnerAccount(v string) *CreateNatIpCidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNatIpCidrRequest) SetResourceOwnerId(v int64) *CreateNatIpCidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNatIpCidrRequest) Validate() error {
	return dara.Validate(s)
}
