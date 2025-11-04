// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingGroupDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutput(v string) *DescribeScalingGroupDetailResponseBody
	GetOutput() *string
	SetRequestId(v string) *DescribeScalingGroupDetailResponseBody
	GetRequestId() *string
	SetScalingGroup(v *DescribeScalingGroupDetailResponseBodyScalingGroup) *DescribeScalingGroupDetailResponseBody
	GetScalingGroup() *DescribeScalingGroupDetailResponseBodyScalingGroup
}

type DescribeScalingGroupDetailResponseBody struct {
	// The output details of the scaling group of the Elastic Container Instance type. Currently, the output is displayed in a Kubernetes Deployment YAML file.
	//
	// example:
	//
	// apiVersion: apps/v1
	//
	// kind: Deployment
	//
	// metadata:
	//
	//   name: nginx-deployment
	//
	//   labels:
	//
	//     app: nginx
	//
	//   spec:
	//
	//     replicas: 3
	//
	//     selector:
	//
	//        matchLabels:
	//
	//         app: nginx
	//
	//     template:
	//
	//       metadata:
	//
	//         labels:
	//
	//           app: nginx
	//
	//         annotations:
	//
	//           k8s.aliyun.com/eip-bandwidth: 10
	//
	//           k8s.aliyun.com/eci-with-eip: true
	//
	//         spec:
	//
	//           containers:
	//
	//           - name: nginx
	//
	//             image: nginx:1.14.2
	//
	//             ports:
	//
	//             - containerPort: 80
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling group.
	ScalingGroup *DescribeScalingGroupDetailResponseBodyScalingGroup `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Struct"`
}

func (s DescribeScalingGroupDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBody) GetOutput() *string {
	return s.Output
}

func (s *DescribeScalingGroupDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingGroupDetailResponseBody) GetScalingGroup() *DescribeScalingGroupDetailResponseBodyScalingGroup {
	return s.ScalingGroup
}

