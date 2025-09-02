// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDDLJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDDLJobStatusResponseBodyData) *GetDDLJobStatusResponseBody
	GetData() *GetDDLJobStatusResponseBodyData
	SetRequestId(v string) *GetDDLJobStatusResponseBody
	GetRequestId() *string
}

type GetDDLJobStatusResponseBody struct {
	// The details of the task.
	Data *GetDDLJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// abc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDDLJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDDLJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDDLJobStatusResponseBody) GetData() *GetDDLJobStatusResponseBodyData {
	return s.Data
}

func (s *GetDDLJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDDLJobStatusResponseBody) SetData(v *GetDDLJobStatusResponseBodyData) *GetDDLJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetDDLJobStatusResponseBody) SetRequestId(v string) *GetDDLJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDDLJobStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDDLJobStatusResponseBodyData struct {
	// The content of the task.
	//
	// example:
	//
	// success
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the ongoing task. If no value is returned for this parameter, all subtasks are complete.
	//
	// example:
	//
	// abc1
	NextTaskId *string `json:"NextTaskId,omitempty" xml:"NextTaskId,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// abc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDDLJobStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDDLJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDDLJobStatusResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetDDLJobStatusResponseBodyData) GetNextTaskId() *string {
	return s.NextTaskId
}

func (s *GetDDLJobStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDDLJobStatusResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDDLJobStatusResponseBodyData) SetContent(v string) *GetDDLJobStatusResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetDDLJobStatusResponseBodyData) SetNextTaskId(v string) *GetDDLJobStatusResponseBodyData {
	s.NextTaskId = &v
	return s
}

func (s *GetDDLJobStatusResponseBodyData) SetStatus(v string) *GetDDLJobStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDDLJobStatusResponseBodyData) SetTaskId(v string) *GetDDLJobStatusResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDDLJobStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
