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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
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
	// The ID of the Virtual Private Cloud (VPC) NAT gateway to which the NAT CIDR block belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT CIDR block whose name and description you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The new description of the NAT CIDR block.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newtest
	NatIpCidrDescription *string `json:"NatIpCidrDescription,omitempty" xml:"NatIpCidrDescription,omitempty"`
	// The new name of the NAT CIDR block.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// newname
	NatIpCidrName *string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway to which the NAT CIDR block belongs.
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
