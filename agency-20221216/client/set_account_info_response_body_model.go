// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetAccountInfoResponseBody
	GetCode() *string
	SetMessage(v string) *SetAccountInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetAccountInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetAccountInfoResponseBody
	GetSuccess() *bool
}

type SetAccountInfoResponseBody struct {
	// The result code. Valid values:
	//
	// - 200: OK.
	//
	// - 1109: A system error occurred.
	//
	// - 3030: The account nickname exceeds the maximum length of 150 bytes.
	//
	// - 3031: The remark exceeds the maximum length of 3000 bytes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Alibaba Cloud uses this ID for error tracking.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Valid values: True and False. Indicates whether the current API call itself is successful. This does not indicate the success of subsequent business operations.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAccountInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccountInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetAccountInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetAccountInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccountInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetAccountInfoResponseBody) SetCode(v string) *SetAccountInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetMessage(v string) *SetAccountInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetRequestId(v string) *SetAccountInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccountInfoResponseBody) SetSuccess(v bool) *SetAccountInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SetAccountInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
