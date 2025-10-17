// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudProvider(v string) *DescribeDBClustersZonalRequest
	GetCloudProvider() *string
	SetConnectionString(v string) *DescribeDBClustersZonalRequest
	GetConnectionString() *string
	SetDBClusterDescription(v string) *DescribeDBClustersZonalRequest
	GetDBClusterDescription() *string
	SetDBClusterIds(v string) *DescribeDBClustersZonalRequest
	GetDBClusterIds() *string
	SetDBClusterStatus(v string) *DescribeDBClustersZonalRequest
	GetDBClusterStatus() *string
	SetDBNodeIds(v string) *DescribeDBClustersZonalRequest
	GetDBNodeIds() *string
	SetDBType(v string) *DescribeDBClustersZonalRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClustersZonalRequest
	GetDBVersion() *string
	SetDescribeType(v string) *DescribeDBClustersZonalRequest
	GetDescribeType() *string
	SetExpired(v string) *DescribeDBClustersZonalRequest
	GetExpired() *string
	SetMaxResults(v int32) *DescribeDBClustersZonalRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDBClustersZonalRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeDBClustersZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClustersZonalRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBClustersZonalRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBClustersZonalRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeDBClustersZonalRequest
	GetPayType() *string
	SetRecentCreationInterval(v int32) *DescribeDBClustersZonalRequest
	GetRecentCreationInterval() *int32
	SetRecentExpirationInterval(v int32) *DescribeDBClustersZonalRequest
	GetRecentExpirationInterval() *int32
	SetRegionId(v string) *DescribeDBClustersZonalRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBClustersZonalRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClustersZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClustersZonalRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeDBClustersZonalRequestTag) *DescribeDBClustersZonalRequest
	GetTag() []*DescribeDBClustersZonalRequestTag
}

type DescribeDBClustersZonalRequest struct {
	// example:
	//
	// AlibabaCloud
	CloudProvider *string `json:"CloudProvider,omitempty" xml:"CloudProvider,omitempty"`
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// pc-****************
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// example:
	//
	// pc-****************
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// example:
	//
	// pi-***************
	DBNodeIds *string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// Simple
	DescribeType *string `json:"DescribeType,omitempty" xml:"DescribeType,omitempty"`
	// example:
	//
	// true
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 7
	RecentCreationInterval *int32 `json:"RecentCreationInterval,omitempty" xml:"RecentCreationInterval,omitempty"`
	// example:
	//
	// 6
	RecentExpirationInterval *int32 `json:"RecentExpirationInterval,omitempty" xml:"RecentExpirationInterval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-**********
	ResourceGroupId      *string                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                               `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Tag                  []*DescribeDBClustersZonalRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersZonalRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersZonalRequest) GetCloudProvider() *string {
	return s.CloudProvider
}

func (s *DescribeDBClustersZonalRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClustersZonalRequest) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeDBClustersZonalRequest) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeDBClustersZonalRequest) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeDBClustersZonalRequest) GetDBNodeIds() *string {
	return s.DBNodeIds
}

func (s *DescribeDBClustersZonalRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClustersZonalRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClustersZonalRequest) GetDescribeType() *string {
	return s.DescribeType
}

func (s *DescribeDBClustersZonalRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBClustersZonalRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDBClustersZonalRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDBClustersZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClustersZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClustersZonalRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBClustersZonalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBClustersZonalRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClustersZonalRequest) GetRecentCreationInterval() *int32 {
	return s.RecentCreationInterval
}

func (s *DescribeDBClustersZonalRequest) GetRecentExpirationInterval() *int32 {
	return s.RecentExpirationInterval
}

func (s *DescribeDBClustersZonalRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClustersZonalRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBClustersZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClustersZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClustersZonalRequest) GetTag() []*DescribeDBClustersZonalRequestTag {
	return s.Tag
}

func (s *DescribeDBClustersZonalRequest) SetCloudProvider(v string) *DescribeDBClustersZonalRequest {
	s.CloudProvider = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetConnectionString(v string) *DescribeDBClustersZonalRequest {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDBClusterDescription(v string) *DescribeDBClustersZonalRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDBClusterIds(v string) *DescribeDBClustersZonalRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDBClusterStatus(v string) *DescribeDBClustersZonalRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDBNodeIds(v string) *DescribeDBClustersZonalRequest {
	s.DBNodeIds = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDBType(v string) *DescribeDBClustersZonalRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDBVersion(v string) *DescribeDBClustersZonalRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetDescribeType(v string) *DescribeDBClustersZonalRequest {
	s.DescribeType = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetExpired(v string) *DescribeDBClustersZonalRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetMaxResults(v int32) *DescribeDBClustersZonalRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetNextToken(v string) *DescribeDBClustersZonalRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetOwnerAccount(v string) *DescribeDBClustersZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetOwnerId(v int64) *DescribeDBClustersZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetPageNumber(v int32) *DescribeDBClustersZonalRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetPageSize(v int32) *DescribeDBClustersZonalRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetPayType(v string) *DescribeDBClustersZonalRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetRecentCreationInterval(v int32) *DescribeDBClustersZonalRequest {
	s.RecentCreationInterval = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetRecentExpirationInterval(v int32) *DescribeDBClustersZonalRequest {
	s.RecentExpirationInterval = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetRegionId(v string) *DescribeDBClustersZonalRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetResourceGroupId(v string) *DescribeDBClustersZonalRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetResourceOwnerId(v int64) *DescribeDBClustersZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersZonalRequest) SetTag(v []*DescribeDBClustersZonalRequestTag) *DescribeDBClustersZonalRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBClustersZonalRequest) Validate() error {
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

type DescribeDBClustersZonalRequestTag struct {
	// example:
	//
	// MySQL
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 5.6
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersZonalRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersZonalRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersZonalRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDBClustersZonalRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDBClustersZonalRequestTag) SetKey(v string) *DescribeDBClustersZonalRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersZonalRequestTag) SetValue(v string) *DescribeDBClustersZonalRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDBClustersZonalRequestTag) Validate() error {
	return dara.Validate(s)
}
