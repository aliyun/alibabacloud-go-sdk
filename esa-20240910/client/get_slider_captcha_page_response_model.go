// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSliderCaptchaPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSliderCaptchaPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSliderCaptchaPageResponse
	GetStatusCode() *int32
	SetBody(v *GetSliderCaptchaPageResponseBody) *GetSliderCaptchaPageResponse
	GetBody() *GetSliderCaptchaPageResponseBody
}

type GetSliderCaptchaPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSliderCaptchaPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSliderCaptchaPageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSliderCaptchaPageResponse) GoString() string {
	return s.String()
}

func (s *GetSliderCaptchaPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSliderCaptchaPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSliderCaptchaPageResponse) GetBody() *GetSliderCaptchaPageResponseBody {
	return s.Body
}

func (s *GetSliderCaptchaPageResponse) SetHeaders(v map[string]*string) *GetSliderCaptchaPageResponse {
	s.Headers = v
	return s
}

func (s *GetSliderCaptchaPageResponse) SetStatusCode(v int32) *GetSliderCaptchaPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSliderCaptchaPageResponse) SetBody(v *GetSliderCaptchaPageResponseBody) *GetSliderCaptchaPageResponse {
	s.Body = v
	return s
}

func (s *GetSliderCaptchaPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
