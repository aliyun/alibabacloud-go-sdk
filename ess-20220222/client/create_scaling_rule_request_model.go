// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustmentType(v string) *CreateScalingRuleRequest
	GetAdjustmentType() *string
	SetAdjustmentValue(v int32) *CreateScalingRuleRequest
	GetAdjustmentValue() *int32
	SetAlarmDimensions(v []*CreateScalingRuleRequestAlarmDimensions) *CreateScalingRuleRequest
	GetAlarmDimensions() []*CreateScalingRuleRequestAlarmDimensions
	SetAlarmOptions(v *CreateScalingRuleRequestAlarmOptions) *CreateScalingRuleRequest
	GetAlarmOptions() *CreateScalingRuleRequestAlarmOptions
	SetCooldown(v int32) *CreateScalingRuleRequest
	GetCooldown() *int32
	SetDisableScaleIn(v bool) *CreateScalingRuleRequest
	GetDisableScaleIn() *bool
	SetEstimatedInstanceWarmup(v int32) *CreateScalingRuleRequest
	GetEstimatedInstanceWarmup() *int32
	SetHybridMetrics(v []*CreateScalingRuleRequestHybridMetrics) *CreateScalingRuleRequest
	GetHybridMetrics() []*CreateScalingRuleRequestHybridMetrics
	SetHybridMonitorNamespace(v string) *CreateScalingRuleRequest
	GetHybridMonitorNamespace() *string
	SetInitialMaxSize(v int32) *CreateScalingRuleRequest
	GetInitialMaxSize() *int32
	SetMetricName(v string) *CreateScalingRuleRequest
	GetMetricName() *string
	SetMetricType(v string) *CreateScalingRuleRequest
	GetMetricType() *string
	SetMinAdjustmentMagnitude(v int32) *CreateScalingRuleRequest
	GetMinAdjustmentMagnitude() *int32
	SetOwnerAccount(v string) *CreateScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateScalingRuleRequest
	GetOwnerId() *int64
	SetPredictiveScalingMode(v string) *CreateScalingRuleRequest
	GetPredictiveScalingMode() *string
	SetPredictiveTaskBufferTime(v int32) *CreateScalingRuleRequest
	GetPredictiveTaskBufferTime() *int32
	SetPredictiveValueBehavior(v string) *CreateScalingRuleRequest
	GetPredictiveValueBehavior() *string
	SetPredictiveValueBuffer(v int32) *CreateScalingRuleRequest
	GetPredictiveValueBuffer() *int32
	SetRegionId(v string) *CreateScalingRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetScaleInEvaluationCount(v int32) *CreateScalingRuleRequest
	GetScaleInEvaluationCount() *int32
	SetScaleOutEvaluationCount(v int32) *CreateScalingRuleRequest
	GetScaleOutEvaluationCount() *int32
	SetScalingGroupId(v string) *CreateScalingRuleRequest
	GetScalingGroupId() *string
	SetScalingRuleName(v string) *CreateScalingRuleRequest
	GetScalingRuleName() *string
	SetScalingRuleType(v string) *CreateScalingRuleRequest
	GetScalingRuleType() *string
	SetStepAdjustments(v []*CreateScalingRuleRequestStepAdjustments) *CreateScalingRuleRequest
	GetStepAdjustments() []*CreateScalingRuleRequestStepAdjustments
	SetTargetValue(v float32) *CreateScalingRuleRequest
	GetTargetValue() *float32
}

