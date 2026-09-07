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
	// The alert rule object.
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
	if s.AlertRule != nil {
		if err := s.AlertRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateAlertRuleResponseBodyAlertRule struct {
	// The check type of the Prometheus alert rule.
	//
	// - `STATIC`: The alert is triggered based on a static threshold.
	//
	// - `CUSTOM`: The alert is triggered based on a custom PromQL expression.
	//
	// example:
	//
	// STATIC
	AlertCheckType *string `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	// The alert group for the Prometheus alert rule.
	//
	// - `-1`: Custom PromQL
	//
	// - `1`: Kubernetes Workloads
	//
	// - `15`: Kubernetes Nodes
	//
	// example:
	//
	// -1
	AlertGroup *int64 `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The ID of the alert rule.
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
	// The content of the alert rule. This applies to application monitoring and browser monitoring.
	AlertRuleContent *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty" type:"Struct"`
	// The status of the alert rule.
	//
	// - `RUNNING`: The alert rule is running.
	//
	// - `STOPPED`: The alert rule is stopped.
	//
	// - `PAUSED`: The alert rule is paused.
	//
	// > The `PAUSED` status indicates that the system has automatically suspended the alert rule due to an abnormality. This can happen if the alert rule generates too many distinct time series or its associated cluster is deleted.
	//
	// example:
	//
	// RUNNING
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// - `APPLICATION_MONITORING_ALERT_RULE`: an alert rule for application monitoring.
	//
	// - `BROWSER_MONITORING_ALERT_RULE`: an alert rule for browser monitoring.
	//
	// - `PROMETHEUS_MONITORING_ALERT_RULE`: an alert rule for Prometheus monitoring.
	//
	// example:
	//
	// APPLICATION_MONITORING_ALERT_RULE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The annotations of the Prometheus alert rule.
	Annotations []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// Indicates whether newly created applications are automatically added to the alert rule. This applies to application monitoring and browser monitoring rules.
	//
	// - `true`: Enabled
	//
	// - `false`: Disabled
	//
	// example:
	//
	// false
	AutoAddNewApplication *bool `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	// The ID of the cluster that is associated with the Prometheus alert rule.
	//
	// example:
	//
	// ceba9b9ea5b924dd0b6726d2de6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The UNIX timestamp, in milliseconds, when the alert rule was created.
	//
	// example:
	//
	// 1641438611000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The duration, in minutes, for which a condition must be true before an alert is triggered. This applies only to Prometheus alert rules.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The extended fields, returned as a JSON string.
	//
	// example:
	//
	// {\\"alarmContext\\":\\"{\\\\\\"content\\\\\\":\\\\\\"报警名称:$报警名称\\\\\\\\n筛选条件: $筛选\\\\\\\\n报警时间: $报警时间\\\\\\\\n报警内容: $报警内容\\\\\\\\n注意！：该报警未收到恢复邮件之前，正在持续报警中，24小时后会再次提醒您！\\\\\\",\\\\\\"subTitle\\\\\\":\\\\\\"\\\\\\"}\\",\\"alertWays\\":\\"[0,1]\\",\\"contactGroupIds\\":\\"381,5075\\",\\"notice\\":\\"{\\\\\\"endTime\\\\\\":1480607940000,\\\\\\"noticeEndTime\\\\\\":1480607940000,\\\\\\"noticeStartTime\\\\\\":1480521600000,\\\\\\"startTime\\\\\\":1480521600000}\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The filters of the alert rule. This applies to application monitoring or browser monitoring.
	Filters *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The labels of the Prometheus alert rule.
	Labels []*CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The severity level of the Prometheus alert rule.
	//
	// - `P1`: Critical. Indicates major issues that affect core business availability and can have severe consequences.
	//
	// - `P2`: Warning. Indicates issues that impact system availability but have a limited scope.
	//
	// - `P3`: Info. Indicates potential issues or alerts from less critical services.
	//
	// - `P4`: Low priority. Indicates informational alerts that do not affect services.
	//
	// - `Default`: The default level used when no specific severity is required.
	//
	// example:
	//
	// P2
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The message of the Prometheus alert rule.
	//
	// example:
	//
	// 命名空间: {{$labels.namespace}} / Pod: {{$labels.pod_name}} / 容器: {{$labels.container}} 内存使用率超过80%, 当前值{{ printf \\\\\\"%.2f\\\\\\" $value }}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metric type of the alert rule. This applies to application monitoring and browser monitoring.
	//
	// example:
	//
	// JVM
	MetricsType *string `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	// The notification mode.
	//
	// example:
	//
	// NORMAL_MODE
	NotifyMode *string `json:"NotifyMode,omitempty" xml:"NotifyMode,omitempty"`
	// The notification policy.
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The PIDs of the applications associated with the alert rule. This applies to application monitoring and browser monitoring rules.
	Pids []*string `json:"Pids,omitempty" xml:"Pids,omitempty" type:"Repeated"`
	// The PromQL expression for the Prometheus alert rule.
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
	// The tags that are added to the alert rule.
	Tags []*CreateOrUpdateAlertRuleResponseBodyAlertRuleTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The UNIX timestamp, in milliseconds, when the alert rule was last updated.
	//
	// example:
	//
	// 1641438611000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The user ID.
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
	if s.AlertRuleContent != nil {
		if err := s.AlertRuleContent.Validate(); err != nil {
			return err
		}
	}
	if s.Annotations != nil {
		for _, item := range s.Annotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent struct {
	// The alert conditions. This applies to application monitoring and browser monitoring alert rules.
	AlertRuleItems []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems `json:"AlertRuleItems,omitempty" xml:"AlertRuleItems,omitempty" type:"Repeated"`
	// The logical operator for combining multiple alert conditions. This applies to application monitoring and browser monitoring.
	//
	// - `OR`: The alert is triggered if any condition is met.
	//
	// - `AND`: The alert is triggered only if all conditions are met.
	//
	// example:
	//
	// OR
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
	if s.AlertRuleItems != nil {
		for _, item := range s.AlertRuleItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems struct {
	// The aggregation method for the alert condition.
	//
	// - `AVG`: average
	//
	// - `SUM`: sum
	//
	// - `MAX`: maximum
	//
	// - `MIN`: minimum
	//
	// example:
	//
	// AVG
	Aggregate *string `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	// The metric that is evaluated by the alert condition.
	//
	// example:
	//
	// JVM非堆总使用内存量
	MetricKey *string `json:"MetricKey,omitempty" xml:"MetricKey,omitempty"`
	// The duration of the time window, in minutes, for evaluating the alert condition.
	//
	// example:
	//
	// 1
	N *float32 `json:"N,omitempty" xml:"N,omitempty"`
	// The operator used to compare the aggregated metric value with the threshold.
	//
	// - `CURRENT_GTE`: greater than or equal to
	//
	// - `CURRENT_LTE`: less than or equal to
	//
	// - `PREVIOUS_UP`: period-over-period increase percentage
	//
	// - `PREVIOUS_DOWN`: period-over-period decrease percentage
	//
	// - `HOH_UP`: hour-over-hour increase percentage
	//
	// - `HOH_DOWN`: hour-over-hour decrease percentage
	//
	// - `DOD_UP`: day-over-day increase percentage
	//
	// - `DOD_DOWN`: day-over-day decrease percentage
	//
	// example:
	//
	// CURRENT_GTE
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold for the alert condition.
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
	// The annotation key.
	//
	// example:
	//
	// 123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The annotation value.
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
	// The custom filter conditions for the browser monitoring alert rule.
	CustomSLSFilters []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters `json:"CustomSLSFilters,omitempty" xml:"CustomSLSFilters,omitempty" type:"Repeated"`
	// The aggregation dimensions.
	CustomSLSGroupByDimensions []*string `json:"CustomSLSGroupByDimensions,omitempty" xml:"CustomSLSGroupByDimensions,omitempty" type:"Repeated"`
	// The configured filter conditions.
	CustomSLSWheres []*string `json:"CustomSLSWheres,omitempty" xml:"CustomSLSWheres,omitempty" type:"Repeated"`
	// The filter conditions of the alert rule. This applies to application monitoring or browser monitoring.
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
	if s.CustomSLSFilters != nil {
		for _, item := range s.CustomSLSFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DimFilters != nil {
		for _, item := range s.DimFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// username
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator for the filter condition.
	//
	// - `=`: equals
	//
	// - `not`: not equal to
	//
	// example:
	//
	// =
	Opt *string `json:"Opt,omitempty" xml:"Opt,omitempty"`
	// Indicates whether the filter condition is displayed on the console.
	//
	// example:
	//
	// false
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
	// Used exclusively to distinguish between log types in browser monitoring. This parameter does not apply to other filter conditions.
	//
	// example:
	//
	// null
	T *string `json:"T,omitempty" xml:"T,omitempty"`
	// The value for the filter condition.
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
	// The operator for the filter condition.
	//
	// example:
	//
	// ALL
	FilterOpt *string `json:"FilterOpt,omitempty" xml:"FilterOpt,omitempty"`
	// The values for the filter condition.
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
	// The label key.
	//
	// example:
	//
	// 123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The label value.
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
