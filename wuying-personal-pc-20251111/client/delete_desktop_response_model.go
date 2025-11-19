// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDesktopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDesktopResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDesktopResponseBody) *DeleteDesktopResponse
	GetBody() *DeleteDesktopResponseBody
}

type DeleteDesktopResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDesktopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDesktopResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopResponse) GoString() string {
	return s.String()
}

func (s *DeleteDesktopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDesktopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDesktopResponse) GetBody() *DeleteDesktopResponseBody {
	return s.Body
}

func (s *DeleteDesktopResponse) SetHeaders(v map[string]*string) *DeleteDesktopResponse {
	s.Headers = v
	return s
}

func (s *DeleteDesktopResponse) SetStatusCode(v int32) *DeleteDesktopResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDesktopResponse) SetBody(v *DeleteDesktopResponseBody) *DeleteDesktopResponse {
	s.Body = v
	return s
}

func (s *DeleteDesktopResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
