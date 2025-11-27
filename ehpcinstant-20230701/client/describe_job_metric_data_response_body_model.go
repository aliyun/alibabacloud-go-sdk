// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataPoints(v string) *DescribeJobMetricDataResponseBody
	GetDataPoints() *string
	SetPeriod(v int32) *DescribeJobMetricDataResponseBody
	GetPeriod() *int32
	SetRequestId(v string) *DescribeJobMetricDataResponseBody
	GetRequestId() *string
}

type DescribeJobMetricDataResponseBody struct {
	// Monitoring statistics points.
	//
	// example:
	//
	// [{"timestamp":1709540685000,"Minimum":28.45,"Maximum":28.45,"Average":28.45}]
	DataPoints *string `json:"DataPoints,omitempty" xml:"DataPoints,omitempty"`
	// The statistical period of the monitoring data. Valid values: 15, 60, 900, and 3600. Unit: seconds.
	//
	// example:
	//
	// 15
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobMetricDataResponseBody) GetDataPoints() *string {
	return s.DataPoints
}

func (s *DescribeJobMetricDataResponseBody) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeJobMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobMetricDataResponseBody) SetDataPoints(v string) *DescribeJobMetricDataResponseBody {
	s.DataPoints = &v
	return s
}

func (s *DescribeJobMetricDataResponseBody) SetPeriod(v int32) *DescribeJobMetricDataResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeJobMetricDataResponseBody) SetRequestId(v string) *DescribeJobMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}
