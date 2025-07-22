// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndPointMetricDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEndPointMetricDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEndPointMetricDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEndPointMetricDataResponseBody) *DescribeEndPointMetricDataResponse
	GetBody() *DescribeEndPointMetricDataResponseBody
}

type DescribeEndPointMetricDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndPointMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndPointMetricDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEndPointMetricDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEndPointMetricDataResponse) GetBody() *DescribeEndPointMetricDataResponseBody {
	return s.Body
}

func (s *DescribeEndPointMetricDataResponse) SetHeaders(v map[string]*string) *DescribeEndPointMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndPointMetricDataResponse) SetStatusCode(v int32) *DescribeEndPointMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndPointMetricDataResponse) SetBody(v *DescribeEndPointMetricDataResponseBody) *DescribeEndPointMetricDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEndPointMetricDataResponse) Validate() error {
	return dara.Validate(s)
}
