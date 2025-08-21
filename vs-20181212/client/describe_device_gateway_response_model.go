// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeviceGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeviceGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeviceGatewayResponseBody) *DescribeDeviceGatewayResponse
	GetBody() *DescribeDeviceGatewayResponseBody
}

type DescribeDeviceGatewayResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeviceGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeviceGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceGatewayResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeviceGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeviceGatewayResponse) GetBody() *DescribeDeviceGatewayResponseBody {
	return s.Body
}

func (s *DescribeDeviceGatewayResponse) SetHeaders(v map[string]*string) *DescribeDeviceGatewayResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceGatewayResponse) SetStatusCode(v int32) *DescribeDeviceGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeviceGatewayResponse) SetBody(v *DescribeDeviceGatewayResponseBody) *DescribeDeviceGatewayResponse {
	s.Body = v
	return s
}

func (s *DescribeDeviceGatewayResponse) Validate() error {
	return dara.Validate(s)
}
