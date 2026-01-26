// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *GetAlertRulesResponseBodyPageBean) *GetAlertRulesResponseBody
	GetPageBean() *GetAlertRulesResponseBodyPageBean
	SetRequestId(v string) *GetAlertRulesResponseBody
	GetRequestId() *string
}

type GetAlertRulesResponseBody struct {
	// The returned pages.
	PageBean *GetAlertRulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 337B8F7E-0A64-5768-9225-E9B3CF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBody) GetPageBean() *GetAlertRulesResponseBodyPageBean {
	return s.PageBean
}

func (s *GetAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertRulesResponseBody) SetPageBean(v *GetAlertRulesResponseBodyPageBean) *GetAlertRulesResponseBody {
	s.PageBean = v
	return s
}

func (s *GetAlertRulesResponseBody) SetRequestId(v string) *GetAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertRulesResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertRulesResponseBodyPageBean struct {
	// The alert rules.
	AlertRules []*GetAlertRulesResponseBodyPageBeanAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	// The number of pages returned.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert rules returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of queried alert rules.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBean) GetAlertRules() []*GetAlertRulesResponseBodyPageBeanAlertRules {
	return s.AlertRules
}

func (s *GetAlertRulesResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *GetAlertRulesResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *GetAlertRulesResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *GetAlertRulesResponseBodyPageBean) SetAlertRules(v []*GetAlertRulesResponseBodyPageBeanAlertRules) *GetAlertRulesResponseBodyPageBean {
	s.AlertRules = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) SetPage(v int64) *GetAlertRulesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) SetSize(v int64) *GetAlertRulesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) SetTotal(v int64) *GetAlertRulesResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) Validate() error {
	if s.AlertRules != nil {
		for _, item := range s.AlertRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlertRulesResponseBodyPageBeanAlertRules struct {
	// The alert check type of the Prometheus alert rule.
	//
	// 	- STATIC: static threshold
	//
	// 	- CUSTOM: custom PromQL
	//
	// example:
	//
	// STATIC
	AlertCheckType *string `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	// The alert contact group ID of the Prometheus alert rule.
	//
	// 	- \\-1: custom PromQL
	//
	// 	- 1: Kubernetes load
	//
	// 	- 15: Kubernetes node
	//
	// example:
	//
	// 1
	AlertGroup *int64 `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The alert rule ID.
	//
	// example:
	//
	// 5730***
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// arms-test
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The content of the Application Monitoring or Browser Monitoring alert rule.
	AlertRuleContent *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty" type:"Struct"`
	// The status of the alert rule. Valid values:
	//
	// 	- RUNNING
	//
	// 	- STOPPED
	//
	// 	- PAUSED
	//
	// >  The PAUSED state indicates that the alert rule is abnormal and has been suspended. This may be because the specified threshold value is excessively large, or the associated cluster has been deleted.
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
	// 	- PROMETHEUS_MONITORING_ALERT_RULE: Prometheus alert rule
	//
	// example:
	//
	// APPLICATION_MONITORING_ALERT_RULE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The annotations of the Prometheus alert rule.
	Annotations []*GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// Indicates whether the alert rule is applied to new applications that are created in Application Monitoring or Browser Monitoring. Valid values:
	//
	// 	- `true`: yes
	//
	// 	- `false`: no
	//
	// example:
	//
	// false
	AutoAddNewApplication *bool `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	// The cluster ID of the Prometheus alert rule.
	//
	// example:
	//
	// ceba9b9ea5b924dd0b6726d2de6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the alert rule was created. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640333981000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The duration of the Prometheus alert rule.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The extended fields.
	//
	// >  For existing Application Monitoring alert rules, the fields contain information such as contacts, alert template, and notification content.
	//
	// example:
	//
	// {\\\\"alarmContext\\\\":\\\\"{\\\\\\\\\\"content\\\\\\\\\\":\\\\\\\\Alert name: $Alert name\\\\\\\\\\\\nFilter condition: $Filter condition\\\\\\\\\\\\nAlert time: $Alert time\\\\\\\\\\\\nAlert content: $Alert content\\\\\\\\\\\\nNote: The alert persists before you receive an email that reminds you to clear the alert. You will be reminded of the alert again 24 hours later. \\\\\\\\\\",\\\\\\\\\\"subTitle\\\\\\\\\\":\\\\\\\\\\"\\\\\\\\\\"}\\\\",\\\\"alertWays\\\\":\\\\"[0,1]\\\\",\\\\"contactGroupIds\\\\":\\\\"381,5075\\\\",\\\\"notice\\\\":\\\\"{\\\\\\\\\\"endTime\\\\\\\\\\":1480607940000,\\\\\\\\\\"noticeEndTime\\\\\\\\\\":1480607940000,\\\\\\\\\\"noticeStartTime\\\\\\\\\\":1480521600000,\\\\\\\\\\"startTime\\\\\\\\\\":1480521600000}\\\\"}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The filter conditions of the Application Monitoring or Browser Monitoring alert rule.
	Filters *GetAlertRulesResponseBodyPageBeanAlertRulesFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The tags of the Prometheus alert rule.
	Labels []*GetAlertRulesResponseBodyPageBeanAlertRulesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
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
	// Namespace: {{$labels.namespace}} / Pod: {{$labels.pod_name}} / Container: {{$labels.container}} CPU usage: {{$labels.metrics_params_opt_label_value}} {{$labels.metrics_params_value}}%. Current value: {{ printf "%.2f" $value }}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The metric type of the Application Monitoring or Browser Monitoring alert rule.
	//
	// example:
	//
	// JVM
	MetricsType *string `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// ALERT_MANAGER
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The process ID (PID) of the application to which the Application Monitoring or Browser Monitoring alert rule is applied.
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
	// The tags of the alert rule.
	Tags []*GetAlertRulesResponseBodyPageBeanAlertRulesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the alert rule was updated. The value is a timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1640333981000
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1131971649******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertCheckType() *string {
	return s.AlertCheckType
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertGroup() *int64 {
	return s.AlertGroup
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertId() *int64 {
	return s.AlertId
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertName() *string {
	return s.AlertName
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertRuleContent() *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent {
	return s.AlertRuleContent
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAlertType() *string {
	return s.AlertType
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAnnotations() []*GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations {
	return s.Annotations
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetAutoAddNewApplication() *bool {
	return s.AutoAddNewApplication
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetDuration() *string {
	return s.Duration
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetExtend() *string {
	return s.Extend
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetFilters() *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	return s.Filters
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetLabels() []*GetAlertRulesResponseBodyPageBeanAlertRulesLabels {
	return s.Labels
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetLevel() *string {
	return s.Level
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetMessage() *string {
	return s.Message
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetMetricsType() *string {
	return s.MetricsType
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetPids() []*string {
	return s.Pids
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetPromQL() *string {
	return s.PromQL
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetTags() []*GetAlertRulesResponseBodyPageBeanAlertRulesTags {
	return s.Tags
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) GetUserId() *string {
	return s.UserId
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertCheckType(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertCheckType = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertGroup(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertGroup = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertId(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertName(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertName = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertRuleContent(v *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRuleContent = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertStatus(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertStatus = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAnnotations(v []*GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Annotations = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAutoAddNewApplication(v bool) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AutoAddNewApplication = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetClusterId(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.ClusterId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetCreatedTime(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.CreatedTime = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetDuration(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Duration = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetExtend(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Extend = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetFilters(v *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Filters = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetLabels(v []*GetAlertRulesResponseBodyPageBeanAlertRulesLabels) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Labels = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetLevel(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Level = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetMessage(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Message = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetMetricsType(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricsType = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetNotifyStrategy(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.NotifyStrategy = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetPids(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Pids = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetPromQL(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.PromQL = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetTags(v []*GetAlertRulesResponseBodyPageBeanAlertRulesTags) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Tags = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetUpdatedTime(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdatedTime = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) Validate() error {
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

type GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent struct {
	// The trigger conditions of the Application Monitoring or Browser Monitoring alert rule.
	AlertRuleItems []*GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems `json:"AlertRuleItems,omitempty" xml:"AlertRuleItems,omitempty" type:"Repeated"`
	// The relationship between multiple alert conditions specified for the Application Monitoring or Browser Monitoring alert rule. Valid values:
	//
	// 	- OR: The alert rule is triggered if one of the conditions is met.
	//
	// 	- AND: The alert rule is triggered if all the conditions are met.
	//
	// example:
	//
	// OR
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) GetAlertRuleItems() []*GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	return s.AlertRuleItems
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) GetCondition() *string {
	return s.Condition
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) SetAlertRuleItems(v []*GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent {
	s.AlertRuleItems = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) SetCondition(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent {
	s.Condition = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) Validate() error {
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

type GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems struct {
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
	// The last N minutes.
	//
	// example:
	//
	// 1
	N *int64 `json:"N,omitempty" xml:"N,omitempty"`
	// The operator that is used to compare the metric value with the threshold. Valid values:
	//
	// 	- CURRENT_GTE: greater than or equal to
	//
	// 	- CURRENT_LTE: less than or equal to
	//
	// 	- PREVIOUS_UP: increase in percentage compared with the previous period
	//
	// 	- PREVIOUS_DOWN: decrease in percentage compared with the previous period
	//
	// 	- HOH_UP: increase in percentage compared with the same period in the previous hour
	//
	// 	- HOH_DOWN: decrease in percentage compared with the same period in the previous hour
	//
	// 	- DOD_UP: increase in percentage compared with the same period in the previous day
	//
	// 	- DOD_DOWN: decrease in percentage compared with the same period in the previous day
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

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GetAggregate() *string {
	return s.Aggregate
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GetMetricKey() *string {
	return s.MetricKey
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GetN() *int64 {
	return s.N
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GetOperator() *string {
	return s.Operator
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GetValue() *string {
	return s.Value
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetAggregate(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.Aggregate = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetMetricKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.MetricKey = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetN(v int64) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.N = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetOperator(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.Operator = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.Value = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) Validate() error {
	return dara.Validate(s)
}

type GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations struct {
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

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) GetName() *string {
	return s.Name
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) GetValue() *string {
	return s.Value
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) SetName(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations {
	s.Name = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations {
	s.Value = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) Validate() error {
	return dara.Validate(s)
}

type GetAlertRulesResponseBodyPageBeanAlertRulesFilters struct {
	// The custom filter condition of the Browser Monitoring alert rule.
	CustomSLSFilters []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters `json:"CustomSLSFilters,omitempty" xml:"CustomSLSFilters,omitempty" type:"Repeated"`
	// The information about the aggregation dimension.
	CustomSLSGroupByDimensions []*string `json:"CustomSLSGroupByDimensions,omitempty" xml:"CustomSLSGroupByDimensions,omitempty" type:"Repeated"`
	// The details of the custom filter condition.
	CustomSLSWheres []*string `json:"CustomSLSWheres,omitempty" xml:"CustomSLSWheres,omitempty" type:"Repeated"`
	// The information about each filter condition of the Application Monitoring or Browser Monitoring alert rule.
	DimFilters []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters `json:"DimFilters,omitempty" xml:"DimFilters,omitempty" type:"Repeated"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFilters) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFilters) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) GetCustomSLSFilters() []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	return s.CustomSLSFilters
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) GetCustomSLSGroupByDimensions() []*string {
	return s.CustomSLSGroupByDimensions
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) GetCustomSLSWheres() []*string {
	return s.CustomSLSWheres
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) GetDimFilters() []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	return s.DimFilters
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetCustomSLSFilters(v []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.CustomSLSFilters = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetCustomSLSGroupByDimensions(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.CustomSLSGroupByDimensions = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetCustomSLSWheres(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.CustomSLSWheres = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetDimFilters(v []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.DimFilters = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) Validate() error {
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

type GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters struct {
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
	// Indicates whether this filter condition is displayed on the frontend.
	//
	// example:
	//
	// false
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
	// The log type of Browser Monitoring. This field is not included in other filter conditions.
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

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GetKey() *string {
	return s.Key
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GetOpt() *string {
	return s.Opt
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GetShow() *bool {
	return s.Show
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GetT() *string {
	return s.T
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GetValue() *string {
	return s.Value
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Key = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetOpt(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Opt = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetShow(v bool) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Show = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetT(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.T = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Value = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) Validate() error {
	return dara.Validate(s)
}

type GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters struct {
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

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) GetFilterKey() *string {
	return s.FilterKey
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) GetFilterOpt() *string {
	return s.FilterOpt
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) GetFilterValues() []*string {
	return s.FilterValues
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) SetFilterKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	s.FilterKey = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) SetFilterOpt(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	s.FilterOpt = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) SetFilterValues(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	s.FilterValues = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) Validate() error {
	return dara.Validate(s)
}

type GetAlertRulesResponseBodyPageBeanAlertRulesLabels struct {
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

func (s GetAlertRulesResponseBodyPageBeanAlertRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesLabels) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) GetName() *string {
	return s.Name
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) GetValue() *string {
	return s.Value
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) SetName(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesLabels {
	s.Name = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesLabels {
	s.Value = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) Validate() error {
	return dara.Validate(s)
}

type GetAlertRulesResponseBodyPageBeanAlertRulesTags struct {
	// The tag key.
	//
	// example:
	//
	// type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesTags) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesTags) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesTags) GetKey() *string {
	return s.Key
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesTags) GetValue() *string {
	return s.Value
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesTags) SetKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesTags {
	s.Key = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesTags) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesTags {
	s.Value = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesTags) Validate() error {
	return dara.Validate(s)
}
