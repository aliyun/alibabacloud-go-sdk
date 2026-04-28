// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimilarImageClusterTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CreateSimilarImageClusterTaskResponseBody
	GetTaskId() *string
}

type CreateSimilarImageClusterTaskResponseBody struct {
	// example:
	//
	// i:SimilarImageClustering-b67d53e7-2fe8-460f-9b95-1e93636923eb
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateSimilarImageClusterTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimilarImageClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusterTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSimilarImageClusterTaskResponseBody) SetTaskId(v string) *CreateSimilarImageClusterTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateSimilarImageClusterTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
