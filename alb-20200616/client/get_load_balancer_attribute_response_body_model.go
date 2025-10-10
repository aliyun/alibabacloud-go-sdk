// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLogConfig(v *GetLoadBalancerAttributeResponseBodyAccessLogConfig) *GetLoadBalancerAttributeResponseBody
	GetAccessLogConfig() *GetLoadBalancerAttributeResponseBodyAccessLogConfig
	SetAddressAllocatedMode(v string) *GetLoadBalancerAttributeResponseBody
	GetAddressAllocatedMode() *string
	SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody
	GetAddressIpVersion() *string
	SetAddressType(v string) *GetLoadBalancerAttributeResponseBody
	GetAddressType() *string
	SetBandwidthPackageId(v string) *GetLoadBalancerAttributeResponseBody
	GetBandwidthPackageId() *string
	SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody
	GetCreateTime() *string
	SetDNSName(v string) *GetLoadBalancerAttributeResponseBody
	GetDNSName() *string
	SetDeletionProtectionConfig(v *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) *GetLoadBalancerAttributeResponseBody
	GetDeletionProtectionConfig() *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig
	SetIpv6AddressType(v string) *GetLoadBalancerAttributeResponseBody
	GetIpv6AddressType() *string
	SetLoadBalancerBillingConfig(v *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerBillingConfig() *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig
	SetLoadBalancerBussinessStatus(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerBussinessStatus() *string
	SetLoadBalancerEdition(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerEdition() *string
	SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerName() *string
	SetLoadBalancerOperationLocks(v []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerOperationLocks() []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks
	SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerStatus() *string
	SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody
	GetModificationProtectionConfig() *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig
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
	// The configuration of the access log feature.
	AccessLogConfig *GetLoadBalancerAttributeResponseBodyAccessLogConfig `json:"AccessLogConfig,omitempty" xml:"AccessLogConfig,omitempty" type:"Struct"`
	// The mode in which IP addresses are allocated. Valid values:
	//
	// 	- **Fixed**: allocates a static IP address to the ALB instance.
	//
	// 	- **Dynamic**: dynamically allocates an IP address to each zone of the ALB instance.
	//
	// example:
	//
	// Dynamic
	AddressAllocatedMode *string `json:"AddressAllocatedMode,omitempty" xml:"AddressAllocatedMode,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// example:
	//
	// DualStack
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The network type of the ALB instance. Valid values:
	//
	// 	- **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	//
	// 	- **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. In this case, the ALB instance can be accessed over the virtual private cloud (VPC) where the ALB instance is deployed.
	//
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the elastic IP address (EIP) bandwidth plan that is associated with the Internet-facing ALB instance.
	//
	// example:
	//
	// cbwp-bp1vevu8h3ieh****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name of the ALB instance.
	//
	// example:
	//
	// alb-95qnr2itwu9orb****.cn-hangzhou.alb.aliyuncs.com
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The configuration of deletion protection.
	DeletionProtectionConfig *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// The type of IPv6 address that is used by the ALB instance. Valid values:
	//
	// 	- **Internet**: The ALB instance uses a public IP address. The domain name of the ALB instance is resolved to the public IP address. Therefore, the ALB instance can be accessed over the Internet.
	//
	// 	- **Intranet**: The ALB instance uses a private IP address. The domain name of the ALB instance is resolved to the private IP address. Therefore, the ALB instance can be accessed over the VPC in which the ALB instance is deployed.
	//
	// example:
	//
	// Intranet
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The billing method of the ALB instance.
	LoadBalancerBillingConfig *GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The service status of the ALB instance. Valid values:
	//
	// 	- **Abnormal**
	//
	// 	- **Normal**
	//
	// example:
	//
	// Normal
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// The edition of the ALB instance. The features and billing rules vary based on the edition of the ALB instance. Valid values:
	//
	// 	- **Basic**
	//
	// 	- **Standard**
	//
	// 	- **StandardWithWaf**
	//
	// example:
	//
	// Standard
	LoadBalancerEdition *string `json:"LoadBalancerEdition,omitempty" xml:"LoadBalancerEdition,omitempty"`
	// The ALB instance ID.
	//
	// example:
	//
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ALB instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// alb1
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The type of the lock. Valid values:
	//
	// 	- **SecurityLocked**: The ALB instance is locked due to security reasons.
	//
	// 	- **RelatedResourceLocked**: The ALB instance is locked due to association issues.
	//
	// 	- **FinancialLocked**: The ALB instance is locked due to overdue payments.
	//
	// 	- **ResidualLocked**: The ALB instance is locked because the associated resources have overdue payments and the resources are released.
	LoadBalancerOperationLocks []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks `json:"LoadBalancerOperationLocks,omitempty" xml:"LoadBalancerOperationLocks,omitempty" type:"Repeated"`
	// The status of the ALB instance. Valid values:
	//
	// 	- **Inactive**: The ALB instance is disabled. ALB instances in the Inactive state do not forward traffic.
	//
	// 	- **Active**: The ALB instance is running.
	//
	// 	- **Provisioning**: The ALB instance is being created.
	//
	// 	- **Configuring**: The ALB instance is being modified.
	//
	// 	- **CreateFailed**: The system failed to create the ALB instance. In this case, you are not charged for the ALB instance. You can only delete the ALB instance.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The configuration read-only mode settings.
	ModificationProtectionConfig *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The region ID of the ALB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the security groups to which the ALB instance is added.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag value cannot contain `http://` or `https://`.
	Tags []*GetLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC in which the ALB instance is deployed.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. At most 10 zones are returned. If the current region supports two or more zones, at least two zones are returned.
	ZoneMappings []*GetLoadBalancerAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) GetAccessLogConfig() *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	return s.AccessLogConfig
}

func (s *GetLoadBalancerAttributeResponseBody) GetAddressAllocatedMode() *string {
	return s.AddressAllocatedMode
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

func (s *GetLoadBalancerAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
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

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerBussinessStatus() *string {
	return s.LoadBalancerBussinessStatus
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerEdition() *string {
	return s.LoadBalancerEdition
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerOperationLocks() []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	return s.LoadBalancerOperationLocks
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *GetLoadBalancerAttributeResponseBody) GetModificationProtectionConfig() *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig {
	return s.ModificationProtectionConfig
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

func (s *GetLoadBalancerAttributeResponseBody) SetAccessLogConfig(v *GetLoadBalancerAttributeResponseBodyAccessLogConfig) *GetLoadBalancerAttributeResponseBody {
	s.AccessLogConfig = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressAllocatedMode(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressAllocatedMode = &v
	return s
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

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
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

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBussinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerEdition(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerEdition = &v
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

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerOperationLocks(v []*GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerOperationLocks = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetModificationProtectionConfig(v *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) *GetLoadBalancerAttributeResponseBody {
	s.ModificationProtectionConfig = v
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
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyAccessLogConfig struct {
	// The Log Service project.
	//
	// example:
	//
	// sls-setter
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The Logstore.
	//
	// example:
	//
	// test
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyAccessLogConfig) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyAccessLogConfig) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) GetLogProject() *string {
	return s.LogProject
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) GetLogStore() *string {
	return s.LogStore
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) SetLogProject(v string) *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	s.LogProject = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) SetLogStore(v string) *GetLoadBalancerAttributeResponseBodyAccessLogConfig {
	s.LogStore = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyAccessLogConfig) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig struct {
	// Indicates whether the deletion protection feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when the deletion protection feature was enabled. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-02T02:49:05Z
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
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

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabled(v bool) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) SetEnabledTime(v string) *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyDeletionProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyLoadBalancerBillingConfig struct {
	// The billing method.
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

type GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks struct {
	// The reason why the ALB instance is locked. This parameter is valid only if **LoadBalancerBussinessStatus*	- is set to **Abnormal**.
	//
	// example:
	//
	// nolock
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The lock type. Valid values:
	//
	// 	- **SecurityLocked**: The ALB instance is locked due to security reasons.
	//
	// 	- **RelatedResourceLocked**: The ALB instance is locked due to other resources that are associated with the ALB instance.
	//
	// 	- **FinancialLocked**: The ALB instance is locked due to overdue payments.
	//
	// 	- **ResidualLocked**: The ALB instance is locked because the associated resources have overdue payments and the resources are released.
	//
	// example:
	//
	// FinancialLocked
	LockType *string `json:"LockType,omitempty" xml:"LockType,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) GetLockReason() *string {
	return s.LockReason
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) GetLockType() *string {
	return s.LockType
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) SetLockReason(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	s.LockReason = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) SetLockType(v string) *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks {
	s.LockType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyLoadBalancerOperationLocks) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyModificationProtectionConfig struct {
	// The reason why the configuration read-only mode is enabled.
	//
	// The name must be 2 to 128 character characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// This parameter takes effect only if **Status*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Specifies whether the configuration read-only mode is enabled. Valid values:
	//
	// 	- **NonProtection**: The configuration read-only mode is disabled. In this case, the value of the **Reason*	- parameter that you specify does not take effect. If you set **Reason**, the value is cleared.
	//
	// 	- **ConsoleProtection**: The configuration read-only mode is enabled. In this case, the value of the **Reason*	- parameter takes effect.****
	//
	// >  If the parameter is set to **ConsoleProtection**, the configuration read-only mode is enabled. You cannot modify the configurations of the ALB instance in the ALB console. However, you can call API operations to modify the configurations of the ALB instance.
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

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *GetLoadBalancerAttributeResponseBodyModificationProtectionConfig) GetStatus() *string {
	return s.Status
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

type GetLoadBalancerAttributeResponseBodyTags struct {
	// The tag key.
	//
	// The tag key can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetLoadBalancerAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	// The address of the ALB instance.
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// The zone status. Valid values:
	//
	// - **Active**: The ALB instance is running.
	//
	// - **Stopped**: The ALB instance is disabled.
	//
	// - **Shifted**: The ALB instance is removed.
	//
	// - **Starting**: The ALB instance is starting.
	//
	// - **Stopping**: The ALB instance is stopping.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of an ALB instance.
	//
	// example:
	//
	// vsw-bp12mw1f8k3jgy****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the ALB instance.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/189196.html) operation to query the most recent zone list.
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
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	// An IPv4 address.
	//
	// This parameter takes effect when **AddressIPVersion*	- is set to **IPv4*	- or **DualStack**. The network type is determined by the value of **AddressType**.
	//
	// example:
	//
	// 192.168.10.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The elastic IP address (EIP).
	//
	// example:
	//
	// eip-uf6wm****1zj9
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
	// 10.0.1.181
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The health status of the private IPv4 address of the ALB instance.
	//
	// This parameter is returned only when the Status of the zone is Active.Valid values:
	//
	// - **Healthy**
	//
	// - **Unhealthy**
	//
	// example:
	//
	// Healthy
	IntranetAddressHcStatus *string `json:"IntranetAddressHcStatus,omitempty" xml:"IntranetAddressHcStatus,omitempty"`
	// The IPv4 link-local addresses. The IP addresses that the ALB instance uses to communicate with the backend servers.
	Ipv4LocalAddresses []*string `json:"Ipv4LocalAddresses,omitempty" xml:"Ipv4LocalAddresses,omitempty" type:"Repeated"`
	// An IPv6 address.
	//
	// This parameter takes effect only when **AddressIPVersion*	- is set to **DualStack**. The network type is determined by the value of **Ipv6AddressType**.
	//
	// example:
	//
	// 2408:XXXX:39d:eb00::/56
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The health status of the private IPv6 address of the ALB instance.
	//
	// This parameter is returned only when the Status of the zone is Active.Valid values:
	//
	// - **Healthy**
	//
	// - **Unhealthy**
	//
	// example:
	//
	// Healthy
	Ipv6AddressHcStatus *string `json:"Ipv6AddressHcStatus,omitempty" xml:"Ipv6AddressHcStatus,omitempty"`
	// The IPv6 link-local addresses. The IP addresses that the ALB instance uses to communicate with the backend servers.
	Ipv6LocalAddresses []*string `json:"Ipv6LocalAddresses,omitempty" xml:"Ipv6LocalAddresses,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetAddress() *string {
	return s.Address
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetAllocationId() *string {
	return s.AllocationId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetEipType() *string {
	return s.EipType
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIntranetAddress() *string {
	return s.IntranetAddress
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIntranetAddressHcStatus() *string {
	return s.IntranetAddressHcStatus
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv4LocalAddresses() []*string {
	return s.Ipv4LocalAddresses
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv6AddressHcStatus() *string {
	return s.Ipv6AddressHcStatus
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetIpv6LocalAddresses() []*string {
	return s.Ipv6LocalAddresses
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAddress(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetAllocationId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.AllocationId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetEipType(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.EipType = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIntranetAddress(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.IntranetAddress = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIntranetAddressHcStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.IntranetAddressHcStatus = &v
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

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv6AddressHcStatus(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv6AddressHcStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetIpv6LocalAddresses(v []*string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.Ipv6LocalAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) Validate() error {
	return dara.Validate(s)
}
