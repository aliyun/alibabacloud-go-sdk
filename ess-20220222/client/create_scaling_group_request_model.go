// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlbServerGroups(v []*CreateScalingGroupRequestAlbServerGroups) *CreateScalingGroupRequest
	GetAlbServerGroups() []*CreateScalingGroupRequestAlbServerGroups
	SetAllocationStrategy(v string) *CreateScalingGroupRequest
	GetAllocationStrategy() *string
	SetAutoRebalance(v bool) *CreateScalingGroupRequest
	GetAutoRebalance() *bool
	SetAzBalance(v bool) *CreateScalingGroupRequest
	GetAzBalance() *bool
	SetBalanceMode(v string) *CreateScalingGroupRequest
	GetBalanceMode() *string
	SetCapacityOptions(v *CreateScalingGroupRequestCapacityOptions) *CreateScalingGroupRequest
	GetCapacityOptions() *CreateScalingGroupRequestCapacityOptions
	SetClientToken(v string) *CreateScalingGroupRequest
	GetClientToken() *string
	SetCompensateWithOnDemand(v bool) *CreateScalingGroupRequest
	GetCompensateWithOnDemand() *bool
	SetContainerGroupId(v string) *CreateScalingGroupRequest
	GetContainerGroupId() *string
	SetCustomPolicyARN(v string) *CreateScalingGroupRequest
	GetCustomPolicyARN() *string
	SetDBInstanceIds(v string) *CreateScalingGroupRequest
	GetDBInstanceIds() *string
	SetDBInstances(v []*CreateScalingGroupRequestDBInstances) *CreateScalingGroupRequest
	GetDBInstances() []*CreateScalingGroupRequestDBInstances
	SetDefaultCooldown(v int32) *CreateScalingGroupRequest
	GetDefaultCooldown() *int32
	SetDesiredCapacity(v int32) *CreateScalingGroupRequest
	GetDesiredCapacity() *int32
	SetGroupDeletionProtection(v bool) *CreateScalingGroupRequest
	GetGroupDeletionProtection() *bool
	SetGroupType(v string) *CreateScalingGroupRequest
	GetGroupType() *string
	SetHealthCheckType(v string) *CreateScalingGroupRequest
	GetHealthCheckType() *string
	SetHealthCheckTypes(v []*string) *CreateScalingGroupRequest
	GetHealthCheckTypes() []*string
	SetInstanceId(v string) *CreateScalingGroupRequest
	GetInstanceId() *string
	SetLaunchTemplateId(v string) *CreateScalingGroupRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateOverrides(v []*CreateScalingGroupRequestLaunchTemplateOverrides) *CreateScalingGroupRequest
	GetLaunchTemplateOverrides() []*CreateScalingGroupRequestLaunchTemplateOverrides
	SetLaunchTemplateVersion(v string) *CreateScalingGroupRequest
	GetLaunchTemplateVersion() *string
	SetLifecycleHooks(v []*CreateScalingGroupRequestLifecycleHooks) *CreateScalingGroupRequest
	GetLifecycleHooks() []*CreateScalingGroupRequestLifecycleHooks
	SetLoadBalancerConfigs(v []*CreateScalingGroupRequestLoadBalancerConfigs) *CreateScalingGroupRequest
	GetLoadBalancerConfigs() []*CreateScalingGroupRequestLoadBalancerConfigs
	SetLoadBalancerIds(v string) *CreateScalingGroupRequest
	GetLoadBalancerIds() *string
	SetMaxInstanceLifetime(v int32) *CreateScalingGroupRequest
	GetMaxInstanceLifetime() *int32
	SetMaxSize(v int32) *CreateScalingGroupRequest
	GetMaxSize() *int32
	SetMinSize(v int32) *CreateScalingGroupRequest
	GetMinSize() *int32
	SetMultiAZPolicy(v string) *CreateScalingGroupRequest
	GetMultiAZPolicy() *string
	SetOnDemandBaseCapacity(v int32) *CreateScalingGroupRequest
	GetOnDemandBaseCapacity() *int32
	SetOnDemandPercentageAboveBaseCapacity(v int32) *CreateScalingGroupRequest
	GetOnDemandPercentageAboveBaseCapacity() *int32
	SetOwnerAccount(v string) *CreateScalingGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateScalingGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateScalingGroupRequest
	GetRegionId() *string
	SetRemovalPolicies(v []*string) *CreateScalingGroupRequest
	GetRemovalPolicies() []*string
	SetResourceGroupId(v string) *CreateScalingGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateScalingGroupRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupName(v string) *CreateScalingGroupRequest
	GetScalingGroupName() *string
	SetScalingPolicy(v string) *CreateScalingGroupRequest
	GetScalingPolicy() *string
	SetServerGroups(v []*CreateScalingGroupRequestServerGroups) *CreateScalingGroupRequest
	GetServerGroups() []*CreateScalingGroupRequestServerGroups
	SetSpotAllocationStrategy(v string) *CreateScalingGroupRequest
	GetSpotAllocationStrategy() *string
	SetSpotInstancePools(v int32) *CreateScalingGroupRequest
	GetSpotInstancePools() *int32
	SetSpotInstanceRemedy(v bool) *CreateScalingGroupRequest
	GetSpotInstanceRemedy() *bool
	SetStopInstanceTimeout(v int32) *CreateScalingGroupRequest
	GetStopInstanceTimeout() *int32
	SetSyncAlarmRuleToCms(v bool) *CreateScalingGroupRequest
	GetSyncAlarmRuleToCms() *bool
	SetTags(v []*CreateScalingGroupRequestTags) *CreateScalingGroupRequest
	GetTags() []*CreateScalingGroupRequestTags
	SetVServerGroups(v []*CreateScalingGroupRequestVServerGroups) *CreateScalingGroupRequest
	GetVServerGroups() []*CreateScalingGroupRequestVServerGroups
	SetVSwitchId(v string) *CreateScalingGroupRequest
	GetVSwitchId() *string
	SetVSwitchIds(v []*string) *CreateScalingGroupRequest
	GetVSwitchIds() []*string
}

