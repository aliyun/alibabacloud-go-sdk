// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityPolicyResponseBody
	GetRequestId() *string
}

type DeleteSecurityPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityPolicyResponseBody) SetRequestId(v string) *DeleteSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
