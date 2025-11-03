// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcGatewayEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcGatewayEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcGatewayEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcGatewayEndpointsResponseBody) *ListVpcGatewayEndpointsResponse
	GetBody() *ListVpcGatewayEndpointsResponseBody
}

type ListVpcGatewayEndpointsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcGatewayEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcGatewayEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcGatewayEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcGatewayEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcGatewayEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcGatewayEndpointsResponse) GetBody() *ListVpcGatewayEndpointsResponseBody {
	return s.Body
}

func (s *ListVpcGatewayEndpointsResponse) SetHeaders(v map[string]*string) *ListVpcGatewayEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcGatewayEndpointsResponse) SetStatusCode(v int32) *ListVpcGatewayEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcGatewayEndpointsResponse) SetBody(v *ListVpcGatewayEndpointsResponseBody) *ListVpcGatewayEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListVpcGatewayEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
