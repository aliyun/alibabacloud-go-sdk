// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuOptions(v *RunInstancesRequestCpuOptions) *RunInstancesRequest
	GetCpuOptions() *RunInstancesRequestCpuOptions
	SetHibernationOptions(v *RunInstancesRequestHibernationOptions) *RunInstancesRequest
	GetHibernationOptions() *RunInstancesRequestHibernationOptions
	SetPrivatePoolOptions(v *RunInstancesRequestPrivatePoolOptions) *RunInstancesRequest
	GetPrivatePoolOptions() *RunInstancesRequestPrivatePoolOptions
	SetSchedulerOptions(v *RunInstancesRequestSchedulerOptions) *RunInstancesRequest
	GetSchedulerOptions() *RunInstancesRequestSchedulerOptions
	SetSecurityOptions(v *RunInstancesRequestSecurityOptions) *RunInstancesRequest
	GetSecurityOptions() *RunInstancesRequestSecurityOptions
	SetSystemDisk(v *RunInstancesRequestSystemDisk) *RunInstancesRequest
	GetSystemDisk() *RunInstancesRequestSystemDisk
	SetAffinity(v string) *RunInstancesRequest
	GetAffinity() *string
	SetAmount(v int32) *RunInstancesRequest
	GetAmount() *int32
	SetArn(v []*RunInstancesRequestArn) *RunInstancesRequest
	GetArn() []*RunInstancesRequestArn
	SetAutoPay(v bool) *RunInstancesRequest
	GetAutoPay() *bool
	SetAutoReleaseTime(v string) *RunInstancesRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *RunInstancesRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *RunInstancesRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *RunInstancesRequest
	GetClientToken() *string
	SetClockOptions(v *RunInstancesRequestClockOptions) *RunInstancesRequest
	GetClockOptions() *RunInstancesRequestClockOptions
	SetCreditSpecification(v string) *RunInstancesRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest
	GetDataDisk() []*RunInstancesRequestDataDisk
	SetDedicatedHostId(v string) *RunInstancesRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *RunInstancesRequest
	GetDeletionProtection() *bool
	SetDeploymentSetGroupNo(v int32) *RunInstancesRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *RunInstancesRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *RunInstancesRequest
	GetDescription() *string
	SetDryRun(v bool) *RunInstancesRequest
	GetDryRun() *bool
	SetHostName(v string) *RunInstancesRequest
	GetHostName() *string
	SetHostNames(v []*string) *RunInstancesRequest
	GetHostNames() []*string
	SetHpcClusterId(v string) *RunInstancesRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *RunInstancesRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *RunInstancesRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *RunInstancesRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *RunInstancesRequest
	GetImageFamily() *string
	SetImageId(v string) *RunInstancesRequest
	GetImageId() *string
	SetImageOptions(v *RunInstancesRequestImageOptions) *RunInstancesRequest
	GetImageOptions() *RunInstancesRequestImageOptions
	SetInstanceChargeType(v string) *RunInstancesRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunInstancesRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunInstancesRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunInstancesRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *RunInstancesRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *RunInstancesRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *RunInstancesRequest
	GetIoOptimized() *string
	SetIpv6Address(v []*string) *RunInstancesRequest
	GetIpv6Address() []*string
	SetIpv6AddressCount(v int32) *RunInstancesRequest
	GetIpv6AddressCount() *int32
	SetIsp(v string) *RunInstancesRequest
	GetIsp() *string
	SetKeyPairName(v string) *RunInstancesRequest
	GetKeyPairName() *string
	SetLaunchTemplateId(v string) *RunInstancesRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *RunInstancesRequest
	GetLaunchTemplateName() *string
	SetLaunchTemplateVersion(v int64) *RunInstancesRequest
	GetLaunchTemplateVersion() *int64
	SetMinAmount(v int32) *RunInstancesRequest
	GetMinAmount() *int32
	SetNetworkInterface(v []*RunInstancesRequestNetworkInterface) *RunInstancesRequest
	GetNetworkInterface() []*RunInstancesRequestNetworkInterface
	SetNetworkInterfaceQueueNumber(v int32) *RunInstancesRequest
	GetNetworkInterfaceQueueNumber() *int32
	SetNetworkOptions(v *RunInstancesRequestNetworkOptions) *RunInstancesRequest
	GetNetworkOptions() *RunInstancesRequestNetworkOptions
	SetOwnerAccount(v string) *RunInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RunInstancesRequest
	GetOwnerId() *int64
	SetPassword(v string) *RunInstancesRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunInstancesRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *RunInstancesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RunInstancesRequest
	GetPeriodUnit() *string
	SetPrivateDnsNameOptions(v *RunInstancesRequestPrivateDnsNameOptions) *RunInstancesRequest
	GetPrivateDnsNameOptions() *RunInstancesRequestPrivateDnsNameOptions
	SetPrivateIpAddress(v string) *RunInstancesRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *RunInstancesRequest
	GetRamRoleName() *string
	SetRegionId(v string) *RunInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RunInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RunInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RunInstancesRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *RunInstancesRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *RunInstancesRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *RunInstancesRequest
	GetSecurityGroupIds() []*string
	SetSpotDuration(v int32) *RunInstancesRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *RunInstancesRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimit(v float32) *RunInstancesRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *RunInstancesRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *RunInstancesRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *RunInstancesRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest
	GetTag() []*RunInstancesRequestTag
	SetTenancy(v string) *RunInstancesRequest
	GetTenancy() *string
	SetUniqueSuffix(v bool) *RunInstancesRequest
	GetUniqueSuffix() *bool
	SetUserData(v string) *RunInstancesRequest
	GetUserData() *string
	SetVSwitchId(v string) *RunInstancesRequest
	GetVSwitchId() *string
	SetZoneId(v string) *RunInstancesRequest
	GetZoneId() *string
}

