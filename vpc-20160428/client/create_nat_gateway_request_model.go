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
	SetAvailabilityMode(v string) *CreateNatGatewayRequest
	GetAvailabilityMode() *string
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
	// The access mode of the VPC NAT gateway for reverse endpoint access.
	//
	// example:
	//
	// MULTI_BINDED
	AccessMode *CreateNatGatewayRequestAccessMode `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" type:"Struct"`
	// Subscription-based public NAT gateways are no longer available for purchase. This parameter is deprecated.
	//
	// example:
	//
	// 无效参数
	AutoPay          *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AvailabilityMode *string `json:"AvailabilityMode,omitempty" xml:"AvailabilityMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can create the token, but you must make sure that the token is unique among different requests.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the NAT gateway.
	//
	// The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testnat
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Subscription-based public NAT gateways are no longer available for purchase. This parameter is deprecated.
	//
	// example:
	//
	// 无效参数
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The mode in which the EIP is associated with the NAT gateway. Valid values:
	//
	// - **MULTI_BINDED*	- (default): the multi-EIP-to-ENI mode.
	//
	// - **NAT**: the EIP-to-NAT gateway mode. This mode is compatible with IPv4 gateways.
	//
	//   > If the EIP is associated with the NAT gateway in EIP-to-NAT gateway mode, the EIP occupies a private IP address of the vSwitch to which the NAT gateway belongs. Make sure that the vSwitch has sufficient private IP addresses. Otherwise, the EIP fails to be associated. In EIP-to-NAT gateway mode, a NAT gateway can be associated with up to 50 EIPs.
	//
	// example:
	//
	// MULTI_BINDED
	EipBindMode *string `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	// Specifies whether to enable ICMP reply. Valid values:
	//
	// - **true*	- (default): enables ICMP reply.
	//
	// - **false**: disables ICMP reply.
	//
	// example:
	//
	// true
	IcmpReplyEnabled *bool `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	// The billing method of the NAT gateway. Set the value to:
	//
	// **PostPaid*	- (default): pay-as-you-go.
	//
	// For more information, see [Billing of public NAT gateways](https://help.aliyun.com/document_detail/48126.html) and [Billing of VPC NAT gateways](https://help.aliyun.com/document_detail/270913.html).
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The billing method of the NAT gateway. Set the value to **PayByLcu**, which indicates that the NAT gateway is a pay-as-you-go NAT gateway and is measured in LCUs.
	//
	// example:
	//
	// PayByLcu
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP address prefix. NAT IP addresses are created from the prefix. Use a reserved CIDR block that is not allocated in the vSwitch to which the NAT gateway belongs.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The name of the NAT gateway.
	//
	// Must be 2 to 128 characters in length, start with a letter or a Chinese character, and can contain digits, underscores (_), and hyphens (-).
	//
	// If you do not specify this parameter, the system automatically specifies a name for the NAT gateway.
	//
	// example:
	//
	// fortest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address of the NAT gateway. Use an idle IP address from the CIDR block of the vSwitch to which the NAT gateway belongs. If this parameter is left empty, an IP address is randomly assigned.
	//
	// example:
	//
	// 192.168.0.2
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The type of NAT gateway. Set the value to **Enhanced**, which specifies an enhanced NAT gateway.
	//
	// example:
	//
	// Enhanced
	NatType *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	// The type of the NAT gateway to be created. Valid values:
	//
	// - **internet**: a public NAT gateway
	//
	// - **intranet**: a VPC NAT gateway
	//
	// example:
	//
	// internet
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Subscription-based public NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// 无效参数
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Specifies whether to enable PrivateLink. Valid values:
	//
	// - true: enables PrivateLink.
	//
	// - false (default): disables PrivateLink.
	//
	// example:
	//
	// false
	PrivateLinkEnabled *bool `json:"PrivateLinkEnabled,omitempty" xml:"PrivateLinkEnabled,omitempty"`
	// The ID of the region in which to create the NAT gateway.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
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
	// - **false*	- (default): disables the firewall feature.
	//
	//   	Notice:
	//
	//   This parameter is deprecated.
	//
	// example:
	//
	// false
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// Subscription-based public NAT gateways are no longer available for purchase. This parameter is deprecated.
	//
	// example:
	//
	// 无效参数
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The tags.
	//
	// example:
	//
	// MULTI_BINDED
	Tag []*CreateNatGatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the NAT gateway belongs.
	//
	// When you create a NAT gateway, you must specify a vSwitch to which the NAT gateway belongs. The system then assigns a private IP address to the NAT gateway from the vSwitch.
	//
	// - To create a NAT gateway in an existing vSwitch, make sure that the zone to which the vSwitch belongs supports NAT gateways and that the vSwitch has idle IP addresses.
	//
	// - If you have not created a vSwitch, create a vSwitch in a zone that supports NAT gateways and then specify the vSwitch.
	//
	// > Call the [ListEnhancedNatGatewayAvailableZones](https://help.aliyun.com/document_detail/182292.html) operation to query available zones and [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) to query the number of idle IP addresses in a vSwitch.
	//
	// example:
	//
	// vsw-bp1e3se98n9fq8hle****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) where you want to create the NAT gateway.
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

func (s *CreateNatGatewayRequest) GetAvailabilityMode() *string {
	return s.AvailabilityMode
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

func (s *CreateNatGatewayRequest) SetAvailabilityMode(v string) *CreateNatGatewayRequest {
	s.AvailabilityMode = &v
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
	// The access mode. Valid values:
	//
	// - **route**
	//
	// - **tunnel**
	//
	// > This parameter is valid only when **PrivateLinkEnabled*	- is set to **true**.
	//
	// example:
	//
	// route
	ModeValue *string `json:"ModeValue,omitempty" xml:"ModeValue,omitempty"`
	// The type of tunnel. Valid value:
	//
	// - **geneve**: Geneve.
	//
	// > This parameter is valid only when the access mode is tunnel.
	//
	// example:
	//
	// geneve
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
	// The tag key. You can specify up to 20 tag keys. The tag key cannot be an empty string. The tag key must be 1 to 128 characters in length and cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The tag value can be an empty string. The tag value must be 0 to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
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
