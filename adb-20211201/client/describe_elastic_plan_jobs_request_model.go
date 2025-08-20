// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeElasticPlanJobsRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *DescribeElasticPlanJobsRequest
	GetElasticPlanName() *string
	SetPageNumber(v int32) *DescribeElasticPlanJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeElasticPlanJobsRequest
	GetPageSize() *int32
	SetResourceGroupName(v string) *DescribeElasticPlanJobsRequest
	GetResourceGroupName() *string
	SetStartTime(v string) *DescribeElasticPlanJobsRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeElasticPlanJobsRequest
	GetStatus() *string
}

type DescribeElasticPlanJobsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// >
	//
	// 	- If you do not specify this parameter, all scaling plans of the cluster are queried.
	//
	// 	- You can call the [DescribeElasticPlans](https://help.aliyun.com/document_detail/601334.html) operation to query the names of scaling plans.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource group.
	//
	// >
	//
	// 	- If you do not specify this parameter, the scaling plans of all resource groups are queried, including the interactive resource group and elastic I/O unit (EIU) types.
	//
	// 	- You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group name for a cluster.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the scaling plan job. Valid values:
	//
	// 	- RUNNING
	//
	// 	- SUCCESSFUL
	//
	// 	- FAILED
	//
	// >  If you do not specify this parameter, the scaling plans in all states are queried.
	//
	// example:
	//
	// SUCCESSFUL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeElasticPlanJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeElasticPlanJobsRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlanJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeElasticPlanJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeElasticPlanJobsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeElasticPlanJobsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeElasticPlanJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticPlanJobsRequest) SetDBClusterId(v string) *DescribeElasticPlanJobsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetElasticPlanName(v string) *DescribeElasticPlanJobsRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetPageNumber(v int32) *DescribeElasticPlanJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetPageSize(v int32) *DescribeElasticPlanJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetResourceGroupName(v string) *DescribeElasticPlanJobsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetStartTime(v string) *DescribeElasticPlanJobsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) SetStatus(v string) *DescribeElasticPlanJobsRequest {
	s.Status = &v
	return s
}

func (s *DescribeElasticPlanJobsRequest) Validate() error {
	return dara.Validate(s)
}
