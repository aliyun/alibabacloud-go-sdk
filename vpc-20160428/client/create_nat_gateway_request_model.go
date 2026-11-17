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
	// The access mode for reverse access to the VPC NAT gateway.
	//
	// example:
	//
	// MULTI_BINDED
	AccessMode *CreateNatGatewayRequestAccessMode `json:"AccessMode,omitempty" xml:"AccessMode,omitempty" type:"Struct"`
	// Subscription Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// 	Notice:  This parameter is deprecated.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The disaster recovery type of the NAT gateway. Valid values:
	//
	// - **CrossAZ**: cross-zone disaster recovery.
	//
	// - **SingleAZ**: single-zone disaster recovery.
	//
	// example:
	//
	// CrossAZ
	AvailabilityMode *string `json:"AvailabilityMode,omitempty" xml:"AvailabilityMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the NAT gateway.
	//
	// The description can be left empty or be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testnat
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Subscription Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// Invalid parameter
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The EIP binding mode of the NAT gateway. Valid values:
	//
	// - **MULTI_BINDED*	- (default): multi-EIP NIC-visible mode.
	//
	// - **NAT**: NAT mode, which is compatible with IPv4 gateways.
	//
	//   > When the EIP binding mode of the NAT gateway is set to NAT mode, the EIP occupies a private IP address of the vSwitch to which the NAT gateway belongs. Make sure that the vSwitch has sufficient private IP addresses. If the vSwitch to which the NAT gateway belongs does not have available idle private IP addresses, you cannot bind new EIPs. When the EIP binding mode of the NAT gateway is set to NAT mode, the NAT gateway supports a maximum of 50 EIPs.
	//
	// example:
	//
	// MULTI_BINDED
	EipBindMode *string `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	// Specifies whether to enable ICMP echo reply. Valid values:
	//
	// - **true*	- (default): enables ICMP echo reply.
	//
	// - **false**: disables ICMP echo reply.
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
	// For more information, see [Internet NAT gateway billing](https://help.aliyun.com/document_detail/48126.html) and [VPC NAT gateway billing](https://help.aliyun.com/document_detail/270913.html).
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The metering method of the NAT gateway. Set the value to **PayByLcu**, which specifies pay-by-LCU.
	//
	// example:
	//
	// PayByLcu
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The IP prefix CIDR block that is used to create NAT IP addresses in batches. Use an unallocated reserved CIDR block in the vSwitch to which the NAT gateway belongs.
	//
	// > This parameter takes effect only for VPC NAT gateways.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// The name of the NAT gateway.
	//
	// The name must be 2 to 128 characters in length. It must start with a letter or a Chinese character, and can contain digits, underscores (_), and hyphens (-).
	//
	// If you do not specify this parameter, the system generates a default name for the NAT gateway.
	//
	// example:
	//
	// fortest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private IP address that is allocated to the NAT gateway. Use an unallocated IP address in the vSwitch CIDR block of the NAT gateway. If this parameter is left empty, an IP address is randomly allocated.
	//
	// > This parameter takes effect only for VPC NAT gateways.
	//
	// example:
	//
	// 192.168.0.2
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The type of the NAT gateway. Set the value to **Enhanced**, which specifies an enhanced NAT gateway.
	//
	// example:
	//
	// Enhanced
	NatType *string `json:"NatType,omitempty" xml:"NatType,omitempty"`
	// The type of the NAT gateway that you want to create. Valid values:
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
	// Subscription Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// Invalid parameter
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
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
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
	// 	Notice:  This parameter is deprecated.
	//
	// example:
	//
	// false
	SecurityProtectionEnabled *bool `json:"SecurityProtectionEnabled,omitempty" xml:"SecurityProtectionEnabled,omitempty"`
	// Subscription Internet NAT gateways are no longer available for purchase. This parameter is no longer used.
	//
	// example:
	//
	// Invalid parameter
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The list of tags.
	//
	// example:
	//
	// MULTI_BINDED
	Tag []*CreateNatGatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the NAT gateway belongs.
	//
	// When you create a NAT gateway, you must specify the vSwitch to which the NAT gateway belongs. The system allocates an idle private IP address from the vSwitch to the NAT gateway.
	//
	// - If you want to create a NAT gateway in an existing vSwitch, make sure that the zone to which the vSwitch belongs supports NAT gateways and the vSwitch has available IP addresses.
	//
	// - If you have not created a vSwitch, create a vSwitch in a zone that supports NAT gateways, and then specify the vSwitch to which the NAT gateway belongs.
	//
	// > You can call the [ListEnhanhcedNatGatewayAvailableZones](https://help.aliyun.com/document_detail/182292.html) operation to query the zones that support NAT gateways, and call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query the available IP address count in a vSwitch.
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
	// - **route**: route mode.
	//
	// - **tunnel**: tunnel mode.
	//
	// > When this field is specified, the **PrivateLinkEnabled*	- field must be set to **true**.
	//
	// example:
	//
	// route
	ModeValue *string `json:"ModeValue,omitempty" xml:"ModeValue,omitempty"`
	// The tunnel mode type. Valid values:
	//
	// - **geneve**: Geneve.
	//
	// > This value is valid only when the access mode is set to tunnel mode.
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
	// The tag key. You can use the Tag.N.Key format when you call API operations. Valid values of N: 1 to 20. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can use the Tag.N.Value format when you call API operations. Valid values of N: 1 to 20. If you specify this parameter, it cannot be an empty string. The tag value can be up to 128 characters in length. It cannot start with aliyun or acs:, and cannot contain http:// or https://.
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
