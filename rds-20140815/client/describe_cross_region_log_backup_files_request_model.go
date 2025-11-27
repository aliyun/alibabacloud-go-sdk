// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionLogBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossBackupRegion(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetCrossBackupRegion() *string
	SetDBInstanceId(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCrossRegionLogBackupFilesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossRegionLogBackupFilesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeCrossRegionLogBackupFilesRequest
	GetStartTime() *string
}

type DescribeCrossRegionLogBackupFilesRequest struct {
	// The ID of the destination region within which the cross-region backup file is stored. You can call the DescribeCrossRegionBackupDBInstance operation to query the region ID.
	//
	// example:
	//
	// cn-shanghai
	CrossBackupRegion *string `json:"CrossBackupRegion,omitempty" xml:"CrossBackupRegion,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-15T12:10:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any non-zero positive integer.
	//
	// Default value: **1**.
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
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-30T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCrossRegionLogBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionLogBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCrossRegionLogBackupFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetCrossBackupRegion(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetDBInstanceId(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetEndTime(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetPageNumber(v int32) *DescribeCrossRegionLogBackupFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetPageSize(v int32) *DescribeCrossRegionLogBackupFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetRegionId(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetResourceOwnerAccount(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetResourceOwnerId(v int64) *DescribeCrossRegionLogBackupFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) SetStartTime(v string) *DescribeCrossRegionLogBackupFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}
