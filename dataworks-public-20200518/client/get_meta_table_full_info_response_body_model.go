// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableFullInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableFullInfoResponseBodyData) *GetMetaTableFullInfoResponseBody
	GetData() *GetMetaTableFullInfoResponseBodyData
	SetErrorCode(v string) *GetMetaTableFullInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableFullInfoResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableFullInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableFullInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableFullInfoResponseBody
	GetSuccess() *bool
}

type GetMetaTableFullInfoResponseBody struct {
	// The business data.
	Data *GetMetaTableFullInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1411515937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableFullInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableFullInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableFullInfoResponseBody) GetData() *GetMetaTableFullInfoResponseBodyData {
	return s.Data
}

func (s *GetMetaTableFullInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableFullInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableFullInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableFullInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableFullInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableFullInfoResponseBody) SetData(v *GetMetaTableFullInfoResponseBodyData) *GetMetaTableFullInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableFullInfoResponseBody) SetErrorCode(v string) *GetMetaTableFullInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBody) SetErrorMessage(v string) *GetMetaTableFullInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBody) SetHttpStatusCode(v int32) *GetMetaTableFullInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBody) SetRequestId(v string) *GetMetaTableFullInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBody) SetSuccess(v bool) *GetMetaTableFullInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableFullInfoResponseBodyData struct {
	// The EMR cluster ID.
	//
	// example:
	//
	// C-010A704DA760****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The fields in the table.
	ColumnList []*GetMetaTableFullInfoResponseBodyDataColumnList `json:"ColumnList,omitempty" xml:"ColumnList,omitempty" type:"Repeated"`
	// The comment on the table.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the table was created. A timestamp is returned for this parameter. You can convert the timestamp to the related date based on the time zone that you use.
	//
	// example:
	//
	// 1589870293000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The size of the storage space that is consumed by the table. Unit: bytes.
	//
	// example:
	//
	// 10
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The type of the environment. Valid values:
	//
	// 	- 0: indicates that the table resides in the development environment.
	//
	// 	- 1: indicates that the table resides in the production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The scope in which the table is visible. Valid values:
	//
	// 	- 0: indicates that the table is visible to workspace members.
	//
	// 	- 1: indicates that the table is visible to users within a tenant.
	//
	// 	- 2: indicates that the table is visible to all tenants.
	//
	// 	- 3: indicates that the table is visible only to the table owner.
	//
	// example:
	//
	// 1
	IsVisible *int32 `json:"IsVisible,omitempty" xml:"IsVisible,omitempty"`
	// The time when the table was last accessed. A timestamp is returned for this parameter. You can convert the timestamp to the related date based on the time zone that you use.
	//
	// example:
	//
	// 1589870294000
	LastAccessTime *int64 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// The time when the schema of the table was last changed. A timestamp is returned for this parameter. You can convert the timestamp to the related date based on the time zone that you use.
	//
	// example:
	//
	// 1589870294000
	LastDdlTime *int64 `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// The time when the table was last updated. A timestamp is returned for this parameter. You can convert the timestamp to the related date based on the time zone that you use.
	//
	// example:
	//
	// 1589870294000
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The lifecycle of the table. Unit: days.
	//
	// example:
	//
	// 5
	LifeCycle *int32 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// The storage path of the Hive table.
	//
	// example:
	//
	// hdfs://localhost:777/user/hadoop/test.txt
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The ID of the table owner.
	//
	// example:
	//
	// 123
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The partition key column.
	//
	// example:
	//
	// abc
	PartitionKeys *string `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty"`
	// The ID of the workspace to which the table belongs.
	//
	// example:
	//
	// 22
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace to which the table belongs.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The schema information of the table.
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The unique identifier of the table.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table_name
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 12345
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The total number of fields.
	//
	// example:
	//
	// 22
	TotalColumnCount *int64 `json:"TotalColumnCount,omitempty" xml:"TotalColumnCount,omitempty"`
}

func (s GetMetaTableFullInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableFullInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableFullInfoResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableFullInfoResponseBodyData) GetColumnList() []*GetMetaTableFullInfoResponseBodyDataColumnList {
	return s.ColumnList
}

func (s *GetMetaTableFullInfoResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetMetaTableFullInfoResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaTableFullInfoResponseBodyData) GetDataSize() *int64 {
	return s.DataSize
}

func (s *GetMetaTableFullInfoResponseBodyData) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableFullInfoResponseBodyData) GetEnvType() *int32 {
	return s.EnvType
}

func (s *GetMetaTableFullInfoResponseBodyData) GetIsVisible() *int32 {
	return s.IsVisible
}

