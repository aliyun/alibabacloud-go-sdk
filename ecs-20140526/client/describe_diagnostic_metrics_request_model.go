// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeDiagnosticMetricsRequest
	GetMaxResults() *int32
	SetMetricIds(v []*string) *DescribeDiagnosticMetricsRequest
	GetMetricIds() []*string
	SetNextToken(v string) *DescribeDiagnosticMetricsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiagnosticMetricsRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeDiagnosticMetricsRequest
	GetResourceType() *string
}

type DescribeDiagnosticMetricsRequest struct {
	// The number of entries per page in a paging query. Maximum value: 100.
	//
	// Default value:
	//
	// - If this parameter is not set, the default value is 10.
	//
	// - If the value is set to a number greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The list of diagnostic metrics.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
	// The pagination token. Set this parameter to the `NextToken` value returned in the previous call. You do not need to set this parameter for the first request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The supported resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeDiagnosticMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiagnosticMetricsRequest) GetMetricIds() []*string {
	return s.MetricIds
}

func (s *DescribeDiagnosticMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticMetricsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosticMetricsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnosticMetricsRequest) SetMaxResults(v int32) *DescribeDiagnosticMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiagnosticMetricsRequest) SetMetricIds(v []*string) *DescribeDiagnosticMetricsRequest {
	s.MetricIds = v
	return s
}

func (s *DescribeDiagnosticMetricsRequest) SetNextToken(v string) *DescribeDiagnosticMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticMetricsRequest) SetRegionId(v string) *DescribeDiagnosticMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosticMetricsRequest) SetResourceType(v string) *DescribeDiagnosticMetricsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnosticMetricsRequest) Validate() error {
	return dara.Validate(s)
}
