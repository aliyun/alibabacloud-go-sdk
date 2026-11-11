// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPolicyBasedRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTransitRouterPolicyBasedRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTransitRouterPolicyBasedRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTransitRouterPolicyBasedRouteResponseBody) *UpdateTransitRouterPolicyBasedRouteResponse
	GetBody() *UpdateTransitRouterPolicyBasedRouteResponseBody
}

type UpdateTransitRouterPolicyBasedRouteResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTransitRouterPolicyBasedRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTransitRouterPolicyBasedRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPolicyBasedRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) GetBody() *UpdateTransitRouterPolicyBasedRouteResponseBody {
	return s.Body
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) SetHeaders(v map[string]*string) *UpdateTransitRouterPolicyBasedRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) SetStatusCode(v int32) *UpdateTransitRouterPolicyBasedRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) SetBody(v *UpdateTransitRouterPolicyBasedRouteResponseBody) *UpdateTransitRouterPolicyBasedRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateTransitRouterPolicyBasedRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
