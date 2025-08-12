// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEscalations(v *PutResourceMetricRuleRequestEscalations) *PutResourceMetricRuleRequest
	GetEscalations() *PutResourceMetricRuleRequestEscalations
	SetCompositeExpression(v *PutResourceMetricRuleRequestCompositeExpression) *PutResourceMetricRuleRequest
	GetCompositeExpression() *PutResourceMetricRuleRequestCompositeExpression
	SetContactGroups(v string) *PutResourceMetricRuleRequest
	GetContactGroups() *string
	SetEffectiveInterval(v string) *PutResourceMetricRuleRequest
	GetEffectiveInterval() *string
	SetEmailSubject(v string) *PutResourceMetricRuleRequest
	GetEmailSubject() *string
	SetInterval(v string) *PutResourceMetricRuleRequest
	GetInterval() *string
	SetLabels(v []*PutResourceMetricRuleRequestLabels) *PutResourceMetricRuleRequest
	GetLabels() []*PutResourceMetricRuleRequestLabels
	SetMetricName(v string) *PutResourceMetricRuleRequest
	GetMetricName() *string
	SetNamespace(v string) *PutResourceMetricRuleRequest
	GetNamespace() *string
	SetNoDataPolicy(v string) *PutResourceMetricRuleRequest
	GetNoDataPolicy() *string
	SetNoEffectiveInterval(v string) *PutResourceMetricRuleRequest
	GetNoEffectiveInterval() *string
	SetPeriod(v string) *PutResourceMetricRuleRequest
	GetPeriod() *string
	SetPrometheus(v *PutResourceMetricRuleRequestPrometheus) *PutResourceMetricRuleRequest
	GetPrometheus() *PutResourceMetricRuleRequestPrometheus
	SetResources(v string) *PutResourceMetricRuleRequest
	GetResources() *string
	SetRuleId(v string) *PutResourceMetricRuleRequest
	GetRuleId() *string
	SetRuleName(v string) *PutResourceMetricRuleRequest
	GetRuleName() *string
	SetSilenceTime(v int32) *PutResourceMetricRuleRequest
	GetSilenceTime() *int32
	SetWebhook(v string) *PutResourceMetricRuleRequest
	GetWebhook() *string
}

