// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobExecutionResponseBody
	GetCode() *int32
	SetData(v *GetJobExecutionResponseBodyData) *GetJobExecutionResponseBody
	GetData() *GetJobExecutionResponseBodyData
	SetMessage(v string) *GetJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobExecutionResponseBody
	GetSuccess() *bool
}

type GetJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetJobExecutionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobExecutionResponseBody) GetData() *GetJobExecutionResponseBodyData {
	return s.Data
}

func (s *GetJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobExecutionResponseBody) SetCode(v int32) *GetJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobExecutionResponseBody) SetData(v *GetJobExecutionResponseBodyData) *GetJobExecutionResponseBody {
	s.Data = v
	return s
}

func (s *GetJobExecutionResponseBody) SetMessage(v string) *GetJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobExecutionResponseBody) SetRequestId(v string) *GetJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobExecutionResponseBody) SetSuccess(v bool) *GetJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionResponseBodyData struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	Attempt  *int32  `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// {\\"Status\\": \\"NORMAL\\", \\"ActiveCount\\": 4, \\"UnavailableCount\\": 0, \\"ExpectedCount\\": 4, \\"RiskCount\\": 0}
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 74
	JobId   *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// /home/avatar/system/services/biz/payment/crontab/monitorpayment.php
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// []
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 1
	RouteStrategy *int32  `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	ScheduleTime  *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// example:
	//
	// 172.3.27.76
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// example:
	//
	// 2025-03-11T00:06:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// example:
	//
	// timer_schedule
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetJobExecutionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobExecutionResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetJobExecutionResponseBodyData) GetAttempt() *int32 {
	return s.Attempt
}

func (s *GetJobExecutionResponseBodyData) GetDataTime() *string {
	return s.DataTime
}

func (s *GetJobExecutionResponseBodyData) GetDuration() *int64 {
	return s.Duration
}

func (s *GetJobExecutionResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobExecutionResponseBodyData) GetExecutor() *string {
	return s.Executor
}

func (s *GetJobExecutionResponseBodyData) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetJobExecutionResponseBodyData) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobExecutionResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *GetJobExecutionResponseBodyData) GetJobType() *string {
	return s.JobType
}

func (s *GetJobExecutionResponseBodyData) GetParameters() *string {
	return s.Parameters
}

func (s *GetJobExecutionResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *GetJobExecutionResponseBodyData) GetRouteStrategy() *int32 {
	return s.RouteStrategy
}

func (s *GetJobExecutionResponseBodyData) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetJobExecutionResponseBodyData) GetServerIp() *string {
	return s.ServerIp
}

func (s *GetJobExecutionResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobExecutionResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetJobExecutionResponseBodyData) GetTimeType() *int32 {
	return s.TimeType
}

func (s *GetJobExecutionResponseBodyData) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *GetJobExecutionResponseBodyData) SetAppName(v string) *GetJobExecutionResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetAttempt(v int32) *GetJobExecutionResponseBodyData {
	s.Attempt = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetDataTime(v string) *GetJobExecutionResponseBodyData {
	s.DataTime = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetDuration(v int64) *GetJobExecutionResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetEndTime(v string) *GetJobExecutionResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetExecutor(v string) *GetJobExecutionResponseBodyData {
	s.Executor = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetJobExecutionId(v string) *GetJobExecutionResponseBodyData {
	s.JobExecutionId = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetJobId(v int64) *GetJobExecutionResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetJobName(v string) *GetJobExecutionResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetJobType(v string) *GetJobExecutionResponseBodyData {
	s.JobType = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetParameters(v string) *GetJobExecutionResponseBodyData {
	s.Parameters = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetResult(v string) *GetJobExecutionResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetRouteStrategy(v int32) *GetJobExecutionResponseBodyData {
	s.RouteStrategy = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetScheduleTime(v string) *GetJobExecutionResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetServerIp(v string) *GetJobExecutionResponseBodyData {
	s.ServerIp = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetStartTime(v string) *GetJobExecutionResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetStatus(v int32) *GetJobExecutionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetTimeType(v int32) *GetJobExecutionResponseBodyData {
	s.TimeType = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) SetTriggerType(v int32) *GetJobExecutionResponseBodyData {
	s.TriggerType = &v
	return s
}

func (s *GetJobExecutionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
