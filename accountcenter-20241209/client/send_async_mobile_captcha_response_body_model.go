// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncMobileCaptchaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SendAsyncMobileCaptchaResponseBody
	GetCode() *string
	SetData(v bool) *SendAsyncMobileCaptchaResponseBody
	GetData() *bool
	SetMessage(v string) *SendAsyncMobileCaptchaResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendAsyncMobileCaptchaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendAsyncMobileCaptchaResponseBody
	GetSuccess() *bool
}

type SendAsyncMobileCaptchaResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// False
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1409E8EE-8F9A-506B-BACB-B9DF3634C287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendAsyncMobileCaptchaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncMobileCaptchaResponseBody) GoString() string {
	return s.String()
}

func (s *SendAsyncMobileCaptchaResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendAsyncMobileCaptchaResponseBody) GetData() *bool {
	return s.Data
}

func (s *SendAsyncMobileCaptchaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendAsyncMobileCaptchaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendAsyncMobileCaptchaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendAsyncMobileCaptchaResponseBody) SetCode(v string) *SendAsyncMobileCaptchaResponseBody {
	s.Code = &v
	return s
}

func (s *SendAsyncMobileCaptchaResponseBody) SetData(v bool) *SendAsyncMobileCaptchaResponseBody {
	s.Data = &v
	return s
}

func (s *SendAsyncMobileCaptchaResponseBody) SetMessage(v string) *SendAsyncMobileCaptchaResponseBody {
	s.Message = &v
	return s
}

func (s *SendAsyncMobileCaptchaResponseBody) SetRequestId(v string) *SendAsyncMobileCaptchaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAsyncMobileCaptchaResponseBody) SetSuccess(v bool) *SendAsyncMobileCaptchaResponseBody {
	s.Success = &v
	return s
}

func (s *SendAsyncMobileCaptchaResponseBody) Validate() error {
	return dara.Validate(s)
}
