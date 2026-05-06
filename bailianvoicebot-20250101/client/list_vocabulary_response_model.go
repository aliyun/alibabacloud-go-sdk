// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabularyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVocabularyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVocabularyResponse
	GetStatusCode() *int32
	SetBody(v *ListVocabularyResponseBody) *ListVocabularyResponse
	GetBody() *ListVocabularyResponseBody
}

type ListVocabularyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVocabularyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVocabularyResponse) GoString() string {
	return s.String()
}

func (s *ListVocabularyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVocabularyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVocabularyResponse) GetBody() *ListVocabularyResponseBody {
	return s.Body
}

func (s *ListVocabularyResponse) SetHeaders(v map[string]*string) *ListVocabularyResponse {
	s.Headers = v
	return s
}

func (s *ListVocabularyResponse) SetStatusCode(v int32) *ListVocabularyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVocabularyResponse) SetBody(v *ListVocabularyResponseBody) *ListVocabularyResponse {
	s.Body = v
	return s
}

func (s *ListVocabularyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
