// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoPlaylistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateVideoPlaylistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateVideoPlaylistResponse
	GetStatusCode() *int32
	SetBody(v *GenerateVideoPlaylistResponseBody) *GenerateVideoPlaylistResponse
	GetBody() *GenerateVideoPlaylistResponseBody
}

type GenerateVideoPlaylistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateVideoPlaylistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateVideoPlaylistResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateVideoPlaylistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateVideoPlaylistResponse) GetBody() *GenerateVideoPlaylistResponseBody {
	return s.Body
}

func (s *GenerateVideoPlaylistResponse) SetHeaders(v map[string]*string) *GenerateVideoPlaylistResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoPlaylistResponse) SetStatusCode(v int32) *GenerateVideoPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateVideoPlaylistResponse) SetBody(v *GenerateVideoPlaylistResponseBody) *GenerateVideoPlaylistResponse {
	s.Body = v
	return s
}

func (s *GenerateVideoPlaylistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
