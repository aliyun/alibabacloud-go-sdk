// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateGeneralResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateGeneralResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateGeneralResponse
	GetStatusCode() *int32
	SetBody(v *TranslateGeneralResponseBody) *TranslateGeneralResponse
	GetBody() *TranslateGeneralResponseBody
}

type TranslateGeneralResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateGeneralResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateGeneralResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralResponse) GoString() string {
	return s.String()
}

func (s *TranslateGeneralResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateGeneralResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateGeneralResponse) GetBody() *TranslateGeneralResponseBody {
	return s.Body
}

func (s *TranslateGeneralResponse) SetHeaders(v map[string]*string) *TranslateGeneralResponse {
	s.Headers = v
	return s
}

func (s *TranslateGeneralResponse) SetStatusCode(v int32) *TranslateGeneralResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateGeneralResponse) SetBody(v *TranslateGeneralResponseBody) *TranslateGeneralResponse {
	s.Body = v
	return s
}

func (s *TranslateGeneralResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
