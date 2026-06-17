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
	SetSendOK(v bool) *PutResourceMetricRuleRequest
	GetSendOK() *bool
	SetSilenceTime(v int32) *PutResourceMetricRuleRequest
	GetSilenceTime() *int32
	SetWebhook(v string) *PutResourceMetricRuleRequest
	GetWebhook() *string
}

type PutResourceMetricRuleRequest struct {
	Escalations *PutResourceMetricRuleRequestEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The alert conditions for multiple metrics.
	//
	// > Single-metric and multi-metric alert conditions are mutually exclusive and cannot be set at the same time.
	CompositeExpression *PutResourceMetricRuleRequestCompositeExpression `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty" type:"Struct"`
	// The alert contact group. Alert notifications are sent to the alert contacts in this alert contact group.
	//
	// > An alert contact group contains one or more alert contacts. For information about how to create alert contacts and alert contact groups, see [PutContact](https://help.aliyun.com/document_detail/114923.html) and [PutContactGroup](https://help.aliyun.com/document_detail/114929.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
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
	// The trigger period of the alert rule. Unit: seconds.
	//
	// > For information about how to query the statistical period of a metric, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The labels that are written to the metric and displayed in alert notifications when the metric meets the alert condition.
	//
	// > This feature is the same as the Label feature in Prometheus alerting.
	Labels []*PutResourceMetricRuleRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the metric. For information about how to query metric names, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > If you create a Prometheus alert rule for Hybrid Cloud Monitoring, this parameter specifies the name of the metric repository. For information about how to obtain the metric repository name, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service. For information about how to query the namespace of an Alibaba Cloud service, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > If you create a Prometheus alert rule for Hybrid Cloud Monitoring, set this parameter to `acs_prometheus`.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The processing method when no monitoring data is found. Valid values:
	//
	// - KEEP_LAST_STATE (default): No action is taken.
	//
	// - INSUFFICIENT_DATA: An alert whose content is "Insufficient data" is triggered.
	//
	// - OK: The status is considered normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time range during which the alert rule is ineffective.
	//
	// example:
	//
	// 00:00-06:00
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The statistical period of the metric. Unit: seconds. The default value is the original reporting period of the metric.
	//
	// > For information about how to query the statistical period of a metric, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The Prometheus alert configuration.
	//
	// > Set this parameter only when you create a Prometheus alert rule for Hybrid Cloud Monitoring.
	Prometheus *PutResourceMetricRuleRequestPrometheus `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	// The resource information, such as `[{"instanceId":"i-uf6j91r34rnwawoo****"}]` or `[{"userId":"100931896542****"}]`.
	//
	// For information about the supported monitoring dimensions, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// [{"instanceId":"i-uf6j91r34rnwawoo****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// You can enter a new alert rule ID or use the ID of an existing alert rule in CloudMonitor. For information about how to query alert rule IDs, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
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
	// You can enter a new alert rule name or use the name of an existing alert rule in CloudMonitor. For information about how to query alert rule names, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
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
	// The mute period. Unit: seconds. Default value: 86400.
	//
	// > The mute period specifies the interval at which an alert notification is re-sent if the alert does not recover to Normal.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The callback URL to which a POST request is sent when an alert is triggered.
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
	// Critical级别阈值比较符。取值：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// Critical级别报警统计方法。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Critical级别报警阈值。
	//
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Critical级别报警重试次数。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
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
	// Info级别阈值比较符。取值：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// Info级别报警统计方法。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Info级别报警阈值。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Info级别报警重试次数。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
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
	// Warn级别阈值比较符。取值：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// Warn级别报警统计方法。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Warn级别报警阈值。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Warn级别报警重试次数。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
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
	// The list of alert conditions created in standard mode.
	ExpressionList []*PutResourceMetricRuleRequestCompositeExpressionExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Repeated"`
	// The relationship between multi-metric alert conditions. Valid values:
	//
	// - `&&`: An alert is triggered only when all metrics meet the alert conditions. An alert is triggered only when all expressions in ExpressionList evaluate to `true`.
	//
	// - `||`: An alert is triggered when any metric meets the alert conditions.
	//
	// example:
	//
	// ||
	ExpressionListJoin *string `json:"ExpressionListJoin,omitempty" xml:"ExpressionListJoin,omitempty"`
	// The alert condition created by using an expression. The following scenarios are supported:
	//
	// - Set an alert blacklist for specific resources. For example, `$instanceId != \\"i-io8kfvcpp7x5****\\" ``&&`` $Average > 50` specifies that no alert is triggered for instance `i-io8kfvcpp7x5****` even if its `Average` exceeds 50.
	//
	// - Set a special alert threshold for a specific instance in the rule. For example, `$Average > ($instanceId == \\"i-io8kfvcpp7x5****\\"? 80: 50)` specifies that an alert is triggered for instance `i-io8kfvcpp7x5****` only when its `Average` exceeds 80, while an alert is triggered for other instances when their `Average` exceeds 50.
	//
	// - Limit the number of instances that exceed the threshold. For example, `count($Average > 20) > 3` specifies that an alert is triggered only when more than three instances have an `Average` greater than 20.
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
	// - INFO: information.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of times that the alert condition must be met before an alert notification is sent.
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
	// The comparison operator for the threshold. Valid values:
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
	// - $Maximum: maximum value.
	//
	// - $Minimum: minimum value.
	//
	// - $Average: average value.
	//
	// - $Availability: active rate (typically used for site monitoring).
	//
	// > `$` is the unified prefix for metrics. For information about supported Alibaba Cloud services, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// > The label value supports template parameters. Template parameters are replaced with actual label values.
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
	// The annotations for Prometheus alerting. The annotation keys and values are rendered to help you understand the metric or alert rule.
	//
	// > This feature is equivalent to the Annotation feature in Prometheus.
	Annotations []*PutResourceMetricRuleRequestPrometheusAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// The alert level. Valid values:
	//
	// - CRITICAL: critical.
	//
	// - WARN: warning.
	//
	// - INFO: information.
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The PromQL query statement.
	//
	// > The data obtained by the PromQL query statement is the alert data. Include the alert threshold in this statement.
	//
	// example:
	//
	// cpuUsage{instanceId="xxxx"}[1m]>90
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The number of times that the alert condition must be met before an alert notification is sent.
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
