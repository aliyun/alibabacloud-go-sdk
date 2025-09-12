// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTranscriptionPhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTranscriptionPhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTranscriptionPhrasesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTranscriptionPhrasesResponseBody) *DeleteTranscriptionPhrasesResponse
	GetBody() *DeleteTranscriptionPhrasesResponseBody
}

type DeleteTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTranscriptionPhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTranscriptionPhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTranscriptionPhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTranscriptionPhrasesResponse) GetBody() *DeleteTranscriptionPhrasesResponseBody {
	return s.Body
}

func (s *DeleteTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *DeleteTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTranscriptionPhrasesResponse) SetStatusCode(v int32) *DeleteTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponse) SetBody(v *DeleteTranscriptionPhrasesResponseBody) *DeleteTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

func (s *DeleteTranscriptionPhrasesResponse) Validate() error {
	return dara.Validate(s)
}
