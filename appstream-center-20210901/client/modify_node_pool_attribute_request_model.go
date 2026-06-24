// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ModifyNodePoolAttributeRequest
	GetBizRegionId() *string
	SetNodeCapacity(v int32) *ModifyNodePoolAttributeRequest
	GetNodeCapacity() *int32
	SetNodePoolStrategy(v *ModifyNodePoolAttributeRequestNodePoolStrategy) *ModifyNodePoolAttributeRequest
	GetNodePoolStrategy() *ModifyNodePoolAttributeRequestNodePoolStrategy
	SetPoolId(v string) *ModifyNodePoolAttributeRequest
	GetPoolId() *string
	SetProductType(v string) *ModifyNodePoolAttributeRequest
	GetProductType() *string
}

type ModifyNodePoolAttributeRequest struct {
	// The region ID of the delivery group. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The number of concurrent sessions, which is the number of sessions that can be simultaneously connected to a single resource. If too many sessions are connected simultaneously, the application experience may degrade. The valid values vary depending on the resource specification. The valid values for each resource specification are as follows:
	//
	// - appstreaming.general.4c8g: 1 to 2.
	//
	// - appstreaming.general.8c16g: 1 to 4.
	//
	// - appstreaming.vgpu.8c16g.4g: 1 to 4.
	//
	// - appstreaming.vgpu.8c31g.16g: 1 to 4.
	//
	// - appstreaming.vgpu.14c93g.12g: 1 to 6.
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The automatic scaling policy of the delivery group.
	NodePoolStrategy *ModifyNodePoolAttributeRequestNodePoolStrategy `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty" type:"Struct"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ew7va2g1wl3vm****
	PoolId *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// The product type.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ModifyNodePoolAttributeRequest) GetNodeCapacity() *int32 {
	return s.NodeCapacity
}

func (s *ModifyNodePoolAttributeRequest) GetNodePoolStrategy() *ModifyNodePoolAttributeRequestNodePoolStrategy {
	return s.NodePoolStrategy
}

func (s *ModifyNodePoolAttributeRequest) GetPoolId() *string {
	return s.PoolId
}

func (s *ModifyNodePoolAttributeRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyNodePoolAttributeRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodePoolStrategy(v *ModifyNodePoolAttributeRequestNodePoolStrategy) *ModifyNodePoolAttributeRequest {
	s.NodePoolStrategy = v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetPoolId(v string) *ModifyNodePoolAttributeRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetProductType(v string) *ModifyNodePoolAttributeRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) Validate() error {
	if s.NodePoolStrategy != nil {
		if err := s.NodePoolStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNodePoolAttributeRequestNodePoolStrategy struct {
	// The maximum number of idle sessions. When this value is specified, automatic scale-out is triggered only when the session usage exceeds `ScalingUsageThreshold` and the number of idle sessions in the current delivery group is less than `MaxIdleAppInstanceAmount`. Otherwise, the idle sessions in the delivery group are considered sufficient, and no automatic scale-out is performed. This parameter can be used to flexibly control elastic scale-out behavior and reduce costs.
	//
	// example:
	//
	// 3
	MaxIdleAppInstanceAmount *int32 `json:"MaxIdleAppInstanceAmount,omitempty" xml:"MaxIdleAppInstanceAmount,omitempty"`
	// The maximum number of resources that can be created during scale-out. This parameter is required when `StrategyType` is set to `NODE_SCALING_BY_USAGE`.
	//
	// example:
	//
	// 10
	MaxScalingAmount *int32 `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	// The number of purchased resources. Valid values: 1 to 100.
	//
	// >
	//
	// - If the resources are subscription resources, this parameter cannot be modified.
	//
	// - If the resources are pay-as-you-go resources, this parameter can be modified when the scaling mode (`StrategyType`) is set to fixed quantity (`NODE_FIXED`) or automatic scaling (`NODE_SCALING_BY_USAGE`).
	//
	// example:
	//
	// 1
	NodeAmount *int32 `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	// The list of policy execution cycles. This parameter is required when `StrategyType` (scaling mode) is set to `NODE_SCALING_BY_SCHEDULE` (scheduled scaling).
	RecurrenceSchedules []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	// The maximum duration (in minutes) that a resource without session connections is retained. When no sessions are connected to a resource, a countdown starts based on the duration specified here. The resource is scaled in when the countdown ends. Valid values: 5 to 120. Default value: 5. The following exceptions apply:
	//
	// - If scale-in would trigger automatic scale-out again, the scale-in is not performed to avoid repeated scale-in and scale-out operations.
	//
	// - If automatic scale-out is triggered by an increase in sessions during this period, the resource is not scaled in as originally planned, and the countdown restarts.
	//
	// example:
	//
	// 5
	ScalingDownAfterIdleMinutes *int32 `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	// The number of resources created per scale-out operation. Valid values: 1 to 10. This parameter is required when `StrategyType` is set to `NODE_SCALING_BY_USAGE`.
	//
	// example:
	//
	// 2
	ScalingStep *int32 `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	// The upper threshold of session usage (%). Automatic scale-out is triggered when the session usage exceeds this threshold. The session usage is calculated by using the following formula: `Session usage = Current sessions ÷ (Total resources × Concurrent sessions per resource) × 100%`. This parameter is required when `StrategyType` is set to `NODE_SCALING_BY_USAGE`. Valid values: 0 to 100. Default value: 85.
	//
	// example:
	//
	// 85
	ScalingUsageThreshold *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	// The date when the policy expires. Format: yyyy-MM-dd. The interval between the expiration date and the effective date must be between 7 days and 1 year, inclusive. This parameter is required when `StrategyType` (scaling mode) is set to `NODE_SCALING_BY_SCHEDULE` (scheduled scaling).
	//
	// example:
	//
	// 2023-01-19
	StrategyDisableDate *string `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	// The date when the policy takes effect. Format: yyyy-MM-dd. The date must be equal to or later than the current date. This parameter is required when `StrategyType` (scaling mode) is set to `NODE_SCALING_BY_SCHEDULE` (scheduled scaling).
	//
	// example:
	//
	// 2023-01-05
	StrategyEnableDate *string `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	// The scaling mode.
	//
	// >
	//
	// - `NODE_FIXED` (fixed quantity): Applicable to subscription and pay-as-you-go resources.
	//
	// - `NODE_SCALING_BY_USAGE` (automatic scaling): Applicable to subscription and pay-as-you-go resources.
	//
	// - `NODE_SCALING_BY_SCHEDULE` (scheduled scaling): Applicable only to pay-as-you-go resources.
	//
	// example:
	//
	// NODE_FIXED
	StrategyType *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// Specifies whether to enable the resource prefetch policy. This parameter is required when `StrategyType` (scaling mode) is set to `NODE_SCALING_BY_SCHEDULE` (scheduled scaling).
	//
	// example:
	//
	// false
	WarmUp *bool `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetMaxIdleAppInstanceAmount() *int32 {
	return s.MaxIdleAppInstanceAmount
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetMaxScalingAmount() *int32 {
	return s.MaxScalingAmount
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetNodeAmount() *int32 {
	return s.NodeAmount
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetRecurrenceSchedules() []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	return s.RecurrenceSchedules
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetScalingDownAfterIdleMinutes() *int32 {
	return s.ScalingDownAfterIdleMinutes
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetScalingStep() *int32 {
	return s.ScalingStep
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetScalingUsageThreshold() *string {
	return s.ScalingUsageThreshold
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetStrategyDisableDate() *string {
	return s.StrategyDisableDate
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetStrategyEnableDate() *string {
	return s.StrategyEnableDate
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetStrategyType() *string {
	return s.StrategyType
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) GetWarmUp() *bool {
	return s.WarmUp
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetMaxIdleAppInstanceAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.MaxIdleAppInstanceAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetMaxScalingAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.MaxScalingAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetNodeAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.NodeAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetRecurrenceSchedules(v []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.RecurrenceSchedules = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingDownAfterIdleMinutes(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingStep(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingStep = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingUsageThreshold(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyDisableDate(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyDisableDate = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyEnableDate(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyEnableDate = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetWarmUp(v bool) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.WarmUp = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) Validate() error {
	if s.RecurrenceSchedules != nil {
		for _, item := range s.RecurrenceSchedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules struct {
	// The type of the policy execution cycle. You must specify both `RecurrenceType` and `RecurrenceValues`.
	//
	// example:
	//
	// weekly
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The list of values for the policy execution cycle.
	RecurrenceValues []*int32 `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	// The list of time periods for the policy execution cycle. Requirements for time period settings:
	//
	// - You can add up to three time periods.
	//
	// - Time periods must not overlap.
	//
	// - The interval between time periods must be at least 5 minutes.
	//
	// - Each time period must be at least 15 minutes long.
	//
	// - All time periods combined must not span across days.
	TimerPeriods []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) GetRecurrenceValues() []*int32 {
	return s.RecurrenceValues
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) GetTimerPeriods() []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	return s.TimerPeriods
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetRecurrenceType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) SetTimerPeriods(v []*ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedules) Validate() error {
	if s.TimerPeriods != nil {
		for _, item := range s.TimerPeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods struct {
	// The resource count.
	//
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The end time. Format: HH:mm.
	//
	// example:
	//
	// 15:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time. Format: HH:mm.
	//
	// example:
	//
	// 12:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) GetAmount() *int32 {
	return s.Amount
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategyRecurrenceSchedulesTimerPeriods) Validate() error {
	return dara.Validate(s)
}
