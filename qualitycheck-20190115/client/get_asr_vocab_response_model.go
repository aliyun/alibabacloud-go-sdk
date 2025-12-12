// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsrVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAsrVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAsrVocabResponse
	GetStatusCode() *int32
	SetBody(v *GetAsrVocabResponseBody) *GetAsrVocabResponse
	GetBody() *GetAsrVocabResponseBody
}

type GetAsrVocabResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsrVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *GetAsrVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAsrVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAsrVocabResponse) GetBody() *GetAsrVocabResponseBody {
	return s.Body
}

func (s *GetAsrVocabResponse) SetHeaders(v map[string]*string) *GetAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *GetAsrVocabResponse) SetStatusCode(v int32) *GetAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsrVocabResponse) SetBody(v *GetAsrVocabResponseBody) *GetAsrVocabResponse {
	s.Body = v
	return s
}

func (s *GetAsrVocabResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
