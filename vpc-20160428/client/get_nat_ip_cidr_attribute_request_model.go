// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNatIpCidrAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetNatIpCidrAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *GetNatIpCidrAttributeRequest
	GetDryRun() *bool
	SetNatGatewayId(v string) *GetNatIpCidrAttributeRequest
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *GetNatIpCidrAttributeRequest
	GetNatIpCidr() *string
	SetOwnerAccount(v string) *GetNatIpCidrAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetNatIpCidrAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetNatIpCidrAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetNatIpCidrAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetNatIpCidrAttributeRequest
	GetResourceOwnerId() *int64
}

type GetNatIpCidrAttributeRequest struct {
	// Client Token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client, ensuring it is unique across different requests. ClientToken only supports ASCII characters.
	//
	// > If not specified, the system automatically uses the RequestId of the API request as the ClientToken identifier. The RequestId is different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Whether to perform a dry run of this request. Values:
	//
	// - true: Sends a check request. The checks include whether the AccessKey is valid, the RAM user\\"s authorization status, and if all required parameters are filled out. If any check fails, the corresponding error is returned. If all checks pass, an error code DryRunOperation is returned.
	//
	// - false (default): Sends a normal request. After passing the checks, a 2xx HTTP status code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the VPC NAT gateway instance to which the queried NAT IP address range belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT IP address range to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.0/24
	NatIpCidr    *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway instance to which the NAT IP address range to be queried belongs. You can obtain the region ID by calling the DescribeRegions interface.
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

func (s GetNatIpCidrAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNatIpCidrAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetNatIpCidrAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetNatIpCidrAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetNatIpCidrAttributeRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetNatIpCidrAttributeRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *GetNatIpCidrAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetNatIpCidrAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetNatIpCidrAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetNatIpCidrAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetNatIpCidrAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetNatIpCidrAttributeRequest) SetClientToken(v string) *GetNatIpCidrAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetDryRun(v bool) *GetNatIpCidrAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetNatGatewayId(v string) *GetNatIpCidrAttributeRequest {
	s.NatGatewayId = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetNatIpCidr(v string) *GetNatIpCidrAttributeRequest {
	s.NatIpCidr = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetOwnerAccount(v string) *GetNatIpCidrAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetOwnerId(v int64) *GetNatIpCidrAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetRegionId(v string) *GetNatIpCidrAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetResourceOwnerAccount(v string) *GetNatIpCidrAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) SetResourceOwnerId(v int64) *GetNatIpCidrAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetNatIpCidrAttributeRequest) Validate() error {
	return dara.Validate(s)
}
