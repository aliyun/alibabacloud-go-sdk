// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressAllocatedMode(v string) *CreateLoadBalancerRequest
	GetAddressAllocatedMode() *string
	SetAddressIpVersion(v string) *CreateLoadBalancerRequest
	GetAddressIpVersion() *string
	SetAddressType(v string) *CreateLoadBalancerRequest
	GetAddressType() *string
	SetClientToken(v string) *CreateLoadBalancerRequest
	GetClientToken() *string
	SetDeletionProtectionEnabled(v bool) *CreateLoadBalancerRequest
	GetDeletionProtectionEnabled() *bool
	SetDryRun(v bool) *CreateLoadBalancerRequest
	GetDryRun() *bool
	SetLoadBalancerBillingConfig(v *CreateLoadBalancerRequestLoadBalancerBillingConfig) *CreateLoadBalancerRequest
	GetLoadBalancerBillingConfig() *CreateLoadBalancerRequestLoadBalancerBillingConfig
	SetLoadBalancerEdition(v string) *CreateLoadBalancerRequest
	GetLoadBalancerEdition() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerRequest
	GetLoadBalancerName() *string
	SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest
	GetModificationProtectionConfig() *CreateLoadBalancerRequestModificationProtectionConfig
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
	// The mode in which IP addresses are allocated to the ALB instance. Valid values:
	//
	// 	- **Fixed*	- (default): a fixed IP address is assigned to the ALB instance in each zone.
	//
	// 	- **Dynamic**: IP addresses are dynamically allocated to the ALB instance in each zone.
	//
	// >  Starting from 00:00:00 on February 25, 2025 (UTC+8), when you call this operation to create an ALB instance, the instance is automatically the [upgraded version](https://help.aliyun.com/document_detail/2864070.html) regardless of the mode you specify. Upgraded ALB instances no longer differentiate between IP modes. Instead, they globally auto-scale IP addresses for providing load balancing services. The ALB instances you created before this date and time are not affected.
	//
	// example:
	//
	// Dynamic
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// The protocol version. Valid values:
	//
	// 	- **IPv4:*	- IPv4.
	//
	// 	- **DualStack:*	- dual stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The type of the address of the ALB instance. Valid values:
	//
	// 	- **Internet:*	- The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. In this case, the ALB instance can be accessed over the Internet.
	//
	// 	- **Intranet:*	- The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the VPC in which the ALB instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses the value of **RequestId*	- as the value of **ClientToken**. The value of the **RequestId*	- parameter may be different for each API request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable deletion protection. Default value: false. Valid values:
	//
	// 	- **true:*	- enables deletion protection.
	//
	// 	- **false:*	- disables deletion protection.
	//
	// example:
	//
	// false
	DeletionProtectionEnabled *bool `json:"DeletionProtectionEnabled,omitempty" xml:"DeletionProtectionEnabled,omitempty"`
	// Specifies whether to perform a dry run. Default value: false. Valid values:
	//
	// 	- **true:*	- performs a dry run. The system checks the required parameters, request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false:*	- performs a dry run and sends the request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The billing method of the ALB instance.
	//
	// This parameter is required.
	LoadBalancerBillingConfig *CreateLoadBalancerRequestLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The edition of the ALB instance. The features and billing rules vary based on the edition of the ALB instance. Valid values:
	//
	// 	- **Basic:*	- basic.
	//
	// 	- **Standard:*	- standard.
	//
	// 	- **StandardWithWaf:*	- WAF-enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Standard
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The name of the ALB instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// alb1
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The configuration read-only mode settings.
	ModificationProtectionConfig *CreateLoadBalancerRequestModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateLoadBalancerRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which you want to create the ALB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones an vSwitches. You can specify at most 10 zones. If the selected region supports two or more zones, select at least two zones to ensure the high availability of your service.
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

func (s *CreateLoadBalancerRequest) GetAddressAllocatedMode() *string {
	return s.AddressAllocatedMode
}

func (s *CreateLoadBalancerRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *CreateLoadBalancerRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateLoadBalancerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLoadBalancerRequest) GetDeletionProtectionEnabled() *bool {
	return s.DeletionProtectionEnabled
}