func (s *DescribeScalingGroupDetailResponseBody) SetOutput(v string) *DescribeScalingGroupDetailResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBody) SetRequestId(v string) *DescribeScalingGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBody) SetScalingGroup(v *DescribeScalingGroupDetailResponseBodyScalingGroup) *DescribeScalingGroupDetailResponseBody {
	s.ScalingGroup = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBody) Validate() error {
	if s.ScalingGroup != nil {
		if err := s.ScalingGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScalingGroupDetailResponseBodyScalingGroup struct {
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
	// The information about the Application Load Balancer (ALB) server groups.
	AlbServerGroups []*DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups `json:"AlbServerGroups,omitempty" xml:"AlbServerGroups,omitempty" type:"Repeated"`
	// The allocation policy of instances. Auto Scaling selects instance types based on the allocation policy to create the required number of instances. You can apply the policy to pay-as-you-go instances and preemptible instances. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling selects instance types based on the specified order to create the required number of instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the lowest unit price of vCPUs to create the required number of instances.
	//
	// example:
	//
	// priority
	AllocationStrategy *string `json:"AllocationStrategy,omitempty" xml:"AllocationStrategy,omitempty"`
	// Indicates whether instances in the scaling group are evenly distributed across zones. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AzBalance *bool `json:"AzBalance,omitempty" xml:"AzBalance,omitempty"`
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
	// The time when the scaling group was created.
	//
	// example:
	//
	// 2014-08-14T10:58Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// >  This parameter is not available for use.
	//
	// example:
	//
	// hostname
	CurrentHostName *string `json:"CurrentHostName,omitempty" xml:"CurrentHostName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the function that is specified in the custom scale-in policy. This parameter takes effect only if you set the first value of RemovalPolicies to CustomPolicy.
	//
	// example:
	//
	// null
	CustomPolicyARN *string `json:"CustomPolicyARN,omitempty" xml:"CustomPolicyARN,omitempty"`
	// The IDs of the ApsaraDB RDS instances that are associated with the scaling group.
	DBInstanceIds []*string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Repeated"`
	// The cooldown period of the scaling group. Unit: seconds.
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
	// true
	EnableDesiredCapacity *bool `json:"EnableDesiredCapacity,omitempty" xml:"EnableDesiredCapacity,omitempty"`
	// Indicates whether Deletion Protection is enabled for the scaling group. Valid values:
	//
	// 	- true: Deletion Protection is enabled for the scaling group. This way, the scaling group cannot be deleted.
	//
	// 	- false: Deletion Protection is disabled for the scaling group.
	//
	// example:
	//
	// true
	GroupDeletionProtection *bool `json:"GroupDeletionProtection,omitempty" xml:"GroupDeletionProtection,omitempty"`
	// The type of the instances that are managed by the scaling group. Valid values:
	//
	// 	- ECS: ECS instances
	//
	// 	- ECI: elastic container instances
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
	// The number of instances that are in the Initialized state and not added to the scaling group.
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
	// The information about the instance types that are extended in the launch template.
	LaunchTemplateOverrides []*DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides `json:"LaunchTemplateOverrides,omitempty" xml:"LaunchTemplateOverrides,omitempty" type:"Repeated"`
	// The version number of the launch template.
	//
	// example:
	//
	// Default
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The status of the scaling group. Valid values:
	//
	// 	- Active: The scaling group is in the Enabled state. Enabled scaling groups can receive requests to execute scaling rules and trigger scaling activities.
	//
	// 	- Inactive: The scaling group is in the Disabled state. Disabled scaling groups cannot receive requests to execute scaling rules.
	//
	// 	- Deleting: The scaling group is being deleted. Scaling groups that are being deleted cannot receive requests to execute scaling rules, and the parameter settings of the scaling groups cannot be modified.
	//
	// example:
	//
	// Active
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The CLB configurations.
	LoadBalancerConfigs []*DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs `json:"LoadBalancerConfigs,omitempty" xml:"LoadBalancerConfigs,omitempty" type:"Repeated"`
	// The IDs of the SLB instances that are associated with the scaling group.
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Repeated"`
	// The maximum life span of an instance in the scaling group. Unit: seconds.
	//
	// Valid values: 0 or from 86400 to `Integer.maxValue`. A value of 0 for MaxInstanceLifetime indicates that any previously set limit has been removed, which effectively disables the maximum instance lifetime constraint.
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
	// The time when the scaling group was modified.
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
	// 	- COST_OPTIMIZED: ECS instances are created based on the unit prices of their vCPUs. Auto Scaling preferentially creates ECS instances that use the lowest-priced vCPUs. If preemptible instance types are specified in the scaling configuration, Auto Scaling preferentially creates preemptible instances. You can also specify CompensateWithOnDemand to allow Auto Scaling to create pay-as-you-go instances in the case that preemptible instances cannot be created due to insufficient inventory of preemptible instance types.
	//
	//     **
	//
	//     **Note*	- The COST_OPTIMIZED setting takes effect only if you specified multiple instance types or preemptible instance types in your scaling configuration.
	//
	// 	- BALANCE: ECS instances are evenly distributed across the zones of the scaling group. If ECS instance are unevenly distributed across the specified zones due to insufficient inventory of instance types, you can call the RebalanceInstance operation to rebalance the distribution of the ECS instances.
	//
	// example:
	//
	// PRIORITY
	MultiAZPolicy *string `json:"MultiAZPolicy,omitempty" xml:"MultiAZPolicy,omitempty"`
	// The minimum number of pay-as-you-go instances that must be contained in the scaling group. Valid values: 0 to 1000. If the number of pay-as-you-go instances in the scaling group is less than the value of this parameter, Auto Scaling preferentially creates pay-as-you-go instances.
	//
	// example:
	//
	// 30
	OnDemandBaseCapacity *int32 `json:"OnDemandBaseCapacity,omitempty" xml:"OnDemandBaseCapacity,omitempty"`
	// The percentage of pay-as-you-go instances among the excess instances when the minimum number of pay-as-you-go instances reaches the requirement. Valid values: 0 to 100.
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
	// 0
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
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance removal policies.
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
	// >  If you specify this parameter, new scaling groups are added to the specified resource group. If you do not specify this parameter, new scaling groups are added to the default resource group.
	//
	// example:
	//
	// rg-aek2epf32c4uyji
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp14wlu85wrpchm0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The name of the scaling group. The name of each scaling group must be unique in a region.
	//
	// The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter or digit.
	//
	// example:
	//
	// dyrSuvBOtO1dEdIlIbp****
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	// The reclaim mode of the scaling group. Valid values:
	//
	// 	- recycle: economical mode
	//
	// 	- release: release mode
	//
	// 	- forcerelease: forced release mode
	//
	// 	- forcerecycle: forced recycle mode
	//
	// For more information, see [RemoveInstances](https://help.aliyun.com/document_detail/25955.html).
	//
	// example:
	//
	// recycle
	ScalingPolicy *string `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty"`
	// The information about the server groups.
	//
	// >  You can use this parameter to obtain information about ALB server groups and Network Load Balancer (NLB) server groups attached to your scaling group.
	ServerGroups []*DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups `json:"ServerGroups,omitempty" xml:"ServerGroups,omitempty" type:"Repeated"`
	// The allocation policy of preemptible instances. Auto Scaling selects instance types based on the allocation policy to create the required number of preemptible instances. You can apply the policy to pay-as-you-go instances and preemptible instances. This parameter takes effect only if you set `MultiAZPolicy` to `COMPOSABLE`. Valid values:
	//
	// 	- priority: Auto Scaling selects instance types based on the specified order to create the required number of preemptible instances.
	//
	// 	- lowestPrice: Auto Scaling selects instance types that have the lowest unit price of vCPUs to create the required number of preemptible instances.
	//
	// example:
	//
	// lowestPrice
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitempty" xml:"SpotAllocationStrategy,omitempty"`
	// The number of instance types that are specified. Preemptible instances of multiple lowest-priced instance types are evenly distributed across the zones of the scaling group. Valid values: 0 to 10.
	//
	// example:
	//
	// 5
	SpotInstancePools *int32 `json:"SpotInstancePools,omitempty" xml:"SpotInstancePools,omitempty"`
	// Indicates whether preemptible instances can be supplemented. If this parameter is set to true, Auto Scaling creates an instance to replace a preemptible instance when Auto Scaling receives the system message which indicates that the preemptible instance is to be reclaimed.
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
	// The number of instances that are stopped in Economical Mode in the scaling group.
	//
	// example:
	//
	// 1
	StoppedCapacity *int32 `json:"StoppedCapacity,omitempty" xml:"StoppedCapacity,omitempty"`
	// The processes that are suspended. If no process is suspended, null is returned. Valid values:
	//
	// 	- ScaleIn: scale-in
	//
	// 	- ScaleOut: scale-out
	//
	// 	- HealthCheck: health check
	//
	// 	- AlarmNotification: event-triggered task
	//
	// 	- ScheduledAction: scheduled task
	SuspendedProcesses []*string `json:"SuspendedProcesses,omitempty" xml:"SuspendedProcesses,omitempty" type:"Repeated"`
	// Indicates whether Auto Scaling stops executing scaling activities in the scaling group. Valid values:
	//
	// 	- true: Auto Scaling stops executing scaling activities in the scaling group if the scaling activities failed for more than seven consecutive days in the scaling group. You must modify the scaling group or scaling configuration to resume the execution of the scaling activities.
	//
	// 	- false: Auto Scaling does not stop executing scaling activities in the scaling group.
	//
	// example:
	//
	// true
	SystemSuspended *bool `json:"SystemSuspended,omitempty" xml:"SystemSuspended,omitempty"`
	// The tags of the scaling group.
	Tags []*DescribeScalingGroupDetailResponseBodyScalingGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total weighted capacity of all ECS instances in the scaling group if Weighted is specified. In other cases, the value of this parameter indicates the total number of ECS instances in the scaling group.
	//
	// example:
	//
	// 1
	TotalCapacity *int32 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The total number of Elastic Compute Service (ECS) instances in the scaling group.
	//
	// example:
	//
	// 1
	TotalInstanceCount *int32 `json:"TotalInstanceCount,omitempty" xml:"TotalInstanceCount,omitempty"`
	// The backend vServer groups.
	VServerGroups []*DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Repeated"`
	// The vSwitch ID of the scaling group.
	//
	// example:
	//
	// vsw-bp1whw2u46cn8zubm****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of the vSwitches that are associated with the scaling group. If you specify VSwitchIds, VSwitchId is ignored.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The virtual private cloud (VPC) ID of the scaling group.
	//
	// example:
	//
	// vpc-wz9fcq97y1vqkd8bijcq6
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroup) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetActiveCapacity() *int32 {
	return s.ActiveCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetActiveScalingConfigurationId() *string {
	return s.ActiveScalingConfigurationId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetAlbServerGroups() []*DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups {
	return s.AlbServerGroups
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetAllocationStrategy() *string {
	return s.AllocationStrategy
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetAzBalance() *bool {
	return s.AzBalance
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetCompensateWithOnDemand() *bool {
	return s.CompensateWithOnDemand
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetCurrentHostName() *string {
	return s.CurrentHostName
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetCustomPolicyARN() *string {
	return s.CustomPolicyARN
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetDBInstanceIds() []*string {
	return s.DBInstanceIds
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetDefaultCooldown() *int32 {
	return s.DefaultCooldown
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetDesiredCapacity() *int32 {
	return s.DesiredCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetEnableDesiredCapacity() *bool {
	return s.EnableDesiredCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetGroupDeletionProtection() *bool {
	return s.GroupDeletionProtection
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetGroupType() *string {
	return s.GroupType
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetHealthCheckType() *string {
	return s.HealthCheckType
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetHealthCheckTypes() []*string {
	return s.HealthCheckTypes
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetInitCapacity() *int32 {
	return s.InitCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetIsElasticStrengthInAlarm() *bool {
	return s.IsElasticStrengthInAlarm
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetLaunchTemplateOverrides() []*DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides {
	return s.LaunchTemplateOverrides
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetLoadBalancerConfigs() []*DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs {
	return s.LoadBalancerConfigs
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetLoadBalancerIds() []*string {
	return s.LoadBalancerIds
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetMaxInstanceLifetime() *int32 {
	return s.MaxInstanceLifetime
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetMinSize() *int32 {
	return s.MinSize
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetMonitorGroupId() *string {
	return s.MonitorGroupId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetMultiAZPolicy() *string {
	return s.MultiAZPolicy
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetOnDemandBaseCapacity() *int32 {
	return s.OnDemandBaseCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetOnDemandPercentageAboveBaseCapacity() *int32 {
	return s.OnDemandPercentageAboveBaseCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetPendingCapacity() *int32 {
	return s.PendingCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetPendingWaitCapacity() *int32 {
	return s.PendingWaitCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetProtectedCapacity() *int32 {
	return s.ProtectedCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetRemovalPolicies() []*string {
	return s.RemovalPolicies
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetRemovingCapacity() *int32 {
	return s.RemovingCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetRemovingWaitCapacity() *int32 {
	return s.RemovingWaitCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetScalingGroupName() *string {
	return s.ScalingGroupName
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetScalingPolicy() *string {
	return s.ScalingPolicy
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetServerGroups() []*DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups {
	return s.ServerGroups
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetSpotAllocationStrategy() *string {
	return s.SpotAllocationStrategy
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetSpotInstancePools() *int32 {
	return s.SpotInstancePools
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetSpotInstanceRemedy() *bool {
	return s.SpotInstanceRemedy
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetStandbyCapacity() *int32 {
	return s.StandbyCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetStoppedCapacity() *int32 {
	return s.StoppedCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetSuspendedProcesses() []*string {
	return s.SuspendedProcesses
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetSystemSuspended() *bool {
	return s.SystemSuspended
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetTags() []*DescribeScalingGroupDetailResponseBodyScalingGroupTags {
	return s.Tags
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetTotalCapacity() *int32 {
	return s.TotalCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetTotalInstanceCount() *int32 {
	return s.TotalInstanceCount
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetVServerGroups() []*DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups {
	return s.VServerGroups
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetActiveCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ActiveCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetActiveScalingConfigurationId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetAlbServerGroups(v []*DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.AlbServerGroups = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetAllocationStrategy(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.AllocationStrategy = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetAzBalance(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.AzBalance = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetCompensateWithOnDemand(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.CompensateWithOnDemand = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetCreationTime(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetCurrentHostName(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.CurrentHostName = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetCustomPolicyARN(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.CustomPolicyARN = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetDBInstanceIds(v []*string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetDefaultCooldown(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.DefaultCooldown = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetDesiredCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.DesiredCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetEnableDesiredCapacity(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.EnableDesiredCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetGroupDeletionProtection(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.GroupDeletionProtection = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetGroupType(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.GroupType = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetHealthCheckType(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.HealthCheckType = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetHealthCheckTypes(v []*string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.HealthCheckTypes = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetInitCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.InitCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetIsElasticStrengthInAlarm(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.IsElasticStrengthInAlarm = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetLaunchTemplateId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetLaunchTemplateOverrides(v []*DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.LaunchTemplateOverrides = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetLaunchTemplateVersion(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetLifecycleState(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetLoadBalancerConfigs(v []*DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.LoadBalancerConfigs = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetLoadBalancerIds(v []*string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.LoadBalancerIds = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetMaxInstanceLifetime(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.MaxInstanceLifetime = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetMaxSize(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetMinSize(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetModificationTime(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ModificationTime = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetMonitorGroupId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.MonitorGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetMultiAZPolicy(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.MultiAZPolicy = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetOnDemandBaseCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.OnDemandBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetOnDemandPercentageAboveBaseCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.OnDemandPercentageAboveBaseCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetPendingCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.PendingCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetPendingWaitCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.PendingWaitCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetProtectedCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ProtectedCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetRegionId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetRemovalPolicies(v []*string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.RemovalPolicies = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetRemovingCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.RemovingCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetRemovingWaitCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.RemovingWaitCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetResourceGroupId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetScalingGroupId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetScalingGroupName(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ScalingGroupName = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetScalingPolicy(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ScalingPolicy = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetServerGroups(v []*DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.ServerGroups = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetSpotAllocationStrategy(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.SpotAllocationStrategy = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetSpotInstancePools(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.SpotInstancePools = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetSpotInstanceRemedy(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.SpotInstanceRemedy = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetStandbyCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.StandbyCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetStoppedCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.StoppedCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetSuspendedProcesses(v []*string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.SuspendedProcesses = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetSystemSuspended(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.SystemSuspended = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetTags(v []*DescribeScalingGroupDetailResponseBodyScalingGroupTags) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.Tags = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetTotalCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetTotalInstanceCount(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.TotalInstanceCount = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetVServerGroups(v []*DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.VServerGroups = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetVSwitchId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.VSwitchId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetVSwitchIds(v []*string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.VSwitchIds = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) SetVpcId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroup {
	s.VpcId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroup) Validate() error {
	if s.AlbServerGroups != nil {
		for _, item := range s.AlbServerGroups {
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

type DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups struct {
	// The ID of the Application Load Balancer (ALB) server group.
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

func (s DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) GetAlbServerGroupId() *string {
	return s.AlbServerGroupId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) SetAlbServerGroupId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups {
	s.AlbServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) SetPort(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) SetWeight(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupAlbServerGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides struct {
	// The instance type. The instance type that is specified by using this parameter overwrites the instance type of the launch template.
	//
	// example:
	//
	// ecs.c5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum bid price of the instance type that is specified by `LaunchTemplateOverride.InstanceType`.
	//
	// >  This parameter takes effect only if you specify `LaunchTemplateId`.
	//
	// example:
	//
	// 0.025
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The weight of the instance type. The value of this parameter indicates the capacity of an instance of the specified instance type in the scaling group. A higher weight indicates that a smaller number of instances of the specified instance type are required to meet the expected capacity requirement.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) SetInstanceType(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) SetSpotPriceLimit(v float32) *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) SetWeightedCapacity(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides {
	s.WeightedCapacity = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLaunchTemplateOverrides) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs struct {
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-2zein3ytoeq49cmkbyxr0
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The weight of a backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) SetLoadBalancerId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) SetWeight(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupLoadBalancerConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups struct {
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

func (s DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) GetPort() *int32 {
	return s.Port
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) GetServerGroupId() *string {
	return s.ServerGroupId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) GetType() *string {
	return s.Type
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) SetPort(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) SetServerGroupId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups {
	s.ServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) SetType(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups {
	s.Type = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) SetWeight(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupServerGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupDetailResponseBodyScalingGroupTags struct {
	// Indicates whether the tags of the scaling group can be propagated to instances. Valid values:
	//
	// 	- true: The tags of the scaling group can be propagated to only instances that are newly created.
	//
	// 	- false: The tags of the scaling group cannot be propagated to any instances.
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

func (s DescribeScalingGroupDetailResponseBodyScalingGroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) GetPropagate() *bool {
	return s.Propagate
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) SetPropagate(v bool) *DescribeScalingGroupDetailResponseBodyScalingGroupTags {
	s.Propagate = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) SetTagKey(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupTags {
	s.TagKey = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) SetTagValue(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupTags {
	s.TagValue = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupTags) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups struct {
	// The ID of the Classic Load Balancer (CLB, formerly known as Server Load Balancer or SLB) instance to which the backend vServer group belongs.
	//
	// example:
	//
	// 147b46d767c-cn-qingdao-cm5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The attributes of the backend vServer group.
	VServerGroupAttributes []*DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes `json:"VServerGroupAttributes,omitempty" xml:"VServerGroupAttributes,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) GetVServerGroupAttributes() []*DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes {
	return s.VServerGroupAttributes
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) SetLoadBalancerId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) SetVServerGroupAttributes(v []*DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups {
	s.VServerGroupAttributes = v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroups) Validate() error {
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

type DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes struct {
	// The port number of a backend vServer.
	//
	// example:
	//
	// 80
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

func (s DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) GetPort() *int32 {
	return s.Port
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) SetPort(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes {
	s.Port = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) SetVServerGroupId(v string) *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) SetWeight(v int32) *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes {
	s.Weight = &v
	return s
}

func (s *DescribeScalingGroupDetailResponseBodyScalingGroupVServerGroupsVServerGroupAttributes) Validate() error {
	return dara.Validate(s)
}
