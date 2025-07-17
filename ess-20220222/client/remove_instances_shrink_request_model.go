// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveInstancesShrinkRequest
	GetClientToken() *string
	SetDecreaseDesiredCapacity(v bool) *RemoveInstancesShrinkRequest
	GetDecreaseDesiredCapacity() *bool
	SetIgnoreInvalidInstance(v bool) *RemoveInstancesShrinkRequest
	GetIgnoreInvalidInstance() *bool
	SetInstanceIds(v []*string) *RemoveInstancesShrinkRequest
	GetInstanceIds() []*string
	SetLifecycleHookContextShrink(v string) *RemoveInstancesShrinkRequest
	GetLifecycleHookContextShrink() *string
	SetOwnerAccount(v string) *RemoveInstancesShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveInstancesShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveInstancesShrinkRequest
	GetRegionId() *string
	SetRemovePolicy(v string) *RemoveInstancesShrinkRequest
	GetRemovePolicy() *string
	SetResourceOwnerAccount(v string) *RemoveInstancesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveInstancesShrinkRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *RemoveInstancesShrinkRequest
	GetScalingGroupId() *string
	SetStopInstanceTimeout(v int32) *RemoveInstancesShrinkRequest
	GetStopInstanceTimeout() *int32
}

type RemoveInstancesShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to adjust the expected number of ECS instances in the scaling group. Valid values:
	//
	// 	- true: After ECS instances are removed from the scaling group, the expected number of ECS instances in the scaling group decreases.
	//
	// 	- false: After ECS instances are removed from the scaling group, the expected number of ECS instances in the scaling group remains unchanged.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DecreaseDesiredCapacity *bool `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	// Specifies whether to ignore invalid instances when you remove a batch of instances from the scaling group. Valid values:
	//
	// 	- true: ignores invalid instances. If invalid instances exist and valid instances are deleted, the corresponding scaling activity enters the Warning state. You can check the scaling activity details to view the invalid instances that are ignored.
	//
	// 	- false: does not ignore invalid instances. If invalid instances exist in the batch of instances that you want to remove from the scaling group, an error is reported.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IgnoreInvalidInstance *bool `json:"IgnoreInvalidInstance,omitempty" xml:"IgnoreInvalidInstance,omitempty"`
	// The IDs of the ECS instances that you want to remove from the scaling group.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The context of the lifecycle hook.
	LifecycleHookContextShrink *string `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty"`
	OwnerAccount               *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The action subsequent to the removal of the Elastic Compute Service (ECS) instances. Valid values:
	//
	// 	- recycle: The ECS instances enter the Economical Mode.
	//
	//     **
	//
	//     **Note*	- This setting is applicable only if you set `ScalingPolicy` to `recycle`.
	//
	// 	- release: The ECS instances are released.
	//
	// ScalingPolicy of the CreateScalingGroup operation specifies the reclaim mode of the scaling group while RemovePolicy of the RemoveInstances operation specifies the subsequent action when an ECS instance is removed from the scaling group. Examples:
	//
	// 	- If you set ScalingPolicy and RemovePolicy to recycle, the ECS instances enter the Economical Mode when they are removed.
	//
	// 	- If you set ScalingPolicy to recycle and RemovePolicy to release, the ECS instances are released when they are removed.
	//
	// 	- If you set ScalingPolicy to release and RemovePolicy to recycle, the ECS instances are released when they are removed.
	//
	// 	- If you set ScalingPolicy and RemovePolicy to release, the ECS instances are released when they are removed.
	//
	// Default value: release.
	//
	// example:
	//
	// release
	RemovePolicy         *string `json:"RemovePolicy,omitempty" xml:"RemovePolicy,omitempty"`
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
	// The period of time required by the ECS instance to enter the Stopped state. Unit: seconds. Valid values: 30 to 240.
	//
	// >
	//
	// 	- By default, this parameter inherits the value of StopInstanceTimeout specified in the CreateScalingGroup or ModifyScalingGroup operation. You can also specify a different value for this parameter in the RemoveInstances operation.
	//
	// 	- This parameter takes effect only if you set RemovePolicy to release.
	//
	// 	- If you specify this parameter, the system waits for the ECS instance to enter the Stopped state only for up to the specified period of time before continuing with the scale-in operation, regardless of the status of the ECS instance.
	//
	// 	- If you do not specify this parameter, the system continues with the scale-in operation until the ECS instance enters the Stopped state. If the ECS instance is not successfully stopped, the scale-in process is rolled back and considered failed.
	//
	// example:
	//
	// 60
	StopInstanceTimeout *int32 `json:"StopInstanceTimeout,omitempty" xml:"StopInstanceTimeout,omitempty"`
}

func (s RemoveInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstancesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveInstancesShrinkRequest) GetDecreaseDesiredCapacity() *bool {
	return s.DecreaseDesiredCapacity
}

func (s *RemoveInstancesShrinkRequest) GetIgnoreInvalidInstance() *bool {
	return s.IgnoreInvalidInstance
}

func (s *RemoveInstancesShrinkRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveInstancesShrinkRequest) GetLifecycleHookContextShrink() *string {
	return s.LifecycleHookContextShrink
}

func (s *RemoveInstancesShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveInstancesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveInstancesShrinkRequest) GetRemovePolicy() *string {
	return s.RemovePolicy
}

func (s *RemoveInstancesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveInstancesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveInstancesShrinkRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *RemoveInstancesShrinkRequest) GetStopInstanceTimeout() *int32 {
	return s.StopInstanceTimeout
}

func (s *RemoveInstancesShrinkRequest) SetClientToken(v string) *RemoveInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetDecreaseDesiredCapacity(v bool) *RemoveInstancesShrinkRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetIgnoreInvalidInstance(v bool) *RemoveInstancesShrinkRequest {
	s.IgnoreInvalidInstance = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetInstanceIds(v []*string) *RemoveInstancesShrinkRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetLifecycleHookContextShrink(v string) *RemoveInstancesShrinkRequest {
	s.LifecycleHookContextShrink = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetOwnerAccount(v string) *RemoveInstancesShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetOwnerId(v int64) *RemoveInstancesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetRegionId(v string) *RemoveInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetRemovePolicy(v string) *RemoveInstancesShrinkRequest {
	s.RemovePolicy = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetResourceOwnerAccount(v string) *RemoveInstancesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetResourceOwnerId(v int64) *RemoveInstancesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetScalingGroupId(v string) *RemoveInstancesShrinkRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) SetStopInstanceTimeout(v int32) *RemoveInstancesShrinkRequest {
	s.StopInstanceTimeout = &v
	return s
}

func (s *RemoveInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
