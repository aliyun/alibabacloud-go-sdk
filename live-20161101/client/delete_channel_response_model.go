// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteChannelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteChannelResponseBody) *DeleteChannelResponse
	GetBody() *DeleteChannelResponseBody
}

type DeleteChannelResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteChannelResponse) GetBody() *DeleteChannelResponseBody {
	return s.Body
}

func (s *DeleteChannelResponse) SetHeaders(v map[string]*string) *DeleteChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteChannelResponse) SetStatusCode(v int32) *DeleteChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChannelResponse) SetBody(v *DeleteChannelResponseBody) *DeleteChannelResponse {
	s.Body = v
	return s
}

func (s *DeleteChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
