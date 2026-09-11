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
	// The ID of the data validation task.
	//
	// example:
	//
	// z9p104ib23***
	CheckJobId *string `json:"CheckJobId,omitempty" xml:"CheckJobId,omitempty"`
	// The data validation method. Valid values:
	//
	// - **1**: full data validation.
	//
	// - **2**: incremental data validation.
	//
	// - **3**: all.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The ID of the data migration instance. You can call the **DescribeMigrationJobs*	- operation to query the ID.
	//
	// example:
	//
	// dtsz9p104ib23e972e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the data migration or data synchronization task.
	//
	// example:
	//
	// zwy_test
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values: **30**, **50**, and **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
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
