// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDesktopImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDesktopImageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDesktopImageResponseBody) *DeleteDesktopImageResponse
	GetBody() *DeleteDesktopImageResponseBody
}

type DeleteDesktopImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDesktopImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDesktopImageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDesktopImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDesktopImageResponse) GetBody() *DeleteDesktopImageResponseBody {
	return s.Body
}

func (s *DeleteDesktopImageResponse) SetHeaders(v map[string]*string) *DeleteDesktopImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopImageResponse) SetStatusCode(v int32) *DeleteDesktopImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDesktopImageResponse) SetBody(v *DeleteDesktopImageResponseBody) *DeleteDesktopImageResponse {
	s.Body = v
	return s
}

func (s *DeleteDesktopImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
