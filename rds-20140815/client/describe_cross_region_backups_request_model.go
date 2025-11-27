// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int32) *DescribeCrossRegionBackupsRequest
	GetBackupId() *int32
	SetCrossBackupId(v int32) *DescribeCrossRegionBackupsRequest
	GetCrossBackupId() *int32
	SetCrossBackupRegion(v string) *DescribeCrossRegionBackupsRequest
	GetCrossBackupRegion() *string
	SetDBInstanceId(v string) *DescribeCrossRegionBackupsRequest
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeCrossRegionBackupsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeCrossRegionBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeCrossRegionBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossRegionBackupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeCrossRegionBackupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCrossRegionBackupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCrossRegionBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCrossRegionBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeCrossRegionBackupsRequest
	GetStartTime() *string
}

type DescribeCrossRegionBackupsRequest struct {
	// The ID of the backup file.
	//
	// example:
	//
	// 603524***
	BackupId *int32 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The ID of the cross-region data backup file.
	//
	// >  You must specify the **CrossBackupId*	- parameter. Alternatively, you must specify the **StartTime*	- and **EndTime*	- parameters.
	//
	// example:
	//
	// 14562
	CrossBackupId *int32 `json:"CrossBackupId,omitempty" xml:"CrossBackupId,omitempty"`
	// The ID of the region in which the cross-region data backup file is stored.
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
	// The number of entries per page. Valid values:
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2019-05-30T12:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCrossRegionBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupsRequest) GetBackupId() *int32 {
	return s.BackupId
}

func (s *DescribeCrossRegionBackupsRequest) GetCrossBackupId() *int32 {
	return s.CrossBackupId
}

func (s *DescribeCrossRegionBackupsRequest) GetCrossBackupRegion() *string {
	return s.CrossBackupRegion
}

func (s *DescribeCrossRegionBackupsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeCrossRegionBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCrossRegionBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossRegionBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossRegionBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossRegionBackupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCrossRegionBackupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCrossRegionBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCrossRegionBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCrossRegionBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCrossRegionBackupsRequest) SetBackupId(v int32) *DescribeCrossRegionBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetCrossBackupId(v int32) *DescribeCrossRegionBackupsRequest {
	s.CrossBackupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetCrossBackupRegion(v string) *DescribeCrossRegionBackupsRequest {
	s.CrossBackupRegion = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetDBInstanceId(v string) *DescribeCrossRegionBackupsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetEndTime(v string) *DescribeCrossRegionBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetOwnerId(v int64) *DescribeCrossRegionBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetPageNumber(v int32) *DescribeCrossRegionBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetPageSize(v int32) *DescribeCrossRegionBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetRegionId(v string) *DescribeCrossRegionBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetResourceGroupId(v string) *DescribeCrossRegionBackupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetResourceOwnerAccount(v string) *DescribeCrossRegionBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetResourceOwnerId(v int64) *DescribeCrossRegionBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) SetStartTime(v string) *DescribeCrossRegionBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCrossRegionBackupsRequest) Validate() error {
	return dara.Validate(s)
}
