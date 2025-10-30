// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionNumLineChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScoreSectionNumLineChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScoreSectionNumLineChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScoreSectionNumLineChartResponseBody) *DescribeScoreSectionNumLineChartResponse
	GetBody() *DescribeScoreSectionNumLineChartResponseBody
}

type DescribeScoreSectionNumLineChartResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScoreSectionNumLineChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScoreSectionNumLineChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionNumLineChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionNumLineChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScoreSectionNumLineChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScoreSectionNumLineChartResponse) GetBody() *DescribeScoreSectionNumLineChartResponseBody {
	return s.Body
}

func (s *DescribeScoreSectionNumLineChartResponse) SetHeaders(v map[string]*string) *DescribeScoreSectionNumLineChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponse) SetStatusCode(v int32) *DescribeScoreSectionNumLineChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponse) SetBody(v *DescribeScoreSectionNumLineChartResponseBody) *DescribeScoreSectionNumLineChartResponse {
	s.Body = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
