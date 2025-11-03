// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskGroupDrivesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDiskGroupDrivesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDiskGroupDrivesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDiskGroupDrivesResponseBody) *DescribeCloudDiskGroupDrivesResponse
	GetBody() *DescribeCloudDiskGroupDrivesResponseBody
}

type DescribeCloudDiskGroupDrivesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDiskGroupDrivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDiskGroupDrivesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskGroupDrivesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskGroupDrivesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDiskGroupDrivesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDiskGroupDrivesResponse) GetBody() *DescribeCloudDiskGroupDrivesResponseBody {
	return s.Body
}

func (s *DescribeCloudDiskGroupDrivesResponse) SetHeaders(v map[string]*string) *DescribeCloudDiskGroupDrivesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponse) SetStatusCode(v int32) *DescribeCloudDiskGroupDrivesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponse) SetBody(v *DescribeCloudDiskGroupDrivesResponseBody) *DescribeCloudDiskGroupDrivesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDiskGroupDrivesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
