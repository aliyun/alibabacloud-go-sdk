// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributionProductsLinkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDistributionProductsLinkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDistributionProductsLinkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDistributionProductsLinkResponseBody) *DescribeDistributionProductsLinkResponse
	GetBody() *DescribeDistributionProductsLinkResponseBody
}

type DescribeDistributionProductsLinkResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDistributionProductsLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDistributionProductsLinkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributionProductsLinkResponse) GoString() string {
	return s.String()
}

func (s *DescribeDistributionProductsLinkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDistributionProductsLinkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDistributionProductsLinkResponse) GetBody() *DescribeDistributionProductsLinkResponseBody {
	return s.Body
}

func (s *DescribeDistributionProductsLinkResponse) SetHeaders(v map[string]*string) *DescribeDistributionProductsLinkResponse {
	s.Headers = v
	return s
}

func (s *DescribeDistributionProductsLinkResponse) SetStatusCode(v int32) *DescribeDistributionProductsLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDistributionProductsLinkResponse) SetBody(v *DescribeDistributionProductsLinkResponseBody) *DescribeDistributionProductsLinkResponse {
	s.Body = v
	return s
}

func (s *DescribeDistributionProductsLinkResponse) Validate() error {
	return dara.Validate(s)
}
