// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewayAssociateNetworkInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNatGatewayAssociateNetworkInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) *DescribeNatGatewayAssociateNetworkInterfacesResponse
	GetBody() *DescribeNatGatewayAssociateNetworkInterfacesResponseBody
}

type DescribeNatGatewayAssociateNetworkInterfacesResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatGatewayAssociateNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewayAssociateNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) GetBody() *DescribeNatGatewayAssociateNetworkInterfacesResponseBody {
	return s.Body
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) SetHeaders(v map[string]*string) *DescribeNatGatewayAssociateNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) SetStatusCode(v int32) *DescribeNatGatewayAssociateNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) SetBody(v *DescribeNatGatewayAssociateNetworkInterfacesResponseBody) *DescribeNatGatewayAssociateNetworkInterfacesResponse {
	s.Body = v
	return s
}

func (s *DescribeNatGatewayAssociateNetworkInterfacesResponse) Validate() error {
	return dara.Validate(s)
}
