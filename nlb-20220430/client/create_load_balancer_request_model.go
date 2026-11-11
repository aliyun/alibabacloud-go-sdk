// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *CreateLoadBalancerRequest
	GetAddressIpVersion() *string
	SetAddressType(v string) *CreateLoadBalancerRequest
	GetAddressType() *string
	SetBandwidthPackageId(v string) *CreateLoadBalancerRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *CreateLoadBalancerRequest
	GetClientToken() *string
	SetDeletionProtectionConfig(v *CreateLoadBalancerRequestDeletionProtectionConfig) *CreateLoadBalancerRequest
	GetDeletionProtectionConfig() *CreateLoadBalancerRequestDeletionProtectionConfig
	SetDryRun(v bool) *CreateLoadBalancerRequest
	GetDryRun() *bool
	SetLoadBalancerBillingConfig(v *CreateLoadBalancerRequestLoadBalancerBillingConfig) *CreateLoadBalancerRequest
	GetLoadBalancerBillingConfig() *CreateLoadBalancerRequestLoadBalancerBillingConfig
	SetLoadBalancerName(v string) *CreateLoadBalancerRequest
	GetLoadBalancerName() *string
	SetLoadBalancerType(v string) *CreateLoadBalancerRequest
	GetLoadBalancerType() *string
	SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest
	GetModificationProtectionConfig() *CreateLoadBalancerRequestModificationProtectionConfig
	SetRegionId(v string) *CreateLoadBalancerRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLoadBalancerRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateLoadBalancerRequestTag) *CreateLoadBalancerRequest
	GetTag() []*CreateLoadBalancerRequestTag
	SetVpcId(v string) *CreateLoadBalancerRequest
	GetVpcId() *string
	SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest
	GetZoneMappings() []*CreateLoadBalancerRequestZoneMappings
}

type CreateLoadBalancerRequest struct {
	// The IP version. Valid values:
	//
	// - **ipv4*	- (default): IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The IPv4 address type of the Network Load Balancing (NLB) instance. Valid values:
	//
	// - **Internet**: The load balancing instance has a public IP address. The DNS domain name is resolved to the public IP address, so it can be accessed over the Internet.
	//
	// - **Intranet**: The load balancing instance has only a private IP address. The DNS domain name is resolved to the private IP address, so it can be accessed only over the internal network of the VPC where the load balancing instance is deployed.
	//
	// > To enable IPv6 access over the Internet for a dual-stack Network Load Balancing (NLB) instance, invoke the [EnableLoadBalancerIpv6Internet](https://help.aliyun.com/document_detail/445878.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the Internet Shared Bandwidth instance that is associated with the Internet-facing instance.
	//
	// example:
	//
	// cbwp-bp1vevu8h3ieh****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deletion protection configuration.
	DeletionProtectionConfig *CreateLoadBalancerRequestDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the instance. The system checks whether required parameters are specified, the request format is valid, and service limits are met. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The billing configuration of the Network Load Balancing (NLB) instance.
	LoadBalancerBillingConfig *CreateLoadBalancerRequestLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The name of the Network Load Balancing (NLB) instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// NLB1
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The type of load balancing. Valid value: **network**, which indicates Network Load Balancing (NLB).
	//
	// example:
	//
	// network
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The modification protection configuration.
	ModificationProtectionConfig *CreateLoadBalancerRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The region ID of the Network Load Balancing (NLB) instance.
	//
	// You can invoke the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	//
	// if can be null:
	// true
	Tag []*CreateLoadBalancerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) where the Network Load Balancing (NLB) instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone and vSwitch mappings. You can add up to 10 zones. If the current region supports two or more zones, you must add at least two zones.
	//
	// This parameter is required.
	ZoneMappings []*CreateLoadBalancerRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *CreateLoadBalancerRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateLoadBalancerRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLoadBalancerRequest) GetDeletionProtectionConfig() *CreateLoadBalancerRequestDeletionProtectionConfig {
	return s.DeletionProtectionConfig
}

