// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkFlowTopNMetricResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkFlowTopNMetricResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkFlowTopNMetricResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkFlowTopNMetricResponseBody) *DescribeNetworkFlowTopNMetricResponse
	GetBody() *DescribeNetworkFlowTopNMetricResponseBody
}

type DescribeNetworkFlowTopNMetricResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkFlowTopNMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkFlowTopNMetricResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkFlowTopNMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkFlowTopNMetricResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkFlowTopNMetricResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkFlowTopNMetricResponse) GetBody() *DescribeNetworkFlowTopNMetricResponseBody {
	return s.Body
}

func (s *DescribeNetworkFlowTopNMetricResponse) SetHeaders(v map[string]*string) *DescribeNetworkFlowTopNMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponse) SetStatusCode(v int32) *DescribeNetworkFlowTopNMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponse) SetBody(v *DescribeNetworkFlowTopNMetricResponseBody) *DescribeNetworkFlowTopNMetricResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkFlowTopNMetricResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
