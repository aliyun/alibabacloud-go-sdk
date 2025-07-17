// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdatePrometheusAlertRuleResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdatePrometheusAlertRuleResponseBody
	GetMessage() *string
	SetPrometheusAlertRule(v *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *UpdatePrometheusAlertRuleResponseBody
	GetPrometheusAlertRule() *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule
	SetRequestId(v string) *UpdatePrometheusAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePrometheusAlertRuleResponseBody
	GetSuccess() *bool
}

type UpdatePrometheusAlertRuleResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
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
	Message             *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PrometheusAlertRule *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	// example:
	//
	// 9FEA6D00-317F-45E3-9004-7FB8B0B7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetPrometheusAlertRule() *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	return s.PrometheusAlertRule
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetCode(v int64) *UpdatePrometheusAlertRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetMessage(v string) *UpdatePrometheusAlertRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *UpdatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetRequestId(v string) *UpdatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetSuccess(v bool) *UpdatePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	// example:
	//
	// 3888704
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// example:
	//
	// Prometheus_Alert
	AlertName   *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10282
	DispatchRuleId *int64 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 100 	- (sum(rate(container_cpu_usage_seconds_total[1m])) by (pod_name) / sum(label_replace(kube_pod_container_resource_limits_cpu_cores, \\"pod_name\\", \\"$1\\", \\"pod\\", \\"(.*)\\")) by (pod_name))>75
	Expression *string                                                           `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels     []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message    *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ALERT_MANAGER
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// 1
	Status *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetAnnotations() []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	return s.Annotations
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetExpression() *string {
	return s.Expression
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetLabels() []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	return s.Labels
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetNotifyType() *string {
	return s.NotifyType
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetStatus() *int32 {
	return s.Status
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GetType() *string {
	return s.Type
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	// example:
	//
	// message
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	// example:
	//
	// severity
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// critical
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}
