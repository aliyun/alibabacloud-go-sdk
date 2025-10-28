// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationScalingRuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAppScalingRule(v *EnableApplicationScalingRuleResponseBodyAppScalingRule) *EnableApplicationScalingRuleResponseBody
  GetAppScalingRule() *EnableApplicationScalingRuleResponseBodyAppScalingRule 
  SetCode(v int32) *EnableApplicationScalingRuleResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnableApplicationScalingRuleResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableApplicationScalingRuleResponseBody
  GetRequestId() *string 
}

type EnableApplicationScalingRuleResponseBody struct {
  // The information about the auto scaling policy.
  AppScalingRule *EnableApplicationScalingRuleResponseBodyAppScalingRule `json:"AppScalingRule,omitempty" xml:"AppScalingRule,omitempty" type:"Struct"`
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
  // a5281053-08e4-47a5-b2ab-5c0323de7b5a
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationScalingRuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBody) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBody) GetAppScalingRule() *EnableApplicationScalingRuleResponseBodyAppScalingRule  {
  return s.AppScalingRule
}

func (s *EnableApplicationScalingRuleResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableApplicationScalingRuleResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableApplicationScalingRuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableApplicationScalingRuleResponseBody) SetAppScalingRule(v *EnableApplicationScalingRuleResponseBodyAppScalingRule) *EnableApplicationScalingRuleResponseBody {
  s.AppScalingRule = v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetCode(v int32) *EnableApplicationScalingRuleResponseBody {
  s.Code = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetMessage(v string) *EnableApplicationScalingRuleResponseBody {
  s.Message = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) SetRequestId(v string) *EnableApplicationScalingRuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBody) Validate() error {
  if s.AppScalingRule != nil {
    if err := s.AppScalingRule.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableApplicationScalingRuleResponseBodyAppScalingRule struct {
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
  Metric *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
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
  Trigger *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
  // The time when the auto scaling policy was last modified.
  // 
  // example:
  // 
  // 23212323123
  UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRule) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRule) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetAppId() *string  {
  return s.AppId
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetCreateTime() *int64  {
  return s.CreateTime
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetLastDisableTime() *int64  {
  return s.LastDisableTime
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetMaxReplicas() *int32  {
  return s.MaxReplicas
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetMetric() *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric  {
  return s.Metric
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetMinReplicas() *int32  {
  return s.MinReplicas
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleEnabled() *bool  {
  return s.ScaleRuleEnabled
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleName() *string  {
  return s.ScaleRuleName
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetScaleRuleType() *string  {
  return s.ScaleRuleType
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetTrigger() *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger  {
  return s.Trigger
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) GetUpdateTime() *int64  {
  return s.UpdateTime
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetAppId(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.AppId = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetCreateTime(v int64) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.CreateTime = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetLastDisableTime(v int64) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.LastDisableTime = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetMaxReplicas(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.MaxReplicas = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetMetric(v *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.Metric = v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetMinReplicas(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.MinReplicas = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleEnabled(v bool) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.ScaleRuleEnabled = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleName(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.ScaleRuleName = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetScaleRuleType(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.ScaleRuleType = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetTrigger(v *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.Trigger = v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) SetUpdateTime(v int64) *EnableApplicationScalingRuleResponseBodyAppScalingRule {
  s.UpdateTime = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRule) Validate() error {
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

type EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric struct {
  // This parameter is deprecated.
  // 
  // example:
  // 
  // 1
  MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
  // This parameter is deprecated.
  Metrics []*EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
  // This parameter is deprecated.
  // 
  // example:
  // 
  // 1
  MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMaxReplicas() *int32  {
  return s.MaxReplicas
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMetrics() []*EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics  {
  return s.Metrics
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) GetMinReplicas() *int32  {
  return s.MinReplicas
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMaxReplicas(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
  s.MaxReplicas = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMetrics(v []*EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
  s.Metrics = v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) SetMinReplicas(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric {
  s.MinReplicas = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetric) Validate() error {
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

type EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics struct {
  // This parameter is deprecated.
  // 
  // example:
  // 
  // 12
  MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
  // This parameter is deprecated.
  // 
  // example:
  // 
  // cpu
  MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricTargetAverageUtilization() *int32  {
  return s.MetricTargetAverageUtilization
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) GetMetricType() *string  {
  return s.MetricType
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricTargetAverageUtilization(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
  s.MetricTargetAverageUtilization = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) SetMetricType(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics {
  s.MetricType = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleMetricMetrics) Validate() error {
  return dara.Validate(s)
}

type EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger struct {
  // The maximum number of replicas. The upper limit is 1000.
  // 
  // example:
  // 
  // 122
  MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
  // The minimum number of replicas. The lower limit is 0.
  // 
  // example:
  // 
  // 1
  MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
  // The list of triggers.
  Triggers []*EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMaxReplicas() *int32  {
  return s.MaxReplicas
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetMinReplicas() *int32  {
  return s.MinReplicas
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) GetTriggers() []*EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers  {
  return s.Triggers
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMaxReplicas(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
  s.MaxReplicas = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetMinReplicas(v int32) *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
  s.MinReplicas = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) SetTriggers(v []*EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger {
  s.Triggers = v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTrigger) Validate() error {
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

type EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers struct {
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

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetMetaData() *string  {
  return s.MetaData
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetName() *string  {
  return s.Name
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) GetType() *string  {
  return s.Type
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetMetaData(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
  s.MetaData = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetName(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
  s.Name = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) SetType(v string) *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers {
  s.Type = &v
  return s
}

func (s *EnableApplicationScalingRuleResponseBodyAppScalingRuleTriggerTriggers) Validate() error {
  return dara.Validate(s)
}

