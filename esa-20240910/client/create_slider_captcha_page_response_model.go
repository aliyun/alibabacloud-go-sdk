// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSliderCaptchaPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSliderCaptchaPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSliderCaptchaPageResponse
	GetStatusCode() *int32
	SetBody(v *CreateSliderCaptchaPageResponseBody) *CreateSliderCaptchaPageResponse
	GetBody() *CreateSliderCaptchaPageResponseBody
}

type CreateSliderCaptchaPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSliderCaptchaPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSliderCaptchaPageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageResponse) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSliderCaptchaPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSliderCaptchaPageResponse) GetBody() *CreateSliderCaptchaPageResponseBody {
	return s.Body
}

func (s *CreateSliderCaptchaPageResponse) SetHeaders(v map[string]*string) *CreateSliderCaptchaPageResponse {
	s.Headers = v
	return s
}

func (s *CreateSliderCaptchaPageResponse) SetStatusCode(v int32) *CreateSliderCaptchaPageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSliderCaptchaPageResponse) SetBody(v *CreateSliderCaptchaPageResponseBody) *CreateSliderCaptchaPageResponse {
	s.Body = v
	return s
}

func (s *CreateSliderCaptchaPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
