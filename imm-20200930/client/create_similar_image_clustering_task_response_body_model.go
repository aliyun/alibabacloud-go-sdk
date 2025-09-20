// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusteringTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateSimilarImageClusteringTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateSimilarImageClusteringTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateSimilarImageClusteringTaskResponseBody
	GetTaskId() *string
}

type CreateSimilarImageClusteringTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 3BF-1UhtFyrua71eOkFlqYq23Co****
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
	// SimilarImageClustering-48d0a0f3-8459-47f4-b8af-ff49c64****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSimilarImageClusteringTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusteringTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusteringTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateSimilarImageClusteringTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSimilarImageClusteringTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSimilarImageClusteringTaskResponseBody) SetEventId(v string) *CreateSimilarImageClusteringTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponseBody) SetRequestId(v string) *CreateSimilarImageClusteringTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponseBody) SetTaskId(v string) *CreateSimilarImageClusteringTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateSimilarImageClusteringTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
