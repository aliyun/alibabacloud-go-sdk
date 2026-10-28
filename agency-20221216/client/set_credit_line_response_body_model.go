// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCreditLineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetCreditLineResponseBody
	GetCode() *string
	SetMessage(v string) *SetCreditLineResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetCreditLineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetCreditLineResponseBody
	GetSuccess() *bool
}

type SetCreditLineResponseBody struct {
	// The result code. Valid values:
	//
	// - 200: OK.
	//
	// - 1109: System error.
	//
	// - 3040: The sub-account is frozen and the operation cannot be completed.
	//
	// - 3041: The credit limit is not a number.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that corresponds to the result code.
	//
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. A unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call is successful. Valid values: True and False. This does not indicate whether the subsequent business operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetCreditLineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCreditLineResponseBody) GoString() string {
	return s.String()
}

func (s *SetCreditLineResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetCreditLineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetCreditLineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCreditLineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetCreditLineResponseBody) SetCode(v string) *SetCreditLineResponseBody {
	s.Code = &v
	return s
}

func (s *SetCreditLineResponseBody) SetMessage(v string) *SetCreditLineResponseBody {
	s.Message = &v
	return s
}

func (s *SetCreditLineResponseBody) SetRequestId(v string) *SetCreditLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCreditLineResponseBody) SetSuccess(v bool) *SetCreditLineResponseBody {
	s.Success = &v
	return s
}

func (s *SetCreditLineResponseBody) Validate() error {
	return dara.Validate(s)
}
