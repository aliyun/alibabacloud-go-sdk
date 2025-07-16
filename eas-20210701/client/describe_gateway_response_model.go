// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGatewayResponseBody) *DescribeGatewayResponse
	GetBody() *DescribeGatewayResponseBody
}

type DescribeGatewayResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGatewayResponse) GetBody() *DescribeGatewayResponseBody {
	return s.Body
}

func (s *DescribeGatewayResponse) SetHeaders(v map[string]*string) *DescribeGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeGatewayResponse) SetStatusCode(v int32) *DescribeGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGatewayResponse) SetBody(v *DescribeGatewayResponseBody) *DescribeGatewayResponse {
	s.Body = v
	return s
}

func (s *DescribeGatewayResponse) Validate() error {
	return dara.Validate(s)
}
