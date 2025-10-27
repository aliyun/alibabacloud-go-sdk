// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChartDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeChartDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeChartDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeChartDataResponseBody) *DescribeChartDataResponse
	GetBody() *DescribeChartDataResponseBody
}

type DescribeChartDataResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChartDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChartDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChartDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeChartDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeChartDataResponse) GetBody() *DescribeChartDataResponseBody {
	return s.Body
}

func (s *DescribeChartDataResponse) SetHeaders(v map[string]*string) *DescribeChartDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChartDataResponse) SetStatusCode(v int32) *DescribeChartDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChartDataResponse) SetBody(v *DescribeChartDataResponseBody) *DescribeChartDataResponse {
	s.Body = v
	return s
}

func (s *DescribeChartDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
