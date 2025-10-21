// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceIds(v string) *DescribeDBInstancesRequest
	GetDBInstanceIds() *string
	SetDBInstanceStatus(v string) *DescribeDBInstancesRequest
	GetDBInstanceStatus() *string
	SetDescription(v string) *DescribeDBInstancesRequest
	GetDescription() *string
	SetPageNumber(v int32) *DescribeDBInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesRequest
	GetResourceGroupId() *string
}

type DescribeDBInstancesRequest struct {
	// The cluster IDs. Separate multiple cluster IDs with commas (,).
	//
	// example:
	//
	// cc-xxxxx,cx-xxxx
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// active
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) GetDBInstanceIds() *string {
	return s.DBInstanceIds
}

func (s *DescribeDBInstancesRequest) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDescription(v string) *DescribeDBInstancesRequest {
	s.Description = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) Validate() error {
	return dara.Validate(s)
}
