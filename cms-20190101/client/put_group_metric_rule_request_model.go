// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutGroupMetricRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEscalations(v *PutGroupMetricRuleRequestEscalations) *PutGroupMetricRuleRequest
	GetEscalations() *PutGroupMetricRuleRequestEscalations
	SetCategory(v string) *PutGroupMetricRuleRequest
	GetCategory() *string
	SetContactGroups(v string) *PutGroupMetricRuleRequest
	GetContactGroups() *string
	SetDimensions(v string) *PutGroupMetricRuleRequest
	GetDimensions() *string
	SetEffectiveInterval(v string) *PutGroupMetricRuleRequest
	GetEffectiveInterval() *string
	SetEmailSubject(v string) *PutGroupMetricRuleRequest
	GetEmailSubject() *string
	SetExtraDimensionJson(v string) *PutGroupMetricRuleRequest
	GetExtraDimensionJson() *string
	SetGroupId(v string) *PutGroupMetricRuleRequest
	GetGroupId() *string
	SetInterval(v string) *PutGroupMetricRuleRequest
	GetInterval() *string
	SetLabels(v []*PutGroupMetricRuleRequestLabels) *PutGroupMetricRuleRequest
	GetLabels() []*PutGroupMetricRuleRequestLabels
	SetMetricName(v string) *PutGroupMetricRuleRequest
	GetMetricName() *string
	SetNamespace(v string) *PutGroupMetricRuleRequest
	GetNamespace() *string
	SetNoDataPolicy(v string) *PutGroupMetricRuleRequest
	GetNoDataPolicy() *string
	SetNoEffectiveInterval(v string) *PutGroupMetricRuleRequest
	GetNoEffectiveInterval() *string
	SetOptions(v string) *PutGroupMetricRuleRequest
	GetOptions() *string
	SetPeriod(v string) *PutGroupMetricRuleRequest
	GetPeriod() *string
	SetRuleId(v string) *PutGroupMetricRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *PutGroupMetricRuleRequest
	GetRuleName() *string
	SetSilenceTime(v int32) *PutGroupMetricRuleRequest
	GetSilenceTime() *int32
	SetWebhook(v string) *PutGroupMetricRuleRequest
	GetWebhook() *string
}

type PutGroupMetricRuleRequest struct {
	Escalations *PutGroupMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The abbreviation of the cloud service name.
	//
	// For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// example:
	//
	// ECS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The alert contact group.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The first-level dimension of the alert rule in the application group.
	//
	// Set the value to a set of key-value pairs, for example, `userId:120886317861****` or `instanceId:i-m5e1qg6uo38rztr4****`.
	//
	// example:
	//
	// [{"instanceId":"i-m5e1qg6uo38rztr4****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The time period during which the alert rule is effective.
	//
	// example:
	//
	// 05:31-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert notification email.
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The second-level or third-level dimension of the alert rule in the application group.
	//
	// Set the value to a set of key-value pairs, for example, `port:80` or `/dev/xvda:d-m5e6yphgzn3aprwu****`.
	//
	// If the first-level dimension of the alert rule is `instanceId:i-m5e1qg6uo38rztr4****`, its second-level dimension is the `/dev/xvda:d-m5e6yphgzn3aprwu****` disk in the instance.
	//
	// example:
	//
	// {"/dev/xvda":"d-m5e6yphgzn3aprwu****"}
	ExtraDimensionJson *string `json:"ExtraDimensionJson,omitempty" xml:"ExtraDimensionJson,omitempty"`
	// The application group ID.
	//
	// For more information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 17285****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The interval at which CloudMonitor checks whether the alert rule is triggered. Unit: seconds.
	//
	// >  We recommend that you set the interval to the data aggregation period. If the interval is shorter than the data aggregation period, alerts cannot be triggered due to insufficient data.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The tags of the alert rule.
	//
	// The specified tag is contained in alert notifications.
	Labels []*PutGroupMetricRuleRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name.
	//
	// For more information about how to obtain the name of a metric, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// For more information about how to obtain the namespace of a cloud service, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The method that is used to handle alerts when no monitoring data is found. Valid values:
	//
	// 	- KEEP_LAST_STATE (default): No operation is performed.
	//
	// 	- INSUFFICIENT_DATA: An alert whose content is "Insufficient data" is triggered.
	//
	// 	- OK: The status is considered normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time period during which the alert rule is ineffective.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	Options             *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The aggregation period of the metric data.
	//
	// Set the `Period` parameter to an integral multiple of 60. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the alert rule.
	//
	// 	- When you create an alert rule for the application group, enter the ID of the alert rule.
	//
	// 	- When you modify a specified alert rule in the application group, you must obtain the ID of the alert rule. For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// 	- When you create an alert rule for the application group, enter the name of the alert rule.
	//
	// 	- When you modify a specified alert rule in the application group, you must obtain the name of the alert rule. For more information about how to obtain the name of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Rule_01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period during which new alerts are not sent even if the trigger conditions are met.
	//
	// Unit: seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The callback URL.
	//
	// The callback URL must be accessible over the Internet. CloudMonitor sends a POST request to push an alert notification to the callback URL that you specify. Only HTTP requests are supported.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutGroupMetricRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequest) GetEscalations() *PutGroupMetricRuleRequestEscalations {
	return s.Escalations
}