func (s *CreateLoadBalancerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerBillingConfig() *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	return s.LoadBalancerBillingConfig
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *CreateLoadBalancerRequest) GetModificationProtectionConfig() *CreateLoadBalancerRequestModificationProtectionConfig {
	return s.ModificationProtectionConfig
}

func (s *CreateLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLoadBalancerRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLoadBalancerRequest) GetTag() []*CreateLoadBalancerRequestTag {
	return s.Tag
}

func (s *CreateLoadBalancerRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLoadBalancerRequest) GetZoneMappings() []*CreateLoadBalancerRequestZoneMappings {
	return s.ZoneMappings
}

func (s *CreateLoadBalancerRequest) SetAddressIpVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetBandwidthPackageId(v string) *CreateLoadBalancerRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeletionProtectionConfig(v *CreateLoadBalancerRequestDeletionProtectionConfig) *CreateLoadBalancerRequest {
	s.DeletionProtectionConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetDryRun(v bool) *CreateLoadBalancerRequest {
	s.DryRun = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerBillingConfig(v *CreateLoadBalancerRequestLoadBalancerBillingConfig) *CreateLoadBalancerRequest {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerType(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest {
	s.ModificationProtectionConfig = v
	return s
}

func (s *CreateLoadBalancerRequest) SetRegionId(v string) *CreateLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetResourceGroupId(v string) *CreateLoadBalancerRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetTag(v []*CreateLoadBalancerRequestTag) *CreateLoadBalancerRequest {
	s.Tag = v
	return s
}

func (s *CreateLoadBalancerRequest) SetVpcId(v string) *CreateLoadBalancerRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetZoneMappings(v []*CreateLoadBalancerRequestZoneMappings) *CreateLoadBalancerRequest {
	s.ZoneMappings = v
	return s
}

func (s *CreateLoadBalancerRequest) Validate() error {
	if s.DeletionProtectionConfig != nil {
		if err := s.DeletionProtectionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LoadBalancerBillingConfig != nil {
		if err := s.LoadBalancerBillingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ModificationProtectionConfig != nil {
		if err := s.ModificationProtectionConfig.Validate(); err != nil {
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
	if s.ZoneMappings != nil {
		for _, item := range s.ZoneMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLoadBalancerRequestDeletionProtectionConfig struct {
	// The deletion protection status. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false*	- (default): disabled.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The reason for changing the deletion protection status. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// Instance is running
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s CreateLoadBalancerRequestDeletionProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) SetEnabled(v bool) *CreateLoadBalancerRequestDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) SetReason(v string) *CreateLoadBalancerRequestDeletionProtectionConfig {
	s.Reason = &v
	return s
}

func (s *CreateLoadBalancerRequestDeletionProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestLoadBalancerBillingConfig struct {
	// The billing method of the Network Load Balancing (NLB) instance.
	//
	// Valid value: **PostPay**, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) GetPayType() *string {
	return s.PayType
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetPayType(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestModificationProtectionConfig struct {
	// The reason for enabling modification protection. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// > This parameter is valid and takes effect only when **Status*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// Business protection period
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The modification protection status of the Network Load Balancing (NLB). Valid values:
	//
	// - **NonProtection**: modification protection is disabled. You cannot specify the **Reason*	- parameter. If you specify the **Reason*	- parameter, the configuration is cleared.
	//
	// - **ConsoleProtection**: console modification protection is enabled. You can specify the **Reason*	- parameter.
	//
	// > If you set the value to **ConsoleProtection**, modification protection is enabled. You cannot modify the instance configurations in the load balancing console, but you can invoke API operations to modify the instance configurations.
	//
	// example:
	//
	// ConsoleProtection
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) GetStatus() *string {
	return s.Status
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetReason(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) SetStatus(v string) *CreateLoadBalancerRequestModificationProtectionConfig {
	s.Status = &v
	return s
}

func (s *CreateLoadBalancerRequestModificationProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestTag struct {
	// The tag key. The key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. The key can contain letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`. The value can contain letters, digits, underscores (_), periods (.), colons (:), forward slashes (/), equal signs (=), plus signs (+), hyphens (-), and at signs (@).
	//
	// You can add up to 20 tags in a single call.
	//
	// example:
	//
	// product
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadBalancerRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLoadBalancerRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLoadBalancerRequestTag) SetKey(v string) *CreateLoadBalancerRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLoadBalancerRequestTag) SetValue(v string) *CreateLoadBalancerRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLoadBalancerRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestZoneMappings struct {
	// The EIP instance ID associated with the Internet-facing instance. Each zone corresponds to one EIP instance ID. You can add up to 10 zones. If the current region supports two or more zones, you must add at least two zones.
	//
	// example:
	//
	// eip-bp1aedxso6u80u0qf****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The list of IPv4 local addresses. These are the addresses used by the NLB to interact with backend services.
	//
	// The number of addresses must be an even number, with a minimum of 2 and a maximum of 8.
	Ipv4LocalAddresses []*string `json:"Ipv4LocalAddresses,omitempty" xml:"Ipv4LocalAddresses,omitempty" type:"Repeated"`
	// The IPv6 VIP. This is the IPv6 address used by the NLB to provide services.
	//
	// example:
	//
	// 2408:400a:d5:3080:b409:840a:ca:e8e5
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The list of IPv6 local addresses. These are the addresses used by the NLB to interact with backend services.
	//
	// The number of addresses must be an even number, with a minimum of 2 and a maximum of 8.
	Ipv6LocalAddresses []*string `json:"Ipv6LocalAddresses,omitempty" xml:"Ipv6LocalAddresses,omitempty" type:"Repeated"`
	// The private IPv4 virtual IP address (VIP). This is the private IPv4 address used by the NLB to provide services.
	//
	// example:
	//
	// 192.168.10.1
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	// The vSwitch that corresponds to the zone. You can use only one vSwitch and one subnet in each zone. You can add up to 10 zones.
	//
	// If the current region supports two or more zones, you must add at least two zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-sersdf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone where the Network Load Balancing (NLB) instance is deployed. You can add up to 10 zones. If the current region supports two or more zones, you must add at least two zones.
	//
	// You can invoke the [DescribeZones](https://help.aliyun.com/document_detail/443890.html) operation to query zone IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLoadBalancerRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerRequestZoneMappings) GetAllocationId() *string {
	return s.AllocationId
}

func (s *CreateLoadBalancerRequestZoneMappings) GetIpv4LocalAddresses() []*string {
	return s.Ipv4LocalAddresses
}

func (s *CreateLoadBalancerRequestZoneMappings) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *CreateLoadBalancerRequestZoneMappings) GetIpv6LocalAddresses() []*string {
	return s.Ipv6LocalAddresses
}

func (s *CreateLoadBalancerRequestZoneMappings) GetPrivateIPv4Address() *string {
	return s.PrivateIPv4Address
}

func (s *CreateLoadBalancerRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLoadBalancerRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLoadBalancerRequestZoneMappings) SetAllocationId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.AllocationId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetIpv4LocalAddresses(v []*string) *CreateLoadBalancerRequestZoneMappings {
	s.Ipv4LocalAddresses = v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetIpv6Address(v string) *CreateLoadBalancerRequestZoneMappings {
	s.Ipv6Address = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetIpv6LocalAddresses(v []*string) *CreateLoadBalancerRequestZoneMappings {
	s.Ipv6LocalAddresses = v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetPrivateIPv4Address(v string) *CreateLoadBalancerRequestZoneMappings {
	s.PrivateIPv4Address = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetVSwitchId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetZoneId(v string) *CreateLoadBalancerRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
