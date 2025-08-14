// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsListResponseBody) *DescribeTagsListResponse
	GetBody() *DescribeTagsListResponseBody
}

type DescribeTagsListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsListResponse) GetBody() *DescribeTagsListResponseBody {
	return s.Body
}

func (s *DescribeTagsListResponse) SetHeaders(v map[string]*string) *DescribeTagsListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsListResponse) SetStatusCode(v int32) *DescribeTagsListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsListResponse) SetBody(v *DescribeTagsListResponseBody) *DescribeTagsListResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsListResponse) Validate() error {
	return dara.Validate(s)
}
