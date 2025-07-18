// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceCategoriesShrink(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceCategoriesShrink() *string
	SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceDescription() *string
	SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceIds() *string
	SetDBInstanceModesShrink(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceModesShrink() *string
	SetDBInstanceStatusesShrink(v string) *DescribeDBInstancesShrinkRequest
	GetDBInstanceStatusesShrink() *string
	SetInstanceDeployTypesShrink(v string) *DescribeDBInstancesShrinkRequest
	GetInstanceDeployTypesShrink() *string
	SetInstanceNetworkType(v string) *DescribeDBInstancesShrinkRequest
	GetInstanceNetworkType() *string
	SetOwnerId(v int64) *DescribeDBInstancesShrinkRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesShrinkRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeDBInstancesShrinkRequestTag) *DescribeDBInstancesShrinkRequest
	GetTag() []*DescribeDBInstancesShrinkRequestTag
	SetVpcId(v string) *DescribeDBInstancesShrinkRequest
	GetVpcId() *string
}

type DescribeDBInstancesShrinkRequest struct {
	// The edition of the instance. Separate multiple values with commas (,).
	DBInstanceCategoriesShrink *string `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance ID. Separate multiple values with commas (,).
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The resource type of the instance. Separate multiple values with commas (,).
	DBInstanceModesShrink *string `json:"DBInstanceModes,omitempty" xml:"DBInstanceModes,omitempty"`
	// The state of the instance.
	DBInstanceStatusesShrink *string `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty"`
	// This parameter is no longer used.
	InstanceDeployTypesShrink *string `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC).
	//
	// 	- **Classic**: classic network.
	//
	// > If you do not specify this parameter, instances of all network types are returned.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*DescribeDBInstancesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID. You can use this parameter to filter instances that reside in the specified VPC.
	//
	// example:
	//
	// vpc-t4nqyp3tc5mx7vy6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDBInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceCategoriesShrink() *string {
	return s.DBInstanceCategoriesShrink
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceIds() *string {
	return s.DBInstanceIds
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceModesShrink() *string {
	return s.DBInstanceModesShrink
}

func (s *DescribeDBInstancesShrinkRequest) GetDBInstanceStatusesShrink() *string {
	return s.DBInstanceStatusesShrink
}

func (s *DescribeDBInstancesShrinkRequest) GetInstanceDeployTypesShrink() *string {
	return s.InstanceDeployTypesShrink
}

func (s *DescribeDBInstancesShrinkRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesShrinkRequest) GetTag() []*DescribeDBInstancesShrinkRequestTag {
	return s.Tag
}

func (s *DescribeDBInstancesShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceCategoriesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceCategoriesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceIds(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceModesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceModesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetDBInstanceStatusesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.DBInstanceStatusesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceDeployTypesShrink(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceDeployTypesShrink = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesShrinkRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetOwnerId(v int64) *DescribeDBInstancesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageNumber(v int32) *DescribeDBInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetPageSize(v int32) *DescribeDBInstancesShrinkRequest {
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

func (s *DescribeDBInstancesShrinkRequest) SetTag(v []*DescribeDBInstancesShrinkRequestTag) *DescribeDBInstancesShrinkRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) SetVpcId(v string) *DescribeDBInstancesShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstancesShrinkRequestTag struct {
	// The key of tag N.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesShrinkRequestTag) SetKey(v string) *DescribeDBInstancesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequestTag) SetValue(v string) *DescribeDBInstancesShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
