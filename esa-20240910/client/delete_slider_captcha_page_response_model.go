// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSliderCaptchaPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSliderCaptchaPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSliderCaptchaPageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSliderCaptchaPageResponseBody) *DeleteSliderCaptchaPageResponse
	GetBody() *DeleteSliderCaptchaPageResponseBody
}

type DeleteSliderCaptchaPageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSliderCaptchaPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSliderCaptchaPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSliderCaptchaPageResponse) GoString() string {
	return s.String()
}

func (s *DeleteSliderCaptchaPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSliderCaptchaPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSliderCaptchaPageResponse) GetBody() *DeleteSliderCaptchaPageResponseBody {
	return s.Body
}

func (s *DeleteSliderCaptchaPageResponse) SetHeaders(v map[string]*string) *DeleteSliderCaptchaPageResponse {
	s.Headers = v
	return s
}

func (s *DeleteSliderCaptchaPageResponse) SetStatusCode(v int32) *DeleteSliderCaptchaPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSliderCaptchaPageResponse) SetBody(v *DeleteSliderCaptchaPageResponseBody) *DeleteSliderCaptchaPageResponse {
	s.Body = v
	return s
}

func (s *DeleteSliderCaptchaPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
