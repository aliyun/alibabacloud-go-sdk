// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDISyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDISyncTaskResponseBodyData) *GetDISyncTaskResponseBody
	GetData() *GetDISyncTaskResponseBodyData
	SetRequestId(v string) *GetDISyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDISyncTaskResponseBody
	GetSuccess() *bool
}

type GetDISyncTaskResponseBody struct {
	// The returned results.
	Data *GetDISyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDISyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponseBody) GetData() *GetDISyncTaskResponseBodyData {
	return s.Data
}

func (s *GetDISyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDISyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDISyncTaskResponseBody) SetData(v *GetDISyncTaskResponseBodyData) *GetDISyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetDISyncTaskResponseBody) SetRequestId(v string) *GetDISyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDISyncTaskResponseBody) SetSuccess(v bool) *GetDISyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetDISyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDISyncTaskResponseBodyData struct {
	// The alert rules that are associated with the real-time synchronization task. The value of this parameter is an array.
	AlarmList []*GetDISyncTaskResponseBodyDataAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Repeated"`
	// 	- If the TaskType parameter is set to DI_REALTIME, the details of the real-time synchronization task are returned.
	//
	// 	- If the TaskType parameter is set to DI_SOLUTION, the value null is returned.
	//
	// example:
	//
	// {"extend":{"mode":"wizard","resourceGroup":"S_res_group_287114642182658_1560324290517"},"nodeDef":{},"order":{"hops":[{"from":"datahub_8htXSsfiS2vtZCVG","to":"datahub_CRHBAyGfhSaLmv2f"}]},"setting":{"errorLimit":{},"jvmOption":""},"steps":[{"stepType":"datahub","category":"writer","displayName":"DataHub1","parameter":{"batchSize":1000,"datasource":"datahub_cloud_dev_test","topic":"dwd_tfc_opt_speed_rid_amap_rt"},"name":"datahub_CRHBAyGfhSaLmv2f","gui":{"x":262,"y":325}},{"stepType":"datahub","displayName":"DataHub2","parameter":{"datasource":"datahub_uric_test","topic":"dwd_tfc_opt_speed_rid_amap_rt_330000","batchSize":1000},"name":"datahub_8htXSsfiS2vtZCVG","gui":{"x":268,"y":160.5},"category":"writer"}]}
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cause of the failure to obtain the details of the real-time synchronization task or data synchronization solution.
	//
	// If the details of the real-time synchronization task or data synchronization solution are obtained, the value null is returned.
	//
	// example:
	//
	// fileId:[100] is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 	- If the TaskType parameter is set to DI_REALTIME, the value null is returned.
	//
	// 	- If the TaskType parameter is set to DI_SOLUTION, the details of the data synchronization solution are returned.
	SolutionDetail *GetDISyncTaskResponseBodyDataSolutionDetail `json:"SolutionDetail,omitempty" xml:"SolutionDetail,omitempty" type:"Struct"`
	// Indicates whether the details of the real-time synchronization task or data synchronization solution are obtained. Valid values:
	//
	// success: The details are obtained. fail: The details fail to be obtained.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDISyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponseBodyData) GetAlarmList() []*GetDISyncTaskResponseBodyDataAlarmList {
	return s.AlarmList
}

func (s *GetDISyncTaskResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *GetDISyncTaskResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *GetDISyncTaskResponseBodyData) GetSolutionDetail() *GetDISyncTaskResponseBodyDataSolutionDetail {
	return s.SolutionDetail
}

func (s *GetDISyncTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDISyncTaskResponseBodyData) SetAlarmList(v []*GetDISyncTaskResponseBodyDataAlarmList) *GetDISyncTaskResponseBodyData {
	s.AlarmList = v
	return s
}

func (s *GetDISyncTaskResponseBodyData) SetCode(v string) *GetDISyncTaskResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDISyncTaskResponseBodyData) SetMessage(v string) *GetDISyncTaskResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetDISyncTaskResponseBodyData) SetSolutionDetail(v *GetDISyncTaskResponseBodyDataSolutionDetail) *GetDISyncTaskResponseBodyData {
	s.SolutionDetail = v
	return s
}

