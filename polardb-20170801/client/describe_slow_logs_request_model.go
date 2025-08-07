// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeSlowLogsRequest
	GetDBClusterId() *string
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
	SetRegionId(v string) *DescribeSlowLogsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSlowLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSlowLogsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeSlowLogsRequest
	GetStartTime() *string
}

type DescribeSlowLogsRequest struct {
	// The ID of cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// PolarDB_MySQL
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The time span between the start time and the end time cannot exceed 31 days. Specify the time in the yyyy-MM-ddZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-30Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 30 to 100. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-05-01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogsRequest) GetDBClusterId() *string {
	return s.DBClusterId
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

func (s *DescribeSlowLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSlowLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSlowLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSlowLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSlowLogsRequest) SetDBClusterId(v string) *DescribeSlowLogsRequest {
	s.DBClusterId = &v
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

func (s *DescribeSlowLogsRequest) SetRegionId(v string) *DescribeSlowLogsRequest {
	s.RegionId = &v
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

func (s *DescribeSlowLogsRequest) SetStartTime(v string) *DescribeSlowLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogsRequest) Validate() error {
	return dara.Validate(s)
}
