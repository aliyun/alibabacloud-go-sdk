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
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// > To create a VPN gateway, we recommend that you enable automatic payment. If you disable automatic payment, you must manually pay the bill to create the VPN gateway.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The maximum bandwidth of the VPN gateway. Unit: Mbit/s.
	//
	// 	- If you want to create a public VPN gateway, valid values are **10**, **100**, **200**, **500**, and **1000**.
	//
	// 	- If you want to create a private VPN gateway, valid values are **200*	- and **1000**.
	//
	// >  The maximum bandwidth supported by VPN gateways in some regions is 500 Mbit/s. For more information, see [VPN gateway limits](https://help.aliyun.com/document_detail/65290.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a value, and you must make sure that each request has a unique token value. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the value of **RequestId*	- as the value of **ClientToken**. The value of **RequestId*	- for each API request is different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The second vSwitch with which you want to associate the VPN gateway.
	//
	// 	- If you call this operation in a region that supports the IPsec-VPN connections in dual-tunnel mode, this parameter is required.
	//
	// 	- You need to specify two vSwitches in different zones in the virtual private cloud (VPC) that is associated with the VPN gateway to implement disaster recovery across zones.
	//
	// 	- For a region that supports only one zone, disaster recovery across zones is not supported. We recommend that you specify two vSwitches in the zone to implement high availability. You can specify the same vSwitch.
	//
	// For more information about the regions and zones that support the IPsec-VPN connections in dual-tunnel mode, see [IPsec-VPN connections support the dual-tunnel mode](https://help.aliyun.com/document_detail/2358946.html).
	//
	// example:
	//
	// vsw-p0wiz7obm0tbimu4r****
	DisasterRecoveryVSwitchId *string `json:"DisasterRecoveryVSwitchId,omitempty" xml:"DisasterRecoveryVSwitchId,omitempty"`
	// Specifies whether to enable IPsec-VPN for the VPN gateway. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableIpsec *bool `json:"EnableIpsec,omitempty" xml:"EnableIpsec,omitempty"`
	// Specifies whether to enable SSL-VPN. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	EnableSsl *bool `json:"EnableSsl,omitempty" xml:"EnableSsl,omitempty"`
	// The billing method of the VPN gateway. Set the value to **POSTPAY**, which specifies the pay-as-you-go billing method.
	//
	// example:
	//
	// Example value for the Alibaba Cloud China site: PREPAY. Example value for the Alibaba Cloud International site: POSTPAY.
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the VPN gateway. The default value is the ID of the VPN gateway.
	//
	// The name must be 2 to 100 characters in length and cannot start with `http://` or `https://`. It must start with a letter and can contain letters, digits, underscores (_), hyphens (-), and periods (.). Other special characters are not supported.
	//
	// example:
	//
	// MYVPN
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type of the VPN gateway. Valid values:
	//
	// 	- **public*	- (default)
	//
	// 	- **private**
	//
	// example:
	//
	// public
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The subscription duration. Unit: month. Valid values: **1*	- to **9**, **12**, **24**, and **36**.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the VPN gateway. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the VPN gateway belongs.
	//
	// 	- You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource group IDs.
	//
	// 	- If you do not specify a resource group ID, the VPN gateway belongs to the default resource group.
	//
	// 	- After the VPN gateway is created, the following resources also belong to the resource group and you cannot change the resource group: SSL servers, SSL client certificates, IPsec servers, and IPsec-VPN connections.
	//
	//     If you move the VPN gateway to a new resource group, the preceding resources are also moved to the new resource group.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum number of clients that can be connected at the same time. Valid values: **5*	- (default), **10**, **20**, **50**, **100**, **200**, **500**, and **1000**.
	//
	// example:
	//
	// 5
	SslConnections *int32 `json:"SslConnections,omitempty" xml:"SslConnections,omitempty"`
	// The vSwitch with which you want to associate the VPN gateway.
	//
	// 	- If you call this operation in a region that supports the IPsec-VPN connections in dual-tunnel mode, this parameter is required. You must specify a vSwitch and specify **DisasterRecoveryVSwitchId**.
	//
	// 	- If you call this operation in a region that supports the IPsec-VPN connections in single-tunnel mode and do not specify a vSwitch, the system automatically specifies a vSwitch.
	//
	// example:
	//
	// vsw-bp1j5miw2bae9s2vt****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) where you want to create the VPN gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1ub1yt9cvakoelj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The type of the VPN gateway. Valid values:
	//
	// Set the value to **Normal*	- (default), which specifies a standard NAT gateway.
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
