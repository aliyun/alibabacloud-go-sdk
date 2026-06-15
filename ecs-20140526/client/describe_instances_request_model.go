// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeInstancesRequestFilter) *DescribeInstancesRequest
	GetFilter() []*DescribeInstancesRequestFilter
	SetAdditionalAttributes(v []*string) *DescribeInstancesRequest
	GetAdditionalAttributes() []*string
	SetDeviceAvailable(v bool) *DescribeInstancesRequest
	GetDeviceAvailable() *bool
	SetDryRun(v bool) *DescribeInstancesRequest
	GetDryRun() *bool
	SetEipAddresses(v string) *DescribeInstancesRequest
	GetEipAddresses() *string
	SetHpcClusterId(v string) *DescribeInstancesRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *DescribeInstancesRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *DescribeInstancesRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *DescribeInstancesRequest
	GetHttpTokens() *string
	SetImageId(v string) *DescribeInstancesRequest
	GetImageId() *string
	SetInnerIpAddresses(v string) *DescribeInstancesRequest
	GetInnerIpAddresses() *string
	SetInstanceChargeType(v string) *DescribeInstancesRequest
	GetInstanceChargeType() *string
	SetInstanceIds(v string) *DescribeInstancesRequest
	GetInstanceIds() *string
	SetInstanceName(v string) *DescribeInstancesRequest
	GetInstanceName() *string
	SetInstanceNetworkType(v string) *DescribeInstancesRequest
	GetInstanceNetworkType() *string
	SetInstanceType(v string) *DescribeInstancesRequest
	GetInstanceType() *string
	SetInstanceTypeFamily(v string) *DescribeInstancesRequest
	GetInstanceTypeFamily() *string
	SetInternetChargeType(v string) *DescribeInstancesRequest
	GetInternetChargeType() *string
	SetIoOptimized(v bool) *DescribeInstancesRequest
	GetIoOptimized() *bool
	SetIpv6Address(v []*string) *DescribeInstancesRequest
	GetIpv6Address() []*string
	SetKeyPairName(v string) *DescribeInstancesRequest
	GetKeyPairName() *string
	SetLockReason(v string) *DescribeInstancesRequest
	GetLockReason() *string
	SetMaxResults(v int32) *DescribeInstancesRequest
	GetMaxResults() *int32
	SetNeedSaleCycle(v bool) *DescribeInstancesRequest
	GetNeedSaleCycle() *bool
	SetNextToken(v string) *DescribeInstancesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesRequest
	GetPageSize() *int32
	SetPrivateIpAddresses(v string) *DescribeInstancesRequest
	GetPrivateIpAddresses() *string
	SetPublicIpAddresses(v string) *DescribeInstancesRequest
	GetPublicIpAddresses() *string
	SetRdmaIpAddresses(v string) *DescribeInstancesRequest
	GetRdmaIpAddresses() *string
	SetRegionId(v string) *DescribeInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstancesRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *DescribeInstancesRequest
	GetSecurityGroupId() *string
	SetStatus(v string) *DescribeInstancesRequest
	GetStatus() *string
	SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest
	GetTag() []*DescribeInstancesRequestTag
	SetVSwitchId(v string) *DescribeInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeInstancesRequest
	GetZoneId() *string
}

