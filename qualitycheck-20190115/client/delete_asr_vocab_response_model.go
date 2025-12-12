// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsrVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAsrVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAsrVocabResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAsrVocabResponseBody) *DeleteAsrVocabResponse
	GetBody() *DeleteAsrVocabResponseBody
}

type DeleteAsrVocabResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAsrVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAsrVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAsrVocabResponse) GetBody() *DeleteAsrVocabResponseBody {
	return s.Body
}

func (s *DeleteAsrVocabResponse) SetHeaders(v map[string]*string) *DeleteAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *DeleteAsrVocabResponse) SetStatusCode(v int32) *DeleteAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAsrVocabResponse) SetBody(v *DeleteAsrVocabResponseBody) *DeleteAsrVocabResponse {
	s.Body = v
	return s
}

func (s *DeleteAsrVocabResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
