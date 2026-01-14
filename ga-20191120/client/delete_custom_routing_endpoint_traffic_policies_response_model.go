// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointTrafficPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointTrafficPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomRoutingEndpointTrafficPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) *DeleteCustomRoutingEndpointTrafficPoliciesResponse
	GetBody() *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody
}

type DeleteCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) GetBody() *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody {
	return s.Body
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *DeleteCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *DeleteCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) *DeleteCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