type PutResourceMetricRuleRequest struct {
	Escalations *PutResourceMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The trigger conditions for multiple metrics.
	//
	// >  The trigger conditions for a single metric and multiple metrics are mutually exclusive. You cannot specify trigger conditions for a single metric and multiple metrics at the same time.
	CompositeExpression *PutResourceMetricRuleRequestCompositeExpression `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty" type:"Struct"`
	// The alert contact groups. Alert notifications are sent to the alert contacts in the alert contact group.
	//
	// >  An alert contact group can contain one or more alert contacts. For information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The period of time during which the alert rule is effective.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert notification email.
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The interval at which alerts are triggered based on the alert rule. Unit: seconds.
	//
	// >  For more information about how to query the statistical periods of metrics, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// If the metric meets the specified condition in the alert rule and CloudMonitor sends an alert notification, the tag is also written to the metric and displayed in the alert notification.
	//
	// >  This parameter is equivalent to the Label parameter of Prometheus alerts.
	Labels []*PutResourceMetricRuleRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name. For more information about how to query metric names, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// >  If you create a Prometheus alert rule for Hybrid Cloud Monitoring, you must set this parameter to the name of the namespace. For more information about how to query the names of namespaces, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service. For more information about how to query the namespaces of cloud services, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// >  If you create a Prometheus alert rule for Hybrid Cloud Monitoring, you must set this parameter to `acs_prometheus`.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The method that is used to handle alerts when no monitoring data is found. Valid value:
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
	// The period of time during which the alert rule is ineffective.
	//
	// example:
	//
	// 00:00-06:00
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The statistical period of the metric. Unit: seconds. The default value is the interval at which the monitoring data of the metric is collected.
	//
	// >  For more information about how to query the statistical periods of metrics, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// Prometheus alerts.
	//
	// >  This parameter is required only if you create a Prometheus alert rule for Hybrid Cloud Monitoring.
	Prometheus *PutResourceMetricRuleRequestPrometheus `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	// The resource information. Examples: `[{"instanceId":"i-uf6j91r34rnwawoo****"}]` and `[{"userId":"100931896542****"}]`.
	//
	// For more information about the supported dimensions that are used to query resources, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// [{"instanceId":"i-uf6j91r34rnwawoo****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// You can specify a new ID or the ID of an existing alert rule. For more information about how to query the IDs of alert rules, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// >  If you specify a new ID, a threshold-triggered alert rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// a151cd6023eacee2f0978e03863cc1697c89508****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// You can specify a new name or the name of an existing alert rule. For more information about how to query the names of alert rules, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// >  If you specify a new name, a threshold-triggered alert rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period during which new alert notifications are not sent even if the trigger conditions are met. Unit: seconds. Default value: 86400.
	//
	// >  If an alert is not cleared after the mute period ends, CloudMonitor resends an alert notification.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered based on the alert rule.
	//
	// example:
	//
	// https://alert.aliyun.com.com:8080/callback
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutResourceMetricRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequest) GetEscalations() *PutResourceMetricRuleRequestEscalations {
	return s.Escalations
}

func (s *PutResourceMetricRuleRequest) GetCompositeExpression() *PutResourceMetricRuleRequestCompositeExpression {
	return s.CompositeExpression
}

func (s *PutResourceMetricRuleRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutResourceMetricRuleRequest) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutResourceMetricRuleRequest) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutResourceMetricRuleRequest) GetInterval() *string {
	return s.Interval
}

func (s *PutResourceMetricRuleRequest) GetLabels() []*PutResourceMetricRuleRequestLabels {
	return s.Labels
}

func (s *PutResourceMetricRuleRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *PutResourceMetricRuleRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutResourceMetricRuleRequest) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *PutResourceMetricRuleRequest) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *PutResourceMetricRuleRequest) GetPeriod() *string {
	return s.Period
}

func (s *PutResourceMetricRuleRequest) GetPrometheus() *PutResourceMetricRuleRequestPrometheus {
	return s.Prometheus
}

func (s *PutResourceMetricRuleRequest) GetResources() *string {
	return s.Resources
}

func (s *PutResourceMetricRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutResourceMetricRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutResourceMetricRuleRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *PutResourceMetricRuleRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *PutResourceMetricRuleRequest) SetEscalations(v *PutResourceMetricRuleRequestEscalations) *PutResourceMetricRuleRequest {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRuleRequest) SetCompositeExpression(v *PutResourceMetricRuleRequestCompositeExpression) *PutResourceMetricRuleRequest {
	s.CompositeExpression = v
	return s
}

func (s *PutResourceMetricRuleRequest) SetContactGroups(v string) *PutResourceMetricRuleRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetEffectiveInterval(v string) *PutResourceMetricRuleRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetEmailSubject(v string) *PutResourceMetricRuleRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetInterval(v string) *PutResourceMetricRuleRequest {
	s.Interval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetLabels(v []*PutResourceMetricRuleRequestLabels) *PutResourceMetricRuleRequest {
	s.Labels = v
	return s
}

func (s *PutResourceMetricRuleRequest) SetMetricName(v string) *PutResourceMetricRuleRequest {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNamespace(v string) *PutResourceMetricRuleRequest {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNoDataPolicy(v string) *PutResourceMetricRuleRequest {
	s.NoDataPolicy = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetNoEffectiveInterval(v string) *PutResourceMetricRuleRequest {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetPeriod(v string) *PutResourceMetricRuleRequest {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetPrometheus(v *PutResourceMetricRuleRequestPrometheus) *PutResourceMetricRuleRequest {
	s.Prometheus = v
	return s
}

func (s *PutResourceMetricRuleRequest) SetResources(v string) *PutResourceMetricRuleRequest {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetRuleId(v string) *PutResourceMetricRuleRequest {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetRuleName(v string) *PutResourceMetricRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetSilenceTime(v int32) *PutResourceMetricRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRuleRequest) SetWebhook(v string) *PutResourceMetricRuleRequest {
	s.Webhook = &v
	return s
}

func (s *PutResourceMetricRuleRequest) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestEscalations struct {
	Critical *PutResourceMetricRuleRequestEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *PutResourceMetricRuleRequestEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *PutResourceMetricRuleRequestEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s PutResourceMetricRuleRequestEscalations) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalations) GetCritical() *PutResourceMetricRuleRequestEscalationsCritical {
	return s.Critical
}

func (s *PutResourceMetricRuleRequestEscalations) GetInfo() *PutResourceMetricRuleRequestEscalationsInfo {
	return s.Info
}

func (s *PutResourceMetricRuleRequestEscalations) GetWarn() *PutResourceMetricRuleRequestEscalationsWarn {
	return s.Warn
}

func (s *PutResourceMetricRuleRequestEscalations) SetCritical(v *PutResourceMetricRuleRequestEscalationsCritical) *PutResourceMetricRuleRequestEscalations {
	s.Critical = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) SetInfo(v *PutResourceMetricRuleRequestEscalationsInfo) *PutResourceMetricRuleRequestEscalations {
	s.Info = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) SetWarn(v *PutResourceMetricRuleRequestEscalationsWarn) *PutResourceMetricRuleRequestEscalations {
	s.Warn = v
	return s
}

func (s *PutResourceMetricRuleRequestEscalations) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestEscalationsCritical struct {
	// The operator that is used to compare the metric value with the threshold for Critical-level alerts. Valid value:
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
	// 	- EqualToThreshold: equal to the threshold
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
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical methods for Critical-level alerts.
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified cloud service. The value of this parameter can be Maximum, Minimum, or Average. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Critical-level alerts.
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before a Critical-level alert is triggered.
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsCritical {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestEscalationsInfo struct {
	// The operator that is used to compare the metric value with the threshold for Info-level alerts. Valid value:
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
	// 	- EqualToThreshold: equal to the threshold
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
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical methods for Info-level alerts.
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified cloud service. The value of this parameter can be Maximum, Minimum, or Average. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Info-level alerts.
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an Info-level alert is triggered.
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsInfo {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestEscalationsWarn struct {
	// The operator that is used to compare the metric value with the threshold for Warn-level alerts. Valid value:
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
	// 	- EqualToThreshold: equal to the threshold
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
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical methods for Warn-level alerts.
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified cloud service. The value of this parameter can be Maximum, Minimum, or Average. For more information about how to obtain the value of this parameter, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Warn-level alerts.
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before a Warn-level alert is triggered.
	//
	// >  You must select at least one of the Critical, Warn, and Info alert levels and specify the Statistics, ComparisonOperator, Threshold, and Times parameters for each alert level.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetStatistics(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetThreshold(v string) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) SetTimes(v int32) *PutResourceMetricRuleRequestEscalationsWarn {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleRequestEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestCompositeExpression struct {
	// The trigger conditions that are created in standard mode.
	ExpressionList []*PutResourceMetricRuleRequestCompositeExpressionExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Repeated"`
	// The relationship between the trigger conditions for multiple metrics. Valid value:
	//
	// 	- `&&`: An alert is triggered only if all metrics meet the trigger conditions. An alert is triggered only if the results of all expressions specified in the ExpressionList parameter are `true`.
	//
	// 	- `||`: An alert is triggered if one of the metrics meets the trigger conditions.
	//
	// example:
	//
	// ||
	ExpressionListJoin *string `json:"ExpressionListJoin,omitempty" xml:"ExpressionListJoin,omitempty"`
	// The trigger conditions that are created by using expressions. You can use expressions to create trigger conditions in the following scenarios:
	//
	// 	- Set an alert blacklist for specific resources. For example, if you specify `$instanceId != \\"i-io8kfvcpp7x5****\\" ``&&`` $Average > 50`, no alert is triggered when the `average metric value` of the `i-io8kfvcpp7x5****` instance exceeds 50.
	//
	// 	- Set a special alert threshold for a specified instance in the rule. For example, if you specify `$Average > ($instanceId == \\"i-io8kfvcpp7x5****\\"? 80: 50)`, an alert is triggered when the `average metric value` of the `i-io8kfvcpp7x5****` instance exceeds 80 or the `average metric value` of other instances exceeds 50.
	//
	// 	- Limit the number of instances whose metric values exceed the threshold. For example, if you specify `count($Average > 20) > 3`, an alert is triggered only when the `average metric value` of more than three instances exceeds 20.
	//
	// example:
	//
	// $Average > ($instanceId == \\"i-io8kfvcpp7x5****\\"? 80: 50)
	ExpressionRaw *string `json:"ExpressionRaw,omitempty" xml:"ExpressionRaw,omitempty"`
	// The alert level. Valid values:
	//
	// 	- Critical
	//
	// 	- Warn
	//
	// 	- Info
	//
	// example:
	//
	// Critical
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of consecutive triggers. If the number of times that the metric values meet the trigger conditions reaches the value of this parameter, CloudMonitor sends alert notifications.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestCompositeExpression) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestCompositeExpression) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestCompositeExpression) GetExpressionList() []*PutResourceMetricRuleRequestCompositeExpressionExpressionList {
	return s.ExpressionList
}

