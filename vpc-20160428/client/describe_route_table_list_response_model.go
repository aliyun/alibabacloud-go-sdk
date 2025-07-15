// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouteTableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouteTableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouteTableListResponseBody) *DescribeRouteTableListResponse
	GetBody() *DescribeRouteTableListResponseBody
}

type DescribeRouteTableListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouteTableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouteTableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteTableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouteTableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouteTableListResponse) GetBody() *DescribeRouteTableListResponseBody {
	return s.Body
}

func (s *DescribeRouteTableListResponse) SetHeaders(v map[string]*string) *DescribeRouteTableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteTableListResponse) SetStatusCode(v int32) *DescribeRouteTableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouteTableListResponse) SetBody(v *DescribeRouteTableListResponseBody) *DescribeRouteTableListResponse {
	s.Body = v
	return s
}

func (s *DescribeRouteTableListResponse) Validate() error {
	return dara.Validate(s)
}
