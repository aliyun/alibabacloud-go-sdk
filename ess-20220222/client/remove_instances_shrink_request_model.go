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
	// A client-generated token to ensure request idempotence. This token must be unique for each request, contain only ASCII characters, and not exceed 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to decrease the desired capacity of the scaling group. Valid values:
	//
	// - `true`: Decreases the desired capacity of the scaling group by the number of removed instances.
	//
	// - `false`: The desired capacity of the scaling group remains unchanged.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DecreaseDesiredCapacity *bool `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	// Specifies whether to ignore invalid instances when you remove multiple instances from a scaling group. Valid values:
	//
	// - `true`: Invalid instances are ignored. If valid instances are removed while invalid ones are present, the scaling activity status is set to Warning. The invalid instances are listed in the scaling activity details.
	//
	// - `false`: The request is rejected and an error is returned if it contains any invalid instances.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IgnoreInvalidInstance *bool `json:"IgnoreInvalidInstance,omitempty" xml:"IgnoreInvalidInstance,omitempty"`
	// The IDs of the ECS instances to remove.
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
	// Specifies the action to take on removed ECS instances. Valid values:
	//
	// - recycle: The ECS instances enter the economical mode.
	//
	//   > This value takes effect only when the `ScalingPolicy` parameter of the scaling group is set to `recycle`.
	//
	// - release: The ECS instances are released.
	//
	// The `ScalingPolicy` parameter of the `CreateScalingGroup` operation specifies the reclamation mode of a scaling group. However, the `RemovePolicy` parameter of the `RemoveInstances` operation determines the action taken when an instance is removed. For example:
	//
	// - If `ScalingPolicy` is `recycle` and `RemovePolicy` is `recycle`, the ECS instances enter the economical mode.
	//
	// - If `ScalingPolicy` is `recycle` and `RemovePolicy` is `release`, the ECS instances are released.
	//
	// - If `ScalingPolicy` is `release` and `RemovePolicy` is `recycle`, the ECS instances are released.
	//
	// - If `ScalingPolicy` is `release` and `RemovePolicy` is `release`, the ECS instances are released.
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
	// The timeout period, in seconds, for an ECS instance to stop during a scale-in. Valid values: 30 to 240.
	//
	// > - By default, this parameter inherits its value from the scaling group, but you can override it when calling the `RemoveInstances` operation.
	//
	// >
	//
	// > - This parameter takes effect only during scale-in events where `RemovePolicy` is set to `release`.
	//
	// >
	//
	// > - If this parameter is set, the system waits for the specified duration for the instance to stop. The scale-in process continues after the timeout expires, regardless of the instance\\"s state.
	//
	// >
	//
	// > - If this parameter is not set, the system waits until the instance stops before continuing the scale-in process. If the instance fails to stop, the scale-in operation is rolled back and fails.
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
