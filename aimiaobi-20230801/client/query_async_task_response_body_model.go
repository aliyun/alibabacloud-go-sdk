// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAsyncTaskResponseBody
	GetCode() *string
	SetData(v *QueryAsyncTaskResponseBodyData) *QueryAsyncTaskResponseBody
	GetData() *QueryAsyncTaskResponseBodyData
	SetHttpStatusCode(v int32) *QueryAsyncTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryAsyncTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAsyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAsyncTaskResponseBody
	GetSuccess() *bool
}

type QueryAsyncTaskResponseBody struct {
	// Status code
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Business data
	Data *QueryAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error description
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Unique request ID
	//
	// example:
	//
	// 867C4ABE-4381-5BC2-9810-5A5F334F71CF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success status: true for success, false for failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAsyncTaskResponseBody) GetData() *QueryAsyncTaskResponseBodyData {
	return s.Data
}

func (s *QueryAsyncTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryAsyncTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAsyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAsyncTaskResponseBody) SetCode(v string) *QueryAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetData(v *QueryAsyncTaskResponseBodyData) *QueryAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetHttpStatusCode(v int32) *QueryAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetMessage(v string) *QueryAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetRequestId(v string) *QueryAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetSuccess(v bool) *QueryAsyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryAsyncTaskResponseBodyData struct {
	// Creation date
	//
	// example:
	//
	// 2021-07-25 14:34:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Creator
	//
	// example:
	//
	// "12121"
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Task ID, indicates the specific task.
	//
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// Task execution error message
	//
	// example:
	//
	// error
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// Unique task ID
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Intermediate task execution result. When a task has multiple steps, save the output of each step here. When resuming from a pause, read the intermediate result and continue from there.
	//
	// example:
	//
	// "{}"
	TaskIntermediateResult *string `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	// Task name
	//
	// example:
	//
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Task execution input parameters, JSON format
	//
	// example:
	//
	// "{\\"fileKey\\":\\"oss://default/xxxx/xxxx/xxx\\",\\"fileName\\":\\"xxxxx.doc\\"}"
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// Task execution progress message
	//
	// example:
	//
	// "{}"
	TaskProgressMessage *string `json:"TaskProgressMessage,omitempty" xml:"TaskProgressMessage,omitempty"`
	// Task execution result message
	//
	// example:
	//
	// "{}"
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// Number of task retries
	//
	// example:
	//
	// "3"
	TaskRetryCount *string `json:"TaskRetryCount,omitempty" xml:"TaskRetryCount,omitempty"`
	// Task execution status: 0-Pending, 1-Executing, 2-Execution successful, 3-Paused, 4-Execution failed (retryable), 5-Execution failed (not retryable), 6-Task canceled.
	//
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// Update date
	//
	// example:
	//
	// 2023-04-27 18:07:43
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Updater
	//
	// example:
	//
	// "12121"
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s QueryAsyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryAsyncTaskResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskCode() *string {
	return s.TaskCode
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskIntermediateResult() *string {
	return s.TaskIntermediateResult
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskParam() *string {
	return s.TaskParam
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskProgressMessage() *string {
	return s.TaskProgressMessage
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskResult() *string {
	return s.TaskResult
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskRetryCount() *string {
	return s.TaskRetryCount
}

func (s *QueryAsyncTaskResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryAsyncTaskResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *QueryAsyncTaskResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *QueryAsyncTaskResponseBodyData) SetCreateTime(v string) *QueryAsyncTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetCreateUser(v string) *QueryAsyncTaskResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskCode(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskErrorMessage(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskId(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskIntermediateResult(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskIntermediateResult = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskName(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskParam(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskParam = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskProgressMessage(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskProgressMessage = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskResult(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskRetryCount(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskRetryCount = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskStatus(v int32) *QueryAsyncTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetUpdateTime(v string) *QueryAsyncTaskResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetUpdateUser(v string) *QueryAsyncTaskResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
