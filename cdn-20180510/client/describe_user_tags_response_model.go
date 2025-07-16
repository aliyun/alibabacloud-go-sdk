// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserTagsResponseBody) *DescribeUserTagsResponse
	GetBody() *DescribeUserTagsResponseBody
}

type DescribeUserTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserTagsResponse) GetBody() *DescribeUserTagsResponseBody {
	return s.Body
}

func (s *DescribeUserTagsResponse) SetHeaders(v map[string]*string) *DescribeUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTagsResponse) SetStatusCode(v int32) *DescribeUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserTagsResponse) SetBody(v *DescribeUserTagsResponseBody) *DescribeUserTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserTagsResponse) Validate() error {
	return dara.Validate(s)
}
