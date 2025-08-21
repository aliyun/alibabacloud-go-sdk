// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudDiskTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudDiskTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudDiskTypesResponseBody) *DescribeCloudDiskTypesResponse
	GetBody() *DescribeCloudDiskTypesResponseBody
}

type DescribeCloudDiskTypesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudDiskTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudDiskTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudDiskTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudDiskTypesResponse) GetBody() *DescribeCloudDiskTypesResponseBody {
	return s.Body
}

func (s *DescribeCloudDiskTypesResponse) SetHeaders(v map[string]*string) *DescribeCloudDiskTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudDiskTypesResponse) SetStatusCode(v int32) *DescribeCloudDiskTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudDiskTypesResponse) SetBody(v *DescribeCloudDiskTypesResponseBody) *DescribeCloudDiskTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudDiskTypesResponse) Validate() error {
	return dara.Validate(s)
}
