// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomizedVoiceJobResponseBody) *CreateCustomizedVoiceJobResponse
	GetBody() *CreateCustomizedVoiceJobResponseBody
}

type CreateCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomizedVoiceJobResponse) GetBody() *CreateCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *CreateCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *CreateCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedVoiceJobResponse) SetStatusCode(v int32) *CreateCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedVoiceJobResponse) SetBody(v *CreateCustomizedVoiceJobResponseBody) *CreateCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *CreateCustomizedVoiceJobResponse) Validate() error {
	return dara.Validate(s)
}
