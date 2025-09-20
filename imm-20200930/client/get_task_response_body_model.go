// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskResponseBody
	GetCode() *string
	SetEndTime(v string) *GetTaskResponseBody
	GetEndTime() *string
	SetEventId(v string) *GetTaskResponseBody
	GetEventId() *string
	SetMessage(v string) *GetTaskResponseBody
	GetMessage() *string
	SetProgress(v int32) *GetTaskResponseBody
	GetProgress() *int32
	SetProjectName(v string) *GetTaskResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetTaskResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetTaskResponseBody
	GetStatus() *string
	SetTags(v map[string]interface{}) *GetTaskResponseBody
	GetTags() map[string]interface{}
	SetTaskId(v string) *GetTaskResponseBody
	GetTaskId() *string
	SetTaskRequestDefinition(v string) *GetTaskResponseBody
	GetTaskRequestDefinition() *string
	SetTaskType(v string) *GetTaskResponseBody
	GetTaskType() *string
	SetUserData(v string) *GetTaskResponseBody
	GetUserData() *string
}

type GetTaskResponseBody struct {
	// The error code of the task.
	//
	// example:
	//
	// ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2021-12-24T03:01:49.480109219Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 2F6-1Bz99Xi93EnRpNEyLudILJm****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The error message of the task.
	//
	// example:
	//
	// The specified resource project is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task progress. Valid values: 0 to 100. Unit: %.
	//
	// >  This parameter is valid only if the task is in the `Running` state.``
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The project name.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2C5C1E0F-D8B8-4DA0-8127-EC32C771****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2021-12-24T03:01:41.662060377Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- RUNNING: The task is running.
	//
	// 	- Succeeded: The task is successful.
	//
	// 	- Failed: The task failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags. This parameter is returned only if you specified Tags when you created the task.
	//
	// example:
	//
	// {"test": "val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The task ID.
	//
	// example:
	//
	// c2b277b9-0d30-4882-ad6d-ad661382****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The initial request parameters used to create the task.
	//
	// example:
	//
	// {
	//
	// 	"ProjectName":"test-project",
	//
	// 	"CompressedFormat":"zip",
	//
	// 	"TargetURI":"oss://test-bucket/output/test.zip",
	//
	// 	"Sources":[{"URI":"oss://test-bucket/input/test.jpg"}]
	//
	// }
	TaskRequestDefinition *string `json:"TaskRequestDefinition,omitempty" xml:"TaskRequestDefinition,omitempty"`
	// The type of the task. For more information, see [Task types](https://help.aliyun.com/document_detail/2743993.html).
	//
	// example:
	//
	// VideoLabelClassification
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The user data of the task.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *GetTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *GetTaskResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBody) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResponseBody) GetTaskRequestDefinition() *string {
	return s.TaskRequestDefinition
}

func (s *GetTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetTaskResponseBody) SetCode(v string) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetEndTime(v string) *GetTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBody) SetEventId(v string) *GetTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetProgress(v int32) *GetTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *GetTaskResponseBody) SetProjectName(v string) *GetTaskResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetStartTime(v string) *GetTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBody) SetStatus(v string) *GetTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBody) SetTags(v map[string]interface{}) *GetTaskResponseBody {
	s.Tags = v
	return s
}

func (s *GetTaskResponseBody) SetTaskId(v string) *GetTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBody) SetTaskRequestDefinition(v string) *GetTaskResponseBody {
	s.TaskRequestDefinition = &v
	return s
}

func (s *GetTaskResponseBody) SetTaskType(v string) *GetTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetTaskResponseBody) SetUserData(v string) *GetTaskResponseBody {
	s.UserData = &v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
