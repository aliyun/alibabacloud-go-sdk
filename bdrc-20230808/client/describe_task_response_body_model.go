// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeTaskResponseBodyData) *DescribeTaskResponseBody
	GetData() *DescribeTaskResponseBodyData
	SetRequestId(v string) *DescribeTaskResponseBody
	GetRequestId() *string
}

type DescribeTaskResponseBody struct {
	// The returned data.
	Data *DescribeTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 14DFF801-A4E3-5136-AAB8-7D246012CD7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) GetData() *DescribeTaskResponseBodyData {
	return s.Data
}

func (s *DescribeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskResponseBody) SetData(v *DescribeTaskResponseBodyData) *DescribeTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTaskResponseBodyData struct {
	// The completion time, formatted as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1724983927
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The error message returned if the task fails.
	//
	// example:
	//
	// too many requests.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The execution ID. This parameter is deprecated and always returns an empty string.
	//
	// example:
	//
	// empty
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The expiration time, formatted as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1719026680
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The task\\"s progress, measured on a scale of 0 to 10,000.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 784076D6-BD6D-5564-9CEA-834EB11F0C62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time, formatted as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1724983927
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task description.
	//
	// example:
	//
	// empty
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// The details of the task execution.
	//
	// example:
	//
	// {"resourceTypes":["ACS::ECS::Instance","ACS::OSS::Bucket","ACS::OTS::Instance","ACS::NAS::FileSystem"]}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// t-xxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test5566
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task priority. Valid values: `HIGH` (typically for user-initiated tasks) and `LOW` (typically for background tasks).
	//
	// example:
	//
	// HIGH
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// The result of the task.
	//
	// example:
	//
	// {"resourceCounts":[{"resourceType":"ACS::OSS::Bucket","count":2},{"resourceType":"ACS::NAS::FileSystem","count":3}]}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The status of the task. Valid values: `CREATED` (Created), `RUNNING` (Running), `COMPLETE` (Completed), `FAILED` (Failed), `EXPIRED` (Expired), and `CANCELED` (Canceled).
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The task type. Valid values: `UPDATE_RESOURCES` (updates resources and their statuses) and `CHECK_RULES`.
	//
	// example:
	//
	// UPDATE_RESOURCES
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodyData) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeTaskResponseBodyData) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *DescribeTaskResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeTaskResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeTaskResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeTaskResponseBodyData) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *DescribeTaskResponseBodyData) GetTaskDetail() *string {
	return s.TaskDetail
}

func (s *DescribeTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeTaskResponseBodyData) GetTaskPriority() *string {
	return s.TaskPriority
}

func (s *DescribeTaskResponseBodyData) GetTaskResult() *string {
	return s.TaskResult
}

func (s *DescribeTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskResponseBodyData) SetCompleteTime(v int64) *DescribeTaskResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetErrorMessage(v string) *DescribeTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetExecutionId(v string) *DescribeTaskResponseBodyData {
	s.ExecutionId = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetExpireTime(v int64) *DescribeTaskResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetProgress(v int32) *DescribeTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetRequestId(v string) *DescribeTaskResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetStartTime(v int64) *DescribeTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskDescription(v string) *DescribeTaskResponseBodyData {
	s.TaskDescription = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskDetail(v string) *DescribeTaskResponseBodyData {
	s.TaskDetail = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskId(v string) *DescribeTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskName(v string) *DescribeTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskPriority(v string) *DescribeTaskResponseBodyData {
	s.TaskPriority = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskResult(v string) *DescribeTaskResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskStatus(v string) *DescribeTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskType(v string) *DescribeTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
