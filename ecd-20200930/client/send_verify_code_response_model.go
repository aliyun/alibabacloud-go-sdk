// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendVerifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendVerifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *SendVerifyCodeResponseBody) *SendVerifyCodeResponse
	GetBody() *SendVerifyCodeResponseBody
}

type SendVerifyCodeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s SendVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *SendVerifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendVerifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendVerifyCodeResponse) GetBody() *SendVerifyCodeResponseBody {
	return s.Body
}

func (s *SendVerifyCodeResponse) SetHeaders(v map[string]*string) *SendVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *SendVerifyCodeResponse) SetStatusCode(v int32) *SendVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerifyCodeResponse) SetBody(v *SendVerifyCodeResponseBody) *SendVerifyCodeResponse {
	s.Body = v
	return s
}

func (s *SendVerifyCodeResponse) Validate() error {
	return dara.Validate(s)
}
