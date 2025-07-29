// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobInstanceResponseBody
	GetCode() *int32
	SetData(v *GetJobInstanceResponseBodyData) *GetJobInstanceResponseBody
	GetData() *GetJobInstanceResponseBodyData
	SetMessage(v string) *GetJobInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobInstanceResponseBody
	GetSuccess() *bool
}

type GetJobInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the job instance.
	Data *GetJobInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned only if the corresponding error occurs.
	//
	// example:
	//
	// jobid: 92583 not match groupId: testSchedulerx.defaultGroup
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobInstanceResponseBody) GetData() *GetJobInstanceResponseBodyData {
	return s.Data
}

func (s *GetJobInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobInstanceResponseBody) SetCode(v int32) *GetJobInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobInstanceResponseBody) SetData(v *GetJobInstanceResponseBodyData) *GetJobInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInstanceResponseBody) SetMessage(v string) *GetJobInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobInstanceResponseBody) SetRequestId(v string) *GetJobInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInstanceResponseBody) SetSuccess(v bool) *GetJobInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobInstanceResponseBodyData struct {
	// The details of the job instance.
	JobInstanceDetail *GetJobInstanceResponseBodyDataJobInstanceDetail `json:"JobInstanceDetail,omitempty" xml:"JobInstanceDetail,omitempty" type:"Struct"`
}

func (s GetJobInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponseBodyData) GetJobInstanceDetail() *GetJobInstanceResponseBodyDataJobInstanceDetail {
	return s.JobInstanceDetail
}

func (s *GetJobInstanceResponseBodyData) SetJobInstanceDetail(v *GetJobInstanceResponseBodyDataJobInstanceDetail) *GetJobInstanceResponseBodyData {
	s.JobInstanceDetail = v
	return s
}

func (s *GetJobInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetJobInstanceResponseBodyDataJobInstanceDetail struct {
	// The data timestamp of the job instance.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the job execution.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The user who executes the job.
	//
	// example:
	//
	// A
	Executor *string `json:"Executor,omitempty" xml:"Executor,omitempty"`
	// The job instance ID.
	//
	// example:
	//
	// 11111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The job name.
	//
	// example:
	//
	// ManualJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The parameters of the job instance.
	//
	// example:
	//
	// {\\"alertId\\":11111}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The progress of the job instance.
	//
	// example:
	//
	// complete
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The execution results of the job instance.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The time when the job was scheduled to run.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The start time of the job execution.
	//
	// example:
	//
	// 2020-07-27 11:52:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the job instance. Valid values:
	//
	// 	- **1**: The job instance is waiting for execution.
	//
	// 	- **3**: The job instance is running.
	//
	// 	- **4**: The job instance is successful.
	//
	// 	- **5**: The job instance failed.
	//
	// 	- **9**: The job instance is rejected.
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.InstanceStatus
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The method that is used to specify the time when to schedule the job instance. Valid values:
	//
	// 	- **1**: cron
	//
	// 	- **3**: fix_rate
	//
	// 	- **4**: second_delay
	//
	// 	- **100**: api
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.TimeType
	//
	// example:
	//
	// 1
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
	// The trace ID, which can be used to query trace details.
	//
	// example:
	//
	// 210e845016596663430048015d0a2c
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The trigger type of the job instance. Valid values:
	//
	// 	- **1**: The job instance was triggered at the scheduled time.
	//
	// 	- **2**: The job instance was triggered due to data update.
	//
	// 	- **3**: The job instance was triggered by an API call.
	//
	// 	- **4**: The job instance was triggered because it is manually rerun.
	//
	// 	- **5**: The job instance was triggered because the system automatically reruns the job instance upon a system exception, such as a database exception.
	//
	// Enumeration class: com.alibaba.schedulerx.common.domain.TriggerType
	//
	// example:
	//
	// 3
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The endpoint of the triggered client. The value is in the IP address:Port number format.
	//
	// example:
	//
	// 192.168.0.0:16
	WorkAddr *string `json:"WorkAddr,omitempty" xml:"WorkAddr,omitempty"`
}

func (s GetJobInstanceResponseBodyDataJobInstanceDetail) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceResponseBodyDataJobInstanceDetail) GoString() string {
	return s.String()
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetDataTime() *string {
	return s.DataTime
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetExecutor() *string {
	return s.Executor
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetJobName() *string {
	return s.JobName
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetParameters() *string {
	return s.Parameters
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetProgress() *string {
	return s.Progress
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetResult() *string {
	return s.Result
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetStatus() *int32 {
	return s.Status
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetTimeType() *int32 {
	return s.TimeType
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetTraceId() *string {
	return s.TraceId
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) GetWorkAddr() *string {
	return s.WorkAddr
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetDataTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.DataTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetEndTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.EndTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetExecutor(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Executor = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetInstanceId(v int64) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.InstanceId = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetJobId(v int64) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetJobName(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.JobName = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetParameters(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Parameters = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetProgress(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Progress = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetResult(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Result = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetScheduleTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.ScheduleTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetStartTime(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.StartTime = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetStatus(v int32) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.Status = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetTimeType(v int32) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.TimeType = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetTraceId(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.TraceId = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetTriggerType(v int32) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.TriggerType = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) SetWorkAddr(v string) *GetJobInstanceResponseBodyDataJobInstanceDetail {
	s.WorkAddr = &v
	return s
}

func (s *GetJobInstanceResponseBodyDataJobInstanceDetail) Validate() error {
	return dara.Validate(s)
}
