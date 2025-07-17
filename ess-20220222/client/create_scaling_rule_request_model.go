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
	// The scaling method of the scaling rule. This parameter is required only if you set the ScalingRuleType parameter to SimpleScalingRule or StepScalingRule. Valid values:
	//
	// 	- QuantityChangeInCapacity: adds the specified number of ECS instances to or removes the specified number of ECS instances from the scaling group.
	//
	// 	- PercentChangeInCapacity: adds the specified percentage of ECS instances to or removes the specified percentage of ECS instances from the scaling group.
	//
	// 	- TotalCapacity: adjusts the number of ECS instances in the scaling group to a specified number.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The number of instances that must be scaled based on the scaling rule. This parameter is required only if you set the ScalingRuleType parameter to SimpleScalingRule or StepScalingRule. The number of ECS instances that are scaled in a single scaling activity cannot exceed 1,000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to QuantityChangeInCapacity: -1000 to 1000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to PercentChangeInCapacity: -100 to 10000.
	//
	// 	- Valid values if you set the AdjustmentType parameter to TotalCapacity: 0 to 2000.
	//
	// example:
	//
	// 100
	AdjustmentValue *int32 `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	// The metric dimensions. This parameter is applicable to target tracking scaling rules. If your predefined metric requires extra dimensions, you must specify this parameter. For example, if you use LoadBalancerRealServerAverageQps as your predefined metric, you must use this parameter to specify the rulePool dimension.
	AlarmDimensions []*CreateScalingRuleRequestAlarmDimensions `json:"AlarmDimensions,omitempty" xml:"AlarmDimensions,omitempty" type:"Repeated"`
	AlarmOptions    *CreateScalingRuleRequestAlarmOptions      `json:"AlarmOptions,omitempty" xml:"AlarmOptions,omitempty" type:"Struct"`
	// The cooldown time of the scaling rule. This parameter is available only if you set the ScalingRuleType parameter to SimpleScalingRule. Valid values: 0 to 86400. Unit: seconds.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 60
	Cooldown *int32 `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	// Specifies whether to disable scale-in. This parameter is available only if you set ScalingRuleType to TargetTrackingScalingRule.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DisableScaleIn *bool `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	// The warmup period of an instance. This parameter is available only if you set the ScalingRuleType parameter to TargetTrackingScalingRule or PredictiveScalingRule. Auto Scaling adds ECS instances that are in the warmup state to a scaling group but does not report monitoring data to CloudMonitor during the warmup period.
	//
	// > Auto Scaling calculates the number of ECS instances that must be scaled. ECS instances in the warmup state are not counted towards the current capacity of the scaling group.
	//
	// Valid values: 0 to 86400. Unit: seconds.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	EstimatedInstanceWarmup *int32 `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	// The Hybrid Cloud Monitoring metrics. For more information, see [Create a custom target tracking scaling rule](https://help.aliyun.com/document_detail/2852162.html).
	HybridMetrics []*CreateScalingRuleRequestHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
	// The ID of the Hybrid Cloud Monitoring namespace.
	//
	// For information about how to manage Hybrid Cloud Monitoring namespaces, see [Manage namespaces](https://help.aliyun.com/document_detail/217606.html).
	//
	// example:
	//
	// aliyun-test
	HybridMonitorNamespace *string `json:"HybridMonitorNamespace,omitempty" xml:"HybridMonitorNamespace,omitempty"`
	// The maximum number of ECS instances that can be contained in the scaling group. If you specify InitialMaxSize, you must specify `PredictiveValueBehavior`.
	//
	// The default value of this parameter is the value of MaxSize.
	//
	// example:
	//
	// 100
	InitialMaxSize *int32 `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	// The predefined metric that you want to monitor. If you set ScalingRuleType to TargetTrackingScalingRule or PredictiveScalingRule, you must specify this parameter.
	//
	// Valid values if you set ScalingRuleType to TargetTrackingScalingRule:
	//
	// 	- CpuUtilizationAgent (recommended): the CPU utilization.
	//
	// 	- MemoryUtilization (recommended): the memory usage.
	//
	// 	- CpuUtilization: the average CPU utilization.
	//
	// 	- IntranetTx: the outbound traffic over an internal network.
	//
	// 	- IntranetRx: the average inbound traffic over an internal network.
	//
	// 	- VpcInternetTx: the outbound traffic from a virtual private cloud (VPC) to the Internet.
	//
	// 	- VpcInternetRx: the inbound traffic from the Internet to a VPC.
	//
	// 	- LoadBalancerRealServerAverageQps:the queries per second (QPS) per Application Load Balancer (ALB) server group.
	//
	// Valid values if you set ScalingRuleType to PredictiveScalingRule:
	//
	// 	- CpuUtilization: the average CPU utilization.
	//
	// 	- IntranetRx: the average inbound traffic over an internal network.
	//
	// 	- IntranetTx: the average outbound traffic over an internal network.
	//
	// For more information, see [Event-triggered tasks of the system monitoring type](https://help.aliyun.com/document_detail/74854.html).
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
	// The minimum number of instances that must be scaled when the AdjustmentType parameter is set to PercentChangeInCapacity. This parameter takes effect only if you set the ScalingRuleType parameter to SimpleScalingRule or StepScalingRule.
	//
	// example:
	//
	// 1
	MinAdjustmentMagnitude *int32  `json:"MinAdjustmentMagnitude,omitempty" xml:"MinAdjustmentMagnitude,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The mode of the predictive scaling rule. Valid values:
	//
	// 	- PredictAndScale: produces predictions and creates prediction tasks.
	//
	// 	- PredictOnly: produces predictions but does not create prediction tasks.
	//
	// Default value: PredictAndScale.
	//
	// example:
	//
	// PredictAndScale
	PredictiveScalingMode *string `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	// The amount of buffer time before the prediction task is executed. By default, all prediction tasks that are automatically created for a predictive scaling rule are executed on the hour. You can specify an amount of buffer time for resource preparation before the prediction tasks are executed. Valid values: 0 to 60. Unit: minutes.
	//
	// Default value: 0.
	//
	// example:
	//
	// 30
	PredictiveTaskBufferTime *int32 `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	// The maximum value for predication tasks. Valid values:
	//
	// 	- MaxOverridePredictiveValue: uses the initial maximum capacity as the maximum value for prediction tasks if the predicted value is greater than the initial maximum capacity.
	//
	// 	- PredictiveValueOverrideMax: uses the predicted value as the maximum value for prediction tasks if the predicted value is greater than the initial maximum capacity.
	//
	// 	- PredictiveValueOverrideMaxWithBuffer: increases the predicted value by a percentage that is specified by the PredictiveValueBuffer parameter. If the predicted value that is increased by the percentage is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks.
	//
	// Default value: MaxOverridePredictiveValue.
	//
	// example:
	//
	// MaxOverridePredictiveValue
	PredictiveValueBehavior *string `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	// The ratio based on which the predicted value is increased when you set `PredictiveValueBehavior` to `PredictiveValueOverrideMaxWithBuffer`. If the predicted value increased by this ratio is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks. Valid values: 0 to 100.
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
	// The number of consecutive times that the event-triggered task created for scale-in activities must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and then associated with the target tracking scaling rule.
	//
	// Default value: 15.
	//
	// example:
	//
	// 15
	ScaleInEvaluationCount *int32 `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	// The number of consecutive times that the event-triggered task created for scale-out activities must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and then associated with the target tracking scaling rule.
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
	// The name of the scaling rule. The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). The name must start with a letter or a digit.
	//
	// The name of each scaling rule must be unique under the same account within a region.
	//
	// If you leave this parameter empty, the scaling rule ID is used.
	//
	// example:
	//
	// scalingrule****
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// The type of the scaling rule. Valid values:
	//
	// 	- SimpleScalingRule: a simple scaling rule. After you execute a simple scaling rule, Auto Scaling adjusts the number of ECS instances or elastic container instances in the scaling group based on the values of AdjustmentType and AdjustmentValue.
	//
	// 	- TargetTrackingScalingRule: a target tracking scaling rule. After you execute a target tracking scaling rule, Auto Scaling dynamically calculates the number of ECS instances or elastic container instances to scale based on the predefined metric (MetricName) and attempts to maintain the metric value close to the expected value (TargetValue).
	//
	// 	- StepScalingRule: a step scaling rule. After you execute a step scaling rule, Auto Scaling scales instances step by step based on the predefined thresholds and metric values.
	//
	// 	- PredictiveScalingRule: a predictive scaling rule. After you execute a predictive scaling rule, Auto Scaling uses machine learning to analyze historical monitoring data of the scaling group and predicts the future values of metrics. In addition, Auto Scaling automatically creates scheduled tasks to adjust the boundary values for the scaling group.
	//
	// Default value: SimpleScalingRule.
	//
	// example:
	//
	// SimpleScalingRule
	ScalingRuleType *string `json:"ScalingRuleType,omitempty" xml:"ScalingRuleType,omitempty"`
	// Details of the step adjustments.
	StepAdjustments []*CreateScalingRuleRequestStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	// The target value. This parameter is required only if you set the ScalingRuleType parameter to TargetTrackingScalingRule or PredictiveScalingRule. The value must be greater than 0 and can have up to three decimal places.
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
	return dara.Validate(s)
}

