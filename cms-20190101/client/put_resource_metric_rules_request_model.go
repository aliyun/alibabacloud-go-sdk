// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *PutResourceMetricRulesRequest
	GetDryRun() *bool
	SetRules(v []*PutResourceMetricRulesRequestRules) *PutResourceMetricRulesRequest
	GetRules() []*PutResourceMetricRulesRequestRules
}

type PutResourceMetricRulesRequest struct {
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: The system only prechecks the request and does not bind a security group to the instance. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): The system sends a normal request. After the request passes the precheck, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The list of threshold alert rules.
	//
	// Valid values of N: 1 to 50.
	//
	// This parameter is required.
	Rules []*PutResourceMetricRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s PutResourceMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *PutResourceMetricRulesRequest) GetRules() []*PutResourceMetricRulesRequestRules {
	return s.Rules
}

func (s *PutResourceMetricRulesRequest) SetDryRun(v bool) *PutResourceMetricRulesRequest {
	s.DryRun = &v
	return s
}

func (s *PutResourceMetricRulesRequest) SetRules(v []*PutResourceMetricRulesRequestRules) *PutResourceMetricRulesRequest {
	s.Rules = v
	return s
}

func (s *PutResourceMetricRulesRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutResourceMetricRulesRequestRules struct {
	Escalations *PutResourceMetricRulesRequestRulesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The alert contact group. Alert notifications are sent to the alert contacts in this alert contact group.
	//
	// Valid values of N: 1 to 50.
	//
	// > An alert contact group is a group of alert contacts and can contain one or more alert contacts. For more information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The effective period of the alert rule.
	//
	// Valid values of N: 1 to 50.
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert email.
	//
	// Valid values of N: 1 to 50.
	//
	// example:
	//
	// ECS instance alert
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The trigger period of the alert rule.
	//
	// Unit: seconds.
	//
	// Valid values of N: 1 to 50.
	//
	// > For more information about how to query the statistical period of a monitoring metric, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The labels that are written to the metric and displayed in the alert notification when the metric meets the alert condition and an alert is triggered.
	Labels []*PutResourceMetricRulesRequestRulesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The monitoring metric name.
	//
	// Valid values of N: 1 to 50.
	//
	// For more information about how to query the monitoring metric name, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The data namespace of the Alibaba Cloud service.
	//
	// Valid values of N: 1 to 50.
	//
	// For more information about how to query the data namespace of an Alibaba Cloud service, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The handling method for alerts when no monitoring data is available. Valid values:
	//
	// - KEEP_LAST_STATE (default): No operations are performed.
	//
	// - INSUFFICIENT_DATA: The alert content indicates that no data is available.
	//
	// - OK: Normal.
	//
	// Valid values of N: 1 to 50.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The inactive time range of the alert rule.
	//
	// Valid values of N: 1 to 50.
	//
	// example:
	//
	// 00:00-06:00
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The statistical period of the monitoring metric.
	//
	// Unit: seconds. Default value: the original reporting period of the monitoring metric.
	//
	// Valid values of N: 1 to 50.
	//
	// > For more information about how to query the statistical period of a monitoring metric, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The resource information. Examples: `[{"instanceId":"i-uf6j91r34rnwawoo****"}]` and `[{"userId":"100931896542****"}]`.
	//
	// Valid values of N: 1 to 50.
	//
	// For more information about the dimensions supported by resource information, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"instanceId":"i-uf6j91r34rnwawoo****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// Valid values of N: 1 to 50.
	//
	// You can enter a new alert rule ID or use an existing alert rule ID in CloudMonitor. For more information about how to query the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// > If you enter a new alert rule ID, a threshold alert rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// a151cd6023eacee2f0978e03863cc1697c89508****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// Valid values of N: 1 to 50.
	//
	// You can enter a new alert rule name or use an existing alert rule name in CloudMonitor. For more information about how to query the name of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// > If you enter a new alert rule name, a threshold alert rule is created.
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
	// The mute epoch.
	//
	// Unit: seconds. Default value: 86400.
	//
	// Valid values of N: 1 to 50.
	//
	// > The mute epoch refers to the interval at which to send an alert notification again if the alert does not recover to Normal after it is triggered.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The webhook URL to which a POST request is sent when an alert callback is triggered.
	//
	// Valid values of N: 1 to 50.
	//
	// example:
	//
	// https://alert.aliyun.com.com:8080/callback
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutResourceMetricRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRules) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRules) GetEscalations() *PutResourceMetricRulesRequestRulesEscalations {
	return s.Escalations
}

