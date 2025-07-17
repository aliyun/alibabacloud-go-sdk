// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribePrometheusAlertRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *DescribePrometheusAlertRuleResponseBody
	GetMessage() *string
	SetPrometheusAlertRule(v *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) *DescribePrometheusAlertRuleResponseBody
	GetPrometheusAlertRule() *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule
	SetRequestId(v string) *DescribePrometheusAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePrometheusAlertRuleResponseBody
	GetSuccess() *bool
}

type DescribePrometheusAlertRuleResponseBody struct {
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
	PrometheusAlertRule *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9FEA6D00-317F-45E3-9004-7FB8B0B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribePrometheusAlertRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePrometheusAlertRuleResponseBody) GetPrometheusAlertRule() *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	return s.PrometheusAlertRule
}

func (s *DescribePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrometheusAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePrometheusAlertRuleResponseBody) SetCode(v int64) *DescribePrometheusAlertRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetMessage(v string) *DescribePrometheusAlertRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) *DescribePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetRequestId(v string) *DescribePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetSuccess(v bool) *DescribePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
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
	Annotations []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// The ID of the cluster.
	//
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the notification policy. This parameter is returned if the **NotifyType*	- parameter is set to `DISPATCH_RULE`.
	//
	// example:
	//
	// 10282
	DispatchRuleId *int64 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// The duration of the alert. Valid values: 1 to 1440. Unit: minutes.
	//
	// example:
	//
	// 1m
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The expression of the alert rule.
	//
	// example:
	//
	// 100 	- (sum(rate(container_cpu_usage_seconds_total[1m])) by (pod_name) / sum(label_replace(kube_pod_container_resource_limits_cpu_cores, \\"pod_name\\", \\"$1\\", \\"pod\\", \\"(.*)\\")) by (pod_name))>75
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The tags of the alert rule.
	Labels []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The alert message. Tags can be referenced in the {{$labels.xxx}} format.
	//
	// example:
	//
	// The CPU utilization of ${{$labels.pod_name}} has exceeded 80%. Current value: {{$value}}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The method of that is used to send alert notifications. Valid values:
	//
	// 	- `ALERT_MANAGER`: Alert notifications are sent by Operation Center.
	//
	// 	- `DISPATCH_RULE`: Alert notifications are sent based on the specified notification policy.
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// Indicates whether the alert rule is enabled. Valid values:
	//
	// 	- `1`: The alert rule is enabled.
	//
	// 	- `0`: The alert rule is disabled.
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

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAnnotations() []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	return s.Annotations
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetExpression() *string {
	return s.Expression
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetLabels() []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	return s.Labels
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetMessage() *string {
	return s.Message
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetNotifyType() *string {
	return s.NotifyType
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetType() *string {
	return s.Type
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) Validate() error {
	return dara.Validate(s)
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
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

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
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

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}
