// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeductOutstandingBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeductOutstandingBalanceResponseBody
	GetCode() *string
	SetMessage(v string) *DeductOutstandingBalanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeductOutstandingBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeductOutstandingBalanceResponseBody
	GetSuccess() *bool
}

type DeductOutstandingBalanceResponseBody struct {
	// The result code. Valid values:
	//
	// 200: OK.
	//
	// 1109: System error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message. The value is the same as the Code parameter value.
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
	// Indicates whether the current API call itself is successful. Valid values: True and False. This does not indicate the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeductOutstandingBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeductOutstandingBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeductOutstandingBalanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeductOutstandingBalanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeductOutstandingBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeductOutstandingBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeductOutstandingBalanceResponseBody) SetCode(v string) *DeductOutstandingBalanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetMessage(v string) *DeductOutstandingBalanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetRequestId(v string) *DeductOutstandingBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) SetSuccess(v bool) *DeductOutstandingBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeductOutstandingBalanceResponseBody) Validate() error {
	return dara.Validate(s)
}
