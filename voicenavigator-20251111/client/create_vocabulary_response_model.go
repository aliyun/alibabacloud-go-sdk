// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVocabularyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVocabularyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVocabularyResponse
	GetStatusCode() *int32
	SetBody(v *CreateVocabularyResponseBody) *CreateVocabularyResponse
	GetBody() *CreateVocabularyResponseBody
}

type CreateVocabularyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVocabularyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVocabularyResponse) GoString() string {
	return s.String()
}

func (s *CreateVocabularyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVocabularyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVocabularyResponse) GetBody() *CreateVocabularyResponseBody {
	return s.Body
}

func (s *CreateVocabularyResponse) SetHeaders(v map[string]*string) *CreateVocabularyResponse {
	s.Headers = v
	return s
}

func (s *CreateVocabularyResponse) SetStatusCode(v int32) *CreateVocabularyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVocabularyResponse) SetBody(v *CreateVocabularyResponseBody) *CreateVocabularyResponse {
	s.Body = v
	return s
}

func (s *CreateVocabularyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
