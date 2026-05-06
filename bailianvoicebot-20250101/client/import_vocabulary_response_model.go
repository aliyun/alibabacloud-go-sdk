// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportVocabularyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportVocabularyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportVocabularyResponse
	GetStatusCode() *int32
	SetBody(v *ImportVocabularyResponseBody) *ImportVocabularyResponse
	GetBody() *ImportVocabularyResponseBody
}

type ImportVocabularyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportVocabularyResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportVocabularyResponse) GoString() string {
	return s.String()
}

func (s *ImportVocabularyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportVocabularyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportVocabularyResponse) GetBody() *ImportVocabularyResponseBody {
	return s.Body
}

func (s *ImportVocabularyResponse) SetHeaders(v map[string]*string) *ImportVocabularyResponse {
	s.Headers = v
	return s
}

func (s *ImportVocabularyResponse) SetStatusCode(v int32) *ImportVocabularyResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportVocabularyResponse) SetBody(v *ImportVocabularyResponseBody) *ImportVocabularyResponse {
	s.Body = v
	return s
}

func (s *ImportVocabularyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
