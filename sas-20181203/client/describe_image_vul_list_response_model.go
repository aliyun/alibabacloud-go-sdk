// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageVulListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageVulListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageVulListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageVulListResponseBody) *DescribeImageVulListResponse
	GetBody() *DescribeImageVulListResponseBody
}

type DescribeImageVulListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageVulListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageVulListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageVulListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageVulListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageVulListResponse) GetBody() *DescribeImageVulListResponseBody {
	return s.Body
}

func (s *DescribeImageVulListResponse) SetHeaders(v map[string]*string) *DescribeImageVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageVulListResponse) SetStatusCode(v int32) *DescribeImageVulListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageVulListResponse) SetBody(v *DescribeImageVulListResponseBody) *DescribeImageVulListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageVulListResponse) Validate() error {
	return dara.Validate(s)
}
