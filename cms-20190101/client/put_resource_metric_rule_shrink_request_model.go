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
	SetDryRun(v bool) *PutResourceMetricRuleShrinkRequest
	GetDryRun() *bool
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
	SetSendOK(v bool) *PutResourceMetricRuleShrinkRequest
	GetSendOK() *bool
	SetSilenceTime(v int32) *PutResourceMetricRuleShrinkRequest
	GetSilenceTime() *int32
	SetWebhook(v string) *PutResourceMetricRuleShrinkRequest
	GetWebhook() *string
}

type PutResourceMetricRuleShrinkRequest struct {
	Escalations *PutResourceMetricRuleShrinkRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The alert conditions for multiple metrics.
	//
	// > Single-metric and multi-metric alert conditions are mutually exclusive and cannot be configured at the same time.
	CompositeExpressionShrink *string `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty"`
	// The alert contact group. Alert notifications are sent to the alert contacts in this alert contact group.
	//
	// > An alert contact group contains one or more alert contacts. For more information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: The request is prechecked but not executed. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): The request is sent and executed. If the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The effective period of the alert rule.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert email.
	//
	// example:
	//
	// ECS instance alert
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The trigger epoch of the alert rule. Unit: seconds.
	//
	// > For more information about how to query the statistical epoch of a metric, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The labels that are added to the metric and displayed in the alert notification when the metric meets the alert condition and an alert is triggered.
	//
	// > This feature is the same as the Label feature in Prometheus alerts.
	Labels []*PutResourceMetricRuleShrinkRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the metric. For more information about how to query the name of a metric, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > When you create a Prometheus alert rule for CloudMonitor, this parameter specifies the name of the metric repository. For more information about how to obtain the name of a metric repository, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The data namespace of the Alibaba Cloud service. For more information about how to query the data namespace of an Alibaba Cloud service, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > When you create a Prometheus alert rule for CloudMonitor, you must configure this parameter to `acs_prometheus`.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The action to take when no monitoring data is found. Valid values:
	//
	// - KEEP_LAST_STATE (default): No action is taken.
	//
	// - INSUFFICIENT_DATA: The alert content indicates that no data is found.
	//
	// - OK: The status is normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time range during which the alert rule is inactive.
	//
	// example:
	//
	// 00:00-06:00
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The statistical epoch of the metric. Unit: seconds. The default value is the original reporting epoch of the metric.
	//
	// > For more information about how to query the statistical epoch of a metric, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The Prometheus alert.
	//
	// > You need to configure this parameter only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
	PrometheusShrink *string `json:"Prometheus,omitempty" xml:"Prometheus,omitempty"`
	// The resource information. Examples: `[{"instanceId":"i-uf6j91r34rnwawoo****"}]` and `[{"userId":"100931896542****"}]`.
	//
	// For more information about the dimensions supported by resource information, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// [{"instanceId":"i-uf6j91r34rnwawoo****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// You can enter a new alert rule ID or use an existing alert rule ID in CloudMonitor. For more information about how to query the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// > Entering a new alert rule ID indicates that a threshold-based alert rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// a151cd6023eacee2f0978e03863cc1697c89508****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// You can enter a new alert rule name or use an existing alert rule name in CloudMonitor. For more information about how to query the name of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// > Entering a new alert rule name indicates that a threshold-based alert rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Specifies whether to send a recovery notification.
	//
	// example:
	//
	// true
	SendOK *bool `json:"SendOK,omitempty" xml:"SendOK,omitempty"`
	// The mute epoch. Unit: seconds. Default value: 86400.
	//
	// > The mute epoch refers to the interval at which to send an alert notification again if the alert does not recover to Normal after it is triggered.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The webhook URL to which a POST request is sent when an alert is triggered.
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

func (s *PutResourceMetricRuleShrinkRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *PutResourceMetricRuleShrinkRequest) GetSendOK() *bool {
	return s.SendOK
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

func (s *PutResourceMetricRuleShrinkRequest) SetDryRun(v bool) *PutResourceMetricRuleShrinkRequest {
	s.DryRun = &v
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

func (s *PutResourceMetricRuleShrinkRequest) SetSendOK(v bool) *PutResourceMetricRuleShrinkRequest {
	s.SendOK = &v
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
	// The comparison operator for the Critical-level threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to the threshold.
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanOrEqualToThreshold: less than or equal to the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// - NotEqualToThreshold: not equal to the threshold.
	//
	// - EqualToThreshold: equal to the threshold.
	//
	// - GreaterThanYesterday: greater than the value at the same time yesterday.
	//
	// - LessThanYesterday: less than the value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the value at the same time last week.
	//
	// - LessThanLastWeek: less than the value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the value in the previous period.
	//
	// - LessThanLastPeriod: less than the value in the previous period.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for the Critical-level alert.
	//
	// The valid values of this parameter are determined by the `Statistics` column corresponding to the `MetricName` of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values for this parameter, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for the Critical-level alert.
	//
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for the Critical-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
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
	// The comparison operator for the Info-level threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to the threshold.
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanOrEqualToThreshold: less than or equal to the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// - NotEqualToThreshold: not equal to the threshold.
	//
	// - EqualToThreshold: equal to the threshold.
	//
	// - GreaterThanYesterday: greater than the value at the same time yesterday.
	//
	// - LessThanYesterday: less than the value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the value at the same time last week.
	//
	// - LessThanLastWeek: less than the value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the value in the previous period.
	//
	// - LessThanLastPeriod: less than the value in the previous period.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for the Info-level alert.
	//
	// The valid values of this parameter are determined by the `Statistics` column corresponding to the `MetricName` of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values for this parameter, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for the Info-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for the Info-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
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
	// The comparison operator for the Warn-level threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to the threshold.
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanOrEqualToThreshold: less than or equal to the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// - NotEqualToThreshold: not equal to the threshold.
	//
	// - EqualToThreshold: equal to the threshold.
	//
	// - GreaterThanYesterday: greater than the value at the same time yesterday.
	//
	// - LessThanYesterday: less than the value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the value at the same time last week.
	//
	// - LessThanLastWeek: less than the value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the value in the previous period.
	//
	// - LessThanLastPeriod: less than the value in the previous period.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for the Warn-level alert.
	//
	// The valid values of this parameter are determined by the `Statistics` column corresponding to the `MetricName` of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values for this parameter, see [Alibaba Cloud service metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for the Warn-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for the Warn-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
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
	// The label key.
	//
	// example:
	//
	// tagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The label value.
	//
	// > The label value supports template parameters. The template parameters are replaced with the actual label values.
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
