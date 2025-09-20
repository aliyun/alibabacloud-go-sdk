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
	Code                  *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	EndTime               *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Message               *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Progress              *int32                 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartTime             *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status                *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                  map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TaskId                *string                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskRequestDefinition *string                `json:"TaskRequestDefinition,omitempty" xml:"TaskRequestDefinition,omitempty"`
	TaskType              *string                `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UserData              *string                `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