type RunInstancesRequest struct {
	CpuOptions         *RunInstancesRequestCpuOptions         `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	HibernationOptions *RunInstancesRequestHibernationOptions `json:"HibernationOptions,omitempty" xml:"HibernationOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *RunInstancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SchedulerOptions   *RunInstancesRequestSchedulerOptions   `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SecurityOptions    *RunInstancesRequestSecurityOptions    `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	SystemDisk         *RunInstancesRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether the instance is associated with a dedicated host. Valid values:
	//
	// - default: The instance is not associated with a dedicated host. When an instance that has economical mode enabled is restarted after it is stopped, the instance is deployed to another dedicated host in the automatic deployment resource pool if the resources of the original dedicated host are insufficient.
	//
	// - host: The instance is associated with a dedicated host. When an instance that has economical mode enabled is restarted after it is stopped, the instance remains on the original dedicated host. If the resources of the original dedicated host are insufficient, the instance fails to restart.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of ECS instances to create. Valid values: 1 to 100.
	//
	// The number of ECS instances that are created depends on the values of Amount and MinAmount:
	//
	// - If MinAmount is not specified, instances are created based on the value of Amount. If the inventory is insufficient, the API returns a failure and no instances are created.
	//
	// - If MinAmount is specified:
	//
	//   - If the available inventory < MinAmount, no ECS instances are created and the API returns a failure.
	//
	//   - If MinAmount ≤ available inventory < Amount, instances are created based on the available inventory and the API returns a success.
	//
	//   - If the available inventory ≥ Amount, instances are created based on the value of Amount and the API returns a success.
	//
	// Default value: 1.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// > This parameter is not publicly available.
	Arn []*RunInstancesRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to automatically complete the payment when you create the instance. Valid values:
	//
	// - true: The payment is automatically completed.
	//
	//     > If the balance of your payment method is insufficient, an abnormal order is generated and can only be canceled. If your payment method has an insufficient balance, set `AutoPay` to `false`. An unpaid order is generated, and you can log on to the ECS console to complete the payment.
	//
	// - false: An order is generated but the payment is not completed.
	//
	//     > When `InstanceChargeType` is set to `PostPaid`, `AutoPay` cannot be set to `false`.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The automatic release time of the pay-as-you-go instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the UTC+0 time zone. The format is `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Valid values:
	//
	// <props="china">
	//
	// - PeriodUnit=Week: 1, 2, and 3.
	//
	// - PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The clock-related property parameters of the instance.
	ClockOptions *RunInstancesRequestClockOptions `json:"ClockOptions,omitempty" xml:"ClockOptions,omitempty" type:"Struct"`
	// The running mode of the burstable instance. Valid values:
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The list of data disk information.
	DataDisk []*RunInstancesRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the dedicated host.
	//
	// <props="china">You can call [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) to query the list of dedicated host IDs.
	//
	// <props="intl">You can call [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) to query the list of dedicated host IDs.
	//
	// 	Notice: Dedicated hosts do not support the creation of spot instances. If you specify the `DedicatedHostId` parameter, the `SpotStrategy` and `SpotPriceLimit` settings in the request are automatically ignored.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The release protection property of the instance. Specifies whether the instance can be released from the console or by calling the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation. Valid values:
	//
	// -  true: Enables release protection for the instance.
	//
	// -  false: Disables release protection for the instance.
	//
	// Default value: false.
	//
	// > This property applies only to pay-as-you-go instances and only restricts manual release operations. It does not take effect on system-initiated release operations.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// If the deployment set uses the high availability group strategy (AvailabilityGroup), you can use this parameter to specify the group number of the instance in the deployment set. Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set.
	//
	// example:
	//
	// ds-bp1brhwhoqinyjd6****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Instance_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// -  true: performs a dry run without creating the instance. The system checks the required parameters, request syntax, business restrictions, and ECS inventory. If the check fails, the corresponding error is returned. If the check passes, the error code `DryRunOperation` is returned.
	//
	// -  false (default): performs a dry run and sends the request. If the check passes, the instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the instance. The following limits apply:
	//
	// example:
	//
	// k8s-node-[1,4]-ecshost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// Specifies a different hostname for each instance when you create multiple instances.
	//
	// example:
	//
	// ecs-host-01
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// The ID of the HPC cluster to which the instance belongs.
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the China Reinforced mode (IMDSv2) for accessing instance metadata. Valid values:
	//
	// - optional: does not forcefully use the China Reinforced mode.
	//
	// - required: forcefully uses the China Reinforced mode. After this value is set, the normal mode cannot be used to access instance metadata.
	//
	// Default value: optional.
	//
	// >For more information about the modes for accessing instance metadata, see [Instance metadata access mode](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available image from the specified image family to create the instance.
	//
	// The name must be 2 to 128 characters in length. The name cannot start with a special character, a digit, http://, or https://. The name can contain only the following special characters: periods (.), underscores (_), hyphens (-), and colons (:).
	//
	// Take note of the following items:
	//
	// - If you set the ImageId parameter, you cannot set this parameter.
	//
	// - If you do not set the ImageId parameter but the launch template specified by LaunchTemplateId or LaunchTemplateName has ImageId configured, you cannot set this parameter.
	//
	// - If you do not set ImageId and the launch template specified by LaunchTemplateId or LaunchTemplateName does not have ImageId configured, you can set this parameter.
	//
	// - If you do not set ImageId and do not set LaunchTemplateId or LaunchTemplateName, you can set this parameter.
	//
	// > For information about image families associated with Alibaba Cloud public images, see [Overview of public images](https://help.aliyun.com/document_detail/108393.html).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The image ID. Specifies the image resource used to start the instance. You can call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available image resources. If you do not specify `LaunchTemplateId` or `LaunchTemplateName` to determine a launch template, and do not specify `ImageFamily` to use the latest available image from an image family, `ImageId` is required.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image-related property information.
	ImageOptions *RunInstancesRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// -  PrePaid: subscription.
	//
	// -  PostPaid: pay-as-you-go.
	//
	// Default value: PostPaid.
	//
	// <props="china">If you set this parameter to PrePaid, make sure that your account supports balance payment or credit payment. Otherwise, the `InvalidPayMethod` error is returned.
	//
	// <props="intl">If you set this parameter to PrePaid, make sure that your account supports credit payment. Otherwise, the `InvalidPayMethod` error is returned.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and can contain characters from the Unicode letter category (including English and Chinese characters) and digits. The name can contain colons (:), underscores (_), periods (.), or hyphens (-). Default value: the `InstanceId` of the instance.
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type of the instance. If you do not specify `LaunchTemplateId` or `LaunchTemplateName` to determine the launch template, `InstanceType` is required.
	//
	// - Instance type selection: See [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the performance data of an instance type. You can also see [Best practices for instance type selection](https://help.aliyun.com/document_detail/58291.html) to learn how to select instance types.
	//
	// - Stock query: Call [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) to query the resource availability in a specific region or zone.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// - PayByBandwidth: pay-by-bandwidth.
	//
	// - PayByTraffic: pay-by-traffic.
	//
	// Default value: PayByTraffic.
	//
	// > In **pay-by-traffic*	- mode, the peak inbound and outbound bandwidths are both upper limits and are not guaranteed. When resource contention occurs, the peak bandwidth may be throttled. If your business requires guaranteed bandwidth, use the **pay-by-bandwidth*	- mode.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth, in Mbit/s. Valid values:
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth, in Mbit/s. Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is an I/O optimized instance. The default value for [retired instance types](https://help.aliyun.com/document_detail/55263.html) is none, which indicates that I/O optimization is disabled. The default value for other instance types is optimized. Valid values:
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// Specifies one or more IPv6 addresses for the primary ENI. You can specify up to 10 IPv6 addresses. Valid values of N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::***`.
	//
	// Take note of the following items:
	//
	// - If you set `Ipv6Address.N`, you must set `Amount` to 1 and cannot set `Ipv6AddressCount`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `Ipv6Addresses.N` or `Ipv6AddressCount`. Instead, set `NetworkInterface.N.Ipv6Addresses.N` or `NetworkInterface.N.Ipv6AddressCount`.
	//
	// example:
	//
	// Ipv6Address.1=2001:db8:1234:1a00::***
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of randomly generated IPv6 addresses to assign to the primary ENI. Valid values: 1 to 10.
	//
	// Take note of the following items:
	//
	// - You cannot specify both `Ipv6Address.N` and `Ipv6AddressCount`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `Ipv6Address.N` or `Ipv6AddressCount`. Instead, specify `NetworkInterface.N.Ipv6Address.N` or `NetworkInterface.N.Ipv6AddressCount`.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The launch template ID. For more information, call [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The launch template name.
	//
	// example:
	//
	// LaunchTemplate_Name
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The launch template version. If you specify `LaunchTemplateId` or `LaunchTemplateName` but do not specify the launch template version, the default version is used.
	//
	// example:
	//
	// 3
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The minimum number of ECS instances to purchase. Valid values: 1 to 100.
	//
	// example:
	//
	// 2
	MinAmount *int32 `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	// The network interface controller (NIC) information.
	NetworkInterface []*RunInstancesRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
	// The number of queues supported by the primary ENI. Take note of the following items:
	//
	// - The value cannot exceed the maximum number of queues per ENI allowed by the instance type.
	//
	// - The total number of queues for all ENIs on the instance cannot exceed the queue quota allowed by the instance type. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the MaximumQueueNumberPerEni and TotalEniQueueQuantity fields for the maximum number of queues per ENI and the total queue quota of an instance type.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `NetworkInterfaceQueueNumber`. Set `NetworkInterface.N.QueueNumber` instead.
	//
	// example:
	//
	// 8
	NetworkInterfaceQueueNumber *int32 `json:"NetworkInterfaceQueueNumber,omitempty" xml:"NetworkInterfaceQueueNumber,omitempty"`
	// The network-related property parameters.
	NetworkOptions *RunInstancesRequestNetworkOptions `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty" type:"Struct"`
	OwnerAccount   *string                            `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the preset password of the image. Valid values:
	//
	// - true: The preset password of the image is used.
	//
	// - false: The preset password of the image is not used.
	//
	// Default value: false.
	//
	// > When you use this parameter, the Password parameter must be empty. Make sure that the image has a preset password.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The duration of the subscription. Unit: specified by PeriodUnit. This parameter is required and takes effect only when InstanceChargeType is set to PrePaid. If DedicatedHostId is specified, the value of this parameter cannot exceed the remaining subscription duration of the dedicated host. Valid values:
	//
	// <props="china">
	//
	// - If PeriodUnit is set to Week, valid values of Period: 1, 2, 3, and 4.
	//
	// - If PeriodUnit is set to Month, valid values of Period: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	//
	//
	// <props="intl">If PeriodUnit is set to Month, valid values of Period: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// <props="china">
	//
	// - Week
	//
	// - Month (default)
	//
	//
	//
	// <props="intl">Month (default).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private DNS name configuration of the instance.
	PrivateDnsNameOptions *RunInstancesRequestPrivateDnsNameOptions `json:"PrivateDnsNameOptions,omitempty" xml:"PrivateDnsNameOptions,omitempty" type:"Struct"`
	// The private IP address of the instance. For a VPC-type ECS instance, the private IP address must be from the idle CIDR block of the vSwitch specified by `VSwitchId`.
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance RAM role. You can call the RAM API [ListRoles](https://help.aliyun.com/document_detail/28713.html) to query the instance RAM roles that you have created.
	//
	// example:
	//
	// RAM_Name
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the new instance belongs. Instances within the same security group can communicate with each other. The maximum number of instances that a security group can contain depends on the security group type. For more information, see the security group section in [Limits](~~25412#SecurityGroupQuota~~).
	//
	// > The `SecurityGroupId` parameter determines the network type of the instance. For example, if the specified security group is of the Virtual Private Cloud (VPC) type, the instance is a VPC-type instance, and you must also specify the `VSwitchId` parameter.
	//
	// If you do not set `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template, the security group ID is required. Take note of the following items:
	//
	// - You can set `SecurityGroupId` to specify a single security group, or set `SecurityGroupIds.N` to specify one or more security groups. However, you cannot specify both `SecurityGroupId` and `SecurityGroupIds.N` at the same time.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `SecurityGroupId` or `SecurityGroupIds.N`. In this case, you can only set `NetworkInterface.N.SecurityGroupId` or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Adds the instance to multiple security groups. The valid values of N depend on the maximum number of security groups to which an instance can belong. For more information, see [Security group limits](https://help.aliyun.com/document_detail/101348.html).
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The retention period of the spot instance, in hours. Valid values:
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption pattern of the spot instance. Valid values:
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the instance. This value supports up to three decimal places. This parameter takes effect when the `SpotStrategy` parameter is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter takes effect only when `InstanceChargeType` is set to `PostPaid`. Valid values:
	//
	// - NoSpot: a regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: a spot instance with a maximum hourly price.
	//
	// - SpotAsPriceGo: a spot instance for which the system automatically bids, following the current market price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. Valid values: greater than or equal to 1.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tag information for the instance, disks, and primary ENI.
	Tag []*RunInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// Specifies whether to automatically append sequential suffixes to `HostName` and `InstanceName` when you create multiple instances. The sequential suffix ranges from 001 to 999. Valid values:
	//
	// example:
	//
	// true
	UniqueSuffix *bool `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
	// The instance user data. The data must be Base64-encoded. The maximum size of the raw data before Base64 encoding is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID. If you are creating a VPC-type ECS instance, you must specify a vSwitch ID. The security group and the vSwitch must belong to the same VPC. You can call [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) to query created vSwitches.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query the zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RunInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequest) GoString() string {
	return s.String()
}

func (s *RunInstancesRequest) GetCpuOptions() *RunInstancesRequestCpuOptions {
	return s.CpuOptions
}

func (s *RunInstancesRequest) GetHibernationOptions() *RunInstancesRequestHibernationOptions {
	return s.HibernationOptions
}

func (s *RunInstancesRequest) GetPrivatePoolOptions() *RunInstancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *RunInstancesRequest) GetSchedulerOptions() *RunInstancesRequestSchedulerOptions {
	return s.SchedulerOptions
}

func (s *RunInstancesRequest) GetSecurityOptions() *RunInstancesRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *RunInstancesRequest) GetSystemDisk() *RunInstancesRequestSystemDisk {
	return s.SystemDisk
}

func (s *RunInstancesRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *RunInstancesRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *RunInstancesRequest) GetArn() []*RunInstancesRequestArn {
	return s.Arn
}

func (s *RunInstancesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RunInstancesRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *RunInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunInstancesRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *RunInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunInstancesRequest) GetClockOptions() *RunInstancesRequestClockOptions {
	return s.ClockOptions
}

func (s *RunInstancesRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *RunInstancesRequest) GetDataDisk() []*RunInstancesRequestDataDisk {
	return s.DataDisk
}

func (s *RunInstancesRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *RunInstancesRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunInstancesRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *RunInstancesRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RunInstancesRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunInstancesRequest) GetHostNames() []*string {
	return s.HostNames
}

func (s *RunInstancesRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *RunInstancesRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *RunInstancesRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *RunInstancesRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *RunInstancesRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *RunInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunInstancesRequest) GetImageOptions() *RunInstancesRequestImageOptions {
	return s.ImageOptions
}

func (s *RunInstancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunInstancesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunInstancesRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *RunInstancesRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *RunInstancesRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *RunInstancesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *RunInstancesRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *RunInstancesRequest) GetIsp() *string {
	return s.Isp
}

func (s *RunInstancesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunInstancesRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *RunInstancesRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *RunInstancesRequest) GetLaunchTemplateVersion() *int64 {
	return s.LaunchTemplateVersion
}

func (s *RunInstancesRequest) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *RunInstancesRequest) GetNetworkInterface() []*RunInstancesRequestNetworkInterface {
	return s.NetworkInterface
}

func (s *RunInstancesRequest) GetNetworkInterfaceQueueNumber() *int32 {
	return s.NetworkInterfaceQueueNumber
}

func (s *RunInstancesRequest) GetNetworkOptions() *RunInstancesRequestNetworkOptions {
	return s.NetworkOptions
}

func (s *RunInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RunInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RunInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *RunInstancesRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunInstancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RunInstancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunInstancesRequest) GetPrivateDnsNameOptions() *RunInstancesRequestPrivateDnsNameOptions {
	return s.PrivateDnsNameOptions
}

func (s *RunInstancesRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RunInstancesRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *RunInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RunInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RunInstancesRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *RunInstancesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunInstancesRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *RunInstancesRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *RunInstancesRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *RunInstancesRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *RunInstancesRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunInstancesRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *RunInstancesRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *RunInstancesRequest) GetTag() []*RunInstancesRequestTag {
	return s.Tag
}

func (s *RunInstancesRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *RunInstancesRequest) GetUniqueSuffix() *bool {
	return s.UniqueSuffix
}

func (s *RunInstancesRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RunInstancesRequest) SetCpuOptions(v *RunInstancesRequestCpuOptions) *RunInstancesRequest {
	s.CpuOptions = v
	return s
}

func (s *RunInstancesRequest) SetHibernationOptions(v *RunInstancesRequestHibernationOptions) *RunInstancesRequest {
	s.HibernationOptions = v
	return s
}

func (s *RunInstancesRequest) SetPrivatePoolOptions(v *RunInstancesRequestPrivatePoolOptions) *RunInstancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *RunInstancesRequest) SetSchedulerOptions(v *RunInstancesRequestSchedulerOptions) *RunInstancesRequest {
	s.SchedulerOptions = v
	return s
}

func (s *RunInstancesRequest) SetSecurityOptions(v *RunInstancesRequestSecurityOptions) *RunInstancesRequest {
	s.SecurityOptions = v
	return s
}

func (s *RunInstancesRequest) SetSystemDisk(v *RunInstancesRequestSystemDisk) *RunInstancesRequest {
	s.SystemDisk = v
	return s
}

func (s *RunInstancesRequest) SetAffinity(v string) *RunInstancesRequest {
	s.Affinity = &v
	return s
}

func (s *RunInstancesRequest) SetAmount(v int32) *RunInstancesRequest {
	s.Amount = &v
	return s
}

func (s *RunInstancesRequest) SetArn(v []*RunInstancesRequestArn) *RunInstancesRequest {
	s.Arn = v
	return s
}

func (s *RunInstancesRequest) SetAutoPay(v bool) *RunInstancesRequest {
	s.AutoPay = &v
	return s
}

func (s *RunInstancesRequest) SetAutoReleaseTime(v string) *RunInstancesRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *RunInstancesRequest) SetAutoRenew(v bool) *RunInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunInstancesRequest) SetAutoRenewPeriod(v int32) *RunInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *RunInstancesRequest) SetClientToken(v string) *RunInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RunInstancesRequest) SetClockOptions(v *RunInstancesRequestClockOptions) *RunInstancesRequest {
	s.ClockOptions = v
	return s
}

func (s *RunInstancesRequest) SetCreditSpecification(v string) *RunInstancesRequest {
	s.CreditSpecification = &v
	return s
}

func (s *RunInstancesRequest) SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest {
	s.DataDisk = v
	return s
}

func (s *RunInstancesRequest) SetDedicatedHostId(v string) *RunInstancesRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *RunInstancesRequest) SetDeletionProtection(v bool) *RunInstancesRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunInstancesRequest) SetDeploymentSetGroupNo(v int32) *RunInstancesRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *RunInstancesRequest) SetDeploymentSetId(v string) *RunInstancesRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunInstancesRequest) SetDescription(v string) *RunInstancesRequest {
	s.Description = &v
	return s
}

func (s *RunInstancesRequest) SetDryRun(v bool) *RunInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *RunInstancesRequest) SetHostName(v string) *RunInstancesRequest {
	s.HostName = &v
	return s
}

func (s *RunInstancesRequest) SetHostNames(v []*string) *RunInstancesRequest {
	s.HostNames = v
	return s
}

func (s *RunInstancesRequest) SetHpcClusterId(v string) *RunInstancesRequest {
	s.HpcClusterId = &v
	return s
}

func (s *RunInstancesRequest) SetHttpEndpoint(v string) *RunInstancesRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *RunInstancesRequest) SetHttpPutResponseHopLimit(v int32) *RunInstancesRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *RunInstancesRequest) SetHttpTokens(v string) *RunInstancesRequest {
	s.HttpTokens = &v
	return s
}

func (s *RunInstancesRequest) SetImageFamily(v string) *RunInstancesRequest {
	s.ImageFamily = &v
	return s
}

func (s *RunInstancesRequest) SetImageId(v string) *RunInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *RunInstancesRequest) SetImageOptions(v *RunInstancesRequestImageOptions) *RunInstancesRequest {
	s.ImageOptions = v
	return s
}

func (s *RunInstancesRequest) SetInstanceChargeType(v string) *RunInstancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceName(v string) *RunInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceType(v string) *RunInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesRequest) SetInternetChargeType(v string) *RunInstancesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetInternetMaxBandwidthIn(v int32) *RunInstancesRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *RunInstancesRequest) SetInternetMaxBandwidthOut(v int32) *RunInstancesRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunInstancesRequest) SetIoOptimized(v string) *RunInstancesRequest {
	s.IoOptimized = &v
	return s
}

func (s *RunInstancesRequest) SetIpv6Address(v []*string) *RunInstancesRequest {
	s.Ipv6Address = v
	return s
}

func (s *RunInstancesRequest) SetIpv6AddressCount(v int32) *RunInstancesRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *RunInstancesRequest) SetIsp(v string) *RunInstancesRequest {
	s.Isp = &v
	return s
}

func (s *RunInstancesRequest) SetKeyPairName(v string) *RunInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateId(v string) *RunInstancesRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateName(v string) *RunInstancesRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateVersion(v int64) *RunInstancesRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *RunInstancesRequest) SetMinAmount(v int32) *RunInstancesRequest {
	s.MinAmount = &v
	return s
}

func (s *RunInstancesRequest) SetNetworkInterface(v []*RunInstancesRequestNetworkInterface) *RunInstancesRequest {
	s.NetworkInterface = v
	return s
}

func (s *RunInstancesRequest) SetNetworkInterfaceQueueNumber(v int32) *RunInstancesRequest {
	s.NetworkInterfaceQueueNumber = &v
	return s
}

func (s *RunInstancesRequest) SetNetworkOptions(v *RunInstancesRequestNetworkOptions) *RunInstancesRequest {
	s.NetworkOptions = v
	return s
}

func (s *RunInstancesRequest) SetOwnerAccount(v string) *RunInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunInstancesRequest) SetOwnerId(v int64) *RunInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RunInstancesRequest) SetPassword(v string) *RunInstancesRequest {
	s.Password = &v
	return s
}

func (s *RunInstancesRequest) SetPasswordInherit(v bool) *RunInstancesRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunInstancesRequest) SetPeriod(v int32) *RunInstancesRequest {
	s.Period = &v
	return s
}

func (s *RunInstancesRequest) SetPeriodUnit(v string) *RunInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunInstancesRequest) SetPrivateDnsNameOptions(v *RunInstancesRequestPrivateDnsNameOptions) *RunInstancesRequest {
	s.PrivateDnsNameOptions = v
	return s
}

func (s *RunInstancesRequest) SetPrivateIpAddress(v string) *RunInstancesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunInstancesRequest) SetRamRoleName(v string) *RunInstancesRequest {
	s.RamRoleName = &v
	return s
}

func (s *RunInstancesRequest) SetRegionId(v string) *RunInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RunInstancesRequest) SetResourceGroupId(v string) *RunInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunInstancesRequest) SetResourceOwnerAccount(v string) *RunInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunInstancesRequest) SetResourceOwnerId(v int64) *RunInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityEnhancementStrategy(v string) *RunInstancesRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityGroupId(v string) *RunInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityGroupIds(v []*string) *RunInstancesRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *RunInstancesRequest) SetSpotDuration(v int32) *RunInstancesRequest {
	s.SpotDuration = &v
	return s
}

func (s *RunInstancesRequest) SetSpotInterruptionBehavior(v string) *RunInstancesRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *RunInstancesRequest) SetSpotPriceLimit(v float32) *RunInstancesRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *RunInstancesRequest) SetSpotStrategy(v string) *RunInstancesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetStorageSetId(v string) *RunInstancesRequest {
	s.StorageSetId = &v
	return s
}

func (s *RunInstancesRequest) SetStorageSetPartitionNumber(v int32) *RunInstancesRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *RunInstancesRequest) SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest {
	s.Tag = v
	return s
}

func (s *RunInstancesRequest) SetTenancy(v string) *RunInstancesRequest {
	s.Tenancy = &v
	return s
}

func (s *RunInstancesRequest) SetUniqueSuffix(v bool) *RunInstancesRequest {
	s.UniqueSuffix = &v
	return s
}

func (s *RunInstancesRequest) SetUserData(v string) *RunInstancesRequest {
	s.UserData = &v
	return s
}

func (s *RunInstancesRequest) SetVSwitchId(v string) *RunInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunInstancesRequest) SetZoneId(v string) *RunInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *RunInstancesRequest) Validate() error {
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.HibernationOptions != nil {
		if err := s.HibernationOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClockOptions != nil {
		if err := s.ClockOptions.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterface != nil {
		for _, item := range s.NetworkInterface {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkOptions != nil {
		if err := s.NetworkOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateDnsNameOptions != nil {
		if err := s.PrivateDnsNameOptions.Validate(); err != nil {
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

type RunInstancesRequestCpuOptions struct {
	// The number of CPU cores.
	//
	// <props="china">Default value: For more information, see [Customize CPU options](https://help.aliyun.com/document_detail/145895.html).
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	Numa *string `json:"Numa,omitempty" xml:"Numa,omitempty"`
	// The number of threads per CPU core. The number of vCPUs of an ECS instance = `CpuOptions.Core` value × `CpuOptions.ThreadsPerCore` value.
	//
	// - `CpuOptions.ThreadsPerCore=1` indicates that hyper-threading is disabled.
	//
	// - Only specific instance types support custom CPU thread counts.
	//
	// <props="china">For the valid values and default value, see [Custom CPU options](https://help.aliyun.com/document_detail/145895.html).
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	// The CPU topology type of the instance. Valid values:
	//
	// - ContinuousCoreToHTMapping: The hyper-threads (HTs) within the same core in the CPU topology of the instance are continuous.
	//
	// - DiscreteCoreToHTMapping: The HTs within the same core in the CPU topology of the instance are discrete.
	//
	// Default value: null.
	//
	// > Only specific instance families support this parameter. For more information about the supported instance families, see [View and modify CPU topology](https://help.aliyun.com/document_detail/2636059.html).
	//
	// example:
	//
	// DiscreteCoreToHTMapping
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// enabled
	NestedVirtualization *string `json:"NestedVirtualization,omitempty" xml:"NestedVirtualization,omitempty"`
}

func (s RunInstancesRequestCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestCpuOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *RunInstancesRequestCpuOptions) GetNuma() *string {
	return s.Numa
}

func (s *RunInstancesRequestCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *RunInstancesRequestCpuOptions) GetTopologyType() *string {
	return s.TopologyType
}

func (s *RunInstancesRequestCpuOptions) GetNestedVirtualization() *string {
	return s.NestedVirtualization
}

func (s *RunInstancesRequestCpuOptions) SetCore(v int32) *RunInstancesRequestCpuOptions {
	s.Core = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetNuma(v string) *RunInstancesRequestCpuOptions {
	s.Numa = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetThreadsPerCore(v int32) *RunInstancesRequestCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetTopologyType(v string) *RunInstancesRequestCpuOptions {
	s.TopologyType = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetNestedVirtualization(v string) *RunInstancesRequestCpuOptions {
	s.NestedVirtualization = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestHibernationOptions struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
}

func (s RunInstancesRequestHibernationOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestHibernationOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestHibernationOptions) GetConfigured() *bool {
	return s.Configured
}

func (s *RunInstancesRequestHibernationOptions) SetConfigured(v bool) *RunInstancesRequestHibernationOptions {
	s.Configured = &v
	return s
}

func (s *RunInstancesRequestHibernationOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestPrivatePoolOptions struct {
	// The private pool ID, which is the ID of the elasticity assurance or capacity reservation.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The private pool options for instance startup. A private pool is generated after an elasticity assurance or capacity reservation takes effect. You can select a private pool when you start an instance. Valid values:
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s RunInstancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *RunInstancesRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *RunInstancesRequestPrivatePoolOptions) SetId(v string) *RunInstancesRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *RunInstancesRequestPrivatePoolOptions) SetMatchCriteria(v string) *RunInstancesRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *RunInstancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSchedulerOptions struct {
	// Specifies the dedicated host cluster for the ECS instance. The system automatically selects a dedicated host from the specified cluster to deploy the ECS instance.
	//
	// > This parameter takes effect only when `Tenancy` is set to `host`.
	//
	// If you specify both a dedicated host (`DedicatedHostId`) and a dedicated host cluster (`SchedulerOptions.DedicatedHostClusterId`):
	//
	// - If the dedicated host belongs to the dedicated host cluster, the ECS instance is preferentially deployed on the specified dedicated host.
	//
	// - If the dedicated host does not belong to the dedicated host cluster, the ECS instance fails to be created.
	//
	// <props="china">You can call [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) to query the list of dedicated host cluster IDs.
	//
	// <props="intl">You can call [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) to query the list of dedicated host cluster IDs.
	//
	// <props="partner">You can call [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) to query the list of dedicated host cluster IDs.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
}

func (s RunInstancesRequestSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSchedulerOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSchedulerOptions) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *RunInstancesRequestSchedulerOptions) SetDedicatedHostClusterId(v string) *RunInstancesRequestSchedulerOptions {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *RunInstancesRequestSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSecurityOptions struct {
	// The confidential computing mode. Valid value: Enclave.
	//
	// example:
	//
	// Enclave
	ConfidentialComputingMode *string `json:"ConfidentialComputingMode,omitempty" xml:"ConfidentialComputingMode,omitempty"`
	// The trusted system mode. Valid value: vTPM.
	//
	// example:
	//
	// vTPM
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
	EnableSecureBoot  *bool   `json:"EnableSecureBoot,omitempty" xml:"EnableSecureBoot,omitempty"`
}

func (s RunInstancesRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *RunInstancesRequestSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *RunInstancesRequestSecurityOptions) GetEnableSecureBoot() *bool {
	return s.EnableSecureBoot
}

func (s *RunInstancesRequestSecurityOptions) SetConfidentialComputingMode(v string) *RunInstancesRequestSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *RunInstancesRequestSecurityOptions) SetTrustedSystemMode(v string) *RunInstancesRequestSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *RunInstancesRequestSecurityOptions) SetEnableSecureBoot(v bool) *RunInstancesRequestSecurityOptions {
	s.EnableSecureBoot = &v
	return s
}

func (s *RunInstancesRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSystemDisk struct {
	// The ID of the automatic snapshot policy applied to the system disk.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The category of the system disk. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD.
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// Default value description:
	//
	// - If InstanceType is set to a retired instance type that is not I/O optimized, the default value is `cloud`.
	//
	// - In other cases, the default value is `cloud_efficiency`.<props="china">After January 30, 2026, for instance types that support only cloud_essd, the default value is changed from cloud_efficiency to cloud_essd PL0. For more information, see [Change announcement](https://www.aliyun.com/notice/117844).
	//
	// >This parameter supports the `cloud_essd_entry` value only when `InstanceType` is set to the [u1, universal instance family](https://help.aliyun.com/document_detail/457079.html) (`ecs.u1`) or the [e, economy instance family](https://help.aliyun.com/document_detail/108489.html) (`ecs.e`).
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SystemDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length and can contain characters from the Unicode letter category (including English and Chinese characters and digits). The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The performance level of the enterprise SSD used as the system disk. Settings for the performance level when you create an enterprise SSD (standard SSD not applicable). Valid values:
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// - Basic disk: 20 to 500.
	//
	// - Enterprise SSD (ESSD):
	//
	//   - PL0: 1 to 2048.
	//
	//   - PL1: 20 to 2048.
	//
	//   - PL2: 461 to 2048.
	//
	//   - PL3: 1261 to 2048.
	//
	// - ESSD AutoPL disk: 1 to 2048.
	//
	// - Other disk types: 20 to 2048.
	//
	// The value of this parameter must be greater than or equal to max{1, ImageSize}.
	//
	// Default value: max{40, size of the image specified by ImageId}.
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// - true: encrypts the system disk.
	//
	// - false: does not encrypt the system disk.
	//
	// Default value: false.
	//
	// >China (Hong Kong) Zone D and Singapore Zone A do not support encrypting the system disk when you create an instance.
	//
	// 	Notice: When you use a shared encrypted image to create a disk based on an encrypted snapshot, you must set the Encrypted parameter to true for the disk to ensure that the created disk uses the key of the account with which the image is shared.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key used for the system disk.
	//
	// > If Encrypted is set to true and KMSKeyId is not specified, the default key is used for encryption, and the KMSKeyId value is returned after the instance is created.
	//
	// > - - The disk is created from a non-shared encrypted snapshot: The encryption key used by the snapshot is used by default.
	//
	// > - - The disk is created from a shared encrypted snapshot: The service key is used by default.
	//
	// > - - The disk is created in a region where block storage account-level default encryption is enabled: The specified account-level key is used by default.
	//
	// > - - Other cases: The service key is used by default.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - baseline performance}.
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The ID of the dedicated block storage cluster. If you want to use a disk in a dedicated block storage cluster as the system disk when you create an ECS instance, specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s RunInstancesRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *RunInstancesRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *RunInstancesRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *RunInstancesRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *RunInstancesRequestSystemDisk) GetSize() *string {
	return s.Size
}

func (s *RunInstancesRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *RunInstancesRequestSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *RunInstancesRequestSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *RunInstancesRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *RunInstancesRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *RunInstancesRequestSystemDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *RunInstancesRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *RunInstancesRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetCategory(v string) *RunInstancesRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetDescription(v string) *RunInstancesRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetDiskName(v string) *RunInstancesRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetPerformanceLevel(v string) *RunInstancesRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetSize(v string) *RunInstancesRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetBurstingEnabled(v bool) *RunInstancesRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetEncryptAlgorithm(v string) *RunInstancesRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetEncrypted(v string) *RunInstancesRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetKMSKeyId(v string) *RunInstancesRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetProvisionedIops(v int64) *RunInstancesRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetStorageClusterId(v string) *RunInstancesRequestSystemDisk {
	s.StorageClusterId = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestArn struct {
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s RunInstancesRequestArn) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestArn) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *RunInstancesRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *RunInstancesRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *RunInstancesRequestArn) SetAssumeRoleFor(v int64) *RunInstancesRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *RunInstancesRequestArn) SetRoleType(v string) *RunInstancesRequestArn {
	s.RoleType = &v
	return s
}

func (s *RunInstancesRequestArn) SetRolearn(v string) *RunInstancesRequestArn {
	s.Rolearn = &v
	return s
}

func (s *RunInstancesRequestArn) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestClockOptions struct {
	// The PTP status. Valid values:
	//
	// example:
	//
	// enabled
	PtpStatus *string `json:"PtpStatus,omitempty" xml:"PtpStatus,omitempty"`
}

func (s RunInstancesRequestClockOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestClockOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestClockOptions) GetPtpStatus() *string {
	return s.PtpStatus
}

func (s *RunInstancesRequestClockOptions) SetPtpStatus(v string) *RunInstancesRequestClockOptions {
	s.PtpStatus = &v
	return s
}

func (s *RunInstancesRequestClockOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestDataDisk struct {
	// The ID of the automatic snapshot policy to apply to the data disk.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of data disk N. Valid values:
	//
	// - cloud_efficiency: ultra disk.
	//
	// - cloud_ssd: standard SSD.
	//
	// - cloud_essd: enterprise SSD.
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_regional_disk_auto: regional Enterprise SSD (ESSD).
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	//   >The `cloud_essd_entry` value is supported only when `InstanceType` is set to an instance type in the `ecs.u1` or `ecs.e` instance family.
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - Standard.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - Premium.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// Default value description:
	//
	// - If InstanceType is set to a retired instance type that is not I/O optimized, the default value is `cloud`.
	//
	// - In other cases, the default value is `cloud_efficiency`.<props="china">After January 30, 2026, if the I/O optimized instance type does not support cloud_auto, the default value is cloud_efficiency. Otherwise, the default value is cloud_auto, and performance burst is enabled by default (which incurs additional fees. For more information, see [Billing examples](~~368372#p_75k_2hp_7gp~~)). For more information, see [Change announcement](https://www.aliyun.com/notice/117844).
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the data disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// DataDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the data disk. The naming conventions for mount points vary based on the number of data disks attached:
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters, digits, and characters categorized as letter in Unicode. The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N. Valid values:
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key for the data disk.
	//
	// > If Encrypted is set to true and KMSKeyId is not specified, the default key is used for encryption, and the KMSKeyId value is returned after the instance is created.
	//
	// > - - The disk is created from a non-shared encrypted snapshot: The encryption key used by the snapshot is used by default.
	//
	// > - - The disk is created from a shared encrypted snapshot: The service key is used by default.
	//
	// > - - The disk is created in a region where block storage account-level default encryption is enabled: The specified account-level key is used by default.
	//
	// > - - Other cases: The service key is used by default.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// Settings for the performance level of the data disk when you create an enterprise SSD as a data disk. The value of N must be the same as that in `DataDisk.N.Category=cloud_essd`. Valid values:
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - baseline performance}.
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// - cloud_efficiency: 20 to 32768.
	//
	// - cloud_ssd: 20 to 32768.
	//
	// - cloud_essd: The valid values depend on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     - PL0: 1 to 65536.
	//
	//     - PL1: 20 to 65536.
	//
	//     - PL2: 461 to 65536.
	//
	//     - PL3: 1261 to 65536.
	//
	// - cloud: 5 to 2000.
	//
	// - cloud_auto: 1 to 65536.
	//
	// - cloud_essd_entry: 10 to 32768.
	//
	// >The value of this parameter must be greater than or equal to the size of the snapshot specified by `SnapshotId`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot to use to create data disk N. Valid values of N: 1 to 16.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster. If you want to use a disk in a dedicated block storage cluster as the data disk when you create an ECS instance, specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s RunInstancesRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestDataDisk) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *RunInstancesRequestDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *RunInstancesRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *RunInstancesRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *RunInstancesRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *RunInstancesRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *RunInstancesRequestDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *RunInstancesRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *RunInstancesRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *RunInstancesRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *RunInstancesRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *RunInstancesRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *RunInstancesRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *RunInstancesRequestDataDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *RunInstancesRequestDataDisk) SetAutoSnapshotPolicyId(v string) *RunInstancesRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetBurstingEnabled(v bool) *RunInstancesRequestDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetCategory(v string) *RunInstancesRequestDataDisk {
	s.Category = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDeleteWithInstance(v bool) *RunInstancesRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDescription(v string) *RunInstancesRequestDataDisk {
	s.Description = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDevice(v string) *RunInstancesRequestDataDisk {
	s.Device = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDiskName(v string) *RunInstancesRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetEncryptAlgorithm(v string) *RunInstancesRequestDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetEncrypted(v string) *RunInstancesRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetKMSKeyId(v string) *RunInstancesRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetPerformanceLevel(v string) *RunInstancesRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetProvisionedIops(v int64) *RunInstancesRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetSize(v int32) *RunInstancesRequestDataDisk {
	s.Size = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetSnapshotId(v string) *RunInstancesRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetStorageClusterId(v string) *RunInstancesRequestDataDisk {
	s.StorageClusterId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestImageOptions struct {
	// Indicates whether instances that use this image support logon as the ecs-user user. Valid values:
	//
	// - true: supported.
	//
	// - false: not supported.
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s RunInstancesRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestImageOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *RunInstancesRequestImageOptions) SetLoginAsNonRoot(v bool) *RunInstancesRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *RunInstancesRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestNetworkInterface struct {
	// Specifies whether to retain the network interface when the instance is released. Valid values:
	//
	// - true: The network interface is not retained.
	//
	// - false: The network interface is retained.
	//
	// Default value: true.
	//
	// >This parameter takes effect only on secondary ENIs.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the network interface controller (NIC).
	//
	// example:
	//
	// Network_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the network interface controller (NIC). The valid values of N cannot exceed the number of network interface controllers (NICs) supported by the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies one or more IPv6 addresses for the primary ENI. You can specify up to 10 IPv6 addresses. Valid values of the second N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::***`
	//
	// Take note of the following items:
	//
	// - This parameter takes effect only when `NetworkInterface.N.InstanceType` is set to `Primary`. If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter cannot be specified.
	//
	// - If you specify this parameter, `Amount` can only be set to 1, and you cannot specify `Ipv6AddressCount`, `Ipv6Address.N`, or `NetworkInterface.N.Ipv6AddressCount`.
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of randomly generated IPv6 addresses for the primary ENI. Valid values: 1 to 10.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The index of the physical network card specified for the ENI.
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
	// The ID of the ENI to attach to the instance.
	//
	// After you set this parameter, the value of `Amount` can only be 1.
	//
	// >This parameter takes effect only for secondary ENIs. After you specify an existing secondary ENI, you cannot configure other ENI creation parameters.
	//
	// example:
	//
	// eni-bp1gn106np8jhxhj****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the ENI. The name must be 2 to 128 characters in length and can contain characters that are categorized as letter in Unicode, including but not limited to English letters, Chinese characters, and digits. The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Take note of the following items:
	//
	// - Valid values of N cannot exceed the maximum number of ENIs supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the maximum number of ENIs supported by the target instance type.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you do not need to set this parameter.
	//
	// example:
	//
	// Network_Name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication mode of the ENI. Valid values:
	//
	// - Standard: Uses the TCP communication mode.
	//
	// - HighPerformance: Enables the Elastic RDMA Interface (ERI) and uses the RDMA communication mode.
	//
	// Default value: Standard.
	//
	// >The number of ENIs in RDMA mode cannot exceed the limit of the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// Adds a network interface controller (NIC) and sets the primary IP address.
	//
	// example:
	//
	// ``172.16.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The number of queues supported by the ENI.
	//
	// Take note of the following items:
	//
	// - The value of N cannot exceed the maximum number of ENIs supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the maximum number of ENIs supported by the target instance type.
	//
	// - The value cannot exceed the maximum number of queues per ENI allowed by the instance type.
	//
	// - The total number of queues across all ENIs on the instance cannot exceed the total queue quota allowed by the instance type. You can call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the MaximumQueueNumberPerEni and TotalEniQueueQuantity fields for the maximum number of queues per ENI and the total queue quota of the instance type.
	//
	// - If NetworkInterface.N.InstanceType is set to Primary and this parameter is specified, you cannot specify the NetworkInterfaceQueueNumber parameter.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queue pairs for the RDMA network interface.
	//
	// If you want to attach multiple RDMA network interfaces to the instance being created, we recommend that you manually specify QueuePairNumber for each network interface based on the upper limit of `QueuePairNumber` supported by the instance type and the number of network interfaces you plan to use. Ensure that the total QueuePairNumber across all network interfaces does not exceed the maximum value allowed by the instance type. Call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the upper limits of the instance type.
	//
	// 	Notice: If QueuePairNumber is not specified for an RDMA network interface, the upper limit of QueuePairNumber for all RDMA network interfaces supported by the instance type is used by default. Therefore, once an RDMA network interface without QueuePairNumber specified is attached, no more RDMA network interfaces can be added (regular network interfaces are not affected by this restriction).</notice>
	//
	// example:
	//
	// 0
	QueuePairNumber *int64 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The inbound queue depth of the network interface controller (NIC).
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The number of secondary private IPv4 addresses for the ENI. Valid values: 1 to 49.
	//
	// example:
	//
	// 10
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The ID of the security group to which the network interface controller (NIC) belongs.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// One or more security group IDs to which the ENI belongs.
	//
	// - The valid values of N for the first index do not exceed the number of ENIs supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of ENIs supported by the target instance type.
	//
	// - The second N indicates that you can specify one or more security group IDs. The valid values of N are related to the quota of security groups that an instance can join. For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// Take note of the following items:
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you must specify this parameter or `NetworkInterface.N.SecurityGroupId`. In this case, this parameter has the same effect as `SecurityGroupIds.N`, but you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupId` at the same time.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the security group to which the ECS instance belongs.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable source/destination checking. Enable this feature to improve network security. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// Default value: false.
	//
	// > Only some regions support this feature. Before you use this feature, read [Source/destination checking](https://help.aliyun.com/document_detail/2863210.html).
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The outbound queue depth of the network interface controller (NIC).
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
	// The ID of the vSwitch to which the network interface controller (NIC) belongs.
	//
	// example:
	//
	// vsw-bp67acfmxazb4p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s RunInstancesRequestNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestNetworkInterface) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *RunInstancesRequestNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequestNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunInstancesRequestNetworkInterface) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *RunInstancesRequestNetworkInterface) GetIpv6AddressCount() *int64 {
	return s.Ipv6AddressCount
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *RunInstancesRequestNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *RunInstancesRequestNetworkInterface) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *RunInstancesRequestNetworkInterface) GetQueuePairNumber() *int64 {
	return s.QueuePairNumber
}

func (s *RunInstancesRequestNetworkInterface) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *RunInstancesRequestNetworkInterface) GetSecondaryPrivateIpAddressCount() *int32 {
	return s.SecondaryPrivateIpAddressCount
}

func (s *RunInstancesRequestNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunInstancesRequestNetworkInterface) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *RunInstancesRequestNetworkInterface) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *RunInstancesRequestNetworkInterface) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *RunInstancesRequestNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunInstancesRequestNetworkInterface) SetDeleteOnRelease(v bool) *RunInstancesRequestNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetDescription(v string) *RunInstancesRequestNetworkInterface {
	s.Description = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetInstanceType(v string) *RunInstancesRequestNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetIpv6Address(v []*string) *RunInstancesRequestNetworkInterface {
	s.Ipv6Address = v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetIpv6AddressCount(v int64) *RunInstancesRequestNetworkInterface {
	s.Ipv6AddressCount = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkCardIndex(v int32) *RunInstancesRequestNetworkInterface {
	s.NetworkCardIndex = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkInterfaceId(v string) *RunInstancesRequestNetworkInterface {
	s.NetworkInterfaceId = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkInterfaceName(v string) *RunInstancesRequestNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *RunInstancesRequestNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetPrimaryIpAddress(v string) *RunInstancesRequestNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetQueueNumber(v int32) *RunInstancesRequestNetworkInterface {
	s.QueueNumber = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetQueuePairNumber(v int64) *RunInstancesRequestNetworkInterface {
	s.QueuePairNumber = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetRxQueueSize(v int32) *RunInstancesRequestNetworkInterface {
	s.RxQueueSize = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSecondaryPrivateIpAddressCount(v int32) *RunInstancesRequestNetworkInterface {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSecurityGroupId(v string) *RunInstancesRequestNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSecurityGroupIds(v []*string) *RunInstancesRequestNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSourceDestCheck(v bool) *RunInstancesRequestNetworkInterface {
	s.SourceDestCheck = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetTxQueueSize(v int32) *RunInstancesRequestNetworkInterface {
	s.TxQueueSize = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetVSwitchId(v string) *RunInstancesRequestNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestNetworkOptions struct {
	// The bandwidth weight value of the instance. The valid values vary by instance type. To query the supported bandwidth weight levels for a specific instance type, call DescribeInstanceTypes. The BandwidthWeighting field in the response indicates the supported bandwidth weight levels. You can use the name field in the returned values as the dictionary value, such as Vpc-L1 or Ebs-L1.
	//
	// example:
	//
	// Default
	BandwidthWeighting *string `json:"BandwidthWeighting,omitempty" xml:"BandwidthWeighting,omitempty"`
	// Specifies whether to enable the Jumbo Frame feature for the instance. Valid values:
	//
	// - false: does not enable Jumbo Frame. The MTU of all NICs (including the primary ENI and secondary ENIs) on the instance is set to 1500.
	//
	// - true: enables Jumbo Frame. The MTU of all NICs (including the primary ENI and secondary ENIs) on the instance is set to 8500.
	//
	// Default value: true.
	//
	// >Only some instance types of the eighth generation or later support the Jumbo Frame feature. For more information, see [ECS instance MTU](https://help.aliyun.com/document_detail/200512.html).
	//
	// example:
	//
	// false
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	EnableNetworkEncryption *bool `json:"EnableNetworkEncryption,omitempty" xml:"EnableNetworkEncryption,omitempty"`
}

func (s RunInstancesRequestNetworkOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestNetworkOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestNetworkOptions) GetBandwidthWeighting() *string {
	return s.BandwidthWeighting
}

func (s *RunInstancesRequestNetworkOptions) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *RunInstancesRequestNetworkOptions) GetEnableNetworkEncryption() *bool {
	return s.EnableNetworkEncryption
}

func (s *RunInstancesRequestNetworkOptions) SetBandwidthWeighting(v string) *RunInstancesRequestNetworkOptions {
	s.BandwidthWeighting = &v
	return s
}

func (s *RunInstancesRequestNetworkOptions) SetEnableJumboFrame(v bool) *RunInstancesRequestNetworkOptions {
	s.EnableJumboFrame = &v
	return s
}

func (s *RunInstancesRequestNetworkOptions) SetEnableNetworkEncryption(v bool) *RunInstancesRequestNetworkOptions {
	s.EnableNetworkEncryption = &v
	return s
}

func (s *RunInstancesRequestNetworkOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestPrivateDnsNameOptions struct {
	// Enables or disables DNS AAAA record resolution from the instance ID-based domain name to the IPv6 address. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableInstanceIdDnsAAAARecord *bool `json:"EnableInstanceIdDnsAAAARecord,omitempty" xml:"EnableInstanceIdDnsAAAARecord,omitempty"`
	// Specifies whether to enable DNS resolution from the instance ID-based domain name to the IPv4 address. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsARecord *bool `json:"EnableInstanceIdDnsARecord,omitempty" xml:"EnableInstanceIdDnsARecord,omitempty"`
	// Specifies whether to enable DNS resolution from the IP-based domain name to the IPv4 address. Valid values:
	//
	// example:
	//
	// true
	EnableIpDnsARecord *bool `json:"EnableIpDnsARecord,omitempty" xml:"EnableIpDnsARecord,omitempty"`
	// Specifies whether to enable reverse DNS resolution from IPv4 addresses to domain names. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableIpDnsPtrRecord *bool `json:"EnableIpDnsPtrRecord,omitempty" xml:"EnableIpDnsPtrRecord,omitempty"`
	// The hostname type. Valid values:
	//
	// example:
	//
	// Custom
	HostnameType *string `json:"HostnameType,omitempty" xml:"HostnameType,omitempty"`
}

func (s RunInstancesRequestPrivateDnsNameOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestPrivateDnsNameOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsAAAARecord() *bool {
	return s.EnableInstanceIdDnsAAAARecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsARecord() *bool {
	return s.EnableInstanceIdDnsARecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableIpDnsARecord() *bool {
	return s.EnableIpDnsARecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableIpDnsPtrRecord() *bool {
	return s.EnableIpDnsPtrRecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetHostnameType() *string {
	return s.HostnameType
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsAAAARecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsAAAARecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsARecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsARecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableIpDnsARecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableIpDnsARecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableIpDnsPtrRecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableIpDnsPtrRecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetHostnameType(v string) *RunInstancesRequestPrivateDnsNameOptions {
	s.HostnameType = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestTag struct {
	// The tag key for the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. The tag key cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value for the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunInstancesRequestTag) SetKey(v string) *RunInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *RunInstancesRequestTag) SetValue(v string) *RunInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *RunInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