type CreateScalingGroupRequest struct {
	// The Application Load Balancer (ALB) server groups to associate with the scaling group.
	AlbServerGroups []*CreateScalingGroupRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The capacity allocation policy determines how the scaling group selects available instance types to meet capacity requirements. The policy applies to both on-demand and preemptible capacity (effective only when the `MultiAZPolicy` parameter is set to `COMPOSABLE`). Valid values:
	//
	// - priority: Creates instances in the order of the configured instance types.
	//
	// - lowestPrice: Create instances based on the price per vCPU of instance types, from lowest to highest.
	//
	// Default value: priority.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Specifies whether to enable automatic balancing for the scaling group. This setting takes effect only when BalancedOnly is enabled for a scaling group that is balanced across availability zones. Value range:
	//
	// - false: Does not enable automatic balancing for the scaling group.
	//
	// - true: When automatic balancing for the scaling group is enabled, the scaling group automatically detects the capacity across availability zones. If the capacity is imbalanced, the scaling group proactively performs scaling across availability zones to rebalance the capacity.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRebalance *bool `json:"AutoRebalance,omitempty" xml:"AutoRebalance,omitempty"`
	// Specifies whether to evenly distribute the capacity of the scaling group across multiple availability zones. This parameter is valid only when `MultiAZPolicy` is set to `COMPOSABLE`. Valid values:
	//
	// - `true`: The capacity of the scaling group is evenly distributed across multiple availability zones.
	//
	// - `false`: The capacity of the scaling group is not evenly distributed across multiple availability zones.
	//
	// > If `MultiAZPolicy` is set to `COMPOSABLE` and `AzBalance` is set to `true`, the effect is the same as setting `MultiAZPolicy` to `BALANCE`.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	AzBalance *bool `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	// The zone balancing mode is effective only when enabled. Valid values:
	//
	// - BalancedBestEffort: If a resource fails to be created in an availability zone, the system falls back to other availability zones to ensure best-effort delivery.
	//
	// - BalancedOnly:
	//
	//   If resource creation fails in an availability zone, the system does not fall back to other availability zones. The scaling activity is partially successful, which prevents an excessive imbalance of resources across different availability zones.
	//
	// Default value: BalancedBestEffort.
	//
	// example:
	//
	// BalancedBestEffort
	BalanceMode *string `json:"BalanceMode,omitempty" xml:"BalanceMode,omitempty"`
	// The capacity options.
	CapacityOptions *CreateScalingGroupRequestCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// A client-generated token to ensure the idempotence of the request.
	//
	// The token must be unique across requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is effective only when `MultiAZPolicy` is set to `COST_OPTIMIZED`. If `true`, Auto Scaling creates on-demand instances to meet capacity requirements when spot instances are unavailable due to price or inventory. Valid values:
	//
	// - `true`: Yes.
	//
	// - `false`: No.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The ID of the ECI instance, also known as the container group ID.
	//
	// example:
	//
	// eci-uf6fonnghi50u374****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The ARN of the custom scale-in policy function. This parameter is valid only when the first removal policy in `RemovalPolicies` is `CustomPolicy`.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// A JSON array of RDS instance IDs.
	//
	// <props="china">
	//
	// The number of RDS instances that you can associate with a single scaling group varies based on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota for **Maximum number of RDS instances that can be associated with a single scaling group**.
	//
	//
	//
	// <props="intl">
	//
	// The number of RDS instances that you can associate with a single scaling group varies based on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota for **Maximum number of RDS instances that can be associated with a single scaling group**.
	//
	//
	//
	// <props="partner">
	//
	// The number of RDS instances that you can associate with a single scaling group varies based on your Auto Scaling usage. Go to Quota Center to view the quota for **Maximum number of RDS instances that can be associated with a single scaling group**.
	//
	// example:
	//
	// ["rm-bp142f86de0t7****", "rm-bp18l1z42ar4o****", "rm-bp1lqr97h4aqk****"]
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The databases that are associated with the scaling group.
	DBInstances []*CreateScalingGroupRequestDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The cooldown period, in seconds, after a scaling activity completes. Valid values: 0 to 86400.
	//
	// During the cooldown period, the scaling group does not execute other scaling activities that are triggered by CloudMonitor alarm tasks.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The desired number of instances in the scaling group. Auto Scaling automatically maintains this number of instances. The value must be less than or equal to `MaxSize` and greater than or equal to `MinSize`.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Specifies whether to enable deletion protection for the scaling group. Valid values:
	//
	// - `true`: Enables deletion protection. The scaling group cannot be deleted.
	//
	// - `false`: Disables deletion protection.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The type of instances managed by the scaling group. Valid values:
	//
	// - `ECS`: The scaling group manages ECS instances.
	//
	// - `ECI`: The scaling group manages ECI instances.
	//
	// Default value: `ECS`.
	//
	// example:
	//
	// ECS
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The health check method for the scaling group. Valid values:
	//
	// - `NONE`: No health checks are performed.
	//
	// - `ECS`: Health checks are performed on instances in the scaling group. This value enables health checks for scaling groups of both the ECS and ECI types.
	//
	// - `LOAD_BALANCER`: The instance health status is based on health check results from the attached load balancer. This option does not support Classic Load Balancer (CLB) instances.
	//
	// Default value: `ECS`.
	//
	// > To enable both instance health checks and load balancer health checks, use the `HealthCheckTypes` parameter.
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check methods for the scaling group.
	//
	// > You can use this parameter to set multiple values and enable multiple health check options. If you set the `HealthCheckType` parameter, this parameter is ignored.
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The ID of an existing instance to use as a template. Auto Scaling uses this instance to create a new scaling configuration for the scaling group.
	//
	// example:
	//
	// i-28wt4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the launch template that provides the configuration for the scaling group.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The instance type information for extending the launch template.
	LaunchTemplateOverrides []*CreateScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the launch template. Valid values:
	//
	// - A specific version number of the template.
	//
	// - `Default`: Uses the default version of the template.
	//
	// - `Latest`: Uses the latest version of the template.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The list of lifecycle hooks.
	LifecycleHooks []*CreateScalingGroupRequestLifecycleHooks `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Repeated"`
	// The load balancer configurations.
	LoadBalancerConfigs []*CreateScalingGroupRequestLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// A JSON array of Classic Load Balancer (CLB) instance IDs.
	//
	// <props="china">
	//
	// The number of CLB instances that you can associate with a single scaling group varies based on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota for **Maximum number of load balancer instances that can be associated with a single scaling group**.
	//
	//
	//
	// <props="intl">
	//
	// The number of CLB instances that you can associate with a single scaling group varies based on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota for **Maximum number of load balancer instances that can be associated with a single scaling group**.
	//
	//
	//
	// <props="partner">
	//
	// The number of CLB instances that you can associate with a single scaling group varies based on your Auto Scaling usage. Go to Quota Center to view the quota for **Maximum number of load balancer instances that can be associated with a single scaling group**.
	//
	// example:
	//
	// ["lb-bp1u7etiogg38yvwz****", "lb-bp168cqrux9ai9l7f****", "lb-bp1jv3m9zvj22ufxp****"]
	LoadBalancerIds *string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	// The maximum lifetime of an instance in the scaling group. Unit: seconds.
	//
	// Value range: [86400, Integer.maxValue].
	//
	// Default value: null.
	//
	// example:
	//
	// null
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of instances in the scaling group. If the total number of instances exceeds this value, Auto Scaling removes instances to meet this maximum.
	//
	// <props="china">
	//
	// The value range of `MaxSize` depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota for **Maximum number of instances per scaling group**.
	//
	//
	//
	// <props="intl">
	//
	// The value range of `MaxSize` depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota for **Maximum number of instances per scaling group**.
	//
	//
	//
	// <props="partner">
	//
	// The value range of `MaxSize` depends on your Auto Scaling usage. Go to Quota Center to view the quota for **Maximum number of instances per scaling group**.
	//
	//
	//
	// If the quota for **Maximum number of instances per scaling group*	- is 2,000, the value of `MaxSize` can range from 0 to 2,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of instances in the scaling group. If the total number of instances falls below this value, Auto Scaling adds instances to meet this minimum.
	//
	// > The value of `MinSize` must be less than or equal to the value of `MaxSize`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The scaling policy for ECS instances in a multi-zone scaling group. Valid values:
	//
	// - `PRIORITY`: Auto Scaling prioritizes the vSwitches specified in `VSwitchIds`. If an operation fails in a higher-priority availability zone, Auto Scaling automatically attempts it in the next-highest-priority zone.
	//
	// - `COST_OPTIMIZED`: During scale-out, creates instances from the instance types with the lowest vCPU unit price. During scale-in, removes instances from the instance types with the highest vCPU unit price. If the scaling configuration includes multiple spot instance types, spot instances are prioritized for creation. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically create on-demand instances when spot instances cannot be created due to reasons such as insufficient inventory.
	//
	//   > The `COST_OPTIMIZED` policy takes effect only when the scaling configuration specifies multiple instance types or includes spot instances.
	//
	// - `BALANCE`: Distributes ECS instances evenly across the specified availability zones in the scaling group. If the distribution of instances becomes uneven due to insufficient inventory, you can call the [RebalanceInstance](https://help.aliyun.com/document_detail/71516.html) API operation to rebalance the instances.
	//
	//   > If `MultiAZPolicy` is set to `BALANCE`, the effect is the same as setting `MultiAZPolicy` to `COMPOSABLE` and `AzBalance` to `true`.
	//
	// - `COMPOSABLE`: A composite policy that allows you to combine the preceding policies for multi-zone scaling groups as needed. You can also specify additional parameters to gain finer control over the capacity of your scaling group.
	//
	// Default value: `PRIORITY`.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of on-demand instances required in the scaling group. Valid values: 0 to 1,000. If the number of on-demand instances is less than this value, Auto Scaling preferentially creates on-demand instances.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of on-demand instances among the excess instances after the minimum number of on-demand instances (`OnDemandBaseCapacity`) is met. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32  `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	OwnerAccount                        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the scaling group resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance removal policies. Valid values:
	//
	// - `OldestInstance`: Removes the ECS instance that was first added to the scaling group.
	//
	// - `NewestInstance`: Removes the ECS instance that was most recently added to the scaling group.
	//
	// - `OldestScalingConfiguration`: Removes the ECS instance that was created based on the earliest scaling configuration.
	//
	// - `CustomPolicy`: Removes ECS instances based on a custom scale-in policy defined by a function.
	//
	// The term `scaling configuration` in `OldestScalingConfiguration` refers to the source of instance configuration information, which includes both scaling configurations and launch templates. `CustomPolicy` can only be set as the first removal policy. If you specify `CustomPolicy`, you must also specify the `CustomPolicyARN` parameter.
	//
	// > The removal of instances is also affected by the scaling group\\"s multi-AZ policy (`MultiAZPolicy`). For more information, see [Configure a combination of removal policies](https://help.aliyun.com/document_detail/254822.html).
	RemovalPolicies []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	// The ID of the resource group to which the new scaling group belongs.
	//
	// > If you do not specify this parameter, the new scaling group is added to the default resource group.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The name of the scaling group. The name must be unique within a region.
	//
	// The name must be 2 to 64 characters in length. It must start with a letter, a digit, or a Chinese character and can contain digits, underscores (_), hyphens (-), and periods (.).
	//
	// If you do not specify this parameter, the value of `ScalingGroupId` is used.
	//
	// example:
	//
	// scalinggroup****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The reclamation mode of the scaling group. Valid values:
	//
	// - `recycle`: The reclamation mode is Economical Mode.
	//
	// - `release`: The reclamation mode is Release Mode.
	//
	// - `forcerelease`: The reclamation mode is Force Release Mode.
	//
	//   > A forced release is equivalent to a power-off operation, which erases data in the memory and ephemeral storage of the instances. This data cannot be recovered. Use this option with caution.
	//
	// - `forcerecycle`: The reclamation mode is Force Economical Mode.
	//
	//   > A forced stop is equivalent to a power-off operation, which erases data in the memory and ephemeral storage of the instances. This data cannot be recovered. Use this option with caution.
	//
	// `ScalingPolicy` specifies the reclamation mode of the scaling group. The specific action taken when an instance is removed from the scaling group is determined by the `RemovePolicy` parameter of the `RemoveInstances` operation. For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The load balancer server groups.
	//
	// > You cannot specify the same server group in both `AlbServerGroups` and `ServerGroups`.
	ServerGroups []*CreateScalingGroupRequestServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The distribution strategy for spot capacity. You can use this parameter to specify a separate strategy for spot capacity (effective only when the `MultiAZPolicy` parameter is set to `COMPOSABLE`). Valid values:
	//
	// - priority: Creates instances in the order of the configured instance types.
	//
	// - lowestPrice: Creates instances in ascending order of the price per vCPU of the instance types.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of instance types to use. The scaling group creates spot instances in a balanced manner across the specified number of lowest-cost instance types. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// If set to `true`, Auto Scaling attempts to create a new instance to replace a spot instance that is about to be reclaimed.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The timeout period for the system to wait for an ECS instance to be stopped during a scale-in event. Unit: seconds.
	//
	// Valid values: 30 to 240.
	//
	// > - This parameter takes effect only during scale-in events when `ScalingPolicy` is set to `release`.
	//
	// >
	//
	// > - If this parameter is set, the system waits for the specified `StopInstanceTimeout` period for the instance to be stopped. If the instance is not stopped after the timeout period, the scale-in process continues regardless of the instance status.
	//
	// >
	//
	// > - If this parameter is not set, the system waits for an extended period for the instance to be stopped. The scale-in process continues only after the instance is stopped. If the instance fails to stop, the scale-in process is rolled back, and the scale-in event fails.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// > This parameter is not yet available.
	//
	// example:
	//
	// false
	SyncAlarmRuleToCms *bool `json:"SyncAlarmRuleToCms,omitempty" xml:"SyncAlarmRuleToCms,omitempty"`
	// The tags to apply to the scaling group.
	Tags []*CreateScalingGroupRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vServer groups to associate with the scaling group.
	VServerGroups []*CreateScalingGroupRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	// The ID of the vSwitch. If you specify this parameter, the network type of the scaling group is Virtual Private Cloud (VPC).
	//
	// > If you do not specify the `VSwitchId` or `VSwitchIds` parameter, the network type of the scaling group defaults to classic network.
	//
	// example:
	//
	// vsw-bp14zolna43z266bq****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of one or more vSwitches. If you specify this parameter, the `VSwitchId` parameter is ignored. If you specify this parameter, the network type of the scaling group is Virtual Private Cloud (VPC).
	//
	// If you specify multiple vSwitches:
	//
	// - They must belong to the same VPC.
	//
	// - They can be in different availability zones.
	//
	// - The vSwitches are prioritized based on their order in the list, with the first vSwitch having the highest priority. If an instance cannot be created in the availability zone of a higher-priority vSwitch, Auto Scaling automatically attempts to create the instance in the availability zone of the next-highest-priority vSwitch.
	//
	// > If you do not specify the `VSwitchId` or `VSwitchIds` parameter, the network type of the scaling group defaults to classic network.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s CreateScalingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequest) GetAlbServerGroups() []*CreateScalingGroupRequestAlbServerGroups {
	return s.AlbServerGroups
}

