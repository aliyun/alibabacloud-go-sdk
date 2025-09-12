// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTranscriptionPhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTranscriptionPhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTranscriptionPhrasesResponse
	GetStatusCode() *int32
	SetBody(v *CreateTranscriptionPhrasesResponseBody) *CreateTranscriptionPhrasesResponse
	GetBody() *CreateTranscriptionPhrasesResponseBody
}

type CreateTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTranscriptionPhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTranscriptionPhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTranscriptionPhrasesResponse) GetBody() *CreateTranscriptionPhrasesResponseBody {
	return s.Body
}

func (s *CreateTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *CreateTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *CreateTranscriptionPhrasesResponse) SetStatusCode(v int32) *CreateTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponse) SetBody(v *CreateTranscriptionPhrasesResponseBody) *CreateTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

func (s *CreateTranscriptionPhrasesResponse) Validate() error {
	return dara.Validate(s)
}
