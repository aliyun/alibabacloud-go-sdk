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
	// The Application Load Balancer (ALB) server groups.
	AlbServerGroups []*CreateScalingGroupRequestAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The allocation policy of instances. Auto Scaling selects instance types based on the allocation policy to create the required number of instances. The policy can be applied to pay-as-you-go instances and preemptible instances. This parameter takes effect only when you set the `MultiAZPolicy` parameter to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling selects instance types based on the specified order of the instance types to create the required number of instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the lowest unit price of vCPUs to create the required number of instances.
	//
	// Default value: priority.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	AutoRebalance      *bool   `json:"AutoRebalance,omitempty" xml:"AutoRebalance,omitempty"`
	// Specifies whether to evenly distribute instances in the scaling group across multiple zones. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  If you set `MultiAZPolicy` to `COMPOSABLE` and enable `AzBalance` to `true`, this setting has an equivalent effect to setting `MultiAZPolicy` to `BALANCE`.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AzBalance   *bool   `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	BalanceMode *string `json:"BalanceMode,omitempty" xml:"BalanceMode,omitempty"`
	// The capacity options.
	CapacityOptions *CreateScalingGroupRequestCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the requirement on the number of ECS instances when the expected capacity of preemptible instances cannot be provided due to reasons such as cost-related issues and insufficient resources. This parameter is available only if you set the MultiAZPolicy parameter to COST_OPTIMIZED. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The ID of the elastic container instance.
	//
	// example:
	//
	// eci-uf6fonnghi50u374****
	ContainerGroupId *string `json:"ContainerGroupId,omitempty" xml:"ContainerGroupId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the custom scale-in policy (Function). This parameter is available only if you specify CustomPolicy as the first step to remove instances.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The IDs of the ApsaraDB RDS instances that you want to associate with the scaling group. The value can be a JSON array that contains multiple ApsaraDB RDS instance IDs. Separate multiple IDs with commas (,).
	//
	// You can associate only a limited number of ApsaraDB RDS instances with a scaling group. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the maximum number of ApsaraDB RDS instances that you can associate with a scaling group.
	//
	// example:
	//
	// ["rm-bp142f86de0t7****", "rm-bp18l1z42ar4o****", "rm-bp1lqr97h4aqk****"]
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The databases that you want to attach to the scaling group.
	DBInstances []*CreateScalingGroupRequestDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The cooldown period of the scaling group after a scaling activity is complete in the scaling group. Valid values: 0 to 86400. Unit: seconds.
	//
	// During the cooldown period, Auto Scaling does not execute scaling activities that are triggered by CloudMonitor event-triggered tasks.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The expected number of ECS instances in the scaling group. Auto Scaling automatically maintains the specified expected number of ECS instances. The DesiredCapacity value cannot be greater than the MaxSize value or less than the MinSize value.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Specifies whether to enable deletion protection for the scaling group. Valid values:
	//
	// 	- true: enables deletion protection for the scaling group. This way, the scaling group cannot be deleted.
	//
	// 	- false: disables deletion protection for the scaling group.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The type of the instances that are managed by the scaling group. Valid values:
	//
	// 	- ECS: ECS instances.
	//
	// 	- ECI: elastic container instances.
	//
	// Default value: ECS.
	//
	// example:
	//
	// ECS
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// 	- NONE: Auto Scaling does not check the health status of instances.
	//
	// 	- ECS: Auto Scaling checks the health status of instances in the scaling group. If you want to enable instance health check, you can set the value to ECS, regardless of whether the scaling group is of ECS type or Elastic Container Instance type.
	//
	// 	- LOAD_BALANCER: Auto Scaling checks the health status of instances in the scaling group based on the health check results of load balancers. The health check results of Classic Load Balancer (CLB) instances are not supported as the health check basis for instances in the scaling group.
	//
	// Default value: ECS.
	//
	// >  If you want to enable instance health check and load balancer health check at the same time, we recommend that you specify `HealthCheckTypes`.
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check mode of the scaling group.
	//
	// >  You can specify multiple values for this parameter to enable multiple health check options at the same time. If you specify `HealthCheckType`, this parameter is ignored.
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The ID of the ECS instance. When you create a scaling group, you can specify an existing ECS instance. Auto Scaling obtains the configurations of the ECS instance and automatically creates a scaling configuration from the obtained configurations.
	//
	// example:
	//
	// i-28wt4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the launch template that provides instance configurations for Auto Scaling to create instances.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// Details of the instance types that you specify by using the Extended Configurations feature of the launch template.
	LaunchTemplateOverrides []*CreateScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version number of the launch template. Valid values:
	//
	// 	- A fixed template version number.
	//
	// 	- Default: the default template version.
	//
	// 	- Latest: the latest template version.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The lifecycle hooks.
	LifecycleHooks []*CreateScalingGroupRequestLifecycleHooks `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Repeated"`
	// The load balancer configurations.
	LoadBalancerConfigs []*CreateScalingGroupRequestLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The IDs of the CLB instances that you want to associate with the scaling group. The value can be a JSON array that contains multiple CLB instance IDs. Separate multiple IDs with commas (,).
	//
	// You can associate only a limited number of CLB instances with a scaling group. Go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check the maximum number of CLB instances that you can associate with a scaling group.
	//
	// example:
	//
	// ["lb-bp1u7etiogg38yvwz****", "lb-bp168cqrux9ai9l7f****", "lb-bp1jv3m9zvj22ufxp****"]
	LoadBalancerIds *string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	// The maximum life span of an instance in the scaling group. Unit: seconds.
	//
	// Valid values: 86400 to the value of the Integer.maxValue parameter.
	//
	// Default value: null.
	//
	// example:
	//
	// null
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of instances that can be contained in the scaling group. When the total number of ECS instances in the scaling group exceeds the value of MaxSize, Auto Scaling automatically removes ECS instances from the scaling group until the total number equals the maximum number.
	//
	// The value range of MaxSize is directly correlated with the degree of dependency your business has on Auto Scaling. You can go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check **the maximum number of instances that a single scaling group can contain.**
	//
	// If **a single scaling group can contain up to 2,000 ECS instances**, the value range of MaxSize is 0 to 2,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of instances that must be contained in the scaling group. When the total number of ECS instances in the scaling group is less than the value of MinSize, Auto Scaling automatically creates ECS instances in the scaling group until the total number reaches the minimum number.
	//
	// >  The value of MinSize must be less than or equal to the value of MaxSize.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The scaling policy for ECS instances in the multi-zone scaling group. Valid values:
	//
	// 	- PRIORITY: scale ECS instances based on the priority of the vSwitches specified by VSwitchIds. Auto Scaling preferentially scales instances in the zone where the vSwitch of the highest priority resides. If the scaling fails, Auto Scaling scales instances in the zone where the vSwitch of the next highest priority resides.
	//
	// 	- COST_OPTIMIZED: create ECS instances that have the lowest unit price of vCPUs during scale-out events and removes ECS instances that have the highest unit price of vCPUs during scale-in events. If you specify preemptible instance types in your scaling configuration, Auto Scaling will preferentially create preemptible instances. You can also specify CompensateWithOnDemand to allow Auto Scaling to create pay-as-you-go instances in the case that preemptible instances cannot be created due to limited stock.
	//
	//     **
	//
	//     **Note*	- The COST_OPTIMIZED setting takes effect only when your scaling configuration contains multiple instance types or specifically contains preemptible instance types.
	//
	// 	- BALANCE: evenly distribute ECS instances across the zones that are specified for the scaling group. If ECS instances are unevenly distributed across the specified zones due to insufficient inventory, you can call the [RebalanceInstance](https://help.aliyun.com/document_detail/71516.html) operation to evenly distribute the instances across the zones.
	//
	//     **
	//
	//     **Note*	- When you set `MultiAZPolicy` to `BALANCE`, this setting has an equivalent effect to setting `MultiAZPolicy` to `COMPOSABLE` and enabling `AzBalance` to `true`.
	//
	// 	- COMPOSABLE: combine the preceding policies into a custom scaling policy based on your business requirements. Alternatively, you can specify custom parameters to finely control the capacity of the scaling group.
	//
	// Default value: PRIORITY.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be contained in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than the value of this parameter, Auto Scaling preferentially creates pay-as-you-go instances.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances in the excess instances when the minimum number of pay-as-you-go instances reaches the requirement. Valid values: 0 to 100.
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
	// The instance removal policies. Valid values:
	//
	// 	- OldestInstance: removes ECS instances that are added at the earliest point in time to the scaling group.
	//
	// 	- NewestInstance: removes ECS instances that are most recently added to the scaling group.
	//
	// 	- OldestScalingConfiguration: removes ECS instances that are created based on the earliest scaling configuration.
	//
	// 	- CustomPolicy: removes ECS instances based on the custom scale-in policy (Function).
	//
	// The scaling configuration source specified by the OldestScalingConfiguration setting can be a scaling configuration or a launch template. The CustomPolicy setting takes effect only if you specify it as the first step to remove instances. If you specify CustomPolicy, you must also specify the CustomPolicyARN parameter.
	//
	// > The removal of ECS instances from a scaling group is also affected by the value of the MultiAZPolicy parameter. For more information, see the [Configure a combination policy for removing instances](https://help.aliyun.com/document_detail/254822.html) topic.
	RemovalPolicies []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	// The ID of the resource group to which you want to add the scaling group.
	//
	// > If you specify this parameter, new scaling groups are added to the specified resource group. If you do not specify this parameter, new scaling groups are added to the default resource group.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The name of the scaling group. The name of each scaling group must be unique in a region.
	//
	// The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). The name must start with a letter or a digit.
	//
	// If you do not specify this parameter, the value of the ScalingGroupId parameter is used.
	//
	// example:
	//
	// scalinggroup****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The reclaim mode of the scaling group. Valid values:
	//
	// 	- recycle: the economical mode
	//
	// 	- release: the release mode
	//
	// 	- forcerelease: the forced release mode
	//
	//     **
	//
	//     **Note*	- If you set the value to forcerelease, Auto Scaling will forcibly release the ECS instances that are in the `Running` state during the scale-out events. Forced release equates to an immediate power-off, resulting in the irreversible deletion of all ephemeral data stored on the instance. Once executed, this action cannot be undone and the lost data cannot be recovered. Exercise caution when you select this option.
	//
	// 	- forcerecycle: the forced recycle mode
	//
	//     **
	//
	//     **Note*	- If you set the value to forcerecycle, Auto Scaling will forcibly shut down the ECS instances that are in the `Running` state during the scale-out events. Forced recycle equates to an immediate power-off, resulting in the irreversible deletion of all ephemeral data stored on the instance. Once executed, this action cannot be undone and the lost data cannot be recovered. Exercise caution when you select this option.
	//
	// ScalingPolicy specifies the reclaim mode of the scaling group. RemovePolicy of the RemoveInstances operation specifies the specific instance removal action. For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The server groups.
	//
	// >  You cannot use AlbServerGroups and ServerGroups to specify the same server group.
	ServerGroups []*CreateScalingGroupRequestServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The allocation policy of preemptible instances. You can use this parameter to individually specify the allocation policy of preemptible instances. This parameter takes effect only if you set the `MultiAZPolicy` parameter to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling selects instance types based on the specified order of the instance types to create the required number of preemptible instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the lowest unit price of vCPUs to create the required number of preemptible instances.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of available instance types. Auto Scaling evenly creates preemptible instances of multiple instance types that are provided at the lowest cost in the scaling group. Valid values: 1 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Specifies whether to supplement preemptible instances. If you set this parameter to true, Auto Scaling creates an instance to replace a preemptible instance when Auto Scaling receives a system message which indicates that the preemptible instance is to be reclaimed.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The period of time required by the ECS instance to enter the Stopped state. Unit: seconds. Valid values: 30 to 240.
	//
	// >
	//
	// 	- This parameter takes effect only if you set ScalingPolicy to release.
	//
	// 	- If you specify this parameter, the system will wait for the ECS instance to enter the Stopped state for the specified period of time before continuing with the scale-in operation, regardless of the status of the ECS instance.
	//
	// 	- If you do not specify this parameter, the system will wait for the ECS instance to stop before continuing with the scale-in operation. If the ECS instance is not successfully stopped, the scale-in process will be rolled back and considered failed.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// > This parameter is unavailable.
	//
	// example:
	//
	// false
	SyncAlarmRuleToCms *bool `json:"SyncAlarmRuleToCms,omitempty" xml:"SyncAlarmRuleToCms,omitempty"`
	// The tags that you want to add to the scaling group.
	Tags []*CreateScalingGroupRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The backend vServer group that you want to associate with the scaling group.
	VServerGroups []*CreateScalingGroupRequestVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	// The ID of the vSwitch. If you specify the VSwitchId parameter, the network type of the scaling group is VPC.
	//
	// > If you do not specify the VSwitchId or VSwitchIds parameter, the network type of the scaling group is classic network.
	//
	// example:
	//
	// vsw-bp14zolna43z266bq****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of the vSwitches. If you specify VSwitchIds, VSwitchId is ignored. If you specify VSwitchIds, the network type of the scaling group is VPC.
	//
	// If you specify multiple vSwitches, take note of the following items:
	//
	// 	- The vSwitches must belong to the same VPC.
	//
	// 	- The vSwitches can belong to different zones.
	//
	// 	- vSwitches are sorted in ascending order based on their priorities. The first vSwitch has the highest priority. If Auto Scaling fails to create ECS instances in the zone where the vSwitch of the highest priority resides, Auto Scaling attempts to create ECS instances in the zone where the vSwitch of the next highest priority resides.
	//
	// >  If you do not specify VSwitchId or VSwitchIds for your scaling group, the network type of the scaling group is classic network.
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
	// You can attach only a limited number of ALB server groups to a scaling group. To view the predefined quota limit or manually request a quota increase, go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas).
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port number used by each ECS instance as a backend server in the ALB server group. Valid values: 1 to 65535.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The weight of an ECS instance as a backend server in the ALB server group. If you increase the weight for an ECS instance, the number of requests that are forwarded to the ECS instance also increases. If you set the weight for an ECS instance to 0, no requests are forwarded to the ECS instance. Valid values: 0 to 100.
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
	// Specifies whether to automatically create pay-as-you-go ECS instances to reach the required number of ECS instances when preemptible ECS instances cannot be created due to high prices or insufficient inventory of resources. This parameter takes effect when you set `MultiAZPolicy` to `COST_OPTIMIZED`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. When the number of pay-as-you-go instances drops below the value of this parameter, Auto Scaling preferentially creates pay-as-you-go instances. Valid values: 0 to 1000.
	//
	// If you set `MultiAZPolicy` to `COMPOSABLE`, the default value is 0.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances in the excess instances when the minimum number of pay-as-you-go instances is reached. `OnDemandBaseCapacity` specifies the minimum number of pay-as-you-go instances that must be contained in the scaling group. Valid values: 0 to 100.
	//
	// If you set `MultiAZPolicy` to `COMPOSABLE`, the default value is 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The cost comparison method. Valid values:
	//
	// 	- PricePerUnit: Prices are compared based on the price per instance capacity.
	//
	//     Capacity is determined by the weights assigned to instance types in the scaling group. If no weight is specified, a default weight of 1 is used, meaning each instance is assigned a capacity of 1.
	//
	// 	- PricePerVCpu: Prices are compared based on the price per vCPU.
	//
	// Default value: PricePerUnit.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// Specifies whether to replace pay-as-you-go instances with preemptible instances. If you specify `CompensateWithOnDemand`, it may result in a higher percentage of pay-as-you-go instances compared to the value of `OnDemandPercentageAboveBaseCapacity`. In this scenario, Auto Scaling will try to deploy preemptible instances to replace the surplus pay-as-you-go instances. When `CompensateWithOnDemand` is specified, Auto Scaling creates pay-as-you-go instances if there are not enough preemptible instance types. To avoid keeping these pay-as-you-go ECS instances for long periods, Auto Scaling tries to replace them with preemptible instances as soon as enough of preemptible instance types become available. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The mode in which you want to attach the database to the scaling group. Valid values:
	//
	// 	- SecurityIp: the mode in which Auto Scaling automatically adds the private IP addresses of ECS instances to the IP address whitelist of the database during scale-out events. You can set the value to SecurityIp only if you set Type to RDS.
	//
	// 	- SecurityGroup: the mode in which Auto Scaling adds the security group of the applied scaling configuration to the security group whitelist of the database. This setting allows ECS instances created from the scaling configuration to access the database.
	//
	// example:
	//
	// SecurityIp
	AttachMode *string `json:"AttachMode,omitempty" xml:"AttachMode,omitempty"`
	// The database ID.
	//
	// example:
	//
	// rm-m5eqju85s45mu0***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database type. Valid values:
	//
	// 	- RDS
	//
	// 	- Redis
	//
	// 	- MongoDB
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
	// The instance type that you want to use to override the instance type that is specified in the launch template.
	//
	// If you want to scale instances based on the weighted capacities of the instances, you must specify both the InstanceType and WeightedCapacity parameters.
	//
	// > This parameter is available only if you specify the LaunchTemplateId parameter.
	//
	// You can use the InstanceType parameter to specify only instance types that are available for purchase.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum bid price of the instance type that is specified by the `InstanceType` parameter. You can specify 1 to 10 instance types by using the Extended Configurations feature of the launch template.
	//
	// > This parameter is available only if you specify the `LaunchTemplateId` parameter.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The weight of the instance type. The weight specifies the capacity of an instance of the specified instance type in the scaling group. If you want to scale instances based on the weighted capacities of the instances, you must specify the WeightedCapacity parameter after you specify the InstanceType parameter.
	//
	// A higher weight specifies that a smaller number of instances of the specified instance type are required to meet the expected capacity requirement.
	//
	// Performance metrics, such as the number of vCPUs and the memory size of each instance type, may vary. You can specify different weights for different instance types based on your business requirements.
	//
	// Example:
	//
	// 	- Current capacity: 0
	//
	// 	- Expected capacity: 6
	//
	// 	- Capacity of ecs.c5.xlarge: 4
	//
	// To meet the expected capacity requirement, Auto Scaling must create and add two ecs.c5.xlarge instances.
	//
	// > The capacity of the scaling group cannot exceed the sum of the maximum number of instances that is specified by the MaxSize parameter and the maximum weight of the instance types.
	//
	// Valid values of the WeightedCapacity parameter: 1 to 500.
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
	// The action that Auto Scaling performs when the lifecycle hook times out. Valid values:
	//
	// 	- CONTINUE: Auto Scaling continues to respond to the scaling request.
	//
	// 	- ABANDON: Auto Scaling releases ECS instances that are created during scale-out events, or removes ECS instances from the scaling group during scale-in events.
	//
	// If multiple lifecycle hooks in the scaling group are triggered during scale-in events, and you set DefaultResult to ABANDON for one of the lifecycle hooks, Auto Scaling immediately performs the action after the lifecycle hook whose DefaultResult is set to ABANDON times out. In this case, other lifecycle hooks time out ahead of schedule. In other cases, Auto Scaling performs the action only after all lifecycle hooks time out. The action that Auto Scaling performs is determined by the value of DefaultResult that you specify for the lifecycle hook that most recently times out.
	//
	// Default value: CONTINUE.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The period of time before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action that is specified by DefaultResult. Valid values: 30 to 21600. Unit: seconds.
	//
	// After you create a lifecycle hook, you can call the RecordLifecycleActionHeartbeat operation to extend the timeout period of the lifecycle hook. You can also call the CompleteLifecycleAction operation to end the timeout period of the lifecycle hook ahead of schedule.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The name of the lifecycle hook. After you specify this parameter, you cannot change the name of the lifecycle hook. If you do not specify this parameter, the name of the lifecycle hook is the same as the ID of the lifecycle hook.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The type of the scaling activity to which you want to apply the lifecycle hook. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// >  If you specify lifecycle hooks for the scaling group, you must specify LifecycleTransition. Other parameters are optional.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient party. You can specify a Simple Message Queue (SMQ, formerly MNS) topic or queue as the recipient party. The value is in the acs:ess:{region}:{account-id}:{resource-relative-id} format.
	//
	// 	- region: the region ID of the scaling group
	//
	// 	- account-id: the ID of your Alibaba Cloud account.
	//
	// Examples:
	//
	// 	- SMQ queue: acs:ess:{region}:{account-id}:queue/{queuename}
	//
	// 	- SMQ topic: acs:ess:{region}:{account-id}:topic/{topicname}
	//
	// example:
	//
	// acs:ess:cn-hangzhou:1111111111:queue/queue2
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The fixed string that you want to include in notifications. When a lifecycle hook takes effect, Auto Scaling sends a notification. The fixed string can contain up to 4,096 characters in length. When Auto Scaling sends a notification to the recipient party, it includes predefined notification metadata into the notification. This helps in managing and labeling notifications of different categories. NotificationMetadata takes effect only if you specify NotificationArn.
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
	// lb-2zen1olhfg9yw3f4qgte4
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of each ECS instance as a backend server in the CLB backend server group. If you increase the weight for an ECS instance, the number of requests that are forwarded to the ECS instance also increases. If you set the weight for an ECS instance to 0, no requests are forwarded to the ECS instance. Valid values: 0 to 100.
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
	// The port number used by each ECS instance as backend server in the vServer group. Valid values: 1 to 65535.
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
	// 	- ALB
	//
	// 	- NLB
	//
	// example:
	//
	// ALB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of each ECS instance as a backend server in the server group. Valid values: 0 to 100.
	//
	// If you increase the weight for an ECS instance, the number of requests that are forwarded to the ECS instance also increases. If you set the weight for an ECS instance to 0, no requests are forwarded to the ECS instance.
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
	// The tag key that you want to add to the scaling group.
	//
	// example:
	//
	// Department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies whether to propagate the tag that you want to add to the scaling group. Valid values:
	//
	// 	- true: propagates the tag to only instances that are newly created.
	//
	// 	- false: does not propagate the tag to any instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Propagate *bool `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	// The tag value that you want to add to the scaling group.
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
	// The ID of the CLB instance to which the backend vServer group belongs.
	//
	// example:
	//
	// lb-bp1u7etiogg38yvwz****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the backend vServer group.
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
	// The port number used by each ECS instance as a backend server in the vServer group. Valid values: 1 to 65535.
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
	// The weight of each ECS instance as a backend server in the vServer group. If you increase the weight for an ECS instance, the number of requests that are forwarded to the ECS instance also increases. If you set the weight for an ECS instance to 0, no requests are forwarded to the ECS instance. Valid values: 0 to 100.
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