func (s *CreateScalingGroupRequest) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *CreateScalingGroupRequest) GetAutoRebalance() *bool {
	return s.AutoRebalance
}

func (s *CreateScalingGroupRequest) GetAzBalance() *bool {
	return s.AzBalance
}

func (s *CreateScalingGroupRequest) GetBalanceMode() *string {
	return s.BalanceMode
}

func (s *CreateScalingGroupRequest) GetCapacityOptions() *CreateScalingGroupRequestCapacityOptions {
	return s.CapacityOptions
}

func (s *CreateScalingGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateScalingGroupRequest) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *CreateScalingGroupRequest) GetContainerGroupId() *string {
	return s.ContainerGroupId
}

func (s *CreateScalingGroupRequest) GetCustomPolicyARN() *string {
	return s.CustomPolicyARN
}

func (s *CreateScalingGroupRequest) GetDBInstanceIds() *string {
	return s.DBInstanceIds
}

func (s *CreateScalingGroupRequest) GetDBInstances() []*CreateScalingGroupRequestDBInstances {
	return s.DBInstances
}

func (s *CreateScalingGroupRequest) GetDefaultCooldown() *int32 {
	return s.DefaultCooldown
}

func (s *CreateScalingGroupRequest) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *CreateScalingGroupRequest) GetGroupDeletionProtection() *bool {
	return s.GroupDeletionProtection
}

