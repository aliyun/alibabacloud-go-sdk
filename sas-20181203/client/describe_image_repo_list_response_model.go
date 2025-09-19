// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRepoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageRepoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageRepoListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageRepoListResponseBody) *DescribeImageRepoListResponse
	GetBody() *DescribeImageRepoListResponseBody
}

type DescribeImageRepoListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageRepoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageRepoListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRepoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageRepoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageRepoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageRepoListResponse) GetBody() *DescribeImageRepoListResponseBody {
	return s.Body
}

func (s *DescribeImageRepoListResponse) SetHeaders(v map[string]*string) *DescribeImageRepoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageRepoListResponse) SetStatusCode(v int32) *DescribeImageRepoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageRepoListResponse) SetBody(v *DescribeImageRepoListResponseBody) *DescribeImageRepoListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageRepoListResponse) Validate() error {
	return dara.Validate(s)
}
