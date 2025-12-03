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
	// The page number.
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
	// The total number of entries returned.
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
	// The number of ECS instances that are in the In Service state in the scaling group.
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
	// The Application Load Balancer (ALB) server groups.
	AlbServerGroups []*DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The instance allocation policy. Auto Scaling selects instance types based on the allocation policy to create the required number of preemptible instances. The policy is suitable for pay-as-you-go instances and preemptible instances. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling adopts the predefined instance type sequence to create the required number of preemptible instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the most economical vCPU pricing to create the required number of instances.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Whether to enable automatic rebalancing for the scaling group. This takes effect only when BalancedOnly is enabled for the scaling group. Valid value:
	//
	// 	- false: Auto rebalancing is disabled for the scaling group.
	//
	// 	- true: If Auto rebalancing is enabled, the scaling group automatically detects the capacity of the zone. If the capacity of the zone is unbalanced, the scaling group actively scales out the zone and re-balances the capacity of the zone.
	//
	// example:
	//
	// false
	AutoRebalance *bool `json:"AutoRebalance,omitempty" xml:"AutoRebalance,omitempty"`
	// Indicates whether instances in the scaling group are evenly distributed across the specified zones. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AzBalance *bool `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
	// The zone balancing mode. This mode takes effect only when the zone balancing mode is enabled. Valid value:
	//
	// 	- Default value: BalancedBestEffort. If a resource fails to be created in a zone, the resource is downgraded to another zone. This ensures best-effort delivery of the resource.
	//
	// 	- BalancedOnly: If a resource fails to be created in a zone, the resource is not downgraded to another zone. The scale-out activity is partially successful to avoid excessive imbalance of resources in different zones.
	//
	// example:
	//
	// BalancedBestEffort
	BalanceMode *string `json:"BalanceMode,omitempty" xml:"BalanceMode,omitempty"`
	// The capacity options.
	CapacityOptions *DescribeScalingGroupsResponseBodyScalingGroupsCapacityOptions `json:"CapacityOptions,omitempty" xml:"CapacityOptions,omitempty" type:"Struct"`
	// Indicates whether Auto Scaling can create pay-as-you-go instances to supplement preemptible instances if preemptible instances cannot be created due to price-related factors or insufficient inventory when MultiAZPolicy is set to COST_OPTIMIZED. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The time when the scaling group was created.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// >  This parameter is unavailable.
	//
	// example:
	//
	// hostname
	CurrentHostName *string `json:"CurrentHostName,omitempty" xml:"CurrentHostName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the function that is specified in the custom scale-in policy. This parameter takes effect only if you set the first value of RemovalPolicies to CustomPolicy.
	//
	// example:
	//
	// acs:fc:cn-zhangjiakou:16145688****:services/ess_custom_terminate_policy.LATEST/functions/ess_custom_terminate_policy_name
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The IDs of the ApsaraDB RDS instances that are attached to the scaling group.
	DBInstanceIds []*string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
	// The databases that are attached to the scaling group.
	DBInstances []*DescribeScalingGroupsResponseBodyScalingGroupsDBInstances `json:"DBInstances,omitempty" xml:"DBInstances,omitempty" type:"Repeated"`
	// The cooldown period of the scaling group. During the cooldown period, Auto Scaling does not execute the scaling activities that are triggered by [CloudMonitor](https://help.aliyun.com/document_detail/35170.html) event-triggered tasks.
	//
	// example:
	//
	// 60
	DefaultCooldown *int32 `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	// The expected number of ECS instances in the scaling group. Auto Scaling automatically maintains the expected number of ECS instances.
	//
	// example:
	//
	// 5
	DesiredCapacity *int32 `json:"DesiredCapacity,omitempty" xml:"DesiredCapacity,omitempty"`
	// Indicates whether the Expected Number of Instances feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableDesiredCapacity *bool `json:"EnableDesiredCapacity,omitempty" xml:"EnableDesiredCapacity,omitempty"`
	// Indicates whether the Deletion Protection feature is enabled for the scaling group. Valid values:
	//
	// 	- true: The Deletion Protection feature is enabled for the scaling group. The scaling group cannot be deleted.
	//
	// 	- false: The Deletion Protection feature is disabled for the scaling group.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The type of instances that are managed by the scaling group.
	//
	// example:
	//
	// ECS
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// 	- NONE: Auto Scaling does not perform health checks.
	//
	// 	- ECS: Auto Scaling checks the health status of instances in the scaling group. If you want to enable instance health check, you can set the value to ECS, regardless of whether the scaling group is of ECS type or Elastic Container Instance type.
	//
	// 	- LOAD_BALANCER: Auto Scaling checks the health status of instances in the scaling group based on the health check results of load balancers. The health check results of Classic Load Balancer (CLB) instances are not supported as the health check basis for instances in the scaling group.
	//
	// example:
	//
	// ECS
	HealthCheckType *string `json:"HealthCheckType,omitempty" xml:"HealthCheckType,omitempty"`
	// The health check mode of the scaling group. Valid values:
	//
	// 	- NONE: Auto Scaling does not perform health checks.
	//
	// 	- ECS: Auto Scaling checks the health status of instances in the scaling group. If you want to enable instance health check, you can set the value to ECS, regardless of whether the scaling group is of ECS type or Elastic Container Instance type.
	//
	// 	- LOAD_BALANCER: Auto Scaling checks the health status of instances in the scaling group based on the health check results of load balancers. The health check results of CLB instances are not supported as the health check basis for instances in the scaling group.
	HealthCheckTypes []*string `json:"HealthCheckTypes,omitempty" xml:"HealthCheckTypes,omitempty" type:"Repeated"`
	// The number of instances that are initialized before they are added into the scaling group.
	//
	// example:
	//
	// 0
	InitCapacity *int32 `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	// >  This parameter is not available for use.
	//
	// example:
	//
	// false
	IsElasticStrengthInAlarm *bool `json:"IsElasticStrengthInAlarm,omitempty" xml:"IsElasticStrengthInAlarm,omitempty"`
	// The ID of the launch template that is used by the scaling group.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The instance types that are extended in the launch template.
	LaunchTemplateOverrides []*DescribeScalingGroupsResponseBodyScalingGroupsLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version of the launch template that is used by the scaling group.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The status of the scaling group. Valid values:
	//
	// 	- Active: The scaling group is active. Active scaling groups can receive requests to execute scaling rules and trigger scaling activities.
	//
	// 	- Inactive: The scaling group is in the Disabled state. Disabled scaling groups cannot receive requests to execute scaling rules.
	//
	// 	- Deleting: The scaling group is being deleted. Scaling groups that are being deleted cannot receive requests to execute scaling rules, and the parameter settings of the scaling groups cannot be modified.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The load balancer configurations.
	LoadBalancerConfigs []*DescribeScalingGroupsResponseBodyScalingGroupsLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The IDs of the load balancers that are attached to the scaling group.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum life span of each ECS instance in the scaling group. Unit: seconds.
	//
	// Valid values: 0 or `[86400, Integer.maxValue]`. A value of 0 for MaxInstanceLifetime indicates that a previously set limit has been removed. This effectively disables the maximum instance lifetime constraint.
	//
	// Default value: null.
	//
	// >  This parameter is not supported by scaling groups of the Elastic Container Instance type and scaling groups whose ScalingPolicy is set to Recycle.
	//
	// example:
	//
	// null
	MaxInstanceLifetime *int32 `json:"MaxInstanceLifetime,omitempty" xml:"MaxInstanceLifetime,omitempty"`
	// The maximum number of ECS instances that can be contained in the scaling group.
	//
	// example:
	//
	// 2
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The minimum number of ECS instances that must be contained in the scaling group.
	//
	// example:
	//
	// 1
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The time when the scaling group was last modified.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The ID of the CloudMonitor application group that is associated with the scaling group.
	//
	// example:
	//
	// 1497****
	MonitorGroupId *string `json:"MonitorGroupId,omitempty" xml:"MonitorGroupId,omitempty"`
	// The scaling policy of the ECS instances in the multi-zone scaling group. Valid values:
	//
	// 	- PRIORITY: ECS instances are created based on the value of VSwitchIds. If Auto Scaling cannot create ECS instances in the zone where the vSwitch of the highest priority resides, Auto Scaling creates ECS instances in the zone where the vSwitch of the next highest priority resides.
	//
	// 	- COST_OPTIMIZED: ECS instances are created based on the unit prices of their vCPUs. Auto Scaling preferentially creates ECS instances whose vCPUs are provided at the lowest price. If preemptible instance types are specified in the scaling configuration, Auto Scaling preferentially creates preemptible instances. You can also specify CompensateWithOnDemand to allow Auto Scaling to create pay-as-you-go instances if preemptible instances cannot be created due to limited stock.
	//
	//     **
	//
	//     **Note*	- The COST_OPTIMIZED setting takes effect only if your scaling configuration contains multiple instance types or contains preemptible instance types.
	//
	// 	- BALANCE: ECS instances are evenly distributed across the zones that are specified for the scaling group. If ECS instances become unevenly distributed across the specified zones due to limited instance type availability, you can call the RebalanceInstance operation to balance the distribution of the ECS instances.
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
	// The percentage of pay-as-you-go instances in excess when the minimum number of pay-as-you-go instances reaches the threshold. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// The number of ECS instances that are being added to the scaling group and configured.
	//
	// example:
	//
	// 0
	PendingCapacity *int32 `json:"PendingCapacity,omitempty" xml:"PendingCapacity,omitempty"`
	// The number of ECS instances that are in the Pending Add state in the scaling group.
	//
	// example:
	//
	// 1
	PendingWaitCapacity *int32 `json:"PendingWaitCapacity,omitempty" xml:"PendingWaitCapacity,omitempty"`
	// The number of ECS instances that are in the Protected state in the scaling group.
	//
	// example:
	//
	// 1
	ProtectedCapacity *int32 `json:"ProtectedCapacity,omitempty" xml:"ProtectedCapacity,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance removal policies. Valid values:
	//
	// 	- OldestInstance: Auto Scaling removes ECS instances that are added at the earliest point in time to the scaling group.
	//
	// 	- NewestInstance: Auto Scaling removes ECS instances that are most recently added to the scaling group.
	//
	// 	- OldestScalingConfiguration: Auto Scaling removes ECS instances that are created from the earliest scaling configuration.
	RemovalPolicies []*string `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Repeated"`
	// The number of ECS instances that are being removed from the scaling group.
	//
	// example:
	//
	// 0
	RemovingCapacity *int32 `json:"RemovingCapacity,omitempty" xml:"RemovingCapacity,omitempty"`
	// The number of ECS instances that are in the Pending Remove state in the scaling group.
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
	// The instance reclaim mode of the scaling group. Valid values:
	//
	// 	- recycle: economical mode.
	//
	// 	- release: release mode.
	//
	// 	- forcerelease: forced release mode.
	//
	// For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The server groups.
	//
	// >  You can use this parameter to obtain information about ALB server groups and Network Load Balancer (NLB) server groups attached to the scaling group.
	ServerGroups []*DescribeScalingGroupsResponseBodyScalingGroupsServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The allocation policy of preemptible instances. This parameter indicates the manner in which Auto Scaling selects instance types to create the required number of preemptible instances. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling adopts the predefined instance type sequence to create the required number of preemptible instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the most economical vCPU pricing to create the required number of preemptible instances.
	//
	// Default value: priority.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of preemptible instances in the scaling group.
	//
	// example:
	//
	// 0
	SpotCapacity *int32 `json:"SpotCapacity,omitempty" xml:"SpotCapacity,omitempty"`
	// The number of instance types in the scaling group. Auto Scaling evenly creates preemptible instances of multiple instance types that are provided at the lowest price across the zones of the scaling group. Valid values: 0 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Indicates whether preemptible instances can be supplemented. If this parameter is set to true, Auto Scaling proactively creates instances to replace the preemptible instances for reclamation when Auto Scaling receives a system notification.
	//
	// example:
	//
	// true
	SpotInstanceRemedy *bool `json:"SpotInstanceRemedy,omitempty" xml:"SpotInstanceRemedy,omitempty"`
	// The number of ECS instances that are in the Standby state in the scaling group.
	//
	// example:
	//
	// 1
	StandbyCapacity *int32 `json:"StandbyCapacity,omitempty" xml:"StandbyCapacity,omitempty"`
	// The period of time that is required by the Elastic Compute Service (ECS) instance to enter the Stopped state during the scale-in process. Unit: seconds.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
	// The number of instances that was stopped in Economical Mode in the scaling group.
	//
	// example:
	//
	// 1
	StoppedCapacity *int32 `json:"StoppedCapacity,omitempty" xml:"StoppedCapacity,omitempty"`
	// The processes that are suspended. If no process is suspended, null is returned. Valid values:
	//
	// 	- ScaleIn: scale-in processes.
	//
	// 	- ScaleOut: scale-out processes.
	//
	// 	- HealthCheck: health check processes.
	//
	// 	- AlarmNotification: event-triggered task processes.
	//
	// 	- ScheduledAction: scheduled task processes.
	SuspendedProcesses []*string `json:"SuspendedProcesses,omitempty" xml:"SuspendedProcesses,omitempty" type:"Repeated"`
	// Indicates whether Auto Scaling stops executing scaling activities in the scaling group. Valid values:
	//
	// 	- true: Auto Scaling stops executing scaling activities in the scaling group if the scaling activities failed for more than seven consecutive days in the scaling group. In this case, you must modify the scaling group or scaling configuration to resume the scaling activities.
	//
	// 	- false: Auto Scaling does not stop executing scaling activities in the scaling group.
	//
	// example:
	//
	// true
	SystemSuspended *bool `json:"SystemSuspended,omitempty" xml:"SystemSuspended,omitempty"`
	// The tags of the scaling group.
	Tags []*DescribeScalingGroupsResponseBodyScalingGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total weighted capacity of all ECS instances in the scaling group if Weighted is specified. In other cases, this parameter specifies the total number of ECS instances in the scaling group.
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
	// The backend vServer groups.
	VServerGroups []*DescribeScalingGroupsResponseBodyScalingGroupsVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	// The vSwitch ID of the scaling group.
	//
	// example:
	//
	// vsw-bp1whw2u46cn8zubm****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of the vSwitches that are associated with the scaling group. If you specify VSwitchIds, VSwitchId is ignored.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which the scaling group resides.
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	return nil
}

