// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudCenterInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudCenterInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudCenterInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudCenterInstancesResponseBody) *DescribeCloudCenterInstancesResponse
	GetBody() *DescribeCloudCenterInstancesResponseBody
}

type DescribeCloudCenterInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudCenterInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudCenterInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudCenterInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudCenterInstancesResponse) GetBody() *DescribeCloudCenterInstancesResponseBody {
	return s.Body
}

func (s *DescribeCloudCenterInstancesResponse) SetHeaders(v map[string]*string) *DescribeCloudCenterInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudCenterInstancesResponse) SetStatusCode(v int32) *DescribeCloudCenterInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudCenterInstancesResponse) SetBody(v *DescribeCloudCenterInstancesResponseBody) *DescribeCloudCenterInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudCenterInstancesResponse) Validate() error {
	return dara.Validate(s)
}
