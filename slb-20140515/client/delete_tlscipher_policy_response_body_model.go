// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTLSCipherPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTLSCipherPolicyResponseBody
	GetRequestId() *string
}

type DeleteTLSCipherPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTLSCipherPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTLSCipherPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTLSCipherPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTLSCipherPolicyResponseBody) SetRequestId(v string) *DeleteTLSCipherPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTLSCipherPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
