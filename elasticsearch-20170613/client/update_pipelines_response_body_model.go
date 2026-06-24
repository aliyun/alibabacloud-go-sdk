// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePipelinesResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdatePipelinesResponseBody
	GetResult() *bool
}

type UpdatePipelinesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The update was successful.
	//
	// - false: The update failed.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdatePipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelinesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdatePipelinesResponseBody) SetRequestId(v string) *UpdatePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelinesResponseBody) SetResult(v bool) *UpdatePipelinesResponseBody {
	s.Result = &v
	return s
}

func (s *UpdatePipelinesResponseBody) Validate() error {
	return dara.Validate(s)
}
