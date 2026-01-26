// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListPrometheusAlertRulesResponseBody
	GetCode() *int64
	SetMessage(v string) *ListPrometheusAlertRulesResponseBody
	GetMessage() *string
	SetPrometheusAlertRules(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) *ListPrometheusAlertRulesResponseBody
	GetPrometheusAlertRules() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules
	SetRequestId(v string) *ListPrometheusAlertRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPrometheusAlertRulesResponseBody
	GetSuccess() *bool
}

type ListPrometheusAlertRulesResponseBody struct {
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned struct.
	PrometheusAlertRules []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules `json:"PrometheusAlertRules,omitempty" xml:"PrometheusAlertRules,omitempty" type:"Repeated"`
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
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListPrometheusAlertRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusAlertRulesResponseBody) GetPrometheusAlertRules() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	return s.PrometheusAlertRules
}

func (s *ListPrometheusAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusAlertRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPrometheusAlertRulesResponseBody) SetCode(v int64) *ListPrometheusAlertRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetMessage(v string) *ListPrometheusAlertRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetPrometheusAlertRules(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) *ListPrometheusAlertRulesResponseBody {
	s.PrometheusAlertRules = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetRequestId(v string) *ListPrometheusAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetSuccess(v bool) *ListPrometheusAlertRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) Validate() error {
	if s.PrometheusAlertRules != nil {
		for _, item := range s.PrometheusAlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRules struct {
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
	Annotations []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// The ID of the cluster.
	//
	// example:
	//
	// c0bad479465464e1d8c1e641b0afb****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the notification policy. This parameter is returned if the NotifyType parameter is set to `DISPATCH_RULE`.
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
	Labels []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The alert message. Tags can be referenced in the {{$labels.xxx}} format.
	//
	// example:
	//
	// The CPU utilization of ${{$labels.pod_name}} exceeds 80%. Current value: {{$value}}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The method that is used to send alert notifications. Valid values:
	//
	// - ALERT_MANAGER: Alert notifications are sent by Operation Center.
	//
	// - DISPATCH_RULE: Alert notifications are
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// Indicates whether the alert rule is enabled. Valid values:
	//
	// - 1: The alert rule is enabled.
	//
	// - 0: The alert rule is disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the alert rule.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetAlertId() *int64 {
	return s.AlertId
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetAlertName() *string {
	return s.AlertName
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetAnnotations() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	return s.Annotations
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetDispatchRuleId() *int64 {
	return s.DispatchRuleId
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetDuration() *string {
	return s.Duration
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetExpression() *string {
	return s.Expression
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetLabels() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	return s.Labels
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetNotifyType() *string {
	return s.NotifyType
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetStatus() *int32 {
	return s.Status
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetTags() []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags {
	return s.Tags
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GetType() *string {
	return s.Type
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAnnotations(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetClusterId(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDispatchRuleId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.DispatchRuleId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDuration(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetExpression(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetLabels(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetMessage(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Message = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetNotifyType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.NotifyType = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetStatus(v int32) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetTags(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Tags = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) Validate() error {
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

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations struct {
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
	// The CPU utilization of ${{$labels.pod_name}} exceeds 80%. Current value: {{$value}}%
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels struct {
	// The name of the tag.
	//
	// example:
	//
	// severity
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the tag associated with the instance.
	//
	// example:
	//
	// critical
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GetName() *string {
	return s.Name
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) Validate() error {
	return dara.Validate(s)
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) GetKey() *string {
	return s.Key
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) GetValue() *string {
	return s.Value
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) SetKey(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags {
	s.Key = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags {
	s.Value = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesTags) Validate() error {
	return dara.Validate(s)
}
