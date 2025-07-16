// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateImageResponse
	GetStatusCode() *int32
	SetBody(v *TranslateImageResponseBody) *TranslateImageResponse
	GetBody() *TranslateImageResponseBody
}

type TranslateImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateImageResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageResponse) GoString() string {
	return s.String()
}

func (s *TranslateImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateImageResponse) GetBody() *TranslateImageResponseBody {
	return s.Body
}

func (s *TranslateImageResponse) SetHeaders(v map[string]*string) *TranslateImageResponse {
	s.Headers = v
	return s
}

func (s *TranslateImageResponse) SetStatusCode(v int32) *TranslateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateImageResponse) SetBody(v *TranslateImageResponseBody) *TranslateImageResponse {
	s.Body = v
	return s
}

func (s *TranslateImageResponse) Validate() error {
	return dara.Validate(s)
}
