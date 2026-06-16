// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLanguageDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LanguageDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LanguageDetectResponse
	GetStatusCode() *int32
	SetBody(v *LanguageDetectResponseBody) *LanguageDetectResponse
	GetBody() *LanguageDetectResponseBody
}

type LanguageDetectResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LanguageDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LanguageDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s LanguageDetectResponse) GoString() string {
	return s.String()
}

func (s *LanguageDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LanguageDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LanguageDetectResponse) GetBody() *LanguageDetectResponseBody {
	return s.Body
}

func (s *LanguageDetectResponse) SetHeaders(v map[string]*string) *LanguageDetectResponse {
	s.Headers = v
	return s
}

func (s *LanguageDetectResponse) SetStatusCode(v int32) *LanguageDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *LanguageDetectResponse) SetBody(v *LanguageDetectResponseBody) *LanguageDetectResponse {
	s.Body = v
	return s
}

func (s *LanguageDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