func (s *PutResourceMetricRuleRequestCompositeExpression) GetExpressionListJoin() *string {
	return s.ExpressionListJoin
}

func (s *PutResourceMetricRuleRequestCompositeExpression) GetExpressionRaw() *string {
	return s.ExpressionRaw
}

func (s *PutResourceMetricRuleRequestCompositeExpression) GetLevel() *string {
	return s.Level
}

func (s *PutResourceMetricRuleRequestCompositeExpression) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleRequestCompositeExpression) SetExpressionList(v []*PutResourceMetricRuleRequestCompositeExpressionExpressionList) *PutResourceMetricRuleRequestCompositeExpression {
	s.ExpressionList = v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpression) SetExpressionListJoin(v string) *PutResourceMetricRuleRequestCompositeExpression {
	s.ExpressionListJoin = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpression) SetExpressionRaw(v string) *PutResourceMetricRuleRequestCompositeExpression {
	s.ExpressionRaw = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpression) SetLevel(v string) *PutResourceMetricRuleRequestCompositeExpression {
	s.Level = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpression) SetTimes(v int32) *PutResourceMetricRuleRequestCompositeExpression {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpression) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestCompositeExpressionExpressionList struct {
	// The operator that is used to compare the metric value with the threshold. Valid value:
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
	// 	- EqualToThreshold: equal to the threshold
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
	// The metric that is used to monitor the cloud service.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The aggregation period of the metric.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 60
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The statistical method of the metric. Valid value:
	//
	// 	- $Maximum: the maximum value
	//
	// 	- $Minimum: the minimum value
	//
	// 	- $Average: the average value
	//
	// 	- $Availability: the availability rate (usually used for site monitoring)
	//
	// >  `$` is the prefix of the metric. For information about the Alibaba Cloud services that are supported by CloudMonitor, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// $Maximum
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert threshold.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s PutResourceMetricRuleRequestCompositeExpressionExpressionList) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestCompositeExpressionExpressionList) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) GetMetricName() *string {
	return s.MetricName
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) GetPeriod() *int64 {
	return s.Period
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) SetComparisonOperator(v string) *PutResourceMetricRuleRequestCompositeExpressionExpressionList {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) SetMetricName(v string) *PutResourceMetricRuleRequestCompositeExpressionExpressionList {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) SetPeriod(v int64) *PutResourceMetricRuleRequestCompositeExpressionExpressionList {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) SetStatistics(v string) *PutResourceMetricRuleRequestCompositeExpressionExpressionList {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) SetThreshold(v string) *PutResourceMetricRuleRequestCompositeExpressionExpressionList {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleRequestCompositeExpressionExpressionList) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestLabels struct {
	// The tag key.
	//
	// example:
	//
	// tagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// >  You can use a template parameter to specify a tag value. CloudMonitor replaces the value of the template parameter with an actual tag value.
	//
	// example:
	//
	// ECS
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutResourceMetricRuleRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestLabels) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestLabels) GetKey() *string {
	return s.Key
}

