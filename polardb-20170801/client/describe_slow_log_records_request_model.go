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
	// Cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) interface to view all cluster information in the target region, including the Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Database name.
	//
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// End time of the query, which must be later than the start time, and the time interval between the start and end times must not exceed 24 hours. The format is `YYYY-MM-DDThh:mmZ` (UTC time).
	//
	// > The input is UTC time (i.e., 0 timezone). If your service is currently in a different timezone, please perform a time conversion. For example, if the current timezone of your service is Beijing Time (UTC+8) at 12:00, and you need to query the slow logs between 08:00-12:00 Beijing Time, you should input 00:00-04:00.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-16T04:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Node ID
	//
	// example:
	//
	// pi-**********
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be an integer that is greater than 0.
	//
	// Default value: **1**.
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
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) interface to view the available regions under the target account, including the Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Unique identifier of the SQL statement in the slow log statistics, which can be used to obtain the detailed slow logs for that SQL statement.
	//
	// example:
	//
	// U2FsdGVk****
	SQLHASH *string `json:"SQLHASH,omitempty" xml:"SQLHASH,omitempty"`
	// Start time of the query. The format is `YYYY-MM-DDThh:mmZ` (UTC time).
	//
	// > 	- Supports viewing slow log information up to 30 days.
	//
	// > 	- The input is UTC time (i.e., 0 timezone). If your service is currently in a different timezone, please perform a time conversion. For example, if the current timezone of your service is Beijing Time (UTC+8) at 12:00, and you need to query the slow logs between 08:00-12:00 Beijing Time, you should input 00:00-04:00.
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
