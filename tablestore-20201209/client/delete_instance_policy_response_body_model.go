// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstancePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteInstancePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstancePolicyResponseBody
	GetRequestId() *string
}

type DeleteInstancePolicyResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 3104C83E-6E82-57FB-BB88-8C64CCFDEF89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstancePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstancePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstancePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstancePolicyResponseBody) SetCode(v string) *DeleteInstancePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstancePolicyResponseBody) SetMessage(v string) *DeleteInstancePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstancePolicyResponseBody) SetRequestId(v string) *DeleteInstancePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstancePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
