// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLocationDateClusteringTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateLocationDateClusteringTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateLocationDateClusteringTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateLocationDateClusteringTaskResponseBody
	GetTaskId() *string
}

type CreateLocationDateClusteringTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 25B-1W2ChgujA3Q8MbBY6mSp2mh****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B121940C-9794-4EE3-8D6E-F8EC525F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// LocationDateClustering-c10dce07-1de7-4da7-abee-1a3aba7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateLocationDateClusteringTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLocationDateClusteringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLocationDateClusteringTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateLocationDateClusteringTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLocationDateClusteringTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateLocationDateClusteringTaskResponseBody) SetEventId(v string) *CreateLocationDateClusteringTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponseBody) SetRequestId(v string) *CreateLocationDateClusteringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponseBody) SetTaskId(v string) *CreateLocationDateClusteringTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateLocationDateClusteringTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