func (s *PutResourceMetricRuleRequestLabels) GetValue() *string {
	return s.Value
}

func (s *PutResourceMetricRuleRequestLabels) SetKey(v string) *PutResourceMetricRuleRequestLabels {
	s.Key = &v
	return s
}

func (s *PutResourceMetricRuleRequestLabels) SetValue(v string) *PutResourceMetricRuleRequestLabels {
	s.Value = &v
	return s
}

func (s *PutResourceMetricRuleRequestLabels) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestPrometheus struct {
	// The annotations of the Prometheus alert rule. When a Prometheus alert is triggered, the system renders the annotated keys and values to help you understand the metrics and alert rule.
	//
	// >  This parameter is equivalent to the annotations parameter of open source Prometheus.
	Annotations []*PutResourceMetricRuleRequestPrometheusAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// The alert level. Valid values:
	//
	// 	- Critical
	//
	// 	- Warn
	//
	// 	- Info
	//
	// example:
	//
	// Critical
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// PromQL statements are supported.
	//
	// >  The data obtained by using the PromQL query statement is the monitoring data. You must include the alert threshold in this statement.
	//
	// example:
	//
	// cpuUsage{instanceId="xxxx"}[1m]>90
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The number of consecutive triggers. If the number of times that the metric values meet the trigger conditions reaches the value of this parameter, CloudMonitor sends alert notifications.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRuleRequestPrometheus) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestPrometheus) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestPrometheus) GetAnnotations() []*PutResourceMetricRuleRequestPrometheusAnnotations {
	return s.Annotations
}

