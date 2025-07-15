// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomerGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomerGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomerGatewaysResponseBody) *DescribeCustomerGatewaysResponse
	GetBody() *DescribeCustomerGatewaysResponseBody
}

type DescribeCustomerGatewaysResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomerGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomerGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomerGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomerGatewaysResponse) GetBody() *DescribeCustomerGatewaysResponseBody {
	return s.Body
}

func (s *DescribeCustomerGatewaysResponse) SetHeaders(v map[string]*string) *DescribeCustomerGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomerGatewaysResponse) SetStatusCode(v int32) *DescribeCustomerGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomerGatewaysResponse) SetBody(v *DescribeCustomerGatewaysResponseBody) *DescribeCustomerGatewaysResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomerGatewaysResponse) Validate() error {
	return dara.Validate(s)
}
