// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDriveGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDriveGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDriveGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDriveGroupsResponseBody) *DescribeCloudDriveGroupsResponse
	GetBody() *DescribeCloudDriveGroupsResponseBody
}

type DescribeCloudDriveGroupsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDriveGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDriveGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDriveGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDriveGroupsResponse) GetBody() *DescribeCloudDriveGroupsResponseBody {
	return s.Body
}

func (s *DescribeCloudDriveGroupsResponse) SetHeaders(v map[string]*string) *DescribeCloudDriveGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDriveGroupsResponse) SetStatusCode(v int32) *DescribeCloudDriveGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDriveGroupsResponse) SetBody(v *DescribeCloudDriveGroupsResponseBody) *DescribeCloudDriveGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDriveGroupsResponse) Validate() error {
	return dara.Validate(s)
}
