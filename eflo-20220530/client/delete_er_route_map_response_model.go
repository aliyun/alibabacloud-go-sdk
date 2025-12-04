// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteErRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteErRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *DeleteErRouteMapResponseBody) *DeleteErRouteMapResponse
	GetBody() *DeleteErRouteMapResponseBody
}

type DeleteErRouteMapResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteErRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteErRouteMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteErRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteErRouteMapResponse) GetBody() *DeleteErRouteMapResponseBody {
	return s.Body
}

func (s *DeleteErRouteMapResponse) SetHeaders(v map[string]*string) *DeleteErRouteMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteErRouteMapResponse) SetStatusCode(v int32) *DeleteErRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErRouteMapResponse) SetBody(v *DeleteErRouteMapResponseBody) *DeleteErRouteMapResponse {
	s.Body = v
	return s
}

func (s *DeleteErRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
