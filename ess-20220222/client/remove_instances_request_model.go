// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveInstancesRequest
	GetClientToken() *string
	SetDecreaseDesiredCapacity(v bool) *RemoveInstancesRequest
	GetDecreaseDesiredCapacity() *bool
	SetIgnoreInvalidInstance(v bool) *RemoveInstancesRequest
	GetIgnoreInvalidInstance() *bool
	SetInstanceIds(v []*string) *RemoveInstancesRequest
	GetInstanceIds() []*string
	SetLifecycleHookContext(v *RemoveInstancesRequestLifecycleHookContext) *RemoveInstancesRequest
	GetLifecycleHookContext() *RemoveInstancesRequestLifecycleHookContext
	SetOwnerAccount(v string) *RemoveInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveInstancesRequest
	GetRegionId() *string
	SetRemovePolicy(v string) *RemoveInstancesRequest
	GetRemovePolicy() *string
	SetResourceOwnerAccount(v string) *RemoveInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveInstancesRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *RemoveInstancesRequest
	GetScalingGroupId() *string
	SetStopInstanceTimeout(v int32) *RemoveInstancesRequest
	GetStopInstanceTimeout() *int32
}

type RemoveInstancesRequest struct {
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
	LifecycleHookContext *RemoveInstancesRequestLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	OwnerAccount         *string                                     `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s RemoveInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveInstancesRequest) GetDecreaseDesiredCapacity() *bool {
	return s.DecreaseDesiredCapacity
}

func (s *RemoveInstancesRequest) GetIgnoreInvalidInstance() *bool {
	return s.IgnoreInvalidInstance
}

func (s *RemoveInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveInstancesRequest) GetLifecycleHookContext() *RemoveInstancesRequestLifecycleHookContext {
	return s.LifecycleHookContext
}

func (s *RemoveInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveInstancesRequest) GetRemovePolicy() *string {
	return s.RemovePolicy
}

func (s *RemoveInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *RemoveInstancesRequest) GetStopInstanceTimeout() *int32 {
	return s.StopInstanceTimeout
}

func (s *RemoveInstancesRequest) SetClientToken(v string) *RemoveInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveInstancesRequest) SetDecreaseDesiredCapacity(v bool) *RemoveInstancesRequest {
	s.DecreaseDesiredCapacity = &v
	return s
}

func (s *RemoveInstancesRequest) SetIgnoreInvalidInstance(v bool) *RemoveInstancesRequest {
	s.IgnoreInvalidInstance = &v
	return s
}

func (s *RemoveInstancesRequest) SetInstanceIds(v []*string) *RemoveInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveInstancesRequest) SetLifecycleHookContext(v *RemoveInstancesRequestLifecycleHookContext) *RemoveInstancesRequest {
	s.LifecycleHookContext = v
	return s
}

func (s *RemoveInstancesRequest) SetOwnerAccount(v string) *RemoveInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetOwnerId(v int64) *RemoveInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveInstancesRequest) SetRegionId(v string) *RemoveInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveInstancesRequest) SetRemovePolicy(v string) *RemoveInstancesRequest {
	s.RemovePolicy = &v
	return s
}

func (s *RemoveInstancesRequest) SetResourceOwnerAccount(v string) *RemoveInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetResourceOwnerId(v int64) *RemoveInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveInstancesRequest) SetScalingGroupId(v string) *RemoveInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *RemoveInstancesRequest) SetStopInstanceTimeout(v int32) *RemoveInstancesRequest {
	s.StopInstanceTimeout = &v
	return s
}

func (s *RemoveInstancesRequest) Validate() error {
	if s.LifecycleHookContext != nil {
		if err := s.LifecycleHookContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveInstancesRequestLifecycleHookContext struct {
	// Specifies whether to disable all lifecycle hooks for the scaling activity. Valid values:
	//
	// - `true`: Disables all lifecycle hooks.
	//
	// - `false`: Does not disable lifecycle hooks.
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// A list of lifecycle hook IDs to ignore for the scaling activity.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
	LifecycleHookResult     *string   `json:"LifecycleHookResult,omitempty" xml:"LifecycleHookResult,omitempty"`
}

func (s RemoveInstancesRequestLifecycleHookContext) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstancesRequestLifecycleHookContext) GoString() string {
	return s.String()
}

func (s *RemoveInstancesRequestLifecycleHookContext) GetDisableLifecycleHook() *bool {
	return s.DisableLifecycleHook
}

func (s *RemoveInstancesRequestLifecycleHookContext) GetIgnoredLifecycleHookIds() []*string {
	return s.IgnoredLifecycleHookIds
}

func (s *RemoveInstancesRequestLifecycleHookContext) GetLifecycleHookResult() *string {
	return s.LifecycleHookResult
}

func (s *RemoveInstancesRequestLifecycleHookContext) SetDisableLifecycleHook(v bool) *RemoveInstancesRequestLifecycleHookContext {
	s.DisableLifecycleHook = &v
	return s
}

func (s *RemoveInstancesRequestLifecycleHookContext) SetIgnoredLifecycleHookIds(v []*string) *RemoveInstancesRequestLifecycleHookContext {
	s.IgnoredLifecycleHookIds = v
	return s
}

func (s *RemoveInstancesRequestLifecycleHookContext) SetLifecycleHookResult(v string) *RemoveInstancesRequestLifecycleHookContext {
	s.LifecycleHookResult = &v
	return s
}

func (s *RemoveInstancesRequestLifecycleHookContext) Validate() error {
	return dara.Validate(s)
}
