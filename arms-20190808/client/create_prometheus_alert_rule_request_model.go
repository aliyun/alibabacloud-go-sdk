// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *CreatePrometheusAlertRuleRequest
	GetAlertName() *string
	SetAnnotations(v string) *CreatePrometheusAlertRuleRequest
	GetAnnotations() *string
	SetClusterId(v string) *CreatePrometheusAlertRuleRequest
	GetClusterId() *string
	SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleRequest
	GetDispatchRuleId() *int64
	SetDuration(v string) *CreatePrometheusAlertRuleRequest
	GetDuration() *string
	SetExpression(v string) *CreatePrometheusAlertRuleRequest
	GetExpression() *string
	SetLabels(v string) *CreatePrometheusAlertRuleRequest
	GetLabels() *string
	SetMessage(v string) *CreatePrometheusAlertRuleRequest
	GetMessage() *string
	SetNotifyType(v string) *CreatePrometheusAlertRuleRequest
	GetNotifyType() *string
	SetRegionId(v string) *CreatePrometheusAlertRuleRequest
	GetRegionId() *string
	SetTags(v []*CreatePrometheusAlertRuleRequestTags) *CreatePrometheusAlertRuleRequest
	GetTags() []*CreatePrometheusAlertRuleRequestTags
	SetType(v string) *CreatePrometheusAlertRuleRequest
	GetType() *string
}

type CreatePrometheusAlertRuleRequest struct {
	// The name of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Prometheus_Alert
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The annotations that are described in a JSON string. You must specify the name and value of each annotation.
	//
	// example:
	//
	// [{"Value": "xxx","Name": "description"}]
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the notification policy. This parameter is required if the NotifyType parameter is set to `DISPATCH_RULE`.
	//
	// example:
	//
	// 10282
	DispatchRuleId *int64 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	// The duration. The value ranges from 1 to 1440 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10m
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The expression of the alert rule. The expression must follow the PromQL syntax.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100 	- (sum(rate(container_cpu_usage_seconds_total[1m])) by (pod_name) / sum(label_replace(kube_pod_container_resource_limits_cpu_cores, \\"pod_name\\", \\"$1\\", \\"pod\\", \\"(.*)\\")) by (pod_name))>75
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The tags that are described in a JSON string. You must specify the name and value of each tag.
	//
	// example:
	//
	// [{"Value": "critical","Name": "severity"}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The content of the alert notification. Tags can be referenced in the {{$labels.xxx}} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// The CPU utilization of ${{$labels.pod_name}} has exceeded 80%. Current value: {{$value}}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The method that is used to send alert notifications. Valid values:
	//
	// - `ALERT_MANAGER`: Alert notifications are sent by Operation Center. This is the default value.
	//
	// - `DISPATCH_RULE`: Alert notifications are sent based on the specified notification policy.
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	Tags []*CreatePrometheusAlertRuleRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the alert rule. Valid values:
	//
	// - 99: custom alert
	//
	// - 101: Prometheus Service alert
	//
	// example:
	//
	// 101
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *CreatePrometheusAlertRuleRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *CreatePrometheusAlertRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePrometheusAlertRuleRequest) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *CreatePrometheusAlertRuleRequest) GetDuration() *string {
	return s.Duration
}

func (s *CreatePrometheusAlertRuleRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreatePrometheusAlertRuleRequest) GetLabels() *string {
	return s.Labels
}

func (s *CreatePrometheusAlertRuleRequest) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusAlertRuleRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *CreatePrometheusAlertRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrometheusAlertRuleRequest) GetTags() []*CreatePrometheusAlertRuleRequestTags {
	return s.Tags
}

func (s *CreatePrometheusAlertRuleRequest) GetType() *string {
	return s.Type
}

func (s *CreatePrometheusAlertRuleRequest) SetAlertName(v string) *CreatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetAnnotations(v string) *CreatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetClusterId(v string) *CreatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDuration(v string) *CreatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetExpression(v string) *CreatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetLabels(v string) *CreatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetMessage(v string) *CreatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetNotifyType(v string) *CreatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetRegionId(v string) *CreatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetTags(v []*CreatePrometheusAlertRuleRequestTags) *CreatePrometheusAlertRuleRequest {
	s.Tags = v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetType(v string) *CreatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePrometheusAlertRuleRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusAlertRuleRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreatePrometheusAlertRuleRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusAlertRuleRequestTags) SetKey(v string) *CreatePrometheusAlertRuleRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequestTags) SetValue(v string) *CreatePrometheusAlertRuleRequestTags {
	s.Value = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequestTags) Validate() error {
	return dara.Validate(s)
}
