// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGrantedToExpressConnectRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceGrantedToExpressConnectRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceGrantedToExpressConnectRouterResponseBody) *DescribeInstanceGrantedToExpressConnectRouterResponse
	GetBody() *DescribeInstanceGrantedToExpressConnectRouterResponseBody
}

type DescribeInstanceGrantedToExpressConnectRouterResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceGrantedToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) GetBody() *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	return s.Body
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetBody(v *DescribeInstanceGrantedToExpressConnectRouterResponseBody) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) Validate() error {
	return dara.Validate(s)
}
