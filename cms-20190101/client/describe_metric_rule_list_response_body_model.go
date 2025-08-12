// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarms(v *DescribeMetricRuleListResponseBodyAlarms) *DescribeMetricRuleListResponseBody
	GetAlarms() *DescribeMetricRuleListResponseBodyAlarms
	SetCode(v int32) *DescribeMetricRuleListResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMetricRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricRuleListResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeMetricRuleListResponseBody
	GetTotal() *string
}

type DescribeMetricRuleListResponseBody struct {
	// The queried alert rules.
	Alarms *DescribeMetricRuleListResponseBodyAlarms `json:"Alarms,omitempty" xml:"Alarms,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 386C6712-335F-5054-930A-CC92B851ECBA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMetricRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBody) GetAlarms() *DescribeMetricRuleListResponseBodyAlarms {
	return s.Alarms
}

func (s *DescribeMetricRuleListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMetricRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleListResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeMetricRuleListResponseBody) SetAlarms(v *DescribeMetricRuleListResponseBodyAlarms) *DescribeMetricRuleListResponseBody {
	s.Alarms = v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetCode(v int32) *DescribeMetricRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetMessage(v string) *DescribeMetricRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetRequestId(v string) *DescribeMetricRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetSuccess(v bool) *DescribeMetricRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) SetTotal(v string) *DescribeMetricRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMetricRuleListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarms struct {
	Alarm []*DescribeMetricRuleListResponseBodyAlarmsAlarm `json:"Alarm,omitempty" xml:"Alarm,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarms) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarms) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarms) GetAlarm() []*DescribeMetricRuleListResponseBodyAlarmsAlarm {
	return s.Alarm
}

