// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShardingNetworkAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeShardingNetworkAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeShardingNetworkAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeShardingNetworkAddressResponseBody) *DescribeShardingNetworkAddressResponse
	GetBody() *DescribeShardingNetworkAddressResponseBody
}

type DescribeShardingNetworkAddressResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeShardingNetworkAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeShardingNetworkAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeShardingNetworkAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeShardingNetworkAddressResponse) GetBody() *DescribeShardingNetworkAddressResponseBody {
	return s.Body
}

func (s *DescribeShardingNetworkAddressResponse) SetHeaders(v map[string]*string) *DescribeShardingNetworkAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeShardingNetworkAddressResponse) SetStatusCode(v int32) *DescribeShardingNetworkAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponse) SetBody(v *DescribeShardingNetworkAddressResponseBody) *DescribeShardingNetworkAddressResponse {
	s.Body = v
	return s
}

func (s *DescribeShardingNetworkAddressResponse) Validate() error {
	return dara.Validate(s)
}
