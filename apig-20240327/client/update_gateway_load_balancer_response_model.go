// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayLoadBalancerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGatewayLoadBalancerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGatewayLoadBalancerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGatewayLoadBalancerResponseBody) *UpdateGatewayLoadBalancerResponse
	GetBody() *UpdateGatewayLoadBalancerResponseBody
}

type UpdateGatewayLoadBalancerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGatewayLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGatewayLoadBalancerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGatewayLoadBalancerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGatewayLoadBalancerResponse) GetBody() *UpdateGatewayLoadBalancerResponseBody {
	return s.Body
}

func (s *UpdateGatewayLoadBalancerResponse) SetHeaders(v map[string]*string) *UpdateGatewayLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *UpdateGatewayLoadBalancerResponse) SetStatusCode(v int32) *UpdateGatewayLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGatewayLoadBalancerResponse) SetBody(v *UpdateGatewayLoadBalancerResponseBody) *UpdateGatewayLoadBalancerResponse {
	s.Body = v
	return s
}

func (s *UpdateGatewayLoadBalancerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
