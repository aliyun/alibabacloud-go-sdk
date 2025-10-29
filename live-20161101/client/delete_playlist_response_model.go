// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePlaylistResponse
	GetStatusCode() *int32
	SetBody(v *DeletePlaylistResponseBody) *DeletePlaylistResponse
	GetBody() *DeletePlaylistResponseBody
}

type DeletePlaylistResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaylistResponse) GoString() string {
	return s.String()
}

func (s *DeletePlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePlaylistResponse) GetBody() *DeletePlaylistResponseBody {
	return s.Body
}

func (s *DeletePlaylistResponse) SetHeaders(v map[string]*string) *DeletePlaylistResponse {
	s.Headers = v
	return s
}

func (s *DeletePlaylistResponse) SetStatusCode(v int32) *DeletePlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePlaylistResponse) SetBody(v *DeletePlaylistResponseBody) *DeletePlaylistResponse {
	s.Body = v
	return s
}

func (s *DeletePlaylistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
