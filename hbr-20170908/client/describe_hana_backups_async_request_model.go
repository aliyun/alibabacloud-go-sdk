// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHanaBackupsAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeHanaBackupsAsyncRequest
	GetClusterId() *string
	SetDatabaseName(v string) *DescribeHanaBackupsAsyncRequest
	GetDatabaseName() *string
	SetIncludeDifferential(v bool) *DescribeHanaBackupsAsyncRequest
	GetIncludeDifferential() *bool
	SetIncludeIncremental(v bool) *DescribeHanaBackupsAsyncRequest
	GetIncludeIncremental() *bool
	SetIncludeLog(v bool) *DescribeHanaBackupsAsyncRequest
	GetIncludeLog() *bool
	SetLogPosition(v int64) *DescribeHanaBackupsAsyncRequest
	GetLogPosition() *int64
	SetMode(v string) *DescribeHanaBackupsAsyncRequest
	GetMode() *string
	SetPageNumber(v int32) *DescribeHanaBackupsAsyncRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHanaBackupsAsyncRequest
	GetPageSize() *int32
	SetRecoveryPointInTime(v int64) *DescribeHanaBackupsAsyncRequest
	GetRecoveryPointInTime() *int64
	SetResourceGroupId(v string) *DescribeHanaBackupsAsyncRequest
	GetResourceGroupId() *string
	SetSource(v string) *DescribeHanaBackupsAsyncRequest
	GetSource() *string
	SetSourceClusterId(v string) *DescribeHanaBackupsAsyncRequest
	GetSourceClusterId() *string
	SetSystemCopy(v bool) *DescribeHanaBackupsAsyncRequest
	GetSystemCopy() *bool
	SetUseBackint(v bool) *DescribeHanaBackupsAsyncRequest
	GetUseBackint() *bool
	SetVaultId(v string) *DescribeHanaBackupsAsyncRequest
	GetVaultId() *string
	SetVolumeId(v int32) *DescribeHanaBackupsAsyncRequest
	GetVolumeId() *int32
}

