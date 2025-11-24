// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIstioGatewayRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIstioGatewayRoutesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIstioGatewayRoutesResponseBody) *DescribeIstioGatewayRoutesResponse
	GetBody() *DescribeIstioGatewayRoutesResponseBody
}

type DescribeIstioGatewayRoutesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIstioGatewayRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIstioGatewayRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRoutesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIstioGatewayRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIstioGatewayRoutesResponse) GetBody() *DescribeIstioGatewayRoutesResponseBody {
	return s.Body
}

func (s *DescribeIstioGatewayRoutesResponse) SetHeaders(v map[string]*string) *DescribeIstioGatewayRoutesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponse) SetStatusCode(v int32) *DescribeIstioGatewayRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIstioGatewayRoutesResponse) SetBody(v *DescribeIstioGatewayRoutesResponseBody) *DescribeIstioGatewayRoutesResponse {
	s.Body = v
	return s
}

func (s *DescribeIstioGatewayRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
