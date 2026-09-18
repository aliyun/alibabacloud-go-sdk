// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnMapping(v map[string]*string) *CreateMmsTimerRequest
	GetColumnMapping() map[string]*string
	SetEnableDataMigration(v bool) *CreateMmsTimerRequest
	GetEnableDataMigration() *bool
	SetEnableSchemaMigration(v bool) *CreateMmsTimerRequest
	GetEnableSchemaMigration() *bool
	SetEnableVerification(v bool) *CreateMmsTimerRequest
	GetEnableVerification() *bool
	SetName(v string) *CreateMmsTimerRequest
	GetName() *string
	SetOthers(v map[string]interface{}) *CreateMmsTimerRequest
	GetOthers() map[string]interface{}
	SetPartitionFilters(v map[string]*string) *CreateMmsTimerRequest
	GetPartitionFilters() map[string]*string
	SetPartitions(v []*int64) *CreateMmsTimerRequest
	GetPartitions() []*int64
	SetScheduleType(v string) *CreateMmsTimerRequest
	GetScheduleType() *string
	SetSourceId(v int64) *CreateMmsTimerRequest
	GetSourceId() *int64
	SetSrcDbName(v string) *CreateMmsTimerRequest
	GetSrcDbName() *string
	SetTableBlackList(v []*string) *CreateMmsTimerRequest
	GetTableBlackList() []*string
	SetTableMapping(v map[string]*string) *CreateMmsTimerRequest
	GetTableMapping() map[string]*string
	SetTableWhiteList(v []*string) *CreateMmsTimerRequest
	GetTableWhiteList() []*string
	SetTables(v []*string) *CreateMmsTimerRequest
	GetTables() []*string
	SetValue(v string) *CreateMmsTimerRequest
	GetValue() *string
}

type CreateMmsTimerRequest struct {
	// A map of source column names to target column names.
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// Specifies whether to migrate table data.
	//
	// example:
	//
	// true
	EnableDataMigration *bool `json:"enableDataMigration,omitempty" xml:"enableDataMigration,omitempty"`
	// Specifies whether to migrate the table schema.
	//
	// example:
	//
	// true
	EnableSchemaMigration *bool `json:"enableSchemaMigration,omitempty" xml:"enableSchemaMigration,omitempty"`
	// Specifies whether to enable data verification. If set to `true`, the system runs a `SELECT COUNT(*)` query on both the source and target tables and compares the row counts.
	//
	// example:
	//
	// false
	EnableVerification *bool `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	// The name of the scheduled task.
	//
	// example:
	//
	// planA
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Other configuration settings.
	Others map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	// A map of table names to their corresponding partition filter expressions.
	PartitionFilters map[string]*string `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	// A list of IDs for the table partitions to migrate. This parameter takes effect only when the `type` parameter is set to `Partitions`.
	Partitions []*int64 `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// The schedule type for the task.
	//
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// 2000014
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The name of the source database.
	//
	// example:
	//
	// src_db
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// A blacklist of tables to exclude from the migration. This parameter takes effect only when the `type` parameter is set to `Database`.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// A map of source table names to target table names.
	TableMapping map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	// A whitelist of tables to migrate. This parameter takes effect only when the `type` parameter is set to `Database`. If omitted, all tables in the source database are migrated.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// A list of table names to migrate. This parameter takes effect only when the `type` parameter is set to `Tables`.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// The time to run the scheduled task. If `scheduleType` is set to `Daily`, the value is the time in `HH:MM` format. If `scheduleType` is set to `Hourly`, the value is the minute of the hour (`MM`).
	//
	// example:
	//
	// 12:00
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMmsTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsTimerRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsTimerRequest) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *CreateMmsTimerRequest) GetEnableDataMigration() *bool {
	return s.EnableDataMigration
}

func (s *CreateMmsTimerRequest) GetEnableSchemaMigration() *bool {
	return s.EnableSchemaMigration
}

func (s *CreateMmsTimerRequest) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *CreateMmsTimerRequest) GetName() *string {
	return s.Name
}

func (s *CreateMmsTimerRequest) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *CreateMmsTimerRequest) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *CreateMmsTimerRequest) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *CreateMmsTimerRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateMmsTimerRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *CreateMmsTimerRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *CreateMmsTimerRequest) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *CreateMmsTimerRequest) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *CreateMmsTimerRequest) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *CreateMmsTimerRequest) GetTables() []*string {
	return s.Tables
}

func (s *CreateMmsTimerRequest) GetValue() *string {
	return s.Value
}

func (s *CreateMmsTimerRequest) SetColumnMapping(v map[string]*string) *CreateMmsTimerRequest {
	s.ColumnMapping = v
	return s
}

func (s *CreateMmsTimerRequest) SetEnableDataMigration(v bool) *CreateMmsTimerRequest {
	s.EnableDataMigration = &v
	return s
}

func (s *CreateMmsTimerRequest) SetEnableSchemaMigration(v bool) *CreateMmsTimerRequest {
	s.EnableSchemaMigration = &v
	return s
}

func (s *CreateMmsTimerRequest) SetEnableVerification(v bool) *CreateMmsTimerRequest {
	s.EnableVerification = &v
	return s
}

func (s *CreateMmsTimerRequest) SetName(v string) *CreateMmsTimerRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsTimerRequest) SetOthers(v map[string]interface{}) *CreateMmsTimerRequest {
	s.Others = v
	return s
}

func (s *CreateMmsTimerRequest) SetPartitionFilters(v map[string]*string) *CreateMmsTimerRequest {
	s.PartitionFilters = v
	return s
}

func (s *CreateMmsTimerRequest) SetPartitions(v []*int64) *CreateMmsTimerRequest {
	s.Partitions = v
	return s
}

func (s *CreateMmsTimerRequest) SetScheduleType(v string) *CreateMmsTimerRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateMmsTimerRequest) SetSourceId(v int64) *CreateMmsTimerRequest {
	s.SourceId = &v
	return s
}

func (s *CreateMmsTimerRequest) SetSrcDbName(v string) *CreateMmsTimerRequest {
	s.SrcDbName = &v
	return s
}

func (s *CreateMmsTimerRequest) SetTableBlackList(v []*string) *CreateMmsTimerRequest {
	s.TableBlackList = v
	return s
}

func (s *CreateMmsTimerRequest) SetTableMapping(v map[string]*string) *CreateMmsTimerRequest {
	s.TableMapping = v
	return s
}

func (s *CreateMmsTimerRequest) SetTableWhiteList(v []*string) *CreateMmsTimerRequest {
	s.TableWhiteList = v
	return s
}

func (s *CreateMmsTimerRequest) SetTables(v []*string) *CreateMmsTimerRequest {
	s.Tables = v
	return s
}

func (s *CreateMmsTimerRequest) SetValue(v string) *CreateMmsTimerRequest {
	s.Value = &v
	return s
}

func (s *CreateMmsTimerRequest) Validate() error {
	return dara.Validate(s)
}
