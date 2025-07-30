// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckJobId(v string) *DescribeCheckJobsRequest
	GetCheckJobId() *string
	SetCheckType(v int32) *DescribeCheckJobsRequest
	GetCheckType() *int32
	SetInstanceId(v string) *DescribeCheckJobsRequest
	GetInstanceId() *string
	SetJobName(v string) *DescribeCheckJobsRequest
	GetJobName() *string
	SetPageNumber(v int32) *DescribeCheckJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCheckJobsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeCheckJobsRequest
	GetResourceGroupId() *string
}

type DescribeCheckJobsRequest struct {
	// Check the task job ID.
	//
	// example:
	//
	// z9p104ib23***
	CheckJobId *string `json:"CheckJobId,omitempty" xml:"CheckJobId,omitempty"`
	// The type of the check
	//
	// >>1 full quantity, 2 incremental, 3 all
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// Data migration instance ID, which can be queried by calling the **describemigrationjobs*	- API.
	//
	// example:
	//
	// dtsz9p104ib23e972e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the data migration or synchronization job.
	//
	// example:
	//
	// zwy_test
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The number of the page to return. The value must be an integer that is greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCheckJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckJobsRequest) GetCheckJobId() *string {
	return s.CheckJobId
}

func (s *DescribeCheckJobsRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeCheckJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCheckJobsRequest) GetJobName() *string {
	return s.JobName
}

func (s *DescribeCheckJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCheckJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCheckJobsRequest) SetCheckJobId(v string) *DescribeCheckJobsRequest {
	s.CheckJobId = &v
	return s
}

func (s *DescribeCheckJobsRequest) SetCheckType(v int32) *DescribeCheckJobsRequest {
	s.CheckType = &v
	return s
}

func (s *DescribeCheckJobsRequest) SetInstanceId(v string) *DescribeCheckJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCheckJobsRequest) SetJobName(v string) *DescribeCheckJobsRequest {
	s.JobName = &v
	return s
}

func (s *DescribeCheckJobsRequest) SetPageNumber(v int32) *DescribeCheckJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCheckJobsRequest) SetPageSize(v int32) *DescribeCheckJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckJobsRequest) SetResourceGroupId(v string) *DescribeCheckJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCheckJobsRequest) Validate() error {
	return dara.Validate(s)
}
