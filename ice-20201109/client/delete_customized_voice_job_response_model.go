// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizedVoiceJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomizedVoiceJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomizedVoiceJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomizedVoiceJobResponseBody) *DeleteCustomizedVoiceJobResponse
	GetBody() *DeleteCustomizedVoiceJobResponseBody
}

type DeleteCustomizedVoiceJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizedVoiceJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizedVoiceJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizedVoiceJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizedVoiceJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomizedVoiceJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomizedVoiceJobResponse) GetBody() *DeleteCustomizedVoiceJobResponseBody {
	return s.Body
}

func (s *DeleteCustomizedVoiceJobResponse) SetHeaders(v map[string]*string) *DeleteCustomizedVoiceJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizedVoiceJobResponse) SetStatusCode(v int32) *DeleteCustomizedVoiceJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizedVoiceJobResponse) SetBody(v *DeleteCustomizedVoiceJobResponseBody) *DeleteCustomizedVoiceJobResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomizedVoiceJobResponse) Validate() error {
	return dara.Validate(s)
}