func (s *CreateScalingGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateScalingGroupRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *CreateScalingGroupRequest) GetHealthCheckTypes() []*string {
	return s.HealthCheckTypes
}

func (s *CreateScalingGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScalingGroupRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateScalingGroupRequest) GetLaunchTemplateOverrides() []*CreateScalingGroupRequestLaunchTemplateOverrides {
	return s.LaunchTemplateOverrides
}

func (s *CreateScalingGroupRequest) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *CreateScalingGroupRequest) GetLifecycleHooks() []*CreateScalingGroupRequestLifecycleHooks {
	return s.LifecycleHooks
}

func (s *CreateScalingGroupRequest) GetLoadBalancerConfigs() []*CreateScalingGroupRequestLoadBalancerConfigs {
	return s.LoadBalancerConfigs
}

func (s *CreateScalingGroupRequest) GetLoadBalancerIds() *string {
	return s.LoadBalancerIds
}

func (s *CreateScalingGroupRequest) GetMaxInstanceLifetime() *int32 {
	return s.MaxInstanceLifetime
}

func (s *CreateScalingGroupRequest) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *CreateScalingGroupRequest) GetMinSize() *int32 {
	return s.MinSize
}

func (s *CreateScalingGroupRequest) GetMultiAZPolicy() *string {
	return s.MultiAZPolicy
}

