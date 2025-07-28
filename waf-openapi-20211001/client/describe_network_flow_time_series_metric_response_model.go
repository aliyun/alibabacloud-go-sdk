// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTimeSeriesMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkFlowTimeSeriesMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkFlowTimeSeriesMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkFlowTimeSeriesMetricResponseBody) *DescribeNetworkFlowTimeSeriesMetricResponse
	GetBody() *DescribeNetworkFlowTimeSeriesMetricResponseBody
}

type DescribeNetworkFlowTimeSeriesMetricResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkFlowTimeSeriesMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkFlowTimeSeriesMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTimeSeriesMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) GetBody() *DescribeNetworkFlowTimeSeriesMetricResponseBody {
	return s.Body
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) SetHeaders(v map[string]*string) *DescribeNetworkFlowTimeSeriesMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) SetStatusCode(v int32) *DescribeNetworkFlowTimeSeriesMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) SetBody(v *DescribeNetworkFlowTimeSeriesMetricResponseBody) *DescribeNetworkFlowTimeSeriesMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkFlowTimeSeriesMetricResponse) Validate() error {
	return dara.Validate(s)
}
