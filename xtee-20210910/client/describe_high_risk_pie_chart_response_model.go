// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighRiskPieChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHighRiskPieChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHighRiskPieChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHighRiskPieChartResponseBody) *DescribeHighRiskPieChartResponse
	GetBody() *DescribeHighRiskPieChartResponseBody
}

type DescribeHighRiskPieChartResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHighRiskPieChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHighRiskPieChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHighRiskPieChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHighRiskPieChartResponse) GetBody() *DescribeHighRiskPieChartResponseBody {
	return s.Body
}

func (s *DescribeHighRiskPieChartResponse) SetHeaders(v map[string]*string) *DescribeHighRiskPieChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeHighRiskPieChartResponse) SetStatusCode(v int32) *DescribeHighRiskPieChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHighRiskPieChartResponse) SetBody(v *DescribeHighRiskPieChartResponseBody) *DescribeHighRiskPieChartResponse {
	s.Body = v
	return s
}

func (s *DescribeHighRiskPieChartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
