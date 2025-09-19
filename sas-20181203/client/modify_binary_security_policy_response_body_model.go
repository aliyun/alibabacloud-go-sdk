// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBinarySecurityPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBinarySecurityPolicyResponseBody
	GetRequestId() *string
}

type ModifyBinarySecurityPolicyResponseBody struct {
	// The request ID of the result.
	//
	// example:
	//
	// C699E4E4-F2F4-58FC-A949-457FFE59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBinarySecurityPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBinarySecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBinarySecurityPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBinarySecurityPolicyResponseBody) SetRequestId(v string) *ModifyBinarySecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBinarySecurityPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
