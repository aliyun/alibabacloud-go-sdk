// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArchiveFileInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateArchiveFileInspectionTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateArchiveFileInspectionTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateArchiveFileInspectionTaskResponseBody
	GetTaskId() *string
}

type CreateArchiveFileInspectionTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 0ED-1Bz8z71k5TtsUejT4UJ16Es*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC564A9A-BA5C-4499-A087-D9B9E76E*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ArchiveFileInspection-8475218e-d86e-4c66-b3cf-50e74d6c****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateArchiveFileInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArchiveFileInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArchiveFileInspectionTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateArchiveFileInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArchiveFileInspectionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateArchiveFileInspectionTaskResponseBody) SetEventId(v string) *CreateArchiveFileInspectionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponseBody) SetRequestId(v string) *CreateArchiveFileInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponseBody) SetTaskId(v string) *CreateArchiveFileInspectionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateArchiveFileInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
