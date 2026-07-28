// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatIpCidrAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyNatIpCidrAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyNatIpCidrAttributeRequest
	GetDryRun() *bool
	SetNatGatewayId(v string) *ModifyNatIpCidrAttributeRequest
	GetNatGatewayId() *string
	SetNatIpCidr(v string) *ModifyNatIpCidrAttributeRequest
	GetNatIpCidr() *string
	SetNatIpCidrDescription(v string) *ModifyNatIpCidrAttributeRequest
	GetNatIpCidrDescription() *string
	SetNatIpCidrName(v string) *ModifyNatIpCidrAttributeRequest
	GetNatIpCidrName() *string
	SetOwnerAccount(v string) *ModifyNatIpCidrAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNatIpCidrAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNatIpCidrAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNatIpCidrAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNatIpCidrAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyNatIpCidrAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the NAT CIDR block information. The system checks whether your AccessKey pair is valid, whether Resource Access Management (RAM) user authorization is granted, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a Normal request. If the check succeeds, a 2xx HTTP status code is returned and the NAT CIDR block information is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the VPC NAT gateway to which the NAT CIDR block belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT CIDR block to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the NAT CIDR block to modify.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newtest
	NatIpCidrDescription *string `json:"NatIpCidrDescription,omitempty" xml:"NatIpCidrDescription,omitempty"`
	// The name of the NAT CIDR block to modify.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// newname
	NatIpCidrName *string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway instance to which the NAT CIDR block belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s ModifyNatIpCidrAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatIpCidrAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatIpCidrAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNatIpCidrAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyNatIpCidrAttributeRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatIpCidrAttributeRequest) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *ModifyNatIpCidrAttributeRequest) GetNatIpCidrDescription() *string {
	return s.NatIpCidrDescription
}

func (s *ModifyNatIpCidrAttributeRequest) GetNatIpCidrName() *string {
	return s.NatIpCidrName
}

func (s *ModifyNatIpCidrAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNatIpCidrAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNatIpCidrAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatIpCidrAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNatIpCidrAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNatIpCidrAttributeRequest) SetClientToken(v string) *ModifyNatIpCidrAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetDryRun(v bool) *ModifyNatIpCidrAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetNatGatewayId(v string) *ModifyNatIpCidrAttributeRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetNatIpCidr(v string) *ModifyNatIpCidrAttributeRequest {
	s.NatIpCidr = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetNatIpCidrDescription(v string) *ModifyNatIpCidrAttributeRequest {
	s.NatIpCidrDescription = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetNatIpCidrName(v string) *ModifyNatIpCidrAttributeRequest {
	s.NatIpCidrName = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetOwnerAccount(v string) *ModifyNatIpCidrAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetOwnerId(v int64) *ModifyNatIpCidrAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetRegionId(v string) *ModifyNatIpCidrAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetResourceOwnerAccount(v string) *ModifyNatIpCidrAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) SetResourceOwnerId(v int64) *ModifyNatIpCidrAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNatIpCidrAttributeRequest) Validate() error {
	return dara.Validate(s)
}
