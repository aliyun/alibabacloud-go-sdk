// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskLineChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskLineChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskLineChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskLineChartResponseBody) *DescribeRiskLineChartResponse
	GetBody() *DescribeRiskLineChartResponseBody
}

type DescribeRiskLineChartResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskLineChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskLineChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLineChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskLineChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskLineChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskLineChartResponse) GetBody() *DescribeRiskLineChartResponseBody {
	return s.Body
}

func (s *DescribeRiskLineChartResponse) SetHeaders(v map[string]*string) *DescribeRiskLineChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskLineChartResponse) SetStatusCode(v int32) *DescribeRiskLineChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskLineChartResponse) SetBody(v *DescribeRiskLineChartResponseBody) *DescribeRiskLineChartResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskLineChartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
