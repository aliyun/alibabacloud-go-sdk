// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveScalingConfigurationId(v string) *ModifyScalingGroupRequest
	GetActiveScalingConfigurationId() *string
	SetAllocationStrategy(v string) *ModifyScalingGroupRequest
	GetAllocationStrategy() *string
	SetAzBalance(v bool) *ModifyScalingGroupRequest
	GetAzBalance() *bool
	SetCapacityOptions(v *ModifyScalingGroupRequestCapacityOptions) *ModifyScalingGroupRequest
	GetCapacityOptions() *ModifyScalingGroupRequestCapacityOptions
	SetCompensateWithOnDemand(v bool) *ModifyScalingGroupRequest
	GetCompensateWithOnDemand() *bool
	SetCustomPolicyARN(v string) *ModifyScalingGroupRequest
	GetCustomPolicyARN() *string
	SetDefaultCooldown(v int32) *ModifyScalingGroupRequest
	GetDefaultCooldown() *int32
	SetDesiredCapacity(v int32) *ModifyScalingGroupRequest
	GetDesiredCapacity() *int32
	SetDisableDesiredCapacity(v bool) *ModifyScalingGroupRequest
	GetDisableDesiredCapacity() *bool
	SetGroupDeletionProtection(v bool) *ModifyScalingGroupRequest
	GetGroupDeletionProtection() *bool
	SetHealthCheckType(v string) *ModifyScalingGroupRequest
	GetHealthCheckType() *string
	SetHealthCheckTypes(v []*string) *ModifyScalingGroupRequest
	GetHealthCheckTypes() []*string
	SetLaunchTemplateId(v string) *ModifyScalingGroupRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateOverrides(v []*ModifyScalingGroupRequestLaunchTemplateOverrides) *ModifyScalingGroupRequest
	GetLaunchTemplateOverrides() []*ModifyScalingGroupRequestLaunchTemplateOverrides
	SetLaunchTemplateVersion(v string) *ModifyScalingGroupRequest
	GetLaunchTemplateVersion() *string
	SetMaxInstanceLifetime(v int32) *ModifyScalingGroupRequest
	GetMaxInstanceLifetime() *int32
	SetMaxSize(v int32) *ModifyScalingGroupRequest
	GetMaxSize() *int32
	SetMinSize(v int32) *ModifyScalingGroupRequest
	GetMinSize() *int32
	SetMultiAZPolicy(v string) *ModifyScalingGroupRequest
	GetMultiAZPolicy() *string
	SetOnDemandBaseCapacity(v int32) *ModifyScalingGroupRequest
	GetOnDemandBaseCapacity() *int32
	SetOnDemandPercentageAboveBaseCapacity(v int32) *ModifyScalingGroupRequest
	GetOnDemandPercentageAboveBaseCapacity() *int32
	SetOwnerAccount(v string) *ModifyScalingGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyScalingGroupRequest
	GetOwnerId() *int64
	SetRemovalPolicies(v []*string) *ModifyScalingGroupRequest
	GetRemovalPolicies() []*string
	SetResourceOwnerAccount(v string) *ModifyScalingGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyScalingGroupRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *ModifyScalingGroupRequest
	GetScalingGroupId() *string
	SetScalingGroupName(v string) *ModifyScalingGroupRequest
	GetScalingGroupName() *string
	SetScalingPolicy(v string) *ModifyScalingGroupRequest
	GetScalingPolicy() *string
	SetSpotAllocationStrategy(v string) *ModifyScalingGroupRequest
	GetSpotAllocationStrategy() *string
	SetSpotInstancePools(v int32) *ModifyScalingGroupRequest
	GetSpotInstancePools() *int32
	SetSpotInstanceRemedy(v bool) *ModifyScalingGroupRequest
	GetSpotInstanceRemedy() *bool
	SetStopInstanceTimeout(v int32) *ModifyScalingGroupRequest
	GetStopInstanceTimeout() *int32
	SetVSwitchIds(v []*string) *ModifyScalingGroupRequest
	GetVSwitchIds() []*string
}

