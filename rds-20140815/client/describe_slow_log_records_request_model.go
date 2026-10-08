// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest
	GetDBInstanceId() *string
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
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end time of the query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-06-18T16:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The node ID.
	//
	// 	Notice: This parameter is applicable only to Cluster Edition instances. You can specify this parameter to query the logs of a specific node. If you do not specify this parameter, the logs of the primary node are returned by default.
	//
	// example:
	//
	// rn-p1fm78s90x5****
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any integer greater than 0 that does not exceed the maximum value of Integer.
	//
	// > Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**.
	//
	// > Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique identifier of the SQL statement in the slow query log statistics. You can use this parameter to obtain the slow query log details of the SQL statement.
	//
	// example:
	//
	// U2FsdGVk****
	SQLHASH *string `json:"SQLHASH,omitempty" xml:"SQLHASH,omitempty"`
	// The start time of the query. The start time must be within the last 30 days. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm</i>Z format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-06-17T16:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
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

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceId = &v
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
