// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTaskResponseBody
	GetCode() *string
	SetData(v *CreateTaskResponseBodyData) *CreateTaskResponseBody
	GetData() *CreateTaskResponseBodyData
	SetMessage(v string) *CreateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
}

type CreateTaskResponseBody struct {
	// example:
	//
	// 0
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTaskResponseBody) GetData() *CreateTaskResponseBodyData {
	return s.Data
}

func (s *CreateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) SetCode(v string) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetData(v *CreateTaskResponseBodyData) *CreateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTaskResponseBodyData struct {
	MeetingJoinUrl *string `json:"MeetingJoinUrl,omitempty" xml:"MeetingJoinUrl,omitempty"`
	// example:
	//
	// c5394c6ee0fb474899d42215a3925c7e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// task_tingwu_123
	TaskKey    *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s CreateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyData) GetMeetingJoinUrl() *string {
	return s.MeetingJoinUrl
}

func (s *CreateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskResponseBodyData) GetTaskKey() *string {
	return s.TaskKey
}

func (s *CreateTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *CreateTaskResponseBodyData) SetMeetingJoinUrl(v string) *CreateTaskResponseBodyData {
	s.MeetingJoinUrl = &v
	return s
}

func (s *CreateTaskResponseBodyData) SetTaskId(v string) *CreateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBodyData) SetTaskKey(v string) *CreateTaskResponseBodyData {
	s.TaskKey = &v
	return s
}

func (s *CreateTaskResponseBodyData) SetTaskStatus(v string) *CreateTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *CreateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
