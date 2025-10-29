// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPlaylistResponse
	GetStatusCode() *int32
	SetBody(v *ListPlaylistResponseBody) *ListPlaylistResponse
	GetBody() *ListPlaylistResponseBody
}

type ListPlaylistResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPlaylistResponse) GoString() string {
	return s.String()
}

func (s *ListPlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPlaylistResponse) GetBody() *ListPlaylistResponseBody {
	return s.Body
}

func (s *ListPlaylistResponse) SetHeaders(v map[string]*string) *ListPlaylistResponse {
	s.Headers = v
	return s
}

func (s *ListPlaylistResponse) SetStatusCode(v int32) *ListPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlaylistResponse) SetBody(v *ListPlaylistResponseBody) *ListPlaylistResponse {
	s.Body = v
	return s
}

func (s *ListPlaylistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
