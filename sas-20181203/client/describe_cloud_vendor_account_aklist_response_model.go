// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorAccountAKListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudVendorAccountAKListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudVendorAccountAKListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudVendorAccountAKListResponseBody) *DescribeCloudVendorAccountAKListResponse
	GetBody() *DescribeCloudVendorAccountAKListResponseBody
}

type DescribeCloudVendorAccountAKListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudVendorAccountAKListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudVendorAccountAKListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorAccountAKListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorAccountAKListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudVendorAccountAKListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudVendorAccountAKListResponse) GetBody() *DescribeCloudVendorAccountAKListResponseBody {
	return s.Body
}

func (s *DescribeCloudVendorAccountAKListResponse) SetHeaders(v map[string]*string) *DescribeCloudVendorAccountAKListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponse) SetStatusCode(v int32) *DescribeCloudVendorAccountAKListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponse) SetBody(v *DescribeCloudVendorAccountAKListResponseBody) *DescribeCloudVendorAccountAKListResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudVendorAccountAKListResponse) Validate() error {
	return dara.Validate(s)
}
