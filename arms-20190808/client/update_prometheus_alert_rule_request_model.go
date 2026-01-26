// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *UpdatePrometheusAlertRuleRequest
	GetAlertId() *int64
	SetAlertName(v string) *UpdatePrometheusAlertRuleRequest
	GetAlertName() *string
	SetAnnotations(v string) *UpdatePrometheusAlertRuleRequest
	GetAnnotations() *string
	SetClusterId(v string) *UpdatePrometheusAlertRuleRequest
	GetClusterId() *string
	SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleRequest
	GetDispatchRuleId() *int64
	SetDuration(v string) *UpdatePrometheusAlertRuleRequest
	GetDuration() *string
	SetExpression(v string) *UpdatePrometheusAlertRuleRequest
	GetExpression() *string
	SetLabels(v string) *UpdatePrometheusAlertRuleRequest
	GetLabels() *string
	SetMessage(v string) *UpdatePrometheusAlertRuleRequest
	GetMessage() *string
	SetNotifyType(v string) *UpdatePrometheusAlertRuleRequest
	GetNotifyType() *string
	SetRegionId(v string) *UpdatePrometheusAlertRuleRequest
	GetRegionId() *string
	SetTags(v []*UpdatePrometheusAlertRuleRequestTags) *UpdatePrometheusAlertRuleRequest
	GetTags() []*UpdatePrometheusAlertRuleRequestTags
	SetType(v string) *UpdatePrometheusAlertRuleRequest
	GetType() *string
}

type UpdatePrometheusAlertRuleRequest struct {
	// The ID of the alert rule. You can call the ListPrometheusAlertRules operation to query the ID of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3888704
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
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
	// The cluster ID.
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
	// 1
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
	// The alert message. Tags can be referenced in the {{$labels.xxx}} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// The CPU utilization of ${{$labels.pod_name}} exceeds 80%. Current value: {{$value}}%
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	Tags []*UpdatePrometheusAlertRuleRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the alert rule.
	//
	// example:
	//
	// Kubernetes component alert
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *UpdatePrometheusAlertRuleRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *UpdatePrometheusAlertRuleRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *UpdatePrometheusAlertRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusAlertRuleRequest) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *UpdatePrometheusAlertRuleRequest) GetDuration() *string {
	return s.Duration
}

func (s *UpdatePrometheusAlertRuleRequest) GetExpression() *string {
	return s.Expression
}

func (s *UpdatePrometheusAlertRuleRequest) GetLabels() *string {
	return s.Labels
}

func (s *UpdatePrometheusAlertRuleRequest) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusAlertRuleRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *UpdatePrometheusAlertRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusAlertRuleRequest) GetTags() []*UpdatePrometheusAlertRuleRequestTags {
	return s.Tags
}

func (s *UpdatePrometheusAlertRuleRequest) GetType() *string {
	return s.Type
}

func (s *UpdatePrometheusAlertRuleRequest) SetAlertId(v int64) *UpdatePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetAlertName(v string) *UpdatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetAnnotations(v string) *UpdatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetClusterId(v string) *UpdatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetDuration(v string) *UpdatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetExpression(v string) *UpdatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetLabels(v string) *UpdatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetMessage(v string) *UpdatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetNotifyType(v string) *UpdatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetRegionId(v string) *UpdatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetTags(v []*UpdatePrometheusAlertRuleRequestTags) *UpdatePrometheusAlertRuleRequest {
	s.Tags = v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetType(v string) *UpdatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrometheusAlertRuleRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusAlertRuleRequestTags) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdatePrometheusAlertRuleRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdatePrometheusAlertRuleRequestTags) SetKey(v string) *UpdatePrometheusAlertRuleRequestTags {
	s.Key = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequestTags) SetValue(v string) *UpdatePrometheusAlertRuleRequestTags {
	s.Value = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequestTags) Validate() error {
	return dara.Validate(s)
}
