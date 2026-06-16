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
	// The method used to create the instance in the scaling group. Valid values:
	//
	// - AutoCreated: The ECS instance is created by automatic creation based on the instance configuration source in Auto Scaling.
	//
	// - Attached: The ECS instance is not created by Auto Scaling but manually added to the scaling group.
	//
	// - Managed: The managed instance is not created by Auto Scaling but manually added to the scaling group.
	//
	// example:
	//
	// AutoCreated
	CreationType *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	// The methods used to create instances in the scaling group. You can specify only one of this parameter and CreationType.
	CreationTypes []*string `json:"CreationTypes,omitempty" xml:"CreationTypes,omitempty" type:"Repeated"`
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
	// The IDs of the ECS instances.
	//
	// Invalid InstanceId values are ignored in the query results, and no error is returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
	// The lifecycle states of ECS instances in the scaling group. You can specify only one of this parameter and LifecycleState. We recommend that you use this parameter.
	LifecycleStates []*string `json:"LifecycleStates,omitempty" xml:"LifecycleStates,omitempty" type:"Repeated"`
	OwnerAccount    *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the ECS instance list. Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Settings: Maximum value: 100.
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
	// The ID of the associated scaling configuration.
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
