// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessModeShrink(v string) *CreateNatGatewayShrinkRequest
	GetAccessModeShrink() *string
	SetAutoPay(v bool) *CreateNatGatewayShrinkRequest
	GetAutoPay() *bool
	SetAvailabilityMode(v string) *CreateNatGatewayShrinkRequest
	GetAvailabilityMode() *string
	SetClientToken(v string) *CreateNatGatewayShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateNatGatewayShrinkRequest
	GetDescription() *string
	SetDuration(v string) *CreateNatGatewayShrinkRequest
	GetDuration() *string
	SetEipBindMode(v string) *CreateNatGatewayShrinkRequest
	GetEipBindMode() *string
	SetIcmpReplyEnabled(v bool) *CreateNatGatewayShrinkRequest
	GetIcmpReplyEnabled() *bool
	SetInstanceChargeType(v string) *CreateNatGatewayShrinkRequest
	GetInstanceChargeType() *string
	SetInternetChargeType(v string) *CreateNatGatewayShrinkRequest
	GetInternetChargeType() *string
	SetIpv4Prefix(v string) *CreateNatGatewayShrinkRequest
	GetIpv4Prefix() *string
	SetName(v string) *CreateNatGatewayShrinkRequest
	GetName() *string
	SetNatIp(v string) *CreateNatGatewayShrinkRequest
	GetNatIp() *string
	SetNatType(v string) *CreateNatGatewayShrinkRequest
	GetNatType() *string
	SetNetworkType(v string) *CreateNatGatewayShrinkRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateNatGatewayShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNatGatewayShrinkRequest
	GetOwnerId() *int64
	SetPricingCycle(v string) *CreateNatGatewayShrinkRequest
	GetPricingCycle() *string
	SetPrivateLinkEnabled(v bool) *CreateNatGatewayShrinkRequest
	GetPrivateLinkEnabled() *bool
	SetRegionId(v string) *CreateNatGatewayShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNatGatewayShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNatGatewayShrinkRequest
	GetResourceOwnerId() *int64
	SetSecurityProtectionEnabled(v bool) *CreateNatGatewayShrinkRequest
	GetSecurityProtectionEnabled() *bool
	SetSpec(v string) *CreateNatGatewayShrinkRequest
	GetSpec() *string
	SetTag(v []*CreateNatGatewayShrinkRequestTag) *CreateNatGatewayShrinkRequest
	GetTag() []*CreateNatGatewayShrinkRequestTag
	SetVSwitchId(v string) *CreateNatGatewayShrinkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateNatGatewayShrinkRequest
	GetVpcId() *string
}