func (s *GetDISyncTaskResponseBodyData) SetStatus(v string) *GetDISyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDISyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDISyncTaskResponseBodyDataAlarmList struct {
	// The alert notification settings. The value of this parameter is an array.
	AlarmRuleList []*GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList `json:"AlarmRuleList,omitempty" xml:"AlarmRuleList,omitempty" type:"Repeated"`
	// The description of the alert rule.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the alert rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// 45242
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The alert type. Valid values:
	//
	// 	- taskStatus
	//
	// 	- bizDelay
	//
	// 	- taskFailoverCount
	//
	// 	- ddlUnsupport
	//
	// 	- ddlReport
	//
	// 	- totalDirtyRecordWriteInLines
	//
	// example:
	//
	// taskStatus
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The settings for alert notification rules. The value of this parameter is an array.
	NotifyRule *GetDISyncTaskResponseBodyDataAlarmListNotifyRule `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty" type:"Struct"`
	// The name of the alert rule.
	//
	// example:
	//
	// Delay alert rule name 1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetDISyncTaskResponseBodyDataAlarmList) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponseBodyDataAlarmList) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetAlarmRuleList() []*GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList {
	return s.AlarmRuleList
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetDescription() *string {
	return s.Description
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetId() *int64 {
	return s.Id
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetMetric() *string {
	return s.Metric
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetNotifyRule() *GetDISyncTaskResponseBodyDataAlarmListNotifyRule {
	return s.NotifyRule
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetAlarmRuleList(v []*GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) *GetDISyncTaskResponseBodyDataAlarmList {
	s.AlarmRuleList = v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetDescription(v string) *GetDISyncTaskResponseBodyDataAlarmList {
	s.Description = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetEnabled(v bool) *GetDISyncTaskResponseBodyDataAlarmList {
	s.Enabled = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetId(v int64) *GetDISyncTaskResponseBodyDataAlarmList {
	s.Id = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetMetric(v string) *GetDISyncTaskResponseBodyDataAlarmList {
	s.Metric = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetNotifyRule(v *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) *GetDISyncTaskResponseBodyDataAlarmList {
	s.NotifyRule = v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) SetRuleName(v string) *GetDISyncTaskResponseBodyDataAlarmList {
	s.RuleName = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmList) Validate() error {
	return dara.Validate(s)
}

type GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList struct {
	// The calculation method of a metric. Valid values:
	//
	// 	- avg
	//
	// 	- max
	//
	// example:
	//
	// avg
	Aggregator *string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The comparison operator, which indicates the method used to compare a metric with the alert rule.
	//
	// 	- \\"=\\"
	//
	// 	- \\"<\\"
	//
	// 	- \\">\\"
	//
	// example:
	//
	// =
	Comparator *string `json:"Comparator,omitempty" xml:"Comparator,omitempty"`
	// The duration that a condition is met before an alert is triggered. Unit: minutes.
	//
	// example:
	//
	// 3
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 	- WARNING
	//
	// 	- CRITICAL
	//
	// example:
	//
	// WARNING
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The threshold for the comparison between a metric and the alert rule.
	//
	// example:
	//
	// 1
	Threshold *int64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) GetAggregator() *string {
	return s.Aggregator
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) GetComparator() *string {
	return s.Comparator
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) GetDuration() *int64 {
	return s.Duration
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) GetLevel() *string {
	return s.Level
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) GetThreshold() *int64 {
	return s.Threshold
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) SetAggregator(v string) *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList {
	s.Aggregator = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) SetComparator(v string) *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList {
	s.Comparator = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) SetDuration(v int64) *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList {
	s.Duration = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) SetLevel(v string) *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList {
	s.Level = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) SetThreshold(v int64) *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList {
	s.Threshold = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListAlarmRuleList) Validate() error {
	return dara.Validate(s)
}

type GetDISyncTaskResponseBodyDataAlarmListNotifyRule struct {
	// The settings for Critical-level alert notifications.
	Critical []*string `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Repeated"`
	// The alert interval. Unit: minutes.
	//
	// example:
	//
	// 5
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The settings for Warning-level alert notifications.
	Warning []*string `json:"Warning,omitempty" xml:"Warning,omitempty" type:"Repeated"`
}

func (s GetDISyncTaskResponseBodyDataAlarmListNotifyRule) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponseBodyDataAlarmListNotifyRule) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) GetCritical() []*string {
	return s.Critical
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) GetInterval() *int64 {
	return s.Interval
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) GetWarning() []*string {
	return s.Warning
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) SetCritical(v []*string) *GetDISyncTaskResponseBodyDataAlarmListNotifyRule {
	s.Critical = v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) SetInterval(v int64) *GetDISyncTaskResponseBodyDataAlarmListNotifyRule {
	s.Interval = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) SetWarning(v []*string) *GetDISyncTaskResponseBodyDataAlarmListNotifyRule {
	s.Warning = v
	return s
}

func (s *GetDISyncTaskResponseBodyDataAlarmListNotifyRule) Validate() error {
	return dara.Validate(s)
}

type GetDISyncTaskResponseBodyDataSolutionDetail struct {
	// The creator of the data synchronization solution.
	//
	// example:
	//
	// dataworks_di
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The ID of the data synchronization solution.
	//
	// example:
	//
	// 100
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the data synchronization solution.
	//
	// example:
	//
	// holo_20211206161025
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configuration details of the data synchronization solution.
	//
	// example:
	//
	// {"holoDatasource":"holo","offlineSyncConfig":{"nodeNameRule":"oneclick_holo_di_${db_table_name_src}_to_${db_table_name_dest}","resourceGroup":"group_219193793999490"},"processRuleId":1007,"readerConcurrent":10,"realtimeSyncConfig":{"content":{"order":{"hops":[{"from":"reader","to":"writer"}]},"setting":{"speed":{"readerConcurrent":10}},"steps":[{"stepType":"mysql","name":"reader","category":"reader","parameter":{"connection":[{"datasource":"mm","datasourceType":"mysql","table":[]}]}},{"stepType":"holo","name":"writer","category":"writer","parameter":{"datasource":"holo","writeMode":"replay","datasourceSchema":"public","tableMappingRule":{"datasource":[{"tableRule":[],"srcDatasourceName":"mm"}]}}}]},"extend":{"mode":"migration_holo","resourceGroup":"group_219193793999490"}},"setting":{"autoCreateWorkflow":true,"userDefinedFileNameExpression":"oneclick"},"srcType":"mysql","tableMappingRuleFromRealtimeSyncConfig":{"datasource":[{"srcDatasourceName":"mm","tableRule":[]}]}}
	ProcessContent *string `json:"ProcessContent,omitempty" xml:"ProcessContent,omitempty"`
	// The additional parameters of the data synchronization solution.
	//
	// example:
	//
	// {"processType":"new","tableNum":300}
	ProcessExtra *string `json:"ProcessExtra,omitempty" xml:"ProcessExtra,omitempty"`
	// The ID of the project to which the data synchronization solution belongs.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the source of the data synchronization solution.
	//
	// example:
	//
	// mysql
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the data synchronization solution.
	//
	// example:
	//
	// 2021-12-07 14:40:51
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the data synchronization solution. Valid values:
	//
	// 	- 0: successful
	//
	// 	- 1: not running
	//
	// 	- 2: running
	//
	// 	- 3: failed
	//
	// 	- 4: committed
	//
	// 	- 5: pending manual confirmation
	//
	// 	- 6: manually confirmed
	//
	// 	- 7: others
	//
	// 	- 8: waiting
	//
	// 	- 9: deleted
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the data synchronization solution was committed.
	//
	// example:
	//
	// 2021-12-07 14:40:51
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// The type of the data synchronization solution.
	//
	// example:
	//
	// holo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDISyncTaskResponseBodyDataSolutionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDISyncTaskResponseBodyDataSolutionDetail) GoString() string {
	return s.String()
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetId() *int64 {
	return s.Id
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetName() *string {
	return s.Name
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetProcessContent() *string {
	return s.ProcessContent
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetProcessExtra() *string {
	return s.ProcessExtra
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetSourceType() *string {
	return s.SourceType
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetStatus() *string {
	return s.Status
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) GetType() *string {
	return s.Type
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetCreatorName(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.CreatorName = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetId(v int64) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.Id = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetName(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.Name = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetProcessContent(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.ProcessContent = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetProcessExtra(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.ProcessExtra = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetProjectId(v int64) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.ProjectId = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetSourceType(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.SourceType = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetStartTime(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.StartTime = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetStatus(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.Status = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetSubmitTime(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.SubmitTime = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) SetType(v string) *GetDISyncTaskResponseBodyDataSolutionDetail {
	s.Type = &v
	return s
}

func (s *GetDISyncTaskResponseBodyDataSolutionDetail) Validate() error {
	return dara.Validate(s)
}
