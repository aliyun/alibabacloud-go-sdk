// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCustomizedVoiceJobResponseBody) *SubmitCustomizedVoiceJobResponse
	GetBody() *SubmitCustomizedVoiceJobResponseBody
}

type SubmitCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCustomizedVoiceJobResponse) GetBody() *SubmitCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *SubmitCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *SubmitCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCustomizedVoiceJobResponse) SetStatusCode(v int32) *SubmitCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCustomizedVoiceJobResponse) SetBody(v *SubmitCustomizedVoiceJobResponseBody) *SubmitCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *SubmitCustomizedVoiceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
