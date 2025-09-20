// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoModerationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateVideoModerationTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateVideoModerationTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateVideoModerationTaskResponseBody
	GetTaskId() *string
}

type CreateVideoModerationTaskResponseBody struct {
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
	// VideoModeration-9442a216-4691-4a48-846d-76daccaf*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoModerationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoModerationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoModerationTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateVideoModerationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoModerationTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoModerationTaskResponseBody) SetEventId(v string) *CreateVideoModerationTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateVideoModerationTaskResponseBody) SetRequestId(v string) *CreateVideoModerationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoModerationTaskResponseBody) SetTaskId(v string) *CreateVideoModerationTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVideoModerationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