type DescribeHanaBackupsAsyncRequest struct {
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-00098******yuqvu
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// BPD
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Specifies whether to include differential backups in the query results. Valid values:
	//
	// 	- true: includes differential backups.
	//
	// 	- false: excludes differential backups.
	//
	// example:
	//
	// false
	IncludeDifferential *bool `json:"IncludeDifferential,omitempty" xml:"IncludeDifferential,omitempty"`
	// Specifies whether to include incremental backups in the query results. Valid values:
	//
	// 	- true: includes incremental backups.
	//
	// 	- false: excludes incremental backups.
	//
	// example:
	//
	// true
	IncludeIncremental *bool `json:"IncludeIncremental,omitempty" xml:"IncludeIncremental,omitempty"`
	// Specifies whether to include log backups in the query results. Valid values:
	//
	// 	- true: includes log backups.
	//
	// 	- false: excludes log backups.
	//
	// example:
	//
	// true
	IncludeLog *bool `json:"IncludeLog,omitempty" xml:"IncludeLog,omitempty"`
	// The log position to which you want to restore the database. This parameter is valid only if you set the Mode parameter to **RECOVERY_TO_LOG_POSITION**.
	//
	// example:
	//
	// 0
	LogPosition *int64 `json:"LogPosition,omitempty" xml:"LogPosition,omitempty"`
	// The recovery mode. Valid values:
	//
	// 	- **RECOVERY_TO_MOST_RECENT**: restores the database to the recently available state to which the database has been backed up.
	//
	// 	- **RECOVERY_TO_POINT_IN_TIME**: restores the database to a specified point in time.
	//
	// 	- **RECOVERY_TO_SPECIFIC_BACKUP**: restores the database to a specified backup.
	//
	// 	- **RECOVERY_TO_LOG_POSITION**: restores the database to a specified log position.
	//
	// example:
	//
	// RECOVERY_TO_SPECIFIC_BACKUP
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The point in time to which you want to restore the database. This parameter is valid only if you set the Mode parameter to **RECOVERY_TO_POINT_IN_TIME**. Cloud Backup restores the database to a state closest to the specified point in time.
	//
	// example:
	//
	// 1649851200
	RecoveryPointInTime *int64 `json:"RecoveryPointInTime,omitempty" xml:"RecoveryPointInTime,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmz7mced2ldhy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the source system. This parameter specifies the name of the source database that you want to restore. You must set the parameter in the `<Source database name>@SID` format.
	//
	// example:
	//
	// P01@HP1
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the source SAP HANA instance.
	//
	// example:
	//
	// cl-0000g3m******5cj
	SourceClusterId *string `json:"SourceClusterId,omitempty" xml:"SourceClusterId,omitempty"`
	// Specifies whether to restore the database to a different instance.
	//
	// 	- true: restores the database to a different instance.
	//
	// 	- false: restores the database within the same instance.
	//
	// example:
	//
	// true
	SystemCopy *bool `json:"SystemCopy,omitempty" xml:"SystemCopy,omitempty"`
	// Specifies whether Backint is used. Valid values:
	//
	// 	- true: Backint is used.
	//
	// 	- false: Backint is not used.
	//
	// example:
	//
	// false
	UseBackint *bool `json:"UseBackint,omitempty" xml:"UseBackint,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000270c******pi81
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	// The ID of the volume that you want to restore. This parameter is valid only if you set the Mode parameter to **RECOVERY_TO_LOG_POSITION**.
	//
	// example:
	//
	// 0
	VolumeId *int32 `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s DescribeHanaBackupsAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHanaBackupsAsyncRequest) GoString() string {
	return s.String()
}

func (s *DescribeHanaBackupsAsyncRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeHanaBackupsAsyncRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeHanaBackupsAsyncRequest) GetIncludeDifferential() *bool {
	return s.IncludeDifferential
}

func (s *DescribeHanaBackupsAsyncRequest) GetIncludeIncremental() *bool {
	return s.IncludeIncremental
}

func (s *DescribeHanaBackupsAsyncRequest) GetIncludeLog() *bool {
	return s.IncludeLog
}

func (s *DescribeHanaBackupsAsyncRequest) GetLogPosition() *int64 {
	return s.LogPosition
}

func (s *DescribeHanaBackupsAsyncRequest) GetMode() *string {
	return s.Mode
}

func (s *DescribeHanaBackupsAsyncRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHanaBackupsAsyncRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHanaBackupsAsyncRequest) GetRecoveryPointInTime() *int64 {
	return s.RecoveryPointInTime
}

func (s *DescribeHanaBackupsAsyncRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeHanaBackupsAsyncRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeHanaBackupsAsyncRequest) GetSourceClusterId() *string {
	return s.SourceClusterId
}

func (s *DescribeHanaBackupsAsyncRequest) GetSystemCopy() *bool {
	return s.SystemCopy
}

func (s *DescribeHanaBackupsAsyncRequest) GetUseBackint() *bool {
	return s.UseBackint
}

func (s *DescribeHanaBackupsAsyncRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeHanaBackupsAsyncRequest) GetVolumeId() *int32 {
	return s.VolumeId
}

func (s *DescribeHanaBackupsAsyncRequest) SetClusterId(v string) *DescribeHanaBackupsAsyncRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetDatabaseName(v string) *DescribeHanaBackupsAsyncRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetIncludeDifferential(v bool) *DescribeHanaBackupsAsyncRequest {
	s.IncludeDifferential = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetIncludeIncremental(v bool) *DescribeHanaBackupsAsyncRequest {
	s.IncludeIncremental = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetIncludeLog(v bool) *DescribeHanaBackupsAsyncRequest {
	s.IncludeLog = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetLogPosition(v int64) *DescribeHanaBackupsAsyncRequest {
	s.LogPosition = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetMode(v string) *DescribeHanaBackupsAsyncRequest {
	s.Mode = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetPageNumber(v int32) *DescribeHanaBackupsAsyncRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetPageSize(v int32) *DescribeHanaBackupsAsyncRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetRecoveryPointInTime(v int64) *DescribeHanaBackupsAsyncRequest {
	s.RecoveryPointInTime = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetResourceGroupId(v string) *DescribeHanaBackupsAsyncRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetSource(v string) *DescribeHanaBackupsAsyncRequest {
	s.Source = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetSourceClusterId(v string) *DescribeHanaBackupsAsyncRequest {
	s.SourceClusterId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetSystemCopy(v bool) *DescribeHanaBackupsAsyncRequest {
	s.SystemCopy = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetUseBackint(v bool) *DescribeHanaBackupsAsyncRequest {
	s.UseBackint = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetVaultId(v string) *DescribeHanaBackupsAsyncRequest {
	s.VaultId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) SetVolumeId(v int32) *DescribeHanaBackupsAsyncRequest {
	s.VolumeId = &v
	return s
}

func (s *DescribeHanaBackupsAsyncRequest) Validate() error {
	return dara.Validate(s)
}
