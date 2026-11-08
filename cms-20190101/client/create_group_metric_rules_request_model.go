// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *CreateGroupMetricRulesRequest
	GetGroupId() *int64
	SetGroupMetricRules(v []*CreateGroupMetricRulesRequestGroupMetricRules) *CreateGroupMetricRulesRequest
	GetGroupMetricRules() []*CreateGroupMetricRulesRequestGroupMetricRules
	SetRegionId(v string) *CreateGroupMetricRulesRequest
	GetRegionId() *string
}

type CreateGroupMetricRulesRequest struct {
	// The ID of the application group.
	//
	// For more information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3607****
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of metric rules for the application group.
	GroupMetricRules []*CreateGroupMetricRulesRequestGroupMetricRules `json:"GroupMetricRules,omitempty" xml:"GroupMetricRules,omitempty" type:"Repeated"`
	RegionId         *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGroupMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGroupMetricRulesRequest) GetGroupMetricRules() []*CreateGroupMetricRulesRequestGroupMetricRules {
	return s.GroupMetricRules
}

func (s *CreateGroupMetricRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGroupMetricRulesRequest) SetGroupId(v int64) *CreateGroupMetricRulesRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupMetricRulesRequest) SetGroupMetricRules(v []*CreateGroupMetricRulesRequestGroupMetricRules) *CreateGroupMetricRulesRequest {
	s.GroupMetricRules = v
	return s
}

