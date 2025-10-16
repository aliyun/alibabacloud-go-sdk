// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDriveUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDriveUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDriveUsersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDriveUsersResponseBody) *DescribeCloudDriveUsersResponse
	GetBody() *DescribeCloudDriveUsersResponseBody
}

type DescribeCloudDriveUsersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDriveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDriveUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDriveUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDriveUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDriveUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDriveUsersResponse) GetBody() *DescribeCloudDriveUsersResponseBody {
	return s.Body
}

func (s *DescribeCloudDriveUsersResponse) SetHeaders(v map[string]*string) *DescribeCloudDriveUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDriveUsersResponse) SetStatusCode(v int32) *DescribeCloudDriveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDriveUsersResponse) SetBody(v *DescribeCloudDriveUsersResponseBody) *DescribeCloudDriveUsersResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDriveUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
