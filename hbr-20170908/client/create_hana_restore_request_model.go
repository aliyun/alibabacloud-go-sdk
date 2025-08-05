// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaRestoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupId(v int64) *CreateHanaRestoreRequest
	GetBackupId() *int64
	SetBackupPrefix(v string) *CreateHanaRestoreRequest
	GetBackupPrefix() *string
	SetCheckAccess(v bool) *CreateHanaRestoreRequest
	GetCheckAccess() *bool
	SetClearLog(v bool) *CreateHanaRestoreRequest
	GetClearLog() *bool
	SetClusterId(v string) *CreateHanaRestoreRequest
	GetClusterId() *string
	SetDatabaseName(v string) *CreateHanaRestoreRequest
	GetDatabaseName() *string
	SetLogPosition(v int64) *CreateHanaRestoreRequest
	GetLogPosition() *int64
	SetMasterClientId(v string) *CreateHanaRestoreRequest
	GetMasterClientId() *string
	SetMode(v string) *CreateHanaRestoreRequest
	GetMode() *string
	SetRecoveryPointInTime(v int64) *CreateHanaRestoreRequest
	GetRecoveryPointInTime() *int64
	SetSidAdmin(v string) *CreateHanaRestoreRequest
	GetSidAdmin() *string
	SetSource(v string) *CreateHanaRestoreRequest
	GetSource() *string
	SetSourceClusterId(v string) *CreateHanaRestoreRequest
	GetSourceClusterId() *string
	SetSystemCopy(v bool) *CreateHanaRestoreRequest
	GetSystemCopy() *bool
	SetUseCatalog(v bool) *CreateHanaRestoreRequest
	GetUseCatalog() *bool
	SetUseDelta(v bool) *CreateHanaRestoreRequest
	GetUseDelta() *bool
	SetVaultId(v string) *CreateHanaRestoreRequest
	GetVaultId() *string
	SetVolumeId(v int32) *CreateHanaRestoreRequest
	GetVolumeId() *int32
}

type CreateHanaRestoreRequest struct {
	// The ID of the backup.
	//
	// example:
	//
	// 1645628400235
	BackupId *int64 `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The backup prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// COMPLETE_DATA_BACKUP_2022_05_02_15_39
	BackupPrefix *string `json:"BackupPrefix,omitempty" xml:"BackupPrefix,omitempty"`
	// Specifies whether to validate the differential backup and log backup. Valid values: true and false. If you set the value to true, HBR checks whether the required differential backup and log backup are available before the restore job starts. If the differential backup or log backup is unavailable, HBR does not start the restore job.
	//
	// example:
	//
	// false
	CheckAccess *bool `json:"CheckAccess,omitempty" xml:"CheckAccess,omitempty"`
	// Specifies whether to delete all log entries from the log area after the log entries are restored. Valid values: true and false. If you set the value to false, all log entries are deleted from the log area after the log entries are restored.
	//
	// example:
	//
	// false
	ClearLog *bool `json:"ClearLog,omitempty" xml:"ClearLog,omitempty"`
	// The ID of the SAP HANA instance that you want to restore.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000fbrs5******ka9w
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the database that you want to restore.
	//
	// This parameter is required.
	//
	// example:
	//
	// TS2
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The log position to which you want to restore the database. This parameter is valid only if you set the Mode parameter to **RECOVERY_TO_LOG_POSITION**.
	//
	// example:
	//
	// 0
	LogPosition *int64 `json:"LogPosition,omitempty" xml:"LogPosition,omitempty"`
	// The ID of the client where the primary node of the SAP HANA resides.
	//
	// example:
	//
	// c-000ii8tzv**********
	MasterClientId *string `json:"MasterClientId,omitempty" xml:"MasterClientId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// RECOVERY_TO_POINT_IN_TIME
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The point in time to which you want to restore the database. This parameter is valid only if you set the Mode parameter to **RECOVERY_TO_POINT_IN_TIME**. HBR restores the database to a state closest to the specified point in time.
	//
	// example:
	//
	// 1635315505
	RecoveryPointInTime *int64 `json:"RecoveryPointInTime,omitempty" xml:"RecoveryPointInTime,omitempty"`
	// The SID admin account that is created by SAP HANA.
	//
	// example:
	//
	// DB
	SidAdmin *string `json:"SidAdmin,omitempty" xml:"SidAdmin,omitempty"`
	// The name of the source system. This parameter specifies the name of the source database that you want to restore. You must set the parameter in the `<Source database name>@SID` format.
	//
	// example:
	//
	// HNP@HNP
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the source SAP HANA instance.
	//
	// example:
	//
	// cl-000ii8tzv******xm0t
	SourceClusterId *string `json:"SourceClusterId,omitempty" xml:"SourceClusterId,omitempty"`
	// Specifies whether to restore the database to a different instance.
	//
	// example:
	//
	// false
	SystemCopy *bool `json:"SystemCopy,omitempty" xml:"SystemCopy,omitempty"`
	// Specifies whether to use a catalog backup to restore the database. This parameter is required only if you set the Mode parameter to **RECOVERY_TO_SPECIFIC_BACKUP**. If you turn off Use Catalog, you must specify the prefix of a backup file. Then, Cloud Backup finds the backup file based on the specified prefix and restores the backup file.
	//
	// example:
	//
	// false
	UseCatalog *bool `json:"UseCatalog,omitempty" xml:"UseCatalog,omitempty"`
	// Specifies whether to use a differential backup or an incremental backup to restore the database. Valid values: true and false. If you want to use a differential backup or an incremental backup to restore the database, set the value to true. If you set the value to false, HBR uses a log backup to restore the database.
	//
	// example:
	//
	// true
	UseDelta *bool `json:"UseDelta,omitempty" xml:"UseDelta,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-000************yqr
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	// The ID of the volume that you want to restore. This parameter is valid only if you set the Mode parameter to **RECOVERY_TO_LOG_POSITION**.
	//
	// example:
	//
	// 0
	VolumeId *int32 `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
}

