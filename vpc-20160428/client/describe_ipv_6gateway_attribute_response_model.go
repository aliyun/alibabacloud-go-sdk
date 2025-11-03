// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6GatewayAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpv6GatewayAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpv6GatewayAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpv6GatewayAttributeResponseBody) *DescribeIpv6GatewayAttributeResponse
	GetBody() *DescribeIpv6GatewayAttributeResponseBody
}

type DescribeIpv6GatewayAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpv6GatewayAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpv6GatewayAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewayAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewayAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpv6GatewayAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpv6GatewayAttributeResponse) GetBody() *DescribeIpv6GatewayAttributeResponseBody {
	return s.Body
}

func (s *DescribeIpv6GatewayAttributeResponse) SetHeaders(v map[string]*string) *DescribeIpv6GatewayAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponse) SetStatusCode(v int32) *DescribeIpv6GatewayAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponse) SetBody(v *DescribeIpv6GatewayAttributeResponseBody) *DescribeIpv6GatewayAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
