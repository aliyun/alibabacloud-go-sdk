// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v *CreateNatGatewayRequestAccessMode) *CreateNatGatewayRequest
	GetAccessMode() *CreateNatGatewayRequestAccessMode
	SetAutoPay(v bool) *CreateNatGatewayRequest
	GetAutoPay() *bool
	SetClientToken(v string) *CreateNatGatewayRequest
	GetClientToken() *string
	SetDescription(v string) *CreateNatGatewayRequest
	GetDescription() *string
	SetDuration(v string) *CreateNatGatewayRequest
	GetDuration() *string
	SetEipBindMode(v string) *CreateNatGatewayRequest
	GetEipBindMode() *string
	SetIcmpReplyEnabled(v bool) *CreateNatGatewayRequest
	GetIcmpReplyEnabled() *bool
	SetInstanceChargeType(v string) *CreateNatGatewayRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *CreateNatGatewayRequest
	GetInternetChargeType() *string
	SetIpv4Prefix(v string) *CreateNatGatewayRequest
	GetIpv4Prefix() *string
	SetName(v string) *CreateNatGatewayRequest
	GetName() *string
	SetNatIp(v string) *CreateNatGatewayRequest
	GetNatIp() *string
	SetNatType(v string) *CreateNatGatewayRequest
	GetNatType() *string
	SetNetworkType(v string) *CreateNatGatewayRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateNatGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNatGatewayRequest
	GetOwnerId() *int64
	SetPricingCycle(v string) *CreateNatGatewayRequest
	GetPricingCycle() *string
	SetPrivateLinkEnabled(v bool) *CreateNatGatewayRequest
	GetPrivateLinkEnabled() *bool
	SetRegionId(v string) *CreateNatGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNatGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNatGatewayRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionEnabled(v bool) *CreateNatGatewayRequest
	GetSecurityProtectionEnabled() *bool
	SetSpec(v string) *CreateNatGatewayRequest
	GetSpec() *string
	SetTag(v []*CreateNatGatewayRequestTag) *CreateNatGatewayRequest
	GetTag() []*CreateNatGatewayRequestTag
	SetVSwitchId(v string) *CreateNatGatewayRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateNatGatewayRequest
	GetVpcId() *string
}

