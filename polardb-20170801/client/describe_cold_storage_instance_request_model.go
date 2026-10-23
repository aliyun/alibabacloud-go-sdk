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
	// The PolarDB cluster ID.
	//
	// example:
	//
	// pc-wz9062015ly7526jc
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The supported engine types. The return value is the sum of the following engine type values:
	//
	// - 1: search engine
	//
	// - 2: time series engine
	//
	// - 4: wide table engine
	//
	// - 8: file engine
	//
	// > For example, if the EngineType value is 15, 15 = 8 + 4 + 2 + 1, which indicates that the instance supports the search engine, time series engine, wide table engine, and file engine. If the EngineType value is 6, 6 = 4 + 2, which indicates that the instance supports the time series engine and wide table engine.
	//
	// example:
	//
	// 2
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the cluster. Note: A specific value is returned only for clusters whose billing method is Prepaid (subscription). For Postpaid (pay-as-you-go) clusters, a null value is returned.
	//
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// - If you do not specify the **MaxResults*	- parameter, batch query is not required. The value of the **MaxResults*	- parameter returned after the API call indicates the total number of list entries.
	//
	// - If you specify the **MaxResults*	- parameter, batch query is required. The value of **MaxResults*	- indicates the number of entries displayed per batch. Valid values: **1*	- to **100**. The value of the **MaxResults*	- parameter returned after the API call indicates the number of list entries in the current query batch. Set the **MaxResults*	- parameter to **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned from the previous API call. If there are no more queries, do not specify this parameter.
	//
	// example:
	//
	// c2FpXzIwMjIwNjI5X2Jhay9zYWlfc3VtbWVyX3RyZWFzdXJlX3Bvb2xfbG9nLkNTVg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The object type. Valid values: TABLE, PARTITION_TABLE, and LOB.
	//
	// example:
	//
	// TABLE
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > 	- You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the region information of all clusters under the target account.
	//
	// 	- If this parameter is left empty, scheduled tasks in all regions under the current account are queried by default.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the data table.
	//
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
