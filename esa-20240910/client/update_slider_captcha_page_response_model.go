// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSliderCaptchaPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSliderCaptchaPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSliderCaptchaPageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSliderCaptchaPageResponseBody) *UpdateSliderCaptchaPageResponse
	GetBody() *UpdateSliderCaptchaPageResponseBody
}

type UpdateSliderCaptchaPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSliderCaptchaPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSliderCaptchaPageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageResponse) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSliderCaptchaPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSliderCaptchaPageResponse) GetBody() *UpdateSliderCaptchaPageResponseBody {
	return s.Body
}

func (s *UpdateSliderCaptchaPageResponse) SetHeaders(v map[string]*string) *UpdateSliderCaptchaPageResponse {
	s.Headers = v
	return s
}

func (s *UpdateSliderCaptchaPageResponse) SetStatusCode(v int32) *UpdateSliderCaptchaPageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSliderCaptchaPageResponse) SetBody(v *UpdateSliderCaptchaPageResponseBody) *UpdateSliderCaptchaPageResponse {
	s.Body = v
	return s
}

func (s *UpdateSliderCaptchaPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
