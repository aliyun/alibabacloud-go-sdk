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
	// > Single-metric and multi-metric alert conditions are mutually exclusive and cannot be set at the same time.
	CompositeExpressionShrink *string `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty"`
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
	Labels []*PutResourceMetricRuleShrinkRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
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
	PrometheusShrink *string `json:"Prometheus,omitempty" xml:"Prometheus,omitempty"`
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
	// > The label value supports template parameters. Template parameters are replaced with actual label values.
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
