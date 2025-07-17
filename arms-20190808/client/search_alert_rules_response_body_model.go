// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchAlertRulesResponseBodyPageBean) *SearchAlertRulesResponseBody
	GetPageBean() *SearchAlertRulesResponseBodyPageBean
	SetRequestId(v string) *SearchAlertRulesResponseBody
	GetRequestId() *string
}

type SearchAlertRulesResponseBody struct {
	// The returned struct.
	PageBean *SearchAlertRulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBody) GetPageBean() *SearchAlertRulesResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAlertRulesResponseBody) SetPageBean(v *SearchAlertRulesResponseBodyPageBean) *SearchAlertRulesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertRulesResponseBody) SetRequestId(v string) *SearchAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAlertRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBean struct {
	// The details of the alert rules.
	AlertRules []*SearchAlertRulesResponseBodyPageBeanAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBean) GetAlertRules() []*SearchAlertRulesResponseBodyPageBeanAlertRules {
	return s.AlertRules
}

func (s *SearchAlertRulesResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAlertRulesResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertRulesResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchAlertRulesResponseBodyPageBean) SetAlertRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRules) *SearchAlertRulesResponseBodyPageBean {
	s.AlertRules = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRules struct {
	// The format of the alert notification.
	AlarmContext *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext `json:"AlarmContext,omitempty" xml:"AlarmContext,omitempty" type:"Struct"`
	// The severity of the alerts. Only the value `WARN` is supported.
	//
	// example:
	//
	// WARN
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The conditions of the alert rule. Multiple conditions are separated by the AND or OR logical operators.
	AlertRule *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	// The name of the alert rule.
	//
	// example:
	//
	// TestAlertRule
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// 	- `1`: custom alert rules that are used to monitor drill-down data sets
	//
	// 	- `3`: custom alert rules that are used to monitor tiled data sets
	//
	// 	- `4`: alert rules that are used to monitor the browser, including the default frontend alert rules
	//
	// 	- `5`: alert rules that are used to monitor applications, including the default application alert rules
	//
	// 	- `6`: the default browser alert rules
	//
	// 	- `7`: the default application alert rules
	//
	// 	- `8`: Tracing Analysis alert rules
	//
	// 	- `101`: Prometheus alert rules
	//
	// example:
	//
	// 4
	AlertType *int32 `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The version of the alert rule. Default value: `1`.
	//
	// example:
	//
	// 1
	AlertVersion *int32 `json:"AlertVersion,omitempty" xml:"AlertVersion,omitempty"`
	// Sending method of alarm notification.
	AlertWays []*string `json:"AlertWays,omitempty" xml:"AlertWays,omitempty" type:"Repeated"`
	// The configuration items of the alert rule. The value is a JSON string.
	//
	// The configuration item **continuous*	- indicates whether alert notifications are continuously sent. Valid values:
	//
	// 	- `true`: Alert notifications are sent every minute.
	//
	// 	- `false`: The alert silence feature is enabled.
	//
	// The configuration item **dataRevision*	- indicates the data revision policy that is used if no data is obtained or the data is null. Default value: 2. Valid values:
	//
	// 	- `0`: overwrites the data by using the value 0
	//
	// 	- `1`: overwrites the data by using the value 1
	//
	// 	- `2`: overwrites the data by using the value null. This value indicates that no alert is triggered if no data exists
	//
	// example:
	//
	// {\\"continuous\\":true,\\"dataRevision\\":2}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the contact group. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 381*,572*
	ContactGroupIdList *string `json:"ContactGroupIdList,omitempty" xml:"ContactGroupIdList,omitempty"`
	// The IDs of the alert contact groups. The value is a JSON array.
	//
	// example:
	//
	// [123, 234]
	ContactGroupIds *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	// The timestamp that shows when the alert rule was created.
	//
	// example:
	//
	// 1579508519683
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the alert is sent through the alert center. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	HostByAlertManager *bool `json:"HostByAlertManager,omitempty" xml:"HostByAlertManager,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the application that is associated with the alert rule.
	MetricParam *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam `json:"MetricParam,omitempty" xml:"MetricParam,omitempty" type:"Struct"`
	// The time ranges when the alert rule takes effect and when alert notifications are sent.
	Notice *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice `json:"Notice,omitempty" xml:"Notice,omitempty" type:"Struct"`
	// The ID of the region to which the alert rule belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the alert rule. `RUNNING`: The alert rule is enabled. `STOPPED`: The alert rule is disabled.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the Application Real-Time Monitoring Service (ARMS) task that is associated with the alert rule.
	//
	// example:
	//
	// 123
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The status of the task. This parameter is hidden from users.
	//
	// example:
	//
	// ""
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// AlertTest
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The timestamp that shows when the alert rule was updated.
	//
	// example:
	//
	// 1480521600000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user to which the alert rule belongs.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlarmContext() *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	return s.AlarmContext
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertRule() *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	return s.AlertRule
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertType() *int32 {
	return s.AlertType
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertVersion() *int32 {
	return s.AlertVersion
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetAlertWays() []*string {
	return s.AlertWays
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetConfig() *string {
	return s.Config
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetContactGroupIdList() *string {
	return s.ContactGroupIdList
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetContactGroupIds() *string {
	return s.ContactGroupIds
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetHostByAlertManager() *bool {
	return s.HostByAlertManager
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetId() *int64 {
	return s.Id
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetMetricParam() *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	return s.MetricParam
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetNotice() *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	return s.Notice
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetStatus() *string {
	return s.Status
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetTaskId() *int64 {
	return s.TaskId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetTitle() *string {
	return s.Title
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) GetUserId() *string {
	return s.UserId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlarmContext(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlarmContext = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertLevel(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertLevel = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertRule(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRule = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertVersion(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertVersion = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWays(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWays = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetConfig(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Config = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIdList(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIdList = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIds(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetCreateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetHostByAlertManager(v bool) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.HostByAlertManager = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Id = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetMetricParam(v *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricParam = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetNotice(v *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Notice = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetResourceGroupId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Status = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskStatus = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUpdateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext struct {
	// The sub-title of the alert notification content.
	//
	// example:
	//
	// TestSubTitle
	AlarmContentSubTitle *string `json:"AlarmContentSubTitle,omitempty" xml:"AlarmContentSubTitle,omitempty"`
	// The template of the alert notification.
	//
	// example:
	//
	// Alert name: $Alert name\\nFilter condition: $Filter\\nAlert time: $Alert time\\nAlert content: $Alert content\\nNote: The alert persists until a reply email is received. The system will remind you again in 24 hours.
	AlarmContentTemplate *string `json:"AlarmContentTemplate,omitempty" xml:"AlarmContentTemplate,omitempty"`
	// The content of the alert notification.
	//
	// example:
	//
	// Alert name: $Alert name\\nFilter condition: $Filter\\nAlert time: $Alert time\\nAlert content: $Alert content\\nNote: The alert persists until a reply email is received. The system will remind you again in 24 hours.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The sub-title of the alert notification.
	//
	// example:
	//
	// test
	SubTitle *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetAlarmContentSubTitle() *string {
	return s.AlarmContentSubTitle
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetAlarmContentTemplate() *string {
	return s.AlarmContentTemplate
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetContent() *string {
	return s.Content
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GetSubTitle() *string {
	return s.SubTitle
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentSubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentTemplate(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentTemplate = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetContent(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.Content = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.SubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule struct {
	// The logical operator between conditions. Valid values: `&`: AND. `|`: OR.
	//
	// example:
	//
	// |
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The condition of the alert rule.
	Rules []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GetOperator() *string {
	return s.Operator
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GetRules() []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	return s.Rules
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Rules = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules struct {
	// The aggregation logic of the metric data of the alert rule. Valid values:
	//
	// 	- `AVG`: calculates the average value for each minute
	//
	// 	- `SUM`: calculates the total value for each minute
	//
	// 	- `MAX`: calculates the maximum value for each minute
	//
	// 	- `MIN`: calculates the minimum value for each minute
	//
	// example:
	//
	// AVG
	Aggregates *string `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	// The displayed description of the alert metric.
	//
	// example:
	//
	// response time_ms
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The metric based on which alerts are triggered. For more information, see the "[Alert metrics](https://help.aliyun.com/document_detail/175825.html#h2-url-4)" section in this topic.
	//
	// example:
	//
	// appstat.jvm.SystemDiskFree
	Measure *string `json:"Measure,omitempty" xml:"Measure,omitempty"`
	// The time range when data is requested. Unit: minutes. For example, a value of 5 indicates that the alert rule applies to the data in the last 5 minutes.
	//
	// example:
	//
	// 5
	NValue *int32 `json:"NValue,omitempty" xml:"NValue,omitempty"`
	// The operation logic of the condition. Valid values:
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
	// The threshold of the condition.
	//
	// example:
	//
	// 30
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetAggregates() *string {
	return s.Aggregates
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetAlias() *string {
	return s.Alias
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetMeasure() *string {
	return s.Measure
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetNValue() *int32 {
	return s.NValue
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetOperator() *string {
	return s.Operator
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GetValue() *float32 {
	return s.Value
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAggregates(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Aggregates = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAlias(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Alias = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetMeasure(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Measure = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetNValue(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.NValue = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetValue(v float32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Value = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam struct {
	// The ID of the application group that is associated with the alert rule. This parameter is applicable to Enterprise Distributed Application Service (EDAS) applications.
	//
	// example:
	//
	// DEFAULT
	AppGroupId *string `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	// The auto-increment ID of the ARMS application. You can ignore this ID.
	//
	// example:
	//
	// 123
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The dimensions in the condition.
	Dimensions []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	// The PID of the application that is associated with the alert rule.
	//
	// example:
	//
	// 9870ca99-8105-4da7-a3a4-d72dd1b1****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The type of the metric. Valid values:
	//
	// 	- `txn`: the number of API calls during application monitoring
	//
	// 	- `txn_type`: the types of API calls during application monitoring
	//
	// 	- `db`: database metrics
	//
	// 	- `jvm`: Java virtual machine (JVM) metrics
	//
	// 	- `host`: host metrics
	//
	// 	- `exception`: API call errors
	//
	// example:
	//
	// DB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetAppGroupId() *string {
	return s.AppGroupId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetAppId() *string {
	return s.AppId
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetDimensions() []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	return s.Dimensions
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetPid() *string {
	return s.Pid
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GetType() *string {
	return s.Type
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppGroupId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppGroupId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetDimensions(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Dimensions = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetPid(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions struct {
	// The key of the dimension. Valid values:
	//
	// 	- `rpc`: the name of the API
	//
	// 	- `rpcType`: the type of the API call, such as HTTP or DUBBO
	//
	// 	- `endpoint`: the name of the database
	//
	// 	- `rootIp`: the IP address of the host
	//
	// example:
	//
	// rootIp
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the dimension. Valid values:
	//
	// 	- `STATIC`: checks only the value of this dimension. In this case, you must set the **dimensions.value*	- parameter.
	//
	// 	- `ALL`: checks the values of all dimensions. The metrics of all API calls are checked. If an API call triggers an alert, the name of the API is displayed in the alert notification. In this case, you do not need to set the **dimensions.value*	- parameter.
	//
	// 	- `DISABLE`: aggregates the values of all dimensions. In this case, you do not need to set the **dimensions.value*	- parameter.
	//
	// example:
	//
	// DISABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the dimension.
	//
	// example:
	//
	// "127.0.0.1"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GetKey() *string {
	return s.Key
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GetType() *string {
	return s.Type
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GetValue() *string {
	return s.Value
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetKey(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Key = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetValue(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Value = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) Validate() error {
	return dara.Validate(s)
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesNotice struct {
	// The end of the time range when the alert rule takes effect within 24 hours per day. This value is a UNIX timestamp. The year, month, and day that are indicated by the timestamp are not displayed in this value. Only the hour, minute, and second are displayed.
	//
	// example:
	//
	// 1480607940000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The end of the time range when alert notifications are sent based on the alert rule within 24 hours per day. This value is a UNIX timestamp. The year, month, and day that are indicated by the timestamp are not displayed in this value. Only the hour, minute, and second are displayed.
	//
	// example:
	//
	// 1480607940000
	NoticeEndTime *int64 `json:"NoticeEndTime,omitempty" xml:"NoticeEndTime,omitempty"`
	// The beginning of the time range when alert notifications are sent based on the alert rule within 24 hours per day. This value is a UNIX timestamp. The year, month, and day that are indicated by the timestamp are not displayed in this value. Only the hour, minute, and second are displayed.
	//
	// example:
	//
	// 1480521600000
	NoticeStartTime *int64 `json:"NoticeStartTime,omitempty" xml:"NoticeStartTime,omitempty"`
	// The beginning of the time range when the alert rule takes effect within 24 hours per day. This value is a UNIX timestamp. The year, month, and day that are indicated by the timestamp are not displayed in this value. Only the hour, minute, and second are displayed.
	//
	// example:
	//
	// 1480521600000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetNoticeEndTime() *int64 {
	return s.NoticeEndTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetNoticeStartTime() *int64 {
	return s.NoticeStartTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.EndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeEndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeStartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.StartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) Validate() error {
	return dara.Validate(s)
}
