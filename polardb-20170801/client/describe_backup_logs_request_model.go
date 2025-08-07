// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupRegion(v string) *DescribeBackupLogsRequest
	GetBackupRegion() *string
	SetDBClusterId(v string) *DescribeBackupLogsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeBackupLogsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeBackupLogsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBackupLogsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeBackupLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupLogsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeBackupLogsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBackupLogsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeBackupLogsRequest
	GetStartTime() *string
}

type DescribeBackupLogsRequest struct {
	// The region for the backup data.
	//
	// example:
	//
	// cn-hangzhou
	BackupRegion *string `json:"BackupRegion,omitempty" xml:"BackupRegion,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-12T15:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. The value must be an integer that is larger than 0. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
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
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-02-01T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBackupLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsRequest) GetBackupRegion() *string {
	return s.BackupRegion
}

func (s *DescribeBackupLogsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeBackupLogsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBackupLogsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBackupLogsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBackupLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupLogsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBackupLogsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBackupLogsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBackupLogsRequest) SetBackupRegion(v string) *DescribeBackupLogsRequest {
	s.BackupRegion = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetDBClusterId(v string) *DescribeBackupLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetEndTime(v string) *DescribeBackupLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetOwnerAccount(v string) *DescribeBackupLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetOwnerId(v int64) *DescribeBackupLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetPageNumber(v int32) *DescribeBackupLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetPageSize(v int32) *DescribeBackupLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetResourceOwnerAccount(v string) *DescribeBackupLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetResourceOwnerId(v int64) *DescribeBackupLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetStartTime(v string) *DescribeBackupLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupLogsRequest) Validate() error {
	return dara.Validate(s)
}