func (s *CreateGroupMetricRulesRequest) SetRegionId(v string) *CreateGroupMetricRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGroupMetricRulesRequest) Validate() error {
	if s.GroupMetricRules != nil {
		for _, item := range s.GroupMetricRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateGroupMetricRulesRequestGroupMetricRules struct {
	Escalations *CreateGroupMetricRulesRequestGroupMetricRulesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The abbreviation of the cloud service name.
	//
	// Valid values of N: 1 to 200.
	//
	// For more information about how to obtain the abbreviation of a cloud service name, see the `metricCategory` field in the `Labels` response parameter of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The alert contact group.
	//
	// Valid values of N: 1 to 200.
	//
	// For more information about how to obtain alert contact groups, see [DescribeContactGroupList](https://help.aliyun.com/document_detail/114922.html).
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The monitoring dimensions of the specified resource.
	//
	// Format: a collection of `key:value` pairs. Examples: `{"userId":"120886317861****"}` and `{"instanceId":"i-2ze2d6j5uhg20x47****"}`.
	//
	// example:
	//
	// [{"instanceId":"i-m5e1qg6uo38rztr4****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The effective period of the alert rule. Valid values of N: 1 to 200.
	//
	// example:
	//
	// 05:31-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert email.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// ECS instance
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The detection period of the alert rule.
	//
	// Valid values of N: 1 to 200.
	//
	// Unit: seconds. The default value is the minimum frequency of the metric.
	//
	// > Set the detection period of the alert rule to the same value as the data reporting period. If the detection period is shorter than the data reporting period, alerts may not be triggered due to insufficient data.
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The tags of the alert rule.
	Labels []*CreateGroupMetricRulesRequestGroupMetricRulesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// Valid values of N: 1 to 200.
	//
	// For more information about how to obtain the name of a metric, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The data namespace of the Alibaba Cloud service.
	//
	// Valid values of N: 1 to 200.
	//
	// For more information about how to obtain the data namespace of an Alibaba Cloud service, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The handling method when no monitoring data is available. Valid values:
	//
	// - KEEP_LAST_STATE: No action is taken. This is the default value.
	//
	// - INSUFFICIENT_DATA: The alert content indicates that no data is available.
	//
	// - OK: The status is normal.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time range during which the alert rule is inactive. Valid values of N: 1 to 200.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The advanced settings.
	//
	// Format: `{"key1":"value1","key2":"value2"}`. Example: `{"NotSendOK":true}`, which specifies whether to send alert recovery notifications. The key is NotSendOK, and the value is true (do not send) or false (send, which is the default value).
	//
	// example:
	//
	// {
	//
	//       "NotSendOK": true
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The reporting period of monitoring data.
	//
	// Valid values of N: 1 to 200.
	//
	// The value of `Period` must be 60 or an integer multiple of 60. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the alert rule.
	//
	// Valid values of N: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 456789
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// Valid values of N: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period for alert notifications. Valid values of N: 1 to 200.
	//
	// Unit: seconds. Default value: 86400. Minimum value: 3600.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The webhook URL that is called when an alert is triggered. Valid values of N: 1 to 200.
	//
	// Specify a publicly accessible URL. CloudMonitor pushes alert notifications to this URL by sending POST requests. Only HTTP is supported.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRules) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRules) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetEscalations() *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	return s.Escalations
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetCategory() *string {
	return s.Category
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetDimensions() *string {
	return s.Dimensions
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetInterval() *string {
	return s.Interval
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetLabels() []*CreateGroupMetricRulesRequestGroupMetricRulesLabels {
	return s.Labels
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetOptions() *string {
	return s.Options
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetPeriod() *string {
	return s.Period
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetWebhook() *string {
	return s.Webhook
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEscalations(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Escalations = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetCategory(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Category = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetContactGroups(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.ContactGroups = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetDimensions(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Dimensions = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEffectiveInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.EffectiveInterval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEmailSubject(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.EmailSubject = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Interval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetLabels(v []*CreateGroupMetricRulesRequestGroupMetricRulesLabels) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Labels = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetMetricName(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.MetricName = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNamespace(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Namespace = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNoDataPolicy(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.NoDataPolicy = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNoEffectiveInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.NoEffectiveInterval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetOptions(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Options = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetPeriod(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Period = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetRuleId(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.RuleId = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetRuleName(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.RuleName = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetSilenceTime(v int32) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.SilenceTime = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetWebhook(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Webhook = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) Validate() error {
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

type CreateGroupMetricRulesRequestGroupMetricRulesEscalations struct {
	Critical *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalations) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GetCritical() *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	return s.Critical
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GetInfo() *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	return s.Info
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GetWarn() *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	return s.Warn
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetCritical(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Critical = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetInfo(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Info = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetWarn(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Warn = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) Validate() error {
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

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical struct {
	// The comparison operator for the critical-level threshold. Valid values:
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
	// - GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// - LessThanYesterday: less than the metric value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the metric value at the same time last week.
	//
	// - LessThanLastWeek: less than the metric value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the metric value in the previous period.
	//
	// - LessThanLastPeriod: less than the metric value in the previous period.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *string `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical method for critical-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// The valid values of this parameter are determined by the `Statistics` column that corresponds to the `MetricName` parameter of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for triggering critical-level alert notifications.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times that the metric must reach the alert threshold to send a critical-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetN() *string {
	return s.N
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetN(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.N = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetPreCondition(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo struct {
	// The comparison operator for the info-level threshold. Valid values:
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
	// - GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// - LessThanYesterday: less than the metric value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the metric value at the same time last week.
	//
	// - LessThanLastWeek: less than the metric value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the metric value in the previous period.
	//
	// - LessThanLastPeriod: less than the metric value in the previous period.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *string `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical method for info-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// The valid values of this parameter are determined by the `Statistics` column that corresponds to the `MetricName` parameter of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for info-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 10
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times that the metric must reach the alert threshold to send an info-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 1
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetN() *string {
	return s.N
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetN(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.N = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetPreCondition(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn struct {
	// The comparison operator for the warning-level threshold. Valid values:
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
	// - GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// - LessThanYesterday: less than the metric value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the metric value at the same time last week.
	//
	// - LessThanLastWeek: less than the metric value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the metric value in the previous period.
	//
	// - LessThanLastPeriod: less than the metric value in the previous period.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *string `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical method for warning-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// The valid values of this parameter are determined by the `Statistics` column that corresponds to the `MetricName` parameter of the specified Alibaba Cloud service. Examples: Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [Alibaba Cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for warning-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 20
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times that the metric must reach the alert threshold to send a warning-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. For the specified alert level, you must specify the Statistics, ComparisonOperator, Threshold, and Times parameters at the same time.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetN() *string {
	return s.N
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetN(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.N = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetPreCondition(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesLabels struct {
	// The tag key of the alert rule. Alert notifications include tags.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the alert rule. Alert notifications include tags.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesLabels) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) GetKey() *string {
	return s.Key
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) GetValue() *string {
	return s.Value
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) SetKey(v string) *CreateGroupMetricRulesRequestGroupMetricRulesLabels {
	s.Key = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) SetValue(v string) *CreateGroupMetricRulesRequestGroupMetricRulesLabels {
	s.Value = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) Validate() error {
	return dara.Validate(s)
}
