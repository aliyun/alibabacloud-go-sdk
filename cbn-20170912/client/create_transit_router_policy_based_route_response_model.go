// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPolicyBasedRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterPolicyBasedRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterPolicyBasedRouteResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterPolicyBasedRouteResponseBody) *CreateTransitRouterPolicyBasedRouteResponse
	GetBody() *CreateTransitRouterPolicyBasedRouteResponseBody
}

type CreateTransitRouterPolicyBasedRouteResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterPolicyBasedRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterPolicyBasedRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPolicyBasedRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) GetBody() *CreateTransitRouterPolicyBasedRouteResponseBody {
	return s.Body
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) SetHeaders(v map[string]*string) *CreateTransitRouterPolicyBasedRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) SetStatusCode(v int32) *CreateTransitRouterPolicyBasedRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) SetBody(v *CreateTransitRouterPolicyBasedRouteResponseBody) *CreateTransitRouterPolicyBasedRouteResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterPolicyBasedRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
