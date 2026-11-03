// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCollectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCollectorResponseBody
	GetRequestId() *string
	SetResult(v bool) *StartCollectorResponseBody
	GetResult() *bool
}

type StartCollectorResponseBody struct {
	// The ID of the request.
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

func (s StartCollectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *StartCollectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCollectorResponseBody) GetResult() *bool {
	return s.Result
}

func (s *StartCollectorResponseBody) SetRequestId(v string) *StartCollectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCollectorResponseBody) SetResult(v bool) *StartCollectorResponseBody {
	s.Result = &v
	return s
}

func (s *StartCollectorResponseBody) Validate() error {
	return dara.Validate(s)
}
