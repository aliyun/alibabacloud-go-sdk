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
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryAiCallTaskPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B99C5955-5664-573D-97BE-A7CC1AFD8401
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	List []*QueryAiCallTaskPageResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 37
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 74
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// example:
	//
	// 11213132123123*****
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// example:
	//
	// 示例值
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// example:
	//
	// 72
	CallingCount *int64 `json:"CallingCount,omitempty" xml:"CallingCount,omitempty"`
	// example:
	//
	// 70%
	CompleteRate *string `json:"CompleteRate,omitempty" xml:"CompleteRate,omitempty"`
	// example:
	//
	// 10
	ConcurrentCount *int64 `json:"ConcurrentCount,omitempty" xml:"ConcurrentCount,omitempty"`
	// example:
	//
	// 1748932499000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 22
	DayCallCount *int64 `json:"DayCallCount,omitempty" xml:"DayCallCount,omitempty"`
	// example:
	//
	// 74.14%
	DayConnectRate *string `json:"DayConnectRate,omitempty" xml:"DayConnectRate,omitempty"`
	// example:
	//
	// 400
	DayImportCount *int64 `json:"DayImportCount,omitempty" xml:"DayImportCount,omitempty"`
	// example:
	//
	// 61
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// example:
	//
	// 95.89%
	HistoryConnectRate *string `json:"HistoryConnectRate,omitempty" xml:"HistoryConnectRate,omitempty"`
	// example:
	//
	// 1748932499000
	RealStartTime *int64 `json:"RealStartTime,omitempty" xml:"RealStartTime,omitempty"`
	// example:
	//
	// 示例值示例值
	StartFailedReason *string `json:"StartFailedReason,omitempty" xml:"StartFailedReason,omitempty"`
	// example:
	//
	// 1748932499000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 36
	SucceedCount *int64 `json:"SucceedCount,omitempty" xml:"SucceedCount,omitempty"`
	// example:
	//
	// 11121221121*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 58
	TotalCallCount *int64 `json:"TotalCallCount,omitempty" xml:"TotalCallCount,omitempty"`
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
