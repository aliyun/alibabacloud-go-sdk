// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsTableResponseBodyData) *GetMmsTableResponseBody
	GetData() *GetMmsTableResponseBodyData
	SetRequestId(v string) *GetMmsTableResponseBody
	GetRequestId() *string
}

type GetMmsTableResponseBody struct {
	Data *GetMmsTableResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// E7FB14F1-4ACD-5C73-A755-B302D70AB9AD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBody) GetData() *GetMmsTableResponseBodyData {
	return s.Data
}

func (s *GetMmsTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsTableResponseBody) SetData(v *GetMmsTableResponseBodyData) *GetMmsTableResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsTableResponseBody) SetRequestId(v string) *GetMmsTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTableResponseBodyData struct {
	// example:
	//
	// 3
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// mms_test
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// example:
	//
	// test
	DstName *string `json:"dstName,omitempty" xml:"dstName,omitempty"`
	// example:
	//
	// mms_test
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
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
	// 22
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
	// 233232
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// outputFormat
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
	// 80
	PartitionsDone *int32 `json:"partitionsDone,omitempty" xml:"partitionsDone,omitempty"`
	// example:
	//
	// 0
	PartitionsFailed *int32                             `json:"partitionsFailed,omitempty" xml:"partitionsFailed,omitempty"`
	Schema           *GetMmsTableResponseBodyDataSchema `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	// serde
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
	Serde *string `json:"serde,omitempty" xml:"serde,omitempty"`
	// example:
	//
	// 23232
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
	// MANAGED_TABLED
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s GetMmsTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyData) GetDbId() *int64 {
	return s.DbId
}

func (s *GetMmsTableResponseBodyData) GetDbName() *string {
	return s.DbName
}

func (s *GetMmsTableResponseBodyData) GetDstName() *string {
	return s.DstName
}

func (s *GetMmsTableResponseBodyData) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *GetMmsTableResponseBodyData) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *GetMmsTableResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetMmsTableResponseBodyData) GetHasPartitions() *bool {
	return s.HasPartitions
}

func (s *GetMmsTableResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsTableResponseBodyData) GetInputFormat() *string {
	return s.InputFormat
}

func (s *GetMmsTableResponseBodyData) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *GetMmsTableResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *GetMmsTableResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMmsTableResponseBodyData) GetNumRows() *int64 {
	return s.NumRows
}

func (s *GetMmsTableResponseBodyData) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *GetMmsTableResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetMmsTableResponseBodyData) GetPartitions() *int32 {
	return s.Partitions
}

func (s *GetMmsTableResponseBodyData) GetPartitionsDoing() *int32 {
	return s.PartitionsDoing
}

func (s *GetMmsTableResponseBodyData) GetPartitionsDone() *int32 {
	return s.PartitionsDone
}

func (s *GetMmsTableResponseBodyData) GetPartitionsFailed() *int32 {
	return s.PartitionsFailed
}

func (s *GetMmsTableResponseBodyData) GetSchema() *GetMmsTableResponseBodyDataSchema {
	return s.Schema
}

func (s *GetMmsTableResponseBodyData) GetSerde() *string {
	return s.Serde
}

func (s *GetMmsTableResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetMmsTableResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsTableResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *GetMmsTableResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsTableResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsTableResponseBodyData) GetUpdated() *bool {
	return s.Updated
}

func (s *GetMmsTableResponseBodyData) SetDbId(v int64) *GetMmsTableResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetDbName(v string) *GetMmsTableResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetDstName(v string) *GetMmsTableResponseBodyData {
	s.DstName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetDstProjectName(v string) *GetMmsTableResponseBodyData {
	s.DstProjectName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetDstSchemaName(v string) *GetMmsTableResponseBodyData {
	s.DstSchemaName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetExtra(v string) *GetMmsTableResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetHasPartitions(v bool) *GetMmsTableResponseBodyData {
	s.HasPartitions = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetId(v int64) *GetMmsTableResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetInputFormat(v string) *GetMmsTableResponseBodyData {
	s.InputFormat = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetLastDdlTime(v string) *GetMmsTableResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetLocation(v string) *GetMmsTableResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetName(v string) *GetMmsTableResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetNumRows(v int64) *GetMmsTableResponseBodyData {
	s.NumRows = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetOutputFormat(v string) *GetMmsTableResponseBodyData {
	s.OutputFormat = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetOwner(v string) *GetMmsTableResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitions(v int32) *GetMmsTableResponseBodyData {
	s.Partitions = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitionsDoing(v int32) *GetMmsTableResponseBodyData {
	s.PartitionsDoing = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitionsDone(v int32) *GetMmsTableResponseBodyData {
	s.PartitionsDone = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetPartitionsFailed(v int32) *GetMmsTableResponseBodyData {
	s.PartitionsFailed = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSchema(v *GetMmsTableResponseBodyDataSchema) *GetMmsTableResponseBodyData {
	s.Schema = v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSerde(v string) *GetMmsTableResponseBodyData {
	s.Serde = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSize(v int64) *GetMmsTableResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSourceId(v int64) *GetMmsTableResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetSourceName(v string) *GetMmsTableResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetStatus(v string) *GetMmsTableResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetType(v string) *GetMmsTableResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsTableResponseBodyData) SetUpdated(v bool) *GetMmsTableResponseBodyData {
	s.Updated = &v
	return s
}

func (s *GetMmsTableResponseBodyData) Validate() error {
	if s.Schema != nil {
		if err := s.Schema.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTableResponseBodyDataSchema struct {
	Columns []*GetMmsTableResponseBodyDataSchemaColumns `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// example:
	//
	// for mms test
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// test
	Name       *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	Partitions []*GetMmsTableResponseBodyDataSchemaPartitions `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s GetMmsTableResponseBodyDataSchema) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTableResponseBodyDataSchema) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyDataSchema) GetColumns() []*GetMmsTableResponseBodyDataSchemaColumns {
	return s.Columns
}

func (s *GetMmsTableResponseBodyDataSchema) GetComment() *string {
	return s.Comment
}

func (s *GetMmsTableResponseBodyDataSchema) GetName() *string {
	return s.Name
}

func (s *GetMmsTableResponseBodyDataSchema) GetPartitions() []*GetMmsTableResponseBodyDataSchemaPartitions {
	return s.Partitions
}

func (s *GetMmsTableResponseBodyDataSchema) SetColumns(v []*GetMmsTableResponseBodyDataSchemaColumns) *GetMmsTableResponseBodyDataSchema {
	s.Columns = v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) SetComment(v string) *GetMmsTableResponseBodyDataSchema {
	s.Comment = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) SetName(v string) *GetMmsTableResponseBodyDataSchema {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) SetPartitions(v []*GetMmsTableResponseBodyDataSchemaPartitions) *GetMmsTableResponseBodyDataSchema {
	s.Partitions = v
	return s
}

func (s *GetMmsTableResponseBodyDataSchema) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Partitions != nil {
		for _, item := range s.Partitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMmsTableResponseBodyDataSchemaColumns struct {
	// example:
	//
	// user id
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// 10
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

func (s GetMmsTableResponseBodyDataSchemaColumns) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTableResponseBodyDataSchemaColumns) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) GetComment() *string {
	return s.Comment
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) GetName() *string {
	return s.Name
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) GetNullable() *bool {
	return s.Nullable
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) GetType() *string {
	return s.Type
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetComment(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Comment = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetDefaultValue(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.DefaultValue = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetName(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetNullable(v bool) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Nullable = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) SetType(v string) *GetMmsTableResponseBodyDataSchemaColumns {
	s.Type = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaColumns) Validate() error {
	return dara.Validate(s)
}

type GetMmsTableResponseBodyDataSchemaPartitions struct {
	// example:
	//
	// first partition level
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

func (s GetMmsTableResponseBodyDataSchemaPartitions) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTableResponseBodyDataSchemaPartitions) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) GetComment() *string {
	return s.Comment
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) GetName() *string {
	return s.Name
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) GetNullable() *bool {
	return s.Nullable
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) GetType() *string {
	return s.Type
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetComment(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Comment = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetDefaultValue(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.DefaultValue = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetName(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Name = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetNullable(v bool) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Nullable = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) SetType(v string) *GetMmsTableResponseBodyDataSchemaPartitions {
	s.Type = &v
	return s
}

func (s *GetMmsTableResponseBodyDataSchemaPartitions) Validate() error {
	return dara.Validate(s)
}
