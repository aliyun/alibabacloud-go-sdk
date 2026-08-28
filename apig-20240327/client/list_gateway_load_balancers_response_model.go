// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayLoadBalancersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayLoadBalancersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayLoadBalancersResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayLoadBalancersResponseBody) *ListGatewayLoadBalancersResponse
	GetBody() *ListGatewayLoadBalancersResponseBody
}

type ListGatewayLoadBalancersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayLoadBalancersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayLoadBalancersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayLoadBalancersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayLoadBalancersResponse) GetBody() *ListGatewayLoadBalancersResponseBody {
	return s.Body
}

func (s *ListGatewayLoadBalancersResponse) SetHeaders(v map[string]*string) *ListGatewayLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayLoadBalancersResponse) SetStatusCode(v int32) *ListGatewayLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayLoadBalancersResponse) SetBody(v *ListGatewayLoadBalancersResponseBody) *ListGatewayLoadBalancersResponse {
	s.Body = v
	return s
}

func (s *ListGatewayLoadBalancersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
