// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachInstancesRequest
	GetClientToken() *string
	SetDecreaseDesiredCapacity(v bool) *DetachInstancesRequest
	GetDecreaseDesiredCapacity() *bool
	SetDetachOption(v string) *DetachInstancesRequest
	GetDetachOption() *string
	SetIgnoreInvalidInstance(v bool) *DetachInstancesRequest
	GetIgnoreInvalidInstance() *bool
	SetInstanceIds(v []*string) *DetachInstancesRequest
	GetInstanceIds() []*string
	SetLifecycleHook(v bool) *DetachInstancesRequest
	GetLifecycleHook() *bool
	SetOwnerAccount(v string) *DetachInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachInstancesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DetachInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachInstancesRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DetachInstancesRequest
	GetScalingGroupId() *string
}

type DetachInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to adjust the expected number of instances in the scaling group. Valid values:
	//
	// 	- true: After a specific number of instances are removed from the scaling group, the expected number of instances in the scaling group decreases.
	//
	// 	- false: After a specific number of instances are removed from the scaling group, the expected number of instances in the scaling group remains unchanged.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DecreaseDesiredCapacity *bool `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	// Specifies whether to detach the ECS instances or elastic container instances that are marked for removal from the associated load balancers, and whether to remove the private IP addresses of these instances from the IP address whitelists of the associated ApsaraDB RDS instances.
	//
	// Both: detaches the ECS instances or elastic container instances that are marked for removal from the associated load balancers and removes the private IP addresses of these instances from the IP address whitelists of the associated ApsaraDB RDS instances.
	//
	// >  This parameter is not supported if you want to remove Alibaba Cloud-hosted third-party instances from a scaling group.
	//
	// example:
	//
	// both
	DetachOption *string `json:"DetachOption,omitempty" xml:"DetachOption,omitempty"`
	// Specifies whether to ignore invalid instances when you remove a batch of instances from the scaling group. Valid values:
	//
	// 	- true: ignores invalid instances. If invalid instances exist and valid instances are removed from the scaling group, the corresponding scaling activity enters the Warning state. You can check the scaling activity details to view the invalid instances that are ignored.
	//
	// 	- false: does not ignore invalid instances. If invalid instances exist in the batch of instances that you want to remove from the scaling group, an error is reported.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IgnoreInvalidInstance *bool `json:"IgnoreInvalidInstance,omitempty" xml:"IgnoreInvalidInstance,omitempty"`
	// The IDs of the ECS instances, elastic container instances, or Aliababa Cloud-managed third-party instances that you want to remove from a scaling group.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// Specifies whether to trigger a lifecycle hook for scale-in purposes when ECS instances or elastic container instances are removed from the scaling group. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is not supported if you want to remove Alibaba Cloud-hosted third-party instances from a scaling group.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	LifecycleHook        *bool   `json:"LifecycleHook,omitempty" xml:"LifecycleHook,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachInstancesRequest) GoString() string {
	return s.String()
}

func (s *DetachInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachInstancesRequest) GetDecreaseDesiredCapacity() *bool {
	return s.DecreaseDesiredCapacity
}

func (s *DetachInstancesRequest) GetDetachOption() *string {
	return s.DetachOption
}

func (s *DetachInstancesRequest) GetIgnoreInvalidInstance() *bool {
	return s.IgnoreInvalidInstance
}

func (s *DetachInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DetachInstancesRequest) GetLifecycleHook() *bool {
	return s.LifecycleHook
}

func (s *DetachInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DetachInstancesRequest) SetClientToken(v string) *DetachInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachInstancesRequest) SetDecreaseDesiredCapacity(v bool) *DetachInstancesRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *DetachInstancesRequest) SetDetachOption(v string) *DetachInstancesRequest {
	s.DetachOption = &v
	return s
}

func (s *DetachInstancesRequest) SetIgnoreInvalidInstance(v bool) *DetachInstancesRequest {
	s.IgnoreInvalidInstance = &v
	return s
}

func (s *DetachInstancesRequest) SetInstanceIds(v []*string) *DetachInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachInstancesRequest) SetLifecycleHook(v bool) *DetachInstancesRequest {
	s.LifecycleHook = &v
	return s
}

func (s *DetachInstancesRequest) SetOwnerAccount(v string) *DetachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetOwnerId(v int64) *DetachInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachInstancesRequest) SetResourceOwnerAccount(v string) *DetachInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetResourceOwnerId(v int64) *DetachInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachInstancesRequest) SetScalingGroupId(v string) *DetachInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DetachInstancesRequest) Validate() error {
	return dara.Validate(s)
}
