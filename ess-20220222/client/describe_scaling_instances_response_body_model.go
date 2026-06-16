// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeScalingInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeScalingInstancesResponseBody
	GetRequestId() *string
	SetScalingInstances(v []*DescribeScalingInstancesResponseBodyScalingInstances) *DescribeScalingInstancesResponseBody
	GetScalingInstances() []*DescribeScalingInstancesResponseBodyScalingInstances
	SetTotalCount(v int32) *DescribeScalingInstancesResponseBody
	GetTotalCount() *int32
	SetTotalSpotCount(v int32) *DescribeScalingInstancesResponseBody
	GetTotalSpotCount() *int32
}

type DescribeScalingInstancesResponseBody struct {
	// The page number of the returned page.
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
	// B13527BF-1FBD-4334-A512-20F5E9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The collection of ECS instance information.
	ScalingInstances []*DescribeScalingInstancesResponseBodyScalingInstances `json:"ScalingInstances,omitempty" xml:"ScalingInstances,omitempty" type:"Repeated"`
	// The total number of ECS instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of running spot instances in the current scaling group.
	//
	// example:
	//
	// 4
	TotalSpotCount *int32 `json:"TotalSpotCount,omitempty" xml:"TotalSpotCount,omitempty"`
}

func (s DescribeScalingInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingInstancesResponseBody) GetScalingInstances() []*DescribeScalingInstancesResponseBodyScalingInstances {
	return s.ScalingInstances
}

func (s *DescribeScalingInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScalingInstancesResponseBody) GetTotalSpotCount() *int32 {
	return s.TotalSpotCount
}

