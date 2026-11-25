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
	SetAutoRebalance(v bool) *ModifyScalingGroupRequest
	GetAutoRebalance() *bool
	SetAzBalance(v bool) *ModifyScalingGroupRequest
	GetAzBalance() *bool
	SetBalanceMode(v string) *ModifyScalingGroupRequest
	GetBalanceMode() *string
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
	SetZoneShiftPolicy(v *ModifyScalingGroupRequestZoneShiftPolicy) *ModifyScalingGroupRequest
	GetZoneShiftPolicy() *ModifyScalingGroupRequestZoneShiftPolicy
}

type ModifyScalingGroupRequest struct {
	// The ID of the active scaling configuration in the scaling group.
	//
	// example:
	//
	// asc-bp17pelvl720x5ub****
	ActiveScalingConfigurationId *string `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	// The capacity allocation policy, which determines how the scaling group selects available instance types to meet the capacity requirements. The capacity distribution policy applies to both pay-as-you-go and spot capacities. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - priority: creates instances in the order of the configured instance types.
	//
	// - lowestPrice: creates instances in ascending order of the vCPU unit price of the instance types.
	//
	// Default value: priority.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Specifies whether to enable automatic balancing for the scaling group. This parameter takes effect only when BalancedOnly is enabled for a zone-balanced scaling group. Valid values:
	//
	// - false: disables automatic balancing for the scaling group.
	//
	// - true: enables automatic balancing for the scaling group. The scaling group automatically detects the capacity across zones. If the capacity is unbalanced across zones, the scaling group proactively scales across zones to rebalance the capacity.
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
	// - BalancedBestEffort: If resource creation fails in a zone, the system degrades to other zones to ensure best-effort delivery of resources.
	//
	// - BalancedOnly: If resource creation fails in a zone, the system does not degrade to other zones. The scale-out activity partially succeeds to avoid excessive imbalance of resources across different zones.
	//
	// Default value: BalancedBestEffort.
	//
	// example:
	//
	// BalancedBestEffort
	BalanceMode *string `json:"BalanceMode,omitempty" xml:"BalanceMode,omitempty"`
	// The capacity options.
	CapacityOptions *ModifyScalingGroupRequestCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// Specifies whether to automatically attempt to create pay-as-you-go instances to meet the required number of ECS instances if sufficient spot instances cannot be created due to price or inventory reasons when the MultiAZPolicy parameter of the CreateScalingGroup operation is set to COST_OPTIMIZED. Valid values:
	//
	// - true: allows the creation of pay-as-you-go instances.
	//
	// - false: does not allow the creation of pay-as-you-go instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the Function Compute function for the custom scale-in policy. This parameter takes effect only when the first removal policy configured in RemovalPolicies is CustomPolicy.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The default cooldown time of the scaling group. This parameter applies only to simple scaling rules. Valid values: 0 to 86400. Unit: seconds.
	//
	// During the cooldown time, the scaling group does not execute any other scaling activities. This applies only to scaling activities triggered by CloudMonitor event-triggered tasks.
	//
	// example:
	//
	// 600
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The expected number of ECS instances or elastic container instances in the scaling group. The scaling group automatically maintains the number of ECS instances or elastic container instances at the expected number. Valid values: MinSize to MaxSize.
	//
	// > If the expected number of instances feature is disabled and then re-enabled, you must reset the DesiredCapacity value.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Specifies whether to disable the expected number of instances feature for the scaling group. Valid values:
	//
	// - false: enables the expected number of instances feature.
	//
	// - true: disables the expected number of instances feature.
	//
	// > You can change the Settings of this parameter to true (that is, shutdown the expected number of instances feature) only when no scaling activity is in progress in the scaling group. When you shutdown the expected number of instances feature, the current DesiredCapacity property of the scaling group is cleared, but the current number of instances in the scaling group remains unchanged.
	//
	// example:
	//
	// false
	DisableDesiredCapacity *bool `json:"DisableDesiredCapacity,omitempty" xml:"DisableDesiredCapacity,omitempty"`
	// Specifies whether to enable deletion protection for the scaling group. Valid values:
	//
	// - true: enables deletion protection for the scaling group. In this case, the scaling group cannot be deleted.
	//
	// - false: disables deletion protection for the scaling group.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// - NONE: disables health checks.
	//
	// - ECS: performs health checks on instances in the scaling group. This value applies to both ECS and elastic container instance scaling groups.
	//
	// - LOAD_BALANCER: determines the health status of instances based on the health check results of load balancing. Classic Load Balancer (CLB) is not supported.
	//
	// Default value: ECS.
	//
	// > To enable both instance health checks and load balancer health checks, use the HealthCheckTypes parameter.
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check modes of the scaling group.
	//
	// > You can use this parameter to specify multiple values to enable multiple health check options. If the HealthCheckType parameter is configured, this parameter is ignored.
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The ID of the instance launch template, which is used to specify that the scaling group obtains launch configuration information from the instance launch template.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The extended instance type information of the launch template.
	LaunchTemplateOverrides []*ModifyScalingGroupRequestLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the instance launch template. Valid values:
	//
	//
	//
	// - A specific template version number.
	//
	// - Default: always uses the default version of the template.
	//
	// - Latest: always uses the latest version of the template.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The maximum lifetime of an instance in the scaling group. Unit: seconds.
	//
	// Valid values: 0 or [86400, Integer.maxValue]. A value of 0 indicates that the configured value of MaxInstanceLifetime is cleared (that is, MaxInstanceLifetime is not configured).
	//
	// Default value: null.
	//
	// > This parameter cannot be configured for scaling groups of the elastic container instance type or scaling groups that have the no fees for stopped instances feature enabled.
	//
	// example:
	//
	// 86400
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of ECS instances or elastic container instances in the scaling group. If the number of ECS instances or elastic container instances in the scaling group exceeds the value of MaxSize, Auto Scaling automatically removes ECS instances or elastic container instances.
	//
	// <props="china">The value range of MaxSize varies based on the usage of Auto Scaling. Go to the [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota value for the **maximum number of instances in a single scaling group**.
	//
	// <props="intl">The value range of MaxSize varies based on the usage of Auto Scaling. Go to the [Quota Center](https://quotas.console.aliyun.com/products/ess/quotas) to view the quota value for the **maximum number of instances in a single scaling group**.
	//
	// <props="partner">The value range of MaxSize varies based on the usage of Auto Scaling. Go to the Quota Center to view the quota value for the **maximum number of instances in a single scaling group**.
	//
	// For example, if the quota value for the **maximum number of instances in a single scaling group*	- is 2000, the value range of MaxSize is 0 to 2000.
	//
	// example:
	//
	// 99
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of ECS instances or elastic container instances in the scaling group. If the number of ECS instances or elastic container instances in the scaling group is less than the value of MinSize, Auto Scaling initiates the automatic creation of ECS instances or elastic container instances.
	//
	// > The value of MinSize must be less than or equal to the value of MaxSize.
	//
	// example:
	//
	// 1
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The scaling policy for ECS instances in a multi-zone scaling group. Valid values:
	//
	// - PRIORITY: scales the scaling group based on the priorities of vSwitches (VSwitchIds.N). Auto Scaling preferentially attempts to scale the scaling group in the zone where the vSwitch with the highest priority resides. If the scaling fails, Auto Scaling automatically attempts to scale the scaling group in the zone where the vSwitch with the next highest priority resides.
	//
	// - COST_OPTIMIZED: during scale-out, Auto Scaling attempts to create ECS instances in ascending order of vCPU unit price. During scale-in, Auto Scaling attempts to remove ECS instances in descending order of vCPU unit price. If the scaling configuration specifies multiple instance types that use the spot billing method, spot instances are preferentially created. You can use the CompensateWithOnDemand parameter to specify whether Auto Scaling automatically attempts to create pay-as-you-go instances when spot instances cannot be created.
	//
	// > COST_OPTIMIZED takes effect only when the scaling configuration specifies multiple instance types or uses spot instances.
	//
	// - BALANCE: evenly allocate ECS instances across the active zones specified for the scaling group. If the instances are unevenly distributed across zones due to insufficient inventory, you can call the RebalanceInstance operation to balance the resources.
	//
	// - COMPOSABLE: a composite policy that allows you to combine the preceding scaling policies for ECS instances in a multi-zone scaling group based on your needs.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances is less than this value, pay-as-you-go instances are preferentially created.
	//
	// Default value: 0 when the MultiAZPolicy parameter is set to COMPOSABLE.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances among the extra instances that exceed the minimum number of pay-as-you-go instances specified by OnDemandBaseCapacity. Valid values: 0 to 100.
	//
	// Default value: 100 when the MultiAZPolicy parameter is set to COMPOSABLE.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32  `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	OwnerAccount                        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The scale-in policy of the scaling group, which specifies the ECS instances to be removed. Valid values:
	//
	// - OldestInstance: removes the ECS instances that were added to the scaling group at the earliest point in time.
	//
	// - NewestInstance: removes the ECS instances that were added to the scaling group at the latest point in time.
	//
	// - OldestScalingConfiguration: removes the ECS instances that were created by using the earliest scaling configuration.
	RemovalPolicies      []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The name of the scaling group. The name must be unique within a region. The name must be 2 to 64 characters in length and start with a letter, digit, or Chinese character. It can contain digits, underscores (_), hyphens (-), and periods (.).
	//
	// example:
	//
	// scalinggroup****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The revoke pattern of the scaling group. Valid values:
	//
	// - recycle: The revoke pattern of the scaling group is shutdown and revoke pattern.
	//
	// - release: The revoke pattern of the scaling group is release pattern.
	//
	// - forcerelease: The revoke pattern of the scaling group is forcible release pattern.
	//
	//     > If you select forcerelease, the system forcibly releases instances in the Running state during scale-down activities. Forcibly releasing an instance is equivalent to a power-off operation, which erases temporary data in the memory and storage of the instance. The data cannot be recovered. Proceed with caution.
	//
	// - forcerecycle: The revoke pattern of the scaling group is forcible shutdown and revoke pattern.
	//
	//     > If you select forcerecycle, the system forcibly shuts down instances in the Running state during scale-down activities. Forcible shutdown is equivalent to a power-off operation, which erases temporary data in the memory and storage of the instance. The data cannot be recovered. Proceed with caution.
	//
	// The ScalingPolicy parameter specifies the revoke pattern of the scaling group. However, the specific action when an instance is removed from the scaling group is determined by the RemovePolicy parameter of the RemoveInstances operation. For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The spot capacity distribution policy. You can use this parameter to separately specify the distribution policy for spot capacity. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - priority: creates instances in the order of the configured instance types.
	//
	// - lowestPrice: creates instances in ascending order of the vCPU unit price of the instance types.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of available instance types. The scaling group evenly creates spot instances across multiple instance types with the lowest cost. Valid values: 0 to 10.
	//
	// Default value: 2 when the MultiAZPolicy parameter is set to COMPOSABLE.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Specifies whether to enable the spot instance replenishment feature. After this feature is enabled, when the scaling group receives a system message indicating that a spot instance will be reclaimed, the scaling group attempts to create a new instance to replace the spot instance to be reclaimed.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The timeout period for waiting for an ECS instance to be successfully stopped during the scale-in flow. Unit: seconds.
	//
	// Valid values: 30 to 240.
	//
	// > - This parameter takes effect only in the scale-in flow when ScalingPolicy is set to the release pattern.
	//
	// > - If the Settings of this parameter are configured, the system waits for a maximum of the StopInstanceTimeout period for the instance to be successfully stopped. After the timeout period elapses, the scale-in flow continues regardless of whether the instance is successfully stopped.
	//
	// > - If the Settings of this parameter are not configured, the system waits for a long time for the instance to be successfully stopped. The scale-in flow continues only after the instance is successfully stopped. If the instance stop has failed, the scale-in flow will rollback and the scale-in has failed.
	//
	// > - A value of 0 indicates that the parameter settings of this parameter are canceled.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// The IDs of one or more vSwitches.
	//
	// This parameter takes effect only when the network type of the scaling group is VPC. The VPC to which the specified vSwitches belong must be the same as the VPC to which the scaling group belongs.
	//
	// The vSwitches can reside in different zones. The priorities of vSwitches are sorted in ascending order of numbers. 1 indicates the highest priority. If ECS instances cannot be created in the zone where the vSwitch with a higher priority resides, the vSwitch with the next highest priority is automatically selected to create ECS instances.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The behavior policy of the scaling group during zone shift.
	ZoneShiftPolicy *ModifyScalingGroupRequestZoneShiftPolicy `json:"ZoneShiftPolicy,omitempty" xml:"ZoneShiftPolicy,omitempty" type:"Struct"`
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

func (s *ModifyScalingGroupRequest) GetAutoRebalance() *bool {
	return s.AutoRebalance
}

func (s *ModifyScalingGroupRequest) GetAzBalance() *bool {
	return s.AzBalance
}

func (s *ModifyScalingGroupRequest) GetBalanceMode() *string {
	return s.BalanceMode
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

func (s *ModifyScalingGroupRequest) GetZoneShiftPolicy() *ModifyScalingGroupRequestZoneShiftPolicy {
	return s.ZoneShiftPolicy
}

func (s *ModifyScalingGroupRequest) SetActiveScalingConfigurationId(v string) *ModifyScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAllocationStrategy(v string) *ModifyScalingGroupRequest {
	s.AllocationStrategy = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAutoRebalance(v bool) *ModifyScalingGroupRequest {
	s.AutoRebalance = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetAzBalance(v bool) *ModifyScalingGroupRequest {
	s.AzBalance = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetBalanceMode(v string) *ModifyScalingGroupRequest {
	s.BalanceMode = &v
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

func (s *ModifyScalingGroupRequest) SetZoneShiftPolicy(v *ModifyScalingGroupRequestZoneShiftPolicy) *ModifyScalingGroupRequest {
	s.ZoneShiftPolicy = v
	return s
}

func (s *ModifyScalingGroupRequest) Validate() error {
	if s.CapacityOptions != nil {
		if err := s.CapacityOptions.Validate(); err != nil {
			return err
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
	if s.ZoneShiftPolicy != nil {
		if err := s.ZoneShiftPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyScalingGroupRequestCapacityOptions struct {
	// Specifies whether to automatically attempt to create pay-as-you-go instances to meet the required number of ECS instances if sufficient spot instances cannot be created due to price or inventory reasons when the MultiAZPolicy parameter of the CreateScalingGroup operation is set to COST_OPTIMIZED. Valid values:
	//
	// - true: allows the creation of pay-as-you-go instances.
	//
	// - false: does not allow the creation of pay-as-you-go instances.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. If the number of pay-as-you-go instances in the scaling group is less than this value, the system preferentially creates pay-as-you-go instances. Valid values: 0 to 1000.
	//
	// Default value: 0 when the MultiAZPolicy parameter is set to COMPOSABLE.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances among the extra instances that exceed the minimum number of pay-as-you-go instances specified by OnDemandBaseCapacity. Valid values: 0 to 100.
	//
	// Default value: 100 when the MultiAZPolicy parameter is set to COMPOSABLE.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The price comparison mode for the cost optimization policy of the scaling group. Valid values:
	//
	// - PricePerUnit: compares prices based on the price per unit of capacity.
	//
	//   The capacity of an instance in the scaling group equals the weight configured for the instance type. The default weight is 1, which means one ECS instance occupies a capacity of 1.
	//
	// - PricePerVCpu: compares prices based on the price per vCPU.
	//
	// Default value: PricePerUnit.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// Specifies whether to attempt to use spot instances to replace pay-as-you-go capacity when the percentage of pay-as-you-go instances exceeds the OnDemandPercentageAboveBaseCapacity value after CompensateWithOnDemand is enabled. In common scenarios where CompensateWithOnDemand downgrades spot instances to pay-as-you-go instances due to spot inventory or price reasons, this feature attempts to use spot instances to replace the excess pay-as-you-go capacity to avoid long-term retention of pay-as-you-go instances. Valid values:
	//
	// - true: allows the replacement.
	//
	// - false: does not allow the replacement.
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
	// When you need the scaling group to scale based on the instance type capacity, specify both this parameter and LaunchTemplateOverrides.WeightedCapacity.
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
	// This parameter is used to specify the maximum hourly price for the instance type that overrides the launch template (that is, LaunchTemplateOverride.N.InstanceType). You can specify N values for this parameter to extend the launch template to support N instance types. Valid values of N: 1 to 20.
	//
	// > This parameter takes effect only when the LaunchTemplateId parameter is specified.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// When you need the scaling group to scale based on the instance type capacity, specify this parameter after you specify LaunchTemplateOverrides.InstanceType. The two parameters correspond to each other.
	//
	// This parameter is used to specify the weight of the instance type, which indicates the capacity that a single instance of the instance type represents in the scaling group.
	//
	// A greater weight indicates that fewer instances of this instance type are required to meet the expected capacity.
	//
	// Because performance metrics such as the number of vCPUs and memory size vary with instance types, you can configure different weights for different instance types based on your needs.
	//
	// Example:
	//
	// - Current capacity: 0.
	//
	// - Expected capacity: 6.
	//
	// - Capacity of the ecs.c5.xlarge instance type: 4.
	//
	// To meet the expected capacity, the scaling group scales out two ecs.c5.xlarge instances.
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

type ModifyScalingGroupRequestZoneShiftPolicy struct {
	// The health check behavior in impaired zones. Valid values:
	//
	// - ReplaceUnhealthy: normally replaces unhealthy instances in all zones.
	//
	// - IgnoreUnhealthy: does not replace unhealthy instances in impaired zones. Healthy zones are not affected.
	//
	// example:
	//
	// ReplaceUnhealthy
	ImpairedZoneHealthCheckBehavior *string `json:"ImpairedZoneHealthCheckBehavior,omitempty" xml:"ImpairedZoneHealthCheckBehavior,omitempty"`
	// Blocks scale-in activities triggered by specified types of scaling rules. The blocking scope covers all zones under the scaling group, not just the impaired zones. This prevents false scale-in activities caused by fake improvements in monitoring metrics due to unreachable instances in impaired zones. Valid values:
	//
	// - StepScalingRule: blocks scale-in activities triggered by step scaling rules.
	//
	// - TargetTrackingScalingRule: blocks scale-in activities triggered by target tracking scaling rules.
	//
	// - None: does not block any scale-in activities.
	//
	// example:
	//
	// StepScalingRule
	ScaleInBlockedRuleTypes []*string `json:"ScaleInBlockedRuleTypes,omitempty" xml:"ScaleInBlockedRuleTypes,omitempty" type:"Repeated"`
	// Specifies whether scale-out creates instances only in healthy zones during zone shift. Valid values:
	//
	// - true: creates instances only in healthy zones during scale-out and skips impaired zones.
	//
	// - false: does not restrict zones, and scale-out can create instances in any zone under the scaling group.
	//
	// example:
	//
	// true
	ScaleOutInHealthyZonesOnly *bool `json:"ScaleOutInHealthyZonesOnly,omitempty" xml:"ScaleOutInHealthyZonesOnly,omitempty"`
}

func (s ModifyScalingGroupRequestZoneShiftPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingGroupRequestZoneShiftPolicy) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) GetImpairedZoneHealthCheckBehavior() *string {
	return s.ImpairedZoneHealthCheckBehavior
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) GetScaleInBlockedRuleTypes() []*string {
	return s.ScaleInBlockedRuleTypes
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) GetScaleOutInHealthyZonesOnly() *bool {
	return s.ScaleOutInHealthyZonesOnly
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) SetImpairedZoneHealthCheckBehavior(v string) *ModifyScalingGroupRequestZoneShiftPolicy {
	s.ImpairedZoneHealthCheckBehavior = &v
	return s
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) SetScaleInBlockedRuleTypes(v []*string) *ModifyScalingGroupRequestZoneShiftPolicy {
	s.ScaleInBlockedRuleTypes = v
	return s
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) SetScaleOutInHealthyZonesOnly(v bool) *ModifyScalingGroupRequestZoneShiftPolicy {
	s.ScaleOutInHealthyZonesOnly = &v
	return s
}

func (s *ModifyScalingGroupRequestZoneShiftPolicy) Validate() error {
	return dara.Validate(s)
}
