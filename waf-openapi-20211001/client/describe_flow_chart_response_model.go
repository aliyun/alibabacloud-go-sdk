// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFlowChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFlowChartResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFlowChartResponseBody) *DescribeFlowChartResponse
	GetBody() *DescribeFlowChartResponseBody
}

type DescribeFlowChartResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFlowChartResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowChartResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFlowChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFlowChartResponse) GetBody() *DescribeFlowChartResponseBody {
	return s.Body
}

func (s *DescribeFlowChartResponse) SetHeaders(v map[string]*string) *DescribeFlowChartResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowChartResponse) SetStatusCode(v int32) *DescribeFlowChartResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowChartResponse) SetBody(v *DescribeFlowChartResponseBody) *DescribeFlowChartResponse {
	s.Body = v
	return s
}

func (s *DescribeFlowChartResponse) Validate() error {
	return dara.Validate(s)
}
