// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyIntelligentCaptchaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyIntelligentCaptchaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyIntelligentCaptchaResponse
	GetStatusCode() *int32
	SetBody(v *VerifyIntelligentCaptchaResponseBody) *VerifyIntelligentCaptchaResponse
	GetBody() *VerifyIntelligentCaptchaResponseBody
}

type VerifyIntelligentCaptchaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyIntelligentCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyIntelligentCaptchaResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyIntelligentCaptchaResponse) GoString() string {
	return s.String()
}

func (s *VerifyIntelligentCaptchaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyIntelligentCaptchaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyIntelligentCaptchaResponse) GetBody() *VerifyIntelligentCaptchaResponseBody {
	return s.Body
}

func (s *VerifyIntelligentCaptchaResponse) SetHeaders(v map[string]*string) *VerifyIntelligentCaptchaResponse {
	s.Headers = v
	return s
}

func (s *VerifyIntelligentCaptchaResponse) SetStatusCode(v int32) *VerifyIntelligentCaptchaResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyIntelligentCaptchaResponse) SetBody(v *VerifyIntelligentCaptchaResponseBody) *VerifyIntelligentCaptchaResponse {
	s.Body = v
	return s
}

func (s *VerifyIntelligentCaptchaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
