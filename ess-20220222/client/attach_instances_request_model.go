// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachInstancesRequest
	GetClientToken() *string
	SetEntrusted(v bool) *AttachInstancesRequest
	GetEntrusted() *bool
	SetIgnoreInvalidInstance(v bool) *AttachInstancesRequest
	GetIgnoreInvalidInstance() *bool
	SetInstanceIds(v []*string) *AttachInstancesRequest
	GetInstanceIds() []*string
	SetLifecycleHook(v bool) *AttachInstancesRequest
	GetLifecycleHook() *bool
	SetLoadBalancerWeights(v []*int32) *AttachInstancesRequest
	GetLoadBalancerWeights() []*int32
	SetOwnerAccount(v string) *AttachInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachInstancesRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *AttachInstancesRequest
	GetScalingGroupId() *string
}

type AttachInstancesRequest struct {
	// Ensures the idempotence of the request.
	//
	// You can generate a parameter value from your client to ensure that the value is unique among different requests. The value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to entrust the lifecycle of an existing instance to the scaling group when the instance is manually added to the scaling group. Valid values:
	//
	// - true: Entrusted. The lifecycle of the instance is managed by Auto Scaling, which is the same as that of instances automatically created by the scaling group. The instance is automatically released when it is removed from the scaling group (excluding instances removed when you invoke the DetachInstances operation).
	//
	// - false: Not entrusted. The instance is not released when it is removed from the scaling group.
	//
	// > This parameter is not supported for subscription instances, managed instances, and instances in the Economical Mode in the scaling group.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// Specifies whether to ignore invalid instances when a batch of instances is added to the scaling group. Valid values:
	//
	// - true: Invalid instances are ignored when a batch of instances is added to the scaling group. If invalid instances exist and valid instances are successfully added, the execution status of the scaling activity is displayed as Warning. You can view the invalid instances in the scaling activity details.
	//
	// - false: Invalid instances are not ignored when a batch of instances is added to the scaling group. If invalid instances exist in the batch, an error is reported.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IgnoreInvalidInstance *bool `json:"IgnoreInvalidInstance,omitempty" xml:"IgnoreInvalidInstance,omitempty"`
	// The IDs of the ECS instances, ECI instances, managed instances, or instances in the Economical Mode in the scaling group to be added.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to trigger the scale-out lifecycle hook when instances are added to the scaling group. Valid values:
	//
	// - true: Triggered.
	//
	// - false: Not triggered.
	//
	// > This parameter is not supported for managed instances and instances in the Economical Mode in the scaling group.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	LifecycleHook *bool `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty"`
	// The weight of the ECS instance or ECI instance as a backend server of the load balancing instance.
	LoadBalancerWeights []*int32 `json:"LoadBalancerWeights,omitempty" xml:"LoadBalancerWeights,omitempty" type:"Repeated"`
	OwnerAccount        *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachInstancesRequest) GetEntrusted() *bool {
	return s.Entrusted
}

func (s *AttachInstancesRequest) GetIgnoreInvalidInstance() *bool {
	return s.IgnoreInvalidInstance
}

func (s *AttachInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachInstancesRequest) GetLifecycleHook() *bool {
	return s.LifecycleHook
}

func (s *AttachInstancesRequest) GetLoadBalancerWeights() []*int32 {
	return s.LoadBalancerWeights
}

func (s *AttachInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *AttachInstancesRequest) SetClientToken(v string) *AttachInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachInstancesRequest) SetEntrusted(v bool) *AttachInstancesRequest {
	s.Entrusted = &v
	return s
}

func (s *AttachInstancesRequest) SetIgnoreInvalidInstance(v bool) *AttachInstancesRequest {
	s.IgnoreInvalidInstance = &v
	return s
}

func (s *AttachInstancesRequest) SetInstanceIds(v []*string) *AttachInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachInstancesRequest) SetLifecycleHook(v bool) *AttachInstancesRequest {
	s.LifecycleHook = &v
	return s
}

func (s *AttachInstancesRequest) SetLoadBalancerWeights(v []*int32) *AttachInstancesRequest {
	s.LoadBalancerWeights = v
	return s
}

func (s *AttachInstancesRequest) SetOwnerAccount(v string) *AttachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetOwnerId(v int64) *AttachInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachInstancesRequest) SetRegionId(v string) *AttachInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *AttachInstancesRequest) SetResourceOwnerAccount(v string) *AttachInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetResourceOwnerId(v int64) *AttachInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachInstancesRequest) SetScalingGroupId(v string) *AttachInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *AttachInstancesRequest) Validate() error {
	return dara.Validate(s)
}
