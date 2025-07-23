// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigBackupTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupHourInDay(v int64) *ConfigBackupTaskRequest
	GetBackupHourInDay() *int64
	SetBackupId(v string) *ConfigBackupTaskRequest
	GetBackupId() *string
	SetBackupPeriod(v int64) *ConfigBackupTaskRequest
	GetBackupPeriod() *int64
	SetManual2PeriodicList(v []*string) *ConfigBackupTaskRequest
	GetManual2PeriodicList() []*string
	SetPeriodic2ManualList(v []*string) *ConfigBackupTaskRequest
	GetPeriodic2ManualList() []*string
}

type ConfigBackupTaskRequest struct {
	// The backup time in the 24-hour format. Valid values: 1 to 24.
	//
	// Enumeration values:
	//
	// 	- 0
	//
	// 	- 1
	//
	// 	- 2
	//
	// 	- 3
	//
	// 	- 4
	//
	// 	- 5
	//
	// 	- 6
	//
	// 	- 7
	//
	// 	- 8
	//
	// 	- 9
	//
	// 	- 10
	//
	// 	- 11
	//
	// 	- 12
	//
	// 	- 13
	//
	// 	- 14
	//
	// 	- 15
	//
	// 	- 16
	//
	// 	- 17
	//
	// 	- 18
	//
	// 	- 19
	//
	// 	- 20
	//
	// 	- 21
	//
	// 	- 22
	//
	// 	- 23
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	BackupHourInDay *int64 `json:"BackupHourInDay,omitempty" xml:"BackupHourInDay,omitempty"`
	// The ID of the backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// backup-173620705****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The automatic backup cycle. Unit: days. Valid values: 1, 3, 7, and 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	BackupPeriod *int64 `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The IDs of images for which the manual backup mode is updated to the automatic backup mode.
	Manual2PeriodicList []*string `json:"Manual2PeriodicList,omitempty" xml:"Manual2PeriodicList,omitempty" type:"Repeated"`
	// The IDs of images for which the automatic backup mode is updated to the manual backup mode.
	Periodic2ManualList []*string `json:"Periodic2ManualList,omitempty" xml:"Periodic2ManualList,omitempty" type:"Repeated"`
}

func (s ConfigBackupTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigBackupTaskRequest) GoString() string {
	return s.String()
}

func (s *ConfigBackupTaskRequest) GetBackupHourInDay() *int64 {
	return s.BackupHourInDay
}

func (s *ConfigBackupTaskRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *ConfigBackupTaskRequest) GetBackupPeriod() *int64 {
	return s.BackupPeriod
}

func (s *ConfigBackupTaskRequest) GetManual2PeriodicList() []*string {
	return s.Manual2PeriodicList
}

func (s *ConfigBackupTaskRequest) GetPeriodic2ManualList() []*string {
	return s.Periodic2ManualList
}

func (s *ConfigBackupTaskRequest) SetBackupHourInDay(v int64) *ConfigBackupTaskRequest {
	s.BackupHourInDay = &v
	return s
}

func (s *ConfigBackupTaskRequest) SetBackupId(v string) *ConfigBackupTaskRequest {
	s.BackupId = &v
	return s
}

func (s *ConfigBackupTaskRequest) SetBackupPeriod(v int64) *ConfigBackupTaskRequest {
	s.BackupPeriod = &v
	return s
}

func (s *ConfigBackupTaskRequest) SetManual2PeriodicList(v []*string) *ConfigBackupTaskRequest {
	s.Manual2PeriodicList = v
	return s
}

func (s *ConfigBackupTaskRequest) SetPeriodic2ManualList(v []*string) *ConfigBackupTaskRequest {
	s.Periodic2ManualList = v
	return s
}

func (s *ConfigBackupTaskRequest) Validate() error {
	return dara.Validate(s)
}
