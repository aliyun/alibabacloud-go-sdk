// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNatGatewayNatTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateNatGatewayNatTypeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateNatGatewayNatTypeRequest
	GetDryRun() *bool
	SetNatGatewayId(v string) *UpdateNatGatewayNatTypeRequest
	GetNatGatewayId() *string
	SetNatType(v string) *UpdateNatGatewayNatTypeRequest
	GetNatType() *string
	SetOwnerAccount(v string) *UpdateNatGatewayNatTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateNatGatewayNatTypeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateNatGatewayNatTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateNatGatewayNatTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateNatGatewayNatTypeRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *UpdateNatGatewayNatTypeRequest
	GetVSwitchId() *string
}

type UpdateNatGatewayNatTypeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// **true**: performs a dry run without upgrading the Internet NAT gateway type. The system checks whether your AccessKey pair is valid, whether Resource Access Management (RAM) users are granted permissions, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// **false*	- (default): performs a dry run and sends the request. If the check succeeds, a 2xx HTTP status code is returned and the Internet NAT gateway type is upgraded.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the standard Internet NAT gateway that you want to upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1b0lic8uz4r6vf2****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The type of the Internet NAT gateway. Set the value to **Enhanced**, which specifies an enhanced Internet NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enhanced
	NatType      *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the standard Internet NAT gateway that you want to upgrade.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vSwitch to which the enhanced Internet NAT gateway belongs after the upgrade.
	//
	// >If you do not set this parameter, the system randomly creates the enhanced Internet NAT gateway on any vSwitch in the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp17nszybg8epodke****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s UpdateNatGatewayNatTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNatGatewayNatTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateNatGatewayNatTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateNatGatewayNatTypeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateNatGatewayNatTypeRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *UpdateNatGatewayNatTypeRequest) GetNatType() *string {
	return s.NatType
}

func (s *UpdateNatGatewayNatTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateNatGatewayNatTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateNatGatewayNatTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNatGatewayNatTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateNatGatewayNatTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateNatGatewayNatTypeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateNatGatewayNatTypeRequest) SetClientToken(v string) *UpdateNatGatewayNatTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetDryRun(v bool) *UpdateNatGatewayNatTypeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetNatGatewayId(v string) *UpdateNatGatewayNatTypeRequest {
	s.NatGatewayId = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetNatType(v string) *UpdateNatGatewayNatTypeRequest {
	s.NatType = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetOwnerAccount(v string) *UpdateNatGatewayNatTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetOwnerId(v int64) *UpdateNatGatewayNatTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetRegionId(v string) *UpdateNatGatewayNatTypeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetResourceOwnerAccount(v string) *UpdateNatGatewayNatTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetResourceOwnerId(v int64) *UpdateNatGatewayNatTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) SetVSwitchId(v string) *UpdateNatGatewayNatTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateNatGatewayNatTypeRequest) Validate() error {
	return dara.Validate(s)
}
