// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnMapping(v map[string]*string) *CreateMmsJobRequest
	GetColumnMapping() map[string]*string
	SetDstDbName(v string) *CreateMmsJobRequest
	GetDstDbName() *string
	SetDstSchemaName(v string) *CreateMmsJobRequest
	GetDstSchemaName() *string
	SetEnableDataMigration(v bool) *CreateMmsJobRequest
	GetEnableDataMigration() *bool
	SetEnableSchemaMigration(v bool) *CreateMmsJobRequest
	GetEnableSchemaMigration() *bool
	SetEnableVerification(v bool) *CreateMmsJobRequest
	GetEnableVerification() *bool
	SetEta(v string) *CreateMmsJobRequest
	GetEta() *string
	SetIncrement(v bool) *CreateMmsJobRequest
	GetIncrement() *bool
	SetName(v string) *CreateMmsJobRequest
	GetName() *string
	SetOthers(v map[string]interface{}) *CreateMmsJobRequest
	GetOthers() map[string]interface{}
	SetPartitionFilters(v map[string]*string) *CreateMmsJobRequest
	GetPartitionFilters() map[string]*string
	SetPartitions(v []*int64) *CreateMmsJobRequest
	GetPartitions() []*int64
	SetSchemaOnly(v bool) *CreateMmsJobRequest
	GetSchemaOnly() *bool
	SetSourceId(v int64) *CreateMmsJobRequest
	GetSourceId() *int64
	SetSourceName(v string) *CreateMmsJobRequest
	GetSourceName() *string
	SetSrcDbName(v string) *CreateMmsJobRequest
	GetSrcDbName() *string
	SetSrcSchemaName(v string) *CreateMmsJobRequest
	GetSrcSchemaName() *string
	SetTableBlackList(v []*string) *CreateMmsJobRequest
	GetTableBlackList() []*string
	SetTableMapping(v map[string]*string) *CreateMmsJobRequest
	GetTableMapping() map[string]*string
	SetTableWhiteList(v []*string) *CreateMmsJobRequest
	GetTableWhiteList() []*string
	SetTables(v []*string) *CreateMmsJobRequest
	GetTables() []*string
	SetTaskType(v string) *CreateMmsJobRequest
	GetTaskType() *string
}

type CreateMmsJobRequest struct {
	// {Source column name: Destination column name}
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// The destination MaxCompute project.
	//
	// example:
	//
	// mms_test
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// The destination MaxCompute schema.
	//
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// Specifies whether to migrate table data.
	//
	// example:
	//
	// true
	EnableDataMigration *bool `json:"enableDataMigration,omitempty" xml:"enableDataMigration,omitempty"`
	// Specifies whether to migrate table schemas.
	//
	// example:
	//
	// true
	EnableSchemaMigration *bool `json:"enableSchemaMigration,omitempty" xml:"enableSchemaMigration,omitempty"`
	// Specifies whether to enable data verification. The current verification method is to execute SELECT COUNT(\\*) on the source and destination to compare the number of rows.
	//
	// example:
	//
	// false
	EnableVerification *bool `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	// The expected completion time of the migration. Note: A smaller eta value gives the migration task higher priority.
	//
	// example:
	//
	// 2025-02-04
	Eta *string `json:"eta,omitempty" xml:"eta,omitempty"`
	// Specifies whether to perform an incremental migration. In an incremental migration, only new or changed partitions are migrated. Note that changed partitions are re-migrated.
	//
	// example:
	//
	// true
	Increment *bool `json:"increment,omitempty" xml:"increment,omitempty"`
	// The name of the migration job.
	//
	// example:
	//
	// migrate_db_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Other configuration information.
	Others map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	// {Table name: Partition filter expression}
	PartitionFilters map[string]*string `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	// The list of partition IDs.
	Partitions []*int64 `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// Specifies whether to migrate only metadata.
	//
	// example:
	//
	// false
	SchemaOnly *bool `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	// The ID of the data source.
	//
	// example:
	//
	// 2000014
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The name of the source database.
	//
	// example:
	//
	// src_db
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// The name of the source schema. This is the schema in a Layer 3 namespace.
	//
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// The blacklist of tables.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// {Source table: Destination table}
	TableMapping map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	// The whitelist of tables. Note: If you configure both a whitelist and a blacklist, only the blacklist takes effect.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// The list of table names.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// The type of the migration task.
	//
	// example:
	//
	// BIGQUERY
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateMmsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsJobRequest) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *CreateMmsJobRequest) GetDstDbName() *string {
	return s.DstDbName
}

