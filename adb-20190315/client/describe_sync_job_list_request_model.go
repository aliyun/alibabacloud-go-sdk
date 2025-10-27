// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncJobListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSyncJobListRequest
	GetDBClusterId() *string
	SetGetSourceDetail(v bool) *DescribeSyncJobListRequest
	GetGetSourceDetail() *bool
	SetOwnerAccount(v string) *DescribeSyncJobListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSyncJobListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSyncJobListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSyncJobListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSyncJobListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSyncJobListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSyncJobListRequest
	GetResourceOwnerId() *int64
	SetSourceDBClusterDescription(v string) *DescribeSyncJobListRequest
	GetSourceDBClusterDescription() *string
	SetSourceDBClusterId(v string) *DescribeSyncJobListRequest
	GetSourceDBClusterId() *string
	SetSourceDBType(v string) *DescribeSyncJobListRequest
	GetSourceDBType() *string
}

type DescribeSyncJobListRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-8vb39udfi356l9psq
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to obtain details about the source instance or cluster.
	//
	// example:
	//
	// true
	GetSourceDetail *bool   `json:"GetSourceDetail,omitempty" xml:"GetSourceDetail,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
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
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the source cluster.
	//
	// example:
	//
	// test
	SourceDBClusterDescription *string `json:"SourceDBClusterDescription,omitempty" xml:"SourceDBClusterDescription,omitempty"`
	// The ID of the source cluster. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query backup set IDs.
	//
	// >  If you want to restore the data of an ApsaraDB for ClickHouse cluster, this parameter is required.
	//
	// example:
	//
	// pc-t4n766v2llx852n81
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// The source database type.
	//
	// example:
	//
	// sls
	SourceDBType *string `json:"SourceDBType,omitempty" xml:"SourceDBType,omitempty"`
}

func (s DescribeSyncJobListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncJobListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncJobListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSyncJobListRequest) GetGetSourceDetail() *bool {
	return s.GetSourceDetail
}

func (s *DescribeSyncJobListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSyncJobListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSyncJobListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSyncJobListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSyncJobListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyncJobListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSyncJobListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSyncJobListRequest) GetSourceDBClusterDescription() *string {
	return s.SourceDBClusterDescription
}

func (s *DescribeSyncJobListRequest) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *DescribeSyncJobListRequest) GetSourceDBType() *string {
	return s.SourceDBType
}

func (s *DescribeSyncJobListRequest) SetDBClusterId(v string) *DescribeSyncJobListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetGetSourceDetail(v bool) *DescribeSyncJobListRequest {
	s.GetSourceDetail = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetOwnerAccount(v string) *DescribeSyncJobListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetOwnerId(v int64) *DescribeSyncJobListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetPageNumber(v int32) *DescribeSyncJobListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetPageSize(v int32) *DescribeSyncJobListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetRegionId(v string) *DescribeSyncJobListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetResourceOwnerAccount(v string) *DescribeSyncJobListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetResourceOwnerId(v int64) *DescribeSyncJobListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetSourceDBClusterDescription(v string) *DescribeSyncJobListRequest {
	s.SourceDBClusterDescription = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetSourceDBClusterId(v string) *DescribeSyncJobListRequest {
	s.SourceDBClusterId = &v
	return s
}

func (s *DescribeSyncJobListRequest) SetSourceDBType(v string) *DescribeSyncJobListRequest {
	s.SourceDBType = &v
	return s
}

func (s *DescribeSyncJobListRequest) Validate() error {
	return dara.Validate(s)
}
