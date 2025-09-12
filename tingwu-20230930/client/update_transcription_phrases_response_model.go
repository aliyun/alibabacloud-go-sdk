// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTranscriptionPhrasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTranscriptionPhrasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTranscriptionPhrasesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTranscriptionPhrasesResponseBody) *UpdateTranscriptionPhrasesResponse
	GetBody() *UpdateTranscriptionPhrasesResponseBody
}

type UpdateTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTranscriptionPhrasesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTranscriptionPhrasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTranscriptionPhrasesResponse) GetBody() *UpdateTranscriptionPhrasesResponseBody {
	return s.Body
}

func (s *UpdateTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *UpdateTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTranscriptionPhrasesResponse) SetStatusCode(v int32) *UpdateTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponse) SetBody(v *UpdateTranscriptionPhrasesResponseBody) *UpdateTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

func (s *UpdateTranscriptionPhrasesResponse) Validate() error {
	return dara.Validate(s)
}
