// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDbNames(v string) *SetBackupPolicyRequest
	GetBackupDbNames() *string
	SetBackupLevel(v string) *SetBackupPolicyRequest
	GetBackupLevel() *string
	SetBackupLog(v string) *SetBackupPolicyRequest
	GetBackupLog() *string
	SetBackupMode(v string) *SetBackupPolicyRequest
	GetBackupMode() *string
	SetDataBackupRetentionPeriod(v string) *SetBackupPolicyRequest
	GetDataBackupRetentionPeriod() *string
	SetDrdsInstanceId(v string) *SetBackupPolicyRequest
	GetDrdsInstanceId() *string
	SetLogBackupRetentionPeriod(v string) *SetBackupPolicyRequest
	GetLogBackupRetentionPeriod() *string
	SetPreferredBackupEndTime(v string) *SetBackupPolicyRequest
	GetPreferredBackupEndTime() *string
	SetPreferredBackupPeriod(v string) *SetBackupPolicyRequest
	GetPreferredBackupPeriod() *string
	SetPreferredBackupStartTime(v string) *SetBackupPolicyRequest
	GetPreferredBackupStartTime() *string
}

type SetBackupPolicyRequest struct {
	// The databases to be backed up. Separate multiple databases with commas (,).
	//
	// >  This parameter takes effect only when the backup level is database level.
	//
	// example:
	//
	// test1,test2
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The level of the backup. Valid values:
	//
	// 	- db: The database type.
	//
	// 	- instance: instance
	//
	// example:
	//
	// db
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// Specifies whether to enable log Backup. Valid values:
	//
	// 	- 1: enabled.
	//
	// 	- 0: disabled.
	//
	// example:
	//
	// 1
	BackupLog *string `json:"BackupLog,omitempty" xml:"BackupLog,omitempty"`
	// The backup mode. Valid values:
	//
	// 	- **Logic **: logical backup
	//
	// 	- **phy**: physical backup
	//
	// example:
	//
	// phy
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The retention period of the backup data. Value range: 7 to 730.
	//
	// example:
	//
	// 7
	DataBackupRetentionPeriod *string `json:"DataBackupRetentionPeriod,omitempty" xml:"DataBackupRetentionPeriod,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The log retention period. Valid values: 7 to 730. This value must be less than or equal to the number of data backup days.
	//
	// example:
	//
	// 7
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// The end time of the backup.
	//
	// example:
	//
	// 12:30:30
	PreferredBackupEndTime *string `json:"PreferredBackupEndTime,omitempty" xml:"PreferredBackupEndTime,omitempty"`
	// The backup cycle. Valid values:
	//
	// 	- 0: Monday
	//
	// 	- 1: Tuesday
	//
	// 	- 2: Wednesday
	//
	// 	- 3: Thursday
	//
	// 	- 4: Friday
	//
	// 	- 5: Saturday
	//
	// 	- 6: Sunday
	//
	// example:
	//
	// 0
	PreferredBackupPeriod *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	// The start time of the backup.
	//
	// example:
	//
	// 11:30:30
	PreferredBackupStartTime *string `json:"PreferredBackupStartTime,omitempty" xml:"PreferredBackupStartTime,omitempty"`
}

func (s SetBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetBackupPolicyRequest) GetBackupDbNames() *string {
	return s.BackupDbNames
}

func (s *SetBackupPolicyRequest) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *SetBackupPolicyRequest) GetBackupLog() *string {
	return s.BackupLog
}

func (s *SetBackupPolicyRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *SetBackupPolicyRequest) GetDataBackupRetentionPeriod() *string {
	return s.DataBackupRetentionPeriod
}

func (s *SetBackupPolicyRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SetBackupPolicyRequest) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *SetBackupPolicyRequest) GetPreferredBackupEndTime() *string {
	return s.PreferredBackupEndTime
}

func (s *SetBackupPolicyRequest) GetPreferredBackupPeriod() *string {
	return s.PreferredBackupPeriod
}

func (s *SetBackupPolicyRequest) GetPreferredBackupStartTime() *string {
	return s.PreferredBackupStartTime
}

func (s *SetBackupPolicyRequest) SetBackupDbNames(v string) *SetBackupPolicyRequest {
	s.BackupDbNames = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLevel(v string) *SetBackupPolicyRequest {
	s.BackupLevel = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupLog(v string) *SetBackupPolicyRequest {
	s.BackupLog = &v
	return s
}

func (s *SetBackupPolicyRequest) SetBackupMode(v string) *SetBackupPolicyRequest {
	s.BackupMode = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDataBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.DataBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetDrdsInstanceId(v string) *SetBackupPolicyRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *SetBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupEndTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupEndTime = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupPeriod(v string) *SetBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *SetBackupPolicyRequest) SetPreferredBackupStartTime(v string) *SetBackupPolicyRequest {
	s.PreferredBackupStartTime = &v
	return s
}

func (s *SetBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
