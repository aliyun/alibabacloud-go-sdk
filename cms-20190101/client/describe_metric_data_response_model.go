// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetricDataResponseBody) *DescribeMetricDataResponse
	GetBody() *DescribeMetricDataResponseBody
}

type DescribeMetricDataResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetricDataResponse) GetBody() *DescribeMetricDataResponseBody {
	return s.Body
}

func (s *DescribeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricDataResponse) SetStatusCode(v int32) *DescribeMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetricDataResponse) SetBody(v *DescribeMetricDataResponseBody) *DescribeMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeMetricDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
