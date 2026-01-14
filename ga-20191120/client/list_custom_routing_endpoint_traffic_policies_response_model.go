// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomRoutingEndpointTrafficPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomRoutingEndpointTrafficPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomRoutingEndpointTrafficPoliciesResponseBody) *ListCustomRoutingEndpointTrafficPoliciesResponse
	GetBody() *ListCustomRoutingEndpointTrafficPoliciesResponseBody
}

type ListCustomRoutingEndpointTrafficPoliciesResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomRoutingEndpointTrafficPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomRoutingEndpointTrafficPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) GetBody() *ListCustomRoutingEndpointTrafficPoliciesResponseBody {
	return s.Body
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) SetHeaders(v map[string]*string) *ListCustomRoutingEndpointTrafficPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) SetStatusCode(v int32) *ListCustomRoutingEndpointTrafficPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) SetBody(v *ListCustomRoutingEndpointTrafficPoliciesResponseBody) *ListCustomRoutingEndpointTrafficPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListCustomRoutingEndpointTrafficPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
