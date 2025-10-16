// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowMetricResponseBody) *DescribeFlowMetricResponse
	GetBody() *DescribeFlowMetricResponseBody
}

type DescribeFlowMetricResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowMetricResponse) GetBody() *DescribeFlowMetricResponseBody {
	return s.Body
}

func (s *DescribeFlowMetricResponse) SetHeaders(v map[string]*string) *DescribeFlowMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowMetricResponse) SetStatusCode(v int32) *DescribeFlowMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowMetricResponse) SetBody(v *DescribeFlowMetricResponseBody) *DescribeFlowMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
