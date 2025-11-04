// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomizedVoiceJobResponseBody) *GetCustomizedVoiceJobResponse
	GetBody() *GetCustomizedVoiceJobResponseBody
}

type GetCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomizedVoiceJobResponse) GetBody() *GetCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *GetCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *GetCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *GetCustomizedVoiceJobResponse) SetStatusCode(v int32) *GetCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomizedVoiceJobResponse) SetBody(v *GetCustomizedVoiceJobResponseBody) *GetCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *GetCustomizedVoiceJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
