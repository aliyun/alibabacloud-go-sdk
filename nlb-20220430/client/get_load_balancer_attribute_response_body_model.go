// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody
	GetAddressIpVersion() *string
	SetAddressType(v string) *GetLoadBalancerAttributeResponseBody
	GetAddressType() *string
	SetBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBody
	GetBandwidthPackageId() *string
	SetCps(v int32) *GetLoadBalancerAttributeResponseBody
	GetCps() *int32
	SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody
	GetCreateTime() *string
	SetCrossZoneEnabled(v bool) *GetLoadBalancerAttributeResponseBody
	GetCrossZoneEnabled() *bool
	SetDNSName(v string) *GetLoadBalancerAttributeResponseBody
	GetDNSName() *string
	SetDeletionProtectionConfig(v *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) *GetLoadBalancerAttributeResponseBody
	GetDeletionProtectionConfig() *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig
	SetIpv6AddressType(v string) *GetLoadBalancerAttributeResponseBody
	GetIpv6AddressType() *string
	SetLoadBalancerBillingConfig(v *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerBillingConfig() *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig
	SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerBusinessStatus() *string
	SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerName() *string
	SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerStatus() *string
	SetLoadBalancerType(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerType() *string
	SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody
	GetModificationProtectionConfig() *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig
	SetOperationLocks(v []*GetLoadBalancerAttributeResponseBodyOperationLocks) *GetLoadBalancerAttributeResponseBody
	GetOperationLocks() []*GetLoadBalancerAttributeResponseBodyOperationLocks
	SetRegionId(v string) *GetLoadBalancerAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetLoadBalancerAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody
	GetResourceGroupId() *string
	SetSecurityGroupIds(v []*string) *GetLoadBalancerAttributeResponseBody
	GetSecurityGroupIds() []*string
	SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody
	GetTags() []*GetLoadBalancerAttributeResponseBodyTags
	SetVpcId(v string) *GetLoadBalancerAttributeResponseBody
	GetVpcId() *string
	SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody
	GetZoneMappings() []*GetLoadBalancerAttributeResponseBodyZoneMappings
}

type GetLoadBalancerAttributeResponseBody struct {
	// The protocol version. Valid values:
	//
	// - **ipv4**: IPv4.
	//
	// - **DualStack**: dual stack.
	//
	// example:
	//
	// ipv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The IPv4 network address type of the NLB instance. Valid values:
	//
	// - **Internet**: The load balancing instance has a public IP address. The DNS domain name is resolved to the public IP address. Therefore, the instance can be accessed over the Internet.
	//
	// - **Intranet**: The load balancing instance has only a private IP address. The DNS domain name is resolved to the private IP address. Therefore, the instance can be accessed only within the internal network of the VPC where the load balancing instance is deployed.
	//
	// example:
	//
	// Internet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the Internet Shared Bandwidth instance associated with the Internet-facing NLB instance.
	//
	// example:
	//
	// cbwp-bp1vevu8h3ieh****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The maximum number of new connections per second that the NLB instance can process per zone (VIP). Valid values: **0*	- to **1000000**.
	//
	// **0*	- indicates that no limit is imposed.
	//
	// example:
	//
	// 100
	Cps *int32 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The time when the resource was created. The time is displayed in UTC in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether cross-zone is enabled for the NLB instance. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	CrossZoneEnabled *bool `json:"CrossZoneEnabled,omitempty" xml:"CrossZoneEnabled,omitempty"`
	// The domain name of the NLB instance.
	//
	// example:
	//
	// nlb-wb7r6dlwetvt5j****.cn-hangzhou.nlb.aliyuncs.com
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The deletion protection configuration.
	DeletionProtectionConfig *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// The IPv6 address type of the NLB instance. Valid values:
	//
	// - **Internet**: The load balancing instance has a public IP address. The DNS domain name is resolved to the public IP address. Therefore, the instance can be accessed over the Internet.
	//
	// - **Intranet**: The load balancing instance has only a private IP address. The DNS domain name is resolved to the private IP address. Therefore, the instance can be accessed only within the internal network of the VPC where the load balancing instance is deployed.
	//
	// example:
	//
	// Internet
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The billing configuration of the load balancing instance.
	LoadBalancerBillingConfig *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The business status of the NLB instance. Valid values:
	//
	// - **Abnormal**: The NLB instance is in an abnormal state.
	//
	// - **Normal**: The NLB instance is in a normal state.
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The NLB instance ID.
	//
	// example:
	//
	// nlb-83ckzc8d4xlp8o****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the NLB instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter or a Chinese character.
	//
	// example:
	//
	// NLB1
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The NLB instance status. Valid values:
	//
	// - **Inactive**: The NLB instance is stopped. The listeners of the instance no longer forward traffic.
	//
	// - **Active**: The NLB instance is running.
	//
	// - **Provisioning**: The NLB instance is being created.
	//
	// - **Configuring**: The NLB instance is being reconfigured (Upgrade/Downgrade).
	//
	// - **CreateFailed**: The NLB instance failed to be created. In this case, no fees are incurred, and the instance can only be deleted.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The load balancing type. Only **network*	- is returned, which indicates Network Load Balancer (NLB).
	//
	// example:
	//
	// network
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The modification protection configuration.
	ModificationProtectionConfig *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The lock state list of the NLB instance. This parameter is returned only when `LoadBalancerBussinessStatus` is **Abnormal**.
	OperationLocks []*GetLoadBalancerAttributeResponseBodyOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Repeated"`
	// The region ID of the NLB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security group IDs associated with the NLB instance.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The tag list.
	Tags []*GetLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The VPC ID of the NLB instance.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. A maximum of 10 zones are returned. If the current region supports two or more zones, at least two zones are returned.
	ZoneMappings []*GetLoadBalancerAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *GetLoadBalancerAttributeResponseBody) GetAddressType() *string {
	return s.AddressType
}

func (s *GetLoadBalancerAttributeResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *GetLoadBalancerAttributeResponseBody) GetCps() *int32 {
	return s.Cps
}

func (s *GetLoadBalancerAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLoadBalancerAttributeResponseBody) GetCrossZoneEnabled() *bool {
	return s.CrossZoneEnabled
}

func (s *GetLoadBalancerAttributeResponseBody) GetDNSName() *string {
	return s.DNSName
}

func (s *GetLoadBalancerAttributeResponseBody) GetDeletionProtectionConfig() *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	return s.DeletionProtectionConfig
}

