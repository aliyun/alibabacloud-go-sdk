// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncEmailCaptchaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAsyncEmailCaptchaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAsyncEmailCaptchaResponse
	GetStatusCode() *int32
	SetBody(v *SendAsyncEmailCaptchaResponseBody) *SendAsyncEmailCaptchaResponse
	GetBody() *SendAsyncEmailCaptchaResponseBody
}

type SendAsyncEmailCaptchaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAsyncEmailCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAsyncEmailCaptchaResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncEmailCaptchaResponse) GoString() string {
	return s.String()
}

func (s *SendAsyncEmailCaptchaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAsyncEmailCaptchaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAsyncEmailCaptchaResponse) GetBody() *SendAsyncEmailCaptchaResponseBody {
	return s.Body
}

func (s *SendAsyncEmailCaptchaResponse) SetHeaders(v map[string]*string) *SendAsyncEmailCaptchaResponse {
	s.Headers = v
	return s
}

func (s *SendAsyncEmailCaptchaResponse) SetStatusCode(v int32) *SendAsyncEmailCaptchaResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAsyncEmailCaptchaResponse) SetBody(v *SendAsyncEmailCaptchaResponseBody) *SendAsyncEmailCaptchaResponse {
	s.Body = v
	return s
}

func (s *SendAsyncEmailCaptchaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
