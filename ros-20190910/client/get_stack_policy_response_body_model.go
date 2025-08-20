// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackPolicyResponseBody
	GetRequestId() *string
	SetStackPolicyBody(v map[string]interface{}) *GetStackPolicyResponseBody
	GetStackPolicyBody() map[string]interface{}
}

type GetStackPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
	//
	// example:
	//
	// {"Statement": [{"Action": "Update:*", "Effect": "Allow","Principal": "*","Resource": "*"}]}
	StackPolicyBody map[string]interface{} `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
}

func (s GetStackPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackPolicyResponseBody) GetStackPolicyBody() map[string]interface{} {
	return s.StackPolicyBody
}

func (s *GetStackPolicyResponseBody) SetRequestId(v string) *GetStackPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackPolicyResponseBody) SetStackPolicyBody(v map[string]interface{}) *GetStackPolicyResponseBody {
	s.StackPolicyBody = v
	return s
}

func (s *GetStackPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
