// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBinarySecurityPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteBinarySecurityPolicyResponseBody
	GetRequestId() *string
}

type DeleteBinarySecurityPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A6CFADC0-1167-521A-9284-8CD8034C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBinarySecurityPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBinarySecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBinarySecurityPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBinarySecurityPolicyResponseBody) SetRequestId(v string) *DeleteBinarySecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBinarySecurityPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
