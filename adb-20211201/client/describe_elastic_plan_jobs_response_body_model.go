// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobs(v []*DescribeElasticPlanJobsResponseBodyJobs) *DescribeElasticPlanJobsResponseBody
	GetJobs() []*DescribeElasticPlanJobsResponseBodyJobs
	SetPageNumber(v int32) *DescribeElasticPlanJobsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeElasticPlanJobsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeElasticPlanJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeElasticPlanJobsResponseBody
	GetTotalCount() *int32
}

type DescribeElasticPlanJobsResponseBody struct {
	// The queried scaling plan jobs.
	Jobs []*DescribeElasticPlanJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5C433C2-001F-58E3-99F5-3274C14DF8BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of scaling plan jobs.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticPlanJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsResponseBody) GetJobs() []*DescribeElasticPlanJobsResponseBodyJobs {
	return s.Jobs
}

func (s *DescribeElasticPlanJobsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeElasticPlanJobsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeElasticPlanJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticPlanJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeElasticPlanJobsResponseBody) SetJobs(v []*DescribeElasticPlanJobsResponseBodyJobs) *DescribeElasticPlanJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetPageNumber(v int32) *DescribeElasticPlanJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetPageSize(v int32) *DescribeElasticPlanJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetRequestId(v string) *DescribeElasticPlanJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) SetTotalCount(v int32) *DescribeElasticPlanJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeElasticPlanJobsResponseBodyJobs struct {
	// The amount of elastic resources.
	//
	// >
	//
	// 	- If Type is set to EXECUTOR, ElasticAcu indicates the amount of elastic resources in the current resource group.
	//
	// 	- If Type is set to WORKER, ElasticAcu indicates the total amount of elastic storage resources in the current cluster.
	//
	// example:
	//
	// 16ACU
	ElasticAcu *string `json:"ElasticAcu,omitempty" xml:"ElasticAcu,omitempty"`
	// The name of the scaling plan.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The end time of the scaling plan job.
	//
	// >  The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-01T12:01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of compute nodes or storage replica sets.
	//
	// >
	//
	// 	- If Type is set to EXECUTOR, InstanceSize indicates the number of compute nodes in the cluster.
	//
	// 	- If Type is set to EXECUTOR, InstanceSize indicates the number of storage replica sets in the cluster.
	//
	// example:
	//
	// 1
	InstanceSize *int32 `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// The amount of reserved resources.
	//
	// >
	//
	// 	- If Type is set to EXECUTOR, ReserveAcu indicates the amount of reserved resources in the current resource group.
	//
	// 	- If Type is set to WORKER, ReserveAcu indicates the total amount of reserved storage resources in the current cluster.
	//
	// example:
	//
	// 16ACU
	ReserveAcu *string `json:"ReserveAcu,omitempty" xml:"ReserveAcu,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The start time of the scaling plan job.
	//
	// >  The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-01-01T11:01:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the scaling plan job. Valid values:
	//
	// 	- RUNNING
	//
	// 	- SUCCESSFUL
	//
	// 	- FAILED
	//
	// example:
	//
	// SUCCESSFUL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The desired specifications of elastic resources after scaling.
	//
	// example:
	//
	// 32ACU
	TargetSize *string `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	// The total amount of resources.
	//
	// >
	//
	// 	- If Type is set to EXECUTOR, TotalAcu indicates the total amount of computing resources in the current resource group.
	//
	// 	- If Type is set to WORKER, TotalAcu indicates the total amount of storage resources in the cluster.
	//
	// example:
	//
	// 32ACU
	TotalAcu *string `json:"TotalAcu,omitempty" xml:"TotalAcu,omitempty"`
	// The type of the scaling plan job. Valid values:
	//
	// 	- EXECUTOR: the interactive resource group type, which indicates the computing resource type.
	//
	// 	- WORKER: the EIU type.
	//
	// example:
	//
	// EXECUTOR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlanJobsResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetElasticAcu() *string {
	return s.ElasticAcu
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetInstanceSize() *int32 {
	return s.InstanceSize
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetReserveAcu() *string {
	return s.ReserveAcu
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetTargetSize() *string {
	return s.TargetSize
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetTotalAcu() *string {
	return s.TotalAcu
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) GetType() *string {
	return s.Type
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetElasticAcu(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ElasticAcu = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetElasticPlanName(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetEndTime(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.EndTime = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetInstanceSize(v int32) *DescribeElasticPlanJobsResponseBodyJobs {
	s.InstanceSize = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetReserveAcu(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ReserveAcu = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetResourceGroupName(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetStartTime(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.StartTime = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetStatus(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetTargetSize(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.TargetSize = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetTotalAcu(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.TotalAcu = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) SetType(v string) *DescribeElasticPlanJobsResponseBodyJobs {
	s.Type = &v
	return s
}

func (s *DescribeElasticPlanJobsResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}
