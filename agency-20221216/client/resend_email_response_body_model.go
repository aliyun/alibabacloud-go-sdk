// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResendEmailResponseBody
	GetCode() *string
	SetMessage(v string) *ResendEmailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResendEmailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResendEmailResponseBody
	GetSuccess() *bool
}

type ResendEmailResponseBody struct {
	// The result code.
	//
	// Error code. Valid values:
	//
	// - 200: OK.
	//
	// - 1109: System error.
	//
	// - 3058: Emails are sent too frequently. A maximum of 10 emails can be sent within 5 minutes.
	//
	// - 3057: InviteId is empty.
	//
	// - 3060: No sending record found for the specified InviteId.
	//
	// - 3061: The registration URL has expired and the email cannot be resent.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result message.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. A unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Valid values: True/False. Indicates whether the current API call itself is successful. This does not indicate the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResendEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailResponseBody) GoString() string {
	return s.String()
}

func (s *ResendEmailResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResendEmailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResendEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResendEmailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResendEmailResponseBody) SetCode(v string) *ResendEmailResponseBody {
	s.Code = &v
	return s
}

func (s *ResendEmailResponseBody) SetMessage(v string) *ResendEmailResponseBody {
	s.Message = &v
	return s
}

func (s *ResendEmailResponseBody) SetRequestId(v string) *ResendEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResendEmailResponseBody) SetSuccess(v bool) *ResendEmailResponseBody {
	s.Success = &v
	return s
}

func (s *ResendEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
