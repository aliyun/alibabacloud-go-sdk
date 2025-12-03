// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TranslateSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TranslateSearchResponse
	GetStatusCode() *int32
	SetBody(v *TranslateSearchResponseBody) *TranslateSearchResponse
	GetBody() *TranslateSearchResponseBody
}

type TranslateSearchResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TranslateSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TranslateSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s TranslateSearchResponse) GoString() string {
	return s.String()
}

func (s *TranslateSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TranslateSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TranslateSearchResponse) GetBody() *TranslateSearchResponseBody {
	return s.Body
}

func (s *TranslateSearchResponse) SetHeaders(v map[string]*string) *TranslateSearchResponse {
	s.Headers = v
	return s
}

func (s *TranslateSearchResponse) SetStatusCode(v int32) *TranslateSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *TranslateSearchResponse) SetBody(v *TranslateSearchResponseBody) *TranslateSearchResponse {
	s.Body = v
	return s
}

func (s *TranslateSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
