// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6GatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpv6GatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpv6GatewaysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpv6GatewaysResponseBody) *DescribeIpv6GatewaysResponse
	GetBody() *DescribeIpv6GatewaysResponseBody
}

type DescribeIpv6GatewaysResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpv6GatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpv6GatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpv6GatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpv6GatewaysResponse) GetBody() *DescribeIpv6GatewaysResponseBody {
	return s.Body
}

func (s *DescribeIpv6GatewaysResponse) SetHeaders(v map[string]*string) *DescribeIpv6GatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpv6GatewaysResponse) SetStatusCode(v int32) *DescribeIpv6GatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpv6GatewaysResponse) SetBody(v *DescribeIpv6GatewaysResponseBody) *DescribeIpv6GatewaysResponse {
	s.Body = v
	return s
}

func (s *DescribeIpv6GatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
