// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysisJobProgress(v []*DescribePreCheckStatusResponseBodyAnalysisJobProgress) *DescribePreCheckStatusResponseBody
	GetAnalysisJobProgress() []*DescribePreCheckStatusResponseBodyAnalysisJobProgress
	SetCode(v string) *DescribePreCheckStatusResponseBody
	GetCode() *string
	SetErrorAnalysisItem(v int32) *DescribePreCheckStatusResponseBody
	GetErrorAnalysisItem() *int32
	SetErrorItem(v int32) *DescribePreCheckStatusResponseBody
	GetErrorItem() *int32
	SetFullNetCheckJobStatus(v []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) *DescribePreCheckStatusResponseBody
	GetFullNetCheckJobStatus() []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatus
	SetHttpStatusCode(v int32) *DescribePreCheckStatusResponseBody
	GetHttpStatusCode() *int32
	SetJobId(v string) *DescribePreCheckStatusResponseBody
	GetJobId() *string
	SetJobName(v string) *DescribePreCheckStatusResponseBody
	GetJobName() *string
	SetJobProgress(v []*DescribePreCheckStatusResponseBodyJobProgress) *DescribePreCheckStatusResponseBody
	GetJobProgress() []*DescribePreCheckStatusResponseBodyJobProgress
	SetNetworkDiagnosisResult(v *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) *DescribePreCheckStatusResponseBody
	GetNetworkDiagnosisResult() *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult
	SetPageNumber(v int64) *DescribePreCheckStatusResponseBody
	GetPageNumber() *int64
	SetPageRecordCount(v int64) *DescribePreCheckStatusResponseBody
	GetPageRecordCount() *int64
	SetRequestId(v string) *DescribePreCheckStatusResponseBody
	GetRequestId() *string
	SetState(v string) *DescribePreCheckStatusResponseBody
	GetState() *string
	SetSubDistributedJobStatus(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatus) *DescribePreCheckStatusResponseBody
	GetSubDistributedJobStatus() []*DescribePreCheckStatusResponseBodySubDistributedJobStatus
	SetSuccess(v bool) *DescribePreCheckStatusResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribePreCheckStatusResponseBody
	GetTotal() *int32
	SetTotalRecordCount(v int64) *DescribePreCheckStatusResponseBody
	GetTotalRecordCount() *int64
}

