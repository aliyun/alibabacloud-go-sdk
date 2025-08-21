// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePipelinesResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeletePipelinesResponseBody
	GetResult() *bool
}

type DeletePipelinesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the pipeline is deleted. Valid values:
	//
	// 	- true: The pipeline is deleted.
	//
	// 	- false: The pipeline fails to be deleted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeletePipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePipelinesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeletePipelinesResponseBody) SetRequestId(v string) *DeletePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelinesResponseBody) SetResult(v bool) *DeletePipelinesResponseBody {
	s.Result = &v
	return s
}

func (s *DeletePipelinesResponseBody) Validate() error {
	return dara.Validate(s)
}
