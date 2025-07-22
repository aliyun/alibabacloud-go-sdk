// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisabledExpressConnectRouterRouteEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDisabledExpressConnectRouterRouteEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) *DescribeDisabledExpressConnectRouterRouteEntriesResponse
	GetBody() *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) GetBody() *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	return s.Body
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetBody(v *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) Validate() error {
	return dara.Validate(s)
}
