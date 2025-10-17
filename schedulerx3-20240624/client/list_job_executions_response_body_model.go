// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListJobExecutionsResponseBody
	GetCode() *int32
	SetData(v *ListJobExecutionsResponseBodyData) *ListJobExecutionsResponseBody
	GetData() *ListJobExecutionsResponseBodyData
	SetMessage(v string) *ListJobExecutionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListJobExecutionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJobExecutionsResponseBody
	GetSuccess() *bool
}

type ListJobExecutionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListJobExecutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6BCE89B3-E882-511D-9A75-D452A56EC4B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListJobExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListJobExecutionsResponseBody) GetData() *ListJobExecutionsResponseBodyData {
	return s.Data
}

func (s *ListJobExecutionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJobExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobExecutionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJobExecutionsResponseBody) SetCode(v int32) *ListJobExecutionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobExecutionsResponseBody) SetData(v *ListJobExecutionsResponseBodyData) *ListJobExecutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListJobExecutionsResponseBody) SetMessage(v string) *ListJobExecutionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobExecutionsResponseBody) SetRequestId(v string) *ListJobExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobExecutionsResponseBody) SetSuccess(v bool) *ListJobExecutionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobExecutionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobExecutionsResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListJobExecutionsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListJobExecutionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJobExecutionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJobExecutionsResponseBodyData) GetRecords() []*ListJobExecutionsResponseBodyDataRecords {
	return s.Records
}

func (s *ListJobExecutionsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListJobExecutionsResponseBodyData) SetPageNumber(v int32) *ListJobExecutionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobExecutionsResponseBodyData) SetPageSize(v int32) *ListJobExecutionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobExecutionsResponseBodyData) SetRecords(v []*ListJobExecutionsResponseBodyDataRecords) *ListJobExecutionsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListJobExecutionsResponseBodyData) SetTotal(v int32) *ListJobExecutionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListJobExecutionsResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobExecutionsResponseBodyDataRecords struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	Attempt *int32 `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	// example:
	//
	// 2024-11-12 14:52:42
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-11-12 14:52:42
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1827811800526000
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// name=zhangsan
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// example:
	//
	// 2024-11-12 14:52:42
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// example:
	//
	// 28.0.168.46
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TimeType    *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// example:
	//
	// 1
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// http://192.168.1.9:9999/
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s ListJobExecutionsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListJobExecutionsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetAttempt() *int32 {
	return s.Attempt
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetDataTime() *string {
	return s.DataTime
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetDuration() *int64 {
	return s.Duration
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetExecutor() *string {
	return s.Executor
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetJobId() *int64 {
	return s.JobId
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetJobName() *string {
	return s.JobName
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetJobType() *string {
	return s.JobType
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetParameters() *string {
	return s.Parameters
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetResult() *string {
	return s.Result
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetServerIp() *string {
	return s.ServerIp
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetStatus() *int32 {
	return s.Status
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetTimeType() *int32 {
	return s.TimeType
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *ListJobExecutionsResponseBodyDataRecords) GetWorkAddr() *string {
	return s.WorkAddr
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetAppName(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetAttempt(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.Attempt = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetDataTime(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.DataTime = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetDuration(v int64) *ListJobExecutionsResponseBodyDataRecords {
	s.Duration = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetEndTime(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.EndTime = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetExecutor(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.Executor = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobExecutionId(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.JobExecutionId = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobId(v int64) *ListJobExecutionsResponseBodyDataRecords {
	s.JobId = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobName(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetJobType(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.JobType = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetParameters(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.Parameters = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetResult(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.Result = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetRouteStrategy(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.RouteStrategy = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetScheduleTime(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.ScheduleTime = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetServerIp(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.ServerIp = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetStatus(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetTimeType(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.TimeType = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetTotalTokens(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.TotalTokens = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetTriggerType(v int32) *ListJobExecutionsResponseBodyDataRecords {
	s.TriggerType = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) SetWorkAddr(v string) *ListJobExecutionsResponseBodyDataRecords {
	s.WorkAddr = &v
	return s
}

func (s *ListJobExecutionsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
