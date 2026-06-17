// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogRecordsRequest
	GetDBClusterId() *string
	SetDBName(v string) *DescribeSlowLogRecordsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeSlowLogRecordsRequest
	GetEndTime() *string
	SetNodeId(v string) *DescribeSlowLogRecordsRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSlowLogRecordsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSlowLogRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSlowLogRecordsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSlowLogRecordsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest
	GetResourceOwnerId() *int64
	SetSQLHASH(v string) *DescribeSlowLogRecordsRequest
	GetSQLHASH() *string
	SetStartTime(v string) *DescribeSlowLogRecordsRequest
	GetStartTime() *string
}

type DescribeSlowLogRecordsRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query all clusters in the target region and their cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the query time range. The end time must be later than the start time. The time range cannot exceed 24 hours. Specify the time in UTC in the `YYYY-MM-DDThh:mmZ` format.
	//
	// > The time must be in UTC. If your service is in a different time zone, you must convert the time. For example, to query data from 08:00 to 12:00 in the UTC+8 time zone, you must set the time range from 00:00 UTC to 04:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-16T04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pi-**********
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer greater than 0.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values:
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
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions and their region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The SQL hash of a slow query. Obtain this hash from slow query log statistics to retrieve the details of a specific slow query.
	//
	// example:
	//
	// U2FsdGVk****
	SQLHASH *string `json:"SQLHASH,omitempty" xml:"SQLHASH,omitempty"`
	// The start of the query time range. Specify the time in UTC in the `YYYY-MM-DDThh:mmZ` format.
	//
	// > - You can query slow query logs from the past 30 days.
	//
	// >
	//
	// > - The time must be in UTC. If your service is in a different time zone, you must convert the time. For example, to query data from 08:00 to 12:00 in the UTC+8 time zone, you must set the time range from 00:00 UTC to 04:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-15T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeSlowLogRecordsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogRecordsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogRecordsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSlowLogRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSlowLogRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlowLogRecordsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlowLogRecordsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlowLogRecordsRequest) GetSQLHASH() *string {
	return s.SQLHASH
}

func (s *DescribeSlowLogRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogRecordsRequest) SetDBClusterId(v string) *DescribeSlowLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetNodeId(v string) *DescribeSlowLogRecordsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLHASH(v string) *DescribeSlowLogRecordsRequest {
	s.SQLHASH = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) Validate() error {
	return dara.Validate(s)
}
