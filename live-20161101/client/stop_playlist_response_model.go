// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPlaylistResponse
	GetStatusCode() *int32
	SetBody(v *StopPlaylistResponseBody) *StopPlaylistResponse
	GetBody() *StopPlaylistResponseBody
}

type StopPlaylistResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPlaylistResponse) GoString() string {
	return s.String()
}

func (s *StopPlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPlaylistResponse) GetBody() *StopPlaylistResponseBody {
	return s.Body
}

func (s *StopPlaylistResponse) SetHeaders(v map[string]*string) *StopPlaylistResponse {
	s.Headers = v
	return s
}

func (s *StopPlaylistResponse) SetStatusCode(v int32) *StopPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPlaylistResponse) SetBody(v *StopPlaylistResponseBody) *StopPlaylistResponse {
	s.Body = v
	return s
}

func (s *StopPlaylistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
