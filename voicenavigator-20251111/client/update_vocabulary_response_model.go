// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabularyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVocabularyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVocabularyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVocabularyResponseBody) *UpdateVocabularyResponse
	GetBody() *UpdateVocabularyResponseBody
}

type UpdateVocabularyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVocabularyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabularyResponse) GoString() string {
	return s.String()
}

func (s *UpdateVocabularyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVocabularyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVocabularyResponse) GetBody() *UpdateVocabularyResponseBody {
	return s.Body
}

func (s *UpdateVocabularyResponse) SetHeaders(v map[string]*string) *UpdateVocabularyResponse {
	s.Headers = v
	return s
}

func (s *UpdateVocabularyResponse) SetStatusCode(v int32) *UpdateVocabularyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVocabularyResponse) SetBody(v *UpdateVocabularyResponseBody) *UpdateVocabularyResponse {
	s.Body = v
	return s
}

func (s *UpdateVocabularyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
