// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunPipelinesResponseBody
	GetRequestId() *string
	SetResult(v bool) *RunPipelinesResponseBody
	GetResult() *bool
}

type RunPipelinesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// 	- true: successful
	//
	// 	- false: failed
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RunPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *RunPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunPipelinesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RunPipelinesResponseBody) SetRequestId(v string) *RunPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunPipelinesResponseBody) SetResult(v bool) *RunPipelinesResponseBody {
	s.Result = &v
	return s
}

func (s *RunPipelinesResponseBody) Validate() error {
	return dara.Validate(s)
}
