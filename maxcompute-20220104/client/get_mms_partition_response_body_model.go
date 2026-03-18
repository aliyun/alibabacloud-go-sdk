// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsPartitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsPartitionResponseBodyData) *GetMmsPartitionResponseBody
	GetData() *GetMmsPartitionResponseBodyData
	SetRequestId(v string) *GetMmsPartitionResponseBody
	GetRequestId() *string
}

type GetMmsPartitionResponseBody struct {
	// The partition object.
	Data *GetMmsPartitionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D9F872FD-5DDE-30A6-8C8A-1B8C6A81059F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsPartitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsPartitionResponseBody) GetData() *GetMmsPartitionResponseBodyData {
	return s.Data
}

func (s *GetMmsPartitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsPartitionResponseBody) SetData(v *GetMmsPartitionResponseBodyData) *GetMmsPartitionResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsPartitionResponseBody) SetRequestId(v string) *GetMmsPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsPartitionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsPartitionResponseBodyData struct {
	// The ID of the database that contains the partition. The name for this metadata layer depends on the data source. For example, it is a Dataset in BigQuery, a Database in Hive, or a Schema in Databricks and MaxCompute.
	//
	// example:
	//
	// 12
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// The name of the database. The name for this metadata layer depends on the data source. For example, it is a Dataset in BigQuery, a Database in Hive, or a Schema in Databricks and MaxCompute.
	//
	// example:
	//
	// d1
	DbName *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	// The name of the destination MaxCompute project.
	//
	// example:
	//
	// mms_test
	DstProjectName *string `json:"dstProjectName,omitempty" xml:"dstProjectName,omitempty"`
	// The name of the destination MaxCompute schema. This parameter is null if the destination MaxCompute project does not use a schema layer.
	//
	// example:
	//
	// default
	DstSchemaName *string `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	// The name of the destination MaxCompute table.
	//
	// example:
	//
	// test
	DstTableName *string `json:"dstTableName,omitempty" xml:"dstTableName,omitempty"`
	// The value of the destination MaxCompute partition. By default, this value is the same as the source partition value.
	//
	// example:
	//
	// p1=1/p2=abc
	DstValue *string `json:"dstValue,omitempty" xml:"dstValue,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 2323
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The time of the last DDL operation.
	//
	// example:
	//
	// 2024-12-17 15:44:42
	LastDdlTime *string `json:"lastDdlTime,omitempty" xml:"lastDdlTime,omitempty"`
	// The number of rows.
	//
	// example:
	//
	// 2323
	NumRows *int64 `json:"numRows,omitempty" xml:"numRows,omitempty"`
	// The data size in bytes.
	//
	// example:
	//
	// 12323
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 200018
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source name.
	//
	// example:
	//
	// demo
	SourceName *string `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	// The migration status.
	//
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The table ID.
	//
	// example:
	//
	// 23
	TableId *int64 `json:"tableId,omitempty" xml:"tableId,omitempty"`
	// The table name.
	//
	// example:
	//
	// t1
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// Indicates whether the metadata is updated.
	//
	// example:
	//
	// false
	Updated *bool `json:"updated,omitempty" xml:"updated,omitempty"`
	// The partition value.
	//
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMmsPartitionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsPartitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsPartitionResponseBodyData) GetDbId() *int64 {
	return s.DbId
}

func (s *GetMmsPartitionResponseBodyData) GetDbName() *string {
	return s.DbName
}

func (s *GetMmsPartitionResponseBodyData) GetDstProjectName() *string {
	return s.DstProjectName
}

func (s *GetMmsPartitionResponseBodyData) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *GetMmsPartitionResponseBodyData) GetDstTableName() *string {
	return s.DstTableName
}

func (s *GetMmsPartitionResponseBodyData) GetDstValue() *string {
	return s.DstValue
}

func (s *GetMmsPartitionResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsPartitionResponseBodyData) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *GetMmsPartitionResponseBodyData) GetNumRows() *int64 {
	return s.NumRows
}

func (s *GetMmsPartitionResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetMmsPartitionResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsPartitionResponseBodyData) GetSourceName() *string {
	return s.SourceName
}

func (s *GetMmsPartitionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsPartitionResponseBodyData) GetTableId() *int64 {
	return s.TableId
}

func (s *GetMmsPartitionResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetMmsPartitionResponseBodyData) GetUpdated() *bool {
	return s.Updated
}

func (s *GetMmsPartitionResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *GetMmsPartitionResponseBodyData) SetDbId(v int64) *GetMmsPartitionResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetDbName(v string) *GetMmsPartitionResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetDstProjectName(v string) *GetMmsPartitionResponseBodyData {
	s.DstProjectName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetDstSchemaName(v string) *GetMmsPartitionResponseBodyData {
	s.DstSchemaName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetDstTableName(v string) *GetMmsPartitionResponseBodyData {
	s.DstTableName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetDstValue(v string) *GetMmsPartitionResponseBodyData {
	s.DstValue = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetId(v int64) *GetMmsPartitionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetLastDdlTime(v string) *GetMmsPartitionResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetNumRows(v int64) *GetMmsPartitionResponseBodyData {
	s.NumRows = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetSize(v int64) *GetMmsPartitionResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetSourceId(v int64) *GetMmsPartitionResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetSourceName(v string) *GetMmsPartitionResponseBodyData {
	s.SourceName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetStatus(v string) *GetMmsPartitionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetTableId(v int64) *GetMmsPartitionResponseBodyData {
	s.TableId = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetTableName(v string) *GetMmsPartitionResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetUpdated(v bool) *GetMmsPartitionResponseBodyData {
	s.Updated = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) SetValue(v string) *GetMmsPartitionResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetMmsPartitionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
