// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeScalingRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeScalingRulesResponseBody
	GetRequestId() *string
	SetScalingRules(v []*DescribeScalingRulesResponseBodyScalingRules) *DescribeScalingRulesResponseBody
	GetScalingRules() []*DescribeScalingRulesResponseBodyScalingRules
	SetTotalCount(v int32) *DescribeScalingRulesResponseBody
	GetTotalCount() *int32
}

type DescribeScalingRulesResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the scaling rules.
	ScalingRules []*DescribeScalingRulesResponseBodyScalingRules `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
	// The total number of scaling rules.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingRulesResponseBody) GetScalingRules() []*DescribeScalingRulesResponseBodyScalingRules {
	return s.ScalingRules
}

func (s *DescribeScalingRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScalingRulesResponseBody) SetPageNumber(v int32) *DescribeScalingRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetPageSize(v int32) *DescribeScalingRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetRequestId(v string) *DescribeScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetScalingRules(v []*DescribeScalingRulesResponseBodyScalingRules) *DescribeScalingRulesResponseBody {
	s.ScalingRules = v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetTotalCount(v int32) *DescribeScalingRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) Validate() error {
	if s.ScalingRules != nil {
		for _, item := range s.ScalingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScalingRulesResponseBodyScalingRules struct {
	// The adjustment type of the scaling rule. Valid values:
	//
	// - QuantityChangeInCapacity: increases or decreases a specified number of ECS instances.
	//
	// - PercentChangeInCapacity: increases or decreases a specified percentage of ECS instances.
	//
	// - TotalCapacity: adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value of the scaling rule.
	//
	// example:
	//
	// 1
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The dimension value of the metric. This parameter applies to target tracking scaling rules and is configured when the metric requires additional dimension information. For example, the LoadBalancerRealServerAverageQps metric requires the rulePool dimension information.
	AlarmDimensions []*DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions `json:"AlarmDimensions,omitempty" xml:"AlarmDimensions,omitempty" type:"Repeated"`
	// The CloudMonitor event-triggered tasks associated with the scaling rule. The associated CloudMonitor event-triggered tasks are returned only when the ShowAlarmRules parameter is set to true. Otherwise, an empty list is returned.
	Alarms []*DescribeScalingRulesResponseBodyScalingRulesAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// The cooldown time of the scaling rule. This parameter applies only to simple scaling rules. Valid values: 0 to 86400. Unit: seconds.
	//
	// example:
	//
	// 20
	Cooldown *int32 `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	// Specifies whether to disable scale-in. This parameter applies only to target tracking scaling rules. Valid values:
	//
	// - true: disables scale-in.
	//
	// - false: allows scale-in.
	//
	// example:
	//
	// true
	DisableScaleIn *bool `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	// The warmup time required for newly created instances to complete a series of preparations. During the warmup period, the started instances do not affect the monitoring metrics of the scaling group.
	//
	// example:
	//
	// 300
	EstimatedInstanceWarmup *int32 `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	// The monitoring metrics settings for Hybrid Cloud Monitoring. For more information about how to configure the settings, see [Use formula operations to customize target tracking scaling rules](https://help.aliyun.com/document_detail/2852162.html).
	HybridMetrics []*DescribeScalingRulesResponseBodyScalingRulesHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
	// The ID of the monitoring metrics repository for Hybrid Cloud Monitoring.
	//
	// To manage monitoring metrics repositories, see [Manage monitoring metrics repositories](https://help.aliyun.com/document_detail/217606.html).
	//
	// example:
	//
	// aliyun-test
	HybridMonitorNamespace *string `json:"HybridMonitorNamespace,omitempty" xml:"HybridMonitorNamespace,omitempty"`
	// The maximum number of instances in the scaling group. This parameter is used in combination with PredictiveValueBehavior.
	//
	// example:
	//
	// 100
	InitialMaxSize *int32 `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	// The maximum number of instances in the scaling group.
	//
	// example:
	//
	// 2
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The name of the metric for the event-triggered task associated with the scaling rule.
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
	// The minimum number of instances to be adjusted by the scaling rule. This parameter takes effect only when the scaling rule type is SimpleScalingRule or StepScalingRule and the AdjustmentType is PercentChangeInCapacity.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The minimum number of instances in the scaling group.
	//
	// example:
	//
	// 1
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The mode of the predictive scaling rule. Valid values:
	//
	// - PredictAndScale: generates prediction results and creates predictive tasks.
	//
	// - PredictOnly: generates prediction results but does not create predictive tasks.
	//
	// example:
	//
	// PredictAndScale
	PredictiveScalingMode *string `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	// The predictive tasks automatically created by the predictive scaling rule are executed on the hour by default. You can set the pre-start time to execute the predictive tasks in advance and prepare resources. Valid values: 0 to 60. Unit: minutes.
	//
	// example:
	//
	// 30
	PredictiveTaskBufferTime *int32 `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	// The processing method for the maximum value of the predictive scaling rule. Valid values:
	//
	// - MaxOverridePredictiveValue: the initial maximum value overrides the predicted value. If the predicted value is greater than the initial maximum value, the maximum value of the predictive task uses the initial maximum value.
	//
	// - PredictiveValueOverrideMax: the predicted value overrides the initial maximum value. If the predicted value is greater than the initial maximum value, the maximum value of the predictive task uses the predicted value.
	//
	// - PredictiveValueOverrideMaxWithBuffer: a certain percentage is added to the predicted value. The predicted value is increased by the percentage specified by PredictiveValueBuffer. If the increased value is greater than the initial maximum value, the increased value is used.
	//
	// example:
	//
	// MaxOverridePredictiveValue
	PredictiveValueBehavior *string `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	// This parameter takes effect when PredictiveValueBehavior is set to PredictiveValueOverrideMaxWithBuffer. The predicted value is increased by this percentage. If the increased value is greater than the initial maximum value, the increased value is used. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	PredictiveValueBuffer *int32 `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	// After a target tracking scaling rule is created, an event-triggered task is automatically created. This parameter specifies the number of consecutive times the threshold condition must be met to trigger the scale-in event-triggered task.
	//
	// example:
	//
	// 15
	ScaleInEvaluationCount *int32 `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	// After a target tracking scaling rule is created, an event-triggered task is automatically created. This parameter specifies the number of consecutive times the threshold condition must be met to trigger the scale-out event-triggered task.
	//
	// example:
	//
	// 3
	ScaleOutEvaluationCount *int32 `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The unique identifier of the scaling rule.
	//
	// example:
	//
	// ari:acs:ess:cn-hangzhou:140692647406****:scalingrule/asr-bp1dvirgwkoowxk7****
	ScalingRuleAri *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	// The ID of the scaling rule.
	//
	// example:
	//
	// asr-bp1dvirgwkoowxk7****
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	// The name of the scaling rule.
	//
	// example:
	//
	// scalingrule****
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// The type of the scaling rule. Valid values:
	//
	// - SimpleScalingRule: simple scaling rule. Adjusts the number of ECS instances based on the adjustment type (AdjustmentType) and adjustment value (AdjustmentValue).
	//
	// - TargetTrackingScalingRule: target tracking scaling rule. Dynamically calculates the number of ECS instances to scale out or scale-in based on a predefined metric (MetricName) and attempts to maintain the metric value near the target value (TargetValue).
	//
	// - StepScalingRule: step scaling rule. Provides a step-by-step extension method based on thresholds and metric values.
	//
	// - PredictiveScalingRule: predictive scaling rule. Analyzes the historical monitoring data of the scaling group based on machine learning capabilities to predict future monitoring metrics, and supports the automatic creation of scheduled tasks to set the boundaries of the scaling group.
	//
	// example:
	//
	// SimpleScalingRule
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	// The step adjustments of the step scaling rule.
	StepAdjustments []*DescribeScalingRulesResponseBodyScalingRulesStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	// The target value. When the scaling rule is a target tracking scaling rule or a predictive scaling rule, the scaling rule maintains the metric near the target value by adding or removing instances.
	//
	// example:
	//
	// 0.125
	TargetValue *float32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetAlarmDimensions() []*DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions {
	return s.AlarmDimensions
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetAlarms() []*DescribeScalingRulesResponseBodyScalingRulesAlarms {
	return s.Alarms
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetCooldown() *int32 {
	return s.Cooldown
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetDisableScaleIn() *bool {
	return s.DisableScaleIn
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetEstimatedInstanceWarmup() *int32 {
	return s.EstimatedInstanceWarmup
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetHybridMetrics() []*DescribeScalingRulesResponseBodyScalingRulesHybridMetrics {
	return s.HybridMetrics
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetHybridMonitorNamespace() *string {
	return s.HybridMonitorNamespace
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetInitialMaxSize() *int32 {
	return s.InitialMaxSize
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetMinAdjustmentMagnitude() *int32 {
	return s.MinAdjustmentMagnitude
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetMinSize() *int32 {
	return s.MinSize
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetPredictiveScalingMode() *string {
	return s.PredictiveScalingMode
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetPredictiveTaskBufferTime() *int32 {
	return s.PredictiveTaskBufferTime
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetPredictiveValueBehavior() *string {
	return s.PredictiveValueBehavior
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetPredictiveValueBuffer() *int32 {
	return s.PredictiveValueBuffer
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScaleInEvaluationCount() *int32 {
	return s.ScaleInEvaluationCount
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScaleOutEvaluationCount() *int32 {
	return s.ScaleOutEvaluationCount
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScalingRuleAri() *string {
	return s.ScalingRuleAri
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetScalingRuleType() *string {
	return s.ScalingRuleType
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetStepAdjustments() []*DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	return s.StepAdjustments
}

func (s *DescribeScalingRulesResponseBodyScalingRules) GetTargetValue() *float32 {
	return s.TargetValue
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAdjustmentType(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAdjustmentValue(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.AdjustmentValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAlarmDimensions(v []*DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) *DescribeScalingRulesResponseBodyScalingRules {
	s.AlarmDimensions = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetAlarms(v []*DescribeScalingRulesResponseBodyScalingRulesAlarms) *DescribeScalingRulesResponseBodyScalingRules {
	s.Alarms = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetCooldown(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.Cooldown = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetDisableScaleIn(v bool) *DescribeScalingRulesResponseBodyScalingRules {
	s.DisableScaleIn = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetEstimatedInstanceWarmup(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetHybridMetrics(v []*DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) *DescribeScalingRulesResponseBodyScalingRules {
	s.HybridMetrics = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetHybridMonitorNamespace(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.HybridMonitorNamespace = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetInitialMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.InitialMaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMetricType(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.MetricType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMinAdjustmentMagnitude(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetMinSize(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveScalingMode(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveScalingMode = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveTaskBufferTime(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveValueBehavior(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetPredictiveValueBuffer(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScaleInEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScaleOutEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingGroupId(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleAri(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleAri = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleId(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleName(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRuleType(v string) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRuleType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetStepAdjustments(v []*DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) *DescribeScalingRulesResponseBodyScalingRules {
	s.StepAdjustments = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetTargetValue(v float32) *DescribeScalingRulesResponseBodyScalingRules {
	s.TargetValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRules) Validate() error {
	if s.AlarmDimensions != nil {
		for _, item := range s.AlarmDimensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Alarms != nil {
		for _, item := range s.Alarms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions struct {
	// The key of the dimension associated with the metric.
	//
	// example:
	//
	// rulePool
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the dimension associated with the metric.
	//
	// example:
	//
	// sgp-l1cbirz451yxu2****
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) SetDimensionKey(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) SetDimensionValue(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions {
	s.DimensionValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRulesAlarms struct {
	// The ID of the event-triggered task associated with the scaling rule.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****_1f9458d1-70e1-4bee-8c7f-7a47695b****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The name of the event-triggered task associated with the scaling rule.
	//
	// example:
	//
	// alarmtask****
	AlarmTaskName *string `json:"AlarmTaskName,omitempty" xml:"AlarmTaskName,omitempty"`
	// The comparison operator used to compare the statistical value of the metric with the threshold for the event-triggered task associated with the scaling rule. This operator specifies the relationship between the statistical value and the threshold that meets the condition. Valid values:
	//
	// - >=: the statistical value of the metric is greater than or equal to the threshold.
	//
	// - <=: the statistical value of the metric is less than or equal to the threshold.
	//
	// - >: the statistical value of the metric is greater than the threshold.
	//
	// - <: the statistical value of the metric is less than the threshold.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The dimension information of the event-triggered task associated with the scaling rule.
	Dimensions []*DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The number of consecutive times the threshold expression must be met for the event-triggered task associated with the scaling rule to reach the alert state.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The name of the metric for the event-triggered task associated with the scaling rule.
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the event-triggered task associated with the scaling rule. Valid values:
	//
	// - system: system monitoring event-triggered task.
	//
	// - custom: custom monitoring event-triggered task.
	//
	// example:
	//
	// system
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The period for collecting metric data in the target tracking scaling rule.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The statistical method of the event-triggered task associated with the scaling rule. Valid values:
	//
	// - Average: the average value.
	//
	// - Maximum: the maximum value.
	//
	// - Minimum: the minimum value.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert threshold of the event-triggered task associated with the scaling rule.
	//
	// example:
	//
	// 50
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarms) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarms) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetAlarmTaskName() *string {
	return s.AlarmTaskName
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetDimensions() []*DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions {
	return s.Dimensions
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetAlarmTaskId(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.AlarmTaskId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetAlarmTaskName(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.AlarmTaskName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetComparisonOperator(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetDimensions(v []*DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Dimensions = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetEvaluationCount(v int32) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetMetricType(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.MetricType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetPeriod(v int32) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Period = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetStatistics(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Statistics = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) SetThreshold(v float32) *DescribeScalingRulesResponseBodyScalingRulesAlarms {
	s.Threshold = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarms) Validate() error {
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

type DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions struct {
	// The key of the dimension associated with the metric. Valid values:
	//
	// - scaling_group: the ID of the scaling group.
	//
	// - userId: the ID of the user account.
	//
	// example:
	//
	// scaling_group
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the dimension associated with the metric.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) SetDimensionKey(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) SetDimensionValue(v string) *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions {
	s.DimensionValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRulesHybridMetrics struct {
	// The dimensions of the metric. Used to specify the resources monitored by the metric.
	Dimensions []*DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The metric operation expression for multiple Hybrid Cloud Monitoring metrics. The calculation result of the expression is used to trigger scaling.
	//
	// The operation expression must meet the Reverse Polish Notation (RPN) specification, and only the `+ - 	- /` operators are supported.
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
	// - Average: the average value of all data points in the specified period.
	//
	// - Minimum: the minimum value of all data points in the specified period.
	//
	// - Maximum: the maximum value of all data points in the specified period.
	//
	// example:
	//
	// Average
	Statistic *string `json:"Statistic,omitempty" xml:"Statistic,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) GetDimensions() []*DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions {
	return s.Dimensions
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) GetExpression() *string {
	return s.Expression
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) GetId() *string {
	return s.Id
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) GetStatistic() *string {
	return s.Statistic
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) SetDimensions(v []*DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics {
	s.Dimensions = v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) SetExpression(v string) *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics {
	s.Expression = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) SetId(v string) *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics {
	s.Id = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) SetMetricName(v string) *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) SetStatistic(v string) *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics {
	s.Statistic = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetrics) Validate() error {
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

type DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions struct {
	// The key of the metric dimension.
	//
	// example:
	//
	// queue
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the metric dimension.
	//
	// example:
	//
	// testQueue
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) SetDimensionKey(v string) *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions {
	s.DimensionKey = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) SetDimensionValue(v string) *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions {
	s.DimensionValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRulesStepAdjustments struct {
	// The lower boundary of the step adjustment. Valid values: -9.999999E18 to 9.999999E18.
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
	// The number of instances to be scaled for the step adjustment.
	//
	// example:
	//
	// 1
	ScalingAdjustment *int32 `json:"ScalingAdjustment,omitempty" xml:"ScalingAdjustment,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) GetMetricIntervalLowerBound() *float32 {
	return s.MetricIntervalLowerBound
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) GetMetricIntervalUpperBound() *float32 {
	return s.MetricIntervalUpperBound
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) GetScalingAdjustment() *int32 {
	return s.ScalingAdjustment
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) SetMetricIntervalLowerBound(v float32) *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	s.MetricIntervalLowerBound = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) SetMetricIntervalUpperBound(v float32) *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) SetScalingAdjustment(v int32) *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments {
	s.ScalingAdjustment = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesStepAdjustments) Validate() error {
	return dara.Validate(s)
}
