// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeColdStorageInstanceRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeColdStorageInstanceRequest
	GetDBName() *string
	SetEngineType(v string) *DescribeColdStorageInstanceRequest
	GetEngineType() *string
	SetExpireTime(v int32) *DescribeColdStorageInstanceRequest
	GetExpireTime() *int32
	SetMaxResults(v int32) *DescribeColdStorageInstanceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeColdStorageInstanceRequest
	GetNextToken() *string
	SetObjectType(v string) *DescribeColdStorageInstanceRequest
	GetObjectType() *string
	SetOwnerAccount(v string) *DescribeColdStorageInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeColdStorageInstanceRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeColdStorageInstanceRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeColdStorageInstanceRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeColdStorageInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeColdStorageInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeColdStorageInstanceRequest
	GetResourceOwnerId() *int64
	SetTableName(v string) *DescribeColdStorageInstanceRequest
	GetTableName() *string
}

type DescribeColdStorageInstanceRequest struct {
	// example:
	//
	// pc-wz9062015ly7526jc
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// 2
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// c2FpXzIwMjIwNjI5X2Jhay9zYWlfc3VtbWVyX3RyZWFzdXJlX3Bvb2xfbG9nLkNTVg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// TABLE
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// account_log
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColdStorageInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageInstanceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeColdStorageInstanceRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeColdStorageInstanceRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeColdStorageInstanceRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *DescribeColdStorageInstanceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeColdStorageInstanceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeColdStorageInstanceRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeColdStorageInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeColdStorageInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeColdStorageInstanceRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeColdStorageInstanceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeColdStorageInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColdStorageInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeColdStorageInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeColdStorageInstanceRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColdStorageInstanceRequest) SetDBClusterId(v string) *DescribeColdStorageInstanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetDBName(v string) *DescribeColdStorageInstanceRequest {
	s.DBName = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetEngineType(v string) *DescribeColdStorageInstanceRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetExpireTime(v int32) *DescribeColdStorageInstanceRequest {
	s.ExpireTime = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetMaxResults(v int32) *DescribeColdStorageInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetNextToken(v string) *DescribeColdStorageInstanceRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetObjectType(v string) *DescribeColdStorageInstanceRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetOwnerAccount(v string) *DescribeColdStorageInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetOwnerId(v int64) *DescribeColdStorageInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetPageNumber(v string) *DescribeColdStorageInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetPageSize(v string) *DescribeColdStorageInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetRegionId(v string) *DescribeColdStorageInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetResourceOwnerAccount(v string) *DescribeColdStorageInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetResourceOwnerId(v int64) *DescribeColdStorageInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) SetTableName(v string) *DescribeColdStorageInstanceRequest {
	s.TableName = &v
	return s
}

func (s *DescribeColdStorageInstanceRequest) Validate() error {
	return dara.Validate(s)
}
