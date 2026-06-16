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
	// The metadata for the scaling activity.
	//
	// example:
	//
	// {"key":"value"}
	ActivityMetadata *string `json:"ActivityMetadata,omitempty" xml:"ActivityMetadata,omitempty"`
	// The method used to adjust the number of instances in a scaling activity. Valid values:
	//
	// - `QuantityChangeInCapacity`: Adds or removes a specified number of ECS instances.
	//
	// - `PercentChangeInCapacity`: Adds or removes a specified percentage of ECS instances.
	//
	// - `TotalCapacity`: Adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// This parameter is required.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value for the scaling activity. A single adjustment cannot add or remove more than 1,000 ECS instances. The valid range depends on `AdjustmentType`:
	//
	// - `QuantityChangeInCapacity`: -1000 to 1000.
	//
	// - `PercentChangeInCapacity`: -100 to 10000.
	//
	// - `TotalCapacity`: 0 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// A client-generated token to ensure the idempotence of the request. This token must be a unique string of up to 64 ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The execution mode. Valid values:
	//
	// - `None`: Executes a standard scaling activity.
	//
	// - `PlanOnly`: Only performs elastic planning and returns the results in `PlanResult` without triggering the scaling activity. The results include details such as instance types, availability zones, billing methods, and the number of new instances.
	//
	// Default value: None.
	//
	// example:
	//
	// PlanOnly
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The lifecycle hook context.
	LifecycleHookContextShrink *string `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty"`
	// The minimum number of instances to adjust in a scaling activity. This parameter takes effect only when `AdjustmentType` is set to `PercentChangeInCapacity`.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The parameters to override when scaling out an ECI scaling group.
	OverridesShrink *string `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether the current scaling activity supports concurrency.
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
	// Specifies whether to execute the scaling activity synchronously. This parameter applies only to scaling groups that are configured with an expected number of instances. Valid values:
	//
	// - `true`: Synchronous execution. The scaling activity is triggered immediately.
	//
	// - `false`: Asynchronous execution. The call updates the expected number of instances without immediately triggering the scaling activity. The activity occurs when the system detects a discrepancy between the new expected number and the current number of instances.
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
