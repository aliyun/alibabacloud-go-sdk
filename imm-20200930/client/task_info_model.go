// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TaskInfo
	GetCode() *string
	SetEndTime(v string) *TaskInfo
	GetEndTime() *string
	SetMessage(v string) *TaskInfo
	GetMessage() *string
	SetProgress(v int32) *TaskInfo
	GetProgress() *int32
	SetStartTime(v string) *TaskInfo
	GetStartTime() *string
	SetStatus(v string) *TaskInfo
	GetStatus() *string
	SetTags(v map[string]interface{}) *TaskInfo
	GetTags() map[string]interface{}
	SetTaskId(v string) *TaskInfo
	GetTaskId() *string
	SetTaskRequestDefinition(v string) *TaskInfo
	GetTaskRequestDefinition() *string
	SetTaskType(v string) *TaskInfo
	GetTaskType() *string
	SetUserData(v string) *TaskInfo
	GetUserData() *string
}

type TaskInfo struct {
	// The error code.
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
	// The error message.
	//
	// example:
	//
	// The specified resource project is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The progress of the task.
	//
	// example:
	//
	// 0
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2021-12-24T03:01:41.662060377Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- Running: The task is running.
	//
	// 	- Succeeded: The task is successful.
	//
	// 	- Failed: The task failed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the task. You can search for tasks by tag.
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// c2b277b9-0d30-4882-ad6d-ad661382****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The parameter definition in the JSON string format. For more information, see the Request parameters section of the topic about an asynchronous processing task.
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
	// The type of the task.
	//
	// example:
	//
	// VideoLabelClassification
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The custom user data.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s TaskInfo) String() string {
	return dara.Prettify(s)
}

func (s TaskInfo) GoString() string {
	return s.String()
}

func (s *TaskInfo) GetCode() *string {
	return s.Code
}

func (s *TaskInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *TaskInfo) GetMessage() *string {
	return s.Message
}

func (s *TaskInfo) GetProgress() *int32 {
	return s.Progress
}

func (s *TaskInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *TaskInfo) GetStatus() *string {
	return s.Status
}

func (s *TaskInfo) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *TaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *TaskInfo) GetTaskRequestDefinition() *string {
	return s.TaskRequestDefinition
}

func (s *TaskInfo) GetTaskType() *string {
	return s.TaskType
}

func (s *TaskInfo) GetUserData() *string {
	return s.UserData
}

func (s *TaskInfo) SetCode(v string) *TaskInfo {
	s.Code = &v
	return s
}

func (s *TaskInfo) SetEndTime(v string) *TaskInfo {
	s.EndTime = &v
	return s
}

func (s *TaskInfo) SetMessage(v string) *TaskInfo {
	s.Message = &v
	return s
}

func (s *TaskInfo) SetProgress(v int32) *TaskInfo {
	s.Progress = &v
	return s
}

func (s *TaskInfo) SetStartTime(v string) *TaskInfo {
	s.StartTime = &v
	return s
}

func (s *TaskInfo) SetStatus(v string) *TaskInfo {
	s.Status = &v
	return s
}

func (s *TaskInfo) SetTags(v map[string]interface{}) *TaskInfo {
	s.Tags = v
	return s
}

func (s *TaskInfo) SetTaskId(v string) *TaskInfo {
	s.TaskId = &v
	return s
}

func (s *TaskInfo) SetTaskRequestDefinition(v string) *TaskInfo {
	s.TaskRequestDefinition = &v
	return s
}

func (s *TaskInfo) SetTaskType(v string) *TaskInfo {
	s.TaskType = &v
	return s
}

func (s *TaskInfo) SetUserData(v string) *TaskInfo {
	s.UserData = &v
	return s
}

func (s *TaskInfo) Validate() error {
	return dara.Validate(s)
}
