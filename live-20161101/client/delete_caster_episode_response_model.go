// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterEpisodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterEpisodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterEpisodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterEpisodeResponseBody) *DeleteCasterEpisodeResponse
	GetBody() *DeleteCasterEpisodeResponseBody
}

type DeleteCasterEpisodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterEpisodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterEpisodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterEpisodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterEpisodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterEpisodeResponse) GetBody() *DeleteCasterEpisodeResponseBody {
	return s.Body
}

func (s *DeleteCasterEpisodeResponse) SetHeaders(v map[string]*string) *DeleteCasterEpisodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterEpisodeResponse) SetStatusCode(v int32) *DeleteCasterEpisodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterEpisodeResponse) SetBody(v *DeleteCasterEpisodeResponseBody) *DeleteCasterEpisodeResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterEpisodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
