// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetricName(v string) *DescribeMetricRuleCountRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricRuleCountRequest
	GetNamespace() *string
	SetRegionId(v string) *DescribeMetricRuleCountRequest
	GetRegionId() *string
}

type DescribeMetricRuleCountRequest struct {
	// The metric name. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMetricRuleCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleCountRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleCountRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricRuleCountRequest) SetMetricName(v string) *DescribeMetricRuleCountRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleCountRequest) SetNamespace(v string) *DescribeMetricRuleCountRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleCountRequest) SetRegionId(v string) *DescribeMetricRuleCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleCountRequest) Validate() error {
	return dara.Validate(s)
}
