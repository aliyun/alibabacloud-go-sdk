// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoDetailListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageRepoDetailListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageRepoDetailListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageRepoDetailListResponseBody) *DescribeImageRepoDetailListResponse
	GetBody() *DescribeImageRepoDetailListResponseBody
}

type DescribeImageRepoDetailListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageRepoDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageRepoDetailListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoDetailListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoDetailListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageRepoDetailListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageRepoDetailListResponse) GetBody() *DescribeImageRepoDetailListResponseBody {
	return s.Body
}

func (s *DescribeImageRepoDetailListResponse) SetHeaders(v map[string]*string) *DescribeImageRepoDetailListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageRepoDetailListResponse) SetStatusCode(v int32) *DescribeImageRepoDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageRepoDetailListResponse) SetBody(v *DescribeImageRepoDetailListResponseBody) *DescribeImageRepoDetailListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageRepoDetailListResponse) Validate() error {
	return dara.Validate(s)
}
