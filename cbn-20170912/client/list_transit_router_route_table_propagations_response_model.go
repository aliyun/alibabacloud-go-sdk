// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTablePropagationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterRouteTablePropagationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterRouteTablePropagationsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterRouteTablePropagationsResponseBody) *ListTransitRouterRouteTablePropagationsResponse
	GetBody() *ListTransitRouterRouteTablePropagationsResponseBody
}

type ListTransitRouterRouteTablePropagationsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterRouteTablePropagationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterRouteTablePropagationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTablePropagationsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTablePropagationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterRouteTablePropagationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterRouteTablePropagationsResponse) GetBody() *ListTransitRouterRouteTablePropagationsResponseBody {
	return s.Body
}

func (s *ListTransitRouterRouteTablePropagationsResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteTablePropagationsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponse) SetStatusCode(v int32) *ListTransitRouterRouteTablePropagationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponse) SetBody(v *ListTransitRouterRouteTablePropagationsResponseBody) *ListTransitRouterRouteTablePropagationsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterRouteTablePropagationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
