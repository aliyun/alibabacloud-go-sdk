// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionPieChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScoreSectionPieChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScoreSectionPieChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScoreSectionPieChartResponseBody) *DescribeScoreSectionPieChartResponse
	GetBody() *DescribeScoreSectionPieChartResponseBody
}

type DescribeScoreSectionPieChartResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScoreSectionPieChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScoreSectionPieChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScoreSectionPieChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScoreSectionPieChartResponse) GetBody() *DescribeScoreSectionPieChartResponseBody {
	return s.Body
}

func (s *DescribeScoreSectionPieChartResponse) SetHeaders(v map[string]*string) *DescribeScoreSectionPieChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeScoreSectionPieChartResponse) SetStatusCode(v int32) *DescribeScoreSectionPieChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponse) SetBody(v *DescribeScoreSectionPieChartResponseBody) *DescribeScoreSectionPieChartResponse {
	s.Body = v
	return s
}

func (s *DescribeScoreSectionPieChartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
