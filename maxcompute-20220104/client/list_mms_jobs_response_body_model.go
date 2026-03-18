// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsJobsResponseBodyData) *ListMmsJobsResponseBody
	GetData() *ListMmsJobsResponseBodyData
	SetRequestId(v string) *ListMmsJobsResponseBody
	GetRequestId() *string
}

type ListMmsJobsResponseBody struct {
	// The details of the returned data.
	Data *ListMmsJobsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1112E7C7-C65F-57A2-A7C7-3B178AA257B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBody) GetData() *ListMmsJobsResponseBodyData {
	return s.Data
}

func (s *ListMmsJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsJobsResponseBody) SetData(v *ListMmsJobsResponseBodyData) *ListMmsJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsJobsResponseBody) SetRequestId(v string) *ListMmsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsJobsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsJobsResponseBodyData struct {
	// The list of migration jobs.
	ObjectList []*ListMmsJobsResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBodyData) GetObjectList() []*ListMmsJobsResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsJobsResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsJobsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsJobsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsJobsResponseBodyData) SetObjectList(v []*ListMmsJobsResponseBodyDataObjectList) *ListMmsJobsResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsJobsResponseBodyData) SetPageNum(v int32) *ListMmsJobsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsJobsResponseBodyData) SetPageSize(v int32) *ListMmsJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsJobsResponseBodyData) SetTotal(v int32) *ListMmsJobsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsJobsResponseBodyData) Validate() error {
	if s.ObjectList != nil {
		for _, item := range s.ObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsJobsResponseBodyDataObjectList struct {
	// The configuration of the migration job.
	Config *ListMmsJobsResponseBodyDataObjectListConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Indicates whether the migration job is stopped.
	//
	// example:
	//
	// 2024-12-17 15:44:17
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The source database ID.
	//
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
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
	// test_table_1
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// The estimated completion time of the migration. A smaller eta value increases the priority of the migration job.
	//
	// example:
	//
	// 2025-05-06
	Eta *string `json:"eta,omitempty" xml:"eta,omitempty"`
	// The migration job ID.
	//
	// example:
	//
	// 18
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
	// 2000015
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
	// test_db_1
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// The source schema name. This is the schema in a Layer 3 namespace.
	//
	// example:
	//
	// test_table_1
	SrcSchemaName *string `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	// The status of the migration job.
	//
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the migration job is stopped.
	//
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// The number of completed migration tasks.
	//
	// example:
	//
	// 10
	TaskDone *int32 `json:"taskDone,omitempty" xml:"taskDone,omitempty"`
	// The number of migration tasks in the job.
	//
	// example:
	//
	// 10
	TaskNum *int32 `json:"taskNum,omitempty" xml:"taskNum,omitempty"`
	// The migration scope. Valid values: Database, Tables, and Partitions.
	//
	// example:
	//
	// Tables
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsJobsResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetConfig() *ListMmsJobsResponseBodyDataObjectListConfig {
	return s.Config
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetDstDbName() *string {
	return s.DstDbName
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetEta() *string {
	return s.Eta
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetName() *string {
	return s.Name
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetSourceName() *string {
	return s.SourceName
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetSrcSchemaName() *string {
	return s.SrcSchemaName
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetStopped() *bool {
	return s.Stopped
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetTaskDone() *int32 {
	return s.TaskDone
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetTaskNum() *int32 {
	return s.TaskNum
}

func (s *ListMmsJobsResponseBodyDataObjectList) GetType() *string {
	return s.Type
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetConfig(v *ListMmsJobsResponseBodyDataObjectListConfig) *ListMmsJobsResponseBodyDataObjectList {
	s.Config = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetDbId(v int64) *ListMmsJobsResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetDstDbName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.DstDbName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetDstSchemaName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.DstSchemaName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetEta(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Eta = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetId(v int64) *ListMmsJobsResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsJobsResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSourceName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSrcDbName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.SrcDbName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetSrcSchemaName(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.SrcSchemaName = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetStatus(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetStopped(v bool) *ListMmsJobsResponseBodyDataObjectList {
	s.Stopped = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetTaskDone(v int32) *ListMmsJobsResponseBodyDataObjectList {
	s.TaskDone = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetTaskNum(v int32) *ListMmsJobsResponseBodyDataObjectList {
	s.TaskNum = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) SetType(v string) *ListMmsJobsResponseBodyDataObjectList {
	s.Type = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectList) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsJobsResponseBodyDataObjectListConfig struct {
	// The mapping from source column names to destination column names.
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// Enable verification. The current method runs SELECT COUNT on both the source and destination and compares the row counts.
	//
	// example:
	//
	// true
	EnableVerification *bool `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	// Enable incremental migration. Only new or modified partitions are migrated. Modified partitions are re-migrated.
	//
	// example:
	//
	// true
	Increment *bool `json:"increment,omitempty" xml:"increment,omitempty"`
	// Other configuration settings.
	//
	// example:
	//
	// {"spark.executor.mem": "2g"}
	Others map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	// The partition filter expression. Specify the partition filter expression for each table.
	PartitionFilters map[string]*string `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	// When type is set to Partitions, specify the partition IDs of the tables to migrate.
	Partitions []*int64 `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// Deprecated.
	//
	// example:
	//
	// false
	SchemaOnly *bool `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	// When type is set to Database, specify the tables to exclude from migration.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// The mapping from source table names to destination table names.
	TableMapping map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	// When type is set to Database, specify the tables to migrate. If you do not specify tableWhiteList, all tables in the specified database are migrated.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// When type is set to Tables, specify the names of the tables to migrate.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// Deprecated. Valid values: MOCK, HIVE (hive udtf task), HIVE_DATAX (hive datax task), COPY_TASK (ODPS Copy Task), ODPS_INSERT_OVERWRITE (ODPS simple insert overwrite task), MC2MC_VERIFY, OSS, HIVE_OSS, HIVE_SPARK, and BIGQUERY.
	//
	// example:
	//
	// BIGQUERY
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	// Deprecated.
	//
	// example:
	//
	// Depcreated
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
}

func (s ListMmsJobsResponseBodyDataObjectListConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMmsJobsResponseBodyDataObjectListConfig) GoString() string {
	return s.String()
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetIncrement() *bool {
	return s.Increment
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetSchemaOnly() *bool {
	return s.SchemaOnly
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetTables() []*string {
	return s.Tables
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetTaskType() *string {
	return s.TaskType
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) GetTunnelQuota() *string {
	return s.TunnelQuota
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetColumnMapping(v map[string]*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.ColumnMapping = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetEnableVerification(v bool) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.EnableVerification = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetIncrement(v bool) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Increment = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetOthers(v map[string]interface{}) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Others = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetPartitionFilters(v map[string]*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.PartitionFilters = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetPartitions(v []*int64) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Partitions = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetSchemaOnly(v bool) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.SchemaOnly = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTableBlackList(v []*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TableBlackList = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTableMapping(v map[string]*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TableMapping = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTableWhiteList(v []*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TableWhiteList = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTables(v []*string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.Tables = v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTaskType(v string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TaskType = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) SetTunnelQuota(v string) *ListMmsJobsResponseBodyDataObjectListConfig {
	s.TunnelQuota = &v
	return s
}

func (s *ListMmsJobsResponseBodyDataObjectListConfig) Validate() error {
	return dara.Validate(s)
}
