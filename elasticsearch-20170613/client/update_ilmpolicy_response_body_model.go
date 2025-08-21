// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateILMPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateILMPolicyResponseBody
	GetRequestId() *string
	SetResult(v string) *UpdateILMPolicyResponseBody
	GetResult() *string
}

type UpdateILMPolicyResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// my-policy
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateILMPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateILMPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateILMPolicyResponseBody) GetResult() *string {
	return s.Result
}

func (s *UpdateILMPolicyResponseBody) SetRequestId(v string) *UpdateILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateILMPolicyResponseBody) SetResult(v string) *UpdateILMPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateILMPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