type CreateScalingRuleRequest struct {
	// The adjustment method of the scaling rule. This parameter is required for simple scaling rules and step scaling rules. Valid values:
	//
	// - QuantityChangeInCapacity: adds or removes a specified number of ECS instances.
	//
	// - PercentChangeInCapacity: adds or removes a specified percentage of ECS instances.
	//
	// - TotalCapacity: adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value of the scaling rule. This parameter is required for simple scaling rules and step scaling rules. The number of ECS instances adjusted in a single operation cannot exceed 1,000. Valid values vary based on the adjustment method:
	//
	// - QuantityChangeInCapacity: -1000 to 1000.
	//
	// - PercentChangeInCapacity: -100 to 10000.
	//
	// - TotalCapacity: 0 to 2000.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The dimension information of the metric. This parameter is applicable to target tracking scaling rules and is configured when the metric requires additional dimension information. For example, the LoadBalancerRealServerAverageQps metric requires the dimension key-value information of rulePool.
	AlarmDimensions []*CreateScalingRuleRequestAlarmDimensions `json:"AlarmDimensions,omitempty" xml:"AlarmDimensions,omitempty" type:"Repeated"`
	// The definition of alert attributes.
	AlarmOptions *CreateScalingRuleRequestAlarmOptions `json:"AlarmOptions,omitempty" xml:"AlarmOptions,omitempty" type:"Struct"`
	// The cooldown time of the scaling rule. This parameter is applicable only to simple scaling rules. Valid values: 0 to 86400. Unit: seconds.
	//
	// Default value: none.
	//
	// example:
	//
	// 60
	Cooldown *int32 `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	// Specifies whether to disable scale-in. This parameter is applicable only to target tracking scaling rules.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DisableScaleIn *bool `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	// The warmup time of instances. This parameter is applicable to target tracking scaling rules and step scaling rules. ECS instances in the warmup state are added to the scaling group normally, but they do not report monitoring data to CloudMonitor during the warmup period.
	//
	// > When dynamically calculating the number of ECS instances to be scaled, instances in the warmup state are not counted in the current number of instances.
	//
	// Valid values: 0 to 86400. Unit: seconds.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	EstimatedInstanceWarmup *int32 `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	// The metric settings for Hybrid Cloud Monitoring. For more information about how to configure the settings, see [Customize target tracking scaling rules by using formula operations](https://help.aliyun.com/document_detail/2852162.html).
	HybridMetrics []*CreateScalingRuleRequestHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
	// The ID of the monitoring metrics repository for Hybrid Cloud Monitoring.
	//
	// For more information about managing monitoring metrics repositories, see [Manage monitoring metrics repositories](https://help.aliyun.com/document_detail/217606.html).
	//
	// example:
	//
	// aliyun-test
	HybridMonitorNamespace *string `json:"HybridMonitorNamespace,omitempty" xml:"HybridMonitorNamespace,omitempty"`
	// The maximum number of instances in the scaling group. This parameter is used in combination with `PredictiveValueBehavior`.
	//
	// Default value: the maximum number of instances in the scaling group (the value of MaxSize).
	//
	// example:
	//
	// 100
	InitialMaxSize *int32 `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	// The predefined metric. This parameter is required for target tracking scaling rules and predictive scaling rules.
	//
	// Valid values for target tracking scaling rules:
	//
	// - CpuUtilizationAgent: (Agent) CPU utilization (recommended).
	//
	// - MemoryUtilization: (Agent) memory usage (recommended).
	//
	// - CpuUtilization: (ECS) average CPU utilization.
	//
	// - IntranetTx: (ECS) average outbound traffic over the internal network.
	//
	// - IntranetRx: (ECS) average inbound traffic over the internal network.
	//
	// - VpcInternetTx: (ECS) average outbound traffic over the Internet.
	//
	// - VpcInternetRx: (ECS) average inbound traffic over the Internet.
	//
	// - LoadBalancerRealServerAverageQps: (ALB) average QPS per backend server in a server group.
	//
	// Valid values for predictive scaling rules:
	//
	// - CpuUtilization: (ECS) average CPU utilization.
	//
	// - IntranetRx: (ECS) average inbound traffic over the internal network.
	//
	// - IntranetTx: (ECS) average outbound traffic over the internal network.
	//
	// For more information, see [System monitoring event-triggered tasks](https://help.aliyun.com/document_detail/74854.html).
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the metric. Valid values:
	//
	// - system: uses system metrics of CloudMonitor.
	//
	// - custom: uses custom metrics reported to CloudMonitor.
	//
	// - hybrid: uses metrics of Hybrid Cloud Monitoring.
	//
	// example:
	//
	// system
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The minimum number of instances to be adjusted. This parameter takes effect only when the scaling rule type is SimpleScalingRule or StepScalingRule and AdjustmentType is set to PercentChangeInCapacity.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32  `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mode of the predictive scaling rule. Valid values:
	//
	// - PredictAndScale: generates prediction results and creates predictive tasks.
	//
	// - PredictOnly: generates prediction results but does not create predictive tasks.
	//
	// Default value: PredictAndScale.
	//
	// example:
	//
	// PredictAndScale
	PredictiveScalingMode *string `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	// The predictive tasks automatically created by predictive scaling rules are executed on the hour by default. You can set the pre-startup time to execute the predictive tasks in advance and prepare resources beforehand. Valid values: 0 to 60. Unit: minutes.
	//
	// Default value: 0.
	//
	// example:
	//
	// 30
	PredictiveTaskBufferTime *int32 `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	// The processing method for the maximum value of the predictive scaling rule. Valid values:
	//
	// - MaxOverridePredictiveValue: The initial maximum value overrides the predicted value. If the predicted value is greater than the initial maximum value, the maximum value of the predictive task uses the initial maximum value.
	//
	// - PredictiveValueOverrideMax: The predicted value overrides the initial maximum value. If the predicted value is greater than the initial maximum value, the maximum value of the predictive task uses the predicted value.
	//
	// - PredictiveValueOverrideMaxWithBuffer: A buffer ratio is added to the predicted value. The predicted value is increased by the ratio specified by PredictiveValueBuffer. If the increased value is greater than the initial maximum value, the increased value is used.
	//
	// Default value: MaxOverridePredictiveValue.
	//
	// example:
	//
	// MaxOverridePredictiveValue
	PredictiveValueBehavior *string `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	// This parameter takes effect only when `PredictiveValueBehavior` is set to `PredictiveValueOverrideMaxWithBuffer`. The predicted value is increased by this ratio. If the increased value is greater than the initial maximum value, the increased value is used. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 50
	PredictiveValueBuffer *int32 `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// After a target tracking scaling rule is created, an event-triggered task is automatically created. This parameter specifies the number of consecutive times the threshold condition must be met before the corresponding scale-in event-triggered task triggers an alert.
	//
	// Default value: 15.
	//
	// example:
	//
	// 15
	ScaleInEvaluationCount *int32 `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	// After a target tracking scaling rule is created, an event-triggered task is automatically created. This parameter specifies the number of consecutive times the threshold condition must be met before the corresponding scale-out event-triggered task triggers an alert.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	ScaleOutEvaluationCount *int32 `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	// The ID of the scaling group to which the scaling rule belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The name of the scaling rule. The name must be 2 to 64 characters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter, digit, or Chinese character.
	//
	// The name of a scaling rule must be unique within a scaling group in a region for an Alibaba Cloud account.
	//
	// If this parameter is not specified, the value of ScalingRuleId is used by default.
	//
	// example:
	//
	// scalingrule****
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// The type of the scaling rule. Valid values:
	//
	// - SimpleScalingRule: simple scaling rule. Adjusts the number of ECS or ECI instances based on the adjustment method (AdjustmentType) and adjustment value (AdjustmentValue).
	//
	// - TargetTrackingScalingRule: target tracking scaling rule. Dynamically calculates the number of ECS or ECI instances for scale-out or scale-in based on a predefined metric (MetricName) to keep the metric value close to the target value (TargetValue).
	//
	// - StepScalingRule: step scaling rule. Provides a step-by-step scaling method based on thresholds and metric values.
	//
	// - PredictiveScalingRule: predictive scaling rule. Analyzes historical monitoring data of the scaling group based on machine learning capabilities to predict future metric values, and supports the automatic creation of scheduled tasks to configure the boundary values of the scaling group.
	//
	// Default value: SimpleScalingRule.
	//
	// example:
	//
	// SimpleScalingRule
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	// The collection of instance scaling information corresponding to step adjustments.
	StepAdjustments []*CreateScalingRuleRequestStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	// The target value. This parameter is required for target tracking scaling rules and predictive scaling rules. The value of TargetValue can be rounded to three decimal places and must be greater than 0.
	//
	// example:
	//
	// 0.125
	TargetValue *float32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s CreateScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequest) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *CreateScalingRuleRequest) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *CreateScalingRuleRequest) GetAlarmDimensions() []*CreateScalingRuleRequestAlarmDimensions {
	return s.AlarmDimensions
}

