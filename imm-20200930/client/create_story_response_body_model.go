// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateStoryResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateStoryResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateStoryResponseBody
	GetTaskId() *string
}

type CreateStoryResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 392-1CqzvESGTEeNZ2OWFbRKIM****
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
	// CreateStory-4ef6ff43-edf3-4612-9cc4-0c7f9e19****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoryResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateStoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStoryResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateStoryResponseBody) SetEventId(v string) *CreateStoryResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateStoryResponseBody) SetRequestId(v string) *CreateStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStoryResponseBody) SetTaskId(v string) *CreateStoryResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
