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
	// The type of the scaling policy. Valid values:
	//
	// 	- QuantityChangeInCapacity: adds the specified number of ECS instances to or removes the specified number of ECS instances from the scaling group.
	//
	// 	- PercentChangeInCapacity: adds the specified percentage of ECS instances to or removes the specified percentage of ECS instances from the scaling group.
	//
	// 	- TotalCapacity: adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// This parameter is required.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The number of instances in each adjustment. The number of ECS instances in each adjustment cannot exceed 1,000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to QuantityChangeInCapacity: -1000 to 1000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to PercentChangeInCapacity: -100 to 10000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to TotalCapacity: 0 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- None: If this is not specified, auto scaling is performed.
	//
	// 	- PlanOnly: Scaling is not triggered. Only elastic planning is performed. The planning result is returned in PlanResult, including the instance type, zone ID, billing type, and number of created instances.
	//
	// Default value: None.
	//
	// example:
	//
	// PlanOnly
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The context of the lifecycle hook.
	LifecycleHookContextShrink *string `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty"`
	// The minimum number of instances allowed in each adjustment. This parameter takes effect only if you set the `AdjustmentType` parameter to `PercentChangeInCapacity`.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The overrides that allow you to adjust the scaling group of the Elastic Container Instance (ECI) type during a scale-out event.
	OverridesShrink *string `json:"Overrides,omitempty" xml:"Overrides,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Whether the current scale-out task supports concurrency.
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
	// Specifies whether to trigger the scaling task in a synchronous manner. This parameter takes effect only on scaling groups for which you specified an expected number of instances. Valid Values:
	//
	// 	- true: triggers the scaling task in a synchronous manner. A scaling activity is triggered at the time when the scaling rule is executed.
	//
	// 	- false: does not trigger the scaling task in a synchronous manner. After you change the expected number of instances for the scaling group, Auto Scaling checks whether the total number of instances in the scaling group matches the new expected number and determines whether to trigger the scaling activity based on the check result.
	//
	// >  For more information, see [Expected number of instances](https://help.aliyun.com/document_detail/146231.html).
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