type DescribeScalingGroupsResponseBodyScalingGroupsAlbServerGroups struct {
	// The ID of the ALB server group.
	//
	// example:
	//
	// sgp-ddwb0y0g6y9bjm****
	AlbServerGroupId *string `json:"AlbServerGroupId,omitempty" xml:"AlbServerGroupId,omitempty"`
	// The port number used by an ECS instance as a backend server in the ALB server group.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The weight of an ECS instance as a backend server in the ALB server group.
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
	// Indicates whether pay-as-you-go ECS instances can be automatically created to reach the required number of ECS instances when preemptible ECS instances cannot be created due to high prices or insufficient inventory of resources. This parameter takes effect when you set `MultiAZPolicy` to `COST_OPTIMIZED`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CompensateWithOnDemand *bool `json:"CompensateWithOnDemand,omitempty" xml:"CompensateWithOnDemand,omitempty"`
	// The minimum number of pay-as-you-go instances required in the scaling group. When the actual number of pay-as-you-go instances drops below the minimum threshold, Auto Scaling preferentially creates pay-as-you-go instances. Valid values: 0 to 1000.
	//
	// example:
	//
	// 0
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances in the excess instances when the minimum number of pay-as-you-go instances is reached. `OnDemandBaseCapacity` specifies the minimum number of pay-as-you-go instances required in the scaling group. Valid values: 0 to 100.
	//
	// example:
	//
	// 0
	OnDemandPercentageAboveBaseCapacity *int32 `json:"OnDemandPercentageAboveBaseCapacity,omitempty" xml:"OnDemandPercentageAboveBaseCapacity,omitempty"`
	// Indicates how prices are compared. Valid values:
	//
	// 	- PricePerUnit: Prices are compared based on the price per instance capacity.
	//
	//     Capacity is determined by the weights assigned to instance types in the scaling group. If no weight is specified, a default weight of 1 is used, meaning each instance is assigned a capacity of 1.
	//
	// 	- PricePerVCpu: Prices are compared based on the price per vCPU.
	//
	// example:
	//
	// PricePerUnit
	PriceComparisonMode *string `json:"PriceComparisonMode,omitempty" xml:"PriceComparisonMode,omitempty"`
	// Specifies whether to replace pay-as-you-go ECS instances with preemptible ECS instances. If you specify `CompensateWithOnDemand`, it may result in a higher percentage of pay-as-you-go instances compared to the value of `OnDemandPercentageAboveBaseCapacity`. In this scenario, Auto Scaling will try to deploy preemptible ECS instances to replace the surplus pay-as-you-go ECS instances. When `CompensateWithOnDemand` is specified, Auto Scaling creates pay-as-you-go ECS instances if there are not enough preemptible instance types available. To avoid keeping these pay-as-you-go ECS instances for long periods, Auto Scaling tries to replace them with preemptible instances as soon as enough of preemptible instance types become available. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// The ID of the database.
	//
	// example:
	//
	// rm-m5eqju85s45mu0***
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The IDs of the security groups that are added to the security group whitelist of the attached database.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The type of the database. Valid values:
	//
	// 	- RDS.
	//
	// 	- Redis.
	//
	// 	- MongoDB.
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
	// The instance type. The instance type that is specified by this parameter overrides the instance type that is specified in the launch template.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum bid price of the instance type that is specified by `LaunchTemplateOverride.InstanceType`.
	//
	// >  This parameter takes effect only if you use `LaunchTemplateId` to specify a launch template.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The weight of the instance type. The value of this parameter indicates the capacity of a single instance of the specified instance type in the scaling group. A higher weight indicates that a smaller number of instances of the specified instance type are required to meet the expected capacity.
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
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-2zep8alpq5zq1a2xwyxxx
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of an ECS instance as a backend server in the CLB server group. An increase in the weight of an ECS instance indicates an increase in the number of access requests that are forwarded to the ECS instance. If you set the weight of an ECS instance to 0, no access requests are forwarded to the ECS instance. Valid values: 0 to 100.
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
	// The port number used by an ECS instance as a backend server in the server group.
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
	// 	- ALB
	//
	// 	- NLB
	//
	// example:
	//
	// ALB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of an ECS instance as a backend server in the server group.
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
	// Indicates whether the tags of the scaling group can be propagated to instances. Valid values:
	//
	// 	- true: The tags of the scaling group can be propagated only to new instances.
	//
	// 	- false: The tags of the scaling group cannot be propagated to instances.
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
	// The ID of the load balancer to which the backend vServer group belongs.
	//
	// example:
	//
	// 147b46d767c-cn-qingdao-cm5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the backend vServer groups.
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
	// The port number that is used by the load balancer to provide external services.
	//
	// example:
	//
	// 22
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend vServer group.
	//
	// example:
	//
	// rsp-bp12bjrny****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The weight of the backend vServer group.
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
