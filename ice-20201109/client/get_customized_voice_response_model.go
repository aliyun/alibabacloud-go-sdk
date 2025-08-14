// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomizedVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomizedVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomizedVoiceResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomizedVoiceResponseBody) *GetCustomizedVoiceResponse
	GetBody() *GetCustomizedVoiceResponseBody
}

type GetCustomizedVoiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomizedVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomizedVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomizedVoiceResponse) GoString() string {
	return s.String()
}

func (s *GetCustomizedVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomizedVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomizedVoiceResponse) GetBody() *GetCustomizedVoiceResponseBody {
	return s.Body
}

func (s *GetCustomizedVoiceResponse) SetHeaders(v map[string]*string) *GetCustomizedVoiceResponse {
	s.Headers = v
	return s
}

func (s *GetCustomizedVoiceResponse) SetStatusCode(v int32) *GetCustomizedVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomizedVoiceResponse) SetBody(v *GetCustomizedVoiceResponseBody) *GetCustomizedVoiceResponse {
	s.Body = v
	return s
}

func (s *GetCustomizedVoiceResponse) Validate() error {
	return dara.Validate(s)
}
