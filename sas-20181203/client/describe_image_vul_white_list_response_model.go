// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageVulWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageVulWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageVulWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageVulWhiteListResponseBody) *DescribeImageVulWhiteListResponse
	GetBody() *DescribeImageVulWhiteListResponseBody
}

type DescribeImageVulWhiteListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageVulWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageVulWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageVulWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageVulWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageVulWhiteListResponse) GetBody() *DescribeImageVulWhiteListResponseBody {
	return s.Body
}

func (s *DescribeImageVulWhiteListResponse) SetHeaders(v map[string]*string) *DescribeImageVulWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageVulWhiteListResponse) SetStatusCode(v int32) *DescribeImageVulWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageVulWhiteListResponse) SetBody(v *DescribeImageVulWhiteListResponseBody) *DescribeImageVulWhiteListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageVulWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
