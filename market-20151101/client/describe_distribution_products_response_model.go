// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDistributionProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDistributionProductsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDistributionProductsResponseBody) *DescribeDistributionProductsResponse
	GetBody() *DescribeDistributionProductsResponseBody
}

type DescribeDistributionProductsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDistributionProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDistributionProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDistributionProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDistributionProductsResponse) GetBody() *DescribeDistributionProductsResponseBody {
	return s.Body
}

func (s *DescribeDistributionProductsResponse) SetHeaders(v map[string]*string) *DescribeDistributionProductsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistributionProductsResponse) SetStatusCode(v int32) *DescribeDistributionProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistributionProductsResponse) SetBody(v *DescribeDistributionProductsResponseBody) *DescribeDistributionProductsResponse {
	s.Body = v
	return s
}

func (s *DescribeDistributionProductsResponse) Validate() error {
	return dara.Validate(s)
}
