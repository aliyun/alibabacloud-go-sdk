// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRestoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDbNames(v string) *StartRestoreRequest
	GetBackupDbNames() *string
	SetBackupId(v string) *StartRestoreRequest
	GetBackupId() *string
	SetBackupLevel(v string) *StartRestoreRequest
	GetBackupLevel() *string
	SetBackupMode(v string) *StartRestoreRequest
	GetBackupMode() *string
	SetDrdsInstanceId(v string) *StartRestoreRequest
	GetDrdsInstanceId() *string
	SetPreferredBackupTime(v string) *StartRestoreRequest
	GetPreferredBackupTime() *string
}

type StartRestoreRequest struct {
	// The name of the database to be restored. Separate multiple databases with commas (,).
	//
	// >  If you do not specify any database name, all databases in the instance are restored by default.
	//
	// example:
	//
	// test1,test2
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The ID of the DRDS backup set.
	//
	// >  If you do not specify this parameter, the system restores data by time (PreferredBackupTime).
	//
	// example:
	//
	// 23***
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The level of the backup. Valid values:
	//
	// 	- db: The database level.
	//
	// 	- instance: the instance level.
	//
	// example:
	//
	// db
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup method. Valid values:
	//
	// 	- logic: the logical backup.
	//
	// 	- phy: fast backup
	//
	// example:
	//
	// phy
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The restoration time of the instance, in the format of`  yyyy-MM-dd HH:mm:ss `.
	//
	// example:
	//
	// 2019-09-10 20:18:18
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s StartRestoreRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRestoreRequest) GoString() string {
	return s.String()
}

func (s *StartRestoreRequest) GetBackupDbNames() *string {
	return s.BackupDbNames
}

func (s *StartRestoreRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *StartRestoreRequest) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *StartRestoreRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *StartRestoreRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *StartRestoreRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *StartRestoreRequest) SetBackupDbNames(v string) *StartRestoreRequest {
	s.BackupDbNames = &v
	return s
}

func (s *StartRestoreRequest) SetBackupId(v string) *StartRestoreRequest {
	s.BackupId = &v
	return s
}

func (s *StartRestoreRequest) SetBackupLevel(v string) *StartRestoreRequest {
	s.BackupLevel = &v
	return s
}

func (s *StartRestoreRequest) SetBackupMode(v string) *StartRestoreRequest {
	s.BackupMode = &v
	return s
}

func (s *StartRestoreRequest) SetDrdsInstanceId(v string) *StartRestoreRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *StartRestoreRequest) SetPreferredBackupTime(v string) *StartRestoreRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *StartRestoreRequest) Validate() error {
	return dara.Validate(s)
}
