// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAudioNoteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAudioNoteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAudioNoteResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAudioNoteResponseBody) *SubmitAudioNoteResponse
	GetBody() *SubmitAudioNoteResponseBody
}

type SubmitAudioNoteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAudioNoteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAudioNoteResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAudioNoteResponse) GoString() string {
	return s.String()
}

func (s *SubmitAudioNoteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAudioNoteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAudioNoteResponse) GetBody() *SubmitAudioNoteResponseBody {
	return s.Body
}

func (s *SubmitAudioNoteResponse) SetHeaders(v map[string]*string) *SubmitAudioNoteResponse {
	s.Headers = v
	return s
}

func (s *SubmitAudioNoteResponse) SetStatusCode(v int32) *SubmitAudioNoteResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAudioNoteResponse) SetBody(v *SubmitAudioNoteResponseBody) *SubmitAudioNoteResponse {
	s.Body = v
	return s
}

func (s *SubmitAudioNoteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
