// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskTagsLineChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRiskTagsLineChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRiskTagsLineChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRiskTagsLineChartResponseBody) *DescribeRiskTagsLineChartResponse
	GetBody() *DescribeRiskTagsLineChartResponseBody
}

type DescribeRiskTagsLineChartResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskTagsLineChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskTagsLineChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTagsLineChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskTagsLineChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRiskTagsLineChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRiskTagsLineChartResponse) GetBody() *DescribeRiskTagsLineChartResponseBody {
	return s.Body
}

func (s *DescribeRiskTagsLineChartResponse) SetHeaders(v map[string]*string) *DescribeRiskTagsLineChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskTagsLineChartResponse) SetStatusCode(v int32) *DescribeRiskTagsLineChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponse) SetBody(v *DescribeRiskTagsLineChartResponseBody) *DescribeRiskTagsLineChartResponse {
	s.Body = v
	return s
}

func (s *DescribeRiskTagsLineChartResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