type ModifyScalingGroupRequest struct {
	// The ID of the active scaling configuration in the scaling group.
	//
	// example:
	//
	// asc-bp17pelvl720x5ub****
	ActiveScalingConfigurationId *string `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	// The allocation policy. Auto Scaling selects instance types based on the allocation policy to create the required number of instances. The policy can be applied to pay-as-you-go instances and preemptible instances at the same time. This parameter takes effect only when you set the MultiAZPolicy parameter to COMPOSABLE. Valid values:
	//
	// 	- priority: Auto Scaling selects instance types based on the specified order to create the required number of instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the lowest unit price of vCPUs to create the required number of instances.
	//
	// Default value: priority.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Specifies whether to evenly distribute instances in the scaling group across zones. This parameter takes effect only when you set the `MultiAZPolicy` parameter to `COMPOSABLE`. Valid values:
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
	AzBalance *bool `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	// The capacity options.
	CapacityOptions *ModifyScalingGroupRequestCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the requirements on the number of ECS instances in the scaling group when the number of preemptible instances cannot be reached due to reasons such as cost-related issues and insufficient resources. This parameter takes effect only if you set `MultiAZPolicy` in the `CreateScalingGroup` operation to `COST_OPTIMIZED`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The ARN of the custom scaling policy (Function). This parameter takes effect only when you specify CustomPolicy as the first step of the instance removal policy.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The cooldown period of the scaling group. This parameter is available only if you set ScalingRuleType to SimpleScalingRule. Valid values: 0 to 86400. Unit: seconds.
	//
	// During the cooldown period, Auto Scaling does not execute scaling activities that are triggered by CloudMonitor event-triggered tasks.
	//
	// example:
	//
	// 600
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The expected number of ECS instances or elastic container instances in the scaling group. Auto Scaling maintains the expected number of ECS instances or elastic container instances in the scaling group. Make sure that you adhere to the following rule when specifying this parameter: Value of MaxSize ≥ Value of DesiredCapacity ≥ Value of MinSize
	//
	// >  If you re-enable the Expected Number of Instances feature, you must specify a value for `DesiredCapacity` again.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// example:
	//
	// false
	DisableDesiredCapacity *bool `json:"DisableDesiredCapacity,omitempty" xml:"DisableDesiredCapacity,omitempty"`
	// Specifies whether to enable deletion protection for the scaling group. Valid values:
	//
	// 	- true: enables deletion protection for the scaling group. This way, the scaling group cannot be deleted.
	//
	// 	- false: disables deletion protection for the scaling group.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// 	- NONE: Auto Scaling does not check the health status of instances.
	//
	// 	- ECS: Auto Scaling checks the health status of instances in the scaling group. If you want to enable instance health check, you can set the value to ECS, regardless of whether the scaling group is of ECS type or Elastic Container Instance type.
	//
	// 	- LOAD_BALANCER: Auto Scaling checks the health status of instances in the scaling group based on the health check results of load balancers. The health check results of Classic Load Balancer (CLB) instances are not supported as the health check basis for instances in the scaling group. Default value: ECS.
	//
	// >  If you want to enable instance health check and load balancer health check at the same time, we recommend that you specify `HealthCheckTypes`.
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check mode of the scaling group.
	//
	// >  You can specify multiple values for this parameter to enable multiple health check options at the same time. If you specify HealthCheckType, this parameter is ignored.
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The ID of the launch template that is used by Auto Scaling to create instances.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// Details of the instance types that are specified in the extended configurations of the launch template.
	LaunchTemplateOverrides []*ModifyScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version number of the launch template. Valid values:
	//
	// 	- A fixed template version number.
	//
	// 	- Default: The default template version is always used.
	//
	// 	- Latest: The latest template version is always used.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum life span of the instance in the scaling group. Unit: seconds.
	//
	// Valid values: 86400 to Integer.maxValue. ``You can also set this parameter to 0. A value of 0 indicates that the instance has an unlimited life span in the scaling group.
	//
	// Default value: null.
	//
	// > You cannot specify this parameter for scaling groups that manage elastic container instances or scaling groups whose ScalingPolicy is set to recycle.
	//
	// example:
	//
	// null
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of ECS instances or elastic container instances that can be contained in the scaling group. If the total number of instances in the scaling group is greater than the value of MaxSize, Auto Scaling proactively removes the surplus instances from the scaling group to restore the total number to match the maximum limit.
	//
	// The value range of MaxSize is directly correlated with the degree of dependency your business has on Auto Scaling. You can go to [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to check **the maximum number of instances that a single scaling group can contain.**
	//
	// For example, if a scaling group can contain up to **2,000*	- instances, the value range of MaxSize is 0 to 2000.
	//
	// example:
	//
	// 99
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of ECS instances or elastic container instances that must be contained in the scaling group. If the total number of instances in the scaling group is less than the value of MinSize, Auto Scaling proactively adds instances to the scaling group to ensure that the total number aligns with the minimum threshold.
	//
	// >  The value of MinSize must be less than or equal to the value of MaxSize.
	//
	// example:
	//
	// 1
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The scaling policy for the multi-zone scaling group that contains ECS instances. Valid values:
	//
	// 	- PRIORITY: ECS instances are scaled based on the vSwitch priority. The first vSwitch specified by using the VSwitchIds parameter has the highest priority. Auto Scaling preferentially scales instances in the zone where the vSwitch that has the highest priority resides. If the scaling fails, Auto Scaling scales instances in the zone where the vSwitch that has the next highest priority resides.
	//
	// 	- COST_OPTIMIZED: During a scale-out activity, Auto Scaling preferentially creates ECS instances of the instance type that has the lowest unit price of vCPU. During a scale-in activity, Auto Scaling preferentially removes ECS instances of the instance types that have the highest unit price of vCPU. Auto Scaling preferentially creates preemptible instances when preemptible instance types are specified in the scaling configuration. You can use the `CompensateWithOnDemand` parameter to specify whether to automatically create pay-as-you-go instances when Auto Scaling fails to create preemptible instances.
	//
	// > The `COST_OPTIMIZED` setting takes effect only when multiple instance types are specified or at least one instance type is specified for preemptible instances.
	//
	// 	- BALANCE: ECS instances are evenly distributed across zones that are specified in the scaling group. If ECS instances are unevenly distributed among zones due to insufficient resources, you can call the RebalanceInstance operation to evenly distribute the instances among the zones.
	//
	// 	- COMPOSABLE: You can flexibly combine the preceding policies based on your business requirements.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be included in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than the value of this parameter, Auto Scaling preferentially creates pay-as-you-go instances.
	//
	// If you set the `MultiAZPolicy` parameter to `COMPOSABLE` Policy, the default value is 0.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The expected percentage of pay-as-you-go instances in the excess instances when the minimum number of pay-as-you-go instances reaches the requirement. Valid values: 0 to 100.
	//
	// If you set the `MultiAZPolicy` parameter to `COMPOSABLE` Policy, the default value is 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32  `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	OwnerAccount                        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy that is used to remove ECS instances from the scaling group. Valid values:
	//
	// 	- OldestInstance: removes ECS instances that are added at the earliest point in time to the scaling group.
	//
	// 	- NewestInstance: removes ECS instances that are most recently added to the scaling group.
	//
	// 	- OldestScalingConfiguration: removes ECS instances that are created based on the earliest scaling configuration.
	RemovalPolicies      []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The name of the scaling group. The name of each scaling group must be unique in a region. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). The name must start with a letter or a digit.
	//
	// example:
	//
	// scalinggroup****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The reclaim mode of the scaling group. Valid values:
	//
	// 	- recycle: economical mode
	//
	// 	- release: release mode
	//
	// 	- forcerelease: forced release mode
	//
	//     **
	//
	//     **Note*	- If you set the value to `forcerelease`, Auto Scaling forcibly releases instances that are in the `Running` state during scale-ins. Forced release is equivalent to power outage. If an instance is forcibly released, ephemeral data on the instance will be cleared and cannot be recovered. Exercise caution when you select this option.
	//
	// 	- forcerecycle: forced recycle mode
	//
	//     **
	//
	//     **Note*	- If you set the value to `forcerecycle`, Auto Scaling forcibly shuts down instances that are in the `Running` state during scale-ins. Forced shutdown is equivalent to power outage. If an instance is forcibly shut down, ephemeral data on the instance will be cleared and cannot be recovered. Exercise caution when you select this option.
	//
	// ScalingPolicy specifies only the reclaim mode of the scaling group. RemovePolicy of the RemoveInstances operation specifies the manner how instances are removed from the scaling group. For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The allocation policy of preemptible instances. You can use this parameter to individually specify the allocation policy of preemptible instances. This parameter takes effect only when you set the `MultiAZPolicy` parameter to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling selects instance types based on the specified order to create the required number of preemptible instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the lowest unit price of vCPUs to create the required number of preemptible instances.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of instance types that you specify. Auto Scaling creates preemptible instances of multiple instance types that are provided at the lowest price. Valid values: 0 to 10.
	//
	// If you set the `MultiAZPolicy` parameter to `COMPOSABLE` Policy, the default value is 2.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Specifies whether to supplement preemptible instances. If this parameter is set to true, Auto Scaling creates an instance to replace a preemptible instance when Auto Scaling receives the system message that the preemptible instance is to be reclaimed.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The period of time that is required by the Elastic Compute Service (ECS) instance to enter the Stopped state during the scale-in process. Unit: seconds. Valid values: 30 to 240.
	//
	// >
	//
	// 	- This parameter takes effect only if you set ScalingPolicy to release.\\
	//
	//     If you specify this parameter, the system proceeds with the scale-in process only after the period of time specified by StopInstanceTimeout ends. In this case, the scale-in operation continues regardless of whether the ECS instance enters the Stopped state or not.\\
	//
	//     If you do not specify this parameter, the system proceeds with the scale-in process only after the ECS instance enters the Stopped state. If the ECS instance fails to enter the Stopped state, the scale-in process rolls back, and the scale-in operation is considered as failed.
	//
	// 	- When you call the ModifyScalingGroup operation, you can set the value to 0. In this case, the system ignores this parameter.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// The IDs of vSwitches.
	//
	// This parameter takes effect only when the network type of the scaling group is virtual private cloud (VPC). The specified vSwitches and the scaling group must reside in the same VPC.
	//
	// The vSwitches can reside in different zones. The vSwitches are sorted in ascending order. The first vSwitch specified by using the VSwitchIds parameter has the highest priority. If Auto Scaling fails to create ECS instances in the zone where the vSwitch that has the highest priority resides, Auto Scaling creates ECS instances in the zone where the vSwitch that has the next highest priority resides.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s ModifyScalingGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequest) GetActiveScalingConfigurationId() *string {
	return s.ActiveScalingConfigurationId
}

