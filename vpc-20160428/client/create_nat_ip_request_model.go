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
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) NAT gateway for which you want to create the NAT IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT IP address that you want to create.
	//
	// If you do not specify an IP address, the system randomly allocates an IP address from the specified CIDR block.
	//
	// example:
	//
	// 192.168.0.34
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The CIDR block to which the NAT IP address belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the NAT IP address.
	//
	// The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	NatIpDescription *string `json:"NatIpDescription,omitempty" xml:"NatIpDescription,omitempty"`
	// The name of the NAT IP address.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newnatip
	NatIpName    *string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway to which the NAT IP address that you want to create belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent list of regions.
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
