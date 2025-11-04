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
	// B13527BF-1FBD-4334-A512-20F5E9D3FB4D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ECS instances.
	ScalingInstances []*DescribeScalingInstancesResponseBodyScalingInstances `json:"ScalingInstances,omitempty" xml:"ScalingInstances,omitempty" type:"Repeated"`
	// The total number of ECS instances in the scaling group.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of preemptible instances that run as expected in the scaling group.
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
	// The time when the ECS instances were added to the scaling group. The value is accurate to the second.
	//
	// example:
	//
	// 2020-05-18T03:11:39Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the ECS instances were added to the scaling group. The value is accurate to the minute.
	//
	// example:
	//
	// 2020-05-18T03:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance creation method. Valid values:
	//
	// 	- AutoCreated: The ECS instances are created by Auto Scaling based on the instance configuration source.
	//
	// 	- Attached: The ECS instances are manually added to the scaling group.
	//
	// example:
	//
	// AutoCreated
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// Indicates whether the scaling group is allowed to manage the instance lifecycles when ECS instances are manually added. If the scaling group is allowed to manage the instance lifecycles, Auto Scaling can release the ECS instances when the instances are automatically removed from the scaling group. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// The health status of the ECS instance in the scaling group. If an ECS instance is not in the Running state, the instance is considered unhealthy. Valid values:
	//
	// 	- Healthy
	//
	// 	- Unhealthy
	//
	// Auto Scaling automatically removes unhealthy ECS instances from the scaling group and then releases the automatically created instances among the unhealthy instances.
	//
	// Unhealthy ECS instances that are manually added to the scaling group are released based on the management mode of the lifecycles of the instances. If the lifecycles of the ECS instances are not managed by the scaling group, Auto Scaling removes the instances from the scaling group but does not release the instances. If the lifecycles of the ECS instances are managed by the scaling group, Auto Scaling removes the instances from the scaling group and releases the instances.
	//
	// >  Make sure that you have sufficient balance within your Alibaba Cloud account. If your Alibaba Cloud account has an overdue payment, all pay-as-you-go ECS instances, including preemptible instances, may be stopped or even released. For information about how the status of ECS instances changes when you have an overdue payment in your Alibaba Cloud account, see [Overdue payments](https://help.aliyun.com/document_detail/170589.html).
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
	// The version number of the launch template.
	//
	// example:
	//
	// 1
	LaunchTemplateVersion *string `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The lifecycle status of the ECS instance in the scaling group. Valid values:
	//
	// 	- InService: The ECS instance is added to the scaling group and provides services as expected.
	//
	// 	- Pending: The ECS instance is being added to the scaling group. When an ECS instance is being added to the scaling group, Auto Scaling also adds the instance to the backend server groups of the attached load balancers and adds the private IP address of the instance to the IP address whitelists of the attached ApsaraDB RDS instances.
	//
	// 	- Pending:Wait: The ECS instance is waiting to be added to the scaling group. If a scale-out lifecycle hook is in effect, the ECS instance remains in the Pending:Wait state until the timeout period for the lifecycle hook expires.
	//
	// 	- Protected: The ECS instance is protected. Protected ECS instances can continue to provide services as expected, but Auto Scaling does not manage the lifecycles of the ECS instances. You must manually manage the lifecycles of the ECS instances.
	//
	// 	- Standby: The ECS instance is on standby. Standby ECS instances do not provide services as expected, and the weights of the ECS instances as backend servers are reset to zero. Auto Scaling does not manage the lifecycles of the ECS instances. Therefore, you must manually manage the lifecycles of the ECS instances.
	//
	// 	- Stopped: The ECS instance is stopped. Stopped ECS instances no longer provide services.
	//
	// 	- Removing: The ECS instance is being removed from the scaling group. When an ECS instance is being removed from the scaling group, Auto Scaling also removes the instance from the backend server groups of the attached load balancers and removes the private IP address of the instance from the IP address whitelists of the attached ApsaraDB RDS instances.
	//
	// 	- Removing:Wait: The ECS instance is waiting to be removed from the scaling group. If a scale-in lifecycle hook is in effect, the ECS instance remains in the Removing:Wait state until the timeout period for the lifecycle hook expires.
	//
	// example:
	//
	// InService
	LifecycleState *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	// The weight of each ECS instance as a backend server.
	//
	// >  This parameter is deprecated and is not recommended.
	//
	// example:
	//
	// 50
	LoadBalancerWeight *int32 `json:"LoadBalancerWeight,omitempty" xml:"LoadBalancerWeight,omitempty"`
	// The private IP address of the ECS instance.
	//
	// example:
	//
	// 1**.2*.1**.2**
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The ID of the scaling activity during which the ECS instances were added to the scaling group.
	//
	// example:
	//
	// asa-bp1c9djwrgxjyk31****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-bp1i65jd06v04vdh****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The ID of the ECS instance or elastic container instance.
	//
	// example:
	//
	// asi-j6cj1gcte640ekhb****
	ScalingInstanceId *string `json:"ScalingInstanceId,omitempty" xml:"ScalingInstanceId,omitempty"`
	// The bidding policy for the preemptible instances. Valid values:
	//
	// 	- SpotWithPriceLimit: The instances are preemptible instances that have a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instances are preemptible instances for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The warm-up status of the ECS instances. Valid values:
	//
	// 	- NoNeedWarmup: The ECS instances do not need to undergo a warm-up process.
	//
	// 	- WaitingForInstanceWarmup: The ECS instances are undergoing the warm-up process.
	//
	// 	- InstanceWarmupFinish: The warm-up process for the ECS instances is complete.
	//
	// example:
	//
	// NoNeedWarmup
	WarmupState *string `json:"WarmupState,omitempty" xml:"WarmupState,omitempty"`
	// The weight of the instance type. The weight indicates the capacity of a single instance of the specified instance type in the scaling group. A higher weight indicates that a smaller number of instances of the instance type are required to meet the expected capacity requirement.
	//
	// example:
	//
	// 4
	WeightedCapacity *int32 `json:"WeightedCapacity,omitempty" xml:"WeightedCapacity,omitempty"`
	// The zone ID of the ECS instances.
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
