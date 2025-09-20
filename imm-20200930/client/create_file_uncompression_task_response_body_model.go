// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUncompressionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateFileUncompressionTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateFileUncompressionTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateFileUncompressionTaskResponseBody
	GetTaskId() *string
}

type CreateFileUncompressionTaskResponseBody struct {
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
	// FileUncompression-16ab5dd6-af02-480e-9ed7-a8d51b1*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFileUncompressionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUncompressionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileUncompressionTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateFileUncompressionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileUncompressionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateFileUncompressionTaskResponseBody) SetEventId(v string) *CreateFileUncompressionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFileUncompressionTaskResponseBody) SetRequestId(v string) *CreateFileUncompressionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileUncompressionTaskResponseBody) SetTaskId(v string) *CreateFileUncompressionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateFileUncompressionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
