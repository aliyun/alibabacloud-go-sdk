// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSmsVerifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSmsVerifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSmsVerifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *CheckSmsVerifyCodeResponseBody) *CheckSmsVerifyCodeResponse
	GetBody() *CheckSmsVerifyCodeResponseBody
}

type CheckSmsVerifyCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSmsVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSmsVerifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSmsVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSmsVerifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSmsVerifyCodeResponse) GetBody() *CheckSmsVerifyCodeResponseBody {
	return s.Body
}

func (s *CheckSmsVerifyCodeResponse) SetHeaders(v map[string]*string) *CheckSmsVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *CheckSmsVerifyCodeResponse) SetStatusCode(v int32) *CheckSmsVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSmsVerifyCodeResponse) SetBody(v *CheckSmsVerifyCodeResponseBody) *CheckSmsVerifyCodeResponse {
	s.Body = v
	return s
}

func (s *CheckSmsVerifyCodeResponse) Validate() error {
	return dara.Validate(s)
}
