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
	// The abbreviation of the Alibaba Cloud service name.
	//
	// For information about how to obtain the abbreviation, see the `metricCategory` tag in the `Labels` response parameter of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
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
	// The first-level dimensions of the alert rule in the application group.
	//
	// Format: a collection of `key:value` pairs, such as `{"userId":"120886317861****"}` and `{"instanceId":"i-m5e1qg6uo38rztr4****"}`.
	//
	// example:
	//
	// [{"instanceId":"i-m5e1qg6uo38rztr4****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The effective period during which the alert rule takes effect.
	//
	// example:
	//
	// 05:31-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert email.
	//
	// example:
	//
	// ECS instance
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The second-level or third-level dimensions of the alert rule in the application group.
	//
	// Format: a collection of `key:value` pairs, such as `port:80` and `/dev/xvda:d-m5e6yphgzn3aprwu****`.
	//
	// If the first-level dimension is `{"instanceId":"i-m5e1qg6uo38rztr4****"}`, the second-level dimension is a cloud disk of the instance: `{"/dev/xvda":"d-m5e6yphgzn3aprwu****"}`.
	//
	// example:
	//
	// {"/dev/xvda":"d-m5e6yphgzn3aprwu****"}
	ExtraDimensionJson *string `json:"ExtraDimensionJson,omitempty" xml:"ExtraDimensionJson,omitempty"`
	// The application group ID.
	//
	// For information about how to obtain the application group ID, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 17285****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The detection period of the alert rule. Unit: seconds.
	//
	// > Keep the detection period consistent with the data reporting period. If the detection period is shorter than the data reporting period, alerts may not be triggered due to insufficient data.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The tags of the alert rule.
	//
	// Tags are included in alert notifications.
	Labels []*PutGroupMetricRuleRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name.
	//
	// For information about how to obtain the metric name, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Cloud service monitoring](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// For information about how to obtain the namespace, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Cloud service monitoring](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The processing method when no monitoring data is found. Valid values:
	//
	// - KEEP_LAST_STATE (default): No action is performed.
	//
	// - INSUFFICIENT_DATA: An alert whose content is "Insufficient Data" is triggered.
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
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The advanced settings.
	//
	// Format: {"key1":"value1","key2":"value2"}. Example: {"NotSendOK":true}. This specifies whether to send a notification when the alert is cleared. The key is NotSendOK, and the value is true (do not send) or false (send, which is the default).
	//
	// example:
	//
	// {"NotSendOK":true}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The reporting period of monitoring data.
	//
	// The value of `Period` must be 60 or a multiple of 60. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The alert rule ID.
	//
	// - To create an alert rule for the application group, enter an alert rule ID.
	//
	// - To modify a specified alert rule in the application group, obtain the alert rule ID. For information about how to obtain the alert rule ID, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The alert rule name.
	//
	// - To create an alert rule for the application group, enter an alert rule name.
	//
	// - To modify a specified alert rule in the application group, obtain the alert rule name. For information about how to obtain the alert rule name, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Rule_01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period.
	//
	// Unit: seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The callback URL to which a request is sent when an alert is triggered.
	//
	// Enter a publicly accessible URL. CloudMonitor sends a POST request to push alert information to this URL. Only the HTTP protocol is supported.
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

type PutGroupMetricRuleRequestEscalationsCritical struct {
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
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// Critical级别报警统计方法。多个统计方法之间用半角逗号（,）分隔。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Critical级别报警阈值。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Critical级别报警重试次数。
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
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// Info级别报警统计方法。多个统计方法之间用半角逗号（,）分隔。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Info级别报警阈值。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Info级别报警重试次数。
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
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// Warn级别报警统计方法。多个统计方法之间用半角逗号（,）分隔。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Warn级别报警阈值。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Warn级别报警重试次数。
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
