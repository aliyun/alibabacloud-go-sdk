// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobInstanceListResponseBody
	GetCode() *int32
	SetData(v *GetJobInstanceListResponseBodyData) *GetJobInstanceListResponseBody
	GetData() *GetJobInstanceListResponseBodyData
	SetMessage(v string) *GetJobInstanceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobInstanceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobInstanceListResponseBody
	GetSuccess() *bool
}

type GetJobInstanceListResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the job instances.
	Data *GetJobInstanceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetJobInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobInstanceListResponseBody) GetData() *GetJobInstanceListResponseBodyData {
	return s.Data
}

func (s *GetJobInstanceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobInstanceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobInstanceListResponseBody) SetCode(v int32) *GetJobInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobInstanceListResponseBody) SetData(v *GetJobInstanceListResponseBodyData) *GetJobInstanceListResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInstanceListResponseBody) SetMessage(v string) *GetJobInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobInstanceListResponseBody) SetRequestId(v string) *GetJobInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInstanceListResponseBody) SetSuccess(v bool) *GetJobInstanceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobInstanceListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobInstanceListResponseBodyData struct {
	// The details of the job instance.
	JobInstanceDetails []*GetJobInstanceListResponseBodyDataJobInstanceDetails `json:"JobInstanceDetails,omitempty" xml:"JobInstanceDetails,omitempty" type:"Repeated"`
}

func (s GetJobInstanceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponseBodyData) GetJobInstanceDetails() []*GetJobInstanceListResponseBodyDataJobInstanceDetails {
	return s.JobInstanceDetails
}

func (s *GetJobInstanceListResponseBodyData) SetJobInstanceDetails(v []*GetJobInstanceListResponseBodyDataJobInstanceDetails) *GetJobInstanceListResponseBodyData {
	s.JobInstanceDetails = v
	return s
}

func (s *GetJobInstanceListResponseBodyData) Validate() error {
	if s.JobInstanceDetails != nil {
		for _, item := range s.JobInstanceDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobInstanceListResponseBodyDataJobInstanceDetails struct {
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
	// The trigger type of the job instance. Valid values:
	//
	// 	- **1**: The job instance was triggered at the scheduled time.
	//
	// 	- **2**: The job instance was triggered due to data updates.
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

func (s GetJobInstanceListResponseBodyDataJobInstanceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetJobInstanceListResponseBodyDataJobInstanceDetails) GoString() string {
	return s.String()
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetDataTime() *string {
	return s.DataTime
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetExecutor() *string {
	return s.Executor
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetJobId() *int64 {
	return s.JobId
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetProgress() *string {
	return s.Progress
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetResult() *string {
	return s.Result
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetStatus() *int32 {
	return s.Status
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetTimeType() *int32 {
	return s.TimeType
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) GetWorkAddr() *string {
	return s.WorkAddr
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetDataTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.DataTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetEndTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.EndTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetExecutor(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Executor = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetInstanceId(v int64) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.InstanceId = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetJobId(v int64) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.JobId = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetProgress(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Progress = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetResult(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Result = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetScheduleTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.ScheduleTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetStartTime(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.StartTime = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetStatus(v int32) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.Status = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetTimeType(v int32) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.TimeType = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetTriggerType(v int32) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.TriggerType = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) SetWorkAddr(v string) *GetJobInstanceListResponseBodyDataJobInstanceDetails {
	s.WorkAddr = &v
	return s
}

func (s *GetJobInstanceListResponseBodyDataJobInstanceDetails) Validate() error {
	return dara.Validate(s)
}
