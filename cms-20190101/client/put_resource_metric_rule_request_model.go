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
	SetClientToken(v string) *PutResourceMetricRuleRequest
	GetClientToken() *string
	SetCompositeExpression(v *PutResourceMetricRuleRequestCompositeExpression) *PutResourceMetricRuleRequest
	GetCompositeExpression() *PutResourceMetricRuleRequestCompositeExpression
	SetContactGroups(v string) *PutResourceMetricRuleRequest
	GetContactGroups() *string
	SetDryRun(v bool) *PutResourceMetricRuleRequest
	GetDryRun() *bool
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
	SetSendOK(v bool) *PutResourceMetricRuleRequest
	GetSendOK() *bool
	SetSilenceTime(v int32) *PutResourceMetricRuleRequest
	GetSilenceTime() *int32
	SetWebhook(v string) *PutResourceMetricRuleRequest
	GetWebhook() *string
}

type PutResourceMetricRuleRequest struct {
	Escalations *PutResourceMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can generate the token on your client. Make sure that the token is unique across different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 6868E9A1-56E6-5FD7-9684-E2963BBAFB5F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The alert conditions for multiple metrics.
	//
	// > Single-metric and multi-metric alert conditions are mutually exclusive and cannot be configured at the same time.
	CompositeExpression *PutResourceMetricRuleRequestCompositeExpression `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty" type:"Struct"`
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
	// - **true**: sends a check request and does not bind a security group to the instance. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the request fails the check, the system returns an error. If the request passes the check, the system returns the error code `DryRunOperation`.
	//
	// - **false**: sends a normal request. If the request passes the check, the system returns an HTTP 2xx status code and performs the operation. This is the default value.
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
	// > For more information about how to query the statistical period of a metric, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The labels that are written to the metric and displayed in the alert notification when the metric meets the alert condition and an alert is triggered.
	//
	// > This feature is the same as the Label feature in Prometheus alerts.
	Labels []*PutResourceMetricRuleRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name. For more information about how to query the metric name, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > When you create a Prometheus alert rule for Hybrid CloudMonitor, this parameter specifies the metric repository name. For more information about how to obtain the metric repository name, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The data namespace of the Alibaba Cloud service. For more information about how to query the data namespace of an Alibaba Cloud service, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > When you create a Prometheus alert rule for Hybrid CloudMonitor, you must set this parameter to `acs_prometheus`.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The action to take when no monitoring data is available. Valid values:
	//
	// - KEEP_LAST_STATE: takes no action. This is the default value.
	//
	// - INSUFFICIENT_DATA: triggers an alert indicating that no data is available.
	//
	// - OK: triggers an alert indicating that the status is normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The period during which the alert rule is inactive.
	//
	// example:
	//
	// 00:00-06:00
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The statistical period of the metric. Unit: seconds. The default value is the original reporting epoch of the metric.
	//
	// > For more information about how to query the statistical period of a metric, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The Prometheus alert.
	//
	// > You need to configure this parameter only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
	Prometheus *PutResourceMetricRuleRequestPrometheus `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	// The resource information. Examples: `[{"instanceId":"i-uf6j91r34rnwawoo****"}]` and `[{"userId":"100931896542****"}]`.
	//
	// For more information about the dimensions supported by the resource information, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// [{"instanceId":"i-uf6j91r34rnwawoo****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The alert rule ID.
	//
	// You can enter a new alert rule ID or use an existing alert rule ID in CloudMonitor. For more information about how to query the alert rule ID, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// > Entering a new alert rule ID indicates that a threshold-based alert rule is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// a151cd6023eacee2f0978e03863cc1697c89508****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The alert rule name.
	//
	// You can enter a new alert rule name or use an existing alert rule name in CloudMonitor. For more information about how to query the alert rule name, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
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
	// > The mute epoch refers to the interval to send an alert notification again if the alert does not recover to Normal after it is triggered.
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

func (s PutResourceMetricRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleRequest) GetEscalations() *PutResourceMetricRuleRequestEscalations {
	return s.Escalations
}

func (s *PutResourceMetricRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PutResourceMetricRuleRequest) GetCompositeExpression() *PutResourceMetricRuleRequestCompositeExpression {
	return s.CompositeExpression
}

func (s *PutResourceMetricRuleRequest) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutResourceMetricRuleRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *PutResourceMetricRuleRequest) GetSendOK() *bool {
	return s.SendOK
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

func (s *PutResourceMetricRuleRequest) SetClientToken(v string) *PutResourceMetricRuleRequest {
	s.ClientToken = &v
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

func (s *PutResourceMetricRuleRequest) SetDryRun(v bool) *PutResourceMetricRuleRequest {
	s.DryRun = &v
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

func (s *PutResourceMetricRuleRequest) SetSendOK(v bool) *PutResourceMetricRuleRequest {
	s.SendOK = &v
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
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	if s.CompositeExpression != nil {
		if err := s.CompositeExpression.Validate(); err != nil {
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
	if s.Prometheus != nil {
		if err := s.Prometheus.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type PutResourceMetricRuleRequestEscalationsCritical struct {
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
	// - GreaterThanLastPeriod: greater than the value in the last period.
	//
	// - LessThanLastPeriod: less than the value in the last period.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for the Critical-level alert.
	//
	// The valid values of this parameter are determined by the Statistics column corresponding to the MetricName of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, the Settings for the Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The Critical-level alert threshold.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for the Critical-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
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
	// - GreaterThanLastPeriod: greater than the value in the last period.
	//
	// - LessThanLastPeriod: less than the value in the last period.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for the Info-level alert.
	//
	// The valid values of this parameter are determined by the Statistics column corresponding to the MetricName of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, the Settings for the Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The Info-level alert threshold.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for the Info-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
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
	// - GreaterThanLastPeriod: greater than the value in the last period.
	//
	// - LessThanLastPeriod: less than the value in the last period.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for the Warn-level alert.
	//
	// The valid values of this parameter are determined by the Statistics column corresponding to the MetricName of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, the Settings for the Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The Warn-level alert threshold.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of retries for the Warn-level alert.
	//
	// > You must configure at least one of the following alert levels: Critical, Warn, and Info. For the configured alert level, you must configure the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
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
	// The list of alert conditions created by using the standard method.
	ExpressionList []*PutResourceMetricRuleRequestCompositeExpressionExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Repeated"`
	// The relationship between the alert conditions for multiple metrics. Valid values:
	//
	// - `&&`: triggers an alert only when all metrics meet the alert conditions. An alert is triggered only when all expressions in ExpressionList evaluate to `true`.
	//
	// - `||`: triggers an alert when any metric meets the alert condition.
	//
	// example:
	//
	// ||
	ExpressionListJoin *string `json:"ExpressionListJoin,omitempty" xml:"ExpressionListJoin,omitempty"`
	// The alert conditions created by using expressions. This includes but is not limited to the following scenarios:
	//
	// - Set an alert blacklist for specific resources. Example: `$instanceId != \\"i-io8kfvcpp7x5****\\" && $Average > 50`. This indicates that no alert is triggered even if the `Average` value of the instance `i-io8kfvcpp7x5****` in the alert rule is greater than 50.
	//
	// - Set a special alert threshold for a specific instance in the rule. Example: `$Average > ($instanceId == \\"i-io8kfvcpp7x5****\\"? 80: 50)`. This indicates that an alert is triggered only when the `Average` value of the instance `i-io8kfvcpp7x5****` in the alert rule is greater than 80. For other instances, an alert is triggered when the `Average` value is greater than 50.
	//
	// - Limit the number of instances that exceed the threshold in the rule. Example: `count($Average > 20) > 3`. This indicates that an alert is triggered only when the number of instances whose `Average` value is greater than 20 in the alert rule is greater than 3.
	//
	// example:
	//
	// $Average > ($instanceId == \\"i-io8kfvcpp7x5****\\"? 80: 50)
	ExpressionRaw *string `json:"ExpressionRaw,omitempty" xml:"ExpressionRaw,omitempty"`
	// The alert level. Valid values:
	//
	// - CRITICAL: critical.
	//
	// - WARN: warning.
	//
	// - INFO: info.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of times the alert condition must be met before an alert notification is sent.
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
	if s.ExpressionList != nil {
		for _, item := range s.ExpressionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutResourceMetricRuleRequestCompositeExpressionExpressionList struct {
	// The threshold comparison operator. Valid values:
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
	// - GreaterThanLastPeriod: greater than the value in the last period.
	//
	// - LessThanLastPeriod: less than the value in the last period.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The metric name of the Alibaba Cloud service.
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
	// The statistical method of the metric. Valid values:
	//
	// - $Maximum: the maximum value.
	//
	// - $Minimum: the minimum value.
	//
	// - $Average: the average value.
	//
	// - $Availability: the active rate. This value is typically used for site monitoring.
	//
	// > `$` is the unified prefix for metrics. For more information about supported Alibaba Cloud services, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The annotations whose keys and values are rendered during Prometheus alerts to help you understand the metrics or alert rules.
	//
	// > This feature is equivalent to the Annotation feature in Prometheus.
	Annotations []*PutResourceMetricRuleRequestPrometheusAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// The alert level. Valid values:
	//
	// - CRITICAL: critical.
	//
	// - WARN: warning.
	//
	// - INFO: info.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The PromQL query statement.
	//
	// > The data obtained by using the PromQL query statement is used as the alert data. You must include the alert threshold in this statement.
	//
	// example:
	//
	// cpuUsage{instanceId="xxxx"}[1m]>90
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The number of times the alert condition must be met before an alert notification is sent.
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
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutResourceMetricRuleRequestPrometheusAnnotations struct {
	// The annotation key.
	//
	// example:
	//
	// summary
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The annotation value.
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
