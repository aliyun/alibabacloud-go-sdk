// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodUserTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodUserTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodUserTagsResponseBody) *DescribeVodUserTagsResponse
	GetBody() *DescribeVodUserTagsResponseBody
}

type DescribeVodUserTagsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodUserTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodUserTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodUserTagsResponse) GetBody() *DescribeVodUserTagsResponseBody {
	return s.Body
}

func (s *DescribeVodUserTagsResponse) SetHeaders(v map[string]*string) *DescribeVodUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserTagsResponse) SetStatusCode(v int32) *DescribeVodUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodUserTagsResponse) SetBody(v *DescribeVodUserTagsResponseBody) *DescribeVodUserTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeVodUserTagsResponse) Validate() error {
	return dara.Validate(s)
}