type CreateNatGatewayShrinkRequest struct {
	// The access mode for reverse access to the VPC NAT gateway.
	//
	// example:
	//
	// MULTI_BINDED
	AccessModeShrink *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// Subscription-based Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// 无效参数
	AutoPay          *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AvailabilityMode *string `json:"AvailabilityMode,omitempty" xml:"AvailabilityMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the NAT gateway.
	//
	// The description can be empty or 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testnat
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Subscription-based Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// 无效参数
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The EIP binding pattern of the NAT gateway. Valid values:
	//
	// - **MULTI_BINDED*	- (default): multi-EIP network interface controller (NIC)-visible pattern.
	//
	// - **NAT**: EIP normal pattern, compatible with IPv4 gateways.
	//
	//   > When the EIP binding pattern of the NAT gateway is set to EIP normal pattern, each EIP occupies a private IP address in the vSwitch to which the NAT gateway belongs. Make sure that the vSwitch has sufficient private IP addresses. If no idle private IP addresses are available in the vSwitch, new EIPs cannot be attached. In EIP normal pattern, a maximum of 50 EIPs can be attached to the NAT gateway.
	//
	// example:
	//
	// MULTI_BINDED
	EipBindMode *string `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	// Specifies whether to enable ICMP echo reply. Valid values:
	//
	// - **true*	- (default): Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	IcmpReplyEnabled *bool `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	// The billing method of the NAT gateway. Valid values:
	//
	// **PostPaid*	- (default): pay-as-you-go.
	//
	//
	//
	// For more information, see [Billing of Internet NAT gateways](https://help.aliyun.com/document_detail/48126.html) and [Billing of VPC NAT gateways](https://help.aliyun.com/document_detail/270913.html).
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the NAT gateway. Valid values: **PayByLcu**: pay-by-data-transfer.
	//
	// example:
	//
	// PayByLcu
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP prefix CIDR block used to create NAT IP addresses in batches. Specify an unassigned reserved CIDR block from the vSwitch to which the NAT gateway belongs.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The name of the NAT gateway.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or Chinese character. It can contain digits, underscores (_), and hyphens (-).
	//
	// If you do not specify this parameter, the system assigns a default name to the NAT gateway.
	//
	// example:
	//
	// fortest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address occupied by the NAT gateway. Specify an unallocated IP address from the vSwitch CIDR block to which the NAT gateway belongs. If this parameter is left empty, a random IP address is allocated.
	//
	// example:
	//
	// 192.168.0.2
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The type of the NAT gateway. Valid values: **Enhanced**: enhanced NAT gateway.
	//
	// example:
	//
	// Enhanced
	NatType *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	// The type of NAT gateway to create. Valid values:
	//
	// - **internet**: Internet NAT gateway.
	//
	// - **intranet**: VPC NAT gateway.
	//
	// example:
	//
	// internet
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Subscription-based Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// 无效参数
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// Specifies whether to enable PrivateLink. Valid values:
	//
	// - true: Enabled.
	//
	// - false (default): Disabled.
	//
	// example:
	//
	// false
	PrivateLinkEnabled *bool `json:"PrivateLinkEnabled,omitempty" xml:"PrivateLinkEnabled,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
	// - **false*	- (default): Disabled.
	//
	// 	Notice:  This parameter is deprecated.
	//
	// example:
	//
	// false
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// Subscription-based Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// 无效参数
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The list of tags.
	//
	// example:
	//
	// MULTI_BINDED
	Tag []*CreateNatGatewayShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the NAT gateway belongs.
	//
	// When you create a NAT gateway, you must specify the vSwitch to which the NAT gateway belongs. The system allocates an available private IP address from the vSwitch to the NAT gateway.
	//
	// - To create a NAT gateway in an existing vSwitch, make sure that the zone to which the vSwitch belongs supports NAT gateways and the vSwitch has active IP addresses.
	//
	// - If you have not created a vSwitch, create a vSwitch in a zone that supports NAT gateways first, and then specify the vSwitch for the NAT gateway.
	//
	// > You can call the [ListEnhanhcedNatGatewayAvailableZones](https://help.aliyun.com/document_detail/182292.html) operation to query the zones that support NAT gateways, and call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query the available IP address count in a vSwitch.
	//
	// example:
	//
	// vsw-bp1e3se98n9fq8hle****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which you want to create the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1di7uewzmtvfuq8****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNatGatewayShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayShrinkRequest) GetAccessModeShrink() *string {
	return s.AccessModeShrink
}

func (s *CreateNatGatewayShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateNatGatewayShrinkRequest) GetAvailabilityMode() *string {
	return s.AvailabilityMode
}

func (s *CreateNatGatewayShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNatGatewayShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNatGatewayShrinkRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreateNatGatewayShrinkRequest) GetEipBindMode() *string {
	return s.EipBindMode
}

func (s *CreateNatGatewayShrinkRequest) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *CreateNatGatewayShrinkRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateNatGatewayShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateNatGatewayShrinkRequest) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *CreateNatGatewayShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateNatGatewayShrinkRequest) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateNatGatewayShrinkRequest) GetNatType() *string {
	return s.NatType
}

func (s *CreateNatGatewayShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateNatGatewayShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNatGatewayShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNatGatewayShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateNatGatewayShrinkRequest) GetPrivateLinkEnabled() *bool {
	return s.PrivateLinkEnabled
}

func (s *CreateNatGatewayShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNatGatewayShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNatGatewayShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNatGatewayShrinkRequest) GetSecurityProtectionEnabled() *bool {
	return s.SecurityProtectionEnabled
}

func (s *CreateNatGatewayShrinkRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateNatGatewayShrinkRequest) GetTag() []*CreateNatGatewayShrinkRequestTag {
	return s.Tag
}

func (s *CreateNatGatewayShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateNatGatewayShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNatGatewayShrinkRequest) SetAccessModeShrink(v string) *CreateNatGatewayShrinkRequest {
	s.AccessModeShrink = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetAutoPay(v bool) *CreateNatGatewayShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetAvailabilityMode(v string) *CreateNatGatewayShrinkRequest {
	s.AvailabilityMode = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetClientToken(v string) *CreateNatGatewayShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetDescription(v string) *CreateNatGatewayShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetDuration(v string) *CreateNatGatewayShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetEipBindMode(v string) *CreateNatGatewayShrinkRequest {
	s.EipBindMode = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetIcmpReplyEnabled(v bool) *CreateNatGatewayShrinkRequest {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetInstanceChargeType(v string) *CreateNatGatewayShrinkRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetInternetChargeType(v string) *CreateNatGatewayShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetIpv4Prefix(v string) *CreateNatGatewayShrinkRequest {
	s.Ipv4Prefix = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetName(v string) *CreateNatGatewayShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetNatIp(v string) *CreateNatGatewayShrinkRequest {
	s.NatIp = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetNatType(v string) *CreateNatGatewayShrinkRequest {
	s.NatType = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetNetworkType(v string) *CreateNatGatewayShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetOwnerAccount(v string) *CreateNatGatewayShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetOwnerId(v int64) *CreateNatGatewayShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetPricingCycle(v string) *CreateNatGatewayShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetPrivateLinkEnabled(v bool) *CreateNatGatewayShrinkRequest {
	s.PrivateLinkEnabled = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetRegionId(v string) *CreateNatGatewayShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetResourceOwnerAccount(v string) *CreateNatGatewayShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetResourceOwnerId(v int64) *CreateNatGatewayShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetSecurityProtectionEnabled(v bool) *CreateNatGatewayShrinkRequest {
	s.SecurityProtectionEnabled = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetSpec(v string) *CreateNatGatewayShrinkRequest {
	s.Spec = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetTag(v []*CreateNatGatewayShrinkRequestTag) *CreateNatGatewayShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetVSwitchId(v string) *CreateNatGatewayShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) SetVpcId(v string) *CreateNatGatewayShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNatGatewayShrinkRequest) Validate() error {
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

type CreateNatGatewayShrinkRequestTag struct {
	// The tag key. Specify this parameter in the Tag.N.Key format. N is the index of the tag, with valid values from 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Specify this parameter in the Tag.N.Value format. N is the index of the tag, with valid values from 1 to 20. The tag value cannot be an empty string. The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNatGatewayShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNatGatewayShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNatGatewayShrinkRequestTag) SetKey(v string) *CreateNatGatewayShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNatGatewayShrinkRequestTag) SetValue(v string) *CreateNatGatewayShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNatGatewayShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