func (s *DescribeScalingInstancesResponseBody) SetPageNumber(v int32) *DescribeScalingInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetPageSize(v int32) *DescribeScalingInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetRequestId(v string) *DescribeScalingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetScalingInstances(v []*DescribeScalingInstancesResponseBodyScalingInstances) *DescribeScalingInstancesResponseBody {
	s.ScalingInstances = v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalSpotCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalSpotCount = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) Validate() error {
	if s.ScalingInstances != nil {
		for _, item := range s.ScalingInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingInstancesResponseBodyScalingInstances struct {
	// The time when the ECS instance was added to the scaling group. The value is accurate to the second.
	//
	// example:
	//
	// 2020-05-18T03:11:39Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the ECS instance was added to the scaling group. The value is accurate to the minute.
	//
	// example:
	//
	// 2020-05-18T03:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The method used to create the ECS instance. Valid values:
	//
	// - AutoCreated: The ECS instance is created by automatic creation based on the instance configuration source in Auto Scaling.
	//
	// - Attached: The ECS instance is not created by Auto Scaling but manually added to the scaling group.
	//
	// example:
	//
	// AutoCreated
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// Indicates whether the manually added instance is managed by the scaling group. A managed manually added instance is released when it is removed from the scaling group (excluding manual removal). Valid values:
	//
	// - true: The instance is managed by the scaling group.
	//
	// - false: The instance is not managed by the scaling group.
	//
	// example:
	//
	// true
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// The health check status of the ECS instance in the scaling group. ECS instances that are not in the Running state are considered unhealthy. Valid values:
	//
	// - Healthy: The ECS instance is healthy.
	//
	// - Unhealthy: The ECS instance is unhealthy.
	//
	// Auto Scaling automatically removes unhealthy ECS instances from the scaling group and releases the ECS instances created by automatic creation.
	//
	// Whether a manually added ECS instance is released depends on its managed state. If the instance lifecycle is not managed by the scaling group, the instance is only removed but not released. If the instance lifecycle is managed by the scaling group, the instance is removed and released.
	//
	// > Make sure that your account has a sufficient available quota. If your account has an overdue payment, all pay-as-you-go ECS instances (including pay-as-you-go instances and spot instances) are stopped or even released. For information about how the status of ECS instances in a scaling group changes after an overdue payment occurs, see [Overdue payments](https://help.aliyun.com/document_detail/170589.html).
	//
	// example:
	//
	// Healthy
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp109k5j3dum1ce6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the launch template.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The lifecycle state of the ECS instance in the scaling group. Valid values:
	//
	//
	//
	// - InService: The ECS instance is added to the scaling group and provides services in the Normal state.
	//
	// - Pending: The ECS instance is being added to the scaling group. During this procedure, the ECS instance is added to the backend server group of the associated load balancing instance and to the access whitelist of the associated ApsaraDB RDS instance.
	//
	// - Pending:Wait: The ECS instance is waiting to be added to the scaling group. If a lifecycle hook that applies to scale-out activities is created for the scaling group, the ECS instance is suspended and waits for the lifecycle hook timeout to end.
	//
	// - Protected: The ECS instance is protected. The ECS instance provides services as expected, but Auto Scaling does not manage the lifecycle of the ECS instance. You must manually manage the lifecycle.
	//
	// - Standby: The ECS instance is in the standby state. The ECS instance does not provide services, the weight of SLB backend server is set to zero, and Auto Scaling does not manage the lifecycle of the ECS instance. You must manually manage the lifecycle.
	//
	// - Stopped: The ECS instance is stopped and does not provide services.
	//
	// - Removing: The ECS instance is being removed from the scaling group. During this procedure, the ECS instance is removed from the backend server group of the associated load balancing instance and from the access whitelist of the associated ApsaraDB RDS instance.
	//
	// - Removing:Wait: The ECS instance is waiting to be removed from the scaling group. If a lifecycle hook that applies to scale-down activities is created for the scaling group, the ECS instance is suspended and waits for the lifecycle hook timeout to end.
	//
	// example:
	//
	// InService
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The weight of the load balancing instance.
	//
	// > This parameter is deprecated and is not recommended.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The private IP address of the instance in the scaling group.
	//
	// example:
	//
	// 1**.2*.1**.2**
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	ReplaceStatus    *string `json:"ReplaceStatus,omitempty" xml:"ReplaceStatus,omitempty"`
	// The ID of the scaling activity during which the ECS instance was added to the scaling group.
	//
	// example:
	//
	// asa-bp1c9djwrgxjyk31****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	// The ID of the associated scaling configuration.
	//
	// example:
	//
	// asc-bp1i65jd06v04vdh****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The ID of the scaling group to which the instance belongs.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The instance identity in the scaling group, which has a one-to-one mapping with the ECS instance ID or Elastic Container Instance (ECI) instance identity.
	//
	// example:
	//
	// asi-j6cj1gcte640ekhb****
	ScalingInstanceId *string `json:"ScalingInstanceId,omitempty" xml:"ScalingInstanceId,omitempty"`
	// The preemption policy of the spot instance. Valid values:
	//
	// - SpotWithPriceLimit: The spot instance has a maximum price limit.
	//
	// - SpotAsPriceGo: The system automatically bids at the current market price.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The warmup state of the ECS instance. Valid values:
	//
	//
	//
	// - NoNeedWarmup: No warmup is required.
	//
	// - WaitingForInstanceWarmup: The instance is waiting for warmup to complete.
	//
	// - InstanceWarmupFinish: Warmup is complete.
	//
	// example:
	//
	// NoNeedWarmup
	WarmupState *string `json:"WarmupState,omitempty" xml:"WarmupState,omitempty"`
	// The weight of the instance type. The weight indicates the capacity that a single instance of this instance type represents in the scaling group. A higher weight means that fewer instances of this type are required to meet the expected capacity.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
	// The zone ID of the ECS instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetCreationType() *string {
	return s.CreationType
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetEntrusted() *bool {
	return s.Entrusted
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetLaunchTemplateVersion() *string {
	return s.LaunchTemplateVersion
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetLoadBalancerWeight() *int32 {
	return s.LoadBalancerWeight
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetReplaceStatus() *string {
	return s.ReplaceStatus
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetScalingInstanceId() *string {
	return s.ScalingInstanceId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetWarmupState() *string {
	return s.WarmupState
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetWeightedCapacity() *int32 {
	return s.WeightedCapacity
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetCreatedTime(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.CreatedTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetCreationTime(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetCreationType(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetEntrusted(v bool) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.Entrusted = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetHealthStatus(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetInstanceId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLaunchTemplateId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LaunchTemplateId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLaunchTemplateVersion(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLifecycleState(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetLoadBalancerWeight(v int32) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.LoadBalancerWeight = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetPrivateIpAddress(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetReplaceStatus(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ReplaceStatus = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingActivityId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingConfigurationId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingGroupId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingInstanceId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingInstanceId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetSpotStrategy(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetWarmupState(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.WarmupState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetWeightedCapacity(v int32) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.WeightedCapacity = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetZoneId(v string) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) Validate() error {
	return dara.Validate(s)
}
