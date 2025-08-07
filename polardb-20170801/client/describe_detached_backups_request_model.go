// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetachedBackupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v string) *DescribeDetachedBackupsRequest
	GetBackupId() *string
	SetBackupMode(v string) *DescribeDetachedBackupsRequest
	GetBackupMode() *string
	SetBackupRegion(v string) *DescribeDetachedBackupsRequest
	GetBackupRegion() *string
	SetBackupStatus(v string) *DescribeDetachedBackupsRequest
	GetBackupStatus() *string
	SetDBClusterId(v string) *DescribeDetachedBackupsRequest
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeDetachedBackupsRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeDetachedBackupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDetachedBackupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDetachedBackupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDetachedBackupsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *DescribeDetachedBackupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDetachedBackupsRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDetachedBackupsRequest
	GetStartTime() *string
}

type DescribeDetachedBackupsRequest struct {
	// The ID of the backup set.
	//
	// example:
	//
	// 111111111
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Automated**
	//
	// 	- **Manual**
	//
	// example:
	//
	// Manual
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The region where the cross-region data backup file of the instance is stored.
	//
	// > This parameter is valid only for PolarDB for MySQL clusters.
	//
	// example:
	//
	// cn-hangzhou
	BackupRegion *string `json:"BackupRegion,omitempty" xml:"BackupRegion,omitempty"`
	// The status of the backup set. Valid values:
	//
	// 	- **Success**
	//
	// 	- **Failed**
	//
	// example:
	//
	// Success
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end of the time range to query. Specify the time in the `YYYY-MM-DDThh:mmZ` format. The time must be in UTC. The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-14T00:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the INTEGER data type. Default value: **1**.
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
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the `yyyy-MM-ddTHH:mmZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-05-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDetachedBackupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeDetachedBackupsRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeDetachedBackupsRequest) GetBackupRegion() *string {
	return s.BackupRegion
}

func (s *DescribeDetachedBackupsRequest) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeDetachedBackupsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDetachedBackupsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDetachedBackupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDetachedBackupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDetachedBackupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDetachedBackupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDetachedBackupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDetachedBackupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDetachedBackupsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDetachedBackupsRequest) SetBackupId(v string) *DescribeDetachedBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupMode(v string) *DescribeDetachedBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupRegion(v string) *DescribeDetachedBackupsRequest {
	s.BackupRegion = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupStatus(v string) *DescribeDetachedBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetDBClusterId(v string) *DescribeDetachedBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetEndTime(v string) *DescribeDetachedBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetOwnerAccount(v string) *DescribeDetachedBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageNumber(v int32) *DescribeDetachedBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageSize(v int32) *DescribeDetachedBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerAccount(v string) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetStartTime(v string) *DescribeDetachedBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) Validate() error {
	return dara.Validate(s)
}
