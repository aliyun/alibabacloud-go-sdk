// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsJobResponseBodyData) *GetMmsJobResponseBody
	GetData() *GetMmsJobResponseBodyData
	SetRequestId(v string) *GetMmsJobResponseBody
	GetRequestId() *string
}

type GetMmsJobResponseBody struct {
	// The migration job object.
	Data *GetMmsJobResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D9F872FD-5DDE-30A6-8C8A-1B8C6A81059F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponseBody) GetData() *GetMmsJobResponseBodyData {
	return s.Data
}

func (s *GetMmsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsJobResponseBody) SetData(v *GetMmsJobResponseBodyData) *GetMmsJobResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsJobResponseBody) SetRequestId(v string) *GetMmsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsJobResponseBodyData struct {
	// The configuration of the migration job.
	Config *GetMmsJobResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The time when the job was created.
	//
	// example:
	//
	// 2024-12-17 15:44:17
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The source database ID.
	//
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// The destination MaxCompute project.
	//
	// example:
	//
	// mms_target
	DstDbName *string `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	// The destination MaxCompute schema.
	//
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// The expected completion time of the migration. Note: A smaller eta value indicates a higher priority for the migration task.
	//
	// example:
	//
	// 2025-05-06
	Eta *string `json:"eta,omitempty" xml:"eta,omitempty"`
	// The migration job ID.
	//
	// example:
	//
	// 10
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the migration job.
	//
	// example:
	//
	// migrate_db_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 2
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
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// The name of the source schema. This parameter specifies the schema in a Layer 3 namespace.
	//
	// example:
	//
	// default
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// The status of the migration task.
	//
	// example:
	//
	// DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Stopped.
	//
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// The number of completed migration tasks.
	//
	// example:
	//
	// 100
	TaskDone *int32 `json:"taskDone,omitempty" xml:"taskDone,omitempty"`
	// The number of migration tasks included in the job.
	//
	// example:
	//
	// 100
	TaskNum *int32 `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
	// The migration scope. Valid values: Database, Tables, and Partitions.
	//
	// example:
	//
	// Tables
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponseBodyData) GetConfig() *GetMmsJobResponseBodyDataConfig {
	return s.Config
}

func (s *GetMmsJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMmsJobResponseBodyData) GetDbId() *int64 {
	return s.DbId
}

func (s *GetMmsJobResponseBodyData) GetDstDbName() *string {
	return s.DstDbName
}

func (s *GetMmsJobResponseBodyData) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *GetMmsJobResponseBodyData) GetEta() *string {
	return s.Eta
}

func (s *GetMmsJobResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsJobResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMmsJobResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsJobResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *GetMmsJobResponseBodyData) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *GetMmsJobResponseBodyData) GetSrcSchemaName() *string {
	return s.SrcSchemaName
}

func (s *GetMmsJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsJobResponseBodyData) GetStopped() *bool {
	return s.Stopped
}

func (s *GetMmsJobResponseBodyData) GetTaskDone() *int32 {
	return s.TaskDone
}

func (s *GetMmsJobResponseBodyData) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *GetMmsJobResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsJobResponseBodyData) SetConfig(v *GetMmsJobResponseBodyDataConfig) *GetMmsJobResponseBodyData {
	s.Config = v
	return s
}

