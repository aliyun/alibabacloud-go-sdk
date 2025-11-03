// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnGatewayResponseBody) *DescribeVpnGatewayResponse
	GetBody() *DescribeVpnGatewayResponseBody
}

type DescribeVpnGatewayResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnGatewayResponse) GetBody() *DescribeVpnGatewayResponseBody {
	return s.Body
}

func (s *DescribeVpnGatewayResponse) SetHeaders(v map[string]*string) *DescribeVpnGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnGatewayResponse) SetStatusCode(v int32) *DescribeVpnGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnGatewayResponse) SetBody(v *DescribeVpnGatewayResponseBody) *DescribeVpnGatewayResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
