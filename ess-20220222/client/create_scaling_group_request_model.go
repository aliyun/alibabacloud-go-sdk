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
	SetZoneShiftPolicy(v *CreateScalingGroupRequestZoneShiftPolicy) *CreateScalingGroupRequest
	GetZoneShiftPolicy() *CreateScalingGroupRequestZoneShiftPolicy
}

type CreateScalingGroupRequest struct {
	// The collection of information about the ALB server groups.
	AlbServerGroups []*CreateScalingGroupRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The capacity allocation policy, which determines how the scaling group selects available instance types to meet the capacity requirements. The capacity distribution policy applies to both pay-as-you-go and spot capacities. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - priority: Creates instances based on the order of the configured instance types.
	//
	// - lowestPrice: Creates instances based on the unit price per vCPU of the instance types in ascending order.
	//
	// Default value: priority.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Specifies whether to enable automatic balancing for the scaling group. This parameter takes effect only when BalancedOnly is enabled for the zone-balanced scaling group. Valid values:
	//
	// - false: Does not enable automatic balancing for the scaling group.
	//
	// - true: Enables automatic balancing for the scaling group. The scaling group automatically detects the capacity across zones. When the capacity is unbalanced across zones, the scaling group proactively scales out or in to rebalance the capacity across zones.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRebalance *bool `json:"AutoRebalance,omitempty" xml:"AutoRebalance,omitempty"`
	// Specifies whether the capacity of the scaling group is evenly distributed across multiple zones. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - true: The capacity of the scaling group is evenly distributed across multiple zones.
	//
	// - false: The capacity of the scaling group is not evenly distributed across multiple zones.
	//
	// > If the MultiAZPolicy parameter is set to COMPOSABLE and the AzBalance parameter is set to true, the effect is the same as when the MultiAZPolicy parameter is set to BALANCE.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AzBalance *bool `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	// The zone balancing mode. This parameter takes effect only when the zone balancing mode is enabled. Valid values:
	//
	//
	//
	// - BalancedBestEffort: When resource creation fails in a zone, the system falls back to other zones to ensure best-effort delivery of resources.
	//
	// - BalancedOnly: When resource creation fails in a zone, the system does not fall back to other zones. The scale-out activity partially succeeds to avoid excessive imbalance of resources across different zones.
	//
	// Default value: BalancedBestEffort.
	//
	// example:
	//
	// BalancedBestEffort
	BalanceMode *string `json:"BalanceMode,omitempty" xml:"BalanceMode,omitempty"`
	// The capacity options.
	CapacityOptions *CreateScalingGroupRequestCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to automatically attempt to create pay-as-you-go instances to meet the required number of ECS instances when the MultiAZPolicy parameter is set to COST_OPTIMIZED and sufficient spot instances cannot be created due to price or inventory reasons. Valid values:
	//
	// - true: allows the creation.
	//
	// - false: does not allow the creation.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The ECI instance ID, which is the container group ID.
	//
	// example:
	//
	// eci-uf6fonnghi50u374****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The ARN of the Function Compute function for the custom scale-in policy. This parameter takes effect only when the first removal policy configured in RemovalPolicies is CustomPolicy.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The IDs of the ApsaraDB RDS instances. The value can be a JSON array that contains multiple ApsaraDB RDS instance IDs. Separate multiple IDs with commas (,).
	//
	// <props="china">The maximum number of ApsaraDB RDS instances that can be associated with a scaling group depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the quota for the **maximum number of ApsaraDB RDS instances that can be associated with a scaling group**.
	//
	// <props="intl">The maximum number of ApsaraDB RDS instances that can be associated with a scaling group depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the quota for the **maximum number of ApsaraDB RDS instances that can be associated with a scaling group**.
	//
	// <props="partner">The maximum number of ApsaraDB RDS instances that can be associated with a scaling group depends on your Auto Scaling usage. Go to Quota Center to check the quota for the **maximum number of ApsaraDB RDS instances that can be associated with a scaling group**.
	//
	// example:
	//
	// ["rm-bp142f86de0t7****", "rm-bp18l1z42ar4o****", "rm-bp1lqr97h4aqk****"]
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The list of databases associated with the scaling group.
	DBInstances []*CreateScalingGroupRequestDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The cooldown time after a scaling activity (adding or removing instances) is completed. Valid values: 0 to 86400. Unit: seconds.
	//
	// During the cooldown time, the scaling group does not execute other scaling activities. This parameter takes effect only for scaling activities triggered by CloudMonitor event-triggered tasks.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The expected number of instances in the scaling group. The scaling group automatically maintains the number of instances at the expected number. The value cannot be greater than MaxSize or less than MinSize.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Specifies whether to enable deletion protection for the scaling group. Valid values:
	//
	// - true: enables deletion protection for the scaling group. In this case, the scaling group cannot be deleted.
	//
	// - false: disables deletion protection for the scaling group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The type of instances managed by the scaling group. Valid values:
	//
	// - ECS: the scaling group manages ECS instances.
	//
	// - ECI: the scaling group manages ECI instances.
	//
	// Default value: ECS.
	//
	// example:
	//
	// ECS
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The health check type of the scaling group. Valid values:
	//
	// - NONE: disables health checks.
	//
	// - ECS: performs health checks on instances in the scaling group. This value is used to enable instance health checks for both ECS and ECI scaling groups.
	//
	// - LOAD_BALANCER: determines the health status of instances based on the health check results of the load balancer. CLB is not supported.
	//
	//
	// Default value: ECS.
	//
	// > If you want to enable both instance health checks and load balancer health checks, use the `HealthCheckTypes` parameter.
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check modes of the scaling group.
	//
	// > You can set multiple values for this parameter to enable multiple health check options. If the HealthCheckType parameter is specified, this parameter is ignored.
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The ID of the instance. When you create a scaling group, Auto Scaling obtains the required configuration information from the specified instance and automatically creates a scaling configuration.
	//
	// example:
	//
	// i-28wt4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the launch template. This parameter is used to specify the launch template from which the scaling group obtains launch configurations.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The extended instance type information of the launch template.
	LaunchTemplateOverrides []*CreateScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the launch template. Valid values:
	//
	// - A specific version number.
	//
	// - Default: always uses the default version of the launch template.
	//
	// - Latest: always uses the latest version of the launch template.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The list of lifecycle hook information.
	LifecycleHooks []*CreateScalingGroupRequestLifecycleHooks `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Repeated"`
	// The list of load balancing configurations.
	LoadBalancerConfigs []*CreateScalingGroupRequestLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The IDs of the Classic Load Balancer (CLB) instances. The value can be a JSON array that contains multiple CLB instance IDs. Separate multiple IDs with commas (,).
	//
	// <props="china">The maximum number of CLB instances that can be associated with a scaling group depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the quota for the **maximum number of load balancing instances that can be associated with a scaling group**.
	//
	// <props="intl">The maximum number of CLB instances that can be associated with a scaling group depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the quota for the **maximum number of load balancing instances that can be associated with a scaling group**.
	//
	// <props="partner">The maximum number of CLB instances that can be associated with a scaling group depends on your Auto Scaling usage. Go to Quota Center to check the quota for the **maximum number of load balancing instances that can be associated with a scaling group**.
	//
	// example:
	//
	// ["lb-bp1u7etiogg38yvwz****", "lb-bp168cqrux9ai9l7f****", "lb-bp1jv3m9zvj22ufxp****"]
	LoadBalancerIds *string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	// The maximum lifetime of an instance in the scaling group. Unit: seconds.
	//
	// Valid values: [86400, Integer.maxValue].
	//
	// Default value: null.
	//
	// example:
	//
	// 86400
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of instances in the scaling group. If the number of instances in the scaling group is greater than the value of MaxSize, Auto Scaling automatically removes instances.
	//
	// <props="china">The value range of MaxSize depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the quota for the **maximum number of instances in a scaling group**.
	//
	// <props="intl">The value range of MaxSize depends on your Auto Scaling usage. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the quota for the **maximum number of instances in a scaling group**.
	//
	// <props="partner">The value range of MaxSize depends on your Auto Scaling usage. Go to Quota Center to check the quota for the **maximum number of instances in a scaling group**.
	//
	// If the quota for the **maximum number of instances in a scaling group*	- is 2000, the value range of MaxSize is 0 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of instances in the scaling group. If the number of instances in the scaling group is less than the value of MinSize, Auto Scaling automatically creates instances.
	//
	// > The value of MinSize must be less than or equal to the value of MaxSize.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The multi-zone scaling policy for ECS instances in the scaling group. Valid values:
	//
	//
	//
	// - PRIORITY: the vSwitches specified by VSwitchIds are prioritized. Auto Scaling preferentially attempts to scale ECS instances in the zone where the highest-priority vSwitch resides. If the scaling fails, Auto Scaling automatically attempts to scale instances in the zone of the next highest-priority vSwitch.
	//
	// - COST_OPTIMIZED: during scale-out, Auto Scaling attempts to create ECS instances in ascending order of vCPU unit price. During scale-in, Auto Scaling attempts to remove ECS instances in descending order of vCPU unit price. If the scaling configuration settings specify multiple instance types that use the preemptible billing method, spot instances are preferentially created. You can use the CompensateWithOnDemand parameter to specify whether to automatically attempt to create pay-as-you-go instances when spot instances cannot be created due to insufficient inventory or other reasons.
	//
	//     > COST_OPTIMIZED takes effect only when the scaling configuration settings specify multiple instance types or use spot instances.
	//
	// - BALANCE: evenly distributes ECS instances across the zones specified for the scaling group. If the instances are unevenly allocated across zones due to insufficient inventory or other reasons, you can invoke the [RebalanceInstance](https://help.aliyun.com/document_detail/71516.html) operation to balance the resources.
	//
	//     > If the `MultiAZPolicy` parameter is set to `BALANCE`, the effect is the same as when the `MultiAZPolicy` parameter is set to `COMPOSABLE` and the `AzBalance` parameter is set to `true`.
	//
	// - COMPOSABLE: a combined policy that allows you to combine the preceding multi-zone scaling policies for ECS instances based on your needs. You can also specify additional parameters to control the capacity of the scaling group in a more fine-grained manner.
	//
	// Default value: PRIORITY.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than this value, pay-as-you-go instances are preferentially created.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances in the excess instances after the scaling group meets the minimum number of pay-as-you-go instances (OnDemandBaseCapacity). Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32  `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	OwnerAccount                        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The filtering policy for the instance removal policy. Valid values:
	//
	// - OldestInstance: removes the ECS instances that were added to the scaling group at the earliest point in time.
	//
	// - NewestInstance: removes the ECS instances that were added to the scaling group at the latest point in time.
	//
	// - OldestScalingConfiguration: removes the ECS instances that were created based on the earliest scaling configuration.
	//
	// - CustomPolicy: removes ECS instances based on a custom scale-in policy function.
	//
	// The scaling configuration mentioned in OldestScalingConfiguration refers to the instance configuration source in the scaling group, including scaling configurations and launch templates. CustomPolicy can only be set as the first removal policy. If you specify CustomPolicy, you must also specify the CustomPolicyARN parameter.
	//
	// > The removal of ECS instances from a scaling group is also affected by the multi-zone scaling policy (MultiAZPolicy) of the scaling group. For more information, see [Configure a combined policy for removing instances](https://help.aliyun.com/document_detail/254822.html).
	RemovalPolicies []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	// The ID of the resource group to which the newly created scaling group belongs.
	//
	// > Assigns the newly created scaling group to the specified resource group. If you do not specify this parameter, the scaling group belongs to the default resource group.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The name of the scaling group. The name must be unique within a region.
	//
	// It must be 2 to 64 characters in length. It must start with a digit, a letter, or a Chinese character, and can contain digits, underscores (_), hyphens (-), and periods (.).
	//
	// Default value: the value of ScalingGroupId.
	//
	// example:
	//
	// scalinggroup****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The reclaim mode of the scaling group. Valid values:
	//
	//
	//
	// - recycle: sets the reclaim mode of the scaling group to shutdown and reclaim mode.
	//
	// - release: sets the reclaim mode of the scaling group to release mode.
	//
	// - forcerelease: sets the reclaim mode of the scaling group to force release mode.
	//
	//     > If you select forcerelease, the system forcibly releases instances in the `Running` state during scale-down activities. Forcibly releasing an instance is equivalent to a power-off operation, which erases temporary data in the memory and storage of the instance. The data cannot be recovered. Proceed with caution.
	//
	// - forcerecycle: sets the reclaim mode of the scaling group to force shutdown and reclaim mode.
	//
	//     > If you select forcerecycle, the system forcibly shuts down instances in the `Running` state during scale-down activities. Forcibly shutting down an instance is equivalent to a power-off operation, which erases temporary data in the memory and storage of the instance. The data cannot be recovered. Proceed with caution.
	//
	//
	// The ScalingPolicy parameter specifies the reclaim mode of the scaling group, but the specific action taken when an instance is removed from the scaling group is determined by the RemovePolicy parameter of RemoveInstances. For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The collection of information about the server groups.
	//
	// > You cannot configure the same server group settings in both AlbServerGroups and ServerGroups at the same time.
	ServerGroups []*CreateScalingGroupRequestServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The spot capacity distribution policy. You can use this parameter to separately specify the distribution policy for spot capacity. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - priority: Creates instances based on the order of the configured instance types.
	//
	// - lowestPrice: Creates instances based on the unit price per vCPU of the instance types in ascending order.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of available instance types. The scaling group evenly creates spot instances across multiple instance types with the lowest cost. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Specifies whether to enable the Auto Scaling policy to supplement spot instances. After this feature is enabled, when a system message indicating that a spot instance is to be reclaimed is received, the scaling group attempts to create a new instance to replace the spot instance to be reclaimed.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The timeout period for waiting for an ECS instance to be stopped during the scale-in process. Unit: seconds.
	//
	// Valid values: 30 to 240.
	//
	// > - This parameter takes effect only in the scale-in process when ScalingPolicy is set to release.
	//
	// > - If this parameter is specified, the system waits for a maximum of the StopInstanceTimeout period for the instance to be stopped. After the timeout period elapses, the scale-in process continues regardless of whether the instance is stopped.
	//
	// > - If this parameter is not specified, the system waits for a long time for the instance to be stopped. The scale-in process continues only after the instance is stopped. If the instance fails to be stopped, the scale-in process is rolled back and the scale-in fails.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// > This parameter is not available.
	//
	// example:
	//
	// false
	SyncAlarmRuleToCms *bool `json:"SyncAlarmRuleToCms,omitempty" xml:"SyncAlarmRuleToCms,omitempty"`
	// The collection of tag information for the scaling group.
	Tags []*CreateScalingGroupRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The collection of information about the virtual server groups associated with the scaling group.
	VServerGroups []*CreateScalingGroupRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	// The ID of the vSwitch. After this parameter is specified, the network type of the scaling group is VPC.
	//
	// > If neither VSwitchId nor VSwitchIds is specified for the scaling group, the network type of the scaling group is classic network by default.
	//
	// example:
	//
	// vsw-bp14zolna43z266bq****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of one or more vSwitches. If you use the VSwitchIds parameter, the VSwitchId parameter is ignored. After this parameter is specified, the network type of the scaling group is VPC.
	//
	// When you specify multiple vSwitches:
	//
	// - The vSwitches must belong to the same VPC.
	//
	// - The vSwitches can belong to different zones.
	//
	// - The vSwitches are prioritized in ascending order of numbers. 1 indicates the highest priority. If instances cannot be created in the zone of a higher-priority vSwitch, the system automatically selects the next highest-priority vSwitch to create instances.
	//
	// > If neither VSwitchId nor VSwitchIds is specified for the scaling group, the network type of the scaling group is classic network by default.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The behavior policy of the scaling group during zone shift.
	ZoneShiftPolicy *CreateScalingGroupRequestZoneShiftPolicy `json:"ZoneShiftPolicy,omitempty" xml:"ZoneShiftPolicy,omitempty" type:"Struct"`
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

func (s *CreateScalingGroupRequest) GetZoneShiftPolicy() *CreateScalingGroupRequestZoneShiftPolicy {
	return s.ZoneShiftPolicy
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

func (s *CreateScalingGroupRequest) SetZoneShiftPolicy(v *CreateScalingGroupRequestZoneShiftPolicy) *CreateScalingGroupRequest {
	s.ZoneShiftPolicy = v
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
	if s.ZoneShiftPolicy != nil {
		if err := s.ZoneShiftPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScalingGroupRequestAlbServerGroups struct {
	// The ID of the ALB server group.
	//
	// A scaling group supports a limited number of associated ALB server groups. To view or manually apply for a quota increase, go to the [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port number used by the instance after Auto Scaling adds the instance to the ALB server group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The weight of the instance as a backend server after Auto Scaling adds the instance to the ALB server group. A higher weight specifies that more requests are distributed to the instance. If the weight is 0, the instance does not receive requests. Valid values: 0 to 100.
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
	// Specifies whether to allow the system to automatically attempt to create pay-as-you-go instances to meet the required number of ECS instances when the MultiAZPolicy parameter is set to COST_OPTIMIZED and sufficient spot instances cannot be created due to price or inventory reasons. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. When the number of pay-as-you-go instances in the scaling group is less than this value, the system preferentially creates pay-as-you-go instances. Valid values: 0 to 1000.
	//
	// When the MultiAZPolicy parameter is set to COMPOSABLE, the default value is 0.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances in the excess instances after the scaling group meets the minimum number of pay-as-you-go instances specified by OnDemandBaseCapacity. Valid values: 0 to 100.
	//
	// When the MultiAZPolicy parameter is set to COMPOSABLE, the default value is 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The price comparison mode for the cost optimization policy of the scaling group. Valid values:
	//
	// - PricePerUnit: Compares prices based on the price per unit of capacity.
	//
	//   The capacity of an instance in the scaling group is equal to the weight set for the instance type. The default weight is 1, which means one ECS instance occupies a capacity of 1.
	//
	// - PricePerVCpu: Compares prices based on the price per vCPU.
	//
	// Default value: PricePerUnit.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// Specifies whether to attempt to use spot instances to replace pay-as-you-go capacity when CompensateWithOnDemand is enabled and the proportion of pay-as-you-go instances exceeds the OnDemandPercentageAboveBaseCapacity value. In common scenarios where CompensateWithOnDemand downgrades spot instances to pay-as-you-go instances due to spot inventory or price reasons, this parameter attempts to use spot instances to replace the excess pay-as-you-go capacity to avoid long-term existence of pay-as-you-go instances. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed.
	//
	// Default value: false.
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
	// The association mode between the scaling group and the database. Valid values:
	//
	// - SecurityIp: Modifies the IP whitelist mode by automatically adding scaled-out instances to the IP whitelist of the database. Only ApsaraDB RDS databases are supported.
	//
	// - SecurityGroup: Uses the security group mode by adding the security group of the scaling configuration to the security group whitelist of the database to allow instances in the security group to access the database.
	//
	// example:
	//
	// SecurityIp
	AttachMode *string `json:"AttachMode,omitempty" xml:"AttachMode,omitempty"`
	// The database instance ID.
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
	// If you want the scaling group to scale based on the capacity of the instance type, specify both this parameter and LaunchTemplateOverrides.WeightedCapacity.
	//
	// This parameter is used to specify the instance type and overrides the instance type in the launch template.
	//
	// > This parameter takes effect only when the LaunchTemplateId parameter is specified.
	//
	// Valid values of InstanceType: available ECS instance types.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is used to specify the maximum price for the overridden instance type (namely, `LaunchTemplateOverride.InstanceType`) in the launch template.
	//
	// > This parameter takes effect only when the `LaunchTemplateId` parameter is specified.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// If you want the scaling group to scale based on the capacity of the instance type, specify this parameter after you specify LaunchTemplateOverrides.InstanceType.
	//
	// This parameter is used to specify the weight of the instance type, which indicates the capacity represented by a single instance of the instance type in the scaling group. A larger weight indicates that fewer instances of this instance type are required to meet the expected capacity.
	//
	// Because performance metrics such as the number of vCPUs and memory size vary among instance types, you can configure different weights for different instance types based on your needs.
	//
	// Example:
	//
	// - Current capacity: 0.
	//
	// - Expected capacity: 6.
	//
	// - Capacity of the ecs.c5.xlarge instance type: 4.
	//
	// To meet the expected capacity, the scaling group scales out two ecs.c5.xlarge instances for you.
	//
	// > During scale-out, the capacity of the scaling group cannot exceed the sum of the maximum capacity (MaxSize) and the maximum weight of the instance type.
	//
	// Valid values of WeightedCapacity: 1 to 500.
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
	// The action to be performed after the wait state ends. Valid values:
	//
	// - CONTINUE: continues the response to the scale-out activity or continues the response to the scale-down activity.
	//
	// - ABANDON: directly releases the instances created by the scale-out activity or directly removes the instances in the scale-down activity from the scaling group.
	//
	// When a scale-down activity (SCALE_IN) occurs in the scaling group and triggers multiple lifecycle hooks, if the wait state triggered by a lifecycle hook whose DefaultResult is ABANDON ends, the other corresponding wait states end in advance. In other cases, the next action is subject to the next action of the last wait state that ends.
	//
	// Default value: CONTINUE.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The wait time set by the lifecycle hook for the scaling group activity. After the wait state times out, the next action is performed. Valid values: 30 to 21600. Unit: seconds.
	//
	// After you create a lifecycle hook, you can call RecordLifecycleActionHeartbeat to extend the wait time of the instance, or call CompleteLifecycleAction to end the wait state of the scaling activity in advance.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The name of the lifecycle hook. The name cannot be modified after it is specified. If you do not specify this parameter, the default value is the same as the lifecycle hook ID.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The type of scaling activity to which the lifecycle hook applies. Valid values:
	//
	// - SCALE_OUT: scale-out activities of the scaling group.
	//
	// - SCALE_IN: scale-down activities of the scaling group.
	//
	// > If you specify a lifecycle hook for the scaling group, this parameter is required, and other related parameters are optional.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The identifier of the notification recipient for the lifecycle hook. Simple Message Queue (formerly MNS) queues or topics are supported. The parameter value format is acs:ess:{region}:{account-id}:{resource-relative-id}.
	//
	// - region: the region where the scaling group resides.
	//
	// - account-id: the Alibaba Cloud account ID.
	//
	// Examples:
	//
	// - Simple Message Queue (formerly MNS) queue: acs:ess:{region}:{account-id}:queue/{queuename}.
	//
	// - Simple Message Queue (formerly MNS) topic: acs:ess:{region}:{account-id}:topic/{topicname}.
	//
	// example:
	//
	// acs:ess:cn-hangzhou:1111111111:queue/queue2
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The fixed string information for the wait state of the scaling activity. The value cannot exceed 4,096 characters in length. Each time Auto Scaling pushes a message to the notification recipient, it also sends the predefined value of the NotificationMetadata parameter to help you manage and mark different types of notification information. The NotificationMetadata parameter takes effect only when you also specify the NotificationArn parameter.
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
	// The weight of the instance as a backend server after Auto Scaling adds the instance to the CLB server group. A higher weight specifies that more requests are distributed to the instance. If the weight is 0, the instance does not receive requests. Valid values: 0 to 100.
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
	// The port number used by the instance after Auto Scaling adds the instance to the server group. Valid values: 1 to 65535.
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
	// - ALB: Application Load Balancer (ALB).
	//
	// - NLB: Network Load Balancer (NLB).
	//
	// - GWLB: Gateway Load Balancer (GWLB).
	//
	// example:
	//
	// ALB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the instance as a backend server after Auto Scaling adds the instance to the server group. Valid values: 0 to 100.
	//
	// A higher weight specifies that more access requests are allocated to the instance. If the weight is 0, the instance does not receive access requests.
	//
	// > This parameter is required for ALB and NLB. You cannot configure this parameter for GWLB.
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
	// The tag key of the scaling group.
	//
	// example:
	//
	// Department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies whether the tag is a propagatable tag. Valid values:
	//
	// - true: tags on the scaling group are propagated only to newly created instances, not to existing instances in the scaling group.
	//
	// - false: tags on the scaling group are not propagated to instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Propagate *bool `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	// The tag value of the scaling group.
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
	// The ID of the Classic Load Balancer (CLB) instance to which the virtual server group belongs.
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
	// The port number used by the instance after Auto Scaling adds the instance to the virtual server group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the virtual server group.
	//
	// example:
	//
	// rsp-bp1443g77****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The weight of the instance as a backend server after Auto Scaling adds the instance to the virtual server group. A higher weight indicates that more access requests are allocated to the instance. If the weight is 0, the instance does not receive access requests. Valid values: 0 to 100.
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

type CreateScalingGroupRequestZoneShiftPolicy struct {
	// The health check behavior in impaired zones. Valid values:
	//
	// - ReplaceUnhealthy: Replaces unhealthy instances in all zones normally.
	//
	// - IgnoreUnhealthy: Does not replace unhealthy instances in impaired zones. Healthy zones are not affected.
	//
	// Default value: ReplaceUnhealthy.
	//
	// example:
	//
	// ReplaceUnhealthy
	ImpairedZoneHealthCheckBehavior *string `json:"ImpairedZoneHealthCheckBehavior,omitempty" xml:"ImpairedZoneHealthCheckBehavior,omitempty"`
	// Blocks scale-in triggered by specified types of scaling rules. The blocking scope applies to all zones under the scaling group, not just the impaired zones. This prevents false scale-in caused by fake improvements in monitoring metrics due to unreachable instances in impaired zones. Valid values:
	//
	// - StepScalingRule: Blocks scale-in triggered by step scaling rules.
	//
	// - TargetTrackingScalingRule: Blocks scale-in triggered by target tracking scaling rules.
	//
	// - None: Does not block.
	//
	// Default value: None.
	//
	// example:
	//
	// StepScalingRule
	ScaleInBlockedRuleTypes []*string `json:"ScaleInBlockedRuleTypes,omitempty" xml:"ScaleInBlockedRuleTypes,omitempty" type:"Repeated"`
	// Specifies whether scale-out creates instances only in healthy zones during zone shift. Valid values:
	//
	// - true: Scale-out creates instances only in healthy zones and skips impaired zones.
	//
	// - false: Scale-out is not restricted by zone and can create instances in any zone of the scaling group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	ScaleOutInHealthyZonesOnly *bool `json:"ScaleOutInHealthyZonesOnly,omitempty" xml:"ScaleOutInHealthyZonesOnly,omitempty"`
}

func (s CreateScalingGroupRequestZoneShiftPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingGroupRequestZoneShiftPolicy) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) GetImpairedZoneHealthCheckBehavior() *string {
	return s.ImpairedZoneHealthCheckBehavior
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) GetScaleInBlockedRuleTypes() []*string {
	return s.ScaleInBlockedRuleTypes
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) GetScaleOutInHealthyZonesOnly() *bool {
	return s.ScaleOutInHealthyZonesOnly
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) SetImpairedZoneHealthCheckBehavior(v string) *CreateScalingGroupRequestZoneShiftPolicy {
	s.ImpairedZoneHealthCheckBehavior = &v
	return s
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) SetScaleInBlockedRuleTypes(v []*string) *CreateScalingGroupRequestZoneShiftPolicy {
	s.ScaleInBlockedRuleTypes = v
	return s
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) SetScaleOutInHealthyZonesOnly(v bool) *CreateScalingGroupRequestZoneShiftPolicy {
	s.ScaleOutInHealthyZonesOnly = &v
	return s
}

func (s *CreateScalingGroupRequestZoneShiftPolicy) Validate() error {
	return dara.Validate(s)
}
