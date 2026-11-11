// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyBasedRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTransitRouterPolicyBasedRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTransitRouterPolicyBasedRouteResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTransitRouterPolicyBasedRouteResponseBody) *DeleteTransitRouterPolicyBasedRouteResponse
	GetBody() *DeleteTransitRouterPolicyBasedRouteResponseBody
}

type DeleteTransitRouterPolicyBasedRouteResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTransitRouterPolicyBasedRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTransitRouterPolicyBasedRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyBasedRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) GetBody() *DeleteTransitRouterPolicyBasedRouteResponseBody {
	return s.Body
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) SetHeaders(v map[string]*string) *DeleteTransitRouterPolicyBasedRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) SetStatusCode(v int32) *DeleteTransitRouterPolicyBasedRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) SetBody(v *DeleteTransitRouterPolicyBasedRouteResponseBody) *DeleteTransitRouterPolicyBasedRouteResponse {
	s.Body = v
	return s
}

func (s *DeleteTransitRouterPolicyBasedRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
