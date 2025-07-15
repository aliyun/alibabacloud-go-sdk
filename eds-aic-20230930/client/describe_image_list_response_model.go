// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageListResponseBody) *DescribeImageListResponse
	GetBody() *DescribeImageListResponseBody
}

type DescribeImageListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageListResponse) GetBody() *DescribeImageListResponseBody {
	return s.Body
}

func (s *DescribeImageListResponse) SetHeaders(v map[string]*string) *DescribeImageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageListResponse) SetStatusCode(v int32) *DescribeImageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageListResponse) SetBody(v *DescribeImageListResponseBody) *DescribeImageListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageListResponse) Validate() error {
	return dara.Validate(s)
}
