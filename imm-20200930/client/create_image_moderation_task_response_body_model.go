// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageModerationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateImageModerationTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateImageModerationTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateImageModerationTaskResponseBody
	GetTaskId() *string
}

type CreateImageModerationTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 2E6-1I0FGn0zFnl5AflRfhzClma*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B3D5E0A-D8B8-4DA0-8127-ED32C851****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ImageModeration-179ef4f8-d583-4f0c-a293-7c0889c*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageModerationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageModerationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageModerationTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateImageModerationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageModerationTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageModerationTaskResponseBody) SetEventId(v string) *CreateImageModerationTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateImageModerationTaskResponseBody) SetRequestId(v string) *CreateImageModerationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageModerationTaskResponseBody) SetTaskId(v string) *CreateImageModerationTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageModerationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