func (s *CreateLoadBalancerRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerBillingConfig() *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	return s.LoadBalancerBillingConfig
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerEdition() *string {
	return s.LoadBalancerEdition
}

func (s *CreateLoadBalancerRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerRequest) GetModificationProtectionConfig() *CreateLoadBalancerRequestModificationProtectionConfig {
	return s.ModificationProtectionConfig
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

func (s *CreateLoadBalancerRequest) SetAddressAllocatedMode(v string) *CreateLoadBalancerRequest {
	s.AddressAllocatedMode = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressIpVersion(v string) *CreateLoadBalancerRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetAddressType(v string) *CreateLoadBalancerRequest {
	s.AddressType = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetClientToken(v string) *CreateLoadBalancerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetDeletionProtectionEnabled(v bool) *CreateLoadBalancerRequest {
	s.DeletionProtectionEnabled = &v
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

func (s *CreateLoadBalancerRequest) SetLoadBalancerEdition(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerEdition = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetLoadBalancerName(v string) *CreateLoadBalancerRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerRequest) SetModificationProtectionConfig(v *CreateLoadBalancerRequestModificationProtectionConfig) *CreateLoadBalancerRequest {
	s.ModificationProtectionConfig = v
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

type CreateLoadBalancerRequestLoadBalancerBillingConfig struct {
	// The ID of the Internet Shared Bandwidth instance that is associated with the Internet-facing ALB instance.
	//
	// example:
	//
	// cbwp-bp1vevu8h3ieh****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The billing method of the instance.
	//
	// Set the value to **PostPay**, which specifies the pay-as-you-go billing method.
	//
	// This parameter is required.
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

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) GetPayType() *string {
	return s.PayType
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetBandwidthPackageId(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.BandwidthPackageId = &v
	return s
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) SetPayType(v string) *CreateLoadBalancerRequestLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

func (s *CreateLoadBalancerRequestLoadBalancerBillingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateLoadBalancerRequestModificationProtectionConfig struct {
	// The reason for enabling the configuration read-only mode.
	//
	// The reason must be 2 to 128 characters in length, can contain letters, digits, periods (.), underscores (_), and hyphens (-), and must start with a letter.
	//
	// >  This parameter takes effect only when **Status*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Specifies whether to enable the configuration read-only mode. Valid values:
	//
	// 	- **NonProtection**: Disables the configuration read-only mode. In this case, the value of the **Reason*	- parameter that you specify does not take effect. If you specify **Reason**, the value of the parameter is cleared.
	//
	// 	- **ConsoleProtection**: Enables the configuration read-only mode. In this case, the value of the **Reason*	- parameter that you specify takes effect.****
	//
	// >  If the parameter is set to **ConsoleProtection**, the configuration read-only mode is enabled. You cannot modify the configurations of the ALB instance in the ALB console. However, you can call API operations to modify the configurations of the ALB instance.
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
	// The tag key can be up to 128 characters in length, and cannot start with acs: or aliyun. It cannot contain http:// or https://.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value can be up to 128 characters in length, and cannot start with acs: or aliyun. It cannot contain http:// or https://.
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
	// The ID of the EIP to be associated with the Internet-facing ALB instance.
	//
	// example:
	//
	// eip-bp1aedxso6u80u0qf****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The type of EIP. Valid values:
	//
	// 	- **Common**: an EIP.
	//
	// 	- **Anycast**: an Anycast EIP.
	//
	// >  For more information about the regions in which ALB supports Anycast EIPs, see [Limits](https://help.aliyun.com/document_detail/460727.html).
	//
	// example:
	//
	// Common
	EipType *string `json:"EipType,omitempty" xml:"EipType,omitempty"`
	// The private IPv4 address.
	//
	// example:
	//
	// 192.168.10.1
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone. You can specify at most 10 zones. If the region supports two or more zones, specify at least two zones.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-sersdf****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the cluster. You can specify at most 10 zones. If the region supports two or more zones, specify at least two zones. You can call the [DescribeZones](https://help.aliyun.com/document_detail/36064.html) operation to query the most recent zone list.
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

func (s *CreateLoadBalancerRequestZoneMappings) GetEipType() *string {
	return s.EipType
}

func (s *CreateLoadBalancerRequestZoneMappings) GetIntranetAddress() *string {
	return s.IntranetAddress
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

func (s *CreateLoadBalancerRequestZoneMappings) SetEipType(v string) *CreateLoadBalancerRequestZoneMappings {
	s.EipType = &v
	return s
}

func (s *CreateLoadBalancerRequestZoneMappings) SetIntranetAddress(v string) *CreateLoadBalancerRequestZoneMappings {
	s.IntranetAddress = &v
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
