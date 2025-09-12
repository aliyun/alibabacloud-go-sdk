// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscriptionPhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTranscriptionPhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTranscriptionPhrasesResponse
	GetStatusCode() *int32
	SetBody(v *GetTranscriptionPhrasesResponseBody) *GetTranscriptionPhrasesResponse
	GetBody() *GetTranscriptionPhrasesResponseBody
}

type GetTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscriptionPhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTranscriptionPhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTranscriptionPhrasesResponse) GetBody() *GetTranscriptionPhrasesResponseBody {
	return s.Body
}

func (s *GetTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *GetTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *GetTranscriptionPhrasesResponse) SetStatusCode(v int32) *GetTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscriptionPhrasesResponse) SetBody(v *GetTranscriptionPhrasesResponseBody) *GetTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

func (s *GetTranscriptionPhrasesResponse) Validate() error {
	return dara.Validate(s)
}
