// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterEndpointsResponseBody) *DescribeDBClusterEndpointsResponse
	GetBody() *DescribeDBClusterEndpointsResponseBody
}

type DescribeDBClusterEndpointsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterEndpointsResponse) GetBody() *DescribeDBClusterEndpointsResponseBody {
	return s.Body
}

func (s *DescribeDBClusterEndpointsResponse) SetHeaders(v map[string]*string) *DescribeDBClusterEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterEndpointsResponse) SetStatusCode(v int32) *DescribeDBClusterEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponse) SetBody(v *DescribeDBClusterEndpointsResponseBody) *DescribeDBClusterEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterEndpointsResponse) Validate() error {
	return dara.Validate(s)
}