type CreateScalingRuleRequestAlarmDimensions struct {
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
	// The metric dimensions. You can use this parameter to specify the monitored resources.
	Dimensions []*CreateScalingRuleRequestHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The metric expression that consists of multiple Hybrid Cloud Monitoring metrics. It calculates a result used to trigger scaling events.
	//
	// The expression must be written in Reverse Polish Notation (RPN) format and supports only the following operators: `+, -, *, /`.
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
	// 	- Average: calculates the average value of all metric values within a specified interval.
	//
	// 	- Minimum: calculates the minimum value of all metric values within a specified interval.
	//
	// 	- Maximum: calculates the maximum value of all metric values within a specified interval.
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
	return dara.Validate(s)
}

type CreateScalingRuleRequestHybridMetricsDimensions struct {
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
	// The lower limit specified in a step adjustment. This parameter is available only if you set the ScalingRuleType parameter to StepScalingRule. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 1.0
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	// The upper limit that is specified in a step adjustment. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 5.0
	MetricIntervalUpperBound *float32 `json:"MetricIntervalUpperBound,omitempty" xml:"MetricIntervalUpperBound,omitempty"`
	// The number of ECS instances that you want to scale in a step adjustment. This parameter is available only if you set the ScalingRuleType parameter to StepScalingRule.
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
