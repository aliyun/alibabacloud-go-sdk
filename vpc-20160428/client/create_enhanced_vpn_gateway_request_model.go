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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may differ for each API request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The second vSwitch associated with the enhanced VPN VPC-connected instance.
	//
	// - This parameter is required.
	//
	// - You must specify two vSwitches in different zones within the VPC associated with the enhanced VPN VPC-connected instance to implement zone-level disaster recovery for IPsec-VPN connections.
	//
	// - For regions that support only one zone, zone-level disaster recovery is not supported. Specify two different vSwitches in the same zone to achieve high availability for IPsec-VPN connections. You can also specify the same vSwitch.
	//
	// example:
	//
	// vsw-p0wiz7obm0tbimu4r****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// The type of the enhanced VPN gateway. Valid values:
	//
	// - **Enhanced.SiteToSite**: enhanced site-to-cloud VPN that covers only IPsec functionality.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enhanced.SiteToSite
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The name of the enhanced VPN gateway. The default value is the ID of the enhanced VPN gateway.
	//
	// The name must be 2 to 100 characters in length. It cannot start with http:// or https://. It must start with an uppercase or lowercase letter and can contain uppercase and lowercase letters, digits, underscores (_), hyphens (-), and periods (.). Other special characters are not supported.
	//
	// example:
	//
	// MYVPN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the VPN gateway. Valid values:
	//
	// - **public*	- (default): public VPN gateway.
	//
	// example:
	//
	// public
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the enhanced VPN gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the enhanced VPN gateway belongs.
	//
	// - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource group IDs.
	//
	// - If you do not specify a resource group ID, the enhanced VPN gateway belongs to the default resource group after it is created.
	//
	// - After the enhanced VPN gateway is created, if you create IPsec-VPN connections under the enhanced VPN gateway (IPsec-VPN connections that are associated with the enhanced VPN gateway), these resources belong to the same resource group as the enhanced VPN gateway and cannot be modified independently. If you change the resource group of the enhanced VPN gateway, the resource group of the preceding resources is also changed.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The vSwitch associated with the enhanced VPN gateway instance.
	//
	// - You must specify a vSwitch and also specify the **DisasterRecoveryVSwitchId*	- parameter.
	//
	// example:
	//
	// vsw-bp1j5miw2bae9s2vt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the enhanced VPN gateway belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1ub1yt9cvakoelj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The type of the enhanced VPN gateway. Valid values:
	//
	// - **Normal*	- (default): standard.
	//
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
