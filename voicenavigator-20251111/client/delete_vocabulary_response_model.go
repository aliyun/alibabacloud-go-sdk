// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVocabularyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVocabularyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVocabularyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVocabularyResponseBody) *DeleteVocabularyResponse
	GetBody() *DeleteVocabularyResponseBody
}

type DeleteVocabularyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVocabularyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVocabularyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVocabularyResponse) GoString() string {
	return s.String()
}

func (s *DeleteVocabularyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVocabularyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVocabularyResponse) GetBody() *DeleteVocabularyResponseBody {
	return s.Body
}

func (s *DeleteVocabularyResponse) SetHeaders(v map[string]*string) *DeleteVocabularyResponse {
	s.Headers = v
	return s
}

func (s *DeleteVocabularyResponse) SetStatusCode(v int32) *DeleteVocabularyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVocabularyResponse) SetBody(v *DeleteVocabularyResponseBody) *DeleteVocabularyResponse {
	s.Body = v
	return s
}

func (s *DeleteVocabularyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
