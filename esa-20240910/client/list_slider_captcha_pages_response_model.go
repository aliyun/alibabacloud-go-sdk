// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSliderCaptchaPagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSliderCaptchaPagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSliderCaptchaPagesResponse
	GetStatusCode() *int32
	SetBody(v *ListSliderCaptchaPagesResponseBody) *ListSliderCaptchaPagesResponse
	GetBody() *ListSliderCaptchaPagesResponseBody
}

type ListSliderCaptchaPagesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSliderCaptchaPagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSliderCaptchaPagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSliderCaptchaPagesResponse) GoString() string {
	return s.String()
}

func (s *ListSliderCaptchaPagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSliderCaptchaPagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSliderCaptchaPagesResponse) GetBody() *ListSliderCaptchaPagesResponseBody {
	return s.Body
}

func (s *ListSliderCaptchaPagesResponse) SetHeaders(v map[string]*string) *ListSliderCaptchaPagesResponse {
	s.Headers = v
	return s
}

func (s *ListSliderCaptchaPagesResponse) SetStatusCode(v int32) *ListSliderCaptchaPagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSliderCaptchaPagesResponse) SetBody(v *ListSliderCaptchaPagesResponseBody) *ListSliderCaptchaPagesResponse {
	s.Body = v
	return s
}

func (s *ListSliderCaptchaPagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