func (s *PutGroupMetricRuleRequest) GetCategory() *string {
	return s.Category
}

func (s *PutGroupMetricRuleRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutGroupMetricRuleRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *PutGroupMetricRuleRequest) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutGroupMetricRuleRequest) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutGroupMetricRuleRequest) GetExtraDimensionJson() *string {
	return s.ExtraDimensionJson
}

func (s *PutGroupMetricRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PutGroupMetricRuleRequest) GetInterval() *string {
	return s.Interval
}

func (s *PutGroupMetricRuleRequest) GetLabels() []*PutGroupMetricRuleRequestLabels {
	return s.Labels
}

func (s *PutGroupMetricRuleRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *PutGroupMetricRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutGroupMetricRuleRequest) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *PutGroupMetricRuleRequest) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *PutGroupMetricRuleRequest) GetOptions() *string {
	return s.Options
}

func (s *PutGroupMetricRuleRequest) GetPeriod() *string {
	return s.Period
}

func (s *PutGroupMetricRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutGroupMetricRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutGroupMetricRuleRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *PutGroupMetricRuleRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *PutGroupMetricRuleRequest) SetEscalations(v *PutGroupMetricRuleRequestEscalations) *PutGroupMetricRuleRequest {
	s.Escalations = v
	return s
}

func (s *PutGroupMetricRuleRequest) SetCategory(v string) *PutGroupMetricRuleRequest {
	s.Category = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetContactGroups(v string) *PutGroupMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetDimensions(v string) *PutGroupMetricRuleRequest {
	s.Dimensions = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetEffectiveInterval(v string) *PutGroupMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetEmailSubject(v string) *PutGroupMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetExtraDimensionJson(v string) *PutGroupMetricRuleRequest {
	s.ExtraDimensionJson = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetGroupId(v string) *PutGroupMetricRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetInterval(v string) *PutGroupMetricRuleRequest {
	s.Interval = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetLabels(v []*PutGroupMetricRuleRequestLabels) *PutGroupMetricRuleRequest {
	s.Labels = v
	return s
}

func (s *PutGroupMetricRuleRequest) SetMetricName(v string) *PutGroupMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetNamespace(v string) *PutGroupMetricRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetNoDataPolicy(v string) *PutGroupMetricRuleRequest {
	s.NoDataPolicy = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetNoEffectiveInterval(v string) *PutGroupMetricRuleRequest {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetOptions(v string) *PutGroupMetricRuleRequest {
	s.Options = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetPeriod(v string) *PutGroupMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetRuleId(v string) *PutGroupMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetRuleName(v string) *PutGroupMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetSilenceTime(v int32) *PutGroupMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutGroupMetricRuleRequest) SetWebhook(v string) *PutGroupMetricRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutGroupMetricRuleRequest) Validate() error {
	return dara.Validate(s)
}

type PutGroupMetricRuleRequestEscalations struct {
	Critical *PutGroupMetricRuleRequestEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *PutGroupMetricRuleRequestEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *PutGroupMetricRuleRequestEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s PutGroupMetricRuleRequestEscalations) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalations) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalations) GetCritical() *PutGroupMetricRuleRequestEscalationsCritical {
	return s.Critical
}

func (s *PutGroupMetricRuleRequestEscalations) GetInfo() *PutGroupMetricRuleRequestEscalationsInfo {
	return s.Info
}

func (s *PutGroupMetricRuleRequestEscalations) GetWarn() *PutGroupMetricRuleRequestEscalationsWarn {
	return s.Warn
}

func (s *PutGroupMetricRuleRequestEscalations) SetCritical(v *PutGroupMetricRuleRequestEscalationsCritical) *PutGroupMetricRuleRequestEscalations {
	s.Critical = v
	return s
}

func (s *PutGroupMetricRuleRequestEscalations) SetInfo(v *PutGroupMetricRuleRequestEscalationsInfo) *PutGroupMetricRuleRequestEscalations {
	s.Info = v
	return s
}

func (s *PutGroupMetricRuleRequestEscalations) SetWarn(v *PutGroupMetricRuleRequestEscalationsWarn) *PutGroupMetricRuleRequestEscalations {
	s.Warn = v
	return s
}

func (s *PutGroupMetricRuleRequestEscalations) Validate() error {
	return dara.Validate(s)
}

type PutGroupMetricRuleRequestEscalationsCritical struct {
	// The operator that is used to compare the metric value with the threshold for Critical-level alerts. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical methods for Critical-level alerts. Separate multiple statistical methods with commas (,).
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified cloud service. The value of this parameter can be Maximum, Minimum, or Average. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Critical-level alerts.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before a Critical-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutGroupMetricRuleRequestEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetComparisonOperator(v string) *PutGroupMetricRuleRequestEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetStatistics(v string) *PutGroupMetricRuleRequestEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetThreshold(v string) *PutGroupMetricRuleRequestEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) SetTimes(v int32) *PutGroupMetricRuleRequestEscalationsCritical {
	s.Times = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type PutGroupMetricRuleRequestEscalationsInfo struct {
	// The operator that is used to compare the metric value with the threshold for Info-level alerts. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical methods for Info-level alerts. Separate multiple statistical methods with commas (,).
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified cloud service. The value of this parameter can be Maximum, Minimum, or Average. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Info-level alerts.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an Info-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutGroupMetricRuleRequestEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetComparisonOperator(v string) *PutGroupMetricRuleRequestEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetStatistics(v string) *PutGroupMetricRuleRequestEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetThreshold(v string) *PutGroupMetricRuleRequestEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) SetTimes(v int32) *PutGroupMetricRuleRequestEscalationsInfo {
	s.Times = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type PutGroupMetricRuleRequestEscalationsWarn struct {
	// The operator that is used to compare the metric value with the threshold for Warn-level alerts. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical methods for Warn-level alerts. Separate multiple statistical methods with commas (,).
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified cloud service. The value of this parameter can be Maximum, Minimum, or Average. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Warn-level alerts.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before a Warn-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutGroupMetricRuleRequestEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleRequestEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetComparisonOperator(v string) *PutGroupMetricRuleRequestEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetStatistics(v string) *PutGroupMetricRuleRequestEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetThreshold(v string) *PutGroupMetricRuleRequestEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) SetTimes(v int32) *PutGroupMetricRuleRequestEscalationsWarn {
	s.Times = &v
	return s
}

func (s *PutGroupMetricRuleRequestEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type PutGroupMetricRuleRequestLabels struct {
	// The tag key of the alert rule.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the alert rule.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutGroupMetricRuleRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s PutGroupMetricRuleRequestLabels) GoString() string {
	return s.String()
}

func (s *PutGroupMetricRuleRequestLabels) GetKey() *string {
	return s.Key
}

func (s *PutGroupMetricRuleRequestLabels) GetValue() *string {
	return s.Value
}

func (s *PutGroupMetricRuleRequestLabels) SetKey(v string) *PutGroupMetricRuleRequestLabels {
	s.Key = &v
	return s
}

func (s *PutGroupMetricRuleRequestLabels) SetValue(v string) *PutGroupMetricRuleRequestLabels {
	s.Value = &v
	return s
}

func (s *PutGroupMetricRuleRequestLabels) Validate() error {
	return dara.Validate(s)
}
