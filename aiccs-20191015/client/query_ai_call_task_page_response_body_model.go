// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallTaskPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryAiCallTaskPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryAiCallTaskPageResponseBody
	GetCode() *string
	SetData(v *QueryAiCallTaskPageResponseBodyData) *QueryAiCallTaskPageResponseBody
	GetData() *QueryAiCallTaskPageResponseBodyData
	SetMessage(v string) *QueryAiCallTaskPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAiCallTaskPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAiCallTaskPageResponseBody
	GetSuccess() *bool
}

type QueryAiCallTaskPageResponseBody struct {
	// The detailed reason for access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task details.
	Data *QueryAiCallTaskPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// Invalid parameter.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B99C5955-5664-573D-97BE-A7CC1AFD8401
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: successful.
	//
	// - **false**: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAiCallTaskPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryAiCallTaskPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAiCallTaskPageResponseBody) GetData() *QueryAiCallTaskPageResponseBodyData {
	return s.Data
}

func (s *QueryAiCallTaskPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAiCallTaskPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAiCallTaskPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAiCallTaskPageResponseBody) SetAccessDeniedDetail(v string) *QueryAiCallTaskPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBody) SetCode(v string) *QueryAiCallTaskPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBody) SetData(v *QueryAiCallTaskPageResponseBodyData) *QueryAiCallTaskPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryAiCallTaskPageResponseBody) SetMessage(v string) *QueryAiCallTaskPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBody) SetRequestId(v string) *QueryAiCallTaskPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBody) SetSuccess(v bool) *QueryAiCallTaskPageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAiCallTaskPageResponseBodyData struct {
	// The task data.
	List []*QueryAiCallTaskPageResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 37
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 74
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 72
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryAiCallTaskPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskPageResponseBodyData) GetList() []*QueryAiCallTaskPageResponseBodyDataList {
	return s.List
}

func (s *QueryAiCallTaskPageResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryAiCallTaskPageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryAiCallTaskPageResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryAiCallTaskPageResponseBodyData) SetList(v []*QueryAiCallTaskPageResponseBodyDataList) *QueryAiCallTaskPageResponseBodyData {
	s.List = v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyData) SetPageNo(v int64) *QueryAiCallTaskPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyData) SetPageSize(v int64) *QueryAiCallTaskPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyData) SetTotal(v int64) *QueryAiCallTaskPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAiCallTaskPageResponseBodyDataList struct {
	// The agent ID.
	//
	// example:
	//
	// 1180**************
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// TestAgent.
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// Sample value.
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// example:
	//
	// Sample value.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// Sample value.
	CallExpireDate *string `json:"CallExpireDate,omitempty" xml:"CallExpireDate,omitempty"`
	// example:
	//
	// 39
	CallExpireMinutes *int64 `json:"CallExpireMinutes,omitempty" xml:"CallExpireMinutes,omitempty"`
	// example:
	//
	// 72
	CallExpireType *int64 `json:"CallExpireType,omitempty" xml:"CallExpireType,omitempty"`
	// The number of ongoing calls.
	//
	// example:
	//
	// 72
	CallingCount *int64 `json:"CallingCount,omitempty" xml:"CallingCount,omitempty"`
	// The task completion rate.
	//
	// example:
	//
	// 70%
	CompleteRate *string `json:"CompleteRate,omitempty" xml:"CompleteRate,omitempty"`
	// The task concurrency.
	//
	// example:
	//
	// 10
	ConcurrentCount *int64 `json:"ConcurrentCount,omitempty" xml:"ConcurrentCount,omitempty"`
	// The creation time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1748932499000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of calls made on the current day.
	//
	// example:
	//
	// 22
	DayCallCount *int64 `json:"DayCallCount,omitempty" xml:"DayCallCount,omitempty"`
	// The daily connection rate. Daily connection rate = number of connections on the current day ÷ number of calls on the current day (DayCallCount).
	//
	// example:
	//
	// 74.14%
	DayConnectRate *string `json:"DayConnectRate,omitempty" xml:"DayConnectRate,omitempty"`
	// The amount of data imported on the current day.
	//
	// example:
	//
	// 400
	DayImportCount *int64 `json:"DayImportCount,omitempty" xml:"DayImportCount,omitempty"`
	// The total number of failed task executions.
	//
	// example:
	//
	// 61
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The historical connection rate. Historical connection rate = historical number of connections ÷ total number of calls (TotalCallCount).
	//
	// example:
	//
	// 95.89%
	HistoryConnectRate *string `json:"HistoryConnectRate,omitempty" xml:"HistoryConnectRate,omitempty"`
	// The actual start time of the task. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1748932499000
	RealStartTime *int64 `json:"RealStartTime,omitempty" xml:"RealStartTime,omitempty"`
	// The reason for startup failure.
	//
	// example:
	//
	// Agent is offline.
	StartFailedReason *string `json:"StartFailedReason,omitempty" xml:"StartFailedReason,omitempty"`
	// The scheduled start time of the task. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1748932499000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status.
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of successful task executions.
	//
	// example:
	//
	// 36
	SucceedCount *int64 `json:"SucceedCount,omitempty" xml:"SucceedCount,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1187**************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// TestTask.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The total number of calls made by the task.
	//
	// example:
	//
	// 58
	TotalCallCount *int64 `json:"TotalCallCount,omitempty" xml:"TotalCallCount,omitempty"`
	// The total number of task items.
	//
	// example:
	//
	// 71
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryAiCallTaskPageResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetCallExpireDate() *string {
	return s.CallExpireDate
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetCallExpireMinutes() *int64 {
	return s.CallExpireMinutes
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetCallExpireType() *int64 {
	return s.CallExpireType
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetCallingCount() *int64 {
	return s.CallingCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetCompleteRate() *string {
	return s.CompleteRate
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetConcurrentCount() *int64 {
	return s.ConcurrentCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetDayCallCount() *int64 {
	return s.DayCallCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetDayConnectRate() *string {
	return s.DayConnectRate
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetDayImportCount() *int64 {
	return s.DayImportCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetHistoryConnectRate() *string {
	return s.HistoryConnectRate
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetRealStartTime() *int64 {
	return s.RealStartTime
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetStartFailedReason() *string {
	return s.StartFailedReason
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetStatus() *int64 {
	return s.Status
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetSucceedCount() *int64 {
	return s.SucceedCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetTotalCallCount() *int64 {
	return s.TotalCallCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetAgentId(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetAgentName(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetApplicationCode(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.ApplicationCode = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetApplicationName(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.ApplicationName = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetCallExpireDate(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.CallExpireDate = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetCallExpireMinutes(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.CallExpireMinutes = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetCallExpireType(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.CallExpireType = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetCallingCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.CallingCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetCompleteRate(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.CompleteRate = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetConcurrentCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.ConcurrentCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetCreateTime(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetDayCallCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.DayCallCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetDayConnectRate(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.DayConnectRate = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetDayImportCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.DayImportCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetFailedCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.FailedCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetHistoryConnectRate(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.HistoryConnectRate = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetRealStartTime(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.RealStartTime = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetStartFailedReason(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.StartFailedReason = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetStartTime(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetStatus(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetSucceedCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.SucceedCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetTaskId(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetTaskName(v string) *QueryAiCallTaskPageResponseBodyDataList {
	s.TaskName = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetTotalCallCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.TotalCallCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) SetTotalCount(v int64) *QueryAiCallTaskPageResponseBodyDataList {
	s.TotalCount = &v
	return s
}

func (s *QueryAiCallTaskPageResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
