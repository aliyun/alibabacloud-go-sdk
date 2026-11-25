// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeScalingGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeScalingGroupsResponseBody
	GetRequestId() *string
	SetScalingGroups(v []*DescribeScalingGroupsResponseBodyScalingGroups) *DescribeScalingGroupsResponseBody
	GetScalingGroups() []*DescribeScalingGroupsResponseBodyScalingGroups
	SetTotalCount(v int32) *DescribeScalingGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeScalingGroupsResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling groups.
	ScalingGroups []*DescribeScalingGroupsResponseBodyScalingGroups `json:"ScalingGroups,omitempty" xml:"ScalingGroups,omitempty" type:"Repeated"`
	// The total number of scaling groups.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingGroupsResponseBody) GetScalingGroups() []*DescribeScalingGroupsResponseBodyScalingGroups {
	return s.ScalingGroups
}

func (s *DescribeScalingGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScalingGroupsResponseBody) SetPageNumber(v int32) *DescribeScalingGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetPageSize(v int32) *DescribeScalingGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetRequestId(v string) *DescribeScalingGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetScalingGroups(v []*DescribeScalingGroupsResponseBodyScalingGroups) *DescribeScalingGroupsResponseBody {
	s.ScalingGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetTotalCount(v int32) *DescribeScalingGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) Validate() error {
	if s.ScalingGroups != nil {
		for _, item := range s.ScalingGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingGroupsResponseBodyScalingGroups struct {
	// The number of ECS instances that have been added to the scaling group and are running normally.
	//
	// example:
	//
	// 1
	ActiveCapacity *int32 `json:"ActiveCapacity,omitempty" xml:"ActiveCapacity,omitempty"`
	// The ID of the active scaling configuration in the scaling group.
	//
	// example:
	//
	// asc-bp1et2qekq3ojr33****
	ActiveScalingConfigurationId *string `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	// The collection of information about Application Load Balancer (ALB) server groups.
	AlbServerGroups []*DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The capacity allocation strategy, which determines how the scaling group selects available instance types to meet the capacity requirements. The capacity distribution strategy applies to both pay-as-you-go and spot capacity. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - priority: Creates instances based on the order of the configured instance types.
	//
	// - lowestPrice: Creates instances based on the unit vCPU price of the instance types in ascending order.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Specifies whether to enable automatic balancing for the scaling group. This parameter takes effect only when BalancedOnly is enabled for the zone balance scaling group. Valid values:
	//
	// - false: Automatic balancing for the scaling group is disabled.
	//
	// - true: Automatic balancing for the scaling group is enabled. The scaling group automatically detects the capacity across zones. If the capacity is unbalanced, the scaling group proactively scales in or out to rebalance the capacity across zones.
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
	// example:
	//
	// false
	AzBalance *bool `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	// The zone balance mode. This parameter takes effect only when the zone balance mode is enabled. Valid values:
	//
	//
	//
	// - BalancedBestEffort: If resource creation fails in a zone, the system falls back to other zones to ensure best-effort delivery of resources.
	//
	// - BalancedOnly: If resource creation fails in a zone, the system does not fall back to other zones. The scale-out activity partially succeeds to avoid excessive imbalance of resources across different zones.
	//
	// example:
	//
	// BalancedBestEffort
	BalanceMode *string `json:"BalanceMode,omitempty" xml:"BalanceMode,omitempty"`
	// The capacity options.
	CapacityOptions *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// Specifies whether to automatically create pay-as-you-go instances to meet the requirement for the number of ECS instances when the MultiAZPolicy parameter is set to COST_OPTIMIZED and spot instances cannot be created due to price or inventory issues. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The time when the scaling group was created. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// > This parameter is not available.
	//
	// example:
	//
	// hostname
	CurrentHostName *string `json:"CurrentHostName,omitempty" xml:"CurrentHostName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the Function Compute function for the custom scale-in policy. This parameter takes effect only when the first removal policy in RemovalPolicies is set to CustomPolicy.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The IDs of the ApsaraDB RDS instances associated with the scaling group.
	DBInstanceIds []*string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
	// The list of databases associated with the scaling group.
	DBInstances []*DescribeScalingGroupsResponseBodyScalingGroupsDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The default cooldown time of the scaling group. During the cooldown time, the scaling group does not execute any other scaling activity. This parameter takes effect only for a scaling activity triggered by [CloudMonitor](https://help.aliyun.com/document_detail/35170.html) event-triggered tasks.
	//
	// example:
	//
	// 60
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The expected number of ECS instances in the scaling group. The scaling group automatically maintains the number of ECS instances at the expected number of instances.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Specifies whether the expected number of instances feature is enabled. Valid values:
	//
	// - true: The expected number of instances feature is enabled.
	//
	// - false: The expected number of instances feature is disabled.
	//
	// example:
	//
	// false
	EnableDesiredCapacity *bool `json:"EnableDesiredCapacity,omitempty" xml:"EnableDesiredCapacity,omitempty"`
	// Indicates whether deletion protection is enabled for the scaling group. Valid values:
	//
	// - true: Deletion protection is enabled. The scaling group cannot be deleted.
	//
	// - false: Deletion protection is disabled.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The type of instances managed by the scaling group.
	//
	// example:
	//
	// ECS
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// - NONE: No health check is performed.
	//
	// - ECS: Health checks are performed on instances in the scaling group. This value is used to enable instance health checks for both ECS and ECI scaling groups.
	//
	// - LOAD_BALANCER: The health status of instances is determined based on the health check results of load balancing (CLB load balancing is not supported).
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// - NONE: No health check is performed.
	//
	// - ECS: Health checks are performed on instances in the scaling group. This value is used to enable instance health checks for both ECS and ECI scaling groups.
	//
	// - LOAD_BALANCER: The health status of instances is determined based on the load balancing health check results (the CLB load balancing type is not currently supported).
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The number of instances in the initialization state before scale-out in the scaling group.
	//
	// example:
	//
	// 0
	InitCapacity *int32 `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	// > This parameter is not available.
	//
	// example:
	//
	// false
	IsElasticStrengthInAlarm *bool `json:"IsElasticStrengthInAlarm,omitempty" xml:"IsElasticStrengthInAlarm,omitempty"`
	// The ID of the launch template used by the scaling group.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The instance type information that extends the launch template.
	LaunchTemplateOverrides []*DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the launch template used by the scaling group.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The lifecycle state of the scaling group. Valid values:
	//
	// - Active: The scaling group is active. Only active scaling groups can accept requests to execute scaling rules and trigger a scaling activity.
	//
	// - Inactive: The scaling group is inactive. Inactive scaling groups do not receive any requests to execute scaling rules.
	//
	// - Deleting: The scaling group is being deleted. Scaling groups in the deleting state do not receive any requests to execute scaling rules, and their parameters cannot be modified.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The list of load balancing configurations.
	LoadBalancerConfigs []*DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The list of IDs of the load balancing instances associated with the scaling group.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum lifetime of an instance in the scaling group. Unit: seconds.
	//
	// Valid values: 0 or [86400, Integer.maxValue]. A value of 0 indicates that the configured value of MaxInstanceLifetime is cleared, which means MaxInstanceLifetime is not configured.
	//
	// Default value: null.
	//
	// >This parameter is not supported for scaling groups of the Elastic Container Instance (ECI) type or scaling groups that have the No Fees for Stopped Instances feature enabled.
	//
	// example:
	//
	// 86400
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of ECS instances in the scaling group.
	//
	// example:
	//
	// 2
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of ECS instances in the scaling group.
	//
	// example:
	//
	// 1
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The time when the scaling group was modified. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The ID of the CloudMonitor application group associated with the scaling group.
	//
	// example:
	//
	// 1497****
	MonitorGroupId *string `json:"MonitorGroupId,omitempty" xml:"MonitorGroupId,omitempty"`
	// The scale-out and scale-in policy for ECS instances in a multi-zone scaling group. Valid values:
	//
	// - PRIORITY: Scales out and scales in instances based on the vSwitches defined in VSwitchIds. If ECS instances cannot be created in the active zone of the vSwitch with the highest priority, the system automatically uses the vSwitch with the next highest priority to create ECS instances.
	//
	// - COST_OPTIMIZED: Attempts to create instances based on the unit price of vCPUs from low to high. When the scaling configuration specifies multiple instance types with the spot billing method, the system preferentially creates spot instances. You can use the CompensateWithOnDemand parameter to specify whether to automatically attempt to create pay-as-you-go instances when spot instances cannot be created due to inventory issues.
	//
	//     > COST_OPTIMIZED takes effect only when the scaling configuration specifies multiple instance types or uses spot instances.
	//
	// - BALANCE: Evenly allocates ECS instances across the active zones specified for the scaling group. If the distribution becomes unbalanced due to insufficient inventory, you can call the RebalanceInstance operation to balance the resources.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 1000. When the number of pay-as-you-go instances in the scaling group is less than this value, the system preferentially creates pay-as-you-go instances.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances to be created when the scaling group exceeds the minimum number of pay-as-you-go instances specified by OnDemandBaseCapacity. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The number of ECS instances that are being added to the scaling group but have not completed the related configurations.
	//
	// example:
	//
	// 0
	PendingCapacity *int32 `json:"PendingCapacity,omitempty" xml:"PendingCapacity,omitempty"`
	// The number of ECS instances that are in the pending wait state in the scaling group.
	//
	// example:
	//
	// 1
	PendingWaitCapacity *int32 `json:"PendingWaitCapacity,omitempty" xml:"PendingWaitCapacity,omitempty"`
	// The number of ECS instances that are in the protected state in the scaling group.
	//
	// example:
	//
	// 1
	ProtectedCapacity *int32 `json:"ProtectedCapacity,omitempty" xml:"ProtectedCapacity,omitempty"`
	// The ID of the region to which the scaling group belongs.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The collection of policies for removing ECS instances from the scaling group. Valid values:
	//
	// - OldestInstance: Removes the ECS instance that was created earliest.
	//
	// - NewestInstance: Removes the ECS instance that was created latest.
	//
	// - OldestScalingConfiguration: Removes the ECS instance that was created earliest based on the scaling configuration.
	RemovalPolicies []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	// The number of ECS instances that are being removed from the scaling group.
	//
	// example:
	//
	// 0
	RemovingCapacity *int32 `json:"RemovingCapacity,omitempty" xml:"RemovingCapacity,omitempty"`
	// The number of ECS instances that are in the removing wait state in the scaling group.
	//
	// example:
	//
	// 1
	RemovingWaitCapacity *int32 `json:"RemovingWaitCapacity,omitempty" xml:"RemovingWaitCapacity,omitempty"`
	// The ID of the resource group to which the scaling group belongs.
	//
	// example:
	//
	// rg-123****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The name of the scaling group.
	//
	// example:
	//
	// dyrSuvBOtO1dEdIlIbp****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The reclaim mode of the scaling group. Valid values:
	//
	// - recycle: The scaling group uses the shutdown and reclaim mode.
	//
	// - release: The scaling group uses the release mode.
	//
	// - forcerelease: The scaling group uses the force release mode.
	//
	// For more information about the actions on removed instances, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The collection of information about load balancing server groups.
	//
	// > You can use this parameter to obtain the collection of information about ALB service groups, NLB server groups, and GWLB service groups associated with the scaling group.
	//
	// >
	ServerGroups []*DescribeScalingGroupsResponseBodyScalingGroupsServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The spot capacity distribution strategy. You can use this parameter to separately specify the distribution strategy for spot capacity. This parameter takes effect only when the MultiAZPolicy parameter is set to COMPOSABLE. Valid values:
	//
	// - priority: Creates instances based on the order of the configured instance types.
	//
	// - lowestPrice: Creates instances based on the unit vCPU price of the instance types in ascending order.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of spot instances in the scaling group.
	//
	// example:
	//
	// 0
	SpotCapacity *int32 `json:"SpotCapacity,omitempty" xml:"SpotCapacity,omitempty"`
	// The number of available instance types specified. The scaling group evenly creates spot instances based on the most cost-effective instance types. Valid values: 0 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Indicates whether the spot instance replacement feature is enabled. After this feature is enabled, when the scaling group receives a system message indicating that a spot instance will be reclaimed, the scaling group attempts to create a new instance to replace the spot instance to be reclaimed.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The number of instances that are in the standby state in the scaling group.
	//
	// example:
	//
	// 1
	StandbyCapacity *int32 `json:"StandbyCapacity,omitempty" xml:"StandbyCapacity,omitempty"`
	// The timeout period for waiting for an ECS instance to be stopped during the scale-in process. Unit: seconds.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// The number of instances in the scaling group that are in the economical mode.
	//
	// example:
	//
	// 1
	StoppedCapacity *int32 `json:"StoppedCapacity,omitempty" xml:"StoppedCapacity,omitempty"`
	// The suspended processes. If no process is suspended, an empty value is returned. Valid values:
	//
	// - ScaleIn: The scale-in process.
	//
	// - ScaleOut: The scale-out process.
	//
	// - HealthCheck: The health check process.
	//
	// - AlarmNotification: The event-triggered task process.
	//
	// - ScheduledAction: The scheduled task process.
	SuspendedProcesses []*string `json:"SuspendedProcesses,omitempty" xml:"SuspendedProcesses,omitempty" type:"Repeated"`
	// Indicates whether Auto Scaling stops executing any scaling activity in the scaling group. Valid values:
	//
	// - true: Auto Scaling stops executing any scaling activity in the scaling group. This indicates that the scaling group has failed to execute scale-out or scale-in for more than seven consecutive days. You must modify the scaling group or scaling configuration to resume the execution of the scaling activity.
	//
	// - false: Auto Scaling executes the scaling activity in the scaling group.
	//
	// example:
	//
	// true
	SystemSuspended *bool `json:"SystemSuspended,omitempty" xml:"SystemSuspended,omitempty"`
	// The tags of the scaling group.
	Tags []*DescribeScalingGroupsResponseBodyScalingGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total weighted capacity of all ECS instances in the scaling group if instance type weights are configured. Otherwise, this parameter indicates the total number of ECS instances in the scaling group.
	//
	// example:
	//
	// 1
	TotalCapacity *int32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The total number of ECS instances in the scaling group.
	//
	// example:
	//
	// 1
	TotalInstanceCount *int32 `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
	// The list of backend server groups.
	VServerGroups []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	// The ID of the vSwitch associated with the scaling group.
	//
	// example:
	//
	// vsw-bp1whw2u46cn8zubm****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The collection of IDs of the vSwitches associated with the scaling group. If you use the VSwitchIds parameter, the VSwitchId parameter is ignored.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC to which the scaling group belongs.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of instances in the warm pool.
	//
	// example:
	//
	// 5
	WarmPoolCapacity *int32 `json:"WarmPoolCapacity,omitempty" xml:"WarmPoolCapacity,omitempty"`
	// The current warm pool configuration of the scaling group.
	WarmPoolConfiguration *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration `json:"WarmPoolConfiguration,omitempty" xml:"WarmPoolConfiguration,omitempty" type:"Struct"`
	// The zone shift options.
	ZoneShiftOptions *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions `json:"ZoneShiftOptions,omitempty" xml:"ZoneShiftOptions,omitempty" type:"Struct"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetActiveCapacity() *int32 {
	return s.ActiveCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetActiveScalingConfigurationId() *string {
	return s.ActiveScalingConfigurationId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetAlbServerGroups() []*DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	return s.AlbServerGroups
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetAutoRebalance() *bool {
	return s.AutoRebalance
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetAzBalance() *bool {
	return s.AzBalance
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetBalanceMode() *string {
	return s.BalanceMode
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetCapacityOptions() *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions {
	return s.CapacityOptions
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetCurrentHostName() *string {
	return s.CurrentHostName
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetCustomPolicyARN() *string {
	return s.CustomPolicyARN
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetDBInstanceIds() []*string {
	return s.DBInstanceIds
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetDBInstances() []*DescribeScalingGroupsResponseBodyScalingGroupsDBInstances {
	return s.DBInstances
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetDefaultCooldown() *int32 {
	return s.DefaultCooldown
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetEnableDesiredCapacity() *bool {
	return s.EnableDesiredCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetGroupDeletionProtection() *bool {
	return s.GroupDeletionProtection
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetHealthCheckTypes() []*string {
	return s.HealthCheckTypes
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetInitCapacity() *int32 {
	return s.InitCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetIsElasticStrengthInAlarm() *bool {
	return s.IsElasticStrengthInAlarm
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetLaunchTemplateOverrides() []*DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides {
	return s.LaunchTemplateOverrides
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetLoadBalancerConfigs() []*DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs {
	return s.LoadBalancerConfigs
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetLoadBalancerIds() []*string {
	return s.LoadBalancerIds
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetMaxInstanceLifetime() *int32 {
	return s.MaxInstanceLifetime
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetMinSize() *int32 {
	return s.MinSize
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetMonitorGroupId() *string {
	return s.MonitorGroupId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetMultiAZPolicy() *string {
	return s.MultiAZPolicy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetPendingCapacity() *int32 {
	return s.PendingCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetPendingWaitCapacity() *int32 {
	return s.PendingWaitCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetProtectedCapacity() *int32 {
	return s.ProtectedCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetRemovalPolicies() []*string {
	return s.RemovalPolicies
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetRemovingCapacity() *int32 {
	return s.RemovingCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetRemovingWaitCapacity() *int32 {
	return s.RemovingWaitCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetScalingGroupName() *string {
	return s.ScalingGroupName
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetServerGroups() []*DescribeScalingGroupsResponseBodyScalingGroupsServerGroups {
	return s.ServerGroups
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetSpotAllocationStrategy() *string {
	return s.SpotAllocationStrategy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetSpotCapacity() *int32 {
	return s.SpotCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetSpotInstancePools() *int32 {
	return s.SpotInstancePools
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetStandbyCapacity() *int32 {
	return s.StandbyCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetStopInstanceTimeout() *int32 {
	return s.StopInstanceTimeout
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetStoppedCapacity() *int32 {
	return s.StoppedCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetSuspendedProcesses() []*string {
	return s.SuspendedProcesses
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetSystemSuspended() *bool {
	return s.SystemSuspended
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetTags() []*DescribeScalingGroupsResponseBodyScalingGroupsTags {
	return s.Tags
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetTotalCapacity() *int32 {
	return s.TotalCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetTotalInstanceCount() *int32 {
	return s.TotalInstanceCount
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetVServerGroups() []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups {
	return s.VServerGroups
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetWarmPoolCapacity() *int32 {
	return s.WarmPoolCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetWarmPoolConfiguration() *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration {
	return s.WarmPoolConfiguration
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) GetZoneShiftOptions() *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions {
	return s.ZoneShiftOptions
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetActiveCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ActiveCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetActiveScalingConfigurationId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAlbServerGroups(v []*DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AlbServerGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAllocationStrategy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AllocationStrategy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAutoRebalance(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AutoRebalance = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetAzBalance(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.AzBalance = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetBalanceMode(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.BalanceMode = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCapacityOptions(v *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CapacityOptions = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCompensateWithOnDemand(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCreationTime(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCurrentHostName(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CurrentHostName = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetCustomPolicyARN(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.CustomPolicyARN = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDBInstanceIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDBInstances(v []*DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DBInstances = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDefaultCooldown(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DefaultCooldown = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetDesiredCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.DesiredCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetEnableDesiredCapacity(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.EnableDesiredCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetGroupDeletionProtection(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.GroupDeletionProtection = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetGroupType(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.GroupType = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetHealthCheckType(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetHealthCheckTypes(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.HealthCheckTypes = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetInitCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.InitCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetIsElasticStrengthInAlarm(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.IsElasticStrengthInAlarm = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLaunchTemplateId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLaunchTemplateOverrides(v []*DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLaunchTemplateVersion(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLifecycleState(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLoadBalancerConfigs(v []*DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LoadBalancerConfigs = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetLoadBalancerIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.LoadBalancerIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMaxInstanceLifetime(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMaxSize(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMinSize(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetModificationTime(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ModificationTime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMonitorGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MonitorGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetMultiAZPolicy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.MultiAZPolicy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetOnDemandBaseCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetOnDemandPercentageAboveBaseCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetPendingCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.PendingCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetPendingWaitCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.PendingWaitCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetProtectedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ProtectedCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRegionId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRemovalPolicies(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RemovalPolicies = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRemovingCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RemovingCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetRemovingWaitCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.RemovingWaitCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetResourceGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingGroupName(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingGroupName = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingPolicy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingPolicy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetServerGroups(v []*DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ServerGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotAllocationStrategy(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotInstancePools(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotInstancePools = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSpotInstanceRemedy(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetStandbyCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.StandbyCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetStopInstanceTimeout(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.StopInstanceTimeout = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetStoppedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.StoppedCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSuspendedProcesses(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SuspendedProcesses = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetSystemSuspended(v bool) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.SystemSuspended = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetTags(v []*DescribeScalingGroupsResponseBodyScalingGroupsTags) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.Tags = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetTotalCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetTotalInstanceCount(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.TotalInstanceCount = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVServerGroups(v []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VServerGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVSwitchId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VSwitchId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVSwitchIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VSwitchIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetVpcId(v string) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.VpcId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetWarmPoolCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.WarmPoolCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetWarmPoolConfiguration(v *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.WarmPoolConfiguration = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetZoneShiftOptions(v *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ZoneShiftOptions = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) Validate() error {
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
	if s.WarmPoolConfiguration != nil {
		if err := s.WarmPoolConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.ZoneShiftOptions != nil {
		if err := s.ZoneShiftOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups struct {
	// The ID of the ALB server group.
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port number used by the ECS instance after Auto Scaling adds the ECS instance to the ALB server group.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to the ALB server group.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) GetAlbServerGroupId() *string {
	return s.AlbServerGroupId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) SetAlbServerGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) SetPort(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) SetWeight(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions struct {
	// Specifies whether to automatically create pay-as-you-go instances to meet the required number of ECS instances when the MultiAZPolicy parameter is set to COST_OPTIMIZED and sufficient spot instances cannot be created due to price or inventory reasons. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. If the number of pay-as-you-go instances in the scaling group is less than this value, the system preferentially creates pay-as-you-go instances. Valid values: 0 to 1000.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances among the excess instances after the scaling group meets the minimum number of pay-as-you-go instances specified by OnDemandBaseCapacity. Valid values: 0 to 100.
	//
	// example:
	//
	// 0
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The price comparison mode of the cost optimization strategy for the scaling group. Valid values:
	//
	// - PricePerUnit: Compares prices based on the price per unit of capacity.
	//
	//   The capacity of an instance in the scaling group is equal to the weight set for the instance type. The default weight is 1, which means one ECS instance occupies one unit of capacity.
	//
	// - PricePerVCpu: Compares prices based on the price per vCPU.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// Specifies whether to use spot instances to replace pay-as-you-go capacity when CompensateWithOnDemand is enabled and the percentage of pay-as-you-go instances exceeds the value of OnDemandPercentageAboveBaseCapacity. In common scenarios where CompensateWithOnDemand downgrades spot instances to pay-as-you-go instances due to spot inventory or price reasons, this feature attempts to use spot instances to replace the excess pay-as-you-go capacity to avoid long-term retention of pay-as-you-go instances. Valid values:
	//
	// - true: Allowed.
	//
	// - false: Not allowed.
	//
	// example:
	//
	// false
	SpotAutoReplaceOnDemand *bool `json:"SpotAutoReplaceOnDemand,omitempty" xml:"SpotAutoReplaceOnDemand,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) GetPriceComparisonMode() *string {
	return s.PriceComparisonMode
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) GetSpotAutoReplaceOnDemand() *bool {
	return s.SpotAutoReplaceOnDemand
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) SetCompensateWithOnDemand(v bool) *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) SetOnDemandBaseCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) SetOnDemandPercentageAboveBaseCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) SetPriceComparisonMode(v string) *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions {
	s.PriceComparisonMode = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) SetSpotAutoReplaceOnDemand(v bool) *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions {
	s.SpotAutoReplaceOnDemand = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsDBInstances struct {
	// The database instance ID.
	//
	// example:
	//
	// rm-m5eqju85s45mu0***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of security group whitelists added to the database by the scaling group.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The type of the database. Valid values:
	//
	// - RDS
	//
	// - Redis
	//
	// - MongoDB
	//
	// example:
	//
	// RDS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) GetType() *string {
	return s.Type
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) SetDBInstanceId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) SetSecurityGroupIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) SetType(v string) *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances {
	s.Type = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsDBInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides struct {
	// The specified instance type, which overrides the instance type in the launch template.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum hourly price for the instance type specified in the launch template override (namely LaunchTemplateOverride.InstanceType).
	//
	// > This parameter takes effect only when the LaunchTemplateId parameter is specified.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The weight of the specified instance type, which indicates the capacity of a single instance of this type in the scaling group. A higher weight indicates that fewer instances of this type are required to meet the expected capacity.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) SetInstanceType(v string) *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) SetSpotPriceLimit(v float32) *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) SetWeightedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs struct {
	// The ID of the Classic Load Balancer (CLB) instance for load balancing.
	//
	// example:
	//
	// lb-2zep8alpq5zq1a2xwyxxx
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to the SLB server group. A higher weight indicates that more access requests are distributed to the ECS instance. If the weight is 0, the ECS instance does not receive access requests. Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) SetLoadBalancerId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) SetWeight(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsServerGroups struct {
	// The port number used by the ECS instance after Auto Scaling adds the ECS instance to the server group.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the server group.
	//
	// example:
	//
	// sgp-i9ouakeaerr*****
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
	// The weight of the ECS instance as a backend server after Auto Scaling adds the ECS instance to the server group.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) GetType() *string {
	return s.Type
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) SetPort(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) SetServerGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) SetType(v string) *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups {
	s.Type = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) SetWeight(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsServerGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsTags struct {
	// Specifies whether the tag is propagated. Valid values:
	//
	// - true: Tags on the scaling group are propagated only to newly created instances, but not to existing instances in the scaling group.
	//
	// - false: Tags on the scaling group are not propagated to instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Propagate *bool `json:"Propagate,omitempty" xml:"Propagate,omitempty"`
	// The tag key of the scaling group.
	//
	// example:
	//
	// Department
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the scaling group.
	//
	// example:
	//
	// Finance
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) GetPropagate() *bool {
	return s.Propagate
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) SetPropagate(v bool) *DescribeScalingGroupsResponseBodyScalingGroupsTags {
	s.Propagate = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) SetTagKey(v string) *DescribeScalingGroupsResponseBodyScalingGroupsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) SetTagValue(v string) *DescribeScalingGroupsResponseBodyScalingGroupsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsTags) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups struct {
	// The ID of the load balancing instance to which the backend server group belongs.
	//
	// example:
	//
	// 147b46d767c-cn-qingdao-cm5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the backend server group.
	VServerGroupAttributes []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) GetVServerGroupAttributes() []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	return s.VServerGroupAttributes
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) SetLoadBalancerId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) SetVServerGroupAttributes(v []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups) Validate() error {
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

type DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes struct {
	// The port number used by the load balancing instance to provide services.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The attributes of the backend server group.
	//
	// example:
	//
	// rsp-bp12bjrny****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The weight of the backend server group.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) GetPort() *int32 {
	return s.Port
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) SetPort(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) SetWeight(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsVServerGroupsVServerGroupAttributes) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration struct {
	// The capacity configuration of the warm pool.
	CapacityOptions *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// The instance reuse policy.
	InstanceReusePolicy *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy `json:"InstanceReusePolicy,omitempty" xml:"InstanceReusePolicy,omitempty" type:"Struct"`
	// The state of instances in the warm pool. Valid values:
	//
	// - Running: The instances are in the running state. Scale-out is faster, but you are charged for the full cost of the instances.
	//
	// - Stopped: The instances are in the economical mode. You are charged only for resources such as cloud disks, elastic IP addresses (EIPs), and snapshots of the instances. However, instance startup may fail during scale-out due to insufficient inventory.
	//
	// example:
	//
	// Running
	PoolState *string `json:"PoolState,omitempty" xml:"PoolState,omitempty"`
	// The status of the warm pool. Valid values:
	//
	// - Active: Active.
	//
	// - Deleting: Deleting.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) GetCapacityOptions() *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions {
	return s.CapacityOptions
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) GetInstanceReusePolicy() *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy {
	return s.InstanceReusePolicy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) GetPoolState() *string {
	return s.PoolState
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) GetStatus() *string {
	return s.Status
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) SetCapacityOptions(v *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration {
	s.CapacityOptions = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) SetInstanceReusePolicy(v *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration {
	s.InstanceReusePolicy = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) SetPoolState(v string) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration {
	s.PoolState = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) SetStatus(v string) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration {
	s.Status = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfiguration) Validate() error {
	if s.CapacityOptions != nil {
		if err := s.CapacityOptions.Validate(); err != nil {
			return err
		}
	}
	if s.InstanceReusePolicy != nil {
		if err := s.InstanceReusePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions struct {
	// The expected number of pre-warmed instances to cover. The system maintains the sum of the current number of instances in the scaling group and the number of instances in the warm pool at this value.
	//
	// example:
	//
	// 5
	DesiredGroupPreparedCapacity *int32 `json:"DesiredGroupPreparedCapacity,omitempty" xml:"DesiredGroupPreparedCapacity,omitempty"`
	// The minimum number of instances in the warm pool. The system always retains no less than this number of pre-warmed instances in the warm pool to ensure rapid scale-out from the warm pool at any time.
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) GetDesiredGroupPreparedCapacity() *int32 {
	return s.DesiredGroupPreparedCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) SetDesiredGroupPreparedCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions {
	s.DesiredGroupPreparedCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) SetMinCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions {
	s.MinCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationCapacityOptions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy struct {
	// Specifies whether instances are returned to the warm pool when the scaling group scales in.
	//
	// example:
	//
	// true
	ReuseOnScaleIn *bool `json:"ReuseOnScaleIn,omitempty" xml:"ReuseOnScaleIn,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy) GetReuseOnScaleIn() *bool {
	return s.ReuseOnScaleIn
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy) SetReuseOnScaleIn(v bool) *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy {
	s.ReuseOnScaleIn = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsWarmPoolConfigurationInstanceReusePolicy) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions struct {
	// The expiration time of the zone shift. After the expiration time, the system automatically ends the shift and restores the related zones. The time is displayed in UTC and follows the ISO 8601 standard in the YYYY-MM-DDThh:mmZ format.
	//
	// example:
	//
	// 2026-01-01T12:05Z
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// The list of impaired zone IDs.
	//
	// example:
	//
	// cn-hangzhou-g
	ImpairedZoneIds []*string `json:"ImpairedZoneIds,omitempty" xml:"ImpairedZoneIds,omitempty" type:"Repeated"`
	// The behavior policy of the scaling group during the zone shift.
	ZoneShiftPolicy *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy `json:"ZoneShiftPolicy,omitempty" xml:"ZoneShiftPolicy,omitempty" type:"Struct"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) GetExpiryTime() *string {
	return s.ExpiryTime
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) GetImpairedZoneIds() []*string {
	return s.ImpairedZoneIds
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) GetZoneShiftPolicy() *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy {
	return s.ZoneShiftPolicy
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) SetExpiryTime(v string) *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions {
	s.ExpiryTime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) SetImpairedZoneIds(v []*string) *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions {
	s.ImpairedZoneIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) SetZoneShiftPolicy(v *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions {
	s.ZoneShiftPolicy = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptions) Validate() error {
	if s.ZoneShiftPolicy != nil {
		if err := s.ZoneShiftPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy struct {
	// The health check behavior in impaired zones. Valid values:
	//
	// - ReplaceUnhealthy: Unhealthy instances in all zones are replaced normally.
	//
	// - IgnoreUnhealthy: Unhealthy instances in impaired zones are not replaced. Healthy zones are not affected.
	//
	// example:
	//
	// ReplaceUnhealthy
	ImpairedZoneHealthCheckBehavior *string `json:"ImpairedZoneHealthCheckBehavior,omitempty" xml:"ImpairedZoneHealthCheckBehavior,omitempty"`
	// Blocks scale-in triggered by specified types of scaling rules. The blocking scope covers all zones under the scaling group, not just the impaired zones. This prevents false scale-in caused by seemingly improved monitoring metrics due to unreachable instances in impaired zones. Valid values:
	//
	// - StepScalingRule: Blocks scale-in triggered by step scaling rules.
	//
	// - TargetTrackingScalingRule: Blocks scale-in triggered by target tracking scaling rules.
	//
	// - None: No blocking.
	//
	// example:
	//
	// StepScalingRule
	ScaleInBlockedRuleTypes []*string `json:"ScaleInBlockedRuleTypes,omitempty" xml:"ScaleInBlockedRuleTypes,omitempty" type:"Repeated"`
	// Specifies whether scale-out creates instances only in healthy zones during the shift. Valid values:
	//
	// - true: Scale-out creates instances only in healthy zones and skips impaired zones.
	//
	// - false: Scale-out is not restricted by zone and can create instances in any zone under the scaling group.
	//
	// example:
	//
	// true
	ScaleOutInHealthyZonesOnly *bool `json:"ScaleOutInHealthyZonesOnly,omitempty" xml:"ScaleOutInHealthyZonesOnly,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) GetImpairedZoneHealthCheckBehavior() *string {
	return s.ImpairedZoneHealthCheckBehavior
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) GetScaleInBlockedRuleTypes() []*string {
	return s.ScaleInBlockedRuleTypes
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) GetScaleOutInHealthyZonesOnly() *bool {
	return s.ScaleOutInHealthyZonesOnly
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) SetImpairedZoneHealthCheckBehavior(v string) *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy {
	s.ImpairedZoneHealthCheckBehavior = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) SetScaleInBlockedRuleTypes(v []*string) *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy {
	s.ScaleInBlockedRuleTypes = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) SetScaleOutInHealthyZonesOnly(v bool) *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy {
	s.ScaleOutInHealthyZonesOnly = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsZoneShiftOptionsZoneShiftPolicy) Validate() error {
	return dara.Validate(s)
}
