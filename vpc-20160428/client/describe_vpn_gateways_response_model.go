// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnGatewaysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnGatewaysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnGatewaysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnGatewaysResponseBody) *DescribeVpnGatewaysResponse
	GetBody() *DescribeVpnGatewaysResponseBody
}

type DescribeVpnGatewaysResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnGatewaysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnGatewaysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnGatewaysResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnGatewaysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnGatewaysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnGatewaysResponse) GetBody() *DescribeVpnGatewaysResponseBody {
	return s.Body
}

func (s *DescribeVpnGatewaysResponse) SetHeaders(v map[string]*string) *DescribeVpnGatewaysResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnGatewaysResponse) SetStatusCode(v int32) *DescribeVpnGatewaysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnGatewaysResponse) SetBody(v *DescribeVpnGatewaysResponseBody) *DescribeVpnGatewaysResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnGatewaysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
