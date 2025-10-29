// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterEpisodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterEpisodeResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterEpisodeResponseBody) *AddCasterEpisodeResponse
	GetBody() *AddCasterEpisodeResponseBody
}

type AddCasterEpisodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterEpisodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterEpisodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeResponse) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterEpisodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterEpisodeResponse) GetBody() *AddCasterEpisodeResponseBody {
	return s.Body
}

func (s *AddCasterEpisodeResponse) SetHeaders(v map[string]*string) *AddCasterEpisodeResponse {
	s.Headers = v
	return s
}

func (s *AddCasterEpisodeResponse) SetStatusCode(v int32) *AddCasterEpisodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterEpisodeResponse) SetBody(v *AddCasterEpisodeResponseBody) *AddCasterEpisodeResponse {
	s.Body = v
	return s
}

func (s *AddCasterEpisodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
