// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAppVerifyCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAppVerifyCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAppVerifyCodeResponse
	GetStatusCode() *int32
	SetBody(v *CheckAppVerifyCodeResponseBody) *CheckAppVerifyCodeResponse
	GetBody() *CheckAppVerifyCodeResponseBody
}

type CheckAppVerifyCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAppVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAppVerifyCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAppVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *CheckAppVerifyCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAppVerifyCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAppVerifyCodeResponse) GetBody() *CheckAppVerifyCodeResponseBody {
	return s.Body
}

func (s *CheckAppVerifyCodeResponse) SetHeaders(v map[string]*string) *CheckAppVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *CheckAppVerifyCodeResponse) SetStatusCode(v int32) *CheckAppVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAppVerifyCodeResponse) SetBody(v *CheckAppVerifyCodeResponseBody) *CheckAppVerifyCodeResponse {
	s.Body = v
	return s
}

func (s *CheckAppVerifyCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
