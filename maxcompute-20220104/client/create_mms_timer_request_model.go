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
	// The mapping between source and destination column names in the {source column name: destination column name} format.
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// Specifies whether to migrate the table data.
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
	// Specifies whether to enable verification. Currently, the verification method executes SELECT COUNT on the source and destination to compare the number of rows.
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
	// Other configuration information.
	Others map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	// The partition filter expression. Configure the partition filter expression for the specified table.
	PartitionFilters map[string]*string `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	// The list of partition IDs for the tables to be migrated. Specify this parameter when type is set to Partitions.
	Partitions []*int64 `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// The schedule type of the scheduled task. Valid values: Daily (runs daily) and Hourly (runs hourly).
	//
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The data source ID.
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
	// The list of tables to exclude from migration. Specify this parameter when type is set to Database.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// The mapping between the source table name and the destination table name.
	TableMapping map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	// The list of tables to migrate. Specify this parameter when type is set to Database. If this parameter is not specified, all tables in the corresponding database are migrated.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// The list of names of the tables to be migrated. Specify this parameter when type is set to Tables.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// The schedule time of the scheduled task. When scheduleType is set to Daily, the value is in the HH:MM format. When scheduleType is set to Hourly, the value is in the MM format.
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
