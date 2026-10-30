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
	SetDryRun(v bool) *PutGroupMetricRuleRequest
	GetDryRun() *bool
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
	// For information about how to obtain the abbreviation of an Alibaba Cloud service name, see the `metricCategory` field in the `Labels` response parameter of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
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
	// Format: a collection of `key:value` pairs. Examples: `{"userId":"120886317861****"}` and `{"instanceId":"i-m5e1qg6uo38rztr4****"}`.
	//
	// example:
	//
	// [{"instanceId":"i-m5e1qg6uo38rztr4****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: sends a check request. The system does not bind a security group to the instance. The check items include whether required parameters are specified, the request format, and service limits. If the check fails, the system returns an error. If the check succeeds, the system returns the error code `DryRunOperation`.
	//
	// - **false*	- (default value): sends a normal request. If the check succeeds, the system returns an HTTP 2xx status code and performs the operation.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The effective period of the alert rule.
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
	// The second-level or third-level dimension of the alert rule in the application group.
	//
	// Format: a collection of `key:value` pairs. Examples: `port:80` and `/dev/xvda:d-m5e6yphgzn3aprwu****`.
	//
	// If the first-level dimension of the alert rule is `{"instanceId":"i-m5e1qg6uo38rztr4****"}`, the second-level dimension is the cloud disk in the instance, such as `{"/dev/xvda":"d-m5e6yphgzn3aprwu****"}`.
	//
	// example:
	//
	// {"/dev/xvda":"d-m5e6yphgzn3aprwu****"}
	ExtraDimensionJson *string `json:"ExtraDimensionJson,omitempty" xml:"ExtraDimensionJson,omitempty"`
	// The ID of the application group.
	//
	// For information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 17285****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The detection period of the alert rule. Unit: seconds.
	//
	// > Make sure the detection period of the alert rule is the same as the data reporting period. If the detection period is shorter than the data reporting period, alerts may not be triggered due to insufficient data.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The tags of the alert rule.
	//
	// Alert notifications include tags.
	Labels []*PutGroupMetricRuleRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// For information about how to obtain the name of a metric, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// For information about how to obtain the namespace of an Alibaba Cloud service, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The action to take when no monitoring data is found. Valid values:
	//
	// - KEEP_LAST_STATE (default value): Takes no action.
	//
	// - INSUFFICIENT_DATA: Sends an alert notification indicating that no data is found.
	//
	// - OK: Sends an alert notification indicating that the status is normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time range during which the alert rule is inactive.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The advanced settings.
	//
	// Format: `{"key1":"value1","key2":"value2"}`. Example: `{"NotSendOK":true}`, which specifies whether to send alert recovery notifications. The key is `NotSendOK`, and the value is `true` (do not send) or `false` (send by default).
	//
	// example:
	//
	// {"NotSendOK":true}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The reporting period of monitoring data.
	//
	// The value of `Period` must be 60 or an integer multiple of 60. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the alert rule.
	//
	// - To create an alert rule for the application group, enter an alert rule ID.
	//
	// - To modify a specified alert rule in the application group, you must obtain the alert rule ID. For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// - To create an alert rule for the application group, enter an alert rule name.
	//
	// - To modify a specified alert rule in the application group, you must obtain the alert rule name. For information about how to obtain the name of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Rule_01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute for epoch.
	//
	// Unit: seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The webhook URL that is called when an alert is triggered.
	//
	// Enter a publicly accessible URL. CloudMonitor pushes alert notifications to this URL by sending POST requests. Only HTTP is supported.
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

func (s *PutGroupMetricRuleRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *PutGroupMetricRuleRequest) SetDryRun(v bool) *PutGroupMetricRuleRequest {
	s.DryRun = &v
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
	// The threshold comparison operator for Critical-level alerts. Valid values:
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
	// The statistical method for Critical-level alerts. Separate multiple statistical methods with commas (,).
	//
	// The valid values of this parameter are determined by the `Statistics` column corresponding to the `MetricName` of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For information about how to obtain the valid values of this parameter, see [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The number of retries for Critical-level alerts.
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
	// The threshold comparison operator for Info-level alerts. Valid values:
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
	// The statistical method for Info-level alerts. Separate multiple statistical methods with commas (,).
	//
	// The valid values of this parameter are determined by the `Statistics` column corresponding to the `MetricName` of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For information about how to obtain the valid values of this parameter, see [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The number of retries for Info-level alerts.
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
	// The threshold comparison operator for Warn-level alerts. Valid values:
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
	// The statistical method for Warn-level alerts. Separate multiple statistical methods with commas (,).
	//
	// The valid values of this parameter are determined by the `Statistics` column corresponding to the `MetricName` of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For information about how to obtain the valid values of this parameter, see [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// The number of retries for Warn-level alerts.
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
