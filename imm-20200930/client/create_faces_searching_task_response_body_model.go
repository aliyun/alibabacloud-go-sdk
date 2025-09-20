// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFacesSearchingTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateFacesSearchingTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateFacesSearchingTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateFacesSearchingTaskResponseBody
	GetTaskId() *string
}

type CreateFacesSearchingTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 0ED-1Bz8z71k5TtsUejT4UJ16****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1E79399-05F7-06D8-95FE-EBE17BA*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// CreateFacesSearchingTask-00bec802-073a-4b61-ba*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFacesSearchingTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFacesSearchingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFacesSearchingTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateFacesSearchingTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFacesSearchingTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateFacesSearchingTaskResponseBody) SetEventId(v string) *CreateFacesSearchingTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFacesSearchingTaskResponseBody) SetRequestId(v string) *CreateFacesSearchingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFacesSearchingTaskResponseBody) SetTaskId(v string) *CreateFacesSearchingTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateFacesSearchingTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