func (s *PutResourceMetricRulesRequestRules) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutResourceMetricRulesRequestRules) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutResourceMetricRulesRequestRules) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutResourceMetricRulesRequestRules) GetInterval() *string {
	return s.Interval
}

func (s *PutResourceMetricRulesRequestRules) GetLabels() []*PutResourceMetricRulesRequestRulesLabels {
	return s.Labels
}

func (s *PutResourceMetricRulesRequestRules) GetMetricName() *string {
	return s.MetricName
}

func (s *PutResourceMetricRulesRequestRules) GetNamespace() *string {
	return s.Namespace
}

func (s *PutResourceMetricRulesRequestRules) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *PutResourceMetricRulesRequestRules) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *PutResourceMetricRulesRequestRules) GetPeriod() *string {
	return s.Period
}

func (s *PutResourceMetricRulesRequestRules) GetResources() *string {
	return s.Resources
}

func (s *PutResourceMetricRulesRequestRules) GetRuleId() *string {
	return s.RuleId
}

func (s *PutResourceMetricRulesRequestRules) GetRuleName() *string {
	return s.RuleName
}

func (s *PutResourceMetricRulesRequestRules) GetSendOK() *bool {
	return s.SendOK
}

func (s *PutResourceMetricRulesRequestRules) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *PutResourceMetricRulesRequestRules) GetWebhook() *string {
	return s.Webhook
}

