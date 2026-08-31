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
	// Specifies whether the instance on a dedicated host is associated with the dedicated host. Valid values:
	//
	// - default: The instance is not associated with the dedicated host. When an instance that has the economical mode enabled is restarted after it is stopped, if the original dedicated host has insufficient resources, the instance is placed on another dedicated host in the automatic deployment resource pool.
	//
	// - host: The instance is associated with the dedicated host. When an instance that has the economical mode enabled is restarted after it is stopped, the instance remains on the original dedicated host. If the original dedicated host has insufficient resources, the instance fails to restart.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The number of ECS instances to create. Valid values: 1 to 100.
	//
	// The number of successfully created ECS instances depends on the specified Amount and minAmount values:
	//
	// - If minAmount is not specified: Instances are created based on the Amount value. If inventory is insufficient, the API returns a failure and no instances are created.
	//
	// - If minAmount is specified:
	//
	//   - If ECS inventory < minAmount: No instances are created and the API returns a failure.
	//
	//   - If minAmount ≤ ECS inventory < Amount: Instances are created based on the available inventory and the API returns success.
	//
	//   - If ECS inventory ≥ Amount: Instances are created based on the specified Amount and the API returns success.
	//
	// Default value: 1.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// >This parameter is not publicly available.
	Arn []*RunInstancesRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to automatically complete automatic payment when you create the instance. Valid values:
	//
	// - true: automatically completes automatic payment.
	//
	//     > Make sure that your payment method has a sufficient balance. Otherwise, an abnormal order is generated and can only be canceled. If your payment method has an insufficient balance, you can set `AutoPay` to `false` to generate an unpaid order. Then, you can log on to the ECS console to pay for the order.
	//
	// - false: generates the order without completing automatic payment.
	//
	//     > If `InstanceChargeType` is set to `PostPaid`, `AutoPay` cannot be set to `false`.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The automatic release time of the pay-as-you-go instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the UTC+0 time zone. The format is `yyyy-MM-ddTHH:mm:ssZ`.
	//
	// - If the seconds (`ss`) value is not `00`, it is automatically set to the start of the current minute (`mm`).
	//
	// - The earliest release time is 30 minutes after the current time.
	//
	// - The latest release time cannot exceed three years from the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal. This parameter takes effect only when `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// - true: Enable auto-renewal.
	//
	// - false: Disable auto-renewal.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period for each renewal. Valid values:
	//
	//
	//
	// <props="china">
	//
	// - When PeriodUnit=Week: 1, 2, or 3.
	//
	// - When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, or 60.
	//
	//
	//
	// <props="intl">When PeriodUnit=Month: 1, 2, 3, 6, 12, 24, 36, 48, or 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// A client token used to ensure the idempotence of the request. Generate a unique value from your client. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length. For more information, refer to [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The clock-related property parameters of the instance.
	ClockOptions *RunInstancesRequestClockOptions `json:"ClockOptions,omitempty" xml:"ClockOptions,omitempty" type:"Struct"`
	// The running mode of the burstable instance. Valid values:
	//
	// - Standard: standard mode. For more information, see the "Performance constrained mode" section in [What are burstable instances?](https://help.aliyun.com/document_detail/59977.html).
	//
	// - Unlimited: unlimited mode. For more information, see the "Unlimited mode" section in [What are burstable instances?](https://help.aliyun.com/document_detail/59977.html).
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
	// 	Notice: Dedicated hosts do not support the creation of spot instances. If you specify `DedicatedHostId`, the `SpotStrategy` and `SpotPriceLimit` settings in the request are automatically ignored.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// Specifies whether to enable release protection for the instance. This parameter determines whether the instance can be released from the console or by calling the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation. Valid values:
	//
	// - true: enables release protection.
	//
	// - false: disables release protection.
	//
	// Default value: false.
	//
	// > This parameter is applicable only to pay-as-you-go instances. It can only restrict manual release operations but does not take effect on system-initiated release operations.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The group number of the instance in the deployment set when the deployment set uses the high availability group strategy (AvailabilityGroup). Valid values: 1 to 7.
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
	// - true: Only a dry run is performed. The system checks whether required parameters are specified, whether the request format is valid, whether business restrictions are met, and whether ECS inventory is sufficient. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - false (default): A request is sent. If the check succeeds, instances are created directly.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the instance. The following limits apply:
	//
	// - It cannot start or end with a period (.) or hyphen (-), and cannot contain consecutive periods or hyphens.
	//
	// - Windows instances: The hostname must be 2 to 15 characters in length and cannot contain periods (.) or consist entirely of digits. It can contain uppercase and lowercase letters, digits, and hyphens (-).
	//
	// - Other instances (such as Linux):
	//
	//     - The hostname must be 2 to 64 characters in length and can contain multiple periods (.). Each segment between periods can contain uppercase and lowercase letters, digits, and hyphens (-).
	//
	//     - You can use the placeholder `${instance_id}` to include the instance ID in the `HostName` parameter. For example, if `HostName=k8s-${instance_id}` and the created ECS instance ID is `i-123abc****`, the hostname is `k8s-i-123abc****`.
	//
	// When creating multiple ECS instances, you can:
	//
	// - Batch configure sequential hostnames. For more information, refer to [Batch configure sequential names or hostnames for instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// - Use the `HostNames.N` parameter to set hostnames for multiple instances individually. Note that `HostName` and `HostNames.N` cannot be set at the same time.
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
	// This parameter is required when you create Super Computing Cluster (SCC) instances. You can create an HPC cluster by referring to [CreateHpcCluster](https://help.aliyun.com/document_detail/109138.html).
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// - enabled: enables the access channel.
	//
	// - disabled: disables the access channel.
	//
	// Default value: enabled.
	//
	// >For information about instance metadata, see [Overview of ECS instance metadata](https://help.aliyun.com/document_detail/49122.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security-hardened mode (IMDSv2) to access instance metadata. Valid values:
	//
	// - optional: does not forcefully use the security-hardened mode.
	//
	// - required: forcefully uses the security-hardened mode. After you set this value, the normal mode cannot be used to access instance metadata.
	//
	// Default value: optional.
	//
	// >For information about the modes for accessing instance metadata, see [Access mode of instance metadata](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. When you set this parameter, the latest available image from the specified image family is used to create the instance.
	//
	// The name must be 2 to 128 characters in length. It cannot start with a special character, digit, http://, or https://. It can contain only the following special characters: periods (.), underscores (_), hyphens (-), and colons (:).
	//
	// Note the following items:
	//
	// - If you set `ImageId`, you cannot set this parameter.
	//
	// - If you do not set `ImageId`, but the launch template specified by `LaunchTemplateId` or `LaunchTemplateName` has `ImageId` configured, you cannot set this parameter.
	//
	// - If you do not set `ImageId`, and the launch template specified by `LaunchTemplateId` or `LaunchTemplateName` does not have `ImageId` configured, you can set this parameter.
	//
	// - If you do not set `ImageId` and do not set `LaunchTemplateId` or `LaunchTemplateName`, you can set this parameter.
	//
	// > For information about image families associated with Alibaba Cloud public images, refer to [Public image overview](https://help.aliyun.com/document_detail/108393.html).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The image ID. Specifies the image resource used to start the instance. You can call [DescribeImages](https://help.aliyun.com/document_detail/25534.html) to query available image resources. If you do not specify `LaunchTemplateId` or `LaunchTemplateName` to use a launch template, and do not specify `ImageFamily` to use the latest available image from an image family, `ImageId` is required.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image-related property information.
	ImageOptions *RunInstancesRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// Default value: PostPaid.
	//
	// <props="china">If you select subscription, make sure that your account supports balance payment or credit payment. Otherwise, the error `InvalidPayMethod` is returned.
	//
	// <props="intl">If you select subscription, make sure that your account supports credit payment. Otherwise, the error `InvalidPayMethod` is returned.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and can contain characters from the Unicode letter category (including English letters, Chinese characters, and digits). It can also contain colons (:), underscores (_), periods (.), or hyphens (-). The default value is the `InstanceId` of the instance.
	//
	// When creating multiple ECS instances, you can batch configure sequential instance names that can contain brackets ([]) and commas (,). For more information, refer to [Batch configure sequential names or hostnames for instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. If you do not specify `LaunchTemplateId` or `LaunchTemplateName` to use a launch template, `InstanceType` is required.
	//
	// - Product selection: Refer to [Instance families](https://help.aliyun.com/document_detail/25378.html) or invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query performance data of the target instance type. You can also refer to [Best practices for instance type selection](https://help.aliyun.com/document_detail/58291.html) to learn how to select an instance type from the appropriate instance family.
	//
	// - Inventory query: Invoke [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) to query active resource availability in a specific region or zone. Use the relevant parameters to filter results.
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
	// > In **pay-by-traffic*	- mode, the peak inbound and outbound bandwidths are upper limits and are not guaranteed. When resource contention occurs, the peak bandwidth may be throttled. If your workloads require guaranteed bandwidth, use **pay-by-bandwidth*	- mode.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth, in Mbit/s. Valid values:
	//
	// - If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s: 1 to 10. Default value: 10.
	//
	// - If the purchased outbound public bandwidth is greater than 10 Mbit/s: 1 to the value of `InternetMaxBandwidthOut`. Default value: the value of `InternetMaxBandwidthOut`.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth, in Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. The default value for [retired instance types](https://help.aliyun.com/document_detail/55263.html) is none. The default value for other instance types is optimized. Valid values:
	//
	// - none: The instance is not I/O optimized.
	//
	// - optimized: The instance is I/O optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// Specifies one or more IPv6 addresses for the primary ENI. You can specify up to 10 IPv6 addresses. Valid values of N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::***`.
	//
	// Note the following items:
	//
	// - If you set `Ipv6Address.N`, the value of `Amount` can only be 1, and you cannot set `Ipv6AddressCount` at the same time.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `Ipv6Addresses.N` or `Ipv6AddressCount`. Instead, set `NetworkInterface.N.Ipv6Addresses.N` or `NetworkInterface.N.Ipv6AddressCount`.
	//
	// example:
	//
	// Ipv6Address.1=2001:db8:1234:1a00::***
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of randomly generated IPv6 addresses to assign to the primary ENI. Valid values: 1 to 10.
	//
	//
	//
	// Take note of the following items:
	//
	// - You cannot set both `Ipv6Address.N` and `Ipv6AddressCount`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `Ipv6Address.N` or `Ipv6AddressCount`. You can only set `NetworkInterface.N.Ipv6Address.N` or `NetworkInterface.N.Ipv6AddressCount`.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// >This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the key pair.
	//
	// >For Windows instances, this parameter is ignored. The default value is empty. Even if you specify this parameter, only the `Password` content is used.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The ID of the launch template. For more information, call [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html).
	//
	// When you use a launch template to create instances, you must specify either `LaunchTemplateId` or `LaunchTemplateName` to determine the launch template.
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template.
	//
	// When you use a launch template to create instances, you must specify either `LaunchTemplateId` or `LaunchTemplateName` to determine the launch template.
	//
	// example:
	//
	// LaunchTemplate_Name
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The version of the launch template. If you specify `LaunchTemplateId` or `LaunchTemplateName` without specifying the launch template version, the default version is used.
	//
	// example:
	//
	// 3
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The minimum number of ECS instances to purchase. Valid values: 1 to 100.
	//
	// The number of successfully created ECS instances depends on the specified Amount and minAmount values:
	//
	// - If minAmount is not specified: Instances are created based on the Amount value. If inventory is insufficient, the API returns a failure and no instances are created.
	//
	// - If minAmount is specified:
	//
	//   - If ECS inventory < minAmount: No instances are created and the API returns a failure.
	//
	//   - If minAmount ≤ ECS inventory < Amount: Instances are created based on the available inventory and the API returns success.
	//
	//   - If ECS inventory ≥ Amount: Instances are created based on the specified Amount and the API returns success.
	//
	// example:
	//
	// 2
	MinAmount *int32 `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	// The network interface controller (NIC) information.
	NetworkInterface []*RunInstancesRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
	// The number of queues supported by the primary ENI. Take note of the following items:
	//
	// - The value cannot exceed the maximum number of queues per ENI allowed for the instance type.
	//
	// - The total number of queues for all ENIs on the instance cannot exceed the queue quota allowed for the instance type. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the `MaximumQueueNumberPerEni` and `TotalEniQueueQuantity` fields for the maximum queue number per ENI and the total queue quota.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `NetworkInterfaceQueueNumber`. You can only set `NetworkInterface.N.QueueNumber`.
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
	// ```
	//
	// ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// ```
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify `Password`, use HTTPS to send the request to avoid password leaks.
	//
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. Valid values:
	//
	// - true: Use the preset password.
	//
	// - false: Do not use the preset password.
	//
	// Default value: false.
	//
	// > When you use this parameter, the Password parameter must be empty. Make sure that the image you use has a password configured.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the resource. The unit is specified by `PeriodUnit`. This parameter takes effect and is required only when `InstanceChargeType` is set to `PrePaid`. If `DedicatedHostId` is specified, the value cannot exceed the subscription duration of the dedicated host. Valid values:
	//
	// <props="china">
	//
	// - When PeriodUnit=Week: 1, 2, 3, or 4.
	//
	// - When PeriodUnit=Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, or 60.
	//
	//
	//
	// <props="intl">When PeriodUnit=Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, or 60.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// <props="china">
	//
	// - Week.
	//
	// - Month (default).
	//
	//
	//
	// <props="intl">Month (default).
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private domain name configuration of the instance.
	//
	//
	// For more information about private private domain resolution, see [ECS private private domain resolution](https://help.aliyun.com/document_detail/2844797.html).
	PrivateDnsNameOptions *RunInstancesRequestPrivateDnsNameOptions `json:"PrivateDnsNameOptions,omitempty" xml:"PrivateDnsNameOptions,omitempty" type:"Struct"`
	// The private IP address of the instance. When you specify system reserved IP address for a VPC-type ECS instance, the IP address must be from the idle CIDR block of the vSwitch (`VSwitchId`).
	//
	// Take note of the following items:
	//
	// - After you set `PrivateIpAddress`:
	//
	//     - If `Amount` is set to 1, system reserved IP address is assigned to the created ECS instance.
	//
	//     - If `Amount` is set to a value greater than 1, consecutive private IP addresses are assigned to the instances in a batch creation, starting from the specified private IP address. In this case, you cannot attach secondary ENIs to the instances (that is, you cannot set `NetworkInterface.N.*` parameters).
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `PrivateIpAddress`. You can only set `NetworkInterface.N.PrimaryIpAddress`.
	//
	// >The first and last three IP addresses of each vSwitch CIDR block are reserved by the system and cannot be specified.
	//
	// For example, if the vSwitch CIDR block is 192.168.1.0/24, the IP addresses 192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255 are reserved by the system.
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
	// - Active: Enable security hardening. This value is applicable only to public images.
	//
	// - Deactive: Disable security hardening. This value is applicable to all image types.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which the new instance belongs. Instances in the same security group can communicate with each other. The maximum number of instances that a security group can contain varies based on the security group type. For more information, refer to the security group section in [Limits](~~25412#SecurityGroupQuota~~).
	//
	// > `SecurityGroupId` determines the network type of the instance. For example, if the specified security group is of the VPC type, the instance is a VPC-type instance, and you must also specify `VSwitchId`.
	//
	// If you do not set `LaunchTemplateId` or `LaunchTemplateName` to use a launch template, the security group ID is required. Note the following items:
	//
	// - You can set one security group by using `SecurityGroupId`, or set one or more security groups by using `SecurityGroupIds.N`. However, you cannot set both `SecurityGroupId` and `SecurityGroupIds.N` at the same time.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `SecurityGroupId` or `SecurityGroupIds.N`. You can only set `NetworkInterface.N.SecurityGroupId` or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// Adds the instance to multiple security groups. The valid values of N depend on the maximum number of security groups to which an instance can belong. For more information, see [Security group limits](https://help.aliyun.com/document_detail/101348.html).
	//
	// Note the following items:
	//
	// - You cannot specify both `SecurityGroupId` and `SecurityGroupIds.N`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `SecurityGroupId` or `SecurityGroupIds.N`. Instead, specify `NetworkInterface.N.SecurityGroupId` or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The protection period of the spot instance, in hours. Valid values:
	//
	// - 1: After the instance is created, Alibaba Cloud guarantees that the instance will not be automatically released for 1 hour. After 1 hour, the system compares the bid price with the marketplace price in real-time and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// - 0: After the instance is created, Alibaba Cloud does not guarantee a runtime. The system compares the bid price with the marketplace price in real-time and checks resource inventory to determine whether to retain or revoke the instance.
	//
	// Default value: 1.
	//
	// >
	//
	// > - This parameter currently supports only the values 0 and 1.
	//
	// > - Spot instances are billed by second. Select an appropriate protection period based on the execution duration of your tasks.
	//
	// > - Alibaba Cloud sends a notification through an ECS system event 5 minutes before the instance is revoked.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the spot instance. Valid values:
	//
	// - Terminate: The instance is directly released.
	//
	// - Stop: The instance enters economical mode.
	//
	//   For more information about economical mode, refer to [Economical mode for pay-as-you-go instances](https://help.aliyun.com/document_detail/63353.html).
	//
	// Default value: Terminate.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the instance. This parameter supports up to three decimal places and takes effect when `SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter takes effect when `InstanceChargeType` is set to `PostPaid`. Valid values:
	//
	// - NoSpot: regular pay-as-you-go instance.
	//
	// - SpotWithPriceLimit: spot instance with a maximum price limit.
	//
	// - SpotAsPriceGo: spot instance priced at the market price at the time of purchase.
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
	// The tags for the instance, disks, and primary ENI.
	Tag []*RunInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// - default: creates a non-dedicated-host instance.
	//
	// - host: creates an instance on a dedicated host. If you do not specify `DedicatedHostId`, Alibaba Cloud automatically selects a dedicated host for the instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// Specifies whether to automatically append sequential suffixes to `HostName` and `InstanceName` when creating multiple instances. Sequential suffixes start from 001 and cannot exceed 999. Valid values:
	//
	// - true: Append sequential suffixes.
	//
	// - false: Do not append sequential suffixes.
	//
	// Default value: false.
	//
	// When `HostName` or `InstanceName` is set in a specified sequential format without the `name_suffix` suffix (that is, the format is `name_prefix[begin_number,bits]`), `UniqueSuffix` does not take effect, and names are ordered only based on the specified sequence.
	//
	// For more information, refer to [Batch configure sequential names or hostnames for instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// true
	UniqueSuffix *bool `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
	// The user data of the instance. The data must be Base64-encoded. The maximum size of the raw data before Base64 encoding is 32 KB.
	//
	// For more information about usage limits, formats, and execution frequency of instance user data, refer to [Instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// >To ensure the security of UserData during transmission, avoid passing sensitive data such as passwords and private keys in plaintext. If you need to pass such information, encrypt it first, encode it in Base64, and then decrypt it inside the instance.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The vSwitch ID. If you are creating a VPC-type ECS instance, you must specify a vSwitch ID. The security group and the vSwitch must belong to the same VPC. You can call [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) to query created vSwitches.
	//
	// Note the following items:
	//
	// - If you set `VSwitchId`, the `ZoneId` parameter must match the zone of the vSwitch. You can also leave `ZoneId` unspecified, and the system automatically selects the zone of the specified vSwitch.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `VSwitchId`. You can only set `NetworkInterface.N.VSwitchId`.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the instance. You can call [DescribeZones](https://help.aliyun.com/document_detail/25610.html) to query available zones.
	//
	// > If you specify `VSwitchId`, the specified `ZoneId` must match the zone of the vSwitch. You can also leave `ZoneId` unspecified, and the system automatically selects the zone of the specified vSwitch.
	//
	// Default value: automatically selected by the system.
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
	// The number of threads per CPU core. The number of vCPUs of the ECS instance = `CpuOptions.Core` value × `CpuOptions.ThreadsPerCore` value.
	//
	// - `CpuOptions.ThreadsPerCore=1` indicates that CPU hyper-threading is disabled.
	//
	// - Only specific instance types support setting the number of threads per CPU core.
	//
	// <props="china">For information about valid values and default values, see [Customize CPU options](https://help.aliyun.com/document_detail/145895.html).
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	// The CPU topology type of the instance. Valid values:
	//
	// - ContinuousCoreToHTMapping: The hyper-threads (HTs) within the same core of the instance CPU topology are continuous.
	//
	// - DiscreteCoreToHTMapping: The HTs within the same core of the instance are discrete.
	//
	// Default value: null.
	//
	// >Only specific instance families support this parameter. For information about supported instance families, see [View and modify the CPU topology structure](https://help.aliyun.com/document_detail/2636059.html).
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
	// >This parameter is in invitational preview and is not publicly available.
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
	// The private pool option for launching the instance. After an elasticity assurance or capacity reservation takes effect, a private pool is generated for the instance to select during launch. Valid values:
	//
	// - Open: open mode. The system automatically matches available open private pool capacity. If no matching private pool capacity is available, public pool resources are used to launch the instance. In this mode, you do not need to set `PrivatePoolOptions.Id`.
	//
	// - Target: specified mode. The instance is launched by using the capacity of the specified private pool. If the specified private pool capacity is unavailable, the instance fails to launch. In this mode, you must specify the private pool ID, that is, `PrivatePoolOptions.Id` is required.
	//
	// - None: none mode. The instance does not use private pool capacity for launch.
	//
	// Default value: None.
	//
	// In any of the following scenarios, the private pool option can only be set to `None` or left empty:
	//
	// - Creating a spot instance.
	//
	// - Creating an ECS instance on a dedicated host.
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
	// Specifies the dedicated host cluster to which the ECS instance belongs. The system automatically selects a dedicated host in the cluster to deploy the ECS instance.
	//
	// > This parameter takes effect only when `Tenancy` is set to `host`.
	//
	// If you specify both a dedicated host (`DedicatedHostId`) and a dedicated host cluster (`SchedulerOptions.DedicatedHostClusterId`):
	//
	// - If the dedicated host belongs to the dedicated host cluster, the ECS instance is preferentially deployed on the specified dedicated host.
	//
	// - If the dedicated host does not belong to the dedicated host cluster, the ECS instance fails to be created.
	//
	// <props="china">You can call the [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) operation to query the list of dedicated host cluster IDs.
	//
	// <props="intl">You can call the [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) operation to query the list of dedicated host cluster IDs.
	//
	// <props="partner">You can call the [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) operation to query the list of dedicated host cluster IDs.
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
	// The confidential computing mode. Set the value to Enclave.
	//
	// When this parameter is set to Enclave, the ECS instance uses Enclave to build a confidential computing environment. Currently, only instance families c7, g7, and r7 support specifying this parameter when you call `RunInstances` to use Enclave confidential computing. Take note of the following items:
	//
	// - The confidential computing feature is in invitational preview.
	//
	// - When you create an ECS instance with Enclave confidential computing by calling an OpenAPI operation, you can only call `RunInstances`. `CreateInstance` does not support the `SecurityOptions.ConfidentialComputingMode` parameter.
	//
	// - Enclave confidential computing relies on the trusted system (vTPM). When you specify that an ECS instance uses Enclave to build a confidential computing environment, the trusted system is also enabled for the instance. Therefore, when you call this operation, if you set `SecurityOptions.ConfidentialComputingMode=Enclave`, the created ECS instance has both Enclave confidential computing mode and the trusted system enabled, regardless of whether you set `SecurityOptions.TrustedSystemMode=vTPM`.
	//
	// For more information about confidential computing, see [Build a confidential computing environment by using Enclave](https://help.aliyun.com/document_detail/203433.html).
	//
	// example:
	//
	// Enclave
	ConfidentialComputingMode *string `json:"ConfidentialComputingMode,omitempty" xml:"ConfidentialComputingMode,omitempty"`
	// The trusted system mode. Set the value to vTPM.
	//
	// The following instance families support the trusted system mode:
	//
	// - g7, c7, and r7.
	//
	// - Security-enhanced instance families (g7t, c7t, and r7t).
	//
	// When you create instances of the preceding instance families, you must set this parameter. Take note of the following items:
	//
	// - To use Alibaba Cloud Trusted System, set this parameter to vTPM. Then, Alibaba Cloud Trusted System performs trusted verification when the instance starts.
	//
	// - If you do not want to use Alibaba Cloud Trusted System, you can leave this parameter empty. However, if the ECS instance that you create uses the Enclave confidential computing mode (`SecurityOptions.ConfidentialComputingMode=Enclave`), the trusted system is also enabled for the instance.
	//
	// - When you create a trusted ECS instance by calling an OpenAPI operation, you can only call `RunInstances`. `CreateInstance` does not support the `SecurityOptions.TrustedSystemMode` parameter.
	//
	// >If you specify the instance as a trusted instance during creation, you can only use images that support the trusted system when you replace the system disk.
	//
	// For more information about the trusted system, see [Overview of the trusted feature for security-enhanced instance families](https://help.aliyun.com/document_detail/201394.html).
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
	// - cloud_essd: enterprise SSD (ESSD).
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	// Default value description:
	//
	// - If InstanceType is a retired instance type that is not I/O optimized, the default value is `cloud`.
	//
	// - In other cases, the default value is `cloud_efficiency`.<props="china"> After January 30, 2026, for instance types that support only cloud_essd, the default value changes from cloud_efficiency to cloud_essd PL0. For more information, refer to [Change notice](https://www.aliyun.com/notice/117844).
	//
	// >This parameter supports the value `cloud_essd_entry` only when `InstanceType` is set to the [u1, universal instance family](https://help.aliyun.com/document_detail/457079.html) (`ecs.u1`) or the [e, economy instance family](https://help.aliyun.com/document_detail/108489.html) (`ecs.e`).
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
	// The name of the system disk. The name must be 2 to 128 characters in length and can contain characters from the Unicode letter category (including English letters, Chinese characters, and digits). It can also contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The performance level of the enterprise SSD used as the system disk. This parameter takes effect only when you create an enterprise SSD as the system disk. Valid values:
	//
	// - PL0: A single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): A single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, refer to [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk, in GiB. Valid values:
	//
	// - Basic disk: 20 to 500.
	//
	// - Enterprise SSD:
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
	// Default value: max{40, size of the image specified by the ImageId parameter}.
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// - true: enables the performance burst feature.
	//
	// - false: does not enable the performance burst feature.
	//
	// >This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// >This parameter is not publicly available.
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
	// >Hong Kong (China) Zone D and Singapore Zone A do not support system disk encryption during instance creation.
	//
	// 	Notice: When you use a shared encrypted image to create a disk based on an encrypted snapshot, you must specify the request parameter Encrypted=true to ensure that the created disk uses the key of the image recipient.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key for the system disk.
	//
	// > If Encrypted is set to true and KMSKeyId is not specified, the default key is used for encryption. The KMSKeyId value is returned after the instance is created.
	//
	// > - - If the disk is created from a non-shared encrypted snapshot: The encryption key used by the snapshot is used by default.
	//
	// > - - If the disk is created from a shared encrypted snapshot: The service key is used by default.
	//
	// > - - If the disk is created in a region where account-level default encryption for block storage is enabled: The specified account-level key is used by default.
	//
	// > - - In other cases: The service key is used by default.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - Baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >This parameter is supported only when `SystemDisk.Category` is set to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The ID of the dedicated block storage cluster. If you want to use a disk in a dedicated block storage cluster as the system disk when you create an ECS instance, set this parameter.
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
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// >This parameter is not publicly available.
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
	// - enabled: enables PTP.
	//
	// - disabled: disables PTP.
	//
	// Default value: disabled.
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
	// - true: enables the performance burst feature.
	//
	// - false: does not enable the performance burst feature.
	//
	// >This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
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
	// - cloud_essd: enterprise SSD (ESSD).
	//
	// - cloud: basic disk.
	//
	// - cloud_auto: ESSD AutoPL disk.
	//
	// - cloud_regional_disk_auto: regional ESSD.
	//
	// - cloud_essd_entry: ESSD Entry disk.
	//
	//   >The `cloud_essd_entry` value is supported only when `InstanceType` is set to an instance type in the `ecs.u1` or `ecs.e` instance family.
	//
	// - elastic_ephemeral_disk_standard: elastic ephemeral disk - Standard Edition.
	//
	// - elastic_ephemeral_disk_premium: elastic ephemeral disk - Premium Edition.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// Default value description:
	//
	// - If InstanceType is a retired instance type that is non-I/O optimized, the default value is `cloud`.
	//
	// - In other cases, the default value is `cloud_efficiency`.<props="china">After January 30, 2026, if the I/O optimized instance type does not support cloud_auto, the default value is cloud_efficiency. Otherwise, the default value is cloud_auto, and performance burst is enabled by default (which incurs additional fees. For more information, see [Billing examples](~~368372#p_75k_2hp_7gp~~)). For more information, see [Change notice](https://www.aliyun.com/notice/117844).
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the data disk when the instance is released. Valid values:
	//
	// - true: releases the data disk when the instance is released.
	//
	// - false: does not release the data disk when the instance is released.
	//
	// Default value: true.
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
	// - 1 to 25 data disks: /dev/xvd`[b-z]`
	//
	// - More than 25 data disks: /dev/xvd`[aa-zz]`. For example, the 26th data disk is named /dev/xvdaa, the 27th data disk is named /dev/xvdab, and so on.
	//
	// > - This parameter is applicable only to full image (system image) scenarios. You can set this parameter to the mount point of a data disk in the full image and modify the corresponding `DataDisk.N.Size` and `DataDisk.N.Category` parameters to change the disk type and size of the data disk in the full image.
	//
	// > - When you use a full image to create an instance, the data disks in the full image are created as the first 1 to n data disks of the ECS instance.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of the data disk. The name must be 2 to 128 characters in length and can contain letters, digits, and characters that are supported by Unicode in the letter category. The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N. Valid values:
	//
	// - true: encrypts the data disk.
	//
	// - false: does not encrypt the data disk.
	//
	// Default value: false.
	//
	//
	// 	Notice: When you use a shared encrypted image to create a disk based on an encrypted snapshot, you must specify the request parameter Encrypted=true to ensure that the created disk uses the key of the image recipient.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key for the data disk.
	//
	// > If Encrypted is set to true and KMSKeyId is not specified, the default key is used for encryption. The KMSKeyId value is returned after the instance is created.
	//
	// > - - If the disk is created from a non-shared encrypted snapshot: The encryption key used by the snapshot is used by default.
	//
	// > - - If the disk is created from a shared encrypted snapshot: The service key is used by default.
	//
	// > - - If the disk is created in a region where account-level default encryption for block storage is enabled: The specified account-level key is used by default.
	//
	// > - - In other cases: The service key is used by default.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// Settings the performance level of the data disk when you create an enterprise SSD as a data disk. The value of N must be consistent with the N in `DataDisk.N.Category=cloud_essd`. Valid values:
	//
	// - PL0: A single disk can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1 (default): A single disk can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about how to select an ESSD performance level, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1000 × Capacity - Baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
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
	// - cloud_essd: The valid value range depends on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     - PL0: 1 to 65,536.
	//
	//     - PL1: 20 to 65,536.
	//
	//     - PL2: 461 to 65,536.
	//
	//     - PL3: 1261 to 65,536.
	//
	// - cloud: 5 to 2000.
	//
	// - cloud_auto: 1 to 65,536.
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
	// After you specify `DataDisk.N.SnapshotId`, `DataDisk.N.Size` is ignored and the disk is created with the size of the specified snapshot. Snapshots created on or before July 15, 2013 cannot be used. Requests that use such snapshots are rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster. If you want to use a disk in a dedicated block storage cluster as the data disk when you create an ECS instance, set this parameter.
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
	// Specifies whether the instance that uses this image supports logon with the ecs-user user. Valid values:
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
	// Specifies whether to retain the ENI when the instance is released. Valid values:
	//
	// - true: does not retain the ENI.
	//
	// - false: retains the ENI.
	//
	// Default value: true.
	//
	// >This parameter takes effect only for secondary ENIs.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the ENI.
	//
	// Note the following items:
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// - The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you do not need to set this parameter.
	//
	// example:
	//
	// Network_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the ENI. The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// Valid values:
	//
	// - Primary: primary ENI.
	//
	// - Secondary: secondary ENI.
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies one or more IPv6 addresses for the primary ENI. You can specify up to 10 IPv6 addresses. Valid values of the second N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::***`
	//
	// Note the following items:
	//
	// - This parameter takes effect only when `NetworkInterface.N.InstanceType` is set to `Primary`. If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, you cannot set this parameter.
	//
	// - After you set this parameter, the value of `Amount` can only be 1, and you cannot set `Ipv6AddressCount`, `Ipv6Address.N`, or `NetworkInterface.N.Ipv6AddressCount`.
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of randomly generated IPv6 addresses for the primary ENI. Valid values: 1 to 10.
	//
	// Note the following items:
	//
	// - This parameter takes effect only when `NetworkInterface.N.InstanceType` is set to `Primary`. If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, you cannot set this parameter.
	//
	// - After you set this parameter, you cannot set `Ipv6AddressCount`, `Ipv6Address.N`, or `NetworkInterface.N.Ipv6Address.N`.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The index of the physical network card specified for the network interface controller (NIC).
	//
	// Note the following items:
	//
	// - Only specific instance types support specifying a physical network card index.
	//
	// - If NetworkInterface.N.InstanceType is set to Primary, for instance types that support physical network cards, this parameter can only be set to 0.
	//
	// - If NetworkInterface.N.InstanceType is set to Secondary or left empty, for instance types that support physical network cards, this parameter can be set based on the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
	// The ID of the ENI to attach to the instance.
	//
	// After you set this parameter, the value of `Amount` can only be 1.
	//
	// >This parameter takes effect only for secondary ENIs. After you specify an existing secondary ENI, you cannot configure other network interface controller (NIC) creation parameters.
	//
	// example:
	//
	// eni-bp1gn106np8jhxhj****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of the ENI. The name must be 2 to 128 characters in length and can contain letters, digits, and characters that are supported by Unicode in the letter categorization. The name can contain colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// Note the following items:
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you do not need to set this parameter.
	//
	// example:
	//
	// Network_Name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication mode of the network interface controller (NIC). Valid values:
	//
	// - Standard: uses the TCP communication mode.
	//
	// - HighPerformance: enables the Elastic RDMA Interface (ERI) and uses the RDMA communication mode.
	//
	// Default value: Standard.
	//
	// >The number of Elastic Network Interfaces (ENIs) in RDMA mode cannot exceed the limit of the instance family. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// Adds an ENI and sets the primary IP address.
	//
	// Note the following items:
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	//     - When you set one ENI, you can set one primary ENI or one secondary ENI. If the value of `Amount` is greater than 1 and you set the primary ENI with this parameter specified, consecutive primary IP addresses starting from the specified IP address are allocated to multiple ECS instances during batch creation. In this case, you cannot attach secondary ENIs to the instances.
	//
	//     - If the value of `Amount` is greater than 1 and this parameter is set for the primary ENI, you cannot set a secondary ENI (that is, you cannot set `NetworkInterface.2.InstanceType=Secondary`).
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, this parameter has the same effect as `PrivateIpAddress`, but you cannot specify the `PrivateIpAddress` parameter at the same time.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter specifies the primary IP address of the secondary ENI. By default, an IP address is randomly selected from the CIDR block of the vSwitch to which the ENI belongs.
	//
	// >- The first and last three IP addresses of each vSwitch CIDR block are system reserved IP addresses and cannot be specified.
	//
	// For example, if the CIDR block of the vSwitch is 192.168.1.0/24, the IP addresses 192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255 are system reserved IP addresses.
	//
	// example:
	//
	// ``172.16.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The number of queues for the ENI.
	//
	// Note the following items:
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// - The value cannot exceed the maximum number of queues per ENI allowed by the instance type.
	//
	// - The total number of queues for all ENIs on the instance cannot exceed the queue quota allowed by the instance type. You can call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the `MaximumQueueNumberPerEni` and `TotalEniQueueQuantity` fields for the maximum number of queues per ENI and the total quota.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary` and this parameter is set, you cannot set the `NetworkInterfaceQueueNumber` parameter.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queues for the RDMA ENI.
	//
	// If you want to attach multiple RDMA ENIs to the instance, we recommend that you manually specify QueuePairNumber for each ENI based on the upper limit of QueuePairNumber supported by the instance type and the number of ENIs you plan to use. Make sure that the total QueuePairNumber of all ENIs does not exceed the maximum value allowed by the instance type. Call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the upper limit of the instance type.
	//
	// 	Notice: If QueuePairNumber is not specified for an RDMA ENI, the upper limit of QueuePairNumber supported by the instance type is used by default. Therefore, after you attach one RDMA ENI without specifying QueuePairNumber, you cannot attach more RDMA ENIs (regular ENIs are not affected by this limit).
	//
	// example:
	//
	// 0
	QueuePairNumber *int64 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The inbound queue depth of the network interface controller (NIC).
	//
	//
	// <props="china">
	//
	// >This parameter is in invitational preview and is not publicly available. If you need to use this feature, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex) to request access.
	//
	//
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is in invitational preview and is not publicly available. If you need to use this feature, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to request access.
	//
	//
	//
	// Note the following items when you use this parameter:
	//
	// - This parameter is applicable only to seventh-generation and later ECS instance types.
	//
	// - This parameter is currently applicable only to Linux images.
	//
	// - A larger inbound queue depth can improve inbound throughput and reduce packet loss, but consumes more memory.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The number of secondary private IPv4 addresses to allocate to the network interface controller (NIC). Valid values: 1 to 49.
	//
	// - The value cannot exceed the IP address limit for the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// - `NetworkInterface.N.SecondaryPrivateIpAddressCount` specifies the number of secondary private IPv4 addresses to allocate to the network interface controller (NIC) (excluding the primary private IP address of the NIC). The system randomly allocates IP addresses from the available CIDR block of the vSwitch (`NetworkInterface.N.VSwitchId`) to which the network interface controller (NIC) belongs.
	//
	// example:
	//
	// 10
	SecondaryPrivateIpAddressCount *int32 `json:"SecondaryPrivateIpAddressCount,omitempty" xml:"SecondaryPrivateIpAddressCount,omitempty"`
	// The ID of the security group to which the ENI belongs.
	//
	// Note the following items:
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you must set this parameter. In this case, this parameter has the same effect as `SecurityGroupId`, but you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the security group of the ECS instance.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of one or more security groups to which the ENI belongs.
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// - The second N indicates that you can specify one or more security group IDs. The valid values of N depend on the maximum number of security groups to which an instance can belong. For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// Note the following items:
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you must set this parameter or `NetworkInterface.N.SecurityGroupId`. In this case, this parameter has the same effect as `SecurityGroupIds.N`, but you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupId`.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the security group of the ECS instance.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable source/destination checking. We recommend that you enable this feature to improve network security. Valid values:
	//
	// - true: enables source/destination checking.
	//
	// - false: disables source/destination checking.
	//
	// Default value: false.
	//
	// > This feature is supported only in specific regions. Before you use this feature, read [Source/destination checking](https://help.aliyun.com/document_detail/2863210.html).
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The outbound queue depth of the network interface controller (NIC).
	//
	//
	// <props="china">
	//
	// >This parameter is in invitational preview and is not publicly available. If you need to use this feature, [submit a ticket](https://selfservice.console.aliyun.com/ticket/createIndex) to request access.
	//
	//
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is in invitational preview and is not publicly available. If you need to use this feature, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl) to request access.
	//
	//
	//
	// Note the following items when you use this parameter:
	//
	// - This parameter is applicable only to seventh-generation and later ECS instance types.
	//
	// - This parameter is currently applicable only to Linux images.
	//
	// - A larger outbound queue depth can improve outbound throughput and reduce packet loss, but consumes more memory.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
	// The ID of the vSwitch to which the ENI belongs.
	//
	// Note the following items:
	//
	// - The valid values of N do not exceed the number of network interface controllers (NICs) supported by the instance type. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html) or call [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) to query the number of network interface controllers (NICs) supported by the target instance type.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Primary`, you must set this parameter. In this case, this parameter has the same effect as `VSwitchId`, but you cannot specify the `VSwitchId` parameter at the same time.
	//
	// - If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, this parameter is optional. Default value: the vSwitch to which the ECS instance belongs.
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
	// The bandwidth weight value of the instance. Different instance types support different value ranges. You can call DescribeInstanceTypes to query the supported bandwidth weight tiers for a specific instance type. The returned BandwidthWeighting field indicates the supported bandwidth weight tiers for that instance type. The dictionary value can be the name field in the returned values, such as Vpc-L1 or Ebs-L1.
	//
	// example:
	//
	// Default
	BandwidthWeighting *string `json:"BandwidthWeighting,omitempty" xml:"BandwidthWeighting,omitempty"`
	// Specifies whether to enable the Jumbo frame feature for the instance. Valid values:
	//
	// - false: disables Jumbo frame. The MTU of all ENIs (including the primary ENI and secondary ENIs) on the instance is set to 1500.
	//
	// - true: enables Jumbo frame. The MTU of all ENIs (including the primary ENI and secondary ENIs) on the instance is set to 8500.
	//
	// Default value: true.
	//
	// >Only some instance types of the eighth generation and later support the Jumbo frame feature. For more information, see [ECS instance MTU](https://help.aliyun.com/document_detail/200512.html).
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
	// Specifies whether to enable DNS resolution from the instance ID-based domain name to the IPv6 address. Valid values:
	//
	// - true: enables the resolution.
	//
	// - false: disables the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableInstanceIdDnsAAAARecord *bool `json:"EnableInstanceIdDnsAAAARecord,omitempty" xml:"EnableInstanceIdDnsAAAARecord,omitempty"`
	// Specifies whether to enable DNS resolution from the instance ID-based domain name to the IPv4 address. Valid values:
	//
	// - true: enables the resolution.
	//
	// - false: disables the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsARecord *bool `json:"EnableInstanceIdDnsARecord,omitempty" xml:"EnableInstanceIdDnsARecord,omitempty"`
	// Specifies whether to enable DNS resolution from the IP-based domain name to the IPv4 address. Valid values:
	//
	// - true: enables the resolution.
	//
	// - false: disables the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	EnableIpDnsARecord *bool `json:"EnableIpDnsARecord,omitempty" xml:"EnableIpDnsARecord,omitempty"`
	// Specifies whether to enable reverse DNS resolution from the IPv4 address to the IP-based domain name. Valid values:
	//
	// - true: enables the resolution.
	//
	// - false: disables the resolution.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableIpDnsPtrRecord *bool `json:"EnableIpDnsPtrRecord,omitempty" xml:"EnableIpDnsPtrRecord,omitempty"`
	// The hostname type. Valid values:
	//
	// - Custom: custom.
	//
	// - IpBased: IP-based hostname.
	//
	// - InstanceIdBased: instance ID-based hostname.
	//
	// Default value: Custom.
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
	// The tag key for the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
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
