// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppScalingRule(v *DisableApplicationScalingRuleResponseBodyAppScalingRule) *DisableApplicationScalingRuleResponseBody
	GetAppScalingRule() *DisableApplicationScalingRuleResponseBodyAppScalingRule
	SetCode(v int32) *DisableApplicationScalingRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *DisableApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableApplicationScalingRuleResponseBody
	GetRequestId() *string
}

type DisableApplicationScalingRuleResponseBody struct {
	// The information about the auto scaling policy.
	AppScalingRule *DisableApplicationScalingRuleResponseBodyAppScalingRule `json:"AppScalingRule,omitempty" xml:"AppScalingRule,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5d6fa0bc-cc3**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBody) GetAppScalingRule() *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	return s.AppScalingRule
}

func (s *DisableApplicationScalingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DisableApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableApplicationScalingRuleResponseBody) SetAppScalingRule(v *DisableApplicationScalingRuleResponseBodyAppScalingRule) *DisableApplicationScalingRuleResponseBody {
	s.AppScalingRule = v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetCode(v int32) *DisableApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetMessage(v string) *DisableApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) SetRequestId(v string) *DisableApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBody) Validate() error {
	if s.AppScalingRule != nil {
		if err := s.AppScalingRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableApplicationScalingRuleResponseBodyAppScalingRule struct {
	// The ID of the application to which the auto scaling policy belongs.
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the auto scaling policy was created.
	//
	// example:
	//
	// 23212323123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the auto scaling policy was last disabled.
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
	Metric *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
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
	// cron-trigger
	ScaleRuleName *string `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	// The type of the auto scaling policy. The value is fixed to trigger.
	//
	// example:
	//
	// trigger
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The configurations of the trigger.
	Trigger *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The time when the auto scaling policy was last modified.
	//
	// example:
	//
	// 23212323123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRule) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRule) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetAppId() *string {
	return s.AppId
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetMetric() *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	return s.Metric
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetTrigger() *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	return s.Trigger
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetAppId(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.AppId = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetCreateTime(v int64) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.CreateTime = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetLastDisableTime(v int64) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.LastDisableTime = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetMaxReplicas(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.MaxReplicas = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetMetric(v *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.Metric = v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetMinReplicas(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.MinReplicas = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleEnabled(v bool) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleName(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleName = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleType(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.ScaleRuleType = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetTrigger(v *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.Trigger = v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) SetUpdateTime(v int64) *DisableApplicationScalingRuleResponseBodyAppScalingRule {
	s.UpdateTime = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRule) Validate() error {
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

type DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 12
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metrics []*DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMetrics() []*DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	return s.Metrics
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMaxReplicas(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.MaxReplicas = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMetrics(v []*DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.Metrics = v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMinReplicas(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
	s.MinReplicas = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetric) Validate() error {
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

type DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics struct {
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
	// 1
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricTargetAverageUtilization(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricType(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger struct {
	// The maximum number of replicas. The upper limit is 1000.
	//
	// example:
	//
	// 12
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of replicas. The lower limit is 0.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The information about the trigger.
	Triggers []*DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetTriggers() []*DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	return s.Triggers
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMaxReplicas(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.MaxReplicas = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMinReplicas(v int32) *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.MinReplicas = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetTriggers(v []*DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
	s.Triggers = v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) Validate() error {
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

type DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers struct {
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
	// cron-trigger
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the trigger. Valid values: cron and app_metric.
	//
	// example:
	//
	// cron
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GoString() string {
	return s.String()
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetMetaData() *string {
	return s.MetaData
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetName() *string {
	return s.Name
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetType() *string {
	return s.Type
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetMetaData(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.MetaData = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetName(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.Name = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetType(v string) *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
	s.Type = &v
	return s
}

func (s *DisableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) Validate() error {
	return dara.Validate(s)
}
