// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustmentType(v string) *ModifyScalingRuleRequest
	GetAdjustmentType() *string
	SetAdjustmentValue(v int32) *ModifyScalingRuleRequest
	GetAdjustmentValue() *int32
	SetAlarmDimensions(v []*ModifyScalingRuleRequestAlarmDimensions) *ModifyScalingRuleRequest
	GetAlarmDimensions() []*ModifyScalingRuleRequestAlarmDimensions
	SetAlarmOptions(v *ModifyScalingRuleRequestAlarmOptions) *ModifyScalingRuleRequest
	GetAlarmOptions() *ModifyScalingRuleRequestAlarmOptions
	SetCooldown(v int32) *ModifyScalingRuleRequest
	GetCooldown() *int32
	SetDisableScaleIn(v bool) *ModifyScalingRuleRequest
	GetDisableScaleIn() *bool
	SetEstimatedInstanceWarmup(v int32) *ModifyScalingRuleRequest
	GetEstimatedInstanceWarmup() *int32
	SetHybridMetrics(v []*ModifyScalingRuleRequestHybridMetrics) *ModifyScalingRuleRequest
	GetHybridMetrics() []*ModifyScalingRuleRequestHybridMetrics
	SetHybridMonitorNamespace(v string) *ModifyScalingRuleRequest
	GetHybridMonitorNamespace() *string
	SetInitialMaxSize(v int32) *ModifyScalingRuleRequest
	GetInitialMaxSize() *int32
	SetMetricName(v string) *ModifyScalingRuleRequest
	GetMetricName() *string
	SetMetricType(v string) *ModifyScalingRuleRequest
	GetMetricType() *string
	SetMinAdjustmentMagnitude(v int32) *ModifyScalingRuleRequest
	GetMinAdjustmentMagnitude() *int32
	SetOwnerAccount(v string) *ModifyScalingRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyScalingRuleRequest
	GetOwnerId() *int64
	SetPredictiveScalingMode(v string) *ModifyScalingRuleRequest
	GetPredictiveScalingMode() *string
	SetPredictiveTaskBufferTime(v int32) *ModifyScalingRuleRequest
	GetPredictiveTaskBufferTime() *int32
	SetPredictiveValueBehavior(v string) *ModifyScalingRuleRequest
	GetPredictiveValueBehavior() *string
	SetPredictiveValueBuffer(v int32) *ModifyScalingRuleRequest
	GetPredictiveValueBuffer() *int32
	SetResourceOwnerAccount(v string) *ModifyScalingRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyScalingRuleRequest
	GetResourceOwnerId() *int64
	SetScaleInEvaluationCount(v int32) *ModifyScalingRuleRequest
	GetScaleInEvaluationCount() *int32
	SetScaleOutEvaluationCount(v int32) *ModifyScalingRuleRequest
	GetScaleOutEvaluationCount() *int32
	SetScalingRuleId(v string) *ModifyScalingRuleRequest
	GetScalingRuleId() *string
	SetScalingRuleName(v string) *ModifyScalingRuleRequest
	GetScalingRuleName() *string
	SetStepAdjustments(v []*ModifyScalingRuleRequestStepAdjustments) *ModifyScalingRuleRequest
	GetStepAdjustments() []*ModifyScalingRuleRequestStepAdjustments
	SetTargetValue(v float32) *ModifyScalingRuleRequest
	GetTargetValue() *float32
}

