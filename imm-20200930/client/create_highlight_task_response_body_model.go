// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighlightTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateHighlightTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateHighlightTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateHighlightTaskResponseBody
	GetTaskId() *string
}

type CreateHighlightTaskResponseBody struct {
	// example:
	//
	// 0ED-1Bz8z71k5TtsUejT4UJ16Es****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFFE0B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Highlight-4d51241b-04d4-4343-aa25-****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateHighlightTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateHighlightTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHighlightTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateHighlightTaskResponseBody) SetEventId(v string) *CreateHighlightTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateHighlightTaskResponseBody) SetRequestId(v string) *CreateHighlightTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHighlightTaskResponseBody) SetTaskId(v string) *CreateHighlightTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateHighlightTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
