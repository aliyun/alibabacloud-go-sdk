// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeTasksResponseBodyData) *DescribeTasksResponseBody
	GetData() *DescribeTasksResponseBodyData
	SetRequestId(v string) *DescribeTasksResponseBody
	GetRequestId() *string
}

type DescribeTasksResponseBody struct {
	// The returned data.
	Data *DescribeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) GetData() *DescribeTasksResponseBodyData {
	return s.Data
}

func (s *DescribeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTasksResponseBody) SetData(v *DescribeTasksResponseBodyData) *DescribeTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTasksResponseBodyData struct {
	// The list of tasks.
	Content []*DescribeTasksResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token used to retrieve the next page of results. If this parameter is empty, no more results are available.
	//
	// example:
	//
	// f4b8c2504545a3b41af5e75147d17d12e3818a0b9b2ff9a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries matching the request criteria. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyData) GetContent() []*DescribeTasksResponseBodyDataContent {
	return s.Content
}

func (s *DescribeTasksResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTasksResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTasksResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeTasksResponseBodyData) SetContent(v []*DescribeTasksResponseBodyDataContent) *DescribeTasksResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeTasksResponseBodyData) SetMaxResults(v int32) *DescribeTasksResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetNextToken(v string) *DescribeTasksResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTotalCount(v int64) *DescribeTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeTasksResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTasksResponseBodyDataContent struct {
	// The task\\"s completion time, represented as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1724983927
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The error message returned if the task fails.
	//
	// example:
	//
	// device not online
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The execution ID. This parameter is currently unused and returns an empty string.
	//
	// example:
	//
	// empty
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The task\\"s expiration time, represented as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1724983927
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The progress of the task, ranging from 0 to 10,000.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task\\"s start time, represented as a Unix timestamp in seconds.
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
	// The task details.
	//
	// example:
	//
	// {"resourceTypes":["ACS::ECS::Instance","ACS::OSS::Bucket","ACS::OTS::Instance","ACS::NAS::FileSystem"]}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-0000e4w0u1v592zdf6s7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// empty
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task priority. Valid values: `HIGH` (high-priority, for user-initiated tasks) and `LOW` (low-priority, for background tasks).
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
	// The task status. Valid values: `CREATED`, `RUNNING`, `COMPLETE`, `FAILED`, `EXPIRED`, and `CANCELED`.
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The task type. Valid values: `TEST`, `UPDATE_RESOURCES`, and `CHECK_RULES`.
	//
	// example:
	//
	// UPDATE_RESOURCES
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeTasksResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeTasksResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyDataContent) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *DescribeTasksResponseBodyDataContent) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeTasksResponseBodyDataContent) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *DescribeTasksResponseBodyDataContent) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeTasksResponseBodyDataContent) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeTasksResponseBodyDataContent) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTasksResponseBodyDataContent) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskDetail() *string {
	return s.TaskDetail
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskPriority() *string {
	return s.TaskPriority
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskResult() *string {
	return s.TaskResult
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeTasksResponseBodyDataContent) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTasksResponseBodyDataContent) SetCompleteTime(v int64) *DescribeTasksResponseBodyDataContent {
	s.CompleteTime = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetErrorMessage(v string) *DescribeTasksResponseBodyDataContent {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetExecutionId(v string) *DescribeTasksResponseBodyDataContent {
	s.ExecutionId = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetExpireTime(v int64) *DescribeTasksResponseBodyDataContent {
	s.ExpireTime = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetProgress(v int32) *DescribeTasksResponseBodyDataContent {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetRequestId(v string) *DescribeTasksResponseBodyDataContent {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetStartTime(v int64) *DescribeTasksResponseBodyDataContent {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskDescription(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskDescription = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskDetail(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskDetail = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskId(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskName(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskName = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskPriority(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskPriority = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskResult(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskResult = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskStatus(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskType(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskType = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
