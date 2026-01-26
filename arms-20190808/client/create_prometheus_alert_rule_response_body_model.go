// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreatePrometheusAlertRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *CreatePrometheusAlertRuleResponseBody
	GetMessage() *string
	SetPrometheusAlertRule(v *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *CreatePrometheusAlertRuleResponseBody
	GetPrometheusAlertRule() *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule
	SetRequestId(v string) *CreatePrometheusAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePrometheusAlertRuleResponseBody
	GetSuccess() *bool
}

type CreatePrometheusAlertRuleResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned struct.
	PrometheusAlertRule *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9FEA6D00-317F-45E3-9004-7FB8B0B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreatePrometheusAlertRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusAlertRuleResponseBody) GetPrometheusAlertRule() *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	return s.PrometheusAlertRule
}

func (s *CreatePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePrometheusAlertRuleResponseBody) SetCode(v int64) *CreatePrometheusAlertRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetMessage(v string) *CreatePrometheusAlertRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *CreatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetRequestId(v string) *CreatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetSuccess(v bool) *CreatePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) Validate() error {
	if s.PrometheusAlertRule != nil {
		if err := s.PrometheusAlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	// The ID of the alert rule.
	//
	// example:
	//
	// 3888704
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// Prometheus_Alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The annotations of the alert rule.
	Annotations []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// The ID of the cluster.
	//
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// 10282
	DispatchRuleId *int64 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// The duration. The value ranges from 1 to 1440 minutes.
	//
	// example:
	//
	// 10m
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The expression of the alert rule.
	//
	// example:
	//
	// 100 	- (sum(rate(container_cpu_usage_seconds_total[1m])) by (pod_name) / sum(label_replace(kube_pod_container_resource_limits_cpu_cores, \\"pod_name\\", \\"$1\\", \\"pod\\", \\"(.*)\\")) by (pod_name))>75
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The tags of the alert rule.
	Labels []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The content of the alert notification. Tags can be referenced in the {{$labels.xxx}} format.
	//
	// example:
	//
	// The CPU utilization of ${{$labels.pod_name}} has exceeded 80%. Current value: {{$value}}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The method that is used to send alert notifications. Valid values:
	//
	// - ALERT_MANAGER: Alert notifications are sent by Operation Center.
	//
	// - DISPATCH_RULE: Alert notifications are sent based on the specified notification policy.
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// Indicates whether the alert rule is enabled. Valid values:
	//
	// - `1`: The alert rule is enabled.
	//
	// - `0`: The alert rule is disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the alert rule.
	//
	// example:
	//
	// Kubernetes component alert
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAnnotations() []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	return s.Annotations
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetExpression() *string {
	return s.Expression
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetLabels() []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	return s.Labels
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetNotifyType() *string {
	return s.NotifyType
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetStatus() *int32 {
	return s.Status
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetType() *string {
	return s.Type
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) Validate() error {
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	// The name of the annotation.
	//
	// example:
	//
	// message
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the annotation.
	//
	// example:
	//
	// The CPU utilization of ${{$labels.pod_name}} has exceeded 80%. Current value: {{$value}}%
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	// The name of the tag.
	//
	// example:
	//
	// severity
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// critical
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}