func (s *CreateMmsJobRequest) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *CreateMmsJobRequest) GetEnableDataMigration() *bool {
	return s.EnableDataMigration
}

func (s *CreateMmsJobRequest) GetEnableSchemaMigration() *bool {
	return s.EnableSchemaMigration
}

func (s *CreateMmsJobRequest) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *CreateMmsJobRequest) GetEta() *string {
	return s.Eta
}

func (s *CreateMmsJobRequest) GetIncrement() *bool {
	return s.Increment
}

func (s *CreateMmsJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateMmsJobRequest) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *CreateMmsJobRequest) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *CreateMmsJobRequest) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *CreateMmsJobRequest) GetSchemaOnly() *bool {
	return s.SchemaOnly
}

func (s *CreateMmsJobRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *CreateMmsJobRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *CreateMmsJobRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *CreateMmsJobRequest) GetSrcSchemaName() *string {
	return s.SrcSchemaName
}

func (s *CreateMmsJobRequest) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *CreateMmsJobRequest) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *CreateMmsJobRequest) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *CreateMmsJobRequest) GetTables() []*string {
	return s.Tables
}

func (s *CreateMmsJobRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateMmsJobRequest) SetColumnMapping(v map[string]*string) *CreateMmsJobRequest {
	s.ColumnMapping = v
	return s
}

func (s *CreateMmsJobRequest) SetDstDbName(v string) *CreateMmsJobRequest {
	s.DstDbName = &v
	return s
}

func (s *CreateMmsJobRequest) SetDstSchemaName(v string) *CreateMmsJobRequest {
	s.DstSchemaName = &v
	return s
}

func (s *CreateMmsJobRequest) SetEnableDataMigration(v bool) *CreateMmsJobRequest {
	s.EnableDataMigration = &v
	return s
}

func (s *CreateMmsJobRequest) SetEnableSchemaMigration(v bool) *CreateMmsJobRequest {
	s.EnableSchemaMigration = &v
	return s
}

func (s *CreateMmsJobRequest) SetEnableVerification(v bool) *CreateMmsJobRequest {
	s.EnableVerification = &v
	return s
}

func (s *CreateMmsJobRequest) SetEta(v string) *CreateMmsJobRequest {
	s.Eta = &v
	return s
}

func (s *CreateMmsJobRequest) SetIncrement(v bool) *CreateMmsJobRequest {
	s.Increment = &v
	return s
}

func (s *CreateMmsJobRequest) SetName(v string) *CreateMmsJobRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsJobRequest) SetOthers(v map[string]interface{}) *CreateMmsJobRequest {
	s.Others = v
	return s
}

func (s *CreateMmsJobRequest) SetPartitionFilters(v map[string]*string) *CreateMmsJobRequest {
	s.PartitionFilters = v
	return s
}

func (s *CreateMmsJobRequest) SetPartitions(v []*int64) *CreateMmsJobRequest {
	s.Partitions = v
	return s
}

func (s *CreateMmsJobRequest) SetSchemaOnly(v bool) *CreateMmsJobRequest {
	s.SchemaOnly = &v
	return s
}

func (s *CreateMmsJobRequest) SetSourceId(v int64) *CreateMmsJobRequest {
	s.SourceId = &v
	return s
}

func (s *CreateMmsJobRequest) SetSourceName(v string) *CreateMmsJobRequest {
	s.SourceName = &v
	return s
}

func (s *CreateMmsJobRequest) SetSrcDbName(v string) *CreateMmsJobRequest {
	s.SrcDbName = &v
	return s
}

func (s *CreateMmsJobRequest) SetSrcSchemaName(v string) *CreateMmsJobRequest {
	s.SrcSchemaName = &v
	return s
}

func (s *CreateMmsJobRequest) SetTableBlackList(v []*string) *CreateMmsJobRequest {
	s.TableBlackList = v
	return s
}

func (s *CreateMmsJobRequest) SetTableMapping(v map[string]*string) *CreateMmsJobRequest {
	s.TableMapping = v
	return s
}

func (s *CreateMmsJobRequest) SetTableWhiteList(v []*string) *CreateMmsJobRequest {
	s.TableWhiteList = v
	return s
}

func (s *CreateMmsJobRequest) SetTables(v []*string) *CreateMmsJobRequest {
	s.Tables = v
	return s
}

func (s *CreateMmsJobRequest) SetTaskType(v string) *CreateMmsJobRequest {
	s.TaskType = &v
	return s
}

func (s *CreateMmsJobRequest) Validate() error {
	return dara.Validate(s)
}
