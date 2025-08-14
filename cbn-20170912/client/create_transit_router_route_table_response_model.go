// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRouteTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterRouteTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterRouteTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterRouteTableResponseBody) *CreateTransitRouterRouteTableResponse
	GetBody() *CreateTransitRouterRouteTableResponseBody
}

type CreateTransitRouterRouteTableResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterRouteTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterRouteTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterRouteTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterRouteTableResponse) GetBody() *CreateTransitRouterRouteTableResponseBody {
	return s.Body
}

func (s *CreateTransitRouterRouteTableResponse) SetHeaders(v map[string]*string) *CreateTransitRouterRouteTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterRouteTableResponse) SetStatusCode(v int32) *CreateTransitRouterRouteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterRouteTableResponse) SetBody(v *CreateTransitRouterRouteTableResponseBody) *CreateTransitRouterRouteTableResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterRouteTableResponse) Validate() error {
	return dara.Validate(s)
}