type DescribePreCheckStatusResponseBody struct {
	// Display list of evaluation tasks
	AnalysisJobProgress []*DescribePreCheckStatusResponseBodyAnalysisJobProgress `json:"AnalysisJobProgress,omitempty" xml:"AnalysisJobProgress,omitempty" type:"Repeated"`
	// The task code that indicates the type of the subtask. Valid values:
	//
	// 	- **01**: precheck.
	//
	// 	- **02**: schema migration or initial schema synchronization.
	//
	// 	- **03**: full data migration or initial full data synchronization.
	//
	// 	- **04**: incremental data migration or synchronization.
	//
	// example:
	//
	// 01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Number of failed evaluation items
	//
	// example:
	//
	// 0
	ErrorAnalysisItem *int32 `json:"ErrorAnalysisItem,omitempty" xml:"ErrorAnalysisItem,omitempty"`
	// The total number of subtask failures.
	//
	// example:
	//
	// 0
	ErrorItem *int32 `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	// Network-wide inspection results.
	FullNetCheckJobStatus []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatus `json:"FullNetCheckJobStatus,omitempty" xml:"FullNetCheckJobStatus,omitempty" type:"Repeated"`
	// The status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// example:
	//
	// b4my3zg929a****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the subtask.
	//
	// example:
	//
	// dtstest
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The subtasks and the progress of each subtask.
	JobProgress []*DescribePreCheckStatusResponseBodyJobProgress `json:"JobProgress,omitempty" xml:"JobProgress,omitempty" type:"Repeated"`
	// Network diagnosis result
	NetworkDiagnosisResult *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult `json:"NetworkDiagnosisResult,omitempty" xml:"NetworkDiagnosisResult,omitempty" type:"Struct"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageRecordCount *int64 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C096FA97-B6BA-4575-899D-61E12B59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the subtask. Valid values:
	//
	// 	- **NotStarted**: The subtask is not started.
	//
	// 	- **Suspending**: The subtask is paused.
	//
	// 	- **Checking**: The subtask is being checked.
	//
	// 	- **Migrating**: The subtask is in progress. Data is being migrated.
	//
	// 	- **Failed**: The subtask failed.
	//
	// 	- **Catched**: The subtask is in progress. Incremental data is being migrated or synchronized.
	//
	// 	- **Finished**: The subtask is complete.
	//
	// example:
	//
	// Finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The information about the distributed subtasks.
	SubDistributedJobStatus []*DescribePreCheckStatusResponseBodySubDistributedJobStatus `json:"SubDistributedJobStatus,omitempty" xml:"SubDistributedJobStatus,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of subtasks.
	//
	// example:
	//
	// 0
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 100
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePreCheckStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBody) GetAnalysisJobProgress() []*DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	return s.AnalysisJobProgress
}

func (s *DescribePreCheckStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePreCheckStatusResponseBody) GetErrorAnalysisItem() *int32 {
	return s.ErrorAnalysisItem
}

func (s *DescribePreCheckStatusResponseBody) GetErrorItem() *int32 {
	return s.ErrorItem
}

func (s *DescribePreCheckStatusResponseBody) GetFullNetCheckJobStatus() []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	return s.FullNetCheckJobStatus
}

func (s *DescribePreCheckStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribePreCheckStatusResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBody) GetJobName() *string {
	return s.JobName
}

func (s *DescribePreCheckStatusResponseBody) GetJobProgress() []*DescribePreCheckStatusResponseBodyJobProgress {
	return s.JobProgress
}

func (s *DescribePreCheckStatusResponseBody) GetNetworkDiagnosisResult() *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult {
	return s.NetworkDiagnosisResult
}

func (s *DescribePreCheckStatusResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribePreCheckStatusResponseBody) GetPageRecordCount() *int64 {
	return s.PageRecordCount
}

func (s *DescribePreCheckStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePreCheckStatusResponseBody) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBody) GetSubDistributedJobStatus() []*DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	return s.SubDistributedJobStatus
}

func (s *DescribePreCheckStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePreCheckStatusResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribePreCheckStatusResponseBody) SetAnalysisJobProgress(v []*DescribePreCheckStatusResponseBodyAnalysisJobProgress) *DescribePreCheckStatusResponseBody {
	s.AnalysisJobProgress = v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetCode(v string) *DescribePreCheckStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetErrorAnalysisItem(v int32) *DescribePreCheckStatusResponseBody {
	s.ErrorAnalysisItem = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetErrorItem(v int32) *DescribePreCheckStatusResponseBody {
	s.ErrorItem = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetFullNetCheckJobStatus(v []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) *DescribePreCheckStatusResponseBody {
	s.FullNetCheckJobStatus = v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetHttpStatusCode(v int32) *DescribePreCheckStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetJobId(v string) *DescribePreCheckStatusResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetJobName(v string) *DescribePreCheckStatusResponseBody {
	s.JobName = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetJobProgress(v []*DescribePreCheckStatusResponseBodyJobProgress) *DescribePreCheckStatusResponseBody {
	s.JobProgress = v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetNetworkDiagnosisResult(v *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) *DescribePreCheckStatusResponseBody {
	s.NetworkDiagnosisResult = v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetPageNumber(v int64) *DescribePreCheckStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetPageRecordCount(v int64) *DescribePreCheckStatusResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetRequestId(v string) *DescribePreCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetState(v string) *DescribePreCheckStatusResponseBody {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetSubDistributedJobStatus(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatus) *DescribePreCheckStatusResponseBody {
	s.SubDistributedJobStatus = v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetSuccess(v bool) *DescribePreCheckStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetTotal(v int32) *DescribePreCheckStatusResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetTotalRecordCount(v int64) *DescribePreCheckStatusResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyAnalysisJobProgress struct {
	// The specific project start time, formatted as <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// example:
	//
	// 2022-03-16T08:01:31.000+00:00
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// Whether to support skipping this sub-item.
	//
	// example:
	//
	// true
	CanSkip *bool `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	// The number of currently running subtasks.
	//
	// example:
	//
	// 0
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The DDL operation to be executed.
	//
	// example:
	//
	// CREATE TABLE ****
	DdlSql *string `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	// Task delay time
	//
	// example:
	//
	// 0
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// Name of the database to which the migration objects in the target instance belong.
	//
	// example:
	//
	// dest
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// This parameter will be deprecated.
	//
	// example:
	//
	// 1
	DiffRow *int64 `json:"DiffRow,omitempty" xml:"DiffRow,omitempty"`
	// Error details when the project encounters an error.
	//
	// example:
	//
	// ANALYSIS_MYSQL
	ErrDetail *string `json:"ErrDetail,omitempty" xml:"ErrDetail,omitempty"`
	// Specific error message.
	//
	// example:
	//
	// ANALYSIS_
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The end time of the evaluation task, formatted as <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// example:
	//
	// 2022-03-16T08:01:31.000+00:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of this evaluation item in the database.
	//
	// example:
	//
	// 123123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Whether to directly ignore this specific item and move to the next one. Return values:
	//
	// - **N**: No. - **Y**: Yes.
	//
	// example:
	//
	// N
	IgnoreFlag *string `json:"IgnoreFlag,omitempty" xml:"IgnoreFlag,omitempty"`
	// Name of the evaluation item
	//
	// example:
	//
	// ANALYSIS_MYSQL_4_ITEM
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The ID of the evaluation task.
	//
	// example:
	//
	// 11234234xc
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Sub-assessment item.
	Logs []*DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Name of the evaluation item
	//
	// example:
	//
	// ANALYSIS_MYSQL_4_DETAIL
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The number of the evaluation item.
	//
	// example:
	//
	// 10
	OrderNum *int32 `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// This parameter will be deprecated.
	//
	// example:
	//
	// demo
	ParentObj *string `json:"ParentObj,omitempty" xml:"ParentObj,omitempty"`
	// Remediation method for the evaluation item.
	//
	// example:
	//
	// ANALYSIS_
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	// If this evaluation item fails, whether you set to skip this item. Return values: 	- **true**: Yes 	- **false**: No
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// Name of the database to which the migration objects in the source instance belong.
	//
	// example:
	//
	// dtstestdata
	SourceSchema *string `json:"SourceSchema,omitempty" xml:"SourceSchema,omitempty"`
	// The result of the evaluation, with return values being: - **Failed**: Failure. - **Success**: Success.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Progress of sub-projects under a specific project. > If it returns <b>[]</b>, it indicates there are no sub-projects.
	//
	// example:
	//
	// []
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// Name of the target object
	//
	// example:
	//
	// testTable
	TargetNames *string `json:"TargetNames,omitempty" xml:"TargetNames,omitempty"`
	// The total number of specific items in the sub-task.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyAnalysisJobProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyAnalysisJobProgress) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetBootTime() *string {
	return s.BootTime
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetCanSkip() *bool {
	return s.CanSkip
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetCurrent() *string {
	return s.Current
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetDdlSql() *string {
	return s.DdlSql
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetDelaySeconds() *int32 {
	return s.DelaySeconds
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetDestSchema() *string {
	return s.DestSchema
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetDiffRow() *int64 {
	return s.DiffRow
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetErrDetail() *string {
	return s.ErrDetail
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetId() *string {
	return s.Id
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetIgnoreFlag() *string {
	return s.IgnoreFlag
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetItem() *string {
	return s.Item
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetLogs() []*DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs {
	return s.Logs
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetNames() *string {
	return s.Names
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetOrderNum() *int32 {
	return s.OrderNum
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetParentObj() *string {
	return s.ParentObj
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetSkip() *bool {
	return s.Skip
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetSourceSchema() *string {
	return s.SourceSchema
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetSub() *string {
	return s.Sub
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetTargetNames() *string {
	return s.TargetNames
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetBootTime(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetCanSkip(v bool) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.CanSkip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetCurrent(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Current = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetDdlSql(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.DdlSql = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetDelaySeconds(v int32) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.DelaySeconds = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetDestSchema(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.DestSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetDiffRow(v int64) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.DiffRow = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetErrDetail(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.ErrDetail = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetFinishTime(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetId(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Id = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetIgnoreFlag(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.IgnoreFlag = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetItem(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Item = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetJobId(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetLogs(v []*DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Logs = v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetNames(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Names = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetOrderNum(v int32) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetParentObj(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.ParentObj = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetRepairMethod(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.RepairMethod = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetSkip(v bool) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Skip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetSourceSchema(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.SourceSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetState(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetSub(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Sub = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetTargetNames(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.TargetNames = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) SetTotal(v int32) *DescribePreCheckStatusResponseBodyAnalysisJobProgress {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgress) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs struct {
	// Error message
	//
	// example:
	//
	// Please modify this object
	ErrData *string `json:"ErrData,omitempty" xml:"ErrData,omitempty"`
	// Error message from DTS when a specific project encounters an error.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: Table \\"customer\\" already exists
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Error type.
	//
	// example:
	//
	// ForeignKey
	ErrType *string `json:"ErrType,omitempty" xml:"ErrType,omitempty"`
	// The level of the log.
	//
	// example:
	//
	// ERROR
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) GetErrData() *string {
	return s.ErrData
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) GetErrType() *string {
	return s.ErrType
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) SetErrData(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs {
	s.ErrData = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) SetErrType(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs {
	s.ErrType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) SetLogLevel(v string) *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs {
	s.LogLevel = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyAnalysisJobProgressLogs) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyFullNetCheckJobStatus struct {
	// Task code, **01*	- represents pre-check.
	//
	// example:
	//
	// 01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ID of the region to which the target network segment belongs.
	//
	// example:
	//
	// cn-hangzhou
	DestRegion *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// Destination network segment.
	//
	// example:
	//
	// 100.104.XX.XXX/XX
	DestRegionCidr *string `json:"DestRegionCidr,omitempty" xml:"DestRegionCidr,omitempty"`
	// The access method of the target instance, with return values as follows: - **ALIYUN**: Access method is **cloud instance**. - **OTHER**: Access method is **public IP**. - **ECS**: Access method is **ECS self-built database**. - **EXPRESS**: Access method is **Express Connect / VPN Gateway / Smart Gateway**. - **CEN**: Access method is **Cloud Enterprise Network (CEN)**. - **DG**: Access method is **Database Gateway (DG)**.
	//
	// example:
	//
	// CEN
	DestinationEndpointType *string `json:"DestinationEndpointType,omitempty" xml:"DestinationEndpointType,omitempty"`
	// Number of pre-check failed items
	//
	// example:
	//
	// 0
	ErrorItem *int32 `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	// The region ID of the instance\\"s running node.
	//
	// example:
	//
	// cn-hangzhou
	HostRegion *string `json:"HostRegion,omitempty" xml:"HostRegion,omitempty"`
	// Task ID.
	//
	// example:
	//
	// l3m1213ye7l****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Task name.
	//
	// example:
	//
	// dts.step.fullnetcheck
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// A list of specific items for the task and their execution progress.
	JobProgress []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress `json:"JobProgress,omitempty" xml:"JobProgress,omitempty" type:"Repeated"`
	// The access method of the source instance, with return values as follows: - **ALIYUN**: Access method is **cloud instance**. - **OTHER**: Access method is **public IP**. - **ECS**: Access method is **ECS self-built database**. - **EXPRESS**: Access method is **dedicated line/VPN gateway/smart gateway**. - **CEN**: Access method is **Cloud Enterprise Network CEN**. - **DG**: Access method is **Database Gateway DG**.
	//
	// example:
	//
	// CEN
	SourceEndpointType *string `json:"SourceEndpointType,omitempty" xml:"SourceEndpointType,omitempty"`
	// ID of the region to which the source network segment belongs.
	//
	// example:
	//
	// cn-hangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// Source network segment.
	//
	// example:
	//
	// 100.104.XX.XXX/XX
	SrcRegionCidr *string `json:"SrcRegionCidr,omitempty" xml:"SrcRegionCidr,omitempty"`
	// Check result, the return value is: - **Failed**: Failure. - **Success**: Completed.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Total number of items in the project.
	//
	// example:
	//
	// 11
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetCode() *string {
	return s.Code
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetDestRegionCidr() *string {
	return s.DestRegionCidr
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetDestinationEndpointType() *string {
	return s.DestinationEndpointType
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetErrorItem() *int32 {
	return s.ErrorItem
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetHostRegion() *string {
	return s.HostRegion
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetJobName() *string {
	return s.JobName
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetJobProgress() []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	return s.JobProgress
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetSourceEndpointType() *string {
	return s.SourceEndpointType
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetSrcRegionCidr() *string {
	return s.SrcRegionCidr
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetCode(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.Code = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetDestRegion(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.DestRegion = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetDestRegionCidr(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.DestRegionCidr = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetDestinationEndpointType(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.DestinationEndpointType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetErrorItem(v int32) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.ErrorItem = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetHostRegion(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.HostRegion = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetJobId(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetJobName(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.JobName = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetJobProgress(v []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.JobProgress = v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetSourceEndpointType(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.SourceEndpointType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetSrcRegion(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.SrcRegion = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetSrcRegionCidr(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.SrcRegionCidr = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetState(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) SetTotal(v int32) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatus) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress struct {
	// The specific project start time, formatted as <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC time).
	//
	// example:
	//
	// 2022-03-30T03:36:11.000+00:00
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// Whether DTS supports skipping a project after it fails. Return values: 	- **true**: Yes 	- **false**: No
	//
	// example:
	//
	// false
	CanSkip *bool `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	// The number of currently running tasks.
	//
	// example:
	//
	// 0
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The DDL operation to be executed.
	//
	// example:
	//
	// CREATE TABLE ****
	DdlSql *string `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	// Task delay time
	//
	// example:
	//
	// 0
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// Name of the database to which the migration objects in the target instance belong.
	//
	// example:
	//
	// dest
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// This parameter will be deprecated.
	//
	// example:
	//
	// 1
	DiffRow *int64 `json:"DiffRow,omitempty" xml:"DiffRow,omitempty"`
	// Details of the error when a specific project fails.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ_DETAIL
	ErrDetail *string `json:"ErrDetail,omitempty" xml:"ErrDetail,omitempty"`
	// Error message prompt when a specific project encounters an error.
	//
	// example:
	//
	// ODPS project does not exist odps.`huijin
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Task completion time, formatted as yyyy-MM-ddTHH:mm:ssZ (UTC time).
	//
	// example:
	//
	// 2022-03-31T03:36:11.000+00:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the record in the metadata database.
	//
	// example:
	//
	// 922305811766881****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Whether to directly ignore this specific item and move to the next one. Return values:
	//
	// - **N**: No. - **Y**: Yes.
	//
	// example:
	//
	// N
	IgnoreFlag *string `json:"IgnoreFlag,omitempty" xml:"IgnoreFlag,omitempty"`
	// Specific project name.
	//
	// example:
	//
	// CHECK_CONN_SRC
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// Task ID.
	//
	// example:
	//
	// l3m1213ye7l****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Error execution log information.
	Logs []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Specific project name.
	//
	// example:
	//
	// CHECK_CONN_SRC_DETAIL
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// Project number.
	//
	// example:
	//
	// 1
	OrderNum *int32 `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// This parameter will be deprecated.
	//
	// example:
	//
	// demo
	ParentObj *string `json:"ParentObj,omitempty" xml:"ParentObj,omitempty"`
	// The corresponding remediation method when the pre-check fails.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ_REPAIR
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	// After this specific item fails, do you set to skip this item. Return values: 	- **true**: Yes 	- **false**: No
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// Name of the database to which the migration objects in the source instance belong.
	//
	// example:
	//
	// dtstestdata
	SourceSchema *string `json:"SourceSchema,omitempty" xml:"SourceSchema,omitempty"`
	// Check result, the return value is: - **Failed**: Failure. - **Success**: Completed.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Progress of sub-projects under a specific project. > If it returns <b>[]</b>, it indicates there are no sub-projects.
	//
	// example:
	//
	// []
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// Name of the target object
	//
	// example:
	//
	// order
	TargetNames *string `json:"TargetNames,omitempty" xml:"TargetNames,omitempty"`
	// The total number of projects.
	//
	// example:
	//
	// 11
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetBootTime() *string {
	return s.BootTime
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetCanSkip() *bool {
	return s.CanSkip
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetCurrent() *string {
	return s.Current
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetDdlSql() *string {
	return s.DdlSql
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetDelaySeconds() *int32 {
	return s.DelaySeconds
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetDestSchema() *string {
	return s.DestSchema
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetDiffRow() *int64 {
	return s.DiffRow
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetErrDetail() *string {
	return s.ErrDetail
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetId() *string {
	return s.Id
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetIgnoreFlag() *string {
	return s.IgnoreFlag
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetItem() *string {
	return s.Item
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetLogs() []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs {
	return s.Logs
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetNames() *string {
	return s.Names
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetOrderNum() *int32 {
	return s.OrderNum
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetParentObj() *string {
	return s.ParentObj
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetSkip() *bool {
	return s.Skip
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetSourceSchema() *string {
	return s.SourceSchema
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetSub() *string {
	return s.Sub
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetTargetNames() *string {
	return s.TargetNames
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetBootTime(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetCanSkip(v bool) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.CanSkip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetCurrent(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Current = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetDdlSql(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.DdlSql = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetDelaySeconds(v int32) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.DelaySeconds = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetDestSchema(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.DestSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetDiffRow(v int64) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.DiffRow = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetErrDetail(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.ErrDetail = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetFinishTime(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetId(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Id = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetIgnoreFlag(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.IgnoreFlag = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetItem(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Item = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetJobId(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetLogs(v []*DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Logs = v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetNames(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Names = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetOrderNum(v int32) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetParentObj(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.ParentObj = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetRepairMethod(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.RepairMethod = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetSkip(v bool) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Skip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetSourceSchema(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.SourceSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetState(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetSub(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Sub = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetTargetNames(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.TargetNames = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) SetTotal(v int32) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgress) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs struct {
	// Error record.
	//
	// example:
	//
	// CREATE TABLE `dtstestdata`.`customer` ****
	ErrData *string `json:"ErrData,omitempty" xml:"ErrData,omitempty"`
	// Specific error message.
	//
	// example:
	//
	// get metric list fail
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// Type of error.
	//
	// example:
	//
	// ForeignKey
	ErrType *string `json:"ErrType,omitempty" xml:"ErrType,omitempty"`
	// The level of the log.
	//
	// example:
	//
	// INFO
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) GetErrData() *string {
	return s.ErrData
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) GetErrType() *string {
	return s.ErrType
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) SetErrData(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs {
	s.ErrData = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) SetErrType(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs {
	s.ErrType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) SetLogLevel(v string) *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs {
	s.LogLevel = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyFullNetCheckJobStatusJobProgressLogs) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyJobProgress struct {
	// The time when the subtask was started. The time is displayed in the yyyy-MM-ddTHH:mm:ssZ format in UTC.
	//
	// example:
	//
	// 2021-03-16T08:01:31.000+00:00
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// Indicates whether the subtask can be ignored if it fails.
	//
	// example:
	//
	// true
	CanSkip *bool `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	// The number of the subtasks that are running.
	//
	// example:
	//
	// 0
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The DDL statements.
	//
	// example:
	//
	// CREATE TABLE `dtstestdata`.`order` (\\n`orderid`  int(11)     COMMENT \\"\\"   NOT NULL   , \\n`username`  char(32)  CHARSET `utf8` COLLATE `utf8_general_ci`    COMMENT \\"\\"   NULL   , \\n`ordertime`  datetime     COMMENT \\"\\"   NULL   , \\n`commodity`  varchar(32)  CHARSET `utf8` COLLATE `utf8_general_ci`    COMMENT \\"\\"   NULL   , \\n`phonenumber`  int(11)     COMMENT \\"\\"   NULL   , \\n`address`  text  CHARSET `utf8mb4` COLLATE `utf8mb4_general_ci`    COMMENT \\"\\"   NULL   \\n, PRIMARY KEY (`orderid`)) engine=InnoDB DEFAULT CHARSET=`gbk` DEFAULT COLLATE `gbk_chinese_ci` ROW_FORMAT= Dynamic comment = \\"\\" ;\\n
	DdlSql *string `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	// The latency of incremental data migration or synchronization.
	//
	// > If you query data migration tasks, the unit of this parameter is milliseconds. If you query data synchronization tasks, the unit of this parameter is seconds.
	//
	// example:
	//
	// 0
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The name of the database to which the object in the destination instance belongs.
	//
	// example:
	//
	// dtstestdata_new
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// This parameter will be removed in the future.
	//
	// example:
	//
	// 1
	DiffRow *int64 `json:"DiffRow,omitempty" xml:"DiffRow,omitempty"`
	// The error details of the subtask failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ_DETAIL
	ErrDetail *string `json:"ErrDetail,omitempty" xml:"ErrDetail,omitempty"`
	// The error message of the subtask failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The time when the subtask was complete. The time is displayed in the yyyy-MM-ddTHH:mm:ssZ format in UTC.
	//
	// example:
	//
	// 2021-03-16T08:01:34.000+00:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the entry in the metadatabase.
	//
	// example:
	//
	// 5632
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether DTS ignores the subtask and proceeds with the next subtask. Valid values:
	//
	// 	- **N**: no.
	//
	// 	- **Y**: yes.
	//
	// example:
	//
	// N
	IgnoreFlag *string `json:"IgnoreFlag,omitempty" xml:"IgnoreFlag,omitempty"`
	// The shortened name of the subtask.
	//
	// example:
	//
	// CHECK_CONN_DEST
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The subtask ID.
	//
	// example:
	//
	// fj1c33ro168****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The logs of subtask failures.
	Logs []*DescribePreCheckStatusResponseBodyJobProgressLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The name of the subtask.
	//
	// example:
	//
	// CHECK_CONN_DEST_DETAIL
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The serial number of the subtask.
	//
	// example:
	//
	// 10
	OrderNum *int32 `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// This parameter will be removed in the future.
	//
	// example:
	//
	// demo
	ParentObj *string `json:"ParentObj,omitempty" xml:"ParentObj,omitempty"`
	// The method to fix the subtask failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ_REPAIR
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	// Indicates whether the subtask is ignored if it fails. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The name of the database to which the object in the source instance belongs.
	//
	// example:
	//
	// dtstestdata
	SourceSchema *string `json:"SourceSchema,omitempty" xml:"SourceSchema,omitempty"`
	// The status of the subtask. Valid values:
	//
	// 	- **NotStarted**: The subtask is not started.
	//
	// 	- **Checking**: The subtask is being checked.
	//
	// 	- **Migrating**: The subtask is in progress. Data is being migrated.
	//
	// 	- **Failed**: The subtask failed.
	//
	// 	- **Warning**: The subtask encounters an exception.
	//
	// 	- **Success**: The subtask is complete.
	//
	// example:
	//
	// Success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The sub-item progress of the subtask.
	//
	// > If \\*\\*[]\\*\\	- is returned, the subtask has no sub-items.
	//
	// example:
	//
	// []
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// The names of the objects that are migrated or synchronized.
	//
	// example:
	//
	// order
	TargetNames *string `json:"TargetNames,omitempty" xml:"TargetNames,omitempty"`
	// The total number of sub-items of the subtask.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyJobProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyJobProgress) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetBootTime() *string {
	return s.BootTime
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetCanSkip() *bool {
	return s.CanSkip
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetCurrent() *string {
	return s.Current
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetDdlSql() *string {
	return s.DdlSql
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetDelaySeconds() *int32 {
	return s.DelaySeconds
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetDestSchema() *string {
	return s.DestSchema
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetDiffRow() *int64 {
	return s.DiffRow
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetErrDetail() *string {
	return s.ErrDetail
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetId() *string {
	return s.Id
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetIgnoreFlag() *string {
	return s.IgnoreFlag
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetItem() *string {
	return s.Item
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetLogs() []*DescribePreCheckStatusResponseBodyJobProgressLogs {
	return s.Logs
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetNames() *string {
	return s.Names
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetOrderNum() *int32 {
	return s.OrderNum
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetParentObj() *string {
	return s.ParentObj
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetSkip() *bool {
	return s.Skip
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetSourceSchema() *string {
	return s.SourceSchema
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetSub() *string {
	return s.Sub
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetTargetNames() *string {
	return s.TargetNames
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetBootTime(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetCanSkip(v bool) *DescribePreCheckStatusResponseBodyJobProgress {
	s.CanSkip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetCurrent(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Current = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDdlSql(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DdlSql = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDelaySeconds(v int32) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DelaySeconds = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDestSchema(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DestSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDiffRow(v int64) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DiffRow = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetErrDetail(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.ErrDetail = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetFinishTime(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetId(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Id = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetIgnoreFlag(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.IgnoreFlag = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetItem(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Item = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetJobId(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetLogs(v []*DescribePreCheckStatusResponseBodyJobProgressLogs) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Logs = v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetNames(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Names = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetOrderNum(v int32) *DescribePreCheckStatusResponseBodyJobProgress {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetParentObj(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.ParentObj = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetRepairMethod(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.RepairMethod = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetSkip(v bool) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Skip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetSourceSchema(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.SourceSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetState(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetSub(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Sub = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetTargetNames(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.TargetNames = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetTotal(v int32) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyJobProgressLogs struct {
	// The error message.
	//
	// example:
	//
	// CREATE TABLE `dtstestdata`.`customer` (\\n`runoob_id`  int(10) unsigned   auto_increment  COMMENT \\"\\"   NOT NULL   , \\n`runoob_title`  varchar(100)  CHARSET `utf8` COLLATE `utf8_general_ci`    COMMENT \\"\\"   NOT NULL   , \\n`runoob_author1216`  varchar(40)  CHARSET `utf8` COLLATE `utf8_general_ci`    COMMENT \\"\\"   NOT NULL   , \\n`submission_date1216`  date     COMMENT \\"\\"   NULL   \\n, PRIMARY KEY (`runoob_id`)) engine=InnoDB AUTO_INCREMENT=200001 DEFAULT CHARSET=`utf8` DEFAULT COLLATE `utf8_general_ci` ROW_FORMAT= Dynamic comment = \\"\\" ;\\n
	ErrData *string `json:"ErrData,omitempty" xml:"ErrData,omitempty"`
	// The error message that is returned when an error occurs on the subtask.
	//
	// example:
	//
	// DTS-1020042 Execute sql error sql: Table \\"customer\\" already exists
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The error type.
	//
	// example:
	//
	// ForeignKey
	ErrType *string `json:"ErrType,omitempty" xml:"ErrType,omitempty"`
	// The level of logs.
	//
	// example:
	//
	// ERROR
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyJobProgressLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyJobProgressLogs) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) GetErrData() *string {
	return s.ErrData
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) GetErrType() *string {
	return s.ErrType
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetErrData(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.ErrData = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetErrType(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.ErrType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetLogLevel(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.LogLevel = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyNetworkDiagnosisResult struct {
	// Network diagnostic report
	Diagnosis []*DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty" type:"Repeated"`
	// Diagnose model version.
	//
	// example:
	//
	// network-v0.2
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) GetDiagnosis() []*DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis {
	return s.Diagnosis
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) SetDiagnosis(v []*DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult {
	s.Diagnosis = v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) SetModelVersion(v string) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult {
	s.ModelVersion = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResult) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis struct {
	// Document address for China region.
	//
	// example:
	//
	// https://***.ali***.com/document_detail/470447.html
	CnDocUrl *string `json:"CnDocUrl,omitempty" xml:"CnDocUrl,omitempty"`
	// Diagnostic code.
	//
	// example:
	//
	// dts.kunlun.diagnosis.network.express_doc
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Access point, the return values are: - **source**: source end. - **destination**: destination end. - **unknown**: unknown.
	//
	// example:
	//
	// source
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// Overseas region document address.
	//
	// example:
	//
	// https://www.ali***.com/help/en/data-transmission-service/latest/how-to-solve-an-error-when-accessing-a-database-instance-to-dts-using-vpn
	InternationalDocUrl *string `json:"InternationalDocUrl,omitempty" xml:"InternationalDocUrl,omitempty"`
	// Reserved field for diagnostic results, default is empty.
	//
	// example:
	//
	// none
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) GetCnDocUrl() *string {
	return s.CnDocUrl
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) GetCode() *string {
	return s.Code
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) GetInternationalDocUrl() *string {
	return s.InternationalDocUrl
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) GetResult() *string {
	return s.Result
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) SetCnDocUrl(v string) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis {
	s.CnDocUrl = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) SetCode(v string) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis {
	s.Code = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) SetEndpointType(v string) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis {
	s.EndpointType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) SetInternationalDocUrl(v string) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis {
	s.InternationalDocUrl = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) SetResult(v string) *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis {
	s.Result = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyNetworkDiagnosisResultDiagnosis) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodySubDistributedJobStatus struct {
	// The task code that indicates the type of the subtask. Valid values:
	//
	// 	- **01**: precheck.
	//
	// 	- **02**: schema migration or initial schema synchronization.
	//
	// 	- **03**: full data migration or initial full data synchronization.
	//
	// 	- **04**: incremental data migration or synchronization.
	//
	// example:
	//
	// 02
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of subtasks that failed.
	//
	// example:
	//
	// 0
	ErrorItem *int32 `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	// The subtask ID.
	//
	// example:
	//
	// n0gm1682j6563np
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of distributed subtasks associated with the subtask.
	//
	// example:
	//
	// dts.step.struct.load
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The subtasks and the progress of each subtask.
	JobProgress []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress `json:"JobProgress,omitempty" xml:"JobProgress,omitempty" type:"Repeated"`
	// The status of the subtask. Valid values:
	//
	// 	- **NotStarted**: The subtask is not started.
	//
	// 	- **Suspending**: The subtask is paused.
	//
	// 	- **Checking**: The subtask is being checked.
	//
	// 	- **Migrating**: The subtask is in progress. Data is being migrated.
	//
	// 	- **Failed**: The subtask failed.
	//
	// 	- **Catched**: The subtask is in progress. Incremental data is being migrated or synchronized.
	//
	// 	- **Finished**: The subtask is complete.
	//
	// example:
	//
	// Finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 11
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatus) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetCode() *string {
	return s.Code
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetErrorItem() *int32 {
	return s.ErrorItem
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetJobName() *string {
	return s.JobName
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetJobProgress() []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	return s.JobProgress
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetCode(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.Code = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetErrorItem(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.ErrorItem = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetJobId(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetJobName(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.JobName = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetJobProgress(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.JobProgress = v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetState(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetTotal(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress struct {
	// The time when the subtask was started. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2022-03-30T03:36:11.000+00:00
	BootTime *string `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	// Indicates whether the subtask can be ignored if it fails. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CanSkip *bool `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	// The number of the subtasks that are running.
	//
	// example:
	//
	// 0
	Current *string `json:"Current,omitempty" xml:"Current,omitempty"`
	// The DDL statements.
	//
	// example:
	//
	// None
	DdlSql *string `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	// The latency of incremental data migration or synchronization.
	//
	// example:
	//
	// 0
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The name of the database to which the object in the destination instance belongs.
	//
	// example:
	//
	// databasetest
	DestSchema *string `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	// This parameter will be removed in the future.
	//
	// example:
	//
	// None
	DiffRow *int64 `json:"DiffRow,omitempty" xml:"DiffRow,omitempty"`
	// The error details of the subtask failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ_DETAIL
	ErrDetail *string `json:"ErrDetail,omitempty" xml:"ErrDetail,omitempty"`
	// The error message of the subtask failure.
	//
	// example:
	//
	// ODPS project does not exist odps.`huijin
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The time when the subtask was complete. The time is displayed in the *yyyy-MM-dd*T*HH:mm:ss*Z format in UTC.
	//
	// example:
	//
	// 2022-03-31T03:36:11.000+00:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the entry in the metadatabase.
	//
	// example:
	//
	// 3890
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether DTS ignores the subtask and proceeds with the next subtask. Valid values:
	//
	// 	- **N**: no.
	//
	// 	- **Y**: yes.
	//
	// example:
	//
	// N
	IgnoreFlag *string `json:"IgnoreFlag,omitempty" xml:"IgnoreFlag,omitempty"`
	// The name of the subtask.
	//
	// example:
	//
	// login_common_time
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The subtask ID.
	//
	// example:
	//
	// l3m1213ye7l****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The operations logs of errors.
	Logs []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The name of the subtask.
	//
	// example:
	//
	// metricRuleTargets-20180308houe
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The serial number of the subtask.
	//
	// example:
	//
	// 1
	OrderNum *int32 `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	// This parameter will be removed in the future.
	//
	// example:
	//
	// None
	ParentObj *string `json:"ParentObj,omitempty" xml:"ParentObj,omitempty"`
	// The method to fix a precheck failure.
	//
	// example:
	//
	// CHECK__ERROR_SAME_OBJ_REPAIR
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	// Indicates whether the subtask was ignored. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The name of the database to which the object in the source instance belongs.
	//
	// example:
	//
	// databasetest
	SourceSchema *string `json:"SourceSchema,omitempty" xml:"SourceSchema,omitempty"`
	// The status of the subtask. Valid values:
	//
	// 	- **NotStarted**: The subtask is not started.
	//
	// 	- **Suspending**: The subtask is paused.
	//
	// 	- **Checking**: The subtask is being checked.
	//
	// 	- **Migrating**: The subtask is in progress. Data is being migrated.
	//
	// 	- **Failed**: The subtask failed.
	//
	// 	- **Catched**: The subtask is in progress. Incremental data is being migrated or synchronized.
	//
	// 	- **Finished**: The subtask is complete.
	//
	// example:
	//
	// Finished
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The sub-item progress of the subtask.
	//
	// > If \\*\\*[]\\*\\	- is returned, the subtask has no sub-item.
	//
	// example:
	//
	// []
	Sub *string `json:"Sub,omitempty" xml:"Sub,omitempty"`
	// The names of the objects that are migrated or synchronized.
	//
	// example:
	//
	// order
	TargetNames *string `json:"TargetNames,omitempty" xml:"TargetNames,omitempty"`
	// The total number of subtasks.
	//
	// example:
	//
	// 11
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetBootTime() *string {
	return s.BootTime
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetCanSkip() *bool {
	return s.CanSkip
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetCurrent() *string {
	return s.Current
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetDdlSql() *string {
	return s.DdlSql
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetDelaySeconds() *int32 {
	return s.DelaySeconds
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetDestSchema() *string {
	return s.DestSchema
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetDiffRow() *int64 {
	return s.DiffRow
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetErrDetail() *string {
	return s.ErrDetail
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetId() *string {
	return s.Id
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetIgnoreFlag() *string {
	return s.IgnoreFlag
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetItem() *string {
	return s.Item
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetJobId() *string {
	return s.JobId
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetLogs() []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	return s.Logs
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetNames() *string {
	return s.Names
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetOrderNum() *int32 {
	return s.OrderNum
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetParentObj() *string {
	return s.ParentObj
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetRepairMethod() *string {
	return s.RepairMethod
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetSkip() *bool {
	return s.Skip
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetSourceSchema() *string {
	return s.SourceSchema
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetState() *string {
	return s.State
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetSub() *string {
	return s.Sub
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetTargetNames() *string {
	return s.TargetNames
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GetTotal() *int32 {
	return s.Total
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetBootTime(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetCanSkip(v bool) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.CanSkip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetCurrent(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Current = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDdlSql(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DdlSql = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDelaySeconds(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DelaySeconds = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDestSchema(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DestSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDiffRow(v int64) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DiffRow = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetErrDetail(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.ErrDetail = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetErrMsg(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetFinishTime(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetId(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Id = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetIgnoreFlag(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.IgnoreFlag = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetItem(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Item = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetJobId(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetLogs(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Logs = v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetNames(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Names = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetOrderNum(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetParentObj(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.ParentObj = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetRepairMethod(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.RepairMethod = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetSkip(v bool) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Skip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetSourceSchema(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.SourceSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetState(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetSub(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Sub = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetTargetNames(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.TargetNames = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetTotal(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) Validate() error {
	return dara.Validate(s)
}

type DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs struct {
	// The record of errors.
	//
	// example:
	//
	// CREATE TABLE `dtstestdata`.`customer` (\\n`runoob_id` int(10) unsigned auto_increment COMMENT \\"\\" NOT NULL , \\n`runoob_title` varchar(100) CHARSET `utf8` COLLATE `utf8_general_ci` COMMENT \\"\\" NOT NULL , \\n`runoob_author1216` varchar(40) CHARSET `utf8` COLLATE `utf8_general_ci` COMMENT \\"\\" NOT NULL , \\n`submission_date1216` date COMMENT \\"\\" NULL \\n, PRIMARY KEY (`runoob_id`)) engine=InnoDB AUTO_INCREMENT=200001 DEFAULT CHARSET=`utf8` DEFAULT COLLATE `utf8_general_ci` ROW_FORMAT= Dynamic comment = \\"\\" ;\\n
	ErrData *string `json:"ErrData,omitempty" xml:"ErrData,omitempty"`
	// The error message.
	//
	// example:
	//
	// get metric list fail
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// The error type.
	//
	// example:
	//
	// ForeignKey
	ErrType *string `json:"ErrType,omitempty" xml:"ErrType,omitempty"`
	// The level of logs.
	//
	// example:
	//
	// INFO
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) GetErrData() *string {
	return s.ErrData
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) GetErrType() *string {
	return s.ErrType
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) GetLogLevel() *string {
	return s.LogLevel
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetErrData(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.ErrData = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetErrMsg(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetErrType(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.ErrType = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetLogLevel(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.LogLevel = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) Validate() error {
	return dara.Validate(s)
}
