// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNatIpRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateNatIpRequest
	GetDryRun() *bool
	SetIpv4Prefix(v string) *CreateNatIpRequest
	GetIpv4Prefix() *string
	SetIpv4PrefixCount(v int32) *CreateNatIpRequest
	GetIpv4PrefixCount() *int32
	SetNatGatewayId(v string) *CreateNatIpRequest
	GetNatGatewayId() *string
	SetNatIp(v string) *CreateNatIpRequest
	GetNatIp() *string
	SetNatIpCidr(v string) *CreateNatIpRequest
	GetNatIpCidr() *string
	SetNatIpDescription(v string) *CreateNatIpRequest
	GetNatIpDescription() *string
	SetNatIpName(v string) *CreateNatIpRequest
	GetNatIpName() *string
	SetOwnerAccount(v string) *CreateNatIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNatIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNatIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNatIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNatIpRequest
	GetResourceOwnerId() *int64
}

type CreateNatIpRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the AccessKey pair, the authorization of the Resource Access Management (RAM) user, and the required parameters. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the request passes the check, a 2xx HTTP status code is returned and the NAT IP address is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP prefix CIDR block to create.
	//
	// The IP prefix CIDR block must be within the reserved CIDR block of the vSwitch where the NAT gateway resides, and the reserved CIDR block must not be in use. The prefix mask must be /28.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The number of IP prefixes to automatically assign.
	//
	// The IP prefixes are randomly assigned from unallocated reserved CIDR blocks of the vSwitch where the NAT gateway resides. Valid values: 1 to 10.
	//
	// example:
	//
	// 1
	Ipv4PrefixCount *int32 `json:"Ipv4PrefixCount,omitempty" xml:"Ipv4PrefixCount,omitempty"`
	// The instance ID of the VPC NAT gateway to which the NAT IP address belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT IP address to create.
	//
	// If you do not specify this parameter, the system randomly assigns an IP address from the NAT CIDR block.
	//
	// example:
	//
	// 192.168.0.34
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The NAT CIDR block from which the NAT IP address is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the NAT IP address.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	NatIpDescription *string `json:"NatIpDescription,omitempty" xml:"NatIpDescription,omitempty"`
	// The name of the NAT IP address.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newnatip
	NatIpName    *string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway instance to which the NAT IP address belongs.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list.
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

func (s CreateNatIpRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpRequest) GoString() string {
	return s.String()
}

func (s *CreateNatIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNatIpRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateNatIpRequest) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *CreateNatIpRequest) GetIpv4PrefixCount() *int32 {
	return s.Ipv4PrefixCount
}

func (s *CreateNatIpRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatIpRequest) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateNatIpRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *CreateNatIpRequest) GetNatIpDescription() *string {
	return s.NatIpDescription
}

func (s *CreateNatIpRequest) GetNatIpName() *string {
	return s.NatIpName
}

func (s *CreateNatIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNatIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNatIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNatIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNatIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNatIpRequest) SetClientToken(v string) *CreateNatIpRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNatIpRequest) SetDryRun(v bool) *CreateNatIpRequest {
	s.DryRun = &v
	return s
}

func (s *CreateNatIpRequest) SetIpv4Prefix(v string) *CreateNatIpRequest {
	s.Ipv4Prefix = &v
	return s
}

func (s *CreateNatIpRequest) SetIpv4PrefixCount(v int32) *CreateNatIpRequest {
	s.Ipv4PrefixCount = &v
	return s
}

func (s *CreateNatIpRequest) SetNatGatewayId(v string) *CreateNatIpRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatIpRequest) SetNatIp(v string) *CreateNatIpRequest {
	s.NatIp = &v
	return s
}

func (s *CreateNatIpRequest) SetNatIpCidr(v string) *CreateNatIpRequest {
	s.NatIpCidr = &v
	return s
}

func (s *CreateNatIpRequest) SetNatIpDescription(v string) *CreateNatIpRequest {
	s.NatIpDescription = &v
	return s
}

func (s *CreateNatIpRequest) SetNatIpName(v string) *CreateNatIpRequest {
	s.NatIpName = &v
	return s
}

func (s *CreateNatIpRequest) SetOwnerAccount(v string) *CreateNatIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNatIpRequest) SetOwnerId(v int64) *CreateNatIpRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNatIpRequest) SetRegionId(v string) *CreateNatIpRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNatIpRequest) SetResourceOwnerAccount(v string) *CreateNatIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNatIpRequest) SetResourceOwnerId(v int64) *CreateNatIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNatIpRequest) Validate() error {
	return dara.Validate(s)
}
