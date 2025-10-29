// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPlaylistResponse
	GetStatusCode() *int32
	SetBody(v *StartPlaylistResponseBody) *StartPlaylistResponse
	GetBody() *StartPlaylistResponseBody
}

type StartPlaylistResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPlaylistResponse) GoString() string {
	return s.String()
}

func (s *StartPlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPlaylistResponse) GetBody() *StartPlaylistResponseBody {
	return s.Body
}

func (s *StartPlaylistResponse) SetHeaders(v map[string]*string) *StartPlaylistResponse {
	s.Headers = v
	return s
}

func (s *StartPlaylistResponse) SetStatusCode(v int32) *StartPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPlaylistResponse) SetBody(v *StartPlaylistResponseBody) *StartPlaylistResponse {
	s.Body = v
	return s
}

func (s *StartPlaylistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
