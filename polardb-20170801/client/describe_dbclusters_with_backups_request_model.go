// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersWithBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeDBClustersWithBackupsRequest
	GetDBClusterIds() *string
	SetDBType(v string) *DescribeDBClustersWithBackupsRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClustersWithBackupsRequest
	GetDBVersion() *string
	SetIsDeleted(v int32) *DescribeDBClustersWithBackupsRequest
	GetIsDeleted() *int32
	SetOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClustersWithBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBClustersWithBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersWithBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBClustersWithBackupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBClustersWithBackupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClustersWithBackupsRequest
	GetResourceOwnerId() *int64
}

type DescribeDBClustersWithBackupsRequest struct {
	// The name of the cluster. The name must meet the following requirements:
	//
	// 	- It cannot start with `http://` or `https://`.
	//
	// 	- It must be 2 to 256 characters in length.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the cluster. If you need to specify multiple cluster IDs, separate the cluster IDs with commas (,).
	//
	// example:
	//
	// pc-**************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// 	- Valid values for the MySQL database engine:
	//
	//     	- **5.6**
	//
	//     	- **5.7**
	//
	//     	- **8.0**
	//
	// 	- Valid values for the PostgreSQL database engine:
	//
	//     	- **11**
	//
	//     	- **14**
	//
	// 	- Valid value for the Oracle database engine: **11**
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// Specifies whether the cluster is deleted. Valid values:
	//
	// 	- **0**: not deleted
	//
	// 	- **1**: deleted
	//
	// example:
	//
	// 0
	IsDeleted    *int32  `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
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
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query information about regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBClustersWithBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersWithBackupsRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeDBClustersWithBackupsRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersWithBackupsRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersWithBackupsRequest) GetIsDeleted() *int32 {
	return s.IsDeleted
}

func (s *DescribeDBClustersWithBackupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClustersWithBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClustersWithBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersWithBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersWithBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersWithBackupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersWithBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClustersWithBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBClusterIds(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBType(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBVersion(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetIsDeleted(v int32) *DescribeDBClustersWithBackupsRequest {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetOwnerId(v int64) *DescribeDBClustersWithBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetPageNumber(v int32) *DescribeDBClustersWithBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetPageSize(v int32) *DescribeDBClustersWithBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetRegionId(v string) *DescribeDBClustersWithBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceGroupId(v string) *DescribeDBClustersWithBackupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceOwnerId(v int64) *DescribeDBClustersWithBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) Validate() error {
	return dara.Validate(s)
}
