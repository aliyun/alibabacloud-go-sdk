// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterChildInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterChildInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterChildInstanceResponseBody) *DescribeExpressConnectRouterChildInstanceResponse
	GetBody() *DescribeExpressConnectRouterChildInstanceResponseBody
}

type DescribeExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) GetBody() *DescribeExpressConnectRouterChildInstanceResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetBody(v *DescribeExpressConnectRouterChildInstanceResponseBody) *DescribeExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) Validate() error {
	return dara.Validate(s)
}
