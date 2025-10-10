// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody
	GetLoadBalancers() []*ListLoadBalancersResponseBodyLoadBalancers
	SetMaxResults(v int32) *ListLoadBalancersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListLoadBalancersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListLoadBalancersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLoadBalancersResponseBody
	GetTotalCount() *int32
}

type ListLoadBalancersResponseBody struct {
	// A list of ALB instances.
	LoadBalancers []*ListLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is used to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBody) GetLoadBalancers() []*ListLoadBalancersResponseBodyLoadBalancers {
	return s.LoadBalancers
}

func (s *ListLoadBalancersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLoadBalancersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLoadBalancersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLoadBalancersResponseBody) SetLoadBalancers(v []*ListLoadBalancersResponseBodyLoadBalancers) *ListLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *ListLoadBalancersResponseBody) SetMaxResults(v int32) *ListLoadBalancersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetNextToken(v string) *ListLoadBalancersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetRequestId(v string) *ListLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLoadBalancersResponseBody) SetTotalCount(v int32) *ListLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLoadBalancersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancers struct {
	// The configurations of access logs.
	AccessLogConfig *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig `json:"AccessLogConfig,omitempty" xml:"AccessLogConfig,omitempty" type:"Struct"`
	// The mode in which IP addresses are allocated. Valid values:
	//
	// 	- **Fixed**: The ALB instance uses a static IP address.
	//
	// 	- **Dynamic**: dynamically allocates an IP address to each zone of the ALB instance.
	//
	// example:
	//
	// Fixed
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
	// The type of IP address that the ALB instance uses to provide services. Valid values:
	//
	// 	- **Internet**: The ALB instance is assigned a public IP address. The domain name is resolved to the public IP address. The ALB instance is accessible over the Internet.
	//
	// 	- **Intranet**: The ALB instance is assigned only a private IP address. The domain name is resolved to the private IP address. The ALB instance is accessible only within the VPC of the ALB instance.
	//
	// example:
	//
	// Intranet
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The ID of the Internet Shared Bandwidth instance that is associated with the Internet-facing ALB instance.
	//
	// example:
	//
	// cbwp-bp1vevu8h3ieh****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name.
	//
	// example:
	//
	// alb-95qnr2itwu9orb****.cn-hangzhou.alb.aliyuncs.com
	DNSName *string `json:"DNSName,omitempty" xml:"DNSName,omitempty"`
	// The configuration of the deletion protection feature.
	DeletionProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig `json:"DeletionProtectionConfig,omitempty" xml:"DeletionProtectionConfig,omitempty" type:"Struct"`
	// The type of IPv6 address used by the ALB instance. Valid values:
	//
	// 	- **Internet*	- The ALB instance is assigned a public IP address. The domain name is resolved to the public IP address. The ALB instance is accessible over the Internet.
	//
	// 	- **Intranet*	- The ALB instance is assigned only a private IP address. The domain name is resolved to the private IP address. The ALB instance is accessible only within the VPC of the ALB instance.
	//
	// example:
	//
	// Intranet
	Ipv6AddressType *string `json:"Ipv6AddressType,omitempty" xml:"Ipv6AddressType,omitempty"`
	// The billing information about the ALB instance.
	LoadBalancerBillingConfig *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig `json:"LoadBalancerBillingConfig,omitempty" xml:"LoadBalancerBillingConfig,omitempty" type:"Struct"`
	// The status of the ALB instance. Valid values:
	//
	// 	- **Abnormal**
	//
	// 	- **Normal**
	//
	// example:
	//
	// Normal
	LoadBalancerBussinessStatus *string `json:"LoadBalancerBussinessStatus,omitempty" xml:"LoadBalancerBussinessStatus,omitempty"`
	// The edition of the ALB instance. The features and billing rules vary based on the edition. Valid values:
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
	// The ID of the ALB instance.
	//
	// example:
	//
	// alb-o9ulmq5hgn68jk****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ALB instance.
	//
	// example:
	//
	// alb-instance-test
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The configuration of the operation lock.
	LoadBalancerOperationLocks []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks `json:"LoadBalancerOperationLocks,omitempty" xml:"LoadBalancerOperationLocks,omitempty" type:"Repeated"`
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
	// 	- **CreateFailed**: The system failed to create the ALB instance.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The configuration read-only mode settings.
	ModificationProtectionConfig *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig `json:"ModificationProtectionConfig,omitempty" xml:"ModificationProtectionConfig,omitempty" type:"Struct"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-atstuj3rtop****
	ResourceGroupId  *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The information about the tags.
	Tags []*ListLoadBalancersResponseBodyLoadBalancersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC in which the ALB instance is deployed.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetAccessLogConfig() *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	return s.AccessLogConfig
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetAddressAllocatedMode() *string {
	return s.AddressAllocatedMode
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetAddressType() *string {
	return s.AddressType
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetDNSName() *string {
	return s.DNSName
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetDeletionProtectionConfig() *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	return s.DeletionProtectionConfig
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetIpv6AddressType() *string {
	return s.Ipv6AddressType
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerBillingConfig() *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig {
	return s.LoadBalancerBillingConfig
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerBussinessStatus() *string {
	return s.LoadBalancerBussinessStatus
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerEdition() *string {
	return s.LoadBalancerEdition
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerOperationLocks() []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	return s.LoadBalancerOperationLocks
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetModificationProtectionConfig() *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	return s.ModificationProtectionConfig
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetTags() []*ListLoadBalancersResponseBodyLoadBalancersTags {
	return s.Tags
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) GetVpcId() *string {
	return s.VpcId
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAccessLogConfig(v *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AccessLogConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressAllocatedMode(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressAllocatedMode = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressIpVersion(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressIpVersion = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetAddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetBandwidthPackageId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetCreateTime(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.CreateTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDNSName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DNSName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetDeletionProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.DeletionProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetIpv6AddressType(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Ipv6AddressType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBillingConfig(v *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBillingConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerBussinessStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerBussinessStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerEdition(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerEdition = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerName(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerName = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerOperationLocks(v []*ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerOperationLocks = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetLoadBalancerStatus(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancerStatus = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetModificationProtectionConfig(v *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ModificationProtectionConfig = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetResourceGroupId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.ResourceGroupId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetSecurityGroupIds(v []*string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.SecurityGroupIds = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetTags(v []*ListLoadBalancersResponseBodyLoadBalancersTags) *ListLoadBalancersResponseBodyLoadBalancers {
	s.Tags = v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) SetVpcId(v string) *ListLoadBalancersResponseBodyLoadBalancers {
	s.VpcId = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancers) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig struct {
	// The Simple Log Service project.
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

func (s ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) GetLogProject() *string {
	return s.LogProject
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) GetLogStore() *string {
	return s.LogStore
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) SetLogProject(v string) *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	s.LogProject = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) SetLogStore(v string) *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig {
	s.LogStore = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersAccessLogConfig) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig struct {
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when deletion protection is enabled.
	//
	// example:
	//
	// 2022-08-02T02:49:05Z
	EnabledTime *string `json:"EnabledTime,omitempty" xml:"EnabledTime,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) GetEnabledTime() *string {
	return s.EnabledTime
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabled(v bool) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.Enabled = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) SetEnabledTime(v string) *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig {
	s.EnabledTime = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersDeletionProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig struct {
	// The billing method. Valid value:
	//
	// **PostPay**: You are charged for the ALB instance on a pay-as-you-go basis.
	//
	// example:
	//
	// PostPay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) GetPayType() *string {
	return s.PayType
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) SetPayType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig {
	s.PayType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerBillingConfig) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks struct {
	// The reason why the ALB instance is locked. This parameter is valid only if **LoadBalancerBussinessStatus*	- is set to **Abnormal**.
	//
	// example:
	//
	// Test LockReason
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The lock type. Valid values:
	//
	// 	- **SecurityLocked**: The ALB instance is locked due to security risks.
	//
	// 	- **RelatedResourceLocked**: The ALB instance is locked due to other resources associated with the ALB instance.
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

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) GetLockReason() *string {
	return s.LockReason
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) GetLockType() *string {
	return s.LockType
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) SetLockReason(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	s.LockReason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) SetLockType(v string) *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks {
	s.LockType = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersLoadBalancerOperationLocks) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig struct {
	// The reason why the configuration read-only mode is enabled.
	//
	// The reason must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter takes effect only if **Status*	- is set to **ConsoleProtection**.
	//
	// example:
	//
	// Test Reason
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether the configuration read-only mode is enabled. Valid values:
	//
	// 	- **NonProtection**: The configuration read-only mode is disabled. In this case, **Reason*	- is not returned. If **Reason*	- is set, the value is cleared.
	//
	// 	- **ConsoleProtection**: The configuration read-only mode is enabled. In this case, **Reason*	- is returned.****
	//
	// >  If the value is **ConsoleProtection**, the configuration read-only mode is enabled. You cannot modify the configurations of the ALB instance in the ALB console. However, you can call API operations to modify the configurations of the ALB instance.
	//
	// example:
	//
	// ConsoleProtection
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) GetReason() *string {
	return s.Reason
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) GetStatus() *string {
	return s.Status
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetReason(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Reason = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) SetStatus(v string) *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig {
	s.Status = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersModificationProtectionConfig) Validate() error {
	return dara.Validate(s)
}

type ListLoadBalancersResponseBodyLoadBalancersTags struct {
	// The tag key of the ALB instance.
	//
	// example:
	//
	// KeyTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the ALB instance.
	//
	// example:
	//
	// alueTest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancersResponseBodyLoadBalancersTags) GoString() string {
	return s.String()
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) GetKey() *string {
	return s.Key
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) GetValue() *string {
	return s.Value
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetKey(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Key = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) SetValue(v string) *ListLoadBalancersResponseBodyLoadBalancersTags {
	s.Value = &v
	return s
}

func (s *ListLoadBalancersResponseBodyLoadBalancersTags) Validate() error {
	return dara.Validate(s)
}
