// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCosyVoiceCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitCosyVoiceCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitCosyVoiceCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitCosyVoiceCustomizedVoiceJobResponseBody) *SubmitCosyVoiceCustomizedVoiceJobResponse
	GetBody() *SubmitCosyVoiceCustomizedVoiceJobResponseBody
}

type SubmitCosyVoiceCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitCosyVoiceCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitCosyVoiceCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitCosyVoiceCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) GetBody() *SubmitCosyVoiceCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *SubmitCosyVoiceCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) SetStatusCode(v int32) *SubmitCosyVoiceCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) SetBody(v *SubmitCosyVoiceCustomizedVoiceJobResponseBody) *SubmitCosyVoiceCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
