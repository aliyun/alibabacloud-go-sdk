// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationType(v string) *DescribeScalingInstancesRequest
	GetCreationType() *string
	SetCreationTypes(v []*string) *DescribeScalingInstancesRequest
	GetCreationTypes() []*string
	SetHealthStatus(v string) *DescribeScalingInstancesRequest
	GetHealthStatus() *string
	SetInstanceIds(v []*string) *DescribeScalingInstancesRequest
	GetInstanceIds() []*string
	SetLifecycleState(v string) *DescribeScalingInstancesRequest
	GetLifecycleState() *string
	SetLifecycleStates(v []*string) *DescribeScalingInstancesRequest
	GetLifecycleStates() []*string
	SetOwnerAccount(v string) *DescribeScalingInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScalingInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScalingInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeScalingInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScalingInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingInstancesRequest
	GetResourceOwnerId() *int64
	SetScalingActivityId(v string) *DescribeScalingInstancesRequest
	GetScalingActivityId() *string
	SetScalingConfigurationId(v string) *DescribeScalingInstancesRequest
	GetScalingConfigurationId() *string
	SetScalingGroupId(v string) *DescribeScalingInstancesRequest
	GetScalingGroupId() *string
}

type DescribeScalingInstancesRequest struct {
	// The instance creation method. Valid values:
	//
	// 	- AutoCreated: The ECS instances are created by Auto Scaling based on the instance configuration source.
	//
	// 	- Attached: The ECS instances are manually added to the scaling group.
	//
	// 	- Managed: The Alibaba Cloud-managed third-party instances are manually added to the scaling group.
	//
	// example:
	//
	// AutoCreated
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// The instance creation methods. If you specify this parameter, you cannot specify CreationType.
	CreationTypes []*string `json:"CreationTypes,omitempty" xml:"CreationTypes,omitempty" type:"Repeated"`
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
	// The IDs of the ECS instances.
	//
	// The IDs of inactive instances are not displayed in the query result, and no errors are returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
	// The lifecycle status of the ECS instances in the scaling group. You can specify only one of LifecycleStates and LifecycleState at a time. We recommend that you specify this parameter.
	LifecycleStates []*string `json:"LifecycleStates,omitempty" xml:"LifecycleStates,omitempty" type:"Repeated"`
	OwnerAccount    *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling activity.
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
}

func (s DescribeScalingInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesRequest) GetCreationType() *string {
	return s.CreationType
}

func (s *DescribeScalingInstancesRequest) GetCreationTypes() []*string {
	return s.CreationTypes
}

func (s *DescribeScalingInstancesRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeScalingInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeScalingInstancesRequest) GetLifecycleState() *string {
	return s.LifecycleState
}

func (s *DescribeScalingInstancesRequest) GetLifecycleStates() []*string {
	return s.LifecycleStates
}

func (s *DescribeScalingInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScalingInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingInstancesRequest) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeScalingInstancesRequest) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *DescribeScalingInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingInstancesRequest) SetCreationType(v string) *DescribeScalingInstancesRequest {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetCreationTypes(v []*string) *DescribeScalingInstancesRequest {
	s.CreationTypes = v
	return s
}

func (s *DescribeScalingInstancesRequest) SetHealthStatus(v string) *DescribeScalingInstancesRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetInstanceIds(v []*string) *DescribeScalingInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeScalingInstancesRequest) SetLifecycleState(v string) *DescribeScalingInstancesRequest {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetLifecycleStates(v []*string) *DescribeScalingInstancesRequest {
	s.LifecycleStates = v
	return s
}

func (s *DescribeScalingInstancesRequest) SetOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetPageNumber(v int32) *DescribeScalingInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetPageSize(v int32) *DescribeScalingInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetRegionId(v string) *DescribeScalingInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetResourceOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetResourceOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingActivityId(v string) *DescribeScalingInstancesRequest {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingConfigurationId(v string) *DescribeScalingInstancesRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingGroupId(v string) *DescribeScalingInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) Validate() error {
	return dara.Validate(s)
}