func (s *ModifyScalingGroupRequest) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *ModifyScalingGroupRequest) GetAzBalance() *bool {
	return s.AzBalance
}

func (s *ModifyScalingGroupRequest) GetCapacityOptions() *ModifyScalingGroupRequestCapacityOptions {
	return s.CapacityOptions
}

func (s *ModifyScalingGroupRequest) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *ModifyScalingGroupRequest) GetCustomPolicyARN() *string {
	return s.CustomPolicyARN
}

func (s *ModifyScalingGroupRequest) GetDefaultCooldown() *int32 {
	return s.DefaultCooldown
}

func (s *ModifyScalingGroupRequest) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *ModifyScalingGroupRequest) GetDisableDesiredCapacity() *bool {
	return s.DisableDesiredCapacity
}

func (s *ModifyScalingGroupRequest) GetGroupDeletionProtection() *bool {
	return s.GroupDeletionProtection
}

func (s *ModifyScalingGroupRequest) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *ModifyScalingGroupRequest) GetHealthCheckTypes() []*string {
	return s.HealthCheckTypes
}

func (s *ModifyScalingGroupRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *ModifyScalingGroupRequest) GetLaunchTemplateOverrides() []*ModifyScalingGroupRequestLaunchTemplateOverrides {
	return s.LaunchTemplateOverrides
}