func (s *CreateScalingRuleRequest) GetAlarmOptions() *CreateScalingRuleRequestAlarmOptions {
	return s.AlarmOptions
}

func (s *CreateScalingRuleRequest) GetCooldown() *int32 {
	return s.Cooldown
}

func (s *CreateScalingRuleRequest) GetDisableScaleIn() *bool {
	return s.DisableScaleIn
}

func (s *CreateScalingRuleRequest) GetEstimatedInstanceWarmup() *int32 {
	return s.EstimatedInstanceWarmup
}

func (s *CreateScalingRuleRequest) GetHybridMetrics() []*CreateScalingRuleRequestHybridMetrics {
	return s.HybridMetrics
}

func (s *CreateScalingRuleRequest) GetHybridMonitorNamespace() *string {
	return s.HybridMonitorNamespace
}

func (s *CreateScalingRuleRequest) GetInitialMaxSize() *int32 {
	return s.InitialMaxSize
}

func (s *CreateScalingRuleRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateScalingRuleRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *CreateScalingRuleRequest) GetMinAdjustmentMagnitude() *int32 {
	return s.MinAdjustmentMagnitude
}

func (s *CreateScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateScalingRuleRequest) GetPredictiveScalingMode() *string {
	return s.PredictiveScalingMode
}