func (s *GetMmsJobResponseBodyData) SetCreateTime(v string) *GetMmsJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetDbId(v int64) *GetMmsJobResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetDstDbName(v string) *GetMmsJobResponseBodyData {
	s.DstDbName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetDstSchemaName(v string) *GetMmsJobResponseBodyData {
	s.DstSchemaName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetEta(v string) *GetMmsJobResponseBodyData {
	s.Eta = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetId(v int64) *GetMmsJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetName(v string) *GetMmsJobResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSourceId(v int64) *GetMmsJobResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSourceName(v string) *GetMmsJobResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSrcDbName(v string) *GetMmsJobResponseBodyData {
	s.SrcDbName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetSrcSchemaName(v string) *GetMmsJobResponseBodyData {
	s.SrcSchemaName = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetStatus(v string) *GetMmsJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetStopped(v bool) *GetMmsJobResponseBodyData {
	s.Stopped = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetTaskDone(v int32) *GetMmsJobResponseBodyData {
	s.TaskDone = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetTaskNum(v int32) *GetMmsJobResponseBodyData {
	s.TaskNum = &v
	return s
}

func (s *GetMmsJobResponseBodyData) SetType(v string) *GetMmsJobResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsJobResponseBodyData) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsJobResponseBodyDataConfig struct {
	// {Source column name: Destination column name}
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// Enables data verification. The current verification method is to execute a SELECT COUNT statement on the source and destination to compare the row counts.
	//
	// example:
	//
	// true
	EnableVerification *bool `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	// Incremental migration. Only new or modified partitions are migrated. Note: Modified partitions are re-migrated.
	//
	// example:
	//
	// true
	Increment *bool `json:"increment,omitempty" xml:"increment,omitempty"`
	// Other configuration information.
	//
	// example:
	//
	// {"spark.executor.mem": "2g"}
	Others map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	// The partition filter expression. This parameter specifies the partition filter expression for a specific table.
	PartitionFilters map[string]*string `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	// If type is set to Partitions, this parameter specifies the list of partition IDs of the table to migrate.
	Partitions []*int64 `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// example:
	//
	// false
	SchemaOnly *bool `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	// If type is set to Database, this parameter specifies the tables to exclude from the migration.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// The mapping from source table names to destination table names.
	TableMapping map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	// If type is set to Database, this parameter specifies the list of tables to migrate. If you do not specify this parameter, all tables in the database are migrated.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// If type is set to Tables, this parameter specifies the list of names of the tables to migrate.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// Deprecated. Valid values: MOCK, HIVE (a Hive user-defined table-valued function (UDTF) task), HIVE_DATAX (a Hive DataX task), COPY_TASK (an ODPS Copy Task), ODPS_INSERT_OVERWRITE (an ODPS simple insert overwrite task), MC2MC_VERIFY, OSS, HIVE_OSS, HIVE_SPARK, and BIGQUERY.
	//
	// example:
	//
	// BIGQUERY
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// Deprecated
	//
	// example:
	//
	// Depcreated
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
}

func (s GetMmsJobResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMmsJobResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetMmsJobResponseBodyDataConfig) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *GetMmsJobResponseBodyDataConfig) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *GetMmsJobResponseBodyDataConfig) GetIncrement() *bool {
	return s.Increment
}

func (s *GetMmsJobResponseBodyDataConfig) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *GetMmsJobResponseBodyDataConfig) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *GetMmsJobResponseBodyDataConfig) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *GetMmsJobResponseBodyDataConfig) GetSchemaOnly() *bool {
	return s.SchemaOnly
}

func (s *GetMmsJobResponseBodyDataConfig) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *GetMmsJobResponseBodyDataConfig) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *GetMmsJobResponseBodyDataConfig) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *GetMmsJobResponseBodyDataConfig) GetTables() []*string {
	return s.Tables
}

func (s *GetMmsJobResponseBodyDataConfig) GetTaskType() *string {
	return s.TaskType
}

func (s *GetMmsJobResponseBodyDataConfig) GetTunnelQuota() *string {
	return s.TunnelQuota
}

func (s *GetMmsJobResponseBodyDataConfig) SetColumnMapping(v map[string]*string) *GetMmsJobResponseBodyDataConfig {
	s.ColumnMapping = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetEnableVerification(v bool) *GetMmsJobResponseBodyDataConfig {
	s.EnableVerification = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetIncrement(v bool) *GetMmsJobResponseBodyDataConfig {
	s.Increment = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetOthers(v map[string]interface{}) *GetMmsJobResponseBodyDataConfig {
	s.Others = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetPartitionFilters(v map[string]*string) *GetMmsJobResponseBodyDataConfig {
	s.PartitionFilters = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetPartitions(v []*int64) *GetMmsJobResponseBodyDataConfig {
	s.Partitions = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetSchemaOnly(v bool) *GetMmsJobResponseBodyDataConfig {
	s.SchemaOnly = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTableBlackList(v []*string) *GetMmsJobResponseBodyDataConfig {
	s.TableBlackList = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTableMapping(v map[string]*string) *GetMmsJobResponseBodyDataConfig {
	s.TableMapping = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTableWhiteList(v []*string) *GetMmsJobResponseBodyDataConfig {
	s.TableWhiteList = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTables(v []*string) *GetMmsJobResponseBodyDataConfig {
	s.Tables = v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTaskType(v string) *GetMmsJobResponseBodyDataConfig {
	s.TaskType = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) SetTunnelQuota(v string) *GetMmsJobResponseBodyDataConfig {
	s.TunnelQuota = &v
	return s
}

func (s *GetMmsJobResponseBodyDataConfig) Validate() error {
	return dara.Validate(s)
}
