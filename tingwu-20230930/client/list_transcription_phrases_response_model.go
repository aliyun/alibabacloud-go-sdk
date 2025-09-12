// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscriptionPhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTranscriptionPhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTranscriptionPhrasesResponse
	GetStatusCode() *int32
	SetBody(v *ListTranscriptionPhrasesResponseBody) *ListTranscriptionPhrasesResponse
	GetBody() *ListTranscriptionPhrasesResponseBody
}

type ListTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTranscriptionPhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTranscriptionPhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTranscriptionPhrasesResponse) GetBody() *ListTranscriptionPhrasesResponseBody {
	return s.Body
}

func (s *ListTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *ListTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *ListTranscriptionPhrasesResponse) SetStatusCode(v int32) *ListTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTranscriptionPhrasesResponse) SetBody(v *ListTranscriptionPhrasesResponseBody) *ListTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

func (s *ListTranscriptionPhrasesResponse) Validate() error {
	return dara.Validate(s)
}
