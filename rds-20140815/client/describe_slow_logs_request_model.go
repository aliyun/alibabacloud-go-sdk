// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSlowLogsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeSlowLogsRequest
	GetDBName() *string
	SetEndTime(v string) *DescribeSlowLogsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeSlowLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSlowLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSlowLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSlowLogsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeSlowLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlowLogsRequest
	GetResourceOwnerId() *int64
	SetSortKey(v string) *DescribeSlowLogsRequest
	GetSortKey() *string
	SetStartTime(v string) *DescribeSlowLogsRequest
	GetStartTime() *string
}

type DescribeSlowLogsRequest struct {
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
	// RDS_MySQL
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end date of the query. The end date must be later than or equal to the start date, and the interval between the start date and the end date cannot exceed 31 days. Format: <i>yyyy-MM-dd</i>Z (UTC).
	//
	// > If the end date is the same as the start date, the query starts from 08:00 on the start date and covers up to 24 hours of slow query log statistics.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-05-30Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30*	- to **100**. Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The sorting criterion. Valid values:
	//
	// 	- **TotalExecutionCounts**: sorted by total number of executions in descending order.
	//
	// 	- **TotalQueryTimes**: sorted by total execution duration in descending order.
	//
	// 	- **TotalLogicalReads**: sorted by total number of logical reads in descending order.
	//
	// 	- **TotalPhysicalReads**: sorted by total number of physical reads in descending order.
	//
	// > This parameter is supported only for SQL Server 2008 R2 instances.
	//
	// example:
	//
	// TotalExecutionCounts
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// The start date of the query. Format: <i>yyyy-MM-dd</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-05-01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSlowLogsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeSlowLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSlowLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSlowLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSlowLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSlowLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSlowLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlowLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlowLogsRequest) GetSortKey() *string {
	return s.SortKey
}

func (s *DescribeSlowLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogsRequest) SetDBInstanceId(v string) *DescribeSlowLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetDBName(v string) *DescribeSlowLogsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetEndTime(v string) *DescribeSlowLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetOwnerAccount(v string) *DescribeSlowLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetOwnerId(v int64) *DescribeSlowLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetPageNumber(v int32) *DescribeSlowLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetPageSize(v int32) *DescribeSlowLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetSortKey(v string) *DescribeSlowLogsRequest {
	s.SortKey = &v
	return s
}

func (s *DescribeSlowLogsRequest) SetStartTime(v string) *DescribeSlowLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogsRequest) Validate() error {
	return dara.Validate(s)
}