func (s *CreateScalingRuleRequest) GetPredictiveTaskBufferTime() *int32 {
	return s.PredictiveTaskBufferTime
}

func (s *CreateScalingRuleRequest) GetPredictiveValueBehavior() *string {
	return s.PredictiveValueBehavior
}

func (s *CreateScalingRuleRequest) GetPredictiveValueBuffer() *int32 {
	return s.PredictiveValueBuffer
}

func (s *CreateScalingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateScalingRuleRequest) GetScaleInEvaluationCount() *int32 {
	return s.ScaleInEvaluationCount
}

func (s *CreateScalingRuleRequest) GetScaleOutEvaluationCount() *int32 {
	return s.ScaleOutEvaluationCount
}

func (s *CreateScalingRuleRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *CreateScalingRuleRequest) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *CreateScalingRuleRequest) GetStepAdjustments() []*CreateScalingRuleRequestStepAdjustments {
	return s.StepAdjustments
}

func (s *CreateScalingRuleRequest) GetTargetValue() *float32 {
	return s.TargetValue
}

func (s *CreateScalingRuleRequest) SetAdjustmentType(v string) *CreateScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetAdjustmentValue(v int32) *CreateScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *CreateScalingRuleRequest) SetAlarmDimensions(v []*CreateScalingRuleRequestAlarmDimensions) *CreateScalingRuleRequest {
	s.AlarmDimensions = v
	return s
}

func (s *CreateScalingRuleRequest) SetAlarmOptions(v *CreateScalingRuleRequestAlarmOptions) *CreateScalingRuleRequest {
	s.AlarmOptions = v
	return s
}

func (s *CreateScalingRuleRequest) SetCooldown(v int32) *CreateScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *CreateScalingRuleRequest) SetDisableScaleIn(v bool) *CreateScalingRuleRequest {
	s.DisableScaleIn = &v
	return s
}

func (s *CreateScalingRuleRequest) SetEstimatedInstanceWarmup(v int32) *CreateScalingRuleRequest {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *CreateScalingRuleRequest) SetHybridMetrics(v []*CreateScalingRuleRequestHybridMetrics) *CreateScalingRuleRequest {
	s.HybridMetrics = v
	return s
}

func (s *CreateScalingRuleRequest) SetHybridMonitorNamespace(v string) *CreateScalingRuleRequest {
	s.HybridMonitorNamespace = &v
	return s
}

