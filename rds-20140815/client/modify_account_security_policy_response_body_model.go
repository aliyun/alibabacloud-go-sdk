// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountSecurityPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAccountSecurityPolicyResponseBody
	GetRequestId() *string
}

type ModifyAccountSecurityPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F2911788-25E8-42E5-A3A3-1B38D263F01E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountSecurityPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountSecurityPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountSecurityPolicyResponseBody) SetRequestId(v string) *ModifyAccountSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountSecurityPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
