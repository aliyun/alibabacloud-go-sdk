// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncMobileCaptchaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAsyncMobileCaptchaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAsyncMobileCaptchaResponse
	GetStatusCode() *int32
	SetBody(v *SendAsyncMobileCaptchaResponseBody) *SendAsyncMobileCaptchaResponse
	GetBody() *SendAsyncMobileCaptchaResponseBody
}

type SendAsyncMobileCaptchaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAsyncMobileCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAsyncMobileCaptchaResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncMobileCaptchaResponse) GoString() string {
	return s.String()
}

func (s *SendAsyncMobileCaptchaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAsyncMobileCaptchaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAsyncMobileCaptchaResponse) GetBody() *SendAsyncMobileCaptchaResponseBody {
	return s.Body
}

func (s *SendAsyncMobileCaptchaResponse) SetHeaders(v map[string]*string) *SendAsyncMobileCaptchaResponse {
	s.Headers = v
	return s
}

func (s *SendAsyncMobileCaptchaResponse) SetStatusCode(v int32) *SendAsyncMobileCaptchaResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAsyncMobileCaptchaResponse) SetBody(v *SendAsyncMobileCaptchaResponseBody) *SendAsyncMobileCaptchaResponse {
	s.Body = v
	return s
}

func (s *SendAsyncMobileCaptchaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
