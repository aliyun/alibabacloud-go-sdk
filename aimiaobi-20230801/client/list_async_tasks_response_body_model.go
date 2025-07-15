// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAsyncTasksResponseBody
	GetCode() *string
	SetCurrent(v int32) *ListAsyncTasksResponseBody
	GetCurrent() *int32
	SetData(v []*ListAsyncTasksResponseBodyData) *ListAsyncTasksResponseBody
	GetData() []*ListAsyncTasksResponseBodyData
	SetHttpStatusCode(v int32) *ListAsyncTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAsyncTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAsyncTasksResponseBody
	GetRequestId() *string
	SetSize(v int32) *ListAsyncTasksResponseBody
	GetSize() *int32
	SetSuccess(v bool) *ListAsyncTasksResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListAsyncTasksResponseBody
	GetTotal() *int32
}

type ListAsyncTasksResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Current *int32                            `json:"Current,omitempty" xml:"Current,omitempty"`
	Data    []*ListAsyncTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 428DCC0D-3C63-5306-BD1B-124396AB97BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAsyncTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAsyncTasksResponseBody) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAsyncTasksResponseBody) GetData() []*ListAsyncTasksResponseBodyData {
	return s.Data
}

func (s *ListAsyncTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAsyncTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAsyncTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsyncTasksResponseBody) GetSize() *int32 {
	return s.Size
}

func (s *ListAsyncTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAsyncTasksResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAsyncTasksResponseBody) SetCode(v string) *ListAsyncTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetCurrent(v int32) *ListAsyncTasksResponseBody {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetData(v []*ListAsyncTasksResponseBodyData) *ListAsyncTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListAsyncTasksResponseBody) SetHttpStatusCode(v int32) *ListAsyncTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetMessage(v string) *ListAsyncTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetRequestId(v string) *ListAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetSize(v int32) *ListAsyncTasksResponseBody {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetSuccess(v bool) *ListAsyncTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetTotal(v int32) *ListAsyncTasksResponseBody {
	s.Total = &v
	return s
}

func (s *ListAsyncTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAsyncTasksResponseBodyData struct {
	// example:
	//
	// 2020-12-23 15:41:58
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1111
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// {}
	TaskDefinition *string `json:"TaskDefinition,omitempty" xml:"TaskDefinition,omitempty"`
	// example:
	//
	// 2023-03-09 00:00:00
	TaskEndTime      *string `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	TaskErrorMessage *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	// example:
	//
	// 2023-10-14 14:30:00
	TaskExecuteTime *string `json:"TaskExecuteTime,omitempty" xml:"TaskExecuteTime,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId                *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInnerErrorMessage *string `json:"TaskInnerErrorMessage,omitempty" xml:"TaskInnerErrorMessage,omitempty"`
	// example:
	//
	// {}
	TaskIntermediateResult *string `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	TaskName               *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// {}
	TaskParam *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	// example:
	//
	// {}
	TaskProgressMessage *string `json:"TaskProgressMessage,omitempty" xml:"TaskProgressMessage,omitempty"`
	// example:
	//
	// {}
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// example:
	//
	// 1
	TaskRetryCount *string `json:"TaskRetryCount,omitempty" xml:"TaskRetryCount,omitempty"`
	// example:
	//
	// 2023-03-20 10:53:00
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// test
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 2023-02-16 10:29:16
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 111
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListAsyncTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAsyncTasksResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListAsyncTasksResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAsyncTasksResponseBodyData) GetTaskCode() *string {
	return s.TaskCode
}

func (s *ListAsyncTasksResponseBodyData) GetTaskDefinition() *string {
	return s.TaskDefinition
}

func (s *ListAsyncTasksResponseBodyData) GetTaskEndTime() *string {
	return s.TaskEndTime
}

func (s *ListAsyncTasksResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *ListAsyncTasksResponseBodyData) GetTaskExecuteTime() *string {
	return s.TaskExecuteTime
}

func (s *ListAsyncTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAsyncTasksResponseBodyData) GetTaskInnerErrorMessage() *string {
	return s.TaskInnerErrorMessage
}

func (s *ListAsyncTasksResponseBodyData) GetTaskIntermediateResult() *string {
	return s.TaskIntermediateResult
}

func (s *ListAsyncTasksResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListAsyncTasksResponseBodyData) GetTaskParam() *string {
	return s.TaskParam
}

func (s *ListAsyncTasksResponseBodyData) GetTaskProgressMessage() *string {
	return s.TaskProgressMessage
}

func (s *ListAsyncTasksResponseBodyData) GetTaskResult() *string {
	return s.TaskResult
}

func (s *ListAsyncTasksResponseBodyData) GetTaskRetryCount() *string {
	return s.TaskRetryCount
}

func (s *ListAsyncTasksResponseBodyData) GetTaskStartTime() *string {
	return s.TaskStartTime
}

func (s *ListAsyncTasksResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListAsyncTasksResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAsyncTasksResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAsyncTasksResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *ListAsyncTasksResponseBodyData) SetCreateTime(v string) *ListAsyncTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetCreateUser(v string) *ListAsyncTasksResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetId(v int64) *ListAsyncTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskCode(v string) *ListAsyncTasksResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskDefinition(v string) *ListAsyncTasksResponseBodyData {
	s.TaskDefinition = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskEndTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskEndTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskErrorMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskExecuteTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskExecuteTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskId(v string) *ListAsyncTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskInnerErrorMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskInnerErrorMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskIntermediateResult(v string) *ListAsyncTasksResponseBodyData {
	s.TaskIntermediateResult = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskName(v string) *ListAsyncTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskParam(v string) *ListAsyncTasksResponseBodyData {
	s.TaskParam = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskProgressMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskProgressMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskResult(v string) *ListAsyncTasksResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskRetryCount(v string) *ListAsyncTasksResponseBodyData {
	s.TaskRetryCount = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskStartTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskStartTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskStatus(v int32) *ListAsyncTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskType(v string) *ListAsyncTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetUpdateTime(v string) *ListAsyncTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetUpdateUser(v string) *ListAsyncTasksResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
