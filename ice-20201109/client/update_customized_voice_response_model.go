// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomizedVoiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomizedVoiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomizedVoiceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomizedVoiceResponseBody) *UpdateCustomizedVoiceResponse
	GetBody() *UpdateCustomizedVoiceResponseBody
}

type UpdateCustomizedVoiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomizedVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomizedVoiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomizedVoiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomizedVoiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomizedVoiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomizedVoiceResponse) GetBody() *UpdateCustomizedVoiceResponseBody {
	return s.Body
}

func (s *UpdateCustomizedVoiceResponse) SetHeaders(v map[string]*string) *UpdateCustomizedVoiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomizedVoiceResponse) SetStatusCode(v int32) *UpdateCustomizedVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomizedVoiceResponse) SetBody(v *UpdateCustomizedVoiceResponseBody) *UpdateCustomizedVoiceResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomizedVoiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
