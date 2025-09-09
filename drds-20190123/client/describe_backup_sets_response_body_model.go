// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupSets(v *DescribeBackupSetsResponseBodyBackupSets) *DescribeBackupSetsResponseBody
	GetBackupSets() *DescribeBackupSetsResponseBodyBackupSets
	SetRequestId(v string) *DescribeBackupSetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupSetsResponseBody
	GetSuccess() *bool
}

type DescribeBackupSetsResponseBody struct {
	// The list of backup sets.
	BackupSets *DescribeBackupSetsResponseBodyBackupSets `json:"BackupSets,omitempty" xml:"BackupSets,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7103AEE3-9025-442F-B82B-BABD0A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackupSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBody) GetBackupSets() *DescribeBackupSetsResponseBodyBackupSets {
	return s.BackupSets
}

func (s *DescribeBackupSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupSetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupSetsResponseBody) SetBackupSets(v *DescribeBackupSetsResponseBodyBackupSets) *DescribeBackupSetsResponseBody {
	s.BackupSets = v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetRequestId(v string) *DescribeBackupSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupSetsResponseBody) SetSuccess(v bool) *DescribeBackupSetsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupSetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupSetsResponseBodyBackupSets struct {
	BackupSet []*DescribeBackupSetsResponseBodyBackupSetsBackupSet `json:"backupSet,omitempty" xml:"backupSet,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetsResponseBodyBackupSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSets) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSets) GetBackupSet() []*DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	return s.BackupSet
}

func (s *DescribeBackupSetsResponseBodyBackupSets) SetBackupSet(v []*DescribeBackupSetsResponseBodyBackupSetsBackupSet) *DescribeBackupSetsResponseBodyBackupSets {
	s.BackupSet = v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSets) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupSetsResponseBodyBackupSetsBackupSet struct {
	// Backup Recovery duration.
	//
	// example:
	//
	// 2020-06-05 11:31:38
	BackupConsitentTime *string `json:"BackupConsitentTime,omitempty" xml:"BackupConsitentTime,omitempty"`
	// The list of backup databases.
	BackupDbs *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs `json:"BackupDbs,omitempty" xml:"BackupDbs,omitempty" type:"Struct"`
	// The end of the backup time which is in timestamp format (measured in millisecond).
	//
	// >  0 indicates not finished.
	//
	// example:
	//
	// 1591327899000
	BackupEndTime *int64 `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	// The level of the backup. Valid values:
	//
	// 	- db: The database level.
	//
	// 	- instance: the instance level.
	//
	// example:
	//
	// instance
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup method. Valid values:
	//
	// 	- logic: the logical backup.
	//
	// 	- phy: fast backup
	//
	// example:
	//
	// logic
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The beginning of the backup time which is in timestamp format (measured in millisecond).
	//
	// example:
	//
	// 1591327754000
	BackupStartTime *int64 `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// The size of the backup set. Unit: MB.
	//
	// example:
	//
	// 93.24
	BackupTotalSize *string `json:"BackupTotalSize,omitempty" xml:"BackupTotalSize,omitempty"`
	// The type of the backup. Valid values:
	//
	// 	- manual: indicates a manual backup.
	//
	// 	- auto: indicates an automatic backup.
	//
	// example:
	//
	// manual
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// Indicates whether the backup set can be restored. Valid values:
	//
	// example:
	//
	// false
	EnableRecovery *bool `json:"EnableRecovery,omitempty" xml:"EnableRecovery,omitempty"`
	// The ID of the data backup file you want to use.
	//
	// example:
	//
	// ba30d5c4-a6dc-11ea-bd40-************
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the backup instance. Valid values:
	//
	// 	- \\-1: Failed
	//
	// 	- 0: Not started
	//
	// 	- 1: The storage instance is running.
	//
	// 	- 2: Success
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSet) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupConsitentTime() *string {
	return s.BackupConsitentTime
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupDbs() *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs {
	return s.BackupDbs
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupEndTime() *int64 {
	return s.BackupEndTime
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupStartTime() *int64 {
	return s.BackupStartTime
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupTotalSize() *string {
	return s.BackupTotalSize
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetEnableRecovery() *bool {
	return s.EnableRecovery
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetId() *string {
	return s.Id
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupConsitentTime(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupConsitentTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupDbs(v *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupDbs = v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupEndTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupLevel(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupLevel = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupMode(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupStartTime(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupTotalSize(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupTotalSize = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetBackupType(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetEnableRecovery(v bool) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.EnableRecovery = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetId(v string) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Id = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) SetStatus(v int64) *DescribeBackupSetsResponseBodyBackupSetsBackupSet {
	s.Status = &v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSet) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs struct {
	BackupDb []*string `json:"backupDb,omitempty" xml:"backupDb,omitempty" type:"Repeated"`
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) GetBackupDb() []*string {
	return s.BackupDb
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) SetBackupDb(v []*string) *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs {
	s.BackupDb = v
	return s
}

func (s *DescribeBackupSetsResponseBodyBackupSetsBackupSetBackupDbs) Validate() error {
	return dara.Validate(s)
}
