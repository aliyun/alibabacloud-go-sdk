// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleWithAdjustmentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityMetadata(v string) *ScaleWithAdjustmentShrinkRequest
	GetActivityMetadata() *string
	SetAdjustmentType(v string) *ScaleWithAdjustmentShrinkRequest
	GetAdjustmentType() *string
	SetAdjustmentValue(v int32) *ScaleWithAdjustmentShrinkRequest
	GetAdjustmentValue() *int32
	SetClientToken(v string) *ScaleWithAdjustmentShrinkRequest
	GetClientToken() *string
	SetExecutionMode(v string) *ScaleWithAdjustmentShrinkRequest
	GetExecutionMode() *string
	SetLifecycleHookContextShrink(v string) *ScaleWithAdjustmentShrinkRequest
	GetLifecycleHookContextShrink() *string
	SetMinAdjustmentMagnitude(v int32) *ScaleWithAdjustmentShrinkRequest
	GetMinAdjustmentMagnitude() *int32
	SetOverridesShrink(v string) *ScaleWithAdjustmentShrinkRequest
	GetOverridesShrink() *string
	SetOwnerId(v int64) *ScaleWithAdjustmentShrinkRequest
	GetOwnerId() *int64
	SetParallelTask(v bool) *ScaleWithAdjustmentShrinkRequest
	GetParallelTask() *bool
	SetResourceOwnerAccount(v string) *ScaleWithAdjustmentShrinkRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ScaleWithAdjustmentShrinkRequest
	GetScalingGroupId() *string
	SetSyncActivity(v bool) *ScaleWithAdjustmentShrinkRequest
	GetSyncActivity() *bool
}

type ScaleWithAdjustmentShrinkRequest struct {
	// The metadata of the scaling activity.
	//
	// example:
	//
	// {"key":"value"}
	ActivityMetadata *string `json:"ActivityMetadata,omitempty" xml:"ActivityMetadata,omitempty"`
	// The adjustment method for the scaling activity. Valid values:
	//
	// - QuantityChangeInCapacity: Increases or decreases a specified number of ECS instances.
	//
	// - PercentChangeInCapacity: Increases or decreases a specified percentage of ECS instances.
	//
	// - TotalCapacity: Adjusts the number of ECS instances in the current scaling group to a specified number.
	//
	// This parameter is required.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value for the scaling activity. In all cases, the number of ECS instances adjusted in a single operation cannot exceed 1,000. Otherwise, the adjustment fails. Valid values vary based on the adjustment method:
	//
	// - QuantityChangeInCapacity: -1,000 to 1,000.
	//
	// - PercentChangeInCapacity: -100 to 10,000.
	//
	// - TotalCapacity: 0 to 2,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and can be up to 64 ASCII characters in length. The value cannot contain non-ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The execution mode. Valid values:
	//
	// - None: Not specified. Normal scaling is performed.
	//
	// - PlanOnly: Does not trigger scaling. Only elastic planning is executed. The planning results, including instance types, zone IDs, billing methods, and the number of instances to be created, are returned in PlanResult.
	//
	// Default value: None.
	//
	// example:
	//
	// PlanOnly
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The context information of the lifecycle hook.
	LifecycleHookContextShrink *string `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty"`
	// The minimum number of instances to be adjusted in the scaling activity. This parameter takes effect only when `AdjustmentType` is set to `PercentChangeInCapacity`.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The override parameters specified for scaling out an Elastic Container Instance (ECI) scaling group.
	OverridesShrink *string `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether the current scaling task supports concurrent execution.
	//
	// example:
	//
	// false
	ParallelTask         *bool   `json:"ParallelTask,omitempty" xml:"ParallelTask,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-j6c1o397427hyjdc****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// Specifies whether to synchronously execute the scaling activity. This parameter is valid only for scaling groups that have the expected number of instances configured. Valid values:
	//
	// - true: Synchronous process. When the scaling rule is executed, the scale-out or scale-in activity of the scaling group is directly triggered.
	//
	// - false: Asynchronous process. When the expected number of instances in the scaling group is modified, the actual scale-out or scale-in activity of the scaling group is not immediately triggered. The system automatically detects the difference between the expected number of instances and the current number of instances in the scaling group, and then triggers the scale-out or scale-in activity.
	//
	// > For more information about the expected number of instances, see [Expected number of instances](https://help.aliyun.com/document_detail/146231.html).
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SyncActivity *bool `json:"SyncActivity,omitempty" xml:"SyncActivity,omitempty"`
}

func (s ScaleWithAdjustmentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleWithAdjustmentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScaleWithAdjustmentShrinkRequest) GetActivityMetadata() *string {
	return s.ActivityMetadata
}

func (s *ScaleWithAdjustmentShrinkRequest) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *ScaleWithAdjustmentShrinkRequest) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *ScaleWithAdjustmentShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ScaleWithAdjustmentShrinkRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *ScaleWithAdjustmentShrinkRequest) GetLifecycleHookContextShrink() *string {
	return s.LifecycleHookContextShrink
}

func (s *ScaleWithAdjustmentShrinkRequest) GetMinAdjustmentMagnitude() *int32 {
	return s.MinAdjustmentMagnitude
}

func (s *ScaleWithAdjustmentShrinkRequest) GetOverridesShrink() *string {
	return s.OverridesShrink
}

func (s *ScaleWithAdjustmentShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ScaleWithAdjustmentShrinkRequest) GetParallelTask() *bool {
	return s.ParallelTask
}

func (s *ScaleWithAdjustmentShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ScaleWithAdjustmentShrinkRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ScaleWithAdjustmentShrinkRequest) GetSyncActivity() *bool {
	return s.SyncActivity
}

func (s *ScaleWithAdjustmentShrinkRequest) SetActivityMetadata(v string) *ScaleWithAdjustmentShrinkRequest {
	s.ActivityMetadata = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetAdjustmentType(v string) *ScaleWithAdjustmentShrinkRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetAdjustmentValue(v int32) *ScaleWithAdjustmentShrinkRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetClientToken(v string) *ScaleWithAdjustmentShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetExecutionMode(v string) *ScaleWithAdjustmentShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetLifecycleHookContextShrink(v string) *ScaleWithAdjustmentShrinkRequest {
	s.LifecycleHookContextShrink = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetMinAdjustmentMagnitude(v int32) *ScaleWithAdjustmentShrinkRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetOverridesShrink(v string) *ScaleWithAdjustmentShrinkRequest {
	s.OverridesShrink = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetOwnerId(v int64) *ScaleWithAdjustmentShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetParallelTask(v bool) *ScaleWithAdjustmentShrinkRequest {
	s.ParallelTask = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetResourceOwnerAccount(v string) *ScaleWithAdjustmentShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetScalingGroupId(v string) *ScaleWithAdjustmentShrinkRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) SetSyncActivity(v bool) *ScaleWithAdjustmentShrinkRequest {
	s.SyncActivity = &v
	return s
}

func (s *ScaleWithAdjustmentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
