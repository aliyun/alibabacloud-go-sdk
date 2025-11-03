// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomerGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomerGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomerGatewayResponseBody) *DescribeCustomerGatewayResponse
	GetBody() *DescribeCustomerGatewayResponseBody
}

type DescribeCustomerGatewayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomerGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomerGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomerGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomerGatewayResponse) GetBody() *DescribeCustomerGatewayResponseBody {
	return s.Body
}

func (s *DescribeCustomerGatewayResponse) SetHeaders(v map[string]*string) *DescribeCustomerGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomerGatewayResponse) SetStatusCode(v int32) *DescribeCustomerGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomerGatewayResponse) SetBody(v *DescribeCustomerGatewayResponseBody) *DescribeCustomerGatewayResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomerGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
