// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodTagResourcesResponseBody) *DescribeVodTagResourcesResponse
	GetBody() *DescribeVodTagResourcesResponseBody
}

type DescribeVodTagResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodTagResourcesResponse) GetBody() *DescribeVodTagResourcesResponseBody {
	return s.Body
}

func (s *DescribeVodTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeVodTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTagResourcesResponse) SetStatusCode(v int32) *DescribeVodTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodTagResourcesResponse) SetBody(v *DescribeVodTagResourcesResponseBody) *DescribeVodTagResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeVodTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}