func (s *CreateScalingGroupRequest) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *CreateScalingGroupRequest) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *CreateScalingGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateScalingGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateScalingGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScalingGroupRequest) GetRemovalPolicies() []*string {
	return s.RemovalPolicies
}

func (s *CreateScalingGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateScalingGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateScalingGroupRequest) GetScalingGroupName() *string {
	return s.ScalingGroupName
}

func (s *CreateScalingGroupRequest) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *CreateScalingGroupRequest) GetServerGroups() []*CreateScalingGroupRequestServerGroups {
	return s.ServerGroups
}

func (s *CreateScalingGroupRequest) GetSpotAllocationStrategy() *string {
	return s.SpotAllocationStrategy
}

func (s *CreateScalingGroupRequest) GetSpotInstancePools() *int32 {
	return s.SpotInstancePools
}

func (s *CreateScalingGroupRequest) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *CreateScalingGroupRequest) GetStopInstanceTimeout() *int32 {
	return s.StopInstanceTimeout
}

func (s *CreateScalingGroupRequest) GetSyncAlarmRuleToCms() *bool {
	return s.SyncAlarmRuleToCms
}

func (s *CreateScalingGroupRequest) GetTags() []*CreateScalingGroupRequestTags {
	return s.Tags
}

func (s *CreateScalingGroupRequest) GetVServerGroups() []*CreateScalingGroupRequestVServerGroups {
	return s.VServerGroups
}

func (s *CreateScalingGroupRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateScalingGroupRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateScalingGroupRequest) SetAlbServerGroups(v []*CreateScalingGroupRequestAlbServerGroups) *CreateScalingGroupRequest {
	s.AlbServerGroups = v
	return s
}

func (s *CreateScalingGroupRequest) SetAllocationStrategy(v string) *CreateScalingGroupRequest {
	s.AllocationStrategy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetAutoRebalance(v bool) *CreateScalingGroupRequest {
	s.AutoRebalance = &v
	return s
}

func (s *CreateScalingGroupRequest) SetAzBalance(v bool) *CreateScalingGroupRequest {
	s.AzBalance = &v
	return s
}

func (s *CreateScalingGroupRequest) SetBalanceMode(v string) *CreateScalingGroupRequest {
	s.BalanceMode = &v
	return s
}

func (s *CreateScalingGroupRequest) SetCapacityOptions(v *CreateScalingGroupRequestCapacityOptions) *CreateScalingGroupRequest {
	s.CapacityOptions = v
	return s
}

func (s *CreateScalingGroupRequest) SetClientToken(v string) *CreateScalingGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateScalingGroupRequest) SetCompensateWithOnDemand(v bool) *CreateScalingGroupRequest {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *CreateScalingGroupRequest) SetContainerGroupId(v string) *CreateScalingGroupRequest {
	s.ContainerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetCustomPolicyARN(v string) *CreateScalingGroupRequest {
	s.CustomPolicyARN = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDBInstanceIds(v string) *CreateScalingGroupRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDBInstances(v []*CreateScalingGroupRequestDBInstances) *CreateScalingGroupRequest {
	s.DBInstances = v
	return s
}

func (s *CreateScalingGroupRequest) SetDefaultCooldown(v int32) *CreateScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDesiredCapacity(v int32) *CreateScalingGroupRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *CreateScalingGroupRequest) SetGroupDeletionProtection(v bool) *CreateScalingGroupRequest {
	s.GroupDeletionProtection = &v
	return s
}

func (s *CreateScalingGroupRequest) SetGroupType(v string) *CreateScalingGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateScalingGroupRequest) SetHealthCheckType(v string) *CreateScalingGroupRequest {
	s.HealthCheckType = &v
	return s
}

func (s *CreateScalingGroupRequest) SetHealthCheckTypes(v []*string) *CreateScalingGroupRequest {
	s.HealthCheckTypes = v
	return s
}

func (s *CreateScalingGroupRequest) SetInstanceId(v string) *CreateScalingGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateId(v string) *CreateScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateOverrides(v []*CreateScalingGroupRequestLaunchTemplateOverrides) *CreateScalingGroupRequest {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *CreateScalingGroupRequest) SetLaunchTemplateVersion(v string) *CreateScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLifecycleHooks(v []*CreateScalingGroupRequestLifecycleHooks) *CreateScalingGroupRequest {
	s.LifecycleHooks = v
	return s
}

func (s *CreateScalingGroupRequest) SetLoadBalancerConfigs(v []*CreateScalingGroupRequestLoadBalancerConfigs) *CreateScalingGroupRequest {
	s.LoadBalancerConfigs = v
	return s
}

