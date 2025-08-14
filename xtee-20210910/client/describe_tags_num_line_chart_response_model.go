// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsNumLineChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsNumLineChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsNumLineChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsNumLineChartResponseBody) *DescribeTagsNumLineChartResponse
	GetBody() *DescribeTagsNumLineChartResponseBody
}

type DescribeTagsNumLineChartResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsNumLineChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsNumLineChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsNumLineChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsNumLineChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsNumLineChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsNumLineChartResponse) GetBody() *DescribeTagsNumLineChartResponseBody {
	return s.Body
}

func (s *DescribeTagsNumLineChartResponse) SetHeaders(v map[string]*string) *DescribeTagsNumLineChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsNumLineChartResponse) SetStatusCode(v int32) *DescribeTagsNumLineChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsNumLineChartResponse) SetBody(v *DescribeTagsNumLineChartResponseBody) *DescribeTagsNumLineChartResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsNumLineChartResponse) Validate() error {
	return dara.Validate(s)
}
