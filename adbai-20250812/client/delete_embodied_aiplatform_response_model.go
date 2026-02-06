// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmbodiedAIPlatformResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEmbodiedAIPlatformResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEmbodiedAIPlatformResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEmbodiedAIPlatformResponseBody) *DeleteEmbodiedAIPlatformResponse
	GetBody() *DeleteEmbodiedAIPlatformResponseBody
}

type DeleteEmbodiedAIPlatformResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEmbodiedAIPlatformResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEmbodiedAIPlatformResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmbodiedAIPlatformResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmbodiedAIPlatformResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEmbodiedAIPlatformResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEmbodiedAIPlatformResponse) GetBody() *DeleteEmbodiedAIPlatformResponseBody {
	return s.Body
}

func (s *DeleteEmbodiedAIPlatformResponse) SetHeaders(v map[string]*string) *DeleteEmbodiedAIPlatformResponse {
	s.Headers = v
	return s
}

func (s *DeleteEmbodiedAIPlatformResponse) SetStatusCode(v int32) *DeleteEmbodiedAIPlatformResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEmbodiedAIPlatformResponse) SetBody(v *DeleteEmbodiedAIPlatformResponseBody) *DeleteEmbodiedAIPlatformResponse {
	s.Body = v
	return s
}

func (s *DeleteEmbodiedAIPlatformResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
