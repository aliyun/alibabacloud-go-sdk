// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateVpnGatewayRequest
	GetAutoPay() *bool
	SetBandwidth(v int32) *CreateVpnGatewayRequest
	GetBandwidth() *int32
	SetClientToken(v string) *CreateVpnGatewayRequest
	GetClientToken() *string
	SetDisasterRecoveryVSwitchId(v string) *CreateVpnGatewayRequest
	GetDisasterRecoveryVSwitchId() *string
	SetEnableIpsec(v bool) *CreateVpnGatewayRequest
	GetEnableIpsec() *bool
	SetEnableSsl(v bool) *CreateVpnGatewayRequest
	GetEnableSsl() *bool
	SetInstanceChargeType(v string) *CreateVpnGatewayRequest
	GetInstanceChargeType() *string
	SetName(v string) *CreateVpnGatewayRequest
	GetName() *string
	SetNetworkType(v string) *CreateVpnGatewayRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateVpnGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpnGatewayRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateVpnGatewayRequest
	GetPeriod() *int32
	SetRegionId(v string) *CreateVpnGatewayRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpnGatewayRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVpnGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpnGatewayRequest
	GetResourceOwnerId() *int64
	SetSslConnections(v int32) *CreateVpnGatewayRequest
	GetSslConnections() *int32
	SetVSwitchId(v string) *CreateVpnGatewayRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateVpnGatewayRequest
	GetVpcId() *string
	SetVpnType(v string) *CreateVpnGatewayRequest
	GetVpnType() *string
}

type CreateVpnGatewayRequest struct {
	// Specifies whether to automatically pay the bill for the VPN gateway. Valid values:
	//
	// - **true**: automatically pays the bill for the VPN gateway.
	//
	// - **false*	- (default): does not automatically pay the bill for the VPN gateway.
	//
	// > To successfully create a VPN gateway instance, enable automatic payment. If you disable automatic payment, you must manually pay the bill to create the VPN gateway instance.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The bandwidth specification of the VPN gateway. Unit: Mbit/s.
	//
	// <props="china">- If you want to create a public VPN gateway, valid values are **5**, **10**, **20**, **50**, **100**, **200**, **500**, and **1000**.
	//
	// <props="china">- If you want to create a private VPN gateway, valid values are **200*	- and **1000**.
	//
	// <props="intl">- If you want to create a public VPN gateway, valid values are **10**, **100**, **200**, **500**, and **1000**.
	//
	// <props="intl">- If you want to create a private VPN gateway, valid values are **200*	- and **1000**.
	//
	// >The maximum bandwidth specification supported by VPN gateways in some regions is 500 Mbit/s. For more information, see [VPN gateway limits](https://help.aliyun.com/document_detail/65290.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The second vSwitch associated with the VPN VPC-connected instance.
	//
	// - If the current region supports dual-tunnel IPsec-VPN connections, this parameter is required.
	//
	// - You must specify two vSwitches in different zones within the VPC associated with the VPN VPC-connected instance to implement zone-level disaster recovery for IPsec-VPN connections.
	//
	// - For regions that support only one zone, zone-level disaster recovery is not supported. Specify two different vSwitches in the same zone to achieve high availability for IPsec-VPN connections. You can also specify the same vSwitch.
	//
	// For information about the regions and zones that support dual-tunnel IPsec-VPN connections, see [Upgrade an IPsec-VPN connection to dual-tunnel mode](https://help.aliyun.com/document_detail/2358946.html).
	//
	// example:
	//
	// vsw-p0wiz7obm0tbimu4r****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// Specifies whether to enable the IPsec-VPN feature. Valid values:
	//
	// - **true*	- (default): enables the IPsec-VPN feature.
	//
	// - **false**: disables the IPsec-VPN feature.
	//
	// example:
	//
	// true
	EnableIpsec *bool `json:"EnableIpsec,omitempty" xml:"EnableIpsec,omitempty"`
	// Specifies whether to enable the SSL-VPN feature. Valid values:
	//
	// - **true**: enables the SSL-VPN feature.
	//
	// - **false*	- (default): disables the SSL-VPN feature.
	//
	// example:
	//
	// false
	EnableSsl *bool `json:"EnableSsl,omitempty" xml:"EnableSsl,omitempty"`
	// <props="china">The billing method of the VPN gateway. Set the value to **PREPAY**, which specifies the subscription billing method.
	//
	// <props="intl">The billing method of the VPN gateway. Set the value to **POSTPAY**, which specifies the pay-as-you-go billing method.
	//
	// <props="partner">The billing method of the VPN gateway. Set the value to **POSTPAY**, which specifies the pay-as-you-go billing method.
	//
	// <props="china">This parameter is required when you create a VPN gateway.
	//
	// example:
	//
	// 中国站示例值：PREPAY，国际站示例值：POSTPAY
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the VPN gateway. The default value is the ID of the VPN gateway.
	//
	// The name must be 2 to 100 characters in length. It cannot start with `http://` or `https://`. It must start with an uppercase or lowercase letter and can contain uppercase and lowercase letters, digits, underscores (_), hyphens (-), and periods (.). Other special characters are not supported.
	//
	// example:
	//
	// MYVPN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the VPN gateway. Valid values:
	//
	// - **public*	- (default): public VPN gateway.
	//
	// - **private**: private VPN gateway.
	//
	// example:
	//
	// public
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: months. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// <props="china">
	//
	// > This parameter is required if **InstanceChargeType*	- is set to **PREPAY**..
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the VPN gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the VPN gateway belongs.
	//
	// - You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource group IDs.
	//
	// - If you do not specify a resource group ID, the VPN gateway belongs to the default resource group after it is created.
	//
	// - After the VPN gateway is created, if you create SSL servers, SSL client certificates, IPsec servers, or IPsec-VPN connections (when the IPsec-VPN connection is associated with the VPN gateway) under the VPN gateway, these resources belong to the same resource group as the VPN gateway. The resource group of these resources cannot be modified.
	//
	//   If you change the resource group of the VPN gateway, the resource group of the preceding resources is also changed.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum number of clients that can be simultaneously connected. Valid values: **5*	- (default), **10**, **20**, **50**, **100**, **200**, **500**, and **1000**.
	//
	// example:
	//
	// 5
	SslConnections *int32 `json:"SslConnections,omitempty" xml:"SslConnections,omitempty"`
	// The vSwitch associated with the VPN gateway instance.
	//
	// - In regions that support dual-tunnel IPsec-VPN connections, this parameter is required. You must specify a vSwitch and also specify the **DisasterRecoveryVSwitchId*	- parameter.
	//
	// - In regions that support only single-tunnel IPsec-VPN connections, if you do not specify a vSwitch, the system automatically selects a vSwitch from the VPC.
	//
	// example:
	//
	// vsw-bp1j5miw2bae9s2vt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC-connected instance to which the VPN gateway belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1ub1yt9cvakoelj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The type of the VPN gateway. Valid values:
	//
	// - **Normal*	- (default): standard.
	//
	// <props="china">- **NationalStandard**: Chinese SM-based..
	//
	// example:
	//
	// Normal
	VpnType *string `json:"VpnType,omitempty" xml:"VpnType,omitempty"`
}

