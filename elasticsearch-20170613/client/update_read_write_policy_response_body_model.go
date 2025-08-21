// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReadWritePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateReadWritePolicyResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateReadWritePolicyResponseBody
	GetResult() *bool
}

type UpdateReadWritePolicyResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateReadWritePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadWritePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReadWritePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateReadWritePolicyResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateReadWritePolicyResponseBody) SetRequestId(v string) *UpdateReadWritePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateReadWritePolicyResponseBody) SetResult(v bool) *UpdateReadWritePolicyResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateReadWritePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
