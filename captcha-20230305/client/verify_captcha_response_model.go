// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCaptchaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCaptchaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCaptchaResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCaptchaResponseBody) *VerifyCaptchaResponse
	GetBody() *VerifyCaptchaResponseBody
}

type VerifyCaptchaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCaptchaResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCaptchaResponse) GoString() string {
	return s.String()
}

func (s *VerifyCaptchaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCaptchaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCaptchaResponse) GetBody() *VerifyCaptchaResponseBody {
	return s.Body
}

func (s *VerifyCaptchaResponse) SetHeaders(v map[string]*string) *VerifyCaptchaResponse {
	s.Headers = v
	return s
}

func (s *VerifyCaptchaResponse) SetStatusCode(v int32) *VerifyCaptchaResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCaptchaResponse) SetBody(v *VerifyCaptchaResponseBody) *VerifyCaptchaResponse {
	s.Body = v
	return s
}

func (s *VerifyCaptchaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