func (s *PutResourceMetricRulesRequestRules) SetEscalations(v *PutResourceMetricRulesRequestRulesEscalations) *PutResourceMetricRulesRequestRules {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetContactGroups(v string) *PutResourceMetricRulesRequestRules {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetEffectiveInterval(v string) *PutResourceMetricRulesRequestRules {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetEmailSubject(v string) *PutResourceMetricRulesRequestRules {
	s.EmailSubject = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetInterval(v string) *PutResourceMetricRulesRequestRules {
	s.Interval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetLabels(v []*PutResourceMetricRulesRequestRulesLabels) *PutResourceMetricRulesRequestRules {
	s.Labels = v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetMetricName(v string) *PutResourceMetricRulesRequestRules {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNamespace(v string) *PutResourceMetricRulesRequestRules {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNoDataPolicy(v string) *PutResourceMetricRulesRequestRules {
	s.NoDataPolicy = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNoEffectiveInterval(v string) *PutResourceMetricRulesRequestRules {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetPeriod(v string) *PutResourceMetricRulesRequestRules {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetResources(v string) *PutResourceMetricRulesRequestRules {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetRuleId(v string) *PutResourceMetricRulesRequestRules {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetRuleName(v string) *PutResourceMetricRulesRequestRules {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetSendOK(v bool) *PutResourceMetricRulesRequestRules {
	s.SendOK = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetSilenceTime(v int32) *PutResourceMetricRulesRequestRules {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetWebhook(v string) *PutResourceMetricRulesRequestRules {
	s.Webhook = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) Validate() error {
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

type PutResourceMetricRulesRequestRulesEscalations struct {
	Critical *PutResourceMetricRulesRequestRulesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *PutResourceMetricRulesRequestRulesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *PutResourceMetricRulesRequestRulesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s PutResourceMetricRulesRequestRulesEscalations) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalations) GetCritical() *PutResourceMetricRulesRequestRulesEscalationsCritical {
	return s.Critical
}

func (s *PutResourceMetricRulesRequestRulesEscalations) GetInfo() *PutResourceMetricRulesRequestRulesEscalationsInfo {
	return s.Info
}

func (s *PutResourceMetricRulesRequestRulesEscalations) GetWarn() *PutResourceMetricRulesRequestRulesEscalationsWarn {
	return s.Warn
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetCritical(v *PutResourceMetricRulesRequestRulesEscalationsCritical) *PutResourceMetricRulesRequestRulesEscalations {
	s.Critical = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetInfo(v *PutResourceMetricRulesRequestRulesEscalationsInfo) *PutResourceMetricRulesRequestRulesEscalations {
	s.Info = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetWarn(v *PutResourceMetricRulesRequestRulesEscalationsWarn) *PutResourceMetricRulesRequestRulesEscalations {
	s.Warn = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) Validate() error {
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

type PutResourceMetricRulesRequestRulesEscalationsCritical struct {
	// The comparison operator for the Critical-level threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to.
	//
	// - GreaterThanThreshold: greater than.
	//
	// - LessThanOrEqualToThreshold: less than or equal to.
	//
	// - LessThanThreshold: less than.
	//
	// - NotEqualToThreshold: not equal to.
	//
	// - EqualToThreshold: equal to.
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
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical method for Critical-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified Alibaba Cloud service, such as Maximum, Minimum, and Average. For more information about how to obtain the value of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure the Settings for at least one of the following alert levels: Critical, Warn (warning), and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Critical-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for Critical-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetN() *int32 {
	return s.N
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetPreCondition() *string {
	return s.PreCondition
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetN(v int32) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.N = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetPreCondition(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRulesRequestRulesEscalationsInfo struct {
	// The comparison operator for the Info-level threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to.
	//
	// - GreaterThanThreshold: greater than.
	//
	// - LessThanOrEqualToThreshold: less than or equal to.
	//
	// - LessThanThreshold: less than.
	//
	// - NotEqualToThreshold: not equal to.
	//
	// - EqualToThreshold: equal to.
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
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical method for Info-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified Alibaba Cloud service, such as Maximum, Minimum, and Average. For more information about how to obtain the value of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure the Settings for at least one of the following alert levels: Critical, Warn (warning), and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Info-level alerts.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for Info-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetN() *int32 {
	return s.N
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetPreCondition() *string {
	return s.PreCondition
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetN(v int32) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.N = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetPreCondition(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRulesRequestRulesEscalationsWarn struct {
	// The comparison operator for the Warn-level threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to.
	//
	// - GreaterThanThreshold: greater than.
	//
	// - LessThanOrEqualToThreshold: less than or equal to.
	//
	// - LessThanThreshold: less than.
	//
	// - NotEqualToThreshold: not equal to.
	//
	// - EqualToThreshold: equal to.
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
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical method for Warn-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// The value of this parameter is determined by the `Statistics` column corresponding to the `MetricName` parameter of the specified Alibaba Cloud service, such as Maximum, Minimum, and Average. For more information about how to obtain the value of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure the Settings for at least one of the following alert levels: Critical, Warn (warning), and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Warn-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for Warn-level alerts.
	//
	// Valid values of N: 1 to 50.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetN() *int32 {
	return s.N
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetPreCondition() *string {
	return s.PreCondition
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetN(v int32) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.N = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetPreCondition(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRulesRequestRulesLabels struct {
	// The tag key.
	//
	// example:
	//
	// tagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// > Tag values support template parameters. The template parameters are replaced with the actual tag values.
	//
	// example:
	//
	// ECS
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesLabels) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesLabels) GetKey() *string {
	return s.Key
}

func (s *PutResourceMetricRulesRequestRulesLabels) GetValue() *string {
	return s.Value
}

func (s *PutResourceMetricRulesRequestRulesLabels) SetKey(v string) *PutResourceMetricRulesRequestRulesLabels {
	s.Key = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesLabels) SetValue(v string) *PutResourceMetricRulesRequestRulesLabels {
	s.Value = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesLabels) Validate() error {
	return dara.Validate(s)
}
