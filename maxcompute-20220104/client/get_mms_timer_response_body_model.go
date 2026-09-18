// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsTimerResponseBodyData) *GetMmsTimerResponseBody
	GetData() *GetMmsTimerResponseBodyData
	SetRequestId(v string) *GetMmsTimerResponseBody
	GetRequestId() *string
}

type GetMmsTimerResponseBody struct {
	// The data returned.
	Data *GetMmsTimerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0a06dfe716674588654372173ec0da
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponseBody) GetData() *GetMmsTimerResponseBodyData {
	return s.Data
}

func (s *GetMmsTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsTimerResponseBody) SetData(v *GetMmsTimerResponseBodyData) *GetMmsTimerResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsTimerResponseBody) SetRequestId(v string) *GetMmsTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsTimerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTimerResponseBodyData struct {
	// The configuration of the migration job.
	Config *GetMmsTimerResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The creation time of the scheduled task. This is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1730946421757
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the source database.
	//
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// The ID of the scheduled task.
	//
	// example:
	//
	// 2523
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the scheduled task.
	//
	// example:
	//
	// sale_detail
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The scheduling type of the scheduled task. Valid values: `Daily` and `Hourly`.
	//
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// 2000017
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The name of the source database.
	//
	// example:
	//
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// Indicates whether the scheduled task is stopped.
	//
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// The type of the scheduled task.
	//
	// example:
	//
	// Daily, Hourly
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The last update time of the scheduled task, in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-01T02:18:01Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The scheduling time. If `scheduleType` is `Daily`, the value is in the `HH:MM` format. If `scheduleType` is `Hourly`, the value is in the `MM` format.
	//
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMmsTimerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponseBodyData) GetConfig() *GetMmsTimerResponseBodyDataConfig {
	return s.Config
}

func (s *GetMmsTimerResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMmsTimerResponseBodyData) GetDbId() *int64 {
	return s.DbId
}

func (s *GetMmsTimerResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsTimerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMmsTimerResponseBodyData) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetMmsTimerResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsTimerResponseBodyData) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *GetMmsTimerResponseBodyData) GetStopped() *bool {
	return s.Stopped
}

func (s *GetMmsTimerResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsTimerResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMmsTimerResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *GetMmsTimerResponseBodyData) SetConfig(v *GetMmsTimerResponseBodyDataConfig) *GetMmsTimerResponseBodyData {
	s.Config = v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetCreateTime(v string) *GetMmsTimerResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetDbId(v int64) *GetMmsTimerResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetId(v int64) *GetMmsTimerResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetName(v string) *GetMmsTimerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetScheduleType(v string) *GetMmsTimerResponseBodyData {
	s.ScheduleType = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetSourceId(v int64) *GetMmsTimerResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetSrcDbName(v string) *GetMmsTimerResponseBodyData {
	s.SrcDbName = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetStopped(v bool) *GetMmsTimerResponseBodyData {
	s.Stopped = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetType(v string) *GetMmsTimerResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetUpdateTime(v string) *GetMmsTimerResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetValue(v string) *GetMmsTimerResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTimerResponseBodyDataConfig struct {
	// A map of source column names to destination column names.
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// Whether to migrate table data.
	EnableDataMigration *bool `json:"enableDataMigration,omitempty" xml:"enableDataMigration,omitempty"`
	// Whether to migrate the table schema.
	EnableSchemaMigration *bool `json:"enableSchemaMigration,omitempty" xml:"enableSchemaMigration,omitempty"`
	// Whether to enable verification. The system performs verification by running a `SELECT COUNT(*)` query on both the source and destination to compare the row count.
	//
	// example:
	//
	// true
	EnableVerification *bool `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	// Other configurations.
	Others map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	// A map of table names to their corresponding partition filter expressions.
	PartitionFilters map[string]*string `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	// If `type` is set to `Partitions`, this parameter specifies a list of partition IDs to migrate.
	Partitions []*int64 `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// If `type` is set to `Database`, this parameter specifies a table deny list. Tables on this list are excluded from the migration.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// A map of source table names to destination table names.
	TableMapping map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	// If `type` is set to `Database`, this parameter specifies a table allowlist. If this parameter is not specified, all tables in the database are migrated.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// If `type` is set to `Tables`, this parameter specifies a list of table names to migrate.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s GetMmsTimerResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponseBodyDataConfig) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *GetMmsTimerResponseBodyDataConfig) GetEnableDataMigration() *bool {
	return s.EnableDataMigration
}

func (s *GetMmsTimerResponseBodyDataConfig) GetEnableSchemaMigration() *bool {
	return s.EnableSchemaMigration
}

func (s *GetMmsTimerResponseBodyDataConfig) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *GetMmsTimerResponseBodyDataConfig) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *GetMmsTimerResponseBodyDataConfig) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *GetMmsTimerResponseBodyDataConfig) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTables() []*string {
	return s.Tables
}

func (s *GetMmsTimerResponseBodyDataConfig) SetColumnMapping(v map[string]*string) *GetMmsTimerResponseBodyDataConfig {
	s.ColumnMapping = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetEnableDataMigration(v bool) *GetMmsTimerResponseBodyDataConfig {
	s.EnableDataMigration = &v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetEnableSchemaMigration(v bool) *GetMmsTimerResponseBodyDataConfig {
	s.EnableSchemaMigration = &v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetEnableVerification(v bool) *GetMmsTimerResponseBodyDataConfig {
	s.EnableVerification = &v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetOthers(v map[string]interface{}) *GetMmsTimerResponseBodyDataConfig {
	s.Others = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetPartitionFilters(v map[string]*string) *GetMmsTimerResponseBodyDataConfig {
	s.PartitionFilters = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetPartitions(v []*int64) *GetMmsTimerResponseBodyDataConfig {
	s.Partitions = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTableBlackList(v []*string) *GetMmsTimerResponseBodyDataConfig {
	s.TableBlackList = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTableMapping(v map[string]*string) *GetMmsTimerResponseBodyDataConfig {
	s.TableMapping = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTableWhiteList(v []*string) *GetMmsTimerResponseBodyDataConfig {
	s.TableWhiteList = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTables(v []*string) *GetMmsTimerResponseBodyDataConfig {
	s.Tables = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) Validate() error {
	return dara.Validate(s)
}
