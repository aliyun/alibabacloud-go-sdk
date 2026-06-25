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
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetJobExecutionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobExecutionResponseBodyData struct {
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of execution attempts.
	//
	// example:
	//
	// 1
	Attempt *int32 `json:"Attempt,omitempty" xml:"Attempt,omitempty"`
	// The data timestamp for the job instance.
	//
	// example:
	//
	// 2025-03-11 00:06:10
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The duration of the job execution.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the job execution ended.
	//
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Details of the executor that ran the job. The value is a JSON string.
	//
	// example:
	//
	// {\\"Status\\": \\"NORMAL\\", \\"ActiveCount\\": 4, \\"UnavailableCount\\": 0, \\"ExpectedCount\\": 4, \\"RiskCount\\": 0}
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The ID of the job execution.
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// The ID of the job.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// 天猫-自动审单
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The type of the job.
	//
	// example:
	//
	// xxljob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The parameters of the job.
	//
	// example:
	//
	// /home/avatar/system/services/biz/payment/crontab/monitorpayment.php
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The result of the job execution. The value is a JSON string.
	//
	// example:
	//
	// []
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The routing strategy. Valid values:
	//
	// - `1`: `Round Robin`
	//
	// - `2`: `Random`
	//
	// - `3`: `First`
	//
	// - `4`: `Last`
	//
	// - `5`: `Least Frequently Used`
	//
	// - `6`: `Least Recently Used`
	//
	// - `7`: `Consistent Hashing`
	//
	// - `8`: `Shard Broadcasting`
	//
	// example:
	//
	// 1
	RouteStrategy *int32 `json:"RouteStrategy,omitempty" xml:"RouteStrategy,omitempty"`
	// The scheduled time for the job execution.
	//
	// example:
	//
	// 2025-03-11 00:06:10
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The IP address of the scheduling server.
	//
	// example:
	//
	// 172.3.27.76
	ServerIp *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	// The time when the job execution started.
	//
	// example:
	//
	// 2025-03-11 00:06:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job execution status. Valid values:
	//
	// - `0`: `UNKNOWN`
	//
	// - `1`: `WAITING`
	//
	// - `2`: `READY`
	//
	// - `3`: `RUNNING`
	//
	// - `4`: `SUCCESS`
	//
	// - `5`: `FAILED`
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduling type of the job. Valid values:
	//
	// - `-1`: `none`
	//
	// - `1`: `cron`
	//
	// - `2`: `fixed_delay`
	//
	// - `3`: `fixed_rate`
	//
	// - `5`: `one_time`
	//
	// - `100`: `api`
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// Indicates how the job was triggered. Valid values:
	//
	// - `0`: `unknown`
	//
	// - `1`: `timer_schedule`
	//
	// - `2`: `rerun`
	//
	// - `3`: `api_run`
	//
	// - `4`: `user_retry`
	//
	// - `5`: `system_retry`
	//
	// - `6`: `manual`
	//
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
