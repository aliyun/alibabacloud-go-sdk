// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutStartBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDbNames(v string) *PutStartBackupRequest
	GetBackupDbNames() *string
	SetBackupLevel(v string) *PutStartBackupRequest
	GetBackupLevel() *string
	SetBackupMode(v string) *PutStartBackupRequest
	GetBackupMode() *string
	SetDrdsInstanceId(v string) *PutStartBackupRequest
	GetDrdsInstanceId() *string
}

type PutStartBackupRequest struct {
	// If you need to back up data at the database level, you must specify the list of databases to be backed up, and separate multiple databases with commas (,).
	//
	// example:
	//
	// db_1, db_2
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The backup level. Valid values:
	//
	// 	- instance: instance
	//
	// 	- db: The database type.
	//
	// example:
	//
	// db
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup mode. For more information, see [backup mode](https://help.aliyun.com/document_detail/108631.html) and the valid values are as follows:
	//
	// 	- phy: fast backup
	//
	// 	- logic: Consistent backup
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
	// drds****c6vxxyzd
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s PutStartBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s PutStartBackupRequest) GoString() string {
	return s.String()
}

func (s *PutStartBackupRequest) GetBackupDbNames() *string {
	return s.BackupDbNames
}

func (s *PutStartBackupRequest) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *PutStartBackupRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *PutStartBackupRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *PutStartBackupRequest) SetBackupDbNames(v string) *PutStartBackupRequest {
	s.BackupDbNames = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupLevel(v string) *PutStartBackupRequest {
	s.BackupLevel = &v
	return s
}

func (s *PutStartBackupRequest) SetBackupMode(v string) *PutStartBackupRequest {
	s.BackupMode = &v
	return s
}

func (s *PutStartBackupRequest) SetDrdsInstanceId(v string) *PutStartBackupRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *PutStartBackupRequest) Validate() error {
	return dara.Validate(s)
}
