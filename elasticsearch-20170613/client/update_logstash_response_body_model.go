// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLogstashResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateLogstashResponseBody
	GetResult() *bool
}

type UpdateLogstashResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLogstashResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateLogstashResponseBody) SetRequestId(v string) *UpdateLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogstashResponseBody) SetResult(v bool) *UpdateLogstashResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateLogstashResponseBody) Validate() error {
	return dara.Validate(s)
}
