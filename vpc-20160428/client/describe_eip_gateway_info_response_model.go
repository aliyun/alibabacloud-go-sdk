// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipGatewayInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEipGatewayInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEipGatewayInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEipGatewayInfoResponseBody) *DescribeEipGatewayInfoResponse
	GetBody() *DescribeEipGatewayInfoResponseBody
}

type DescribeEipGatewayInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEipGatewayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEipGatewayInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipGatewayInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEipGatewayInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEipGatewayInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEipGatewayInfoResponse) GetBody() *DescribeEipGatewayInfoResponseBody {
	return s.Body
}

func (s *DescribeEipGatewayInfoResponse) SetHeaders(v map[string]*string) *DescribeEipGatewayInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEipGatewayInfoResponse) SetStatusCode(v int32) *DescribeEipGatewayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEipGatewayInfoResponse) SetBody(v *DescribeEipGatewayInfoResponseBody) *DescribeEipGatewayInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeEipGatewayInfoResponse) Validate() error {
	return dara.Validate(s)
}
