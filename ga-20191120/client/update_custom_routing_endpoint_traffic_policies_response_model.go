// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointTrafficPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointTrafficPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) *UpdateCustomRoutingEndpointTrafficPoliciesResponse
	GetBody() *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody
}

type UpdateCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) GetBody() *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody {
	return s.Body
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) *UpdateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
