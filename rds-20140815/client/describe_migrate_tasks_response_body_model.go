// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMigrateTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeMigrateTasksResponseBody
	GetDBInstanceId() *string
	SetItems(v *DescribeMigrateTasksResponseBodyItems) *DescribeMigrateTasksResponseBody
	GetItems() *DescribeMigrateTasksResponseBodyItems
	SetPageNumber(v int32) *DescribeMigrateTasksResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeMigrateTasksResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeMigrateTasksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeMigrateTasksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeMigrateTasksResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The details of the migration task.
	Items *DescribeMigrateTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E356DDF-6B83-45DB-99D5-4B1E8A0D286B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMigrateTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTasksResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeMigrateTasksResponseBody) GetItems() *DescribeMigrateTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeMigrateTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMigrateTasksResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeMigrateTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMigrateTasksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeMigrateTasksResponseBody) SetDBInstanceId(v string) *DescribeMigrateTasksResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeMigrateTasksResponseBody) SetItems(v *DescribeMigrateTasksResponseBodyItems) *DescribeMigrateTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMigrateTasksResponseBody) SetPageNumber(v int32) *DescribeMigrateTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrateTasksResponseBody) SetPageRecordCount(v int32) *DescribeMigrateTasksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrateTasksResponseBody) SetRequestId(v string) *DescribeMigrateTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrateTasksResponseBody) SetTotalRecordCount(v int32) *DescribeMigrateTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrateTasksResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMigrateTasksResponseBodyItems struct {
	MigrateTask []*DescribeMigrateTasksResponseBodyItemsMigrateTask `json:"MigrateTask,omitempty" xml:"MigrateTask,omitempty" type:"Repeated"`
}

func (s DescribeMigrateTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTasksResponseBodyItems) GetMigrateTask() []*DescribeMigrateTasksResponseBodyItemsMigrateTask {
	return s.MigrateTask
}

func (s *DescribeMigrateTasksResponseBodyItems) SetMigrateTask(v []*DescribeMigrateTasksResponseBodyItemsMigrateTask) *DescribeMigrateTasksResponseBodyItems {
	s.MigrateTask = v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItems) Validate() error {
	if s.MigrateTask != nil {
		for _, item := range s.MigrateTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMigrateTasksResponseBodyItemsMigrateTask struct {
	// The migration task type. Valid values:
	//
	// 	- **FULL**: The migration task migrates full backup files that can be used to restore the full data of the instance.
	//
	// 	- **UPDF**: The migration task migrates incremental or log backup files that can be used to restore the incremental data of the instance.
	//
	// example:
	//
	// FULL
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The time when the migration task was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-05-30T12:11:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The database name.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The description of the migration task.
	//
	// example:
	//
	// Api description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the migration task was completed. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-05-30T13:11:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the imported data overwrites the existing data.
	//
	// example:
	//
	// True
	IsDBReplaced *string `json:"IsDBReplaced,omitempty" xml:"IsDBReplaced,omitempty"`
	// The migration task ID.
	//
	// example:
	//
	// 564522545
	MigrateTaskId *string `json:"MigrateTaskId,omitempty" xml:"MigrateTaskId,omitempty"`
	// The status of the migration task. Valid values:
	//
	// 	- **NoStart**: The task is not started.
	//
	// 	- **Running**:The task is in progress.
	//
	// 	- **Success**: The task is successful.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Waiting**: The task is waiting for an incremental backup file to be imported.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrateTasksResponseBodyItemsMigrateTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeMigrateTasksResponseBodyItemsMigrateTask) GoString() string {
	return s.String()
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetDBName() *string {
	return s.DBName
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetDescription() *string {
	return s.Description
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetIsDBReplaced() *string {
	return s.IsDBReplaced
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetMigrateTaskId() *string {
	return s.MigrateTaskId
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) GetStatus() *string {
	return s.Status
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetBackupMode(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.BackupMode = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetCreateTime(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.CreateTime = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetDBName(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.DBName = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetDescription(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.Description = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetEndTime(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.EndTime = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetIsDBReplaced(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.IsDBReplaced = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetMigrateTaskId(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.MigrateTaskId = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) SetStatus(v string) *DescribeMigrateTasksResponseBodyItemsMigrateTask {
	s.Status = &v
	return s
}

func (s *DescribeMigrateTasksResponseBodyItemsMigrateTask) Validate() error {
	return dara.Validate(s)
}
