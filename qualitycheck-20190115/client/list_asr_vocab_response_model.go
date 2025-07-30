// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsrVocabResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAsrVocabResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAsrVocabResponse
	GetStatusCode() *int32
	SetBody(v *ListAsrVocabResponseBody) *ListAsrVocabResponse
	GetBody() *ListAsrVocabResponseBody
}

type ListAsrVocabResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsrVocabResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsrVocabResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAsrVocabResponse) GoString() string {
	return s.String()
}

func (s *ListAsrVocabResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAsrVocabResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAsrVocabResponse) GetBody() *ListAsrVocabResponseBody {
	return s.Body
}

func (s *ListAsrVocabResponse) SetHeaders(v map[string]*string) *ListAsrVocabResponse {
	s.Headers = v
	return s
}

func (s *ListAsrVocabResponse) SetStatusCode(v int32) *ListAsrVocabResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsrVocabResponse) SetBody(v *ListAsrVocabResponseBody) *ListAsrVocabResponse {
	s.Body = v
	return s
}

func (s *ListAsrVocabResponse) Validate() error {
	return dara.Validate(s)
}
