// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClustersMergingTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateFigureClustersMergingTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateFigureClustersMergingTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateFigureClustersMergingTaskResponseBody
	GetTaskId() *string
}

type CreateFigureClustersMergingTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 0ED-1Bz8z71k5TtsUejT4UJ16E****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 92376fbb-171f-4259-913f-705f7ee0****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFigureClustersMergingTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClustersMergingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFigureClustersMergingTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateFigureClustersMergingTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFigureClustersMergingTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateFigureClustersMergingTaskResponseBody) SetEventId(v string) *CreateFigureClustersMergingTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponseBody) SetRequestId(v string) *CreateFigureClustersMergingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponseBody) SetTaskId(v string) *CreateFigureClustersMergingTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateFigureClustersMergingTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
