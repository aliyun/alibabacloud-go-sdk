// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncEmailCaptchaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendAsyncEmailCaptchaResponseBody
	GetCode() *string
	SetData(v bool) *SendAsyncEmailCaptchaResponseBody
	GetData() *bool
	SetMessage(v string) *SendAsyncEmailCaptchaResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendAsyncEmailCaptchaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendAsyncEmailCaptchaResponseBody
	GetSuccess() *bool
}

type SendAsyncEmailCaptchaResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Instance invalid-id-123 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5E9636D3-6C10-5FAB-B391-EDD122E28BC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendAsyncEmailCaptchaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncEmailCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *SendAsyncEmailCaptchaResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendAsyncEmailCaptchaResponseBody) GetData() *bool {
	return s.Data
}

func (s *SendAsyncEmailCaptchaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendAsyncEmailCaptchaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendAsyncEmailCaptchaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendAsyncEmailCaptchaResponseBody) SetCode(v string) *SendAsyncEmailCaptchaResponseBody {
	s.Code = &v
	return s
}

func (s *SendAsyncEmailCaptchaResponseBody) SetData(v bool) *SendAsyncEmailCaptchaResponseBody {
	s.Data = &v
	return s
}

func (s *SendAsyncEmailCaptchaResponseBody) SetMessage(v string) *SendAsyncEmailCaptchaResponseBody {
	s.Message = &v
	return s
}

func (s *SendAsyncEmailCaptchaResponseBody) SetRequestId(v string) *SendAsyncEmailCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAsyncEmailCaptchaResponseBody) SetSuccess(v bool) *SendAsyncEmailCaptchaResponseBody {
	s.Success = &v
	return s
}

func (s *SendAsyncEmailCaptchaResponseBody) Validate() error {
	return dara.Validate(s)
}
