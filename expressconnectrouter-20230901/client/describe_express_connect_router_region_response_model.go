// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterRegionResponseBody) *DescribeExpressConnectRouterRegionResponse
	GetBody() *DescribeExpressConnectRouterRegionResponseBody
}

type DescribeExpressConnectRouterRegionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterRegionResponse) GetBody() *DescribeExpressConnectRouterRegionResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterRegionResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) SetBody(v *DescribeExpressConnectRouterRegionResponseBody) *DescribeExpressConnectRouterRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) Validate() error {
	return dara.Validate(s)
}