type DescribeInstancesRequest struct {
	Filter []*DescribeInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The list of additional instance attributes.
	//
	// example:
	//
	// META_OPTIONS
	AdditionalAttributes []*string `json:"AdditionalAttributes,omitempty" xml:"AdditionalAttributes,omitempty" type:"Repeated"`
	// > This parameter is in invitational preview and is not supported.
	//
	// example:
	//
	// false
	DeviceAvailable *bool `json:"DeviceAvailable,omitempty" xml:"DeviceAvailable,omitempty"`
	// Specifies whether to perform only a dry run for the request. Valid values:
	//
	// - true: Only checks the request without querying resources. Checks include AccessKey validity, RAM user permissions, and required parameters. If the check fails, an error is returned. If the check passes, the DryRunOperation error code is returned.
	//
	// - false: Sends a normal request. After passing the checks, a 2XX HTTP status code is returned and resources are queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The Elastic IP addresses of instances. This parameter takes effect only when InstanceNetworkType=vpc. Specify multiple IP addresses as a JSON array. You can specify up to 100 IP addresses. Separate IP addresses with commas (,).
	//
	// example:
	//
	// ["42.1.1.**", "42.1.2.**", … "42.1.10.**"]
	EipAddresses *string `json:"EipAddresses,omitempty" xml:"EipAddresses,omitempty"`
	// The ID of the HPC cluster to which the instance belongs.
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable access to instance metadata. Valid values:
	//
	// - enabled: enabled.
	//
	// - disabled: disabled.
	//
	// Default value: enabled.
	//
	// > For more information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/49122.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// > This parameter is not available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to enforce the use of IMDSv2 when accessing instance metadata. Valid values:
	//
	// - optional: does not enforce IMDSv2.
	//
	// - required: enforces IMDSv2. After this value is set, instance metadata cannot be accessed in standard mode.
	//
	// Default value: optional.
	//
	// > For more information about instance metadata access modes, see [Instance metadata access modes](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-bp67acfmxazb4p****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The private IP addresses of instances in the classic network. This parameter takes effect only when InstanceNetworkType=classic. Specify multiple IP addresses as a JSON array. You can specify up to 100 IP addresses. Separate IP addresses with commas (,).
	//
	// example:
	//
	// ["10.1.1.1", "10.1.2.1", … "10.1.10.1"]
	InnerIpAddresses *string `json:"InnerIpAddresses,omitempty" xml:"InnerIpAddresses,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - PostPaid: pay-as-you-go.
	//
	// - PrePaid: subscription.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The IDs of instances. Specify multiple instance IDs as a JSON array. You can specify up to 100 IDs. Separate IDs with commas (,).
	//
	// example:
	//
	// ["i-bp67acfmxazb4p****", "i-bp67acfmxazb4p****", … "i-bp67acfmxazb4p****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the instance. You can use the wildcard character \\	- for fuzzy search.
	//
	// example:
	//
	// Test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The network type of the instance. Valid values:
	//
	// - classic: classic network.
	//
	// - vpc: Virtual Private Cloud (VPC).
	//
	// example:
	//
	// vpc
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance family.
	//
	// example:
	//
	// ecs.g5
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The billing method for public bandwidth. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-data-transfer.
	//
	// > In **pay-by-data-transfer*	- mode, both inbound and outbound peak bandwidth represent upper limits and are not guaranteed service levels. During resource contention, peak bandwidth may be limited. If your business requires guaranteed bandwidth, use **pay-by-bandwidth*	- mode.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// Indicates whether the instance is I/O optimized. Valid values:
	//
	// - true: yes.
	//
	// - false: no.
	//
	// example:
	//
	// true
	IoOptimized *bool `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The IPv6 addresses assigned to the ENI.
	//
	// if can be null:
	// false
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The name of the SSH key pair used by the instance.
	//
	// example:
	//
	// KeyPairNameTest
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The reason why the resource is locked. Valid values:
	//
	// - financial: The instance is locked due to overdue payment.
	//
	// - security: The instance is locked for security reasons.
	//
	// - Recycling: The spot instance is locked and pending release.
	//
	// - dedicatedhostfinancial: The ECS instance is locked because the dedicated host has an overdue payment.
	//
	// - refunded: The instance is locked due to a refund.
	//
	// example:
	//
	// security
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maximum number of entries to return on each page. Maximum value: 100.
	//
	// Default value:
	//
	// - If you do not specify this parameter or specify a value less than 10, the default value is 10.
	//
	// - If you specify a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// > This parameter is in invitational preview and is not supported.
	//
	// example:
	//
	// false
	NeedSaleCycle *bool `json:"NeedSaleCycle,omitempty" xml:"NeedSaleCycle,omitempty"`
	// The pagination token. Set this parameter to the `NextToken` value returned in the last API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// > This parameter will be deprecated. We recommend that you use NextToken and MaxResults to perform paged queries.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// > This parameter will be deprecated. We recommend that you use NextToken and MaxResults to perform paged queries.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The private IP addresses of instances in a VPC. This parameter takes effect only when InstanceNetworkType=vpc. Specify multiple IP addresses as a JSON array. You can specify up to 100 IP addresses. Separate IP addresses with commas (,).
	//
	// example:
	//
	// ["172.16.1.1", "172.16.2.1", … "172.16.10.1"]
	PrivateIpAddresses *string `json:"PrivateIpAddresses,omitempty" xml:"PrivateIpAddresses,omitempty"`
	// The public IP addresses of instances. Specify multiple IP addresses as a JSON array. You can specify up to 100 IP addresses. Separate IP addresses with commas (,).
	//
	// example:
	//
	// ["42.1.1.**", "42.1.2.**", … "42.1.10.**"]
	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty"`
	// The RDMA IP address of the HPC instance.
	//
	// example:
	//
	// 10.10.10.102
	RdmaIpAddresses *string `json:"RdmaIpAddresses,omitempty" xml:"RdmaIpAddresses,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. When you use this parameter to filter resources, the number of resources cannot exceed 1,000.
	//
	// > Filtering by the default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group to which the instance belongs.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the instance. Valid values:
	//
	// - Pending: The instance is being created.
	//
	// - Running: The instance is running.
	//
	// - Starting: The instance is starting.
	//
	// - Stopping: The instance is stopping.
	//
	// - Stopped: The instance is stopped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tag []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp67acfmxazb4p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the Virtual Private Cloud (VPC).
	//
	// example:
	//
	// v-bp67acfmxazb4p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetFilter() []*DescribeInstancesRequestFilter {
	return s.Filter
}

func (s *DescribeInstancesRequest) GetAdditionalAttributes() []*string {
	return s.AdditionalAttributes
}

func (s *DescribeInstancesRequest) GetDeviceAvailable() *bool {
	return s.DeviceAvailable
}

func (s *DescribeInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeInstancesRequest) GetEipAddresses() *string {
	return s.EipAddresses
}

func (s *DescribeInstancesRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *DescribeInstancesRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *DescribeInstancesRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *DescribeInstancesRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *DescribeInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeInstancesRequest) GetInnerIpAddresses() *string {
	return s.InnerIpAddresses
}

func (s *DescribeInstancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeInstancesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeInstancesRequest) GetIoOptimized() *bool {
	return s.IoOptimized
}

func (s *DescribeInstancesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *DescribeInstancesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *DescribeInstancesRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstancesRequest) GetNeedSaleCycle() *bool {
	return s.NeedSaleCycle
}

func (s *DescribeInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetPrivateIpAddresses() *string {
	return s.PrivateIpAddresses
}

func (s *DescribeInstancesRequest) GetPublicIpAddresses() *string {
	return s.PublicIpAddresses
}

func (s *DescribeInstancesRequest) GetRdmaIpAddresses() *string {
	return s.RdmaIpAddresses
}

func (s *DescribeInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstancesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesRequest) GetTag() []*DescribeInstancesRequestTag {
	return s.Tag
}

func (s *DescribeInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesRequest) SetFilter(v []*DescribeInstancesRequestFilter) *DescribeInstancesRequest {
	s.Filter = v
	return s
}

func (s *DescribeInstancesRequest) SetAdditionalAttributes(v []*string) *DescribeInstancesRequest {
	s.AdditionalAttributes = v
	return s
}

func (s *DescribeInstancesRequest) SetDeviceAvailable(v bool) *DescribeInstancesRequest {
	s.DeviceAvailable = &v
	return s
}

func (s *DescribeInstancesRequest) SetDryRun(v bool) *DescribeInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeInstancesRequest) SetEipAddresses(v string) *DescribeInstancesRequest {
	s.EipAddresses = &v
	return s
}

func (s *DescribeInstancesRequest) SetHpcClusterId(v string) *DescribeInstancesRequest {
	s.HpcClusterId = &v
	return s
}

func (s *DescribeInstancesRequest) SetHttpEndpoint(v string) *DescribeInstancesRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *DescribeInstancesRequest) SetHttpPutResponseHopLimit(v int32) *DescribeInstancesRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *DescribeInstancesRequest) SetHttpTokens(v string) *DescribeInstancesRequest {
	s.HttpTokens = &v
	return s
}

func (s *DescribeInstancesRequest) SetImageId(v string) *DescribeInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInnerIpAddresses(v string) *DescribeInstancesRequest {
	s.InnerIpAddresses = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceChargeType(v string) *DescribeInstancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceNetworkType(v string) *DescribeInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceType(v string) *DescribeInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceTypeFamily(v string) *DescribeInstancesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeInstancesRequest) SetInternetChargeType(v string) *DescribeInstancesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetIoOptimized(v bool) *DescribeInstancesRequest {
	s.IoOptimized = &v
	return s
}

func (s *DescribeInstancesRequest) SetIpv6Address(v []*string) *DescribeInstancesRequest {
	s.Ipv6Address = v
	return s
}

func (s *DescribeInstancesRequest) SetKeyPairName(v string) *DescribeInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *DescribeInstancesRequest) SetLockReason(v string) *DescribeInstancesRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeInstancesRequest) SetMaxResults(v int32) *DescribeInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstancesRequest) SetNeedSaleCycle(v bool) *DescribeInstancesRequest {
	s.NeedSaleCycle = &v
	return s
}

func (s *DescribeInstancesRequest) SetNextToken(v string) *DescribeInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstancesRequest) SetOwnerAccount(v string) *DescribeInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstancesRequest) SetOwnerId(v int64) *DescribeInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetPrivateIpAddresses(v string) *DescribeInstancesRequest {
	s.PrivateIpAddresses = &v
	return s
}

func (s *DescribeInstancesRequest) SetPublicIpAddresses(v string) *DescribeInstancesRequest {
	s.PublicIpAddresses = &v
	return s
}

func (s *DescribeInstancesRequest) SetRdmaIpAddresses(v string) *DescribeInstancesRequest {
	s.RdmaIpAddresses = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceOwnerAccount(v string) *DescribeInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceOwnerId(v int64) *DescribeInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSecurityGroupId(v string) *DescribeInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v string) *DescribeInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstancesRequest) SetVSwitchId(v string) *DescribeInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVpcId(v string) *DescribeInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesRequest) SetZoneId(v string) *DescribeInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type DescribeInstancesRequestFilter struct {
	// The filter key used to query resources. Set this parameter to `CreationStartTime`. When you set both `Filter.1.Key` and `Filter.1.Value`, you can query resources created after the specified point in time.
	//
	// example:
	//
	// CreationStartTime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The filter value used to query resources. You must also specify the `Filter.1.Key` parameter when you specify this parameter. The value must be in the format `yyyy-MM-ddTHH:mmZ` (UTC+0).
	//
	// example:
	//
	// 2017-12-05T22:40Z
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestFilter) SetKey(v string) *DescribeInstancesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestFilter) SetValue(v string) *DescribeInstancesRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestFilter) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesRequestTag struct {
	// The tag key.
	//
	// > To improve compatibility, we recommend that you use the `Tag.N.Key` parameter instead.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
