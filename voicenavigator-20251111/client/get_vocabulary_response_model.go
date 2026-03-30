// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVocabularyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVocabularyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVocabularyResponse
	GetStatusCode() *int32
	SetBody(v *GetVocabularyResponseBody) *GetVocabularyResponse
	GetBody() *GetVocabularyResponseBody
}

type GetVocabularyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVocabularyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVocabularyResponse) GoString() string {
	return s.String()
}

func (s *GetVocabularyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVocabularyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVocabularyResponse) GetBody() *GetVocabularyResponseBody {
	return s.Body
}

func (s *GetVocabularyResponse) SetHeaders(v map[string]*string) *GetVocabularyResponse {
	s.Headers = v
	return s
}

func (s *GetVocabularyResponse) SetStatusCode(v int32) *GetVocabularyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVocabularyResponse) SetBody(v *GetVocabularyResponseBody) *GetVocabularyResponse {
	s.Body = v
	return s
}

func (s *GetVocabularyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
