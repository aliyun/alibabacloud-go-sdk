// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEscalations(v *PutResourceMetricRuleShrinkRequestEscalations) *PutResourceMetricRuleShrinkRequest
	GetEscalations() *PutResourceMetricRuleShrinkRequestEscalations
	SetCompositeExpressionShrink(v string) *PutResourceMetricRuleShrinkRequest
	GetCompositeExpressionShrink() *string
	SetContactGroups(v string) *PutResourceMetricRuleShrinkRequest
	GetContactGroups() *string
	SetEffectiveInterval(v string) *PutResourceMetricRuleShrinkRequest
	GetEffectiveInterval() *string
	SetEmailSubject(v string) *PutResourceMetricRuleShrinkRequest
	GetEmailSubject() *string
	SetInterval(v string) *PutResourceMetricRuleShrinkRequest
	GetInterval() *string
	SetLabels(v []*PutResourceMetricRuleShrinkRequestLabels) *PutResourceMetricRuleShrinkRequest
	GetLabels() []*PutResourceMetricRuleShrinkRequestLabels
	SetMetricName(v string) *PutResourceMetricRuleShrinkRequest
	GetMetricName() *string
	SetNamespace(v string) *PutResourceMetricRuleShrinkRequest
	GetNamespace() *string
	SetNoDataPolicy(v string) *PutResourceMetricRuleShrinkRequest
	GetNoDataPolicy() *string
	SetNoEffectiveInterval(v string) *PutResourceMetricRuleShrinkRequest
	GetNoEffectiveInterval() *string
	SetPeriod(v string) *PutResourceMetricRuleShrinkRequest
	GetPeriod() *string
	SetPrometheusShrink(v string) *PutResourceMetricRuleShrinkRequest
	GetPrometheusShrink() *string
	SetResources(v string) *PutResourceMetricRuleShrinkRequest
	GetResources() *string
	SetRuleId(v string) *PutResourceMetricRuleShrinkRequest
	GetRuleId() *string
	SetRuleName(v string) *PutResourceMetricRuleShrinkRequest
	GetRuleName() *string
	SetSilenceTime(v int32) *PutResourceMetricRuleShrinkRequest
	GetSilenceTime() *int32
	SetWebhook(v string) *PutResourceMetricRuleShrinkRequest
	GetWebhook() *string
}

type PutResourceMetricRuleShrinkRequest struct {
	Escalations *PutResourceMetricRuleShrinkRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The trigger conditions for multiple metrics.
	//
	// >  The trigger conditions for a single metric and multiple metrics are mutually exclusive. You cannot specify trigger conditions for a single metric and multiple metrics at the same time.
	CompositeExpressionShrink *string `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty"`
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
	Labels []*PutResourceMetricRuleShrinkRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
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
	PrometheusShrink *string `json:"Prometheus,omitempty" xml:"Prometheus,omitempty"`
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

func (s PutResourceMetricRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleShrinkRequest) GetEscalations() *PutResourceMetricRuleShrinkRequestEscalations {
	return s.Escalations
}

func (s *PutResourceMetricRuleShrinkRequest) GetCompositeExpressionShrink() *string {
	return s.CompositeExpressionShrink
}

func (s *PutResourceMetricRuleShrinkRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutResourceMetricRuleShrinkRequest) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutResourceMetricRuleShrinkRequest) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutResourceMetricRuleShrinkRequest) GetInterval() *string {
	return s.Interval
}

func (s *PutResourceMetricRuleShrinkRequest) GetLabels() []*PutResourceMetricRuleShrinkRequestLabels {
	return s.Labels
}

func (s *PutResourceMetricRuleShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *PutResourceMetricRuleShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutResourceMetricRuleShrinkRequest) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *PutResourceMetricRuleShrinkRequest) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *PutResourceMetricRuleShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *PutResourceMetricRuleShrinkRequest) GetPrometheusShrink() *string {
	return s.PrometheusShrink
}

func (s *PutResourceMetricRuleShrinkRequest) GetResources() *string {
	return s.Resources
}

func (s *PutResourceMetricRuleShrinkRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *PutResourceMetricRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutResourceMetricRuleShrinkRequest) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *PutResourceMetricRuleShrinkRequest) GetWebhook() *string {
	return s.Webhook
}

