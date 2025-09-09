// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsShardingDbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeDrdsShardingDbsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDrdsShardingDbsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeDrdsShardingDbsResponseBody
	GetRequestId() *string
	SetShardingDbs(v *DescribeDrdsShardingDbsResponseBodyShardingDbs) *DescribeDrdsShardingDbsResponseBody
	GetShardingDbs() *DescribeDrdsShardingDbsResponseBodyShardingDbs
	SetSuccess(v bool) *DescribeDrdsShardingDbsResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeDrdsShardingDbsResponseBody
	GetTotal() *string
}

type DescribeDrdsShardingDbsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of database shards returned per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 509BDE17-505A-4B3B-854D-30D3F092502F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of returned database shards.
	ShardingDbs *DescribeDrdsShardingDbsResponseBodyShardingDbs `json:"ShardingDbs,omitempty" xml:"ShardingDbs,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of returned database shards.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDrdsShardingDbsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDrdsShardingDbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsShardingDbsResponseBody) GetShardingDbs() *DescribeDrdsShardingDbsResponseBodyShardingDbs {
	return s.ShardingDbs
}

func (s *DescribeDrdsShardingDbsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsShardingDbsResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeDrdsShardingDbsResponseBody) SetPageNumber(v string) *DescribeDrdsShardingDbsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetPageSize(v string) *DescribeDrdsShardingDbsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetRequestId(v string) *DescribeDrdsShardingDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetShardingDbs(v *DescribeDrdsShardingDbsResponseBodyShardingDbs) *DescribeDrdsShardingDbsResponseBody {
	s.ShardingDbs = v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetSuccess(v bool) *DescribeDrdsShardingDbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) SetTotal(v string) *DescribeDrdsShardingDbsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsShardingDbsResponseBodyShardingDbs struct {
	ShardingDb []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb `json:"ShardingDb,omitempty" xml:"ShardingDb,omitempty" type:"Repeated"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbs) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbs) GetShardingDb() []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	return s.ShardingDb
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbs) SetShardingDb(v []*DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) *DescribeDrdsShardingDbsResponseBodyShardingDbs {
	s.ShardingDb = v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbs) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb struct {
	// The timeout period for a transaction to wait for the release of the data lock.
	//
	// example:
	//
	// 5000
	BlockingTimeout *int32 `json:"BlockingTimeout,omitempty" xml:"BlockingTimeout,omitempty"`
	// The URL that is used to access the Apsara RDS for MySQL instance.
	//
	// example:
	//
	// 100.100.64.100:11569
	ConnectUrl *string `json:"ConnectUrl,omitempty" xml:"ConnectUrl,omitempty"`
	// The properties of the connection string.
	//
	// example:
	//
	// connectTimeout=3000;autoReconnect=true;failOverReadOnly=false;socketTimeout=900000;rewriteBatchedStatements=true;characterEncoding=utf8
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The ID of the Apsara RDS for MySQL instance that is used as the storage of the database shard.
	//
	// example:
	//
	// rm-bp1hjzn0yv5j2****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The status of the database.
	//
	// example:
	//
	// running
	DbStatus *string `json:"DbStatus,omitempty" xml:"DbStatus,omitempty"`
	// The engine of the database.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The name of group on which the database shard is stored.
	//
	// example:
	//
	// TEMP1_1568171495522SABE_KUP4_0000
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The timeout period of an idle connection.
	//
	// example:
	//
	// 30
	IdleTimeOut *int32 `json:"IdleTimeOut,omitempty" xml:"IdleTimeOut,omitempty"`
	// The maximum size of the connection pool.
	//
	// example:
	//
	// 60
	MaxPoolSize *int32 `json:"MaxPoolSize,omitempty" xml:"MaxPoolSize,omitempty"`
	// The minimum size of the connection pool.
	//
	// example:
	//
	// 5
	MinPoolSize *int32 `json:"MinPoolSize,omitempty" xml:"MinPoolSize,omitempty"`
	// The size of cache for the returned results.
	//
	// example:
	//
	// 0
	PreparedStatementCacheSize *int32 `json:"PreparedStatementCacheSize,omitempty" xml:"PreparedStatementCacheSize,omitempty"`
	// The name of the database shard.
	//
	// example:
	//
	// temp1_zhk1_0000
	ShardingDbName *string `json:"ShardingDbName,omitempty" xml:"ShardingDbName,omitempty"`
	// The username that is used to connect to the ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// pg284mi8
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetBlockingTimeout() *int32 {
	return s.BlockingTimeout
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetConnectUrl() *string {
	return s.ConnectUrl
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetDbStatus() *string {
	return s.DbStatus
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetDbType() *string {
	return s.DbType
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetIdleTimeOut() *int32 {
	return s.IdleTimeOut
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetMaxPoolSize() *int32 {
	return s.MaxPoolSize
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetMinPoolSize() *int32 {
	return s.MinPoolSize
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetPreparedStatementCacheSize() *int32 {
	return s.PreparedStatementCacheSize
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetShardingDbName() *string {
	return s.ShardingDbName
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetBlockingTimeout(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.BlockingTimeout = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectUrl(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectUrl = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetConnectionProperties(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ConnectionProperties = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbInstanceId(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbInstanceId = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbStatus(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbStatus = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetDbType(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.DbType = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetGroupName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.GroupName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetIdleTimeOut(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.IdleTimeOut = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMaxPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MaxPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetMinPoolSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.MinPoolSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetPreparedStatementCacheSize(v int32) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.PreparedStatementCacheSize = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetShardingDbName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.ShardingDbName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) SetUserName(v string) *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb {
	s.UserName = &v
	return s
}

func (s *DescribeDrdsShardingDbsResponseBodyShardingDbsShardingDb) Validate() error {
	return dara.Validate(s)
}
