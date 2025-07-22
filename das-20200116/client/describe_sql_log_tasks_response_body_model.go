// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSqlLogTasksResponseBody
	GetCode() *string
	SetData(v *DescribeSqlLogTasksResponseBodyData) *DescribeSqlLogTasksResponseBody
	GetData() *DescribeSqlLogTasksResponseBodyData
	SetMessage(v string) *DescribeSqlLogTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSqlLogTasksResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSqlLogTasksResponseBody
	GetSuccess() *string
}

type DescribeSqlLogTasksResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeSqlLogTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSqlLogTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSqlLogTasksResponseBody) GetData() *DescribeSqlLogTasksResponseBodyData {
	return s.Data
}

func (s *DescribeSqlLogTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSqlLogTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSqlLogTasksResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSqlLogTasksResponseBody) SetCode(v string) *DescribeSqlLogTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBody) SetData(v *DescribeSqlLogTasksResponseBodyData) *DescribeSqlLogTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSqlLogTasksResponseBody) SetMessage(v string) *DescribeSqlLogTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBody) SetRequestId(v string) *DescribeSqlLogTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBody) SetSuccess(v string) *DescribeSqlLogTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTasksResponseBodyData struct {
	// The details of the data returned.
	List []*DescribeSqlLogTasksResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of tasks.
	//
	// example:
	//
	// 40
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeSqlLogTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksResponseBodyData) GetList() []*DescribeSqlLogTasksResponseBodyDataList {
	return s.List
}

func (s *DescribeSqlLogTasksResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeSqlLogTasksResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSqlLogTasksResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSqlLogTasksResponseBodyData) SetList(v []*DescribeSqlLogTasksResponseBodyDataList) *DescribeSqlLogTasksResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyData) SetPageNo(v int64) *DescribeSqlLogTasksResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyData) SetPageSize(v int64) *DescribeSqlLogTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyData) SetTotal(v int64) *DescribeSqlLogTasksResponseBodyData {
	s.Total = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTasksResponseBodyDataList struct {
	// The time when the analysis task was complete. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1712751923000
	AnalysisTaskFinishTime *int64 `json:"AnalysisTaskFinishTime,omitempty" xml:"AnalysisTaskFinishTime,omitempty"`
	// The state of the analysis task.
	//
	// >  This parameter is a system parameter. You do not need to pay attention to the parameter.
	//
	// example:
	//
	// SCAN_ANALYZE_COMPLETED
	AnalysisTaskStatus *string `json:"AnalysisTaskStatus,omitempty" xml:"AnalysisTaskStatus,omitempty"`
	// The time when the task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1681363254423
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1705975320000
	End *int64 `json:"End,omitempty" xml:"End,omitempty"`
	// Indicates whether the task expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Expire *bool `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The filter parameters.
	Filters     []*DescribeSqlLogTasksResponseBodyDataListFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	InnerResult *string                                           `json:"InnerResult,omitempty" xml:"InnerResult,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// rm-2zew761kf7ho18752
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of log records.
	//
	// example:
	//
	// 99999
	LogCount *int64 `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	// The task name.
	//
	// example:
	//
	// test01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The task progress.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The URL that is returned if the value of TaskType is **Export**.
	//
	// example:
	//
	// https://das-sqllog-download-cn-shanghai.oss-cn-shanghai.aliyuncs.com/la
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The number of files that are scanned.
	//
	// example:
	//
	// 3000
	ScanFileSize *int64 `json:"ScanFileSize,omitempty" xml:"ScanFileSize,omitempty"`
	// The time when the task started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1683859555482
	Start *int64 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The task state. Valid values:
	//
	// 	- **INIT**: The task is to be scheduled.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **FAILED**: The task failed.
	//
	// 	- **CANCELED**: The task is canceled.
	//
	// 	- **COMPLETED**: The task is complete.
	//
	// >  If a task is in the **COMPLETED*	- state, you can view the results of the task.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 9a4f5c4494dbd6713185d87a97aa53e8
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. Valid values:
	//
	// 	- **Export**
	//
	// 	- **Query**
	//
	// example:
	//
	// Export
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeSqlLogTasksResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetAnalysisTaskFinishTime() *int64 {
	return s.AnalysisTaskFinishTime
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetAnalysisTaskStatus() *string {
	return s.AnalysisTaskStatus
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetEnd() *int64 {
	return s.End
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetExpire() *bool {
	return s.Expire
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetFilters() []*DescribeSqlLogTasksResponseBodyDataListFilters {
	return s.Filters
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetInnerResult() *string {
	return s.InnerResult
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetLogCount() *int64 {
	return s.LogCount
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetResult() *string {
	return s.Result
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetScanFileSize() *int64 {
	return s.ScanFileSize
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetStart() *int64 {
	return s.Start
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSqlLogTasksResponseBodyDataList) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetAnalysisTaskFinishTime(v int64) *DescribeSqlLogTasksResponseBodyDataList {
	s.AnalysisTaskFinishTime = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetAnalysisTaskStatus(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.AnalysisTaskStatus = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetCreateTime(v int64) *DescribeSqlLogTasksResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetEnd(v int64) *DescribeSqlLogTasksResponseBodyDataList {
	s.End = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetExpire(v bool) *DescribeSqlLogTasksResponseBodyDataList {
	s.Expire = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetFilters(v []*DescribeSqlLogTasksResponseBodyDataListFilters) *DescribeSqlLogTasksResponseBodyDataList {
	s.Filters = v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetInnerResult(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.InnerResult = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetInstanceId(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetLogCount(v int64) *DescribeSqlLogTasksResponseBodyDataList {
	s.LogCount = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetName(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetProgress(v int32) *DescribeSqlLogTasksResponseBodyDataList {
	s.Progress = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetResult(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.Result = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetScanFileSize(v int64) *DescribeSqlLogTasksResponseBodyDataList {
	s.ScanFileSize = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetStart(v int64) *DescribeSqlLogTasksResponseBodyDataList {
	s.Start = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetStatus(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetTaskId(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) SetTaskType(v string) *DescribeSqlLogTasksResponseBodyDataList {
	s.TaskType = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeSqlLogTasksResponseBodyDataListFilters struct {
	// The name of the filter parameter.
	//
	// >  For more information about the filter parameters, see the **Valid values of Key*	- section of this topic.
	//
	// example:
	//
	// delimiter
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter parameter.
	//
	// example:
	//
	// ,
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSqlLogTasksResponseBodyDataListFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTasksResponseBodyDataListFilters) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTasksResponseBodyDataListFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSqlLogTasksResponseBodyDataListFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSqlLogTasksResponseBodyDataListFilters) SetKey(v string) *DescribeSqlLogTasksResponseBodyDataListFilters {
	s.Key = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataListFilters) SetValue(v string) *DescribeSqlLogTasksResponseBodyDataListFilters {
	s.Value = &v
	return s
}

func (s *DescribeSqlLogTasksResponseBodyDataListFilters) Validate() error {
	return dara.Validate(s)
}
