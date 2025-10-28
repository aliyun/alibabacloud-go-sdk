// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppScalingRule(v *CreateApplicationScalingRuleResponseBodyAppScalingRule) *CreateApplicationScalingRuleResponseBody
	GetAppScalingRule() *CreateApplicationScalingRuleResponseBodyAppScalingRule
	SetCode(v int32) *CreateApplicationScalingRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApplicationScalingRuleResponseBody
	GetRequestId() *string
}

type CreateApplicationScalingRuleResponseBody struct {
	// The information about the auto scaling policy.
	AppScalingRule *CreateApplicationScalingRuleResponseBodyAppScalingRule `json:"AppScalingRule,omitempty" xml:"AppScalingRule,omitempty" type:"Struct"`
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

func (s CreateApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBody) GetAppScalingRule() *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	return s.AppScalingRule
}

func (s *CreateApplicationScalingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationScalingRuleResponseBody) SetAppScalingRule(v *CreateApplicationScalingRuleResponseBodyAppScalingRule) *CreateApplicationScalingRuleResponseBody {
	s.AppScalingRule = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetCode(v int32) *CreateApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetMessage(v string) *CreateApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) SetRequestId(v string) *CreateApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBody) Validate() error {
	if s.AppScalingRule != nil {
		if err := s.AppScalingRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationScalingRuleResponseBodyAppScalingRule struct {
	// The ID of the application to which the auto scaling policy belongs.
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The scaling behavior.
	Behaviour *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour `json:"Behaviour,omitempty" xml:"Behaviour,omitempty" type:"Struct"`
	// The timestamp when the auto scaling policy was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 23212323123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp when the auto scaling policy was last disabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 23212323123
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metric *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
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
	// cpu
	ScaleRuleName *string `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	// The type of the rule. Only trigger is supported.
	//
	// example:
	//
	// trigger
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The configurations of the trigger.
	Trigger *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The timestamp when the auto scaling policy was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 23212323123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRule) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRule) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetBehaviour() *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour {
	return s.Behaviour
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetMetric() *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	return s.Metric
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetTrigger() *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	return s.Trigger
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetAppId(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.AppId = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetBehaviour(v *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.Behaviour = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetCreateTime(v int64) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.CreateTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetLastDisableTime(v int64) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.LastDisableTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetMaxReplicas(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.MaxReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetMetric(v *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.Metric = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetMinReplicas(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.MinReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleEnabled(v bool) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleName(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleName = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleType(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleType = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetTrigger(v *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.Trigger = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) SetUpdateTime(v int64) *CreateApplicationScalingRuleResponseBodyAppScalingRule {
	s.UpdateTime = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRule) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour struct {
	// The behavior configurations of the scale-in.
	ScaleDown *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown `json:"ScaleDown,omitempty" xml:"ScaleDown,omitempty" type:"Struct"`
	// The behavior configurations of the scale-out.
	ScaleUp *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp `json:"ScaleUp,omitempty" xml:"ScaleUp,omitempty" type:"Struct"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) GetScaleDown() *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	return s.ScaleDown
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) GetScaleUp() *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	return s.ScaleUp
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) SetScaleDown(v *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour {
	s.ScaleDown = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) SetScaleUp(v *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour {
	s.ScaleUp = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviour) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown struct {
	// The configurations of the auto scaling policy.
	Policies []*CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// Max
	SelectPolicy *string `json:"SelectPolicy,omitempty" xml:"SelectPolicy,omitempty"`
	// example:
	//
	// 300
	StabilizationWindowSeconds *int32 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GetPolicies() []*CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	return s.Policies
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GetSelectPolicy() *string {
	return s.SelectPolicy
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) SetPolicies(v []*CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	s.Policies = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) SetSelectPolicy(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	s.SelectPolicy = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) SetStabilizationWindowSeconds(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDown) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies struct {
	// The period during which the check is performed. Valid values: 0 to 1800. Unit: seconds.
	//
	// example:
	//
	// 15
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The type of the policy. The value can be Pods or Percent.
	//
	// example:
	//
	// Pods
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the auto scaling policy. The value of this parameter is an integer greater than 0. If the policy type is Pods, the value indicates the number of pods. If the policy type is Percent, the value indicates a percentage. The value is allowed to exceed 100%.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GetType() *string {
	return s.Type
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) GetValue() *string {
	return s.Value
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) SetPeriodSeconds(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) SetType(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	s.Type = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) SetValue(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies {
	s.Value = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleDownPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp struct {
	// The configurations of the auto scaling policy.
	Policies []*CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// Max
	SelectPolicy *string `json:"SelectPolicy,omitempty" xml:"SelectPolicy,omitempty"`
	// example:
	//
	// 0
	StabilizationWindowSeconds *int32 `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GetPolicies() []*CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	return s.Policies
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GetSelectPolicy() *string {
	return s.SelectPolicy
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) SetPolicies(v []*CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	s.Policies = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) SetSelectPolicy(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	s.SelectPolicy = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) SetStabilizationWindowSeconds(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUp) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies struct {
	// The period during which the check is performed. Valid values: 0 to 1800. Unit: seconds.
	//
	// example:
	//
	// 15
	PeriodSeconds *int32 `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	// The type of the policy. The value can be Pods or Percent.
	//
	// example:
	//
	// Pods
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the auto scaling policy. The value of this parameter is an integer greater than 0. If the policy type is Pods, the value indicates the number of pods. If the policy type is Percent, the value indicates a percentage. The value is allowed to exceed 100%.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GetType() *string {
	return s.Type
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) GetValue() *string {
	return s.Value
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) SetPeriodSeconds(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	s.PeriodSeconds = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) SetType(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	s.Type = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) SetValue(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies {
	s.Value = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleBehaviourScaleUpPolicies) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metrics []*CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMetrics() []*CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	return s.Metrics
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMaxReplicas(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.MaxReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMetrics(v []*CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.Metrics = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMinReplicas(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.MinReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetric) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics struct {
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

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricTargetAverageUtilization(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricType(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger struct {
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
	// 2
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The triggers.
	Triggers []*CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetTriggers() []*CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	return s.Triggers
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMaxReplicas(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.MaxReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMinReplicas(v int32) *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.MinReplicas = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetTriggers(v []*CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.Triggers = v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTrigger) Validate() error {
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

type CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers struct {
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
	// The type of the trigger. Only cron and app_metric are supported.
	//
	// example:
	//
	// cron
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GoString() string {
	return s.String()
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetMetaData() *string {
	return s.MetaData
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetName() *string {
	return s.Name
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetType() *string {
	return s.Type
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetMetaData(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.MetaData = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetName(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.Name = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetType(v string) *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.Type = &v
	return s
}

func (s *CreateApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) Validate() error {
	return dara.Validate(s)
}