type ModifyScalingRuleRequest struct {
	// The adjustment method of the scaling rule. This is required when the ScalingRuleType parameter is set to SimpleScalingRule or StepScalingRule. Valid values:
	//
	// 	- QuantityChangeInCapacity: adds the specified number of ECS instances to or removes the specified number of ECS instances from the scaling group.
	//
	// 	- PercentChangeInCapacity: adds the specified percentage of ECS instances to or removes the specified percentage of ECS instances from the scaling group.
	//
	// 	- TotalCapacity: adjusts the number of ECS instances in the scaling group to the specified number.
	//
	// example:
	//
	// QuantityChangeInCapacity
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The target value specified in the scaling rule. This parameter is required when the ScalingRuleType parameter is set to SimpleScalingRule or StepScalingRule. The number of ECS instances that are scaled in a single scaling activity cannot exceed 1,000.
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
	// The dimensions. This parameter is applicable to target tracking scaling rules. You can specify this parameter if your predefined metric requires extra dimensions. For example, if you predefine the LoadBalancerRealServerAverageQps metric, you must use this parameter to specify the rulePool dimension.
	AlarmDimensions []*ModifyScalingRuleRequestAlarmDimensions `json:"AlarmDimensions,omitempty" xml:"AlarmDimensions,omitempty" type:"Repeated"`
	// The alert attributes.
	AlarmOptions *ModifyScalingRuleRequestAlarmOptions `json:"AlarmOptions,omitempty" xml:"AlarmOptions,omitempty" type:"Struct"`
	// The cooldown time of the scaling rule. This parameter is available only if you set the ScalingRuleType parameter to SimpleScalingRule.
	//
	// Valid values: 0 to 86400. Unit: seconds.
	//
	// example:
	//
	// 60
	Cooldown *int32 `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	// Specifies whether to disable scale-in. This parameter is available only if you set the ScalingRuleType parameter to TargetTrackingScalingRule.
	//
	// example:
	//
	// true
	DisableScaleIn *bool `json:"DisableScaleIn,omitempty" xml:"DisableScaleIn,omitempty"`
	// The warmup period of an instance. This parameter is available only if you set the ScalingRuleType parameter to TargetTrackingScalingRule or PredictiveScalingRule. Auto Scaling adds ECS instances that are in the warmup state to a scaling group but does not report monitoring data to CloudMonitor during the warmup period.
	//
	// > Auto Scaling calculates the number of ECS instances that need to be scaled. ECS instances in the warmup state are not counted towards the current capacity of the scaling group.
	//
	// Valid values: 0 to 86400. Unit: seconds.
	//
	// example:
	//
	// 60
	EstimatedInstanceWarmup *int32 `json:"EstimatedInstanceWarmup,omitempty" xml:"EstimatedInstanceWarmup,omitempty"`
	// The Hybrid Cloud Monitoring metrics. For more information, see [Create a custom target tracking scaling rule](https://help.aliyun.com/document_detail/2852162.html).
	HybridMetrics []*ModifyScalingRuleRequestHybridMetrics `json:"HybridMetrics,omitempty" xml:"HybridMetrics,omitempty" type:"Repeated"`
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
	// example:
	//
	// 100
	InitialMaxSize *int32 `json:"InitialMaxSize,omitempty" xml:"InitialMaxSize,omitempty"`
	// The predefined metric. This parameter is required only if you create a target tracking scaling rule or predictive scaling rule.
	//
	// Valid values if you create a target tracking scaling rule:
	//
	// 	- CpuUtilizationAgent (recommended): the CPU utilization.
	//
	// 	- MemoryUtilization (recommended): the memory usage.
	//
	// 	- CpuUtilization: the average CPU utilization.
	//
	// 	- IntranetTx: the average outbound traffic over an internal network.
	//
	// 	- IntranetRx: the average inbound traffic over an internal network.
	//
	// 	- VpcInternetTx: the average outbound traffic from a virtual private cloud (VPC) to the Internet.
	//
	// 	- VpcInternetRx: the average inbound traffic from the Internet to a VPC.
	//
	// 	- LoadBalancerRealServerAverageQps: the queries per second (QPS) per Application Load Balancer (ALB) server group.
	//
	// Valid values if you create a predictive scaling rule:
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
	// The minimum number of instances to scale. This parameter takes effect only if you create a simple scaling rule or step scaling rule and set `AdjustmentType` to `PercentChangeInCapacity`.
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
	// example:
	//
	// PredictAndScale
	PredictiveScalingMode *string `json:"PredictiveScalingMode,omitempty" xml:"PredictiveScalingMode,omitempty"`
	// The amount of buffer time before the prediction task is executed. By default, all prediction tasks that are automatically created for a predictive scaling rule are executed on the hour. You can specify an amount of buffer time for resource preparation before the prediction tasks are executed. Valid values: 0 to 60.
	//
	// example:
	//
	// 30
	PredictiveTaskBufferTime *int32 `json:"PredictiveTaskBufferTime,omitempty" xml:"PredictiveTaskBufferTime,omitempty"`
	// Specifies which one of the initial maximum capacity and the predicted value can be used as the maximum value for prediction tasks. Valid values:
	//
	// 	- MaxOverridePredictiveValue: uses the initial maximum capacity as the maximum value for prediction tasks if the predicted value is greater than the initial maximum capacity.
	//
	// 	- PredictiveValueOverrideMax: uses the predicted value as the maximum value for prediction tasks when the predicted value is greater than the initial maximum capacity.
	//
	// 	- PredictiveValueOverrideMaxWithBuffer: increases the predicted value by a percentage that is specified by the PredictiveValueBuffer parameter. If the predicted value that is increased by the percentage is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks.
	//
	// example:
	//
	// MaxOverridePredictiveValue
	PredictiveValueBehavior *string `json:"PredictiveValueBehavior,omitempty" xml:"PredictiveValueBehavior,omitempty"`
	// The ratio based on which the predicted value is increased when `PredictiveValueBehavior` is set to `PredictiveValueOverrideMaxWithBuffer`. If the predicted value increased by this ratio is greater than the initial maximum capacity, the increased value is used as the maximum value for prediction tasks. Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	PredictiveValueBuffer *int32  `json:"PredictiveValueBuffer,omitempty" xml:"PredictiveValueBuffer,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The number of consecutive times that the event-triggered task created for scale-in activities must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and then associated with the target tracking scaling rule.
	//
	// example:
	//
	// 15
	ScaleInEvaluationCount *int32 `json:"ScaleInEvaluationCount,omitempty" xml:"ScaleInEvaluationCount,omitempty"`
	// The number of consecutive times that the event-triggered task created for scale-out activities must meet the threshold conditions before an alert is triggered. After a target tracking scaling rule is created, an event-triggered task is automatically created and then associated with the target tracking scaling rule.
	//
	// example:
	//
	// 3
	ScaleOutEvaluationCount *int32 `json:"ScaleOutEvaluationCount,omitempty" xml:"ScaleOutEvaluationCount,omitempty"`
	// The ID of the scaling rule that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// asr-bp1dvirgwkoowxk7****
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	// The name of the scaling rule. The name must be 2 to 64 letters in length and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter or digit.
	//
	// The name of each scaling rule must be unique under the same account within the same region.
	//
	// example:
	//
	// scalingrule****
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// Details of the step adjustments.
	StepAdjustments []*ModifyScalingRuleRequestStepAdjustments `json:"StepAdjustments,omitempty" xml:"StepAdjustments,omitempty" type:"Repeated"`
	// The target value. This parameter is available only if you set the ScalingRuleType parameter to TargetTrackingScalingRule or PredictiveScalingRule. The value must be greater than 0 and can have up to three decimal places.
	//
	// example:
	//
	// 0.125
	TargetValue *float32 `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ModifyScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *ModifyScalingRuleRequest) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *ModifyScalingRuleRequest) GetAlarmDimensions() []*ModifyScalingRuleRequestAlarmDimensions {
	return s.AlarmDimensions
}

func (s *ModifyScalingRuleRequest) GetAlarmOptions() *ModifyScalingRuleRequestAlarmOptions {
	return s.AlarmOptions
}

func (s *ModifyScalingRuleRequest) GetCooldown() *int32 {
	return s.Cooldown
}

func (s *ModifyScalingRuleRequest) GetDisableScaleIn() *bool {
	return s.DisableScaleIn
}

func (s *ModifyScalingRuleRequest) GetEstimatedInstanceWarmup() *int32 {
	return s.EstimatedInstanceWarmup
}

func (s *ModifyScalingRuleRequest) GetHybridMetrics() []*ModifyScalingRuleRequestHybridMetrics {
	return s.HybridMetrics
}

func (s *ModifyScalingRuleRequest) GetHybridMonitorNamespace() *string {
	return s.HybridMonitorNamespace
}

func (s *ModifyScalingRuleRequest) GetInitialMaxSize() *int32 {
	return s.InitialMaxSize
}

func (s *ModifyScalingRuleRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyScalingRuleRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *ModifyScalingRuleRequest) GetMinAdjustmentMagnitude() *int32 {
	return s.MinAdjustmentMagnitude
}

func (s *ModifyScalingRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyScalingRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyScalingRuleRequest) GetPredictiveScalingMode() *string {
	return s.PredictiveScalingMode
}

func (s *ModifyScalingRuleRequest) GetPredictiveTaskBufferTime() *int32 {
	return s.PredictiveTaskBufferTime
}

func (s *ModifyScalingRuleRequest) GetPredictiveValueBehavior() *string {
	return s.PredictiveValueBehavior
}

func (s *ModifyScalingRuleRequest) GetPredictiveValueBuffer() *int32 {
	return s.PredictiveValueBuffer
}

func (s *ModifyScalingRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyScalingRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyScalingRuleRequest) GetScaleInEvaluationCount() *int32 {
	return s.ScaleInEvaluationCount
}

func (s *ModifyScalingRuleRequest) GetScaleOutEvaluationCount() *int32 {
	return s.ScaleOutEvaluationCount
}

func (s *ModifyScalingRuleRequest) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *ModifyScalingRuleRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *ModifyScalingRuleRequest) GetStepAdjustments() []*ModifyScalingRuleRequestStepAdjustments {
	return s.StepAdjustments
}

func (s *ModifyScalingRuleRequest) GetTargetValue() *float32 {
	return s.TargetValue
}

func (s *ModifyScalingRuleRequest) SetAdjustmentType(v string) *ModifyScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAdjustmentValue(v int32) *ModifyScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAlarmDimensions(v []*ModifyScalingRuleRequestAlarmDimensions) *ModifyScalingRuleRequest {
	s.AlarmDimensions = v
	return s
}

func (s *ModifyScalingRuleRequest) SetAlarmOptions(v *ModifyScalingRuleRequestAlarmOptions) *ModifyScalingRuleRequest {
	s.AlarmOptions = v
	return s
}

func (s *ModifyScalingRuleRequest) SetCooldown(v int32) *ModifyScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetDisableScaleIn(v bool) *ModifyScalingRuleRequest {
	s.DisableScaleIn = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetEstimatedInstanceWarmup(v int32) *ModifyScalingRuleRequest {
	s.EstimatedInstanceWarmup = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetHybridMetrics(v []*ModifyScalingRuleRequestHybridMetrics) *ModifyScalingRuleRequest {
	s.HybridMetrics = v
	return s
}

func (s *ModifyScalingRuleRequest) SetHybridMonitorNamespace(v string) *ModifyScalingRuleRequest {
	s.HybridMonitorNamespace = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetInitialMaxSize(v int32) *ModifyScalingRuleRequest {
	s.InitialMaxSize = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMetricName(v string) *ModifyScalingRuleRequest {
	s.MetricName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMetricType(v string) *ModifyScalingRuleRequest {
	s.MetricType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetMinAdjustmentMagnitude(v int32) *ModifyScalingRuleRequest {
	s.MinAdjustmentMagnitude = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOwnerId(v int64) *ModifyScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveScalingMode(v string) *ModifyScalingRuleRequest {
	s.PredictiveScalingMode = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveTaskBufferTime(v int32) *ModifyScalingRuleRequest {
	s.PredictiveTaskBufferTime = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveValueBehavior(v string) *ModifyScalingRuleRequest {
	s.PredictiveValueBehavior = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetPredictiveValueBuffer(v int32) *ModifyScalingRuleRequest {
	s.PredictiveValueBuffer = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceOwnerId(v int64) *ModifyScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScaleInEvaluationCount(v int32) *ModifyScalingRuleRequest {
	s.ScaleInEvaluationCount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScaleOutEvaluationCount(v int32) *ModifyScalingRuleRequest {
	s.ScaleOutEvaluationCount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleId(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleName(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetStepAdjustments(v []*ModifyScalingRuleRequestStepAdjustments) *ModifyScalingRuleRequest {
	s.StepAdjustments = v
	return s
}

func (s *ModifyScalingRuleRequest) SetTargetValue(v float32) *ModifyScalingRuleRequest {
	s.TargetValue = &v
	return s
}

func (s *ModifyScalingRuleRequest) Validate() error {
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

type ModifyScalingRuleRequestAlarmDimensions struct {
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

func (s ModifyScalingRuleRequestAlarmDimensions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequestAlarmDimensions) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestAlarmDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *ModifyScalingRuleRequestAlarmDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *ModifyScalingRuleRequestAlarmDimensions) SetDimensionKey(v string) *ModifyScalingRuleRequestAlarmDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ModifyScalingRuleRequestAlarmDimensions) SetDimensionValue(v string) *ModifyScalingRuleRequestAlarmDimensions {
	s.DimensionValue = &v
	return s
}

func (s *ModifyScalingRuleRequestAlarmDimensions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingRuleRequestAlarmOptions struct {
	// The statistical period of the metric data in the target tracking scaling rule. Unit: seconds. Valid values:
	//
	// 	- 15
	//
	// 	- 60
	//
	// 	- 120
	//
	// 	- 300
	//
	// 	- 900
	//
	// >  Default value: 60.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s ModifyScalingRuleRequestAlarmOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequestAlarmOptions) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestAlarmOptions) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyScalingRuleRequestAlarmOptions) SetPeriod(v int32) *ModifyScalingRuleRequestAlarmOptions {
	s.Period = &v
	return s
}

func (s *ModifyScalingRuleRequestAlarmOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingRuleRequestHybridMetrics struct {
	// The metric dimensions. You can use this parameter to specify the monitored resources.
	Dimensions []*ModifyScalingRuleRequestHybridMetricsDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
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

func (s ModifyScalingRuleRequestHybridMetrics) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequestHybridMetrics) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestHybridMetrics) GetDimensions() []*ModifyScalingRuleRequestHybridMetricsDimensions {
	return s.Dimensions
}

func (s *ModifyScalingRuleRequestHybridMetrics) GetExpression() *string {
	return s.Expression
}

func (s *ModifyScalingRuleRequestHybridMetrics) GetId() *string {
	return s.Id
}

func (s *ModifyScalingRuleRequestHybridMetrics) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyScalingRuleRequestHybridMetrics) GetStatistic() *string {
	return s.Statistic
}

func (s *ModifyScalingRuleRequestHybridMetrics) SetDimensions(v []*ModifyScalingRuleRequestHybridMetricsDimensions) *ModifyScalingRuleRequestHybridMetrics {
	s.Dimensions = v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetrics) SetExpression(v string) *ModifyScalingRuleRequestHybridMetrics {
	s.Expression = &v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetrics) SetId(v string) *ModifyScalingRuleRequestHybridMetrics {
	s.Id = &v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetrics) SetMetricName(v string) *ModifyScalingRuleRequestHybridMetrics {
	s.MetricName = &v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetrics) SetStatistic(v string) *ModifyScalingRuleRequestHybridMetrics {
	s.Statistic = &v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetrics) Validate() error {
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

type ModifyScalingRuleRequestHybridMetricsDimensions struct {
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

func (s ModifyScalingRuleRequestHybridMetricsDimensions) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequestHybridMetricsDimensions) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestHybridMetricsDimensions) GetDimensionKey() *string {
	return s.DimensionKey
}

func (s *ModifyScalingRuleRequestHybridMetricsDimensions) GetDimensionValue() *string {
	return s.DimensionValue
}

func (s *ModifyScalingRuleRequestHybridMetricsDimensions) SetDimensionKey(v string) *ModifyScalingRuleRequestHybridMetricsDimensions {
	s.DimensionKey = &v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetricsDimensions) SetDimensionValue(v string) *ModifyScalingRuleRequestHybridMetricsDimensions {
	s.DimensionValue = &v
	return s
}

func (s *ModifyScalingRuleRequestHybridMetricsDimensions) Validate() error {
	return dara.Validate(s)
}

type ModifyScalingRuleRequestStepAdjustments struct {
	// The lower limit that is specified in a step adjustment. This parameter is available only if you set the ScalingRuleType parameter to StepScalingRule. Valid values: -9.999999E18 to 9.999999E18.
	//
	// example:
	//
	// 1.0
	MetricIntervalLowerBound *float32 `json:"MetricIntervalLowerBound,omitempty" xml:"MetricIntervalLowerBound,omitempty"`
	// The upper limit specified in a step adjustment. This parameter is available only if you set the ScalingRuleType parameter to StepScalingRule. Valid values: -9.999999E18 to 9.999999E18.
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

func (s ModifyScalingRuleRequestStepAdjustments) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequestStepAdjustments) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequestStepAdjustments) GetMetricIntervalLowerBound() *float32 {
	return s.MetricIntervalLowerBound
}

func (s *ModifyScalingRuleRequestStepAdjustments) GetMetricIntervalUpperBound() *float32 {
	return s.MetricIntervalUpperBound
}

func (s *ModifyScalingRuleRequestStepAdjustments) GetScalingAdjustment() *int32 {
	return s.ScalingAdjustment
}

func (s *ModifyScalingRuleRequestStepAdjustments) SetMetricIntervalLowerBound(v float32) *ModifyScalingRuleRequestStepAdjustments {
	s.MetricIntervalLowerBound = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustments) SetMetricIntervalUpperBound(v float32) *ModifyScalingRuleRequestStepAdjustments {
	s.MetricIntervalUpperBound = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustments) SetScalingAdjustment(v int32) *ModifyScalingRuleRequestStepAdjustments {
	s.ScalingAdjustment = &v
	return s
}

func (s *ModifyScalingRuleRequestStepAdjustments) Validate() error {
	return dara.Validate(s)
}
