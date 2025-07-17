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
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling rules.
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
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRules struct {
	// The adjustment method of the scaling rule. Valid values:
	//
	// 	- QuantityChangeInCapacity: adds or removes the specified number of Elastic Compute Service (ECS) instances to or from the scaling group.
	//
	// 	- PercentChangeInCapacity: adds or removes the specified percentage of ECS instances to or from the scaling group.
	//
	// 	- TotalCapacity: adjusts the number of ECS instances in the scaling group to the specified number.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The number of instances that must be scaled based on the scaling rule.
	//
	// example:
	//
	// 1
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The dimensions. This parameter is applicable to target tracking scaling rules. You can specify this parameter if your predefined metric requires extra dimensions. For example, if you predefine the LoadBalancerRealServerAverageQps metric, you must use this parameter to specify the rulePool dimension.
	AlarmDimensions []*DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions `json:"AlarmDimensions,omitempty" xml:"AlarmDimensions,omitempty" type:"Repeated"`
	// The event-triggered tasks that are associated with the scaling rule. The value of this parameter is returned only if you set ShowAlarmRules to true. Otherwise, null is returned.
	Alarms []*DescribeScalingRulesResponseBodyScalingRulesAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Repeated"`
	// The cooldown period of the scaling rule. This parameter is available only if you set ScalingRuleType to SimpleScalingRule. Valid values: 0 to 86400. Unit: seconds.
	//
	// example:
	//
	// 20
	Cooldown *int32 `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	// Indicates whether scale-in is disabled. This parameter takes effect only if you set ScalingRuleType to TargetTrackingScalingRule. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DisableScaleIn *bool `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	// The warm-up period of instances. During the warm-up period, a series of preparation measures are taken for the new instances. Performance metrics of instances being warmed up are not counted towards the monitoring range.
	//
	// example:
	//
	// 300
	EstimatedInstanceWarmup *int32 `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	// The Hybrid Cloud Monitoring metrics. For more information, see [Create a custom target tracking scaling rule](https://help.aliyun.com/document_detail/2852162.html).
	HybridMetrics []*DescribeScalingRulesResponseBodyScalingRulesHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
	// The ID of the Hybrid Cloud Monitoring namespace.
	//
	// For information about how to manage Hybrid Cloud Monitoring namespaces, see [Manage namespaces](https://help.aliyun.com/document_detail/217606.html).
	//
	// example:
	//
	// aliyun-test
	HybridMonitorNamespace *string `json:"HybridMonitorNamespace,omitempty" xml:"HybridMonitorNamespace,omitempty"`
	// The maximum number of ECS instances that can be contained in the scaling group. If you specify this parameter, you must also specify PredictiveValueBehavior.
	//
	// example:
	//
	// 100
	InitialMaxSize *int32 `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	// The maximum number of ECS instances that can be contained in the scaling group.
	//
	// example:
	//
	// 2
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The name of the metric of the event-triggered task that is associated with the scaling rule.
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The metric type. Valid values:
	//
	// 	- system: system metrics of CloudMonitor.
	//
	// 	- custom: custom metrics that are reported to CloudMonitor.
	//
	// 	- hybrid: metrics of Hybrid Cloud Monitoring.
	//
	// example:
	//
	// system
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The minimum number of instances that must be scaled. This parameter takes effect only if you set ScalingRuleType to SimpleScalingRule or StepScalingRule and set AdjustmentType to PercentChangeInCapacity.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32 `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	// The minimum number of ECS instances that must be contained in the scaling group.
	//
	// example:
	//
	// 1
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	// The mode of the predictive scaling rule. Valid values:
	//
	// 	- PredictAndScale: provides predictions and creates prediction tasks.
	//
	// 	- PredictOnly: provides predictions but does not create prediction tasks.
	//
	// example:
	//
	// PredictAndScale
	PredictiveScalingMode *string `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	// The amount of buffer time before prediction tasks are executed. By default, all prediction tasks that are automatically created based on a predictive scaling rule are executed on the hour. You can specify a buffer time for resource preparation before prediction tasks are executed. Valid values: 0 to 60. Unit: minutes.
	//
	// example:
	//
	// 30
	PredictiveTaskBufferTime *int32 `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	// The action on the predicted maximum value. Valid values:
	//
	// 	- MaxOverridePredictiveValue: uses the initial maximum capacity as the maximum value for prediction tasks if the predicted value is greater than the initial maximum capacity.
	//
	// 	- PredictiveValueOverrideMax: uses the predicted value as the maximum value for prediction tasks when the predicted value is greater than the initial maximum capacity.
	//
	// 	- PredictiveValueOverrideMaxWithBuffer: increases the predicted value by a ratio that is specified by PredictiveValueBuffer, and uses the increased value as the maximum value for prediction tasks if the predicted value increased by this ratio is greater than the initial maximum capacity.
	//
	// example:
	//
	// MaxOverridePredictiveValue
	PredictiveValueBehavior *string `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	// The ratio based on which the predicted value is increased when PredictiveValueBehavior is set to PredictiveValueOverrideMaxWithBuffer. If the predicted value increased by this ratio is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	PredictiveValueBuffer *int32 `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	// The number of consecutive times that the event-triggered task for scale-in purposes must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
	//
	// example:
	//
	// 15
	ScaleInEvaluationCount *int32 `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	// The number of consecutive times that the event-triggered task created for scale-out purposes must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and associated with the target tracking scaling rule.
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
	// 	- SimpleScalingRule: a simple scaling rule. Once a simple scaling rule is executed, Auto Scaling adjusts the number of ECS instances in the scaling group based on the values of AdjustmentType and AdjustmentValue.
	//
	// 	- TargetTrackingScalingRule: a target tracking scaling rule. Once a target tracking scaling rule is executed, Auto Scaling dynamically calculates the number of ECS instances or elastic container instances to scale based on the predefined metric (MetricName) and attempts to maintain the metric value close to the specified target value (TargetValue).
	//
	// 	- StepScalingRule: a step scaling rule. Once a step scaling rule is executed, Auto Scaling scales instances step by step based on the predefined thresholds and metric values.
	//
	// 	- PredictiveScalingRule: a predictive scaling rule. Once a predictive scaling rule is executed, Auto Scaling analyzes the historical monitoring data based on the machine learning technology and predicts the trends of metric data. Auto Scaling also creates scheduled tasks to enable dynamic adjustment of the boundary values for the scaling group.
	//
	// example:
	//
	// SimpleScalingRule
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	// The step adjustments of the step scaling rule.
	StepAdjustments []*DescribeScalingRulesResponseBodyScalingRulesStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	// The target value of the metric. If you set ScalingRuleType to TargetTrackingScalingRule or PredictiveScalingRule, Auto Scaling keeps the metric value close to the target value by adding instances to or removing instances from the scaling group.
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
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRulesAlarmDimensions struct {
	// The dimension key of the metric.
	//
	// example:
	//
	// rulePool
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The dimension value of the metric.
	//
	// example:
	//
	// sgp-l1cbirz451yxu2dxxx
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
	// The ID of the event-triggered task that is associated with the scaling rule.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****_1f9458d1-70e1-4bee-8c7f-7a47695b****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	// The name of the event-triggered task that is associated with the scaling rule.
	//
	// example:
	//
	// alarmtask****
	AlarmTaskName *string `json:"AlarmTaskName,omitempty" xml:"AlarmTaskName,omitempty"`
	// The comparison operator between the statistical value and the threshold of the metric of the event-triggered task that is associated with the scaling rule. The comparison operator indicates the relationship in which the metric value and the metric threshold can meet the alert condition.
	//
	// 	- Valid value if the metric value is greater than or equal to the threshold: >=
	//
	// 	- Valid value if the metric value is less than or equal to the threshold: <=
	//
	// 	- Valid value if the metric value is greater than the threshold: >
	//
	// 	- Valid value if the metric value is less than the threshold: <
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The dimensions of the event-triggered task that is associated with the scaling rule.
	Dimensions []*DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The number of consecutive times when the event-triggered task that is associated with the scaling rule must meet the alert condition before an alert is triggered.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The name of the metric of the event-triggered task that is associated with the scaling rule.
	//
	// example:
	//
	// CpuUtilization
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The type of the metric of the event-triggered task that is associated with the scaling rule. Valid values:
	//
	// 	- system: system metrics
	//
	// 	- custom: custom metrics
	//
	// example:
	//
	// system
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The statistical period of the metric data in the target tracking scaling rule.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The statistical method of the event-triggered task that is associated with the scaling rule. Valid values:
	//
	// 	- Average
	//
	// 	- Maximum
	//
	// 	- Minimum
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert threshold of the event-triggered task that is associated with the scaling rule.
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
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRulesAlarmsDimensions struct {
	// The key of the dimension that is associated with the metric. Valid values:
	//
	// 	- ScalingGroupId: the ID of the scaling group.
	//
	// 	- userId: the ID of the user account.
	//
	// example:
	//
	// scaling_group
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The value of the dimension that is associated with the metric.
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
	// The metric dimensions. This parameter is used to specify the monitored resources.
	Dimensions []*DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The metric expression that consists of multiple Hybrid Cloud Monitoring metrics. It calculates a result used to trigger scaling events.
	//
	// The expression is written in Reverse Polish Notation (RPN) format and supports only the following operators: `+, -, *, /`.
	//
	// example:
	//
	// (a+b)/2
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The reference ID of the metric in the metric expression.
	//
	// example:
	//
	// a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Hybrid Cloud Monitoring metric.
	//
	// example:
	//
	// AliyunSmq_NumberOfMessagesVisible
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The statistical method of the metric value. Valid values:
	//
	// 	- Average: The average value of all metric values within a specified interval is calculated.
	//
	// 	- Minimum: The minimum value of all metric values within a specified interval is calculated.
	//
	// 	- Maximum: The maximum value of all metric values within a specified interval is calculated.
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
	return dara.Validate(s)
}

type DescribeScalingRulesResponseBodyScalingRulesHybridMetricsDimensions struct {
	// The key of the metric dimension.
	//
	// example:
	//
	// queue
	DimensionKey *string `json:"DimensionKey,omitempty" xml:"DimensionKey,omitempty"`
	// The key of the metric dimension.
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
	// The lower limit of a step adjustment. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 1.0
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	// The upper limit of a step adjustment. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 5.0
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	// The number of ECS instances that are scaled in a step adjustment.
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
