// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsrVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAsrVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAsrVocabResponse
	GetStatusCode() *int32
	SetBody(v *CreateAsrVocabResponseBody) *CreateAsrVocabResponse
	GetBody() *CreateAsrVocabResponseBody
}

type CreateAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAsrVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *CreateAsrVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAsrVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAsrVocabResponse) GetBody() *CreateAsrVocabResponseBody {
	return s.Body
}

func (s *CreateAsrVocabResponse) SetHeaders(v map[string]*string) *CreateAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *CreateAsrVocabResponse) SetStatusCode(v int32) *CreateAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAsrVocabResponse) SetBody(v *CreateAsrVocabResponseBody) *CreateAsrVocabResponse {
	s.Body = v
	return s
}

func (s *CreateAsrVocabResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
