// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileCompressionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateFileCompressionTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateFileCompressionTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateFileCompressionTaskResponseBody
	GetTaskId() *string
}

type CreateFileCompressionTaskResponseBody struct {
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
	// FileCompression-3579a380-6f7a-4a9d-b9d2-65996*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFileCompressionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileCompressionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileCompressionTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateFileCompressionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileCompressionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateFileCompressionTaskResponseBody) SetEventId(v string) *CreateFileCompressionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFileCompressionTaskResponseBody) SetRequestId(v string) *CreateFileCompressionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileCompressionTaskResponseBody) SetTaskId(v string) *CreateFileCompressionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateFileCompressionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
