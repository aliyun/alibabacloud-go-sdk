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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// RDS_MySQL
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time span between the start time and the end time cannot exceed 31 days. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*Z format. The time must be in UTC.
	//
	// >  If the end date of the query is the same as the start date of the query, you can query the logs that are generated at 08:00 on the start date of the query. You can query the slow logs within a maximum time range of 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2011-05-30Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from 1.
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
	// The dimension based on which the system sorts the entries to return. Valid values:
	//
	// 	- **TotalExecutionCounts**: The system sorts the entries to return based on the number of times that SQL statements are executed.
	//
	// 	- **TotalQueryTimes**: The system sorts the entries to return based on the total execution duration.
	//
	// 	- **TotalLogicalReads**: The system sorts the entries to return based on the total number of logical reads.
	//
	// 	- **TotalPhysicalReads**: The system sorts the entries to return based on the total number of physical reads.
	//
	// > This parameter is supported only for instances that run SQL Server 2008 R2.
	//
	// example:
	//
	// TotalExecutionCounts
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*Z format. The time must be in UTC.
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