func (s *GetMetaTableFullInfoResponseBodyData) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *GetMetaTableFullInfoResponseBodyData) GetLastDdlTime() *int64 {
	return s.LastDdlTime
}

func (s *GetMetaTableFullInfoResponseBodyData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *GetMetaTableFullInfoResponseBodyData) GetLifeCycle() *int32 {
	return s.LifeCycle
}

func (s *GetMetaTableFullInfoResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *GetMetaTableFullInfoResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMetaTableFullInfoResponseBodyData) GetPartitionKeys() *string {
	return s.PartitionKeys
}

func (s *GetMetaTableFullInfoResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMetaTableFullInfoResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetMetaTableFullInfoResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetMetaTableFullInfoResponseBodyData) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableFullInfoResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableFullInfoResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetMetaTableFullInfoResponseBodyData) GetTotalColumnCount() *int64 {
	return s.TotalColumnCount
}

func (s *GetMetaTableFullInfoResponseBodyData) SetClusterId(v string) *GetMetaTableFullInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetColumnList(v []*GetMetaTableFullInfoResponseBodyDataColumnList) *GetMetaTableFullInfoResponseBodyData {
	s.ColumnList = v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetComment(v string) *GetMetaTableFullInfoResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetCreateTime(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetDataSize(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.DataSize = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetDatabaseName(v string) *GetMetaTableFullInfoResponseBodyData {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetEnvType(v int32) *GetMetaTableFullInfoResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetIsVisible(v int32) *GetMetaTableFullInfoResponseBodyData {
	s.IsVisible = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetLastAccessTime(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.LastAccessTime = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetLastDdlTime(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetLastModifyTime(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetLifeCycle(v int32) *GetMetaTableFullInfoResponseBodyData {
	s.LifeCycle = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetLocation(v string) *GetMetaTableFullInfoResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetOwnerId(v string) *GetMetaTableFullInfoResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetPartitionKeys(v string) *GetMetaTableFullInfoResponseBodyData {
	s.PartitionKeys = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetProjectId(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetProjectName(v string) *GetMetaTableFullInfoResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetSchema(v string) *GetMetaTableFullInfoResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetTableGuid(v string) *GetMetaTableFullInfoResponseBodyData {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetTableName(v string) *GetMetaTableFullInfoResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetTenantId(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) SetTotalColumnCount(v int64) *GetMetaTableFullInfoResponseBodyData {
	s.TotalColumnCount = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableFullInfoResponseBodyDataColumnList struct {
	// The description of the field.
	//
	// example:
	//
	// data comment
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The unique identifier of the field.
	//
	// example:
	//
	// odps.engine_name.table_name.1
	ColumnGuid *string `json:"ColumnGuid,omitempty" xml:"ColumnGuid,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// 1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The data type of the field.
	//
	// example:
	//
	// string
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The remarks of the field.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether the field is a foreign key. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsForeignKey *bool `json:"IsForeignKey,omitempty" xml:"IsForeignKey,omitempty"`
	// Indicates whether the field is a partition field. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsPartitionColumn *bool `json:"IsPartitionColumn,omitempty" xml:"IsPartitionColumn,omitempty"`
	// Indicates whether the field is a primary key. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsPrimaryKey *bool `json:"IsPrimaryKey,omitempty" xml:"IsPrimaryKey,omitempty"`
	// The sequence number of the field.
	//
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s GetMetaTableFullInfoResponseBodyDataColumnList) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableFullInfoResponseBodyDataColumnList) GoString() string {
	return s.String()
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetCaption() *string {
	return s.Caption
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetColumnGuid() *string {
	return s.ColumnGuid
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetComment() *string {
	return s.Comment
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetIsForeignKey() *bool {
	return s.IsForeignKey
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetIsPartitionColumn() *bool {
	return s.IsPartitionColumn
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetIsPrimaryKey() *bool {
	return s.IsPrimaryKey
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) GetPosition() *int32 {
	return s.Position
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetCaption(v string) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.Caption = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetColumnGuid(v string) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.ColumnGuid = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetColumnName(v string) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.ColumnName = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetColumnType(v string) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.ColumnType = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetComment(v string) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.Comment = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetIsForeignKey(v bool) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.IsForeignKey = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetIsPartitionColumn(v bool) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.IsPartitionColumn = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetIsPrimaryKey(v bool) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.IsPrimaryKey = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) SetPosition(v int32) *GetMetaTableFullInfoResponseBodyDataColumnList {
	s.Position = &v
	return s
}

func (s *GetMetaTableFullInfoResponseBodyDataColumnList) Validate() error {
	return dara.Validate(s)
}
