// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewayAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnGatewayAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnGatewayAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnGatewayAvailableZonesResponseBody) *DescribeVpnGatewayAvailableZonesResponse
	GetBody() *DescribeVpnGatewayAvailableZonesResponseBody
}

type DescribeVpnGatewayAvailableZonesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnGatewayAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnGatewayAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnGatewayAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnGatewayAvailableZonesResponse) GetBody() *DescribeVpnGatewayAvailableZonesResponseBody {
	return s.Body
}

func (s *DescribeVpnGatewayAvailableZonesResponse) SetHeaders(v map[string]*string) *DescribeVpnGatewayAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponse) SetStatusCode(v int32) *DescribeVpnGatewayAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponse) SetBody(v *DescribeVpnGatewayAvailableZonesResponseBody) *DescribeVpnGatewayAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnGatewayAvailableZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
