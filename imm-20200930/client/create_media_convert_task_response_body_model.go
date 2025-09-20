// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConvertTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateMediaConvertTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateMediaConvertTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateMediaConvertTaskResponseBody
	GetTaskId() *string
}

type CreateMediaConvertTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 0ED-1Bz8z71k5TtsUejT4UJ16Es****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFFE0B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// MediaConvert-adb1ee28-c4c9-42a7-9f54-3b8eadcb****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMediaConvertTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateMediaConvertTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaConvertTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMediaConvertTaskResponseBody) SetEventId(v string) *CreateMediaConvertTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateMediaConvertTaskResponseBody) SetRequestId(v string) *CreateMediaConvertTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaConvertTaskResponseBody) SetTaskId(v string) *CreateMediaConvertTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateMediaConvertTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