type CreateNatGatewayRequest struct {
	// The access mode for reverse access to the VPC NAT gateway.
	AccessMode *CreateNatGatewayRequestAccessMode `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" type:"Struct"`
	// Subscription Internet NAT gateways are no longer available for purchase. Ignore this parameter.
	//
	// example:
	//
	// Invalid parameter.
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the NAT gateway.
	//
	// You can leave this parameter empty or enter a description. If you enter a description, the description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testnat
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Subscription Internet NAT gateways are no longer available for purchase. Ignore this parameter.
	//
	// example:
	//
	// Invalid parameter.
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The mode in which the EIP is associated with the NAT gateway. Valid values:
	//
	// - **MULTI_BINDED**(default): the multi-EIP-to-ENI mode.
	//
	// - **NAT**: NAT mode, which is compatible with IPv4 addresses.
	//
	// > If an EIP is associated with a NAT gateway in NAT mode, the EIP occupies a private IP address of the vSwitch where the NAT gateway is deployed. Make sure that the vSwitch has sufficient private IP addresses. Otherwise, EIPs cannot be associated with the NAT gateway. In NAT mode, a maximum number of 50 EIPs can be associated with each NAT gateway.
	//
	// example:
	//
	// MULTI_BINDED
	EipBindMode *string `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	// Specifies whether to enable ICMP retrieval. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IcmpReplyEnabled *bool `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	// The billing method of the NAT gateway.
	//
	// Set the value to **PostPaid*	- (pay-as-you-go), which is the default value.
	//
	// For more information, see [Internet NAT gateway billing](https://help.aliyun.com/document_detail/48126.html) and [VPC NAT gateway billing](https://help.aliyun.com/document_detail/270913.html).
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the NAT gateway. Set the value to **PayByLcu**, which specifies the pay-by-CU metering method.
	//
	// example:
	//
	// PayByLcu
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Ipv4Prefix         *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The name of the NAT gateway.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// If this parameter is not set, the system assigns a default name to the NAT gateway.
	//
	// example:
	//
	// fortest
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The type of NAT gateway. Set the value to **Enhanced**, which specifies enhanced NAT gateway.
	//
	// example:
	//
	// Enhanced
	NatType *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	// The network type of the NAT gateway. Valid values:
	//
	// 	- **internet**: Internet
	//
	// 	- **intranet**: VPC
	//
	// example:
	//
	// internet
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Subscription Internet NAT gateways are no longer available for purchase. Ignore this parameter.
	//
	// example:
	//
	// Invalid parameter.
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// PrivateLink is not supported by default. If you set the value to true, PrivateLink is supported.
	PrivateLinkEnabled *bool `json:"PrivateLinkEnabled,omitempty" xml:"PrivateLinkEnabled,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable the firewall feature. Valid values:
	//
	// 	- **false*	- (default)	Notice: This parameter is deprecated.
	//
	// example:
	//
	// false
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// Subscription Internet NAT gateways are no longer available for purchase. Ignore this parameter.
	//
	// example:
	//
	// Invalid parameter.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The tags.
	//
	// example:
	//
	// MULTI_BINDED
	Tag []*CreateNatGatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the NAT gateway is attached.
	//
	// When you create a NAT gateway, you must specify a vSwitch for the NAT gateway. Then, the system assigns an idle private IP address from the vSwitch to the NAT gateway.
	//
	// 	- To attach the NAT gateway to an existing vSwitch, make sure that the zone to which the vSwitch belongs supports NAT gateways. In addition, the vSwitch must have idle IP addresses.
	//
	// 	- If no vSwitch exists in the VPC, create a vSwitch in a zone that supports NAT gateways. Then, specify the vSwitch for the NAT gateway.
	//
	// >  You can call the [ListEnhanhcedNatGatewayAvailableZones](https://help.aliyun.com/document_detail/182292.html) operation to query zones that support NAT gateways. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query idle IP addresses in a vSwitch.
	//
	// example:
	//
	// vsw-bp1e3se98n9fq8hle****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC where you want to create the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1di7uewzmtvfuq8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNatGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequest) GetAccessMode() *CreateNatGatewayRequestAccessMode {
	return s.AccessMode
}

func (s *CreateNatGatewayRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateNatGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNatGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNatGatewayRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateNatGatewayRequest) GetEipBindMode() *string {
	return s.EipBindMode
}

func (s *CreateNatGatewayRequest) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *CreateNatGatewayRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateNatGatewayRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateNatGatewayRequest) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *CreateNatGatewayRequest) GetName() *string {
	return s.Name
}

func (s *CreateNatGatewayRequest) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateNatGatewayRequest) GetNatType() *string {
	return s.NatType
}

func (s *CreateNatGatewayRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateNatGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNatGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNatGatewayRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateNatGatewayRequest) GetPrivateLinkEnabled() *bool {
	return s.PrivateLinkEnabled
}

func (s *CreateNatGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNatGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNatGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNatGatewayRequest) GetSecurityProtectionEnabled() *bool {
	return s.SecurityProtectionEnabled
}

func (s *CreateNatGatewayRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateNatGatewayRequest) GetTag() []*CreateNatGatewayRequestTag {
	return s.Tag
}

func (s *CreateNatGatewayRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNatGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNatGatewayRequest) SetAccessMode(v *CreateNatGatewayRequestAccessMode) *CreateNatGatewayRequest {
	s.AccessMode = v
	return s
}

func (s *CreateNatGatewayRequest) SetAutoPay(v bool) *CreateNatGatewayRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNatGatewayRequest) SetClientToken(v string) *CreateNatGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNatGatewayRequest) SetDescription(v string) *CreateNatGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateNatGatewayRequest) SetDuration(v string) *CreateNatGatewayRequest {
	s.Duration = &v
	return s
}

func (s *CreateNatGatewayRequest) SetEipBindMode(v string) *CreateNatGatewayRequest {
	s.EipBindMode = &v
	return s
}

func (s *CreateNatGatewayRequest) SetIcmpReplyEnabled(v bool) *CreateNatGatewayRequest {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *CreateNatGatewayRequest) SetInstanceChargeType(v string) *CreateNatGatewayRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateNatGatewayRequest) SetInternetChargeType(v string) *CreateNatGatewayRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNatGatewayRequest) SetIpv4Prefix(v string) *CreateNatGatewayRequest {
	s.Ipv4Prefix = &v
	return s
}

func (s *CreateNatGatewayRequest) SetName(v string) *CreateNatGatewayRequest {
	s.Name = &v
	return s
}

func (s *CreateNatGatewayRequest) SetNatIp(v string) *CreateNatGatewayRequest {
	s.NatIp = &v
	return s
}

func (s *CreateNatGatewayRequest) SetNatType(v string) *CreateNatGatewayRequest {
	s.NatType = &v
	return s
}

func (s *CreateNatGatewayRequest) SetNetworkType(v string) *CreateNatGatewayRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateNatGatewayRequest) SetOwnerAccount(v string) *CreateNatGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNatGatewayRequest) SetOwnerId(v int64) *CreateNatGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetPricingCycle(v string) *CreateNatGatewayRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateNatGatewayRequest) SetPrivateLinkEnabled(v bool) *CreateNatGatewayRequest {
	s.PrivateLinkEnabled = &v
	return s
}

func (s *CreateNatGatewayRequest) SetRegionId(v string) *CreateNatGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetResourceOwnerAccount(v string) *CreateNatGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNatGatewayRequest) SetResourceOwnerId(v int64) *CreateNatGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetSecurityProtectionEnabled(v bool) *CreateNatGatewayRequest {
	s.SecurityProtectionEnabled = &v
	return s
}

func (s *CreateNatGatewayRequest) SetSpec(v string) *CreateNatGatewayRequest {
	s.Spec = &v
	return s
}

func (s *CreateNatGatewayRequest) SetTag(v []*CreateNatGatewayRequestTag) *CreateNatGatewayRequest {
	s.Tag = v
	return s
}

func (s *CreateNatGatewayRequest) SetVSwitchId(v string) *CreateNatGatewayRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNatGatewayRequest) SetVpcId(v string) *CreateNatGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNatGatewayRequest) Validate() error {
	if s.AccessMode != nil {
		if err := s.AccessMode.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNatGatewayRequestAccessMode struct {
	// Access mode. Valid values:
	//
	// - **route**: route mode
	//
	// - **tunnel**: tunnel mode
	//
	// > If this parameter is specified, you must set **PrivateLinkEnabled*	- to **true**.
	//
	// example:
	//
	// route
	ModeValue *string `json:"ModeValue,omitempty" xml:"ModeValue,omitempty"`
	// Tunnel mode type:
	//
	// - **geneve**: Geneve type
	//
	// > This value takes effect if the access mode is the tunnel mode.
	TunnelType *string `json:"TunnelType,omitempty" xml:"TunnelType,omitempty"`
}

func (s CreateNatGatewayRequestAccessMode) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequestAccessMode) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequestAccessMode) GetModeValue() *string {
	return s.ModeValue
}

func (s *CreateNatGatewayRequestAccessMode) GetTunnelType() *string {
	return s.TunnelType
}

func (s *CreateNatGatewayRequestAccessMode) SetModeValue(v string) *CreateNatGatewayRequestAccessMode {
	s.ModeValue = &v
	return s
}

func (s *CreateNatGatewayRequestAccessMode) SetTunnelType(v string) *CreateNatGatewayRequestAccessMode {
	s.TunnelType = &v
	return s
}

func (s *CreateNatGatewayRequestAccessMode) Validate() error {
	return dara.Validate(s)
}

type CreateNatGatewayRequestTag struct {
	// The tag key. The format of Tag.N.Key when you call the operation. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The format of Tag.N.Value when you call the operation. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNatGatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNatGatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNatGatewayRequestTag) SetKey(v string) *CreateNatGatewayRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNatGatewayRequestTag) SetValue(v string) *CreateNatGatewayRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNatGatewayRequestTag) Validate() error {
	return dara.Validate(s)
}
