// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDecodeBlindWatermarkResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetCode() *string
	SetContent(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetContent() *string
	SetEndTime(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetEndTime() *string
	SetEventId(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetEventId() *string
	SetMessage(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetMessage() *string
	SetProjectName(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetTaskId() *string
	SetTaskType(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetTaskType() *string
	SetUserData(v string) *GetDecodeBlindWatermarkResultResponseBody
	GetUserData() *string
}

type GetDecodeBlindWatermarkResultResponseBody struct {
	// The error code of the task.
	//
	// example:
	//
	// ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The watermark content.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The end time of the task.
	//
	// example:
	//
	// 2024-03-03T09:45:56.87Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 2C2-1I0EG57VR37J4rQ8oKG6C9*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The error message of the task.
	//
	// example:
	//
	// The specified resource project is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The project name.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 93126E40-0296-4129-95E3-AFAC709372E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2024-03-03T09:44:31.029Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// DecodeBlindWatermark-c09b0943-ed79-4983-8dbe-7a882574****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// example:
	//
	// DecodeBlindWatermark
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The user data of the task.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetDecodeBlindWatermarkResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDecodeBlindWatermarkResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetDecodeBlindWatermarkResultResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetCode(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetContent(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.Content = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetEndTime(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetEventId(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetMessage(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetProjectName(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetRequestId(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetStartTime(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetStatus(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetTaskId(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetTaskType(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) SetUserData(v string) *GetDecodeBlindWatermarkResultResponseBody {
	s.UserData = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultResponseBody) Validate() error {
	return dara.Validate(s)
}
