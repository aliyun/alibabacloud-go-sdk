// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceIpAddressResponseBody) *DescribeInstanceIpAddressResponse
	GetBody() *DescribeInstanceIpAddressResponseBody
}

type DescribeInstanceIpAddressResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceIpAddressResponse) GetBody() *DescribeInstanceIpAddressResponseBody {
	return s.Body
}

func (s *DescribeInstanceIpAddressResponse) SetHeaders(v map[string]*string) *DescribeInstanceIpAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceIpAddressResponse) SetStatusCode(v int32) *DescribeInstanceIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceIpAddressResponse) SetBody(v *DescribeInstanceIpAddressResponseBody) *DescribeInstanceIpAddressResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceIpAddressResponse) Validate() error {
	return dara.Validate(s)
}
