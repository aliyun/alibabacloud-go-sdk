// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpAddressesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpAddressesResponseBody) *DescribeIpAddressesResponse
	GetBody() *DescribeIpAddressesResponseBody
}

type DescribeIpAddressesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpAddressesResponse) GetBody() *DescribeIpAddressesResponseBody {
	return s.Body
}

func (s *DescribeIpAddressesResponse) SetHeaders(v map[string]*string) *DescribeIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpAddressesResponse) SetStatusCode(v int32) *DescribeIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpAddressesResponse) SetBody(v *DescribeIpAddressesResponseBody) *DescribeIpAddressesResponse {
	s.Body = v
	return s
}

func (s *DescribeIpAddressesResponse) Validate() error {
	return dara.Validate(s)
}
