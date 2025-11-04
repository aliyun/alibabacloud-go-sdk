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
	LifecycleHookContext *RemoveInstancesRequestLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	OwnerAccount         *string                                     `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// Specifies whether to disable the lifecycle hook. Valid Values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// The IDs of the lifecycle hooks that you want to disable.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
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

func (s *RemoveInstancesRequestLifecycleHookContext) SetDisableLifecycleHook(v bool) *RemoveInstancesRequestLifecycleHookContext {
	s.DisableLifecycleHook = &v
	return s
}

func (s *RemoveInstancesRequestLifecycleHookContext) SetIgnoredLifecycleHookIds(v []*string) *RemoveInstancesRequestLifecycleHookContext {
	s.IgnoredLifecycleHookIds = v
	return s
}

func (s *RemoveInstancesRequestLifecycleHookContext) Validate() error {
	return dara.Validate(s)
}