func (s *ModifyScalingGroupRequest) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *ModifyScalingGroupRequest) GetMaxInstanceLifetime() *int32 {
	return s.MaxInstanceLifetime
}

func (s *ModifyScalingGroupRequest) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *ModifyScalingGroupRequest) GetMinSize() *int32 {
	return s.MinSize
}

func (s *ModifyScalingGroupRequest) GetMultiAZPolicy() *string {
	return s.MultiAZPolicy
}

func (s *ModifyScalingGroupRequest) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *ModifyScalingGroupRequest) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *ModifyScalingGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyScalingGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyScalingGroupRequest) GetRemovalPolicies() []*string {
	return s.RemovalPolicies
}

func (s *ModifyScalingGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyScalingGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyScalingGroupRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ModifyScalingGroupRequest) GetScalingGroupName() *string {
	return s.ScalingGroupName
}

func (s *ModifyScalingGroupRequest) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *ModifyScalingGroupRequest) GetSpotAllocationStrategy() *string {
	return s.SpotAllocationStrategy
}

func (s *ModifyScalingGroupRequest) GetSpotInstancePools() *int32 {
	return s.SpotInstancePools
}

func (s *ModifyScalingGroupRequest) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *ModifyScalingGroupRequest) GetStopInstanceTimeout() *int32 {
	return s.StopInstanceTimeout
}

