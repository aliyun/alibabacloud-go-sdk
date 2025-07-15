// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDrivePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDrivePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDrivePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDrivePermissionsResponseBody) *DescribeCloudDrivePermissionsResponse
	GetBody() *DescribeCloudDrivePermissionsResponseBody
}

type DescribeCloudDrivePermissionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDrivePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDrivePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDrivePermissionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDrivePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDrivePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDrivePermissionsResponse) GetBody() *DescribeCloudDrivePermissionsResponseBody {
	return s.Body
}

func (s *DescribeCloudDrivePermissionsResponse) SetHeaders(v map[string]*string) *DescribeCloudDrivePermissionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDrivePermissionsResponse) SetStatusCode(v int32) *DescribeCloudDrivePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDrivePermissionsResponse) SetBody(v *DescribeCloudDrivePermissionsResponseBody) *DescribeCloudDrivePermissionsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDrivePermissionsResponse) Validate() error {
	return dara.Validate(s)
}
