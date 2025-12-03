// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateResponse
	GetStatusCode() *int32
	SetBody(v *TranslateResponseBody) *TranslateResponse
	GetBody() *TranslateResponseBody
}

type TranslateResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateResponse) GoString() string {
	return s.String()
}

func (s *TranslateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateResponse) GetBody() *TranslateResponseBody {
	return s.Body
}

func (s *TranslateResponse) SetHeaders(v map[string]*string) *TranslateResponse {
	s.Headers = v
	return s
}

func (s *TranslateResponse) SetStatusCode(v int32) *TranslateResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateResponse) SetBody(v *TranslateResponseBody) *TranslateResponse {
	s.Body = v
	return s
}

func (s *TranslateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
