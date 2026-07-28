// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatIpAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyNatIpAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyNatIpAttributeRequest
	GetDryRun() *bool
	SetNatIpDescription(v string) *ModifyNatIpAttributeRequest
	GetNatIpDescription() *string
	SetNatIpId(v string) *ModifyNatIpAttributeRequest
	GetNatIpId() *string
	SetNatIpName(v string) *ModifyNatIpAttributeRequest
	GetNatIpName() *string
	SetOwnerAccount(v string) *ModifyNatIpAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNatIpAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNatIpAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNatIpAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNatIpAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyNatIpAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a check request without modifying the NAT IP address information. The system checks the request for potential issues, including invalid AccessKey pairs, the authorization status of the RAM user, and missing parameter values. If the request fails the dry run, the corresponding error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, a 2xx HTTP status code is returned and the NAT IP address information is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the NAT IP address that you want to modify.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	NatIpDescription *string `json:"NatIpDescription,omitempty" xml:"NatIpDescription,omitempty"`
	// The instance ID of the NAT IP address that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpcnatip-gw8e1n11f44wpg****
	NatIpId *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	// The name of the NAT IP address that you want to modify.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newname
	NatIpName    *string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway instance to which the NAT IP address belongs.
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

func (s ModifyNatIpAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatIpAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatIpAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyNatIpAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyNatIpAttributeRequest) GetNatIpDescription() *string {
	return s.NatIpDescription
}

func (s *ModifyNatIpAttributeRequest) GetNatIpId() *string {
	return s.NatIpId
}

func (s *ModifyNatIpAttributeRequest) GetNatIpName() *string {
	return s.NatIpName
}

func (s *ModifyNatIpAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNatIpAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNatIpAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatIpAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNatIpAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNatIpAttributeRequest) SetClientToken(v string) *ModifyNatIpAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetDryRun(v bool) *ModifyNatIpAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetNatIpDescription(v string) *ModifyNatIpAttributeRequest {
	s.NatIpDescription = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetNatIpId(v string) *ModifyNatIpAttributeRequest {
	s.NatIpId = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetNatIpName(v string) *ModifyNatIpAttributeRequest {
	s.NatIpName = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetOwnerAccount(v string) *ModifyNatIpAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetOwnerId(v int64) *ModifyNatIpAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetRegionId(v string) *ModifyNatIpAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetResourceOwnerAccount(v string) *ModifyNatIpAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) SetResourceOwnerId(v int64) *ModifyNatIpAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNatIpAttributeRequest) Validate() error {
	return dara.Validate(s)
}