func (s CreateVpnGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateVpnGatewayRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateVpnGatewayRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateVpnGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpnGatewayRequest) GetDisasterRecoveryVSwitchId() *string {
	return s.DisasterRecoveryVSwitchId
}

func (s *CreateVpnGatewayRequest) GetEnableIpsec() *bool {
	return s.EnableIpsec
}

func (s *CreateVpnGatewayRequest) GetEnableSsl() *bool {
	return s.EnableSsl
}

func (s *CreateVpnGatewayRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateVpnGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateVpnGatewayRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateVpnGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpnGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpnGatewayRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateVpnGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpnGatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpnGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpnGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpnGatewayRequest) GetSslConnections() *int32 {
	return s.SslConnections
}

func (s *CreateVpnGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVpnGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpnGatewayRequest) GetVpnType() *string {
	return s.VpnType
}

func (s *CreateVpnGatewayRequest) SetAutoPay(v bool) *CreateVpnGatewayRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetBandwidth(v int32) *CreateVpnGatewayRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetClientToken(v string) *CreateVpnGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetDisasterRecoveryVSwitchId(v string) *CreateVpnGatewayRequest {
	s.DisasterRecoveryVSwitchId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetEnableIpsec(v bool) *CreateVpnGatewayRequest {
	s.EnableIpsec = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetEnableSsl(v bool) *CreateVpnGatewayRequest {
	s.EnableSsl = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetInstanceChargeType(v string) *CreateVpnGatewayRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetName(v string) *CreateVpnGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetNetworkType(v string) *CreateVpnGatewayRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetOwnerAccount(v string) *CreateVpnGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetOwnerId(v int64) *CreateVpnGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetPeriod(v int32) *CreateVpnGatewayRequest {
	s.Period = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetRegionId(v string) *CreateVpnGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetResourceGroupId(v string) *CreateVpnGatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetResourceOwnerAccount(v string) *CreateVpnGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetResourceOwnerId(v int64) *CreateVpnGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetSslConnections(v int32) *CreateVpnGatewayRequest {
	s.SslConnections = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetVSwitchId(v string) *CreateVpnGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetVpcId(v string) *CreateVpnGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpnGatewayRequest) SetVpnType(v string) *CreateVpnGatewayRequest {
	s.VpnType = &v
	return s
}

func (s *CreateVpnGatewayRequest) Validate() error {
	return dara.Validate(s)
}
