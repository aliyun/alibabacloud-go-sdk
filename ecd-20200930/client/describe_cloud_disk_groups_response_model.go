// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDiskGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDiskGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDiskGroupsResponseBody) *DescribeCloudDiskGroupsResponse
	GetBody() *DescribeCloudDiskGroupsResponseBody
}

type DescribeCloudDiskGroupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDiskGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDiskGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDiskGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDiskGroupsResponse) GetBody() *DescribeCloudDiskGroupsResponseBody {
	return s.Body
}

func (s *DescribeCloudDiskGroupsResponse) SetHeaders(v map[string]*string) *DescribeCloudDiskGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDiskGroupsResponse) SetStatusCode(v int32) *DescribeCloudDiskGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDiskGroupsResponse) SetBody(v *DescribeCloudDiskGroupsResponseBody) *DescribeCloudDiskGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDiskGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