func (s *CreateScalingRuleRequest) SetInitialMaxSize(v int32) *CreateScalingRuleRequest {
	s.InitialMaxSize = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMetricName(v string) *CreateScalingRuleRequest {
	s.MetricName = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMetricType(v string) *CreateScalingRuleRequest {
	s.MetricType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetMinAdjustmentMagnitude(v int32) *CreateScalingRuleRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *CreateScalingRuleRequest) SetOwnerAccount(v string) *CreateScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetOwnerId(v int64) *CreateScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveScalingMode(v string) *CreateScalingRuleRequest {
	s.PredictiveScalingMode = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveTaskBufferTime(v int32) *CreateScalingRuleRequest {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveValueBehavior(v string) *CreateScalingRuleRequest {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *CreateScalingRuleRequest) SetPredictiveValueBuffer(v int32) *CreateScalingRuleRequest {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *CreateScalingRuleRequest) SetRegionId(v string) *CreateScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetResourceOwnerAccount(v string) *CreateScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScaleInEvaluationCount(v int32) *CreateScalingRuleRequest {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScaleOutEvaluationCount(v int32) *CreateScalingRuleRequest {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingGroupId(v string) *CreateScalingRuleRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingRuleName(v string) *CreateScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingRuleType(v string) *CreateScalingRuleRequest {
	s.ScalingRuleType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetStepAdjustments(v []*CreateScalingRuleRequestStepAdjustments) *CreateScalingRuleRequest {
	s.StepAdjustments = v
	return s
}

func (s *CreateScalingRuleRequest) SetTargetValue(v float32) *CreateScalingRuleRequest {
	s.TargetValue = &v
	return s
}

func (s *CreateScalingRuleRequest) Validate() error {
	if s.AlarmDimensions != nil {
		for _, item := range s.AlarmDimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AlarmOptions != nil {
		if err := s.AlarmOptions.Validate(); err != nil {
			return err
		}
	}
	if s.HybridMetrics != nil {
		for _, item := range s.HybridMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StepAdjustments != nil {
		for _, item := range s.StepAdjustments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScalingRuleRequestAlarmDimensions struct {
	// The dimension key associated with the metric.
	//
	// example:
	//
	// rulePool
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension value associated with the metric.
	//
	// example:
	//
	// sgp-l1cbirz451yxuxxx
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s CreateScalingRuleRequestAlarmDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequestAlarmDimensions) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestAlarmDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *CreateScalingRuleRequestAlarmDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *CreateScalingRuleRequestAlarmDimensions) SetDimensionKey(v string) *CreateScalingRuleRequestAlarmDimensions {
	s.DimensionKey = &v
	return s
}

func (s *CreateScalingRuleRequestAlarmDimensions) SetDimensionValue(v string) *CreateScalingRuleRequestAlarmDimensions {
	s.DimensionValue = &v
	return s
}

func (s *CreateScalingRuleRequestAlarmDimensions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingRuleRequestAlarmOptions struct {
	// The period for collecting metric data in the target tracking scaling rule. Unit: seconds. Valid values:
	//
	// - 15
	//
	// - 60
	//
	// - 120
	//
	// - 300
	//
	// - 900
	//
	// > Default value: 60.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s CreateScalingRuleRequestAlarmOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequestAlarmOptions) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestAlarmOptions) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateScalingRuleRequestAlarmOptions) SetPeriod(v int32) *CreateScalingRuleRequestAlarmOptions {
	s.Period = &v
	return s
}

func (s *CreateScalingRuleRequestAlarmOptions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingRuleRequestHybridMetrics struct {
	// The dimensions of the metric. Used to specify the resources monitored by the metric.
	Dimensions []*CreateScalingRuleRequestHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The metric operation expression for multiple Hybrid Cloud Monitoring metrics. The calculation result of the expression is used to trigger scaling.
	//
	// The operation expression must comply with the Reverse Polish Notation (RPN) specification, and only the `+`, `-`, `*`, and `/` operators are supported.
	//
	// example:
	//
	// (a+b)/2
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The reference ID of the metric in the metric operation expression.
	//
	// example:
	//
	// a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the monitoring metrics in the monitoring metrics repository for Hybrid Cloud Monitoring.
	//
	// example:
	//
	// AliyunSmq_NumberOfMessagesVisible
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical method of the metric. Valid values:
	//
	// - Average: the average value of all data points within the specified period.
	//
	// - Minimum: the minimum value of all data points within the specified period.
	//
	// - Maximum: the maximum value of all data points within the specified period.
	//
	// example:
	//
	// Average
	Statistic *string `json:"Statistic,omitempty" xml:"Statistic,omitempty"`
}

func (s CreateScalingRuleRequestHybridMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequestHybridMetrics) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestHybridMetrics) GetDimensions() []*CreateScalingRuleRequestHybridMetricsDimensions {
	return s.Dimensions
}

func (s *CreateScalingRuleRequestHybridMetrics) GetExpression() *string {
	return s.Expression
}

func (s *CreateScalingRuleRequestHybridMetrics) GetId() *string {
	return s.Id
}

func (s *CreateScalingRuleRequestHybridMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateScalingRuleRequestHybridMetrics) GetStatistic() *string {
	return s.Statistic
}

func (s *CreateScalingRuleRequestHybridMetrics) SetDimensions(v []*CreateScalingRuleRequestHybridMetricsDimensions) *CreateScalingRuleRequestHybridMetrics {
	s.Dimensions = v
	return s
}

func (s *CreateScalingRuleRequestHybridMetrics) SetExpression(v string) *CreateScalingRuleRequestHybridMetrics {
	s.Expression = &v
	return s
}

func (s *CreateScalingRuleRequestHybridMetrics) SetId(v string) *CreateScalingRuleRequestHybridMetrics {
	s.Id = &v
	return s
}

func (s *CreateScalingRuleRequestHybridMetrics) SetMetricName(v string) *CreateScalingRuleRequestHybridMetrics {
	s.MetricName = &v
	return s
}

func (s *CreateScalingRuleRequestHybridMetrics) SetStatistic(v string) *CreateScalingRuleRequestHybridMetrics {
	s.Statistic = &v
	return s
}

func (s *CreateScalingRuleRequestHybridMetrics) Validate() error {
	if s.Dimensions != nil {
		for _, item := range s.Dimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScalingRuleRequestHybridMetricsDimensions struct {
	// The dimension key of the metric.
	//
	// example:
	//
	// queue
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension value of the metric.
	//
	// example:
	//
	// testQueue
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s CreateScalingRuleRequestHybridMetricsDimensions) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequestHybridMetricsDimensions) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestHybridMetricsDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *CreateScalingRuleRequestHybridMetricsDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *CreateScalingRuleRequestHybridMetricsDimensions) SetDimensionKey(v string) *CreateScalingRuleRequestHybridMetricsDimensions {
	s.DimensionKey = &v
	return s
}

func (s *CreateScalingRuleRequestHybridMetricsDimensions) SetDimensionValue(v string) *CreateScalingRuleRequestHybridMetricsDimensions {
	s.DimensionValue = &v
	return s
}

func (s *CreateScalingRuleRequestHybridMetricsDimensions) Validate() error {
	return dara.Validate(s)
}

type CreateScalingRuleRequestStepAdjustments struct {
	// The lower boundary of the step adjustment. This parameter is applicable only to step scaling rules. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 1.0
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	// The upper boundary of the step adjustment. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 5.0
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	// The number of instances to be scaled for the step adjustment. This parameter is applicable only to step scaling rules.
	//
	// example:
	//
	// 1
	ScalingAdjustment *int32 `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
}

func (s CreateScalingRuleRequestStepAdjustments) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingRuleRequestStepAdjustments) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequestStepAdjustments) GetMetricIntervalLowerBound() *float32 {
	return s.MetricIntervalLowerBound
}

func (s *CreateScalingRuleRequestStepAdjustments) GetMetricIntervalUpperBound() *float32 {
	return s.MetricIntervalUpperBound
}

func (s *CreateScalingRuleRequestStepAdjustments) GetScalingAdjustment() *int32 {
	return s.ScalingAdjustment
}

func (s *CreateScalingRuleRequestStepAdjustments) SetMetricIntervalLowerBound(v float32) *CreateScalingRuleRequestStepAdjustments {
	s.MetricIntervalLowerBound = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustments) SetMetricIntervalUpperBound(v float32) *CreateScalingRuleRequestStepAdjustments {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustments) SetScalingAdjustment(v int32) *CreateScalingRuleRequestStepAdjustments {
	s.ScalingAdjustment = &v
	return s
}

func (s *CreateScalingRuleRequestStepAdjustments) Validate() error {
	return dara.Validate(s)
}
