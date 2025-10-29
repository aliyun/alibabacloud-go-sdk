// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRouteResponseBody) *DeleteRouteResponse
	GetBody() *DeleteRouteResponseBody
}

type DeleteRouteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRouteResponse) GetBody() *DeleteRouteResponseBody {
	return s.Body
}

func (s *DeleteRouteResponse) SetHeaders(v map[string]*string) *DeleteRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteRouteResponse) SetStatusCode(v int32) *DeleteRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRouteResponse) SetBody(v *DeleteRouteResponseBody) *DeleteRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
