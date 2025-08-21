// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCollectorResponseBody
	GetRequestId() *string
	SetResult(v bool) *StopCollectorResponseBody
	GetResult() *bool
}

type StopCollectorResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StopCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *StopCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCollectorResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StopCollectorResponseBody) SetRequestId(v string) *StopCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCollectorResponseBody) SetResult(v bool) *StopCollectorResponseBody {
	s.Result = &v
	return s
}

func (s *StopCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
