// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDynamicRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDynamicRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDynamicRouteResponseBody) *DeleteDynamicRouteResponse
	GetBody() *DeleteDynamicRouteResponseBody
}

type DeleteDynamicRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDynamicRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteDynamicRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDynamicRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDynamicRouteResponse) GetBody() *DeleteDynamicRouteResponseBody {
	return s.Body
}

func (s *DeleteDynamicRouteResponse) SetHeaders(v map[string]*string) *DeleteDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteDynamicRouteResponse) SetStatusCode(v int32) *DeleteDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDynamicRouteResponse) SetBody(v *DeleteDynamicRouteResponseBody) *DeleteDynamicRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteDynamicRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
