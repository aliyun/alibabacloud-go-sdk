// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceDescription() *string
	SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceIds() *string
	SetDBInstanceStatus(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceStatus() *string
	SetPageNumber(v int64) *DescribeDBInstancesShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDBInstancesShrinkRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeDBInstancesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesShrinkRequest
	GetResourceGroupId() *string
	SetTagShrink(v string) *DescribeDBInstancesShrinkRequest
	GetTagShrink() *string
}

type DescribeDBInstancesShrinkRequest struct {
	// The description of the instance.
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **RESOURCE_CHANGING**: The resource configuration of the instance is being changed.
	//
	// 	- **ORDER_PREPARING**: The order is being confirmed.
	//
	// 	- **READONLY_RESOURCE_CHANGING**: The resource configuration of the instance is being changed and the instance is write-locked.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// example:
	//
	// ACTIVATION
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
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
	TagShrink       *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DescribeDBInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceIds() *string {
	return s.DBInstanceIds
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDBInstancesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDBInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageNumber(v int64) *DescribeDBInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageSize(v int64) *DescribeDBInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetRegionId(v string) *DescribeDBInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetResourceGroupId(v string) *DescribeDBInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetTagShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
