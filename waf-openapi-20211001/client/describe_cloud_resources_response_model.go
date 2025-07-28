// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudResourcesResponseBody) *DescribeCloudResourcesResponse
	GetBody() *DescribeCloudResourcesResponseBody
}

type DescribeCloudResourcesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudResourcesResponse) GetBody() *DescribeCloudResourcesResponseBody {
	return s.Body
}

func (s *DescribeCloudResourcesResponse) SetHeaders(v map[string]*string) *DescribeCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudResourcesResponse) SetStatusCode(v int32) *DescribeCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudResourcesResponse) SetBody(v *DescribeCloudResourcesResponseBody) *DescribeCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudResourcesResponse) Validate() error {
	return dara.Validate(s)
}
