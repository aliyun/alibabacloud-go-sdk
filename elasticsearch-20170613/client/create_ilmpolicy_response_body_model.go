// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateILMPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateILMPolicyResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateILMPolicyResponseBody
	GetResult() *string
}

type CreateILMPolicyResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// my-policy
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateILMPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateILMPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateILMPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateILMPolicyResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateILMPolicyResponseBody) SetRequestId(v string) *CreateILMPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateILMPolicyResponseBody) SetResult(v string) *CreateILMPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *CreateILMPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
