// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateConnectionResponseBody
	GetRequestId() *string
	SetResult(v bool) *ValidateConnectionResponseBody
	GetResult() *bool
}

type ValidateConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the connectivity is normal. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateConnectionResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ValidateConnectionResponseBody) SetRequestId(v string) *ValidateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateConnectionResponseBody) SetResult(v bool) *ValidateConnectionResponseBody {
	s.Result = &v
	return s
}

func (s *ValidateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
