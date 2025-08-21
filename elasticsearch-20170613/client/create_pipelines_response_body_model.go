// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePipelinesResponseBody
	GetRequestId() *string
	SetResult(v bool) *CreatePipelinesResponseBody
	GetResult() *bool
}

type CreatePipelinesResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreatePipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelinesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CreatePipelinesResponseBody) SetRequestId(v string) *CreatePipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelinesResponseBody) SetResult(v bool) *CreatePipelinesResponseBody {
	s.Result = &v
	return s
}

func (s *CreatePipelinesResponseBody) Validate() error {
	return dara.Validate(s)
}
