// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodAppNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodAppNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodAppNameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodAppNameResponseBody) *DescribeVodAppNameResponse
	GetBody() *DescribeVodAppNameResponseBody
}

type DescribeVodAppNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodAppNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodAppNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodAppNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodAppNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodAppNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodAppNameResponse) GetBody() *DescribeVodAppNameResponseBody {
	return s.Body
}

func (s *DescribeVodAppNameResponse) SetHeaders(v map[string]*string) *DescribeVodAppNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodAppNameResponse) SetStatusCode(v int32) *DescribeVodAppNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodAppNameResponse) SetBody(v *DescribeVodAppNameResponseBody) *DescribeVodAppNameResponse {
	s.Body = v
	return s
}

func (s *DescribeVodAppNameResponse) Validate() error {
	return dara.Validate(s)
}
