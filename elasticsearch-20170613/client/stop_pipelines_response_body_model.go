// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopPipelinesResponseBody
	GetRequestId() *string
	SetResult(v bool) *StopPipelinesResponseBody
	GetResult() *bool
}

type StopPipelinesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// 	- true: The pipelines are stopped.
	//
	// 	- false: The pipelines fail to be stopped.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StopPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPipelinesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StopPipelinesResponseBody) SetRequestId(v string) *StopPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelinesResponseBody) SetResult(v bool) *StopPipelinesResponseBody {
	s.Result = &v
	return s
}

func (s *StopPipelinesResponseBody) Validate() error {
	return dara.Validate(s)
}
