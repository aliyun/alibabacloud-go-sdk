// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatIpCidrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteNatIpCidrRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteNatIpCidrRequest
	GetDryRun() *bool
	SetNatGatewayId(v string) *DeleteNatIpCidrRequest
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *DeleteNatIpCidrRequest
	GetNatIpCidr() *string
	SetOwnerAccount(v string) *DeleteNatIpCidrRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteNatIpCidrRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNatIpCidrRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNatIpCidrRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteNatIpCidrRequest
	GetResourceOwnerId() *int64
}

type DeleteNatIpCidrRequest struct {
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
	// The ID of the NAT gateway to which the NAT CIDR block to be deleted belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT CIDR block to be deleted.
	//
	// 	- Before you delete a NAT CIDR block, you must delete all NAT IP addresses from the CIDR block.
	//
	// 	- The default NAT CIDR block cannot be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr    *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway to which the NAT CIDR block to be deleted belongs.
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

func (s DeleteNatIpCidrRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatIpCidrRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatIpCidrRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteNatIpCidrRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteNatIpCidrRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DeleteNatIpCidrRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *DeleteNatIpCidrRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteNatIpCidrRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNatIpCidrRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNatIpCidrRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNatIpCidrRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteNatIpCidrRequest) SetClientToken(v string) *DeleteNatIpCidrRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetDryRun(v bool) *DeleteNatIpCidrRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetNatGatewayId(v string) *DeleteNatIpCidrRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetNatIpCidr(v string) *DeleteNatIpCidrRequest {
	s.NatIpCidr = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetOwnerAccount(v string) *DeleteNatIpCidrRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetOwnerId(v int64) *DeleteNatIpCidrRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetRegionId(v string) *DeleteNatIpCidrRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetResourceOwnerAccount(v string) *DeleteNatIpCidrRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNatIpCidrRequest) SetResourceOwnerId(v int64) *DeleteNatIpCidrRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteNatIpCidrRequest) Validate() error {
	return dara.Validate(s)
}
