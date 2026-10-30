// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteMapResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteMapResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteMapResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteMapResponseBody) *DeleteRouteMapResponse
	GetBody() *DeleteRouteMapResponseBody
}

type DeleteRouteMapResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteMapResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteMapResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteMapResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteMapResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteMapResponse) GetBody() *DeleteRouteMapResponseBody {
	return s.Body
}

func (s *DeleteRouteMapResponse) SetHeaders(v map[string]*string) *DeleteRouteMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteMapResponse) SetStatusCode(v int32) *DeleteRouteMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteMapResponse) SetBody(v *DeleteRouteMapResponseBody) *DeleteRouteMapResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteMapResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
