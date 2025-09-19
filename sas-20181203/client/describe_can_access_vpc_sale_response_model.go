// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCanAccessVpcSaleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCanAccessVpcSaleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCanAccessVpcSaleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCanAccessVpcSaleResponseBody) *DescribeCanAccessVpcSaleResponse
	GetBody() *DescribeCanAccessVpcSaleResponseBody
}

type DescribeCanAccessVpcSaleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCanAccessVpcSaleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCanAccessVpcSaleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanAccessVpcSaleResponse) GoString() string {
	return s.String()
}

func (s *DescribeCanAccessVpcSaleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCanAccessVpcSaleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCanAccessVpcSaleResponse) GetBody() *DescribeCanAccessVpcSaleResponseBody {
	return s.Body
}

func (s *DescribeCanAccessVpcSaleResponse) SetHeaders(v map[string]*string) *DescribeCanAccessVpcSaleResponse {
	s.Headers = v
	return s
}

func (s *DescribeCanAccessVpcSaleResponse) SetStatusCode(v int32) *DescribeCanAccessVpcSaleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCanAccessVpcSaleResponse) SetBody(v *DescribeCanAccessVpcSaleResponseBody) *DescribeCanAccessVpcSaleResponse {
	s.Body = v
	return s
}

func (s *DescribeCanAccessVpcSaleResponse) Validate() error {
	return dara.Validate(s)
}