func (s *PutResourceMetricRuleRequestPrometheus) GetLevel() *string {
	return s.Level
}

func (s *PutResourceMetricRuleRequestPrometheus) GetPromQL() *string {
	return s.PromQL
}

func (s *PutResourceMetricRuleRequestPrometheus) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleRequestPrometheus) SetAnnotations(v []*PutResourceMetricRuleRequestPrometheusAnnotations) *PutResourceMetricRuleRequestPrometheus {
	s.Annotations = v
	return s
}

func (s *PutResourceMetricRuleRequestPrometheus) SetLevel(v string) *PutResourceMetricRuleRequestPrometheus {
	s.Level = &v
	return s
}

func (s *PutResourceMetricRuleRequestPrometheus) SetPromQL(v string) *PutResourceMetricRuleRequestPrometheus {
	s.PromQL = &v
	return s
}

func (s *PutResourceMetricRuleRequestPrometheus) SetTimes(v int32) *PutResourceMetricRuleRequestPrometheus {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleRequestPrometheus) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleRequestPrometheusAnnotations struct {
	// The key of the annotation.
	//
	// example:
	//
	// summary
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the annotation.
	//
	// example:
	//
	// {{ $labels.instance }} CPU usage above 10% {current value: {{ humanizePercentage $value }} }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutResourceMetricRuleRequestPrometheusAnnotations) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequestPrometheusAnnotations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequestPrometheusAnnotations) GetKey() *string {
	return s.Key
}

func (s *PutResourceMetricRuleRequestPrometheusAnnotations) GetValue() *string {
	return s.Value
}

func (s *PutResourceMetricRuleRequestPrometheusAnnotations) SetKey(v string) *PutResourceMetricRuleRequestPrometheusAnnotations {
	s.Key = &v
	return s
}

func (s *PutResourceMetricRuleRequestPrometheusAnnotations) SetValue(v string) *PutResourceMetricRuleRequestPrometheusAnnotations {
	s.Value = &v
	return s
}

func (s *PutResourceMetricRuleRequestPrometheusAnnotations) Validate() error {
	return dara.Validate(s)
}
