// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsActionLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeApsActionLogsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeApsActionLogsRequest
	GetEndTime() *string
	SetKeyword(v string) *DescribeApsActionLogsRequest
	GetKeyword() *string
	SetOwnerAccount(v string) *DescribeApsActionLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeApsActionLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeApsActionLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApsActionLogsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApsActionLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeApsActionLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeApsActionLogsRequest
	GetResourceOwnerId() *int64
	SetStage(v string) *DescribeApsActionLogsRequest
	GetStage() *string
	SetStartTime(v string) *DescribeApsActionLogsRequest
	GetStartTime() *string
	SetState(v string) *DescribeApsActionLogsRequest
	GetState() *string
	SetWorkloadId(v string) *DescribeApsActionLogsRequest
	GetWorkloadId() *string
}

type DescribeApsActionLogsRequest struct {
	// <props="china">The ID of the cluster. The cluster must be of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the logs. Specify the time in the **yyyy-MM-ddTHH:mmZ*	- format. The time must be in UTC.
	//
	// > When you query logs, the end time must be later than the start time, and the time span between the start time and the end time cannot exceed 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-11T09:30Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The keyword for the query. Fuzzy matching is applied during log query.
	//
	// example:
	//
	// table_test
	Keyword      *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that is greater than 0 and does not exceed the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30*	- (default)
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task stage. Valid values:
	//
	// - **StructureMigrate**: Schema migration.
	//
	// - **FullDataSync**: Full data synchronization.
	//
	// - **IncrementalSync**: Incremental data synchronization.
	//
	// > If you do not specify this parameter, logs of all task stages are queried by default.
	//
	// example:
	//
	// FullDataSync
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The start time of the logs. Specify the time in the **yyyy-MM-ddTHH:mmZ*	- format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-11T08:30Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The log type. Separate multiple log types with commas (,). Valid values:
	//
	// - **INFO**: Information.
	//
	// - **WARN**: Warning.
	//
	// - **ERROR**: Error.
	//
	// > If you do not specify this parameter, logs of all types are queried by default.
	//
	// example:
	//
	// INFO,WARN,ERROR
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the data link task.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-hz109vpvt4fg8528d****
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s DescribeApsActionLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsActionLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsActionLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApsActionLogsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeApsActionLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeApsActionLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeApsActionLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApsActionLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApsActionLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsActionLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeApsActionLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeApsActionLogsRequest) GetStage() *string {
	return s.Stage
}

func (s *DescribeApsActionLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApsActionLogsRequest) GetState() *string {
	return s.State
}

func (s *DescribeApsActionLogsRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *DescribeApsActionLogsRequest) SetDBClusterId(v string) *DescribeApsActionLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetEndTime(v string) *DescribeApsActionLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetKeyword(v string) *DescribeApsActionLogsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetOwnerAccount(v string) *DescribeApsActionLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetOwnerId(v int64) *DescribeApsActionLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetPageNumber(v int32) *DescribeApsActionLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetPageSize(v int32) *DescribeApsActionLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetRegionId(v string) *DescribeApsActionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetResourceOwnerAccount(v string) *DescribeApsActionLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetResourceOwnerId(v int64) *DescribeApsActionLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetStage(v string) *DescribeApsActionLogsRequest {
	s.Stage = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetStartTime(v string) *DescribeApsActionLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetState(v string) *DescribeApsActionLogsRequest {
	s.State = &v
	return s
}

func (s *DescribeApsActionLogsRequest) SetWorkloadId(v string) *DescribeApsActionLogsRequest {
	s.WorkloadId = &v
	return s
}

func (s *DescribeApsActionLogsRequest) Validate() error {
	return dara.Validate(s)
}
