// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaBackupSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogBackupParameterFile(v string) *UpdateHanaBackupSettingRequest
	GetCatalogBackupParameterFile() *string
	SetCatalogBackupUsingBackint(v bool) *UpdateHanaBackupSettingRequest
	GetCatalogBackupUsingBackint() *bool
	SetClusterId(v string) *UpdateHanaBackupSettingRequest
	GetClusterId() *string
	SetDataBackupParameterFile(v string) *UpdateHanaBackupSettingRequest
	GetDataBackupParameterFile() *string
	SetDatabaseName(v string) *UpdateHanaBackupSettingRequest
	GetDatabaseName() *string
	SetEnableAutoLogBackup(v bool) *UpdateHanaBackupSettingRequest
	GetEnableAutoLogBackup() *bool
	SetLogBackupParameterFile(v string) *UpdateHanaBackupSettingRequest
	GetLogBackupParameterFile() *string
	SetLogBackupTimeout(v int64) *UpdateHanaBackupSettingRequest
	GetLogBackupTimeout() *int64
	SetLogBackupUsingBackint(v bool) *UpdateHanaBackupSettingRequest
	GetLogBackupUsingBackint() *bool
	SetVaultId(v string) *UpdateHanaBackupSettingRequest
	GetVaultId() *string
}

type UpdateHanaBackupSettingRequest struct {
	// The configuration file for catalog backup.
	//
	// example:
	//
	// /usr/sap/SID/SYS/global/hdb/opt/hdbconfig/param
	CatalogBackupParameterFile *string `json:"CatalogBackupParameterFile,omitempty" xml:"CatalogBackupParameterFile,omitempty"`
	// Specifies whether to use Backint to back up catalogs. Valid values:
	//
	// 	- true: Backint is used to back up catalogs.
	//
	// 	- false: Backint is not used to back up catalogs.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CatalogBackupUsingBackint *bool `json:"CatalogBackupUsingBackint,omitempty" xml:"CatalogBackupUsingBackint,omitempty"`
	// The ID of the SAP HANA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cl-000ezvg******8znz
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configuration file for data backup.
	//
	// example:
	//
	// /usr/sap/SID/SYS/global/hdb/opt/hdbconfig/param
	DataBackupParameterFile *string `json:"DataBackupParameterFile,omitempty" xml:"DataBackupParameterFile,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// SYSTEMDB
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// Specifies whether to enable automatic log backup. Valid values:
	//
	// 	- **true**: enables automatic log backup.
	//
	// 	- **false**: disables automatic log backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableAutoLogBackup *bool `json:"EnableAutoLogBackup,omitempty" xml:"EnableAutoLogBackup,omitempty"`
	// The configuration file for log backup.
	//
	// example:
	//
	// /usr/sap/SID/SYS/global/hdb/opt/hdbconfig/param
	LogBackupParameterFile *string `json:"LogBackupParameterFile,omitempty" xml:"LogBackupParameterFile,omitempty"`
	// The interval at which logs are backed up. Unit: seconds.
	//
	// example:
	//
	// 900
	LogBackupTimeout *int64 `json:"LogBackupTimeout,omitempty" xml:"LogBackupTimeout,omitempty"`
	// Specifies whether to use Backint to back up logs. Valid values:
	//
	// 	- true: Backint is used to back up logs.
	//
	// 	- false: Backint is not used to back up logs.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	LogBackupUsingBackint *bool `json:"LogBackupUsingBackint,omitempty" xml:"LogBackupUsingBackint,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0005n******rluw5
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateHanaBackupSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaBackupSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupSettingRequest) GetCatalogBackupParameterFile() *string {
	return s.CatalogBackupParameterFile
}

func (s *UpdateHanaBackupSettingRequest) GetCatalogBackupUsingBackint() *bool {
	return s.CatalogBackupUsingBackint
}

func (s *UpdateHanaBackupSettingRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateHanaBackupSettingRequest) GetDataBackupParameterFile() *string {
	return s.DataBackupParameterFile
}

func (s *UpdateHanaBackupSettingRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *UpdateHanaBackupSettingRequest) GetEnableAutoLogBackup() *bool {
	return s.EnableAutoLogBackup
}

func (s *UpdateHanaBackupSettingRequest) GetLogBackupParameterFile() *string {
	return s.LogBackupParameterFile
}

func (s *UpdateHanaBackupSettingRequest) GetLogBackupTimeout() *int64 {
	return s.LogBackupTimeout
}

func (s *UpdateHanaBackupSettingRequest) GetLogBackupUsingBackint() *bool {
	return s.LogBackupUsingBackint
}

func (s *UpdateHanaBackupSettingRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateHanaBackupSettingRequest) SetCatalogBackupParameterFile(v string) *UpdateHanaBackupSettingRequest {
	s.CatalogBackupParameterFile = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetCatalogBackupUsingBackint(v bool) *UpdateHanaBackupSettingRequest {
	s.CatalogBackupUsingBackint = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetClusterId(v string) *UpdateHanaBackupSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetDataBackupParameterFile(v string) *UpdateHanaBackupSettingRequest {
	s.DataBackupParameterFile = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetDatabaseName(v string) *UpdateHanaBackupSettingRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetEnableAutoLogBackup(v bool) *UpdateHanaBackupSettingRequest {
	s.EnableAutoLogBackup = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetLogBackupParameterFile(v string) *UpdateHanaBackupSettingRequest {
	s.LogBackupParameterFile = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetLogBackupTimeout(v int64) *UpdateHanaBackupSettingRequest {
	s.LogBackupTimeout = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetLogBackupUsingBackint(v bool) *UpdateHanaBackupSettingRequest {
	s.LogBackupUsingBackint = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) SetVaultId(v string) *UpdateHanaBackupSettingRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateHanaBackupSettingRequest) Validate() error {
	return dara.Validate(s)
}
