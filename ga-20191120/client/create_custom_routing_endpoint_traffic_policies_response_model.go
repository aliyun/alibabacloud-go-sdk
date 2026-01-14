// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointTrafficPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointTrafficPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomRoutingEndpointTrafficPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) *CreateCustomRoutingEndpointTrafficPoliciesResponse
	GetBody() *CreateCustomRoutingEndpointTrafficPoliciesResponseBody
}

type CreateCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) GetBody() *CreateCustomRoutingEndpointTrafficPoliciesResponseBody {
	return s.Body
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *CreateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *CreateCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) *CreateCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
