// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionRatioLineChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScoreSectionRatioLineChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScoreSectionRatioLineChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScoreSectionRatioLineChartResponseBody) *DescribeScoreSectionRatioLineChartResponse
	GetBody() *DescribeScoreSectionRatioLineChartResponseBody
}

type DescribeScoreSectionRatioLineChartResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScoreSectionRatioLineChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScoreSectionRatioLineChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionRatioLineChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionRatioLineChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScoreSectionRatioLineChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScoreSectionRatioLineChartResponse) GetBody() *DescribeScoreSectionRatioLineChartResponseBody {
	return s.Body
}

func (s *DescribeScoreSectionRatioLineChartResponse) SetHeaders(v map[string]*string) *DescribeScoreSectionRatioLineChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponse) SetStatusCode(v int32) *DescribeScoreSectionRatioLineChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponse) SetBody(v *DescribeScoreSectionRatioLineChartResponseBody) *DescribeScoreSectionRatioLineChartResponse {
	s.Body = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponse) Validate() error {
	return dara.Validate(s)
}
