// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConditionalAccessPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConditionalAccessPolicyResponseBody
	GetRequestId() *string
}

type DeleteConditionalAccessPolicyResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConditionalAccessPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConditionalAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConditionalAccessPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConditionalAccessPolicyResponseBody) SetRequestId(v string) *DeleteConditionalAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConditionalAccessPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