func (s CreateHanaRestoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaRestoreRequest) GoString() string {
	return s.String()
}

func (s *CreateHanaRestoreRequest) GetBackupId() *int64 {
	return s.BackupId
}

func (s *CreateHanaRestoreRequest) GetBackupPrefix() *string {
	return s.BackupPrefix
}

func (s *CreateHanaRestoreRequest) GetCheckAccess() *bool {
	return s.CheckAccess
}

func (s *CreateHanaRestoreRequest) GetClearLog() *bool {
	return s.ClearLog
}

func (s *CreateHanaRestoreRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateHanaRestoreRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CreateHanaRestoreRequest) GetLogPosition() *int64 {
	return s.LogPosition
}

func (s *CreateHanaRestoreRequest) GetMasterClientId() *string {
	return s.MasterClientId
}

func (s *CreateHanaRestoreRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateHanaRestoreRequest) GetRecoveryPointInTime() *int64 {
	return s.RecoveryPointInTime
}

func (s *CreateHanaRestoreRequest) GetSidAdmin() *string {
	return s.SidAdmin
}

func (s *CreateHanaRestoreRequest) GetSource() *string {
	return s.Source
}

func (s *CreateHanaRestoreRequest) GetSourceClusterId() *string {
	return s.SourceClusterId
}

func (s *CreateHanaRestoreRequest) GetSystemCopy() *bool {
	return s.SystemCopy
}

func (s *CreateHanaRestoreRequest) GetUseCatalog() *bool {
	return s.UseCatalog
}

func (s *CreateHanaRestoreRequest) GetUseDelta() *bool {
	return s.UseDelta
}

func (s *CreateHanaRestoreRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateHanaRestoreRequest) GetVolumeId() *int32 {
	return s.VolumeId
}

func (s *CreateHanaRestoreRequest) SetBackupId(v int64) *CreateHanaRestoreRequest {
	s.BackupId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetBackupPrefix(v string) *CreateHanaRestoreRequest {
	s.BackupPrefix = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetCheckAccess(v bool) *CreateHanaRestoreRequest {
	s.CheckAccess = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetClearLog(v bool) *CreateHanaRestoreRequest {
	s.ClearLog = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetClusterId(v string) *CreateHanaRestoreRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetDatabaseName(v string) *CreateHanaRestoreRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetLogPosition(v int64) *CreateHanaRestoreRequest {
	s.LogPosition = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetMasterClientId(v string) *CreateHanaRestoreRequest {
	s.MasterClientId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetMode(v string) *CreateHanaRestoreRequest {
	s.Mode = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetRecoveryPointInTime(v int64) *CreateHanaRestoreRequest {
	s.RecoveryPointInTime = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSidAdmin(v string) *CreateHanaRestoreRequest {
	s.SidAdmin = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSource(v string) *CreateHanaRestoreRequest {
	s.Source = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSourceClusterId(v string) *CreateHanaRestoreRequest {
	s.SourceClusterId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetSystemCopy(v bool) *CreateHanaRestoreRequest {
	s.SystemCopy = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetUseCatalog(v bool) *CreateHanaRestoreRequest {
	s.UseCatalog = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetUseDelta(v bool) *CreateHanaRestoreRequest {
	s.UseDelta = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetVaultId(v string) *CreateHanaRestoreRequest {
	s.VaultId = &v
	return s
}

func (s *CreateHanaRestoreRequest) SetVolumeId(v int32) *CreateHanaRestoreRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateHanaRestoreRequest) Validate() error {
	return dara.Validate(s)
}