func (s *GetLoadBalancerAttributeResponseBody) GetIpv6AddressType() *string {
	return s.Ipv6AddressType
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerBillingConfig() *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig {
	return s.LoadBalancerBillingConfig
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerBusinessStatus() *string {
	return s.LoadBalancerBusinessStatus
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *GetLoadBalancerAttributeResponseBody) GetModificationProtectionConfig() *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	return s.ModificationProtectionConfig
}

func (s *GetLoadBalancerAttributeResponseBody) GetOperationLocks() []*GetLoadBalancerAttributeResponseBodyOperationLocks {
	return s.OperationLocks
}

func (s *GetLoadBalancerAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoadBalancerAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLoadBalancerAttributeResponseBody) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *GetLoadBalancerAttributeResponseBody) GetTags() []*GetLoadBalancerAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetLoadBalancerAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLoadBalancerAttributeResponseBody) GetZoneMappings() []*GetLoadBalancerAttributeResponseBodyZoneMappings {
	return s.ZoneMappings
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCps(v int32) *GetLoadBalancerAttributeResponseBody {
	s.Cps = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCrossZoneEnabled(v bool) *GetLoadBalancerAttributeResponseBody {
	s.CrossZoneEnabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDNSName(v string) *GetLoadBalancerAttributeResponseBody {
	s.DNSName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetDeletionProtectionConfig(v *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.DeletionProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetIpv6AddressType(v string) *GetLoadBalancerAttributeResponseBody {
	s.Ipv6AddressType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBillingConfig(v *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerType(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.ModificationProtectionConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetOperationLocks(v []*GetLoadBalancerAttributeResponseBodyOperationLocks) *GetLoadBalancerAttributeResponseBody {
	s.OperationLocks = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRegionId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetSecurityGroupIds(v []*string) *GetLoadBalancerAttributeResponseBody {
	s.SecurityGroupIds = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetVpcId(v string) *GetLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) Validate() error {
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
	if s.OperationLocks != nil {
		for _, item := range s.OperationLocks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
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

type GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig struct {
	// The deletion protection status. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when deletion protection was enabled. The time is displayed in UTC in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2022-11-02T02:49:05Z
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	// The reason for changing the deletion protection status. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The reason must start with a letter or a Chinese character.
	//
	// example:
	//
	// Instance running.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GetEnabledTime() *string {
	return s.EnabledTime
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabled(v bool) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetReason(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Reason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig struct {
	// The billing type of the NLB instance.
	//
	// Only **PostPay*	- is returned, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) GetPayType() *string {
	return s.PayType
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) SetPayType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyModificationProtectionConfig struct {
	// The time when modification protection was enabled. The time is displayed in UTC in the `yyyy-MM-ddTHH:mm:ssZ` format.
	//
	// example:
	//
	// 2022-12-02T02:49:05Z
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
	// The reason for enabling modification protection. The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The reason must start with a letter or a Chinese character.
	//
	// This parameter is valid and required only when **Status*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// Business protection period.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The modification protection status of the Network Load Balancer (NLB) instance. Valid values:
	//
	// - **NonProtection**: Modification protection is disabled. You cannot specify the **Reason*	- parameter. If the **Reason*	- parameter is specified, the value is cleared.
	//
	// - **ConsoleProtection**: Console modification protection is enabled. You can specify the **Reason*	- parameter.
	//
	// > If the value is **ConsoleProtection**, you cannot modify the instance configuration in the load balancing console, but you can modify the instance configuration by invoking API operations.
	//
	// example:
	//
	// ConsoleProtection
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GetEnabledTime() *string {
	return s.EnabledTime
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GetStatus() *string {
	return s.Status
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetReason(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) SetStatus(v string) *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	s.Status = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyOperationLocks struct {
	// The reason for the lock.
	//
	// example:
	//
	// Business exception.
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The type of the lock. Valid values:
	//
	// - **SecurityLocked**: The instance is locked due to security reasons.
	//
	// - **RelatedResourceLocked**: The instance is locked due to associate lock.
	//
	// - **FinancialLocked**: The instance is locked due to overdue payment.
	//
	// - **ResidualLocked**: The instance is locked due to residual resources.
	//
	// example:
	//
	// SecurityLocked
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyOperationLocks) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) GetLockReason() *string {
	return s.LockReason
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) GetLockType() *string {
	return s.LockType
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) SetLockReason(v string) *GetLoadBalancerAttributeResponseBodyOperationLocks {
	s.LockReason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) SetLockType(v string) *GetLoadBalancerAttributeResponseBodyOperationLocks {
	s.LockType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyOperationLocks) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyTags struct {
	// The tag key of the instance.
	//
	// example:
	//
	// KeyTest
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the instance.
	//
	// example:
	//
	// ValueTest
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetLoadBalancerAttributeResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetTagKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetTagValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	// The Network Load Balancing (NLB) instance address information.
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// The zone status. Valid values:
	//
	// - **Active**: enabled.
	//
	// - **Stopped**: stopped. This status can be set only through the CADT disaster recovery management service invocation.
	//
	// - **Shifted**: DNS manually removed. You can call the [CancelShiftLoadBalancerZones](https://help.aliyun.com/document_detail/2412000.html) operation to restore the DNS of the zone.
	//
	// - **AutoShifted**: DNS automatically removed. When a zone encounters an exception, the system automatically removes the IP address of the abnormal zone to ensure that you can access services normally through the NLB domain name. The IP address is automatically restored when the zone recovers. No manual operation is required.
	//
	// - **Starting**: starting. This status can be set only through the CADT disaster recovery management service invocation.
	//
	// - **Stopping**: stopping. This status can be set only through the CADT disaster recovery management service invocation.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch in the zone. Each zone uses one vSwitch and one subnet by default.
	//
	// example:
	//
	// vsw-bp1rmcrwg3erh1fh8****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID. You can call the [DescribeZones](https://help.aliyun.com/document_detail/443890.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetLoadBalancerAddresses() []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	return s.LoadBalancerAddresses
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetStatus() *string {
	return s.Status
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetLoadBalancerAddresses(v []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.Status = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) Validate() error {
	if s.LoadBalancerAddresses != nil {
		for _, item := range s.LoadBalancerAddresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	// The Elastic IP Address (EIP) instance ID.
	//
	// example:
	//
	// eip-bp1aedxso6u80u0qf****
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-bp12f1xhs5yal61a****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The list of IPv4 local addresses. These are the addresses used for communication between the NLB instance and backend servers.
	Ipv4LocalAddresses []*string `json:"Ipv4LocalAddresses,omitempty" xml:"Ipv4LocalAddresses,omitempty" type:"Repeated"`
	// The IPv6 VIP address. The IPv6 address that the NLB instance uses to provide services.
	//
	// example:
	//
	// 2001:db8:1:1:1:1:1:1
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The list of IPv6 local addresses. These are the addresses used for communication between the NLB instance and backend servers.
	Ipv6LocalAddresses []*string `json:"Ipv6LocalAddresses,omitempty" xml:"Ipv6LocalAddresses,omitempty" type:"Repeated"`
	// The private IPv4 VIP address. The private IPv4 address that the NLB instance uses to provide services.
	//
	// example:
	//
	// 192.168.3.32
	PrivateIPv4Address *string `json:"PrivateIPv4Address,omitempty" xml:"PrivateIPv4Address,omitempty"`
	// The health check status of the private IPv4 address of the NLB instance. Valid values:
	//
	// - **Healthy**: healthy.
	//
	// - **Unhealthy**: unhealthy.
	//
	// > This parameter is returned only when the zone status (**Status**) is Active.
	//
	// example:
	//
	// Healthy
	PrivateIPv4HcStatus *string `json:"PrivateIPv4HcStatus,omitempty" xml:"PrivateIPv4HcStatus,omitempty"`
	// The health check status of the IPv6 address of the NLB instance. Valid values:
	//
	// - **Healthy**: healthy.
	//
	// - **Unhealthy**: unhealthy.
	//
	// > This parameter is returned only when the zone status (**Status**) is Active.
	//
	// example:
	//
	// Healthy
	PrivateIPv6HcStatus *string `json:"PrivateIPv6HcStatus,omitempty" xml:"PrivateIPv6HcStatus,omitempty"`
	// The public IPv4 address of the NLB instance.
	//
	// example:
	//
	// 120.XX.XX.69
	PublicIPv4Address *string `json:"PublicIPv4Address,omitempty" xml:"PublicIPv4Address,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetAllocationId() *string {
	return s.AllocationId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetEniId() *string {
	return s.EniId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv4LocalAddresses() []*string {
	return s.Ipv4LocalAddresses
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv6LocalAddresses() []*string {
	return s.Ipv6LocalAddresses
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetPrivateIPv4Address() *string {
	return s.PrivateIPv4Address
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetPrivateIPv4HcStatus() *string {
	return s.PrivateIPv4HcStatus
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetPrivateIPv6HcStatus() *string {
	return s.PrivateIPv6HcStatus
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetPublicIPv4Address() *string {
	return s.PublicIPv4Address
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAllocationId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.AllocationId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetEniId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv4LocalAddresses(v []*string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv4LocalAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv6Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv6Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv6LocalAddresses(v []*string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv6LocalAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIPv4Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIPv4HcStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv4HcStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIPv6HcStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIPv6HcStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPublicIPv4Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PublicIPv4Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) Validate() error {
	return dara.Validate(s)
}
