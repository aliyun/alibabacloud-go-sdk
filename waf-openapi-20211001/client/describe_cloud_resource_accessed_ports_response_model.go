// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceAccessedPortsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudResourceAccessedPortsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudResourceAccessedPortsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudResourceAccessedPortsResponseBody) *DescribeCloudResourceAccessedPortsResponse
	GetBody() *DescribeCloudResourceAccessedPortsResponseBody
}

type DescribeCloudResourceAccessedPortsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourceAccessedPortsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourceAccessedPortsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessedPortsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessedPortsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudResourceAccessedPortsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudResourceAccessedPortsResponse) GetBody() *DescribeCloudResourceAccessedPortsResponseBody {
	return s.Body
}

func (s *DescribeCloudResourceAccessedPortsResponse) SetHeaders(v map[string]*string) *DescribeCloudResourceAccessedPortsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourceAccessedPortsResponse) SetStatusCode(v int32) *DescribeCloudResourceAccessedPortsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourceAccessedPortsResponse) SetBody(v *DescribeCloudResourceAccessedPortsResponseBody) *DescribeCloudResourceAccessedPortsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudResourceAccessedPortsResponse) Validate() error {
	return dara.Validate(s)
}
