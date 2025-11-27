// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoLabelClassificationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoLabelClassificationResultResponseBody
	GetCode() *string
	SetEndTime(v string) *GetVideoLabelClassificationResultResponseBody
	GetEndTime() *string
	SetEventId(v string) *GetVideoLabelClassificationResultResponseBody
	GetEventId() *string
	SetLabels(v []*Label) *GetVideoLabelClassificationResultResponseBody
	GetLabels() []*Label
	SetMessage(v string) *GetVideoLabelClassificationResultResponseBody
	GetMessage() *string
	SetProjectName(v string) *GetVideoLabelClassificationResultResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetVideoLabelClassificationResultResponseBody
	GetRequestId() *string
	SetStartTime(v string) *GetVideoLabelClassificationResultResponseBody
	GetStartTime() *string
	SetStatus(v string) *GetVideoLabelClassificationResultResponseBody
	GetStatus() *string
	SetTaskId(v string) *GetVideoLabelClassificationResultResponseBody
	GetTaskId() *string
	SetTaskType(v string) *GetVideoLabelClassificationResultResponseBody
	GetTaskType() *string
	SetUserData(v string) *GetVideoLabelClassificationResultResponseBody
	GetUserData() *string
}

type GetVideoLabelClassificationResultResponseBody struct {
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
	// 2021-12-24T03:00:42.134971294Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 2F6-1Bz99Xi93EnRpNEyLudILJm****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
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
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7F84C6D9-5AC0-49F9-914D-F02678E3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the task.
	//
	// example:
	//
	// 2021-12-24T03:00:38.892462383Z
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
	// VideoLabelClassification-2f157087-91df-4fda-8c3e-232407ec****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// VideoLabelClassification
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The custom information.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetVideoLabelClassificationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoLabelClassificationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoLabelClassificationResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoLabelClassificationResultResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetVideoLabelClassificationResultResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *GetVideoLabelClassificationResultResponseBody) GetLabels() []*Label {
	return s.Labels
}

func (s *GetVideoLabelClassificationResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoLabelClassificationResultResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetVideoLabelClassificationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoLabelClassificationResultResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetVideoLabelClassificationResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVideoLabelClassificationResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoLabelClassificationResultResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVideoLabelClassificationResultResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *GetVideoLabelClassificationResultResponseBody) SetCode(v string) *GetVideoLabelClassificationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetEndTime(v string) *GetVideoLabelClassificationResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetEventId(v string) *GetVideoLabelClassificationResultResponseBody {
	s.EventId = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetLabels(v []*Label) *GetVideoLabelClassificationResultResponseBody {
	s.Labels = v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetMessage(v string) *GetVideoLabelClassificationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetProjectName(v string) *GetVideoLabelClassificationResultResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetRequestId(v string) *GetVideoLabelClassificationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetStartTime(v string) *GetVideoLabelClassificationResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetStatus(v string) *GetVideoLabelClassificationResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetTaskId(v string) *GetVideoLabelClassificationResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetTaskType(v string) *GetVideoLabelClassificationResultResponseBody {
	s.TaskType = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) SetUserData(v string) *GetVideoLabelClassificationResultResponseBody {
	s.UserData = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponseBody) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
