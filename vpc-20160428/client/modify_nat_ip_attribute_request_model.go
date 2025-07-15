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
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the precheck, a 2xx HTTP status code is returned and the name and description of the NAT IP address are modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the NAT IP address that you want to modify.
	//
	// The description must be 2 to 256 characters in length and start with a letter. The description cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	NatIpDescription *string `json:"NatIpDescription,omitempty" xml:"NatIpDescription,omitempty"`
	// The ID of the NAT IP address that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpcnatip-gw8e1n11f44wpg****
	NatIpId *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	// The name of the NAT IP address that you want to modify.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// newname
	NatIpName    *string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway to which the NAT IP address that you want to modify belongs.
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
