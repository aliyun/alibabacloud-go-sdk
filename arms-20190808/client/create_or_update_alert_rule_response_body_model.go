// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertRule(v *CreateOrUpdateAlertRuleResponseBodyAlertRule) *CreateOrUpdateAlertRuleResponseBody
	GetAlertRule() *CreateOrUpdateAlertRuleResponseBodyAlertRule
	SetRequestId(v string) *CreateOrUpdateAlertRuleResponseBody
	GetRequestId() *string
}

type CreateOrUpdateAlertRuleResponseBody struct {
	// The details of the alert rule.
	AlertRule *CreateOrUpdateAlertRuleResponseBodyAlertRule `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 337B8F7E-0A64-5768-9225-E9B3CF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBody) GetAlertRule() *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	return s.AlertRule
}

func (s *CreateOrUpdateAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateAlertRuleResponseBody) SetAlertRule(v *CreateOrUpdateAlertRuleResponseBodyAlertRule) *CreateOrUpdateAlertRuleResponseBody {
	s.AlertRule = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBody) SetRequestId(v string) *CreateOrUpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRule struct {
	// The alert check type of the Prometheus alert rule. Valid values:
	//
	// 	- STATIC: a static threshold value.
	//
	// 	- CUSTOM: a custom PromQL statement.
	//
	// example:
	//
	// STATIC
	AlertCheckType *string `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	// The alert contact group ID of the Prometheus alert rule. Valid values:
	//
	// 	- \\-1: custom PromQL
	//
	// 	- 1: Kubernetes load
	//
	// 	- 15: Kubernetes node
	//
	// example:
	//
	// -1
	AlertGroup *int64 `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The alert rule ID.
	//
	// example:
	//
	// 5510445
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// arms-test
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The content of the Application Monitoring or Browser Monitoring alert rule.
	AlertRuleContent *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty" type:"Struct"`
	// The status of the alert rule. Valid values:
	//
	// 	- RUNNING
	//
	// 	- STOPPED
	//
	// 	- PAUSED
	//
	// > The PAUSED status indicates that the alert rule is abnormal and is actively paused by the system. The alert rule may be paused because that it is not unique or the associated cluster has been deleted.
	//
	// example:
	//
	// RUNNING
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// 	- APPLICATION_MONITORING_ALERT_RULE: alert rule for Application Monitoring
	//
	// 	- BROWSER_MONITORING_ALERT_RULE: alert rule for Browser Monitoring
	//
	// 	- PROMETHEUS_MONITORING_ALERT_RULE: alert rule for Prometheus Service
	//
	// example:
	//
	// APPLICATION_MONITORING_ALERT_RULE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The annotations of the Prometheus alert rule.
	Annotations []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// Indicates whether the alert rule was applied to new applications that were created in Application Monitoring or Browser Monitoring. Valid values:
	//
	// 	- `true`: enables the health check feature.
	//
	// 	- `false`: disables the automatic backup feature.
	//
	// example:
	//
	// false
	AutoAddNewApplication *bool `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	// The ID of the monitored cluster.
	//
	// example:
	//
	// ceba9b9ea5b924dd0b6726d2de6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The timestamp generated when the alert rule was created. Unit: seconds.
	//
	// example:
	//
	// 1641438611000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The duration of the Prometheus alert rule. Unit: minutes.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The extended fields.
	//
	// example:
	//
	// {\\\\"alarmContext\\\\":\\\\"{\\\\\\\\\\"content\\\\\\\\\\":\\\\\\\\Alert name: $Alert name\\\\\\\\\\\\nFilter condition: $Filter condition\\\\\\\\\\\\nAlert time: $Alert time\\\\\\\\\\\\nAlert content: $Alert content\\\\\\\\\\\\nNote: The alert persists before you receive an email that reminds you to clear the alert. You will be reminded of the alert again 24 hours later. \\\\\\\\\\",\\\\\\\\\\"subTitle\\\\\\\\\\":\\\\\\\\\\"\\\\\\\\\\"}\\\\",\\\\"alertWays\\\\":\\\\"[0,1]\\\\",\\\\"contactGroupIds\\\\":\\\\"381,5075\\\\",\\\\"notice\\\\":\\\\"{\\\\\\\\\\"endTime\\\\\\\\\\":1480607940000,\\\\\\\\\\"noticeEndTime\\\\\\\\\\":1480607940000,\\\\\\\\\\"noticeStartTime\\\\\\\\\\":1480521600000,\\\\\\\\\\"startTime\\\\\\\\\\":1480521600000}\\\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The filter conditions of the Application Monitoring or Browser Monitoring alert rule.
	Filters *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The tags of the Prometheus alert rule.
	Labels []*CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The severity level of the Prometheus alert rule.
	//
	// 	- P1: Alert notifications are sent for major issues that affect the availability of core business, have a huge impact, and may lead to serious consequences.
	//
	// 	- P2: Alert notifications are sent for service errors that affect the system availability with relatively limited impact.
	//
	// 	- P3: Alert notifications are sent for issues that may cause service errors or negative effects, or alert notifications for services that are relatively less important.
	//
	// 	- P4: Alert notifications are sent for low-priority issues that do not affect your business.
	//
	// 	- Default: Alert notifications are sent regardless of alert levels.
	//
	// example:
	//
	// P2
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The alert message of the Prometheus alert rule.
	//
	// example:
	//
	// Namespace: {{$labels.namespace}} / Pod: {{$labels.pod_name}} / Container: {{$labels.container}} Memory usage exceeds 80%. Current value: {{ printf \\\\\\\\\\"%.2f\\\\\\\\\\" $value }}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metric type of the Application Monitoring or Browser Monitoring alert rule.
	//
	// example:
	//
	// JVM
	MetricsType *string `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	// Notification Mode.
	//
	// example:
	//
	// NORMAL_MODE
	NotifyMode *string `json:"NotifyMode,omitempty" xml:"NotifyMode,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The process ID (PID) that was associated with the Application Monitoring or Browser Monitoring alert rule.
	Pids []*string `json:"Pids,omitempty" xml:"Pids,omitempty" type:"Repeated"`
	// The PromQL statement of the Prometheus alert rule.
	//
	// example:
	//
	// node_memory_MemAvailable_bytes{} / node_memory_MemTotal_bytes{} 	- 100
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags.
	Tags []*CreateOrUpdateAlertRuleResponseBodyAlertRuleTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timestamp generated when the alert rule was updated. Unit: seconds.
	//
	// example:
	//
	// 1641438611000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1131971649******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRule) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertCheckType() *string {
	return s.AlertCheckType
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertGroup() *int64 {
	return s.AlertGroup
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertId() *int64 {
	return s.AlertId
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertName() *string {
	return s.AlertName
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertRuleContent() *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent {
	return s.AlertRuleContent
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAlertType() *string {
	return s.AlertType
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAnnotations() []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations {
	return s.Annotations
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetAutoAddNewApplication() *bool {
	return s.AutoAddNewApplication
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetDuration() *string {
	return s.Duration
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetExtend() *string {
	return s.Extend
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetFilters() *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	return s.Filters
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetLabels() []*CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels {
	return s.Labels
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetLevel() *string {
	return s.Level
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetMessage() *string {
	return s.Message
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetMetricsType() *string {
	return s.MetricsType
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetNotifyMode() *string {
	return s.NotifyMode
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetPids() []*string {
	return s.Pids
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetPromQL() *string {
	return s.PromQL
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetTags() []*CreateOrUpdateAlertRuleResponseBodyAlertRuleTags {
	return s.Tags
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) GetUserId() *string {
	return s.UserId
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertCheckType(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertCheckType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertGroup(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertGroup = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertId(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertName(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertName = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertRuleContent(v *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertRuleContent = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertStatus(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertStatus = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertType(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAnnotations(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Annotations = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAutoAddNewApplication(v bool) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AutoAddNewApplication = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetClusterId(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.ClusterId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetCreatedTime(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.CreatedTime = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetDuration(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Duration = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetExtend(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Extend = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetFilters(v *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Filters = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetLabels(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Labels = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetLevel(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Level = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetMessage(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetMetricsType(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.MetricsType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetNotifyMode(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.NotifyMode = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetNotifyStrategy(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.NotifyStrategy = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetPids(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Pids = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetPromQL(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.PromQL = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetRegionId(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetTags(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Tags = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetUpdatedTime(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.UpdatedTime = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetUserId(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.UserId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent struct {
	// The trigger conditions of the Application Monitoring or Browser Monitoring alert rule.
	AlertRuleItems []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems `json:"AlertRuleItems,omitempty" xml:"AlertRuleItems,omitempty" type:"Repeated"`
	// The relationship between multiple alert conditions that were specified for the Application Monitoring or Browser Monitoring alert rule. Valid values:
	//
	// 	- OR: meets any of the specified conditions.
	//
	// 	- AND: meets all the specified conditions.
	//
	// example:
	//
	// "|"
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) GetAlertRuleItems() []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	return s.AlertRuleItems
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) GetCondition() *string {
	return s.Condition
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) SetAlertRuleItems(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent {
	s.AlertRuleItems = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) SetCondition(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent {
	s.Condition = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems struct {
	// The aggregation method of the alert condition. Valid values:
	//
	// 	- AVG: calculates the average value
	//
	// 	- SUM: calculates the total value
	//
	// 	- MAX: selects the maximum value
	//
	// 	- MIN: selects the minimum value
	//
	// example:
	//
	// AVG
	Aggregate *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	// The metric of the alert condition.
	//
	// example:
	//
	// appstat.jvm.non_heap_used
	MetricKey *string `json:"MetricKey,omitempty" xml:"MetricKey,omitempty"`
	// Indicates the last N minutes.
	//
	// example:
	//
	// 1
	N *float32 `json:"N,omitempty" xml:"N,omitempty"`
	// The comparison operator that was used to compare the metric value with the threshold. Valid values:
	//
	// 	- CURRENT_GTE: greater than or equal to
	//
	// 	- CURRENT_LTE: less than or equal to
	//
	// 	- PREVIOUS_UP: the increase percentage compared with the last period
	//
	// 	- PREVIOUS_DOWN: the decrease percentage compared with the last period
	//
	// 	- HOH_UP: the increase percentage compared with the last hour
	//
	// 	- HOH_DOWN: the decrease percentage compared with the last hour
	//
	// 	- DOD_UP: the increase percentage compared with the last day
	//
	// 	- DOD_DOWN: the decrease percentage compared with the last day
	//
	// example:
	//
	// CURRENT_GTE
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold of the alert condition.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GetAggregate() *string {
	return s.Aggregate
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GetMetricKey() *string {
	return s.MetricKey
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GetN() *float32 {
	return s.N
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GetOperator() *string {
	return s.Operator
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetAggregate(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.Aggregate = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetMetricKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.MetricKey = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetN(v float32) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.N = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetOperator(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations struct {
	// The key of the annotation.
	//
	// example:
	//
	// 123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the annotation.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) SetName(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters struct {
	// The custom filter condition of the Browser Monitoring alert rule.
	CustomSLSFilters []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters `json:"CustomSLSFilters,omitempty" xml:"CustomSLSFilters,omitempty" type:"Repeated"`
	// The information of the aggregation dimension.
	CustomSLSGroupByDimensions []*string `json:"CustomSLSGroupByDimensions,omitempty" xml:"CustomSLSGroupByDimensions,omitempty" type:"Repeated"`
	// The details of the custom filter condition.
	CustomSLSWheres []*string `json:"CustomSLSWheres,omitempty" xml:"CustomSLSWheres,omitempty" type:"Repeated"`
	// The information about each filter condition of the Application Monitoring or Browser Monitoring alert rule.
	DimFilters []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters `json:"DimFilters,omitempty" xml:"DimFilters,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) GetCustomSLSFilters() []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	return s.CustomSLSFilters
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) GetCustomSLSGroupByDimensions() []*string {
	return s.CustomSLSGroupByDimensions
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) GetCustomSLSWheres() []*string {
	return s.CustomSLSWheres
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) GetDimFilters() []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	return s.DimFilters
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetCustomSLSFilters(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.CustomSLSFilters = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetCustomSLSGroupByDimensions(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.CustomSLSGroupByDimensions = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetCustomSLSWheres(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.CustomSLSWheres = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetDimFilters(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.DimFilters = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// username
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The logical operator of the filter condition. Valid values:
	//
	// 	- \\=: equal to
	//
	// 	- not: not equal to
	//
	// example:
	//
	// =
	Opt *string `json:"Opt,omitempty" xml:"Opt,omitempty"`
	// Indicates whether this filter condition was displayed on the frontend.
	//
	// example:
	//
	// false
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
	// The log type of Browser Monitoring. This field was not included in other filter conditions.
	//
	// example:
	//
	// null
	T *string `json:"T,omitempty" xml:"T,omitempty"`
	// The value of the filter condition.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GetKey() *string {
	return s.Key
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GetOpt() *string {
	return s.Opt
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GetShow() *bool {
	return s.Show
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GetT() *string {
	return s.T
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetOpt(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Opt = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetShow(v bool) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Show = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetT(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.T = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// rootIp
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	// The logical operator of the filter condition.
	//
	// example:
	//
	// ALL
	FilterOpt *string `json:"FilterOpt,omitempty" xml:"FilterOpt,omitempty"`
	// The details of the filter condition.
	FilterValues []*string `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) GetFilterKey() *string {
	return s.FilterKey
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) GetFilterOpt() *string {
	return s.FilterOpt
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) GetFilterValues() []*string {
	return s.FilterValues
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) SetFilterKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	s.FilterKey = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) SetFilterOpt(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	s.FilterOpt = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) SetFilterValues(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	s.FilterValues = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels struct {
	// The tag key.
	//
	// example:
	//
	// 123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tag value.
	//
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) SetName(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleTags struct {
	// The tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// John
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) GetKey() *string {
	return s.Key
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) SetKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleTags) Validate() error {
	return dara.Validate(s)
}
