// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterRouteTableResponseBody) *DeleteTransitRouterRouteTableResponse
	GetBody() *DeleteTransitRouterRouteTableResponseBody
}

type DeleteTransitRouterRouteTableResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterRouteTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterRouteTableResponse) GetBody() *DeleteTransitRouterRouteTableResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterRouteTableResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterRouteTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterRouteTableResponse) SetStatusCode(v int32) *DeleteTransitRouterRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterRouteTableResponse) SetBody(v *DeleteTransitRouterRouteTableResponseBody) *DeleteTransitRouterRouteTableResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
