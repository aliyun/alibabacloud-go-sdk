// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterRouteTableResponseBody) *UpdateTransitRouterRouteTableResponse
	GetBody() *UpdateTransitRouterRouteTableResponseBody
}

type UpdateTransitRouterRouteTableResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterRouteTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterRouteTableResponse) GetBody() *UpdateTransitRouterRouteTableResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterRouteTableResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterRouteTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterRouteTableResponse) SetStatusCode(v int32) *UpdateTransitRouterRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterRouteTableResponse) SetBody(v *UpdateTransitRouterRouteTableResponseBody) *UpdateTransitRouterRouteTableResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
