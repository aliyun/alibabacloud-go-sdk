// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterRouteEntriesResponseBody) *DescribeExpressConnectRouterRouteEntriesResponse
	GetBody() *DescribeExpressConnectRouterRouteEntriesResponseBody
}

type DescribeExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) GetBody() *DescribeExpressConnectRouterRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetBody(v *DescribeExpressConnectRouterRouteEntriesResponseBody) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
