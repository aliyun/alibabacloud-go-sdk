// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogMonitorAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetricName(v string) *DescribeLogMonitorAttributeRequest
	GetMetricName() *string
	SetRegionId(v string) *DescribeLogMonitorAttributeRequest
	GetRegionId() *string
}

type DescribeLogMonitorAttributeRequest struct {
	// The metric name. Exact match is supported.
	//
	// For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLogMonitorAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogMonitorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogMonitorAttributeRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeLogMonitorAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogMonitorAttributeRequest) SetMetricName(v string) *DescribeLogMonitorAttributeRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeLogMonitorAttributeRequest) SetRegionId(v string) *DescribeLogMonitorAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLogMonitorAttributeRequest) Validate() error {
	return dara.Validate(s)
}
