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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can only contain ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to modify the expected number of instances. Valid values:
	//
	// - true: After the ECS instances are removed from the scaling group, the expected number of instances is decreased by the number of removed instances.
	//
	// - false: After the ECS instances are removed from the scaling group, the expected number of instances remains unchanged.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DecreaseDesiredCapacity *bool `json:"DecreaseDesiredCapacity,omitempty" xml:"DecreaseDesiredCapacity,omitempty"`
	// Specifies whether to ignore invalid instances when removing a batch of instances from the scaling group. Valid values:
	//
	// - true: Invalid instances are ignored when removing a batch of instances from the scaling group. If invalid instances exist and valid instances are successfully removed, the execution status of the scaling activity is displayed as Warning. You can view the invalid instances in the scaling activity details.
	//
	// - false: Invalid instances are not ignored when removing a batch of instances from the scaling group. If invalid instances exist in the batch, an error is reported.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IgnoreInvalidInstance *bool `json:"IgnoreInvalidInstance,omitempty" xml:"IgnoreInvalidInstance,omitempty"`
	// The IDs of the ECS instances to be removed.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The context information of the lifecycle hook.
	LifecycleHookContext *RemoveInstancesRequestLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	OwnerAccount         *string                                     `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the action for the removed ECS instances. Valid values:
	//
	// - recycle: The ECS instances are stopped in economical mode.
	//
	//     > This value takes effect only when `ScalingPolicy` is set to `recycle`.
	//
	// - release: The ECS instances are released.
	//
	// The ScalingPolicy parameter of CreateScalingGroup specifies the reclaim mode of the scaling group, but the specific action when an instance is removed is determined by the RemovePolicy parameter of RemoveInstances. Examples:
	//
	// - If ScalingPolicy is set to recycle and RemovePolicy is set to recycle, the ECS instances are stopped in economical mode.
	//
	// - If ScalingPolicy is set to recycle and RemovePolicy is set to release, the ECS instances are released.
	//
	// - If ScalingPolicy is set to release and RemovePolicy is set to recycle, the ECS instances are released.
	//
	// - If ScalingPolicy is set to release and RemovePolicy is set to release, the ECS instances are released.
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
	// The timeout period for waiting for the ECS instances to be successfully stopped during the scale-in process. Unit: seconds.
	//
	// Valid values: 30 to 240.
	//
	// > - By default, this parameter inherits the value from the scaling group. You can also reset it when calling RemoveInstances.
	//
	// > - This parameter takes effect only in the scale-in process when RemovePolicy is set to release.
	//
	// > - If this parameter is set, the system waits for a maximum of the StopInstanceTimeout period for the instances to be successfully stopped. After the timeout period, the scale-in process continues regardless of whether the instances are successfully stopped.
	//
	// > - If this parameter is not set, the system waits for a long time for the instances to be successfully stopped. The scale-in process continues only after the instances are successfully stopped. If the instances fail to be stopped, the scale-in process is rolled back and the scale-in fails.
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
	// - true: Disables all lifecycle hooks.
	//
	// - false: Does not disable lifecycle hooks.
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// The IDs of the lifecycle hooks disabled for the scaling activity.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
	// The default execution result of the lifecycle hook in this scaling activity. After this parameter is set, it overwrites the original default execution result of the hook.
	//
	// example:
	//
	// CONTINUE
	LifecycleHookResult *string `json:"LifecycleHookResult,omitempty" xml:"LifecycleHookResult,omitempty"`
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
