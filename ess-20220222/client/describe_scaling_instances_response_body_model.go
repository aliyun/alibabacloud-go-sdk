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
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
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
	// The total number of running spot instances in the scaling group.
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
	// The time when the ECS instance was added to the scaling group. The time is accurate to seconds and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-18T03:11:39Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the ECS instance was added to the scaling group. The time is accurate to minutes and follows the ISO 8601 standard in the YYYY-MM-DDTHH:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-18T03:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The creation type of the ECS instance. Valid values:
	//
	// - AutoCreated: The ECS instance is created by Auto Scaling through automatic creation based on the instance configuration source.
	//
	// - Attached: The ECS instance is not created by Auto Scaling, but is manually added to the scaling group.
	//
	// example:
	//
	// AutoCreated
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// Specifies whether the instance is managed by the scaling group when it is manually added to the scaling group. A manually added instance in the managed state is released when it is removed from the scaling group, excluding manual removal. Valid values:
	//
	// - true: The instance is managed by the scaling group.
	//
	// - false: The instance is not managed by the scaling group.
	//
	// example:
	//
	// true
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// The health status of the ECS instance in the scaling group. ECS instances that are not in the Running state are considered unhealthy. Valid values:
	//
	// - Healthy: The ECS instance is healthy.
	//
	// - Unhealthy: The ECS instance is unhealthy.
	//
	// Auto Scaling automatically removes unhealthy ECS instances from the scaling group and releases ECS instances created through automatic creation.
	//
	// Whether to release manually added ECS instances depends on their managed state. If the instance lifecycle is not managed by the scaling group, the instance is only removed but not released. If the instance lifecycle is managed by the scaling group, the instance is removed and released.
	//
	// > Make sure that your account has sufficient active balance. If your account has overdue payments, all pay-as-you-go ECS instances, including pay-as-you-go instances and spot instances, are stopped or even released. For more information about the instance status changes of ECS instances in a scaling group after overdue payments, see [Overdue payments](https://help.aliyun.com/document_detail/170589.html).
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
	// The ID of the instance launch template.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The version of the instance launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The lifecycle state of the ECS instance in the scaling group. Valid values:
	//
	//
	//
	// - InService: The ECS instance is added to the scaling group and provides services normally.
	//
	// - Pending: The ECS instance is being added to the scaling group. The procedure of adding an ECS instance to a scaling group includes adding the instance to the backend server group of a load balancing instance and adding the IP address of the instance to the access whitelist of an ApsaraDB RDS instance.
	//
	// - Pending:Wait: The addition of the ECS instance to the scaling group is suspended. If a lifecycle hook applicable to scale-out activities is created in the scaling group, the ECS instance is suspended when it is added to the scaling group and waits for the hook timeout to end.
	//
	// - Protected: The ECS instance provides services normally, but Auto Scaling does not manage the lifecycle of the ECS instance. You must manually manage the lifecycle.
	//
	// - Standby: The ECS instance does not provide services, the weight of SLB backend server is set to zero, and Auto Scaling does not manage the lifecycle of the ECS instance. You must manually manage the lifecycle.
	//
	// - Stopped: The ECS instance is stopped and does not provide services.
	//
	// - Removing: The ECS instance is being removed from the scaling group. The procedure of removing an ECS instance from a scaling group includes removing the instance from the backend server group of a load balancing instance and removing the IP address of the instance from the access whitelist of an ApsaraDB RDS instance.
	//
	// - Removing:Wait: The removal of the ECS instance from the scaling group is suspended. If a lifecycle hook applicable to scale-down activities is created in the scaling group, the ECS instance is suspended when it is removed from the scaling group and waits for the hook timeout to end.
	//
	// example:
	//
	// InService
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The weight of the SLB backend server.
	//
	// > **[Deprecated]*	- This parameter is deprecated. Do not use it.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The internal IP address of the instance in the scaling group.
	//
	// example:
	//
	// 1**.2*.1**.2**
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The instance replacement process that the instance is in. Valid values:
	//
	// - Normal: The instance is not in any replacement process.
	//
	// - ReBalance: The instance is in the rebalancing replacement process of the scaling group.
	//
	// - SpotRemedy: The instance is in the spot instance remedy replacement process.
	//
	// - MaxInstanceLifetimeExpire: The instance is in the replacement process due to the expiration of the maximum instance lifetime.
	//
	// - InstanceRefresh: The instance is in the instance refresh replacement process.
	//
	// - SpotReplaceOnDemand: The instance is in the process of replacing a spot instance with a pay-as-you-go instance.
	//
	// example:
	//
	// ReBalance
	ReplaceStatus *string `json:"ReplaceStatus,omitempty" xml:"ReplaceStatus,omitempty"`
	// The ID of the scaling activity when the ECS instance was added to the scaling group.
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
	// The instance identity in the scaling group, which corresponds one-to-one with the ID of the ECS instance or elastic container instance.
	//
	// example:
	//
	// asi-j6cj1gcte640ekhb****
	ScalingInstanceId *string `json:"ScalingInstanceId,omitempty" xml:"ScalingInstanceId,omitempty"`
	// The bidding policy for the spot instance. Valid values:
	//
	// - SpotWithPriceLimit: The spot instance with a maximum price.
	//
	// - SpotAsPriceGo: The system automatically bids based on the current market price.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The warmup status of the ECS instance. Valid values:
	//
	//
	//
	// - NoNeedWarmup: No warmup is required.
	//
	// - WaitingForInstanceWarmup: The instance is waiting for the warmup to complete.
	//
	// - InstanceWarmupFinish: The warmup is complete.
	//
	// example:
	//
	// NoNeedWarmup
	WarmupState *string `json:"WarmupState,omitempty" xml:"WarmupState,omitempty"`
	// The weight of the instance type, which indicates the capacity of a single instance of the instance type in the scaling group. A greater weight indicates that fewer instances of the instance type are required to meet the expected capacity.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
	// The ID of the zone to which the ECS instance belongs.
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
