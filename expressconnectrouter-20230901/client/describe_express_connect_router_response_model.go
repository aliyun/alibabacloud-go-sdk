// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpressConnectRouterResponseBody) *DescribeExpressConnectRouterResponse
	GetBody() *DescribeExpressConnectRouterResponseBody
}

type DescribeExpressConnectRouterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpressConnectRouterResponse) GetBody() *DescribeExpressConnectRouterResponseBody {
	return s.Body
}

func (s *DescribeExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponse) SetBody(v *DescribeExpressConnectRouterResponseBody) *DescribeExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *DescribeExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
