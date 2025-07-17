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
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to use the scaling group to manage the lifecycles of manually added instances. Valid values:
	//
	// 	- true: The scaling group manages the lifecycles of manually added instances and automatically created instances in the same manner. In this case, Auto Scaling releases the instances when they are removed from the scaling group. This rule does not apply to instances that are removed by calling the DetachInstances operation.
	//
	// 	- false: The scaling group does not manage the lifecycles of manually added instances. In this case, Auto Scaling does not release the instances when they are removed from the scaling group.
	//
	// >  You cannot specify this parameter for subscription instances, non-Alibaba Cloud instances, and instances in Economical Mode.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Entrusted *bool `json:"Entrusted,omitempty" xml:"Entrusted,omitempty"`
	// Specifies whether to ignore invalid instances when a batch of instances is added to the scaling group. Valid values:
	//
	// 	- true: ignores invalid instances. If invalid instances exist and valid instances are added, the corresponding scaling activity enters the Warning state. You can check the scaling activity details to view the invalid instances that are ignored.
	//
	// 	- false: does not ignore invalid instances. If invalid instances exist in the batch of instances that you want to add to the scaling group, an error is reported.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IgnoreInvalidInstance *bool `json:"IgnoreInvalidInstance,omitempty" xml:"IgnoreInvalidInstance,omitempty"`
	// The IDs of the ECS instances, elastic container instances, non-Alibaba Cloud instances, or instances in Economical Mode.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to trigger the lifecycle hook for scale-outs when you call this operation. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  You cannot specify this parameter for subscription instances and instances in Economical Mode.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	LifecycleHook *bool `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty"`
	// The weight of an ECS instance or elastic container instance as a backend server. You can use this parameter to specify weights for multiple instances at the same time.
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
