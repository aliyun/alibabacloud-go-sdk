// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFigureClusteringTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateFigureClusteringTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateFigureClusteringTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateFigureClusteringTaskResponseBody
	GetTaskId() *string
}

type CreateFigureClusteringTaskResponseBody struct {
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
	// 1B3D5E0A-D8B8-4DA0-8127-ED32C851****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// formatconvert-00bec802-073a-4b61-ba3b-39bc****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateFigureClusteringTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFigureClusteringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFigureClusteringTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateFigureClusteringTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFigureClusteringTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateFigureClusteringTaskResponseBody) SetEventId(v string) *CreateFigureClusteringTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateFigureClusteringTaskResponseBody) SetRequestId(v string) *CreateFigureClusteringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFigureClusteringTaskResponseBody) SetTaskId(v string) *CreateFigureClusteringTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateFigureClusteringTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
