// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterRouteTableAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTransitRouterRouteTableAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTransitRouterRouteTableAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *ListTransitRouterRouteTableAssociationsResponseBody) *ListTransitRouterRouteTableAssociationsResponse
	GetBody() *ListTransitRouterRouteTableAssociationsResponseBody
}

type ListTransitRouterRouteTableAssociationsResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTransitRouterRouteTableAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTransitRouterRouteTableAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterRouteTableAssociationsResponse) GoString() string {
	return s.String()
}

func (s *ListTransitRouterRouteTableAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTransitRouterRouteTableAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTransitRouterRouteTableAssociationsResponse) GetBody() *ListTransitRouterRouteTableAssociationsResponseBody {
	return s.Body
}

func (s *ListTransitRouterRouteTableAssociationsResponse) SetHeaders(v map[string]*string) *ListTransitRouterRouteTableAssociationsResponse {
	s.Headers = v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponse) SetStatusCode(v int32) *ListTransitRouterRouteTableAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponse) SetBody(v *ListTransitRouterRouteTableAssociationsResponseBody) *ListTransitRouterRouteTableAssociationsResponse {
	s.Body = v
	return s
}

func (s *ListTransitRouterRouteTableAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