func (s *CreateScalingGroupRequest) SetLoadBalancerIds(v string) *CreateScalingGroupRequest {
	s.LoadBalancerIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMaxInstanceLifetime(v int32) *CreateScalingGroupRequest {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMaxSize(v int32) *CreateScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMinSize(v int32) *CreateScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMultiAZPolicy(v string) *CreateScalingGroupRequest {
	s.MultiAZPolicy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOnDemandBaseCapacity(v int32) *CreateScalingGroupRequest {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOnDemandPercentageAboveBaseCapacity(v int32) *CreateScalingGroupRequest {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerAccount(v string) *CreateScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerId(v int64) *CreateScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRegionId(v string) *CreateScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRemovalPolicies(v []*string) *CreateScalingGroupRequest {
	s.RemovalPolicies = v
	return s
}

func (s *CreateScalingGroupRequest) SetResourceGroupId(v string) *CreateScalingGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetResourceOwnerAccount(v string) *CreateScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetScalingGroupName(v string) *CreateScalingGroupRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *CreateScalingGroupRequest) SetScalingPolicy(v string) *CreateScalingGroupRequest {
	s.ScalingPolicy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetServerGroups(v []*CreateScalingGroupRequestServerGroups) *CreateScalingGroupRequest {
	s.ServerGroups = v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotAllocationStrategy(v string) *CreateScalingGroupRequest {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotInstancePools(v int32) *CreateScalingGroupRequest {
	s.SpotInstancePools = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSpotInstanceRemedy(v bool) *CreateScalingGroupRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *CreateScalingGroupRequest) SetStopInstanceTimeout(v int32) *CreateScalingGroupRequest {
	s.StopInstanceTimeout = &v
	return s
}

func (s *CreateScalingGroupRequest) SetSyncAlarmRuleToCms(v bool) *CreateScalingGroupRequest {
	s.SyncAlarmRuleToCms = &v
	return s
}

func (s *CreateScalingGroupRequest) SetTags(v []*CreateScalingGroupRequestTags) *CreateScalingGroupRequest {
	s.Tags = v
	return s
}

func (s *CreateScalingGroupRequest) SetVServerGroups(v []*CreateScalingGroupRequestVServerGroups) *CreateScalingGroupRequest {
	s.VServerGroups = v
	return s
}

func (s *CreateScalingGroupRequest) SetVSwitchId(v string) *CreateScalingGroupRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetVSwitchIds(v []*string) *CreateScalingGroupRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateScalingGroupRequest) Validate() error {
	if s.AlbServerGroups != nil {
		for _, item := range s.AlbServerGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CapacityOptions != nil {
		if err := s.CapacityOptions.Validate(); err != nil {
			return err
		}
	}
	if s.DBInstances != nil {
		for _, item := range s.DBInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LaunchTemplateOverrides != nil {
		for _, item := range s.LaunchTemplateOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LifecycleHooks != nil {
		for _, item := range s.LifecycleHooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LoadBalancerConfigs != nil {
		for _, item := range s.LoadBalancerConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServerGroups != nil {
		for _, item := range s.ServerGroups {
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
	if s.VServerGroups != nil {
		for _, item := range s.VServerGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScalingGroupRequestAlbServerGroups struct {
	// The ID of the ALB server group.
	//
	// A scaling group can be associated with a limited number of ALB server groups. To view or request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port number used by an instance after it is added to the ALB server group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The weight of an instance as a backend server after the instance is added to the ALB server group. The higher the weight, the more access requests are distributed to the instance. If the weight is 0, no access requests are distributed to the instance. Valid values: 0 to 100.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateScalingGroupRequestAlbServerGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestAlbServerGroups) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestAlbServerGroups) GetAlbServerGroupId() *string {
	return s.AlbServerGroupId
}

func (s *CreateScalingGroupRequestAlbServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *CreateScalingGroupRequestAlbServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateScalingGroupRequestAlbServerGroups) SetAlbServerGroupId(v string) *CreateScalingGroupRequestAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroups) SetPort(v int32) *CreateScalingGroupRequestAlbServerGroups {
	s.Port = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroups) SetWeight(v int32) *CreateScalingGroupRequestAlbServerGroups {
	s.Weight = &v
	return s
}

func (s *CreateScalingGroupRequestAlbServerGroups) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestCapacityOptions struct {
	// When `MultiAZPolicy` is set to `COST_OPTIMIZED`, this parameter specifies whether to automatically create on-demand instances to meet capacity requirements when spot instances are unavailable due to price or inventory. Valid values:
	//
	// - `true`: Yes.
	//
	// - `false`: No.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The minimum number of on-demand instances required in the scaling group. When the number of on-demand instances in the scaling group is less than this value, the system preferentially creates on-demand instances. Valid values: 0 to 1,000.
	//
	// When `MultiAZPolicy` is set to `COMPOSABLE`, the default value is 0.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of on-demand instances among the excess instances after the `OnDemandBaseCapacity` requirement is met. Valid values: 0 to 100.
	//
	// When `MultiAZPolicy` is set to `COMPOSABLE`, the default value is 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The price comparison mode for the cost optimization strategy of the scaling group. Valid values:
	//
	// - `PricePerUnit`: Compares prices based on per-unit capacity.
	//
	//   The capacity of an instance in a scaling group is equal to the weight set for the instance type, with a default of 1, meaning one ECS instance equals one unit of capacity.
	//
	// - `PricePerVCpu`: Compares prices based on per-vCPU price.
	//
	// Default value: `PricePerUnit`.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// After you enable `CompensateWithOnDemand`, if the on-demand percentage exceeds the `OnDemandPercentageAboveBaseCapacity` ratio, the system attempts to replace on-demand capacity with spot capacity. A common scenario is when `CompensateWithOnDemand` leads to on-demand instances being created due to spot inventory or price issues. To avoid the prolonged existence of these on-demand instances, the system attempts to replace the excess on-demand capacity with spot instances. Valid values:
	//
	// - `true`: Allows replacement.
	//
	// - `false`: Does not allow replacement.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SpotAutoReplaceOnDemand *bool `json:"SpotAutoReplaceOnDemand,omitempty" xml:"SpotAutoReplaceOnDemand,omitempty"`
}

func (s CreateScalingGroupRequestCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestCapacityOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestCapacityOptions) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *CreateScalingGroupRequestCapacityOptions) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *CreateScalingGroupRequestCapacityOptions) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *CreateScalingGroupRequestCapacityOptions) GetPriceComparisonMode() *string {
	return s.PriceComparisonMode
}

func (s *CreateScalingGroupRequestCapacityOptions) GetSpotAutoReplaceOnDemand() *bool {
	return s.SpotAutoReplaceOnDemand
}

func (s *CreateScalingGroupRequestCapacityOptions) SetCompensateWithOnDemand(v bool) *CreateScalingGroupRequestCapacityOptions {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *CreateScalingGroupRequestCapacityOptions) SetOnDemandBaseCapacity(v int32) *CreateScalingGroupRequestCapacityOptions {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *CreateScalingGroupRequestCapacityOptions) SetOnDemandPercentageAboveBaseCapacity(v int32) *CreateScalingGroupRequestCapacityOptions {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *CreateScalingGroupRequestCapacityOptions) SetPriceComparisonMode(v string) *CreateScalingGroupRequestCapacityOptions {
	s.PriceComparisonMode = &v
	return s
}

func (s *CreateScalingGroupRequestCapacityOptions) SetSpotAutoReplaceOnDemand(v bool) *CreateScalingGroupRequestCapacityOptions {
	s.SpotAutoReplaceOnDemand = &v
	return s
}

func (s *CreateScalingGroupRequestCapacityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestDBInstances struct {
	// The method that is used to associate the scaling group with the database. Valid values:
	//
	// - `SecurityIp`: The IP address whitelist mode. This mode automatically adds the scaled-out instances to the IP address whitelist of the database. This mode is supported only by RDS databases.
	//
	// - `SecurityGroup`: The security group mode. This mode adds the security group of the scaling configuration to the security group whitelist of the database. This allows instances in the security group to access the database.
	//
	// example:
	//
	// SecurityIp
	AttachMode *string `json:"AttachMode,omitempty" xml:"AttachMode,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// rm-m5eqju85s45mu0***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of the database. Valid values:
	//
	// - RDS
	//
	// - Redis
	//
	// - MongoDB
	//
	// Default value: RDS.
	//
	// example:
	//
	// RDS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateScalingGroupRequestDBInstances) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestDBInstances) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestDBInstances) GetAttachMode() *string {
	return s.AttachMode
}

func (s *CreateScalingGroupRequestDBInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateScalingGroupRequestDBInstances) GetType() *string {
	return s.Type
}

func (s *CreateScalingGroupRequestDBInstances) SetAttachMode(v string) *CreateScalingGroupRequestDBInstances {
	s.AttachMode = &v
	return s
}

func (s *CreateScalingGroupRequestDBInstances) SetDBInstanceId(v string) *CreateScalingGroupRequestDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *CreateScalingGroupRequestDBInstances) SetType(v string) *CreateScalingGroupRequestDBInstances {
	s.Type = &v
	return s
}

func (s *CreateScalingGroupRequestDBInstances) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestLaunchTemplateOverrides struct {
	// To enable the scaling group to scale based on instance type capacity, you must specify both this parameter and `LaunchTemplateOverrides.WeightedCapacity`.
	//
	// This parameter specifies the instance type, which overrides the instance type specified in the launch template.
	//
	// > This parameter takes effect only when the `LaunchTemplateId` parameter is specified.
	//
	// Must be a valid ECS instance type.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum hourly price for the instance type specified in `LaunchTemplateOverride.InstanceType`.
	//
	// > This parameter takes effect only when the `LaunchTemplateId` parameter is specified.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// To enable the scaling group to scale based on instance type capacity, you must specify this parameter after you specify `LaunchTemplateOverrides.InstanceType`.
	//
	// This parameter specifies the weight of the instance type, which represents the capacity of a single instance of that type in the scaling group. A higher weight means that fewer instances of this type are needed to meet the desired capacity.
	//
	// Because instance types have different performance metrics, such as the number of vCPUs and memory size, you can assign different weights to different instance types based on your requirements.
	//
	// Example:
	//
	// - Current capacity: 0.
	//
	// - Desired capacity: 6.
	//
	// - Capacity of ecs.c5.xlarge: 4.
	//
	// To meet the desired capacity, the scaling group will create two ecs.c5.xlarge instances.
	//
	// > During a scale-out activity, the capacity of the scaling group cannot exceed the sum of the maximum capacity (`MaxSize`) and the maximum weight of an instance type.
	//
	// Valid values: 1 to 500.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s CreateScalingGroupRequestLaunchTemplateOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) SetInstanceType(v string) *CreateScalingGroupRequestLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) SetSpotPriceLimit(v float32) *CreateScalingGroupRequestLaunchTemplateOverrides {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) SetWeightedCapacity(v int32) *CreateScalingGroupRequestLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *CreateScalingGroupRequestLaunchTemplateOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestLifecycleHooks struct {
	// The action to take after the wait state ends. Valid values:
	//
	// - `CONTINUE`: Continues the scale-out or scale-in activity.
	//
	// - `ABANDON`: Aborts the scale-out activity by releasing the created instances, or aborts the scale-in activity by keeping the instances in the scaling group.
	//
	// If a scale-in (SCALE_IN) activity triggers multiple lifecycle hooks, and the `DefaultResult` of one of the lifecycle hooks is `ABANDON`, the wait state of the other lifecycle hooks ends prematurely. In other cases, the action is determined by the last lifecycle hook to complete.
	//
	// Default value: `CONTINUE`.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The wait time that is defined in the lifecycle hook for a scaling activity. After the wait time expires, the next action is performed. Valid values: 30 to 21600. Unit: seconds.
	//
	// After you create a lifecycle hook, you can call the `RecordLifecycleActionHeartbeat` operation to extend the wait time of an instance, or call the `CompleteLifecycleAction` operation to end the wait state of the scaling activity in advance.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The name of the lifecycle hook. The name cannot be modified after it is specified. If you do not specify a name, the ID of the lifecycle hook is used.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The type of scaling activity to which the lifecycle hook applies. Valid values:
	//
	// - `SCALE_OUT`: A scale-out activity.
	//
	// - `SCALE_IN`: A scale-in activity.
	//
	// > This parameter is required if you specify a lifecycle hook for the scaling group. Other related parameters are optional.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient for the lifecycle hook. Message Service (MNS) queues and topics are supported. The format is `acs:ess:{region}:{account-id}:{resource-relative-id}`.
	//
	// - `region`: the region where the scaling group is located.
	//
	// - `account-id`: the ID of your Alibaba Cloud account.
	//
	// Examples:
	//
	// - MNS queue: `acs:ess:{region}:{account-id}:queue/{queuename}`.
	//
	// - MNS topic: `acs:ess:{region}:{account-id}:topic/{topicname}`.
	//
	// example:
	//
	// acs:ess:cn-hangzhou:1111111111:queue/queue2
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// A fixed string of information for the wait state of a scaling activity. The value cannot exceed 4,096 characters in length. When Auto Scaling sends a message to the specified notification recipient, it includes the value of this parameter. This allows you to manage and categorize notifications. This parameter is valid only when you specify the `NotificationArn` parameter.
	//
	// example:
	//
	// Test
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
}

func (s CreateScalingGroupRequestLifecycleHooks) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestLifecycleHooks) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLifecycleHooks) GetDefaultResult() *string {
	return s.DefaultResult
}

func (s *CreateScalingGroupRequestLifecycleHooks) GetHeartbeatTimeout() *int32 {
	return s.HeartbeatTimeout
}

func (s *CreateScalingGroupRequestLifecycleHooks) GetLifecycleHookName() *string {
	return s.LifecycleHookName
}

func (s *CreateScalingGroupRequestLifecycleHooks) GetLifecycleTransition() *string {
	return s.LifecycleTransition
}

func (s *CreateScalingGroupRequestLifecycleHooks) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *CreateScalingGroupRequestLifecycleHooks) GetNotificationMetadata() *string {
	return s.NotificationMetadata
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetDefaultResult(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.DefaultResult = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetHeartbeatTimeout(v int32) *CreateScalingGroupRequestLifecycleHooks {
	s.HeartbeatTimeout = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetLifecycleHookName(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.LifecycleHookName = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetLifecycleTransition(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.LifecycleTransition = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetNotificationArn(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.NotificationArn = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) SetNotificationMetadata(v string) *CreateScalingGroupRequestLifecycleHooks {
	s.NotificationMetadata = &v
	return s
}

func (s *CreateScalingGroupRequestLifecycleHooks) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestLoadBalancerConfigs struct {
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-2zen1olhfg9yw3f4q****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of an instance as a backend server after the instance is added to the SLB server group. The higher the weight, the more access requests are distributed to the instance. If the weight is 0, no access requests are distributed to the instance. Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateScalingGroupRequestLoadBalancerConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestLoadBalancerConfigs) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestLoadBalancerConfigs) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateScalingGroupRequestLoadBalancerConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateScalingGroupRequestLoadBalancerConfigs) SetLoadBalancerId(v string) *CreateScalingGroupRequestLoadBalancerConfigs {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateScalingGroupRequestLoadBalancerConfigs) SetWeight(v int32) *CreateScalingGroupRequestLoadBalancerConfigs {
	s.Weight = &v
	return s
}

func (s *CreateScalingGroupRequestLoadBalancerConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestServerGroups struct {
	// The port number used by an instance after it is added to the server group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// sgp-5yc3bd9lfyh*****
	ServerGroupId *string `json:"ServerGroupId,omitempty" xml:"ServerGroupId,omitempty"`
	// The type of the server group. Valid values:
	//
	// - `ALB`: Application Load Balancer.
	//
	// - `NLB`: Network Load Balancer.
	//
	// - `GWLB`: Gateway Load Balancer.
	//
	// example:
	//
	// ALB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of an instance as a backend server after the instance is added to the server group. Valid values: 0 to 100.
	//
	// A higher weight indicates that more access requests are distributed to the instance. If the weight is 0, no access requests are distributed to the instance.
	//
	// > This parameter is required for ALB and NLB server groups. You cannot set this parameter for GWLB server groups.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateScalingGroupRequestServerGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestServerGroups) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *CreateScalingGroupRequestServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *CreateScalingGroupRequestServerGroups) GetType() *string {
	return s.Type
}

func (s *CreateScalingGroupRequestServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateScalingGroupRequestServerGroups) SetPort(v int32) *CreateScalingGroupRequestServerGroups {
	s.Port = &v
	return s
}

func (s *CreateScalingGroupRequestServerGroups) SetServerGroupId(v string) *CreateScalingGroupRequestServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestServerGroups) SetType(v string) *CreateScalingGroupRequestServerGroups {
	s.Type = &v
	return s
}

func (s *CreateScalingGroupRequestServerGroups) SetWeight(v int32) *CreateScalingGroupRequestServerGroups {
	s.Weight = &v
	return s
}

func (s *CreateScalingGroupRequestServerGroups) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// Department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies whether the tag can be propagated. Valid values:
	//
	// - `true`: The tag is propagated from the scaling group only to newly created instances, not to instances that are already running in the scaling group.
	//
	// - `false`: The tag is not propagated from the scaling group to any instances.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	Propagate *bool `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// Finance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateScalingGroupRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestTags) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateScalingGroupRequestTags) GetPropagate() *bool {
	return s.Propagate
}

func (s *CreateScalingGroupRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateScalingGroupRequestTags) SetKey(v string) *CreateScalingGroupRequestTags {
	s.Key = &v
	return s
}

func (s *CreateScalingGroupRequestTags) SetPropagate(v bool) *CreateScalingGroupRequestTags {
	s.Propagate = &v
	return s
}

func (s *CreateScalingGroupRequestTags) SetValue(v string) *CreateScalingGroupRequestTags {
	s.Value = &v
	return s
}

func (s *CreateScalingGroupRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateScalingGroupRequestVServerGroups struct {
	// The ID of the Classic Load Balancer (CLB) instance to which the vServer group belongs.
	//
	// example:
	//
	// lb-bp1u7etiogg38yvwz****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the backend server group.
	VServerGroupAttributes []*CreateScalingGroupRequestVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s CreateScalingGroupRequestVServerGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestVServerGroups) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestVServerGroups) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateScalingGroupRequestVServerGroups) GetVServerGroupAttributes() []*CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	return s.VServerGroupAttributes
}

func (s *CreateScalingGroupRequestVServerGroups) SetLoadBalancerId(v string) *CreateScalingGroupRequestVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroups) SetVServerGroupAttributes(v []*CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) *CreateScalingGroupRequestVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

func (s *CreateScalingGroupRequestVServerGroups) Validate() error {
	if s.VServerGroupAttributes != nil {
		for _, item := range s.VServerGroupAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScalingGroupRequestVServerGroupsVServerGroupAttributes struct {
	// The port number used by an instance after it is added to the vServer group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the vServer group.
	//
	// example:
	//
	// rsp-bp1443g77****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The weight of an instance as a backend server after the instance is added to the vServer group. The higher the weight, the more access requests are distributed to the instance. If the weight is 0, no access requests are distributed to the instance. Valid values: 0 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) GetPort() *int32 {
	return s.Port
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) SetPort(v int32) *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) SetWeight(v int32) *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

func (s *CreateScalingGroupRequestVServerGroupsVServerGroupAttributes) Validate() error {
	return dara.Validate(s)
}
