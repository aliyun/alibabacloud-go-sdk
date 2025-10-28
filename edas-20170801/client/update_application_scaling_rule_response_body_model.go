// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppScalingRule(v *UpdateApplicationScalingRuleResponseBodyAppScalingRule) *UpdateApplicationScalingRuleResponseBody
	GetAppScalingRule() *UpdateApplicationScalingRuleResponseBodyAppScalingRule
	SetCode(v int32) *UpdateApplicationScalingRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationScalingRuleResponseBody
	GetRequestId() *string
}

type UpdateApplicationScalingRuleResponseBody struct {
	// The information about the auto scaling policy.
	AppScalingRule *UpdateApplicationScalingRuleResponseBodyAppScalingRule `json:"AppScalingRule,omitempty" xml:"AppScalingRule,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBody) GetAppScalingRule() *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	return s.AppScalingRule
}

func (s *UpdateApplicationScalingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationScalingRuleResponseBody) SetAppScalingRule(v *UpdateApplicationScalingRuleResponseBodyAppScalingRule) *UpdateApplicationScalingRuleResponseBody {
	s.AppScalingRule = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetCode(v int32) *UpdateApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetMessage(v string) *UpdateApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) SetRequestId(v string) *UpdateApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBody) Validate() error {
	if s.AppScalingRule != nil {
		if err := s.AppScalingRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRule struct {
	// The ID of the application to which the auto scaling policy belongs.
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId     *string                                                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Behaviour *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour `json:"Behaviour,omitempty" xml:"Behaviour,omitempty" type:"Struct"`
	// The time when the auto scaling policy was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601785
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the auto scaling policy was last disabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601785
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metric *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// Indicates whether the auto scaling policy is enabled. Valid values:
	//
	// 	- **true**: The auto scaling policy is enabled.
	//
	// 	- **false**: The auto scaling policy is disabled.
	//
	// example:
	//
	// true
	ScaleRuleEnabled *bool `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The name of the auto scaling policy.
	//
	// example:
	//
	// cpu-trigger
	ScaleRuleName *string `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	// The type of the auto scaling policy. The value is fixed to trigger.
	//
	// example:
	//
	// trigger
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The configurations of the trigger.
	Trigger *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The time when the auto scaling policy was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1574251601785
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRule) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetBehaviour() *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour {
	return s.Behaviour
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetMetric() *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	return s.Metric
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetTrigger() *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	return s.Trigger
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetAppId(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetBehaviour(v *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.Behaviour = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetCreateTime(v int64) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.CreateTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetLastDisableTime(v int64) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.LastDisableTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetMaxReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.MaxReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetMetric(v *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.Metric = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetMinReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.MinReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleEnabled(v bool) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleName(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleName = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleType(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleType = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetTrigger(v *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.Trigger = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) SetUpdateTime(v int64) *UpdateApplicationScalingRuleResponseBodyAppScalingRule {
	s.UpdateTime = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRule) Validate() error {
	if s.Behaviour != nil {
		if err := s.Behaviour.Validate(); err != nil {
			return err
		}
	}
	if s.Metric != nil {
		if err := s.Metric.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour struct {
	ScaleDown *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown `json:"ScaleDown,omitempty" xml:"ScaleDown,omitempty" type:"Struct"`
	ScaleUp   *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp   `json:"ScaleUp,omitempty" xml:"ScaleUp,omitempty" type:"Struct"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) GetScaleDown() *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	return s.ScaleDown
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) GetScaleUp() *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	return s.ScaleUp
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) SetScaleDown(v *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour {
	s.ScaleDown = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) SetScaleUp(v *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour {
	s.ScaleUp = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) Validate() error {
	if s.ScaleDown != nil {
		if err := s.ScaleDown.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleUp != nil {
		if err := s.ScaleUp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown struct {
	// The configuration of the policy.
	Policies []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The step size policy for the scale-in. Valid values: Max, Min, and Disable.
	//
	// example:
	//
	// Max
	SelectPolicy *string `json:"SelectPolicy,omitempty" xml:"SelectPolicy,omitempty"`
	// The cooldown time of the scale-in. Valid values: 0 to 3600. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	StabilizationWindowSeconds *int32 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GetPolicies() []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	return s.Policies
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GetSelectPolicy() *string {
	return s.SelectPolicy
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) SetPolicies(v []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	s.Policies = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) SetSelectPolicy(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	s.SelectPolicy = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) SetStabilizationWindowSeconds(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies struct {
	// The cycle of the policy check. Valid values: 0 to 1800. Unit: seconds.
	//
	// example:
	//
	// 15
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The type of the policy. Valid values: Pods and Percent.
	//
	// example:
	//
	// Pods
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The policy value of auto scaling. Set the value to an integer greater than zero. If the policy type is Pods, the value of this parameter indicates the number of pods. If the policy type is Percent, the value of this parameter indicates a percentage, which can exceed 100%.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GetType() *string {
	return s.Type
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GetValue() *string {
	return s.Value
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) SetPeriodSeconds(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) SetType(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	s.Type = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) SetValue(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	s.Value = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp struct {
	// The configuration of the policy.
	Policies []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The step size policy for the scale-out. Valid values: Max, Min, and Disable.
	//
	// example:
	//
	// Max
	SelectPolicy *string `json:"SelectPolicy,omitempty" xml:"SelectPolicy,omitempty"`
	// The cooldown time of the scale-out. Valid values: 0 to 3600. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 0
	StabilizationWindowSeconds *int32 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GetPolicies() []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	return s.Policies
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GetSelectPolicy() *string {
	return s.SelectPolicy
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) SetPolicies(v []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	s.Policies = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) SetSelectPolicy(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	s.SelectPolicy = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) SetStabilizationWindowSeconds(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies struct {
	// The cycle of the policy check. Valid values: 0 to 1800. Unit: seconds.
	//
	// example:
	//
	// 15
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The type of the policy. Valid values: Pods and Percent.
	//
	// example:
	//
	// Pods
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The policy value of auto scaling. Set the value to an integer greater than zero. If the policy type is Pods, the value of this parameter indicates the number of pods. If the policy type is Percent, the value of this parameter indicates a percentage, which can exceed 100%.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GetType() *string {
	return s.Type
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GetValue() *string {
	return s.Value
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) SetPeriodSeconds(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	s.PeriodSeconds = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) SetType(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	s.Type = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) SetValue(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	s.Value = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metrics []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMetrics() []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	return s.Metrics
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMaxReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.MaxReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMetrics(v []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.Metrics = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMinReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.MinReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetric) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// cpu
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricTargetAverageUtilization(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricType(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger struct {
	// The maximum number of replicas. The maximum value is 1000.
	//
	// example:
	//
	// 122
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of replicas. The minimum value is 0.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The configurations of the trigger.
	Triggers []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetTriggers() []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	return s.Triggers
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMaxReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.MaxReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMinReplicas(v int32) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.MinReplicas = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetTriggers(v []*UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.Triggers = v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) Validate() error {
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers struct {
	// The metadata of the trigger.
	//
	// example:
	//
	// {"dryRun":true}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// The name of the trigger.
	//
	// example:
	//
	// cpu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the trigger. Valid values: cron and app_metric.
	//
	// example:
	//
	// cron
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GoString() string {
	return s.String()
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetMetaData() *string {
	return s.MetaData
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetType() *string {
	return s.Type
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetMetaData(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.MetaData = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetName(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.Name = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetType(v string) *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.Type = &v
	return s
}

func (s *UpdateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) Validate() error {
	return dara.Validate(s)
}