func (s *DescribeMetricRuleListResponseBodyAlarms) SetAlarm(v []*DescribeMetricRuleListResponseBodyAlarmsAlarm) *DescribeMetricRuleListResponseBodyAlarms {
	s.Alarm = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarms) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarm struct {
	// The status of the alert rule. Valid values:
	//
	// 	- OK: The alert rule has no active alerts.
	//
	// 	- ALARM: The alert rule has active alerts.
	//
	// 	- INSUFFICIENT_DATA: No data is available.
	//
	// example:
	//
	// OK
	AlertState *string `json:"AlertState,omitempty" xml:"AlertState,omitempty"`
	// The trigger conditions for multiple metrics.
	//
	// >  The trigger conditions for a single metric and multiple metrics are mutually exclusive. You cannot specify trigger conditions for a single metric and multiple metrics at the same time.
	CompositeExpression *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression `json:"CompositeExpression,omitempty" xml:"CompositeExpression,omitempty" type:"Struct"`
	// The alert contact group.
	//
	// example:
	//
	// ECS_Alarm
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The dimensions of the alert rule.
	//
	// example:
	//
	// [{"instanceId":"i-2ze2d6j5uhg20x47****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The time period during which the alert rule is effective.
	//
	// example:
	//
	// 05:31-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// Indicates whether the alert rule is enabled. Valid values:
	//
	// 	- true: The alert rule is enabled.
	//
	// 	- false: The alert rule is disabled.
	//
	// example:
	//
	// true
	EnableState *bool `json:"EnableState,omitempty" xml:"EnableState,omitempty"`
	// The conditions for triggering different levels of alerts.
	Escalations *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	GmtCreate   *int64                                                    `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtUpdate   *string                                                   `json:"GmtUpdate,omitempty" xml:"GmtUpdate,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the application group.
	//
	// >  If the alert rule is associated with an application group, the name of the application group is returned in this parameter.
	//
	// example:
	//
	// ECS_Group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The tags of the alert rule.
	Labels *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	// The subject of the alert notification email.
	//
	// example:
	//
	// "${serviceType}-${metricName}-${levelDescription}Notification(${dimensions})"
	MailSubject *string `json:"MailSubject,omitempty" xml:"MailSubject,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The method that is used to handle alerts when no monitoring data is found. Valid values:
	//
	// 	- KEEP_LAST_STATE (default value): No operation is performed.
	//
	// 	- INSUFFICIENT_DATA: An alert whose content is "Insufficient data" is triggered.
	//
	// 	- OK: The status is considered normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time period during which the alert rule is ineffective.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// The statistical period.
	//
	// example:
	//
	// 60
	Period          *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ProductCategory *string `json:"ProductCategory,omitempty" xml:"ProductCategory,omitempty"`
	// The Prometheus alerts.
	//
	// >  This parameter is required only if you create a Prometheus alert rule for Hybrid Cloud Monitoring.
	Prometheus *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus `json:"Prometheus,omitempty" xml:"Prometheus,omitempty" type:"Struct"`
	// The resources that are associated with the alert rule.
	//
	// example:
	//
	// [{\\"instanceId\\":\\"i-2ze2d6j5uhg20x47****\\"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// applyTemplate344cfd42-0f32-4fd6-805a-88d7908a****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// Rule_01
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period during which new alert notifications are not sent even if the trigger conditions are met. Unit: seconds. Default value: 86400. Minimum value: 3600.
	//
	// Only one alert is reported during each mute period even if the metric value consecutively exceeds the alert rule threshold several times.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The type of the alert rule. Valid value: METRIC. This value indicates an alert rule for time series metrics.
	//
	// example:
	//
	// METRIC
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The callback URL. CloudMonitor pushes an alert notification to the specified callback URL by sending an HTTP POST request. Only the HTTP protocol is supported.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarm) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetAlertState() *string {
	return s.AlertState
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetCompositeExpression() *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	return s.CompositeExpression
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetDimensions() *string {
	return s.Dimensions
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetEnableState() *bool {
	return s.EnableState
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetEscalations() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	return s.Escalations
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGmtUpdate() *string {
	return s.GmtUpdate
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetLabels() *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels {
	return s.Labels
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetMailSubject() *string {
	return s.MailSubject
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetPrometheus() *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	return s.Prometheus
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetResources() *string {
	return s.Resources
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) GetWebhook() *string {
	return s.Webhook
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetAlertState(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.AlertState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetCompositeExpression(v *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.CompositeExpression = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetContactGroups(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.ContactGroups = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetDimensions(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Dimensions = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEnableState(v bool) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.EnableState = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetEscalations(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Escalations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGmtCreate(v int64) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGmtUpdate(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GmtUpdate = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetGroupName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.GroupName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetLabels(v *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Labels = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMailSubject(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MailSubject = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetMetricName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNamespace(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNoDataPolicy(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.NoDataPolicy = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetNoEffectiveInterval(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.NoEffectiveInterval = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetPeriod(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Period = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetProductCategory(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.ProductCategory = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetPrometheus(v *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Prometheus = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetResources(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Resources = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleId(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleId = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetRuleName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSilenceTime(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SilenceTime = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetSourceType(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.SourceType = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) SetWebhook(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarm {
	s.Webhook = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarm) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression struct {
	// The trigger conditions that are created in standard mode.
	ExpressionList *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Struct"`
	// The relationship between the trigger conditions for multiple metrics. Valid values:
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
	// 	- Limit the number of instances whose metric values exceed the threshold. For example, if you specify `count($Average > 20) > 3`, an alert is triggered only when the number of instances whose `average metric value` exceeds 20 exceeds three.
	//
	// example:
	//
	// $Average > ($instanceId == \\"i-io8kfvcpp7x5****\\"? 80: 50)
	ExpressionRaw *string `json:"ExpressionRaw,omitempty" xml:"ExpressionRaw,omitempty"`
	// The alert level. Valid values:
	//
	// 	- CRITICAL
	//
	// 	- WARN
	//
	// 	- INFO
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

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetExpressionList() *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList {
	return s.ExpressionList
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetExpressionListJoin() *string {
	return s.ExpressionListJoin
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetExpressionRaw() *string {
	return s.ExpressionRaw
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetLevel() *string {
	return s.Level
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetExpressionList(v *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.ExpressionList = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetExpressionListJoin(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.ExpressionListJoin = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetExpressionRaw(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.ExpressionRaw = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetLevel(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.Level = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpression) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList struct {
	ExpressionList []*DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList `json:"ExpressionList,omitempty" xml:"ExpressionList,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) GetExpressionList() []*DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	return s.ExpressionList
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) SetExpressionList(v []*DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList {
	s.ExpressionList = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionList) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList struct {
	// The operator that is used to compare the metric value with the threshold. Valid values:
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
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The statistical method of the metric. Valid values:
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

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetMetricName(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetPeriod(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.Period = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmCompositeExpressionExpressionListExpressionList) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations struct {
	// The conditions for triggering Critical-level alerts.
	Critical *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The conditions for triggering Info-level alerts.
	Info *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	// The conditions for triggering Warn-level alerts.
	Warn *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GetCritical() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	return s.Critical
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GetInfo() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	return s.Info
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) GetWarn() *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	return s.Warn
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetCritical(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Critical = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetInfo(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Info = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) SetWarn(v *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations {
	s.Warn = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalations) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical struct {
	// The comparison operator that is used to compare the metric value with the threshold. Valid values:
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
	// The additional conditions for triggering Critical-level alerts. The additional conditions take effect when the value of the ComparisonOperator parameter is GreaterThanYesterday, LessThanYesterday, GreaterThanLastWeek, LessThanLastWeek, GreaterThanLastPeriod, or LessThanLastPeriod.
	//
	// For example, the values of the PreCondition, ComparisonOperator, and Threshold parameters are set to $Average>80, GreaterThanYesterday, and 10, respectively. An alert is triggered only when the average metric value is greater than 80 and 10% greater than the average metric value at the same time yesterday.
	//
	// >  $Average is a placeholder that consists of `a dollar sign ($) and the statistical method`. CloudMonitor replaces the placeholder with the aggregated value or original value before value comparison.
	//
	// example:
	//
	// $Average>80
	PreCondition *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical methods for Critical-level alerts.
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
	// The consecutive number of times for which the metric value meets the alert condition before a Critical-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo struct {
	// The comparison operator that is used to compare the metric value with the threshold. Valid values:
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
	// The additional conditions for triggering Info-level alerts. The additional conditions take effect when the value of the ComparisonOperator parameter is GreaterThanYesterday, LessThanYesterday, GreaterThanLastWeek, LessThanLastWeek, GreaterThanLastPeriod, or LessThanLastPeriod.
	//
	// For example, the values of the PreCondition, ComparisonOperator, and Threshold parameters are set to $Average>80, GreaterThanYesterday, and 10, respectively. An alert is triggered only when the average metric value is greater than 80 and 10% greater than the average metric value at the same time yesterday.
	//
	// >  $Average is a placeholder that consists of `a dollar sign ($) and the statistical method`. CloudMonitor replaces the placeholder with the aggregated value or original value before value comparison.
	//
	// example:
	//
	// $Average>80
	PreCondition *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical methods for Info-level alerts.
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
	// The consecutive number of times for which the metric value meets the alert condition before an Info-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn struct {
	// The comparison operator that is used to compare the metric value with the threshold. Valid values:
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
	// The additional conditions for triggering Warn-level alerts. The additional conditions take effect when the value of the ComparisonOperator parameter is GreaterThanYesterday, LessThanYesterday, GreaterThanLastWeek, LessThanLastWeek, GreaterThanLastPeriod, or LessThanLastPeriod.
	//
	// For example, the values of the PreCondition, ComparisonOperator, and Threshold parameters are set to $Average>80, GreaterThanYesterday, and 10, respectively. An alert is triggered only when the average metric value is greater than 80 and 10% greater than the average metric value at the same time yesterday.
	//
	// >  $Average is a placeholder that consists of `a dollar sign ($) and the statistical method`. CloudMonitor replaces the placeholder with the aggregated value or original value before value comparison.
	//
	// example:
	//
	// $Average>80
	PreCondition *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// The statistical methods for Warn-level alerts.
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
	// The consecutive number of times for which the metric value meets the alert condition before a Warn-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetPreCondition() *string {
	return s.PreCondition
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetComparisonOperator(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetPreCondition(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetStatistics(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetThreshold(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) SetTimes(v int32) *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmLabels struct {
	Labels []*DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) GetLabels() []*DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels {
	return s.Labels
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) SetLabels(v []*DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels {
	s.Labels = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels struct {
	// The tag key of the alert rule.
	//
	// example:
	//
	// cmsRuleKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the alert rule.
	//
	// example:
	//
	// cmsRuleValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) SetKey(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels {
	s.Key = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) SetValue(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels {
	s.Value = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmLabelsLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus struct {
	// The annotations of the Prometheus alert rule. When a Prometheus alert is triggered, the system renders the annotated keys and values to help you understand the metrics and alert rule.
	//
	// >  This parameter is equivalent to the annotations parameter of open source Prometheus.
	Annotations *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Struct"`
	// The alert level. Valid values:
	//
	// 	- CRITICAL
	//
	// 	- WARN
	//
	// 	- INFO
	//
	// example:
	//
	// Critical
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The PromQL query statement.
	//
	// >  The data obtained by using the PromQL query statement is the monitoring data. You must include the alert threshold in this statement.
	//
	// example:
	//
	// CpuUsage{instanceId="xxxx"}[1m]>90
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The number of consecutive triggers. If the number of times that the metric values meet the trigger conditions reaches the value of this parameter, CloudMonitor sends alert notifications.
	//
	// example:
	//
	// 3
	Times *int64 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetAnnotations() *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations {
	return s.Annotations
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetLevel() *string {
	return s.Level
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetPromQL() *string {
	return s.PromQL
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) GetTimes() *int64 {
	return s.Times
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetAnnotations(v *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.Annotations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetLevel(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.Level = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetPromQL(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.PromQL = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) SetTimes(v int64) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheus) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations struct {
	Annotations []*DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) GetAnnotations() []*DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations {
	return s.Annotations
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) SetAnnotations(v []*DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations {
	s.Annotations = v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotations) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations struct {
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

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) GetKey() *string {
	return s.Key
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) GetValue() *string {
	return s.Value
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) SetKey(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations {
	s.Key = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) SetValue(v string) *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations {
	s.Value = &v
	return s
}

func (s *DescribeMetricRuleListResponseBodyAlarmsAlarmPrometheusAnnotationsAnnotations) Validate() error {
	return dara.Validate(s)
}