func (s *ModifyScalingGroupRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ModifyScalingGroupRequest) SetActiveScalingConfigurationId(v string) *ModifyScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAllocationStrategy(v string) *ModifyScalingGroupRequest {
	s.AllocationStrategy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAzBalance(v bool) *ModifyScalingGroupRequest {
	s.AzBalance = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetCapacityOptions(v *ModifyScalingGroupRequestCapacityOptions) *ModifyScalingGroupRequest {
	s.CapacityOptions = v
	return s
}

func (s *ModifyScalingGroupRequest) SetCompensateWithOnDemand(v bool) *ModifyScalingGroupRequest {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetCustomPolicyARN(v string) *ModifyScalingGroupRequest {
	s.CustomPolicyARN = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDefaultCooldown(v int32) *ModifyScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDesiredCapacity(v int32) *ModifyScalingGroupRequest {
	s.DesiredCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDisableDesiredCapacity(v bool) *ModifyScalingGroupRequest {
	s.DisableDesiredCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetGroupDeletionProtection(v bool) *ModifyScalingGroupRequest {
	s.GroupDeletionProtection = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetHealthCheckType(v string) *ModifyScalingGroupRequest {
	s.HealthCheckType = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetHealthCheckTypes(v []*string) *ModifyScalingGroupRequest {
	s.HealthCheckTypes = v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateId(v string) *ModifyScalingGroupRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateOverrides(v []*ModifyScalingGroupRequestLaunchTemplateOverrides) *ModifyScalingGroupRequest {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *ModifyScalingGroupRequest) SetLaunchTemplateVersion(v string) *ModifyScalingGroupRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMaxInstanceLifetime(v int32) *ModifyScalingGroupRequest {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMaxSize(v int32) *ModifyScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMinSize(v int32) *ModifyScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMultiAZPolicy(v string) *ModifyScalingGroupRequest {
	s.MultiAZPolicy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOnDemandBaseCapacity(v int32) *ModifyScalingGroupRequest {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOnDemandPercentageAboveBaseCapacity(v int32) *ModifyScalingGroupRequest {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerId(v int64) *ModifyScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetRemovalPolicies(v []*string) *ModifyScalingGroupRequest {
	s.RemovalPolicies = v
	return s
}

func (s *ModifyScalingGroupRequest) SetResourceOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetResourceOwnerId(v int64) *ModifyScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupId(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupName(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingPolicy(v string) *ModifyScalingGroupRequest {
	s.ScalingPolicy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotAllocationStrategy(v string) *ModifyScalingGroupRequest {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotInstancePools(v int32) *ModifyScalingGroupRequest {
	s.SpotInstancePools = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetSpotInstanceRemedy(v bool) *ModifyScalingGroupRequest {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetStopInstanceTimeout(v int32) *ModifyScalingGroupRequest {
	s.StopInstanceTimeout = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetVSwitchIds(v []*string) *ModifyScalingGroupRequest {
	s.VSwitchIds = v
	return s
}

func (s *ModifyScalingGroupRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingGroupRequestCapacityOptions struct {
	// Specifies whether to automatically create pay-as-you-go ECS instances to reach the required number of ECS instances when preemptible ECS instances cannot be created due to high prices or insufficient inventory of resources. This parameter takes effect only if you set `MultiAZPolicy` in the `CreateScalingGroup` operation to `COST_OPTIMIZED`. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The percentage of additional pay-as-you-go instances beyond the minimum required by `OnDemandBaseCapacity` in the scaling group. Valid values: 0 to 100
	//
	// If you set `MultiAZPolicy` to `COMPOSABLE`, the default value is 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The price comparison mode. Valid values:
	//
	// 	- PricePerUnit: compares prices based on capacity.
	//
	//     The capacity of instances in a scaling group is determined by the weights of the instance types used. If no weight is specified, the default weight is 1, which specifies that each instance in the scaling group has a capacity of 1.
	//
	// 	- PricePerVCpu: compares prices based on the price per vCPU.
	//
	// Default value: PricePerUnit.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// Specifies whether to replace pay-as-you-go instances with preemptible instances. If you specify `CompensateWithOnDemand`, it may result in a higher percentage of pay-as-you-go instances compared to the value of `OnDemandPercentageAboveBaseCapacity`. In this case, Auto Scaling will try to deploy preemptible instances to replace the surplus pay-as-you-go instances. When `CompensateWithOnDemand` is specified, Auto Scaling creates pay-as-you-go instances if there are not enough preemptible instance types. To avoid keeping these pay-as-you-go ECS instances for long periods, Auto Scaling tries to replace them with preemptible instances as soon as enough of preemptible instance types become available. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	SpotAutoReplaceOnDemand *bool `json:"SpotAutoReplaceOnDemand,omitempty" xml:"SpotAutoReplaceOnDemand,omitempty"`
}

func (s ModifyScalingGroupRequestCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingGroupRequestCapacityOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequestCapacityOptions) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *ModifyScalingGroupRequestCapacityOptions) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *ModifyScalingGroupRequestCapacityOptions) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *ModifyScalingGroupRequestCapacityOptions) GetPriceComparisonMode() *string {
	return s.PriceComparisonMode
}

func (s *ModifyScalingGroupRequestCapacityOptions) GetSpotAutoReplaceOnDemand() *bool {
	return s.SpotAutoReplaceOnDemand
}

func (s *ModifyScalingGroupRequestCapacityOptions) SetCompensateWithOnDemand(v bool) *ModifyScalingGroupRequestCapacityOptions {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *ModifyScalingGroupRequestCapacityOptions) SetOnDemandBaseCapacity(v int32) *ModifyScalingGroupRequestCapacityOptions {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequestCapacityOptions) SetOnDemandPercentageAboveBaseCapacity(v int32) *ModifyScalingGroupRequestCapacityOptions {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequestCapacityOptions) SetPriceComparisonMode(v string) *ModifyScalingGroupRequestCapacityOptions {
	s.PriceComparisonMode = &v
	return s
}

func (s *ModifyScalingGroupRequestCapacityOptions) SetSpotAutoReplaceOnDemand(v bool) *ModifyScalingGroupRequestCapacityOptions {
	s.SpotAutoReplaceOnDemand = &v
	return s
}

func (s *ModifyScalingGroupRequestCapacityOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingGroupRequestLaunchTemplateOverrides struct {
	// The instance type. The instance type that you specify by using the InstanceType parameter overwrites the instance type that is specified in the launch template.
	//
	// If you want Auto Scaling to scale instances in the scaling group based on the instance type weight, you must specify both the InstanceType and WeightedCapacity parameters.
	//
	// > This parameter takes effect only after you specify the LaunchTemplateId parameter.
	//
	// You can use the InstanceType parameter to specify only instance types that are available for purchase.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The weight of the instance type. The weight specifies the capacity of a single instance of the specified instance type in the scaling group. If you want Auto Scaling to scale instances in the scaling group based on the weighted capacity of instances, you must specify the WeightedCapacity parameter after you specify the InstanceType parameter.
	//
	// A higher weight specifies that a smaller number of instances of the specified instance type are required to meet the expected capacity.
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
	// > The capacity of the scaling group cannot exceed the sum of the maximum number of instances that is specified by the MaxSize parameter and the maximum weight of the instance type.
	//
	// Valid values of the WeightedCapacity parameter: 1 to 500.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s ModifyScalingGroupRequestLaunchTemplateOverrides) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingGroupRequestLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) SetInstanceType(v string) *ModifyScalingGroupRequestLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) SetSpotPriceLimit(v float32) *ModifyScalingGroupRequestLaunchTemplateOverrides {
	s.SpotPriceLimit = &v
	return s
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) SetWeightedCapacity(v int32) *ModifyScalingGroupRequestLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *ModifyScalingGroupRequestLaunchTemplateOverrides) Validate() error {
	return dara.Validate(s)
}
