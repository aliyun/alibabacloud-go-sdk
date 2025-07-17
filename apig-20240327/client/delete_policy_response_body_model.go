// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePolicyResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePolicyResponseBody
	GetRequestId() *string
}

type DeletePolicyResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E7406754***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyResponseBody) SetCode(v string) *DeletePolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyResponseBody) SetMessage(v string) *DeletePolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
