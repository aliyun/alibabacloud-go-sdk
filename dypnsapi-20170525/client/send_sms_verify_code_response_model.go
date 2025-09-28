// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSmsVerifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendSmsVerifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendSmsVerifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *SendSmsVerifyCodeResponseBody) *SendSmsVerifyCodeResponse
	GetBody() *SendSmsVerifyCodeResponseBody
}

type SendSmsVerifyCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSmsVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSmsVerifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SendSmsVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendSmsVerifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendSmsVerifyCodeResponse) GetBody() *SendSmsVerifyCodeResponseBody {
	return s.Body
}

func (s *SendSmsVerifyCodeResponse) SetHeaders(v map[string]*string) *SendSmsVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *SendSmsVerifyCodeResponse) SetStatusCode(v int32) *SendSmsVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSmsVerifyCodeResponse) SetBody(v *SendSmsVerifyCodeResponseBody) *SendSmsVerifyCodeResponse {
	s.Body = v
	return s
}

func (s *SendSmsVerifyCodeResponse) Validate() error {
	return dara.Validate(s)
}