func (s *PutResourceMetricRuleShrinkRequest) SetEscalations(v *PutResourceMetricRuleShrinkRequestEscalations) *PutResourceMetricRuleShrinkRequest {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetCompositeExpressionShrink(v string) *PutResourceMetricRuleShrinkRequest {
	s.CompositeExpressionShrink = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetContactGroups(v string) *PutResourceMetricRuleShrinkRequest {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetEffectiveInterval(v string) *PutResourceMetricRuleShrinkRequest {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetEmailSubject(v string) *PutResourceMetricRuleShrinkRequest {
	s.EmailSubject = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetInterval(v string) *PutResourceMetricRuleShrinkRequest {
	s.Interval = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetLabels(v []*PutResourceMetricRuleShrinkRequestLabels) *PutResourceMetricRuleShrinkRequest {
	s.Labels = v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetMetricName(v string) *PutResourceMetricRuleShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetNamespace(v string) *PutResourceMetricRuleShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetNoDataPolicy(v string) *PutResourceMetricRuleShrinkRequest {
	s.NoDataPolicy = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetNoEffectiveInterval(v string) *PutResourceMetricRuleShrinkRequest {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetPeriod(v string) *PutResourceMetricRuleShrinkRequest {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetPrometheusShrink(v string) *PutResourceMetricRuleShrinkRequest {
	s.PrometheusShrink = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetResources(v string) *PutResourceMetricRuleShrinkRequest {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetRuleId(v string) *PutResourceMetricRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetRuleName(v string) *PutResourceMetricRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetSilenceTime(v int32) *PutResourceMetricRuleShrinkRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) SetWebhook(v string) *PutResourceMetricRuleShrinkRequest {
	s.Webhook = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequest) Validate() error {
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
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

type PutResourceMetricRuleShrinkRequestEscalations struct {
	Critical *PutResourceMetricRuleShrinkRequestEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *PutResourceMetricRuleShrinkRequestEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *PutResourceMetricRuleShrinkRequestEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s PutResourceMetricRuleShrinkRequestEscalations) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleShrinkRequestEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) GetCritical() *PutResourceMetricRuleShrinkRequestEscalationsCritical {
	return s.Critical
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) GetInfo() *PutResourceMetricRuleShrinkRequestEscalationsInfo {
	return s.Info
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) GetWarn() *PutResourceMetricRuleShrinkRequestEscalationsWarn {
	return s.Warn
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) SetCritical(v *PutResourceMetricRuleShrinkRequestEscalationsCritical) *PutResourceMetricRuleShrinkRequestEscalations {
	s.Critical = v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) SetInfo(v *PutResourceMetricRuleShrinkRequestEscalationsInfo) *PutResourceMetricRuleShrinkRequestEscalations {
	s.Info = v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) SetWarn(v *PutResourceMetricRuleShrinkRequestEscalationsWarn) *PutResourceMetricRuleShrinkRequestEscalations {
	s.Warn = v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalations) Validate() error {
	if s.Critical != nil {
		if err := s.Critical.Validate(); err != nil {
			return err
		}
	}
	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return err
		}
	}
	if s.Warn != nil {
		if err := s.Warn.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutResourceMetricRuleShrinkRequestEscalationsCritical struct {
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

func (s PutResourceMetricRuleShrinkRequestEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleShrinkRequestEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRuleShrinkRequestEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) SetStatistics(v string) *PutResourceMetricRuleShrinkRequestEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) SetThreshold(v string) *PutResourceMetricRuleShrinkRequestEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) SetTimes(v int32) *PutResourceMetricRuleShrinkRequestEscalationsCritical {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleShrinkRequestEscalationsInfo struct {
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

func (s PutResourceMetricRuleShrinkRequestEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleShrinkRequestEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRuleShrinkRequestEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) SetStatistics(v string) *PutResourceMetricRuleShrinkRequestEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) SetThreshold(v string) *PutResourceMetricRuleShrinkRequestEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) SetTimes(v int32) *PutResourceMetricRuleShrinkRequestEscalationsInfo {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleShrinkRequestEscalationsWarn struct {
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

func (s PutResourceMetricRuleShrinkRequestEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleShrinkRequestEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRuleShrinkRequestEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) SetStatistics(v string) *PutResourceMetricRuleShrinkRequestEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) SetThreshold(v string) *PutResourceMetricRuleShrinkRequestEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) SetTimes(v int32) *PutResourceMetricRuleShrinkRequestEscalationsWarn {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRuleShrinkRequestLabels struct {
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

func (s PutResourceMetricRuleShrinkRequestLabels) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleShrinkRequestLabels) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleShrinkRequestLabels) GetKey() *string {
	return s.Key
}

func (s *PutResourceMetricRuleShrinkRequestLabels) GetValue() *string {
	return s.Value
}

func (s *PutResourceMetricRuleShrinkRequestLabels) SetKey(v string) *PutResourceMetricRuleShrinkRequestLabels {
	s.Key = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestLabels) SetValue(v string) *PutResourceMetricRuleShrinkRequestLabels {
	s.Value = &v
	return s
}

func (s *PutResourceMetricRuleShrinkRequestLabels) Validate() error {
	return dara.Validate(s)
}
