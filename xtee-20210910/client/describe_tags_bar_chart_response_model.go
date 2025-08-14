// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsBarChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsBarChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsBarChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsBarChartResponseBody) *DescribeTagsBarChartResponse
	GetBody() *DescribeTagsBarChartResponseBody
}

type DescribeTagsBarChartResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsBarChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsBarChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsBarChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsBarChartResponse) GetBody() *DescribeTagsBarChartResponseBody {
	return s.Body
}

func (s *DescribeTagsBarChartResponse) SetHeaders(v map[string]*string) *DescribeTagsBarChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsBarChartResponse) SetStatusCode(v int32) *DescribeTagsBarChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsBarChartResponse) SetBody(v *DescribeTagsBarChartResponseBody) *DescribeTagsBarChartResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsBarChartResponse) Validate() error {
	return dara.Validate(s)
}
