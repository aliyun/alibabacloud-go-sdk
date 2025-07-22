// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterInterRegionTransitModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterInterRegionTransitModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) *DescribeExpressConnectRouterInterRegionTransitModeResponse
	GetBody() *DescribeExpressConnectRouterInterRegionTransitModeResponseBody
}

type DescribeExpressConnectRouterInterRegionTransitModeResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterInterRegionTransitModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) GetBody() *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetBody(v *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) Validate() error {
	return dara.Validate(s)
}
