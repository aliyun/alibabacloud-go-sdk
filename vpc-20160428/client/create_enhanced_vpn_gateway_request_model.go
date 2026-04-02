// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnhancedVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEnhancedVpnGatewayRequest
	GetClientToken() *string
	SetDisasterRecoveryVSwitchId(v string) *CreateEnhancedVpnGatewayRequest
	GetDisasterRecoveryVSwitchId() *string
	SetGatewayType(v string) *CreateEnhancedVpnGatewayRequest
	GetGatewayType() *string
	SetName(v string) *CreateEnhancedVpnGatewayRequest
	GetName() *string
	SetNetworkType(v string) *CreateEnhancedVpnGatewayRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateEnhancedVpnGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateEnhancedVpnGatewayRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateEnhancedVpnGatewayRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateEnhancedVpnGatewayRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateEnhancedVpnGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateEnhancedVpnGatewayRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *CreateEnhancedVpnGatewayRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateEnhancedVpnGatewayRequest
	GetVpcId() *string
	SetVpnType(v string) *CreateEnhancedVpnGatewayRequest
	GetVpnType() *string
}

type CreateEnhancedVpnGatewayRequest struct {
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// vsw-p0wiz7obm0tbimu4r****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enhanced.SiteToSite
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// example:
	//
	// MYVPN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// public
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// vsw-bp1j5miw2bae9s2vt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1ub1yt9cvakoelj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// Normal
	VpnType *string `json:"VpnType,omitempty" xml:"VpnType,omitempty"`
}

func (s CreateEnhancedVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnhancedVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateEnhancedVpnGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEnhancedVpnGatewayRequest) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *CreateEnhancedVpnGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateEnhancedVpnGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateEnhancedVpnGatewayRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateEnhancedVpnGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateEnhancedVpnGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateEnhancedVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnhancedVpnGatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEnhancedVpnGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateEnhancedVpnGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateEnhancedVpnGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEnhancedVpnGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEnhancedVpnGatewayRequest) GetVpnType() *string {
	return s.VpnType
}

func (s *CreateEnhancedVpnGatewayRequest) SetClientToken(v string) *CreateEnhancedVpnGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetDisasterRecoveryVSwitchId(v string) *CreateEnhancedVpnGatewayRequest {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetGatewayType(v string) *CreateEnhancedVpnGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetName(v string) *CreateEnhancedVpnGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetNetworkType(v string) *CreateEnhancedVpnGatewayRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetOwnerAccount(v string) *CreateEnhancedVpnGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetOwnerId(v int64) *CreateEnhancedVpnGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetRegionId(v string) *CreateEnhancedVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetResourceGroupId(v string) *CreateEnhancedVpnGatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetResourceOwnerAccount(v string) *CreateEnhancedVpnGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetResourceOwnerId(v int64) *CreateEnhancedVpnGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetVSwitchId(v string) *CreateEnhancedVpnGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetVpcId(v string) *CreateEnhancedVpnGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) SetVpnType(v string) *CreateEnhancedVpnGatewayRequest {
	s.VpnType = &v
	return s
}

func (s *CreateEnhancedVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
