// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6AddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpv6AddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpv6AddressesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpv6AddressesResponseBody) *DescribeIpv6AddressesResponse
	GetBody() *DescribeIpv6AddressesResponseBody
}

type DescribeIpv6AddressesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpv6AddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpv6AddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpv6AddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpv6AddressesResponse) GetBody() *DescribeIpv6AddressesResponseBody {
	return s.Body
}

func (s *DescribeIpv6AddressesResponse) SetHeaders(v map[string]*string) *DescribeIpv6AddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpv6AddressesResponse) SetStatusCode(v int32) *DescribeIpv6AddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpv6AddressesResponse) SetBody(v *DescribeIpv6AddressesResponseBody) *DescribeIpv6AddressesResponse {
	s.Body = v
	return s
}

func (s *DescribeIpv6AddressesResponse) Validate() error {
	return dara.Validate(s)
}
