// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVocabResponse
	GetStatusCode() *int32
	SetBody(v *ListVocabResponseBody) *ListVocabResponse
	GetBody() *ListVocabResponseBody
}

type ListVocabResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVocabResponse) GoString() string {
	return s.String()
}

func (s *ListVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVocabResponse) GetBody() *ListVocabResponseBody {
	return s.Body
}

func (s *ListVocabResponse) SetHeaders(v map[string]*string) *ListVocabResponse {
	s.Headers = v
	return s
}

func (s *ListVocabResponse) SetStatusCode(v int32) *ListVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVocabResponse) SetBody(v *ListVocabResponseBody) *ListVocabResponse {
	s.Body = v
	return s
}

func (s *ListVocabResponse) Validate() error {
	return dara.Validate(s)
}
