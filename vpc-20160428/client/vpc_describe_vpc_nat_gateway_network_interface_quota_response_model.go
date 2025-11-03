// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse
	GetStatusCode() *int32
	SetBody(v *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse
	GetBody() *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody
}

type VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) GoString() string {
	return s.String()
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) GetBody() *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody {
	return s.Body
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) SetHeaders(v map[string]*string) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse {
	s.Headers = v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) SetStatusCode(v int32) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) SetBody(v *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponseBody) *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse {
	s.Body = v
	return s
}

func (s *VpcDescribeVpcNatGatewayNetworkInterfaceQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
