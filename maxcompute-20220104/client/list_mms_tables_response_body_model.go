// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsTablesResponseBodyData) *ListMmsTablesResponseBody
	GetData() *ListMmsTablesResponseBodyData
	SetRequestId(v string) *ListMmsTablesResponseBody
	GetRequestId() *string
}

type ListMmsTablesResponseBody struct {
	Data *ListMmsTablesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// E7FB14F1-4ACD-5C73-A755-B302D70AB9AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBody) GetData() *ListMmsTablesResponseBodyData {
	return s.Data
}

func (s *ListMmsTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsTablesResponseBody) SetData(v *ListMmsTablesResponseBodyData) *ListMmsTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTablesResponseBody) SetRequestId(v string) *ListMmsTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMmsTablesResponseBodyData struct {
	ObjectList []*ListMmsTablesResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyData) GetObjectList() []*ListMmsTablesResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsTablesResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsTablesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsTablesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsTablesResponseBodyData) SetObjectList(v []*ListMmsTablesResponseBodyDataObjectList) *ListMmsTablesResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsTablesResponseBodyData) SetPageNum(v int32) *ListMmsTablesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsTablesResponseBodyData) SetPageSize(v int32) *ListMmsTablesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsTablesResponseBodyData) SetTotal(v int32) *ListMmsTablesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsTablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListMmsTablesResponseBodyDataObjectList struct {
	// example:
	//
	// 196
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// demo
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// {"mapkey.delim":":","collection.delim":",","serialization.format":"|","field.delim":"|"}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// example:
	//
	// true
	HasPartitions *bool `json:"hasPartitions,omitempty" xml:"hasPartitions,omitempty"`
	// table ID
	//
	// example:
	//
	// 1003476
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// inputFormat
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat
	InputFormat *string `json:"inputFormat,omitempty" xml:"inputFormat,omitempty"`
	// lastDdlTime
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// example:
	//
	// | hdfs://master-1-1.c-c127cd184bb029ea.cn-zhangjiakou.emr.aliyuncs.com:9000/user/hive/warehouse/demo
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 232323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// outFormat
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat
	OutputFormat *string `json:"outputFormat,omitempty" xml:"outputFormat,omitempty"`
	// example:
	//
	// Hive
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 100
	Partitions *int32 `json:"partitions,omitempty" xml:"partitions,omitempty"`
	// example:
	//
	// 20
	PartitionsDoing *int32 `json:"partitionsDoing,omitempty" xml:"partitionsDoing,omitempty"`
	// example:
	//
	// 60
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 40
	PartitionsFailed *int32                                         `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	Schema           *ListMmsTablesResponseBodyDataObjectListSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// serde
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
	Serde *string `json:"serde,omitempty" xml:"serde,omitempty"`
	// example:
	//
	// 2985028
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// example:
	//
	// 2000028
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// MANAGED_TABLE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListMmsTablesResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetDbName() *string {
	return s.DbName
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetExtra() *string {
	return s.Extra
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetHasPartitions() *bool {
	return s.HasPartitions
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetInputFormat() *string {
	return s.InputFormat
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetLocation() *string {
	return s.Location
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetName() *string {
	return s.Name
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetNumRows() *int64 {
	return s.NumRows
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetOwner() *string {
	return s.Owner
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetPartitions() *int32 {
	return s.Partitions
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetPartitionsDoing() *int32 {
	return s.PartitionsDoing
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetPartitionsDone() *int32 {
	return s.PartitionsDone
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetPartitionsFailed() *int32 {
	return s.PartitionsFailed
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetSchema() *ListMmsTablesResponseBodyDataObjectListSchema {
	return s.Schema
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetSerde() *string {
	return s.Serde
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetSize() *int64 {
	return s.Size
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetSourceName() *string {
	return s.SourceName
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetType() *string {
	return s.Type
}

func (s *ListMmsTablesResponseBodyDataObjectList) GetUpdated() *bool {
	return s.Updated
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetDbId(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.DbId = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetDbName(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.DbName = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetExtra(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Extra = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetHasPartitions(v bool) *ListMmsTablesResponseBodyDataObjectList {
	s.HasPartitions = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetId(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetInputFormat(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.InputFormat = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetLastDdlTime(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.LastDdlTime = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetLocation(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Location = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetName(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetNumRows(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.NumRows = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetOutputFormat(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.OutputFormat = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetOwner(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Owner = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitions(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.Partitions = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitionsDoing(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.PartitionsDoing = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitionsDone(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.PartitionsDone = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetPartitionsFailed(v int32) *ListMmsTablesResponseBodyDataObjectList {
	s.PartitionsFailed = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSchema(v *ListMmsTablesResponseBodyDataObjectListSchema) *ListMmsTablesResponseBodyDataObjectList {
	s.Schema = v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSerde(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Serde = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSize(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.Size = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSourceId(v int64) *ListMmsTablesResponseBodyDataObjectList {
	s.SourceId = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetSourceName(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.SourceName = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetStatus(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetType(v string) *ListMmsTablesResponseBodyDataObjectList {
	s.Type = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) SetUpdated(v bool) *ListMmsTablesResponseBodyDataObjectList {
	s.Updated = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectList) Validate() error {
	return dara.Validate(s)
}

type ListMmsTablesResponseBodyDataObjectListSchema struct {
	Columns []*ListMmsTablesResponseBodyDataObjectListSchemaColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// example:
	//
	// for mms test
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// test
	Name       *string                                                    `json:"name,omitempty" xml:"name,omitempty"`
	Partitions []*ListMmsTablesResponseBodyDataObjectListSchemaPartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s ListMmsTablesResponseBodyDataObjectListSchema) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectListSchema) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) GetColumns() []*ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	return s.Columns
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) GetComment() *string {
	return s.Comment
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) GetName() *string {
	return s.Name
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) GetPartitions() []*ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	return s.Partitions
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetColumns(v []*ListMmsTablesResponseBodyDataObjectListSchemaColumns) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Columns = v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetComment(v string) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Comment = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetName(v string) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) SetPartitions(v []*ListMmsTablesResponseBodyDataObjectListSchemaPartitions) *ListMmsTablesResponseBodyDataObjectListSchema {
	s.Partitions = v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchema) Validate() error {
	return dara.Validate(s)
}

type ListMmsTablesResponseBodyDataObjectListSchemaColumns struct {
	// example:
	//
	// user id
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// ""
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// user_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// bigint
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaColumns) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) GetName() *string {
	return s.Name
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) GetNullable() *bool {
	return s.Nullable
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) GetType() *string {
	return s.Type
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetComment(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Comment = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetDefaultValue(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.DefaultValue = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetName(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetNullable(v bool) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Nullable = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) SetType(v string) *ListMmsTablesResponseBodyDataObjectListSchemaColumns {
	s.Type = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaColumns) Validate() error {
	return dara.Validate(s)
}

type ListMmsTablesResponseBodyDataObjectListSchemaPartitions struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// abc
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// p1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Nullable *bool `json:"nullable,omitempty" xml:"nullable,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaPartitions) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GoString() string {
	return s.String()
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GetComment() *string {
	return s.Comment
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GetName() *string {
	return s.Name
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GetNullable() *bool {
	return s.Nullable
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) GetType() *string {
	return s.Type
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetComment(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Comment = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetDefaultValue(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.DefaultValue = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetName(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Name = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetNullable(v bool) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Nullable = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) SetType(v string) *ListMmsTablesResponseBodyDataObjectListSchemaPartitions {
	s.Type = &v
	return s
}

func (s *ListMmsTablesResponseBodyDataObjectListSchemaPartitions) Validate() error {
	return dara.Validate(s)
}
