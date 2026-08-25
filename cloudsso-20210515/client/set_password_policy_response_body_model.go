// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetPasswordPolicyResponseBody
	GetRequestId() *string
}

type SetPasswordPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 768F908D-A66A-5A5D-816C-20C93CBBFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPasswordPolicyResponseBody) SetRequestId(v string) *SetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPasswordPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
