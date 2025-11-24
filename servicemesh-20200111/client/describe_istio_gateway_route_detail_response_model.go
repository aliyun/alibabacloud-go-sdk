// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIstioGatewayRouteDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIstioGatewayRouteDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIstioGatewayRouteDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIstioGatewayRouteDetailResponseBody) *DescribeIstioGatewayRouteDetailResponse
	GetBody() *DescribeIstioGatewayRouteDetailResponseBody
}

type DescribeIstioGatewayRouteDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIstioGatewayRouteDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIstioGatewayRouteDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIstioGatewayRouteDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeIstioGatewayRouteDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIstioGatewayRouteDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIstioGatewayRouteDetailResponse) GetBody() *DescribeIstioGatewayRouteDetailResponseBody {
	return s.Body
}

func (s *DescribeIstioGatewayRouteDetailResponse) SetHeaders(v map[string]*string) *DescribeIstioGatewayRouteDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponse) SetStatusCode(v int32) *DescribeIstioGatewayRouteDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponse) SetBody(v *DescribeIstioGatewayRouteDetailResponseBody) *DescribeIstioGatewayRouteDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeIstioGatewayRouteDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
