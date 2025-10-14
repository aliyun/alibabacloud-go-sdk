// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskAvailableResourceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDiskAvailableResourceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDiskAvailableResourceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDiskAvailableResourceInfoResponseBody) *DescribeCloudDiskAvailableResourceInfoResponse
	GetBody() *DescribeCloudDiskAvailableResourceInfoResponseBody
}

type DescribeCloudDiskAvailableResourceInfoResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDiskAvailableResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDiskAvailableResourceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskAvailableResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) GetBody() *DescribeCloudDiskAvailableResourceInfoResponseBody {
	return s.Body
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeCloudDiskAvailableResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) SetStatusCode(v int32) *DescribeCloudDiskAvailableResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) SetBody(v *DescribeCloudDiskAvailableResourceInfoResponseBody) *DescribeCloudDiskAvailableResourceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDiskAvailableResourceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
