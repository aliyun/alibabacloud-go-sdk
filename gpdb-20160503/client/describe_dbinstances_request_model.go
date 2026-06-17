// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceCategories(v []*string) *DescribeDBInstancesRequest
	GetDBInstanceCategories() []*string
	SetDBInstanceDescription(v string) *DescribeDBInstancesRequest
	GetDBInstanceDescription() *string
	SetDBInstanceIds(v string) *DescribeDBInstancesRequest
	GetDBInstanceIds() *string
	SetDBInstanceModes(v []*string) *DescribeDBInstancesRequest
	GetDBInstanceModes() []*string
	SetDBInstanceStatuses(v []*string) *DescribeDBInstancesRequest
	GetDBInstanceStatuses() []*string
	SetInstanceDeployTypes(v []*string) *DescribeDBInstancesRequest
	GetInstanceDeployTypes() []*string
	SetInstanceNetworkType(v string) *DescribeDBInstancesRequest
	GetInstanceNetworkType() *string
	SetOwnerId(v int64) *DescribeDBInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest
	GetTag() []*DescribeDBInstancesRequestTag
	SetVpcId(v string) *DescribeDBInstancesRequest
	GetVpcId() *string
}

type DescribeDBInstancesRequest struct {
	// The instance edition. You can specify multiple editions. Separate multiple editions with a comma (,).
	DBInstanceCategories []*string `json:"DBInstanceCategories,omitempty" xml:"DBInstanceCategories,omitempty" type:"Repeated"`
	// The instance description.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The instance IDs. Separate multiple IDs with commas (,).
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceIds *string `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	// The resource type of the instance. You can specify multiple resource types. Separate multiple resource types with a comma (,).
	DBInstanceModes []*string `json:"DBInstanceModes,omitempty" xml:"DBInstanceModes,omitempty" type:"Repeated"`
	// The state of the instance.
	DBInstanceStatuses []*string `json:"DBInstanceStatuses,omitempty" xml:"DBInstanceStatuses,omitempty" type:"Repeated"`
	// This parameter is deprecated. Do not specify this parameter.
	InstanceDeployTypes []*string `json:"InstanceDeployTypes,omitempty" xml:"InstanceDeployTypes,omitempty" type:"Repeated"`
	// The instance network type. Valid values:
	//
	// - **VPC**: VPC
	//
	// - **Classic**: classic network
	//
	// > If you omit this parameter, the operation returns instances of all network types.
	//
	// example:
	//
	// VPC
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available region IDs.
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
	// The tags of the instance.
	Tag []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The VPC ID. You can specify this parameter to query instances in a VPC.
	//
	// example:
	//
	// vpc-t4nqyp3tc5mx7vy6****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) GetDBInstanceCategories() []*string {
	return s.DBInstanceCategories
}

func (s *DescribeDBInstancesRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstancesRequest) GetDBInstanceIds() *string {
	return s.DBInstanceIds
}

func (s *DescribeDBInstancesRequest) GetDBInstanceModes() []*string {
	return s.DBInstanceModes
}

func (s *DescribeDBInstancesRequest) GetDBInstanceStatuses() []*string {
	return s.DBInstanceStatuses
}

func (s *DescribeDBInstancesRequest) GetInstanceDeployTypes() []*string {
	return s.InstanceDeployTypes
}

func (s *DescribeDBInstancesRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
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

func (s *DescribeDBInstancesRequest) GetTag() []*DescribeDBInstancesRequestTag {
	return s.Tag
}

func (s *DescribeDBInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesRequest) SetDBInstanceCategories(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceCategories = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceModes(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceModes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatuses(v []*string) *DescribeDBInstancesRequest {
	s.DBInstanceStatuses = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceDeployTypes(v []*string) *DescribeDBInstancesRequest {
	s.InstanceDeployTypes = v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
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

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesRequest) SetVpcId(v string) *DescribeDBInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancesRequestTag) SetKey(v string) *DescribeDBInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) SetValue(v string) *DescribeDBInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
