// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomMetricRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *PutCustomMetricRuleRequest
	GetComparisonOperator() *string
	SetContactGroups(v string) *PutCustomMetricRuleRequest
	GetContactGroups() *string
	SetEffectiveInterval(v string) *PutCustomMetricRuleRequest
	GetEffectiveInterval() *string
	SetEmailSubject(v string) *PutCustomMetricRuleRequest
	GetEmailSubject() *string
	SetEvaluationCount(v int32) *PutCustomMetricRuleRequest
	GetEvaluationCount() *int32
	SetGroupId(v string) *PutCustomMetricRuleRequest
	GetGroupId() *string
	SetLevel(v string) *PutCustomMetricRuleRequest
	GetLevel() *string
	SetMetricName(v string) *PutCustomMetricRuleRequest
	GetMetricName() *string
	SetPeriod(v string) *PutCustomMetricRuleRequest
	GetPeriod() *string
	SetResources(v string) *PutCustomMetricRuleRequest
	GetResources() *string
	SetRuleId(v string) *PutCustomMetricRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *PutCustomMetricRuleRequest
	GetRuleName() *string
	SetSilenceTime(v int32) *PutCustomMetricRuleRequest
	GetSilenceTime() *int32
	SetStatistics(v string) *PutCustomMetricRuleRequest
	GetStatistics() *string
	SetThreshold(v string) *PutCustomMetricRuleRequest
	GetThreshold() *string
	SetWebhook(v string) *PutCustomMetricRuleRequest
	GetWebhook() *string
}

type PutCustomMetricRuleRequest struct {
	// The operator that is used to compare the metric value with the threshold. Valid values:
	//
	// 	- `>=`
	//
	// 	- `=`
	//
	// 	- `<=`
	//
	// 	- `>`
	//
	// 	- `<`
	//
	// 	- `!=`
	//
	// This parameter is required.
	//
	// example:
	//
	// >=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The alert contact groups. Separate multiple alert contact groups with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The period of time during which the alert rule is effective. Valid values: 00:00 to 23:59.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert notification email.
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an alert is triggered.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	EvaluationCount *int32 `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The ID of the application group to which the custom monitoring data belongs.
	//
	// >  The value 0 indicates that the reported custom monitoring data does not belong to an application group.
	//
	// example:
	//
	// 7378****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The alert level. Valid values:
	//
	// 	- CRITICAL
	//
	// 	- WARN
	//
	// 	- INFO
	//
	// This parameter is required.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The metric name.
	//
	// >  For more information about how to obtain the metric name, see [DescribeCustomMetricList](https://help.aliyun.com/document_detail/115005.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The cycle that is used to aggregate custom monitoring data. Unit: seconds Set the value to an integral multiple of 60. The original reporting cycle of custom monitoring data is used by default.
	//
	// example:
	//
	// 300
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The custom monitoring data to which the alert rule applies. The value includes the application group ID to which the custom monitoring data belongs and the dimension to which the metric belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"groupId":7378****,"dimension":"instanceId=i-hp3543t5e4sudb3s****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// >  You can specify an existing ID to modify the corresponding alert rule or specify a new ID to create an alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyRuleId1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// CpuUsage
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period during which new alert notifications are not sent even if the trigger conditions are met. Unit: seconds. Default value: 86400, which is equivalent to one day.
	//
	// >  Only one alert notification is sent during each mute period even if the metric value exceeds the alert threshold several times.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The method used to calculate the metric value based on which alerts are triggered.
	//
	// This parameter is required.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutCustomMetricRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRuleRequest) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutCustomMetricRuleRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutCustomMetricRuleRequest) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutCustomMetricRuleRequest) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutCustomMetricRuleRequest) GetEvaluationCount() *int32 {
	return s.EvaluationCount
}

func (s *PutCustomMetricRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PutCustomMetricRuleRequest) GetLevel() *string {
	return s.Level
}

func (s *PutCustomMetricRuleRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *PutCustomMetricRuleRequest) GetPeriod() *string {
	return s.Period
}

func (s *PutCustomMetricRuleRequest) GetResources() *string {
	return s.Resources
}

func (s *PutCustomMetricRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutCustomMetricRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutCustomMetricRuleRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *PutCustomMetricRuleRequest) GetStatistics() *string {
	return s.Statistics
}

func (s *PutCustomMetricRuleRequest) GetThreshold() *string {
	return s.Threshold
}

func (s *PutCustomMetricRuleRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *PutCustomMetricRuleRequest) SetComparisonOperator(v string) *PutCustomMetricRuleRequest {
	s.ComparisonOperator = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetContactGroups(v string) *PutCustomMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetEffectiveInterval(v string) *PutCustomMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetEmailSubject(v string) *PutCustomMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetEvaluationCount(v int32) *PutCustomMetricRuleRequest {
	s.EvaluationCount = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetGroupId(v string) *PutCustomMetricRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetLevel(v string) *PutCustomMetricRuleRequest {
	s.Level = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetMetricName(v string) *PutCustomMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetPeriod(v string) *PutCustomMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetResources(v string) *PutCustomMetricRuleRequest {
	s.Resources = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetRuleId(v string) *PutCustomMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetRuleName(v string) *PutCustomMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetSilenceTime(v int32) *PutCustomMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetStatistics(v string) *PutCustomMetricRuleRequest {
	s.Statistics = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetThreshold(v string) *PutCustomMetricRuleRequest {
	s.Threshold = &v
	return s
}

func (s *PutCustomMetricRuleRequest) SetWebhook(v string) *PutCustomMetricRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutCustomMetricRuleRequest) Validate() error {
	return dara.Validate(s)
}
