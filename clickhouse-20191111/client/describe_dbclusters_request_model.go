// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *DescribeDBClustersRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeDBClustersRequest
	GetDBClusterIds() *string
	SetDBClusterStatus(v string) *DescribeDBClustersRequest
	GetDBClusterStatus() *string
	SetOwnerAccount(v string) *DescribeDBClustersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClustersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBClustersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBClustersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClustersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClustersRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest
	GetTag() []*DescribeDBClustersRequestTag
}

type DescribeDBClustersRequest struct {
	// The description of the cluster.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The cluster ID.
	//
	// >  If you do not specify this parameter, the information about all clusters is queried.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The state of the cluster. Valid values:
	//
	// 	- **Preparing**: The cluster is being prepared.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **Running**: The cluster is running.
	//
	// 	- **Deleting**: The cluster is being deleted.
	//
	// 	- **SCALING_OUT**: The storage capacity of the cluster is being expanded.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Valid values:
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeDBClustersRequest) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClustersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClustersRequest) GetTag() []*DescribeDBClustersRequestTag {
	return s.Tag
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerId(v int64) *DescribeDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersRequest) Validate() error {
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

type DescribeDBClustersRequestTag struct {
	// The tag name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// it
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersRequestTag) Validate() error {
	return dara.Validate(s)
}
