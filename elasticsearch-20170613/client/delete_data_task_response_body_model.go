// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteDataTaskResponseBody
	GetResult() *bool
}

type DeleteDataTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteDataTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteDataTaskResponseBody) SetRequestId(v string) *DeleteDataTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataTaskResponseBody) SetResult(v bool) *DeleteDataTaskResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDataTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
