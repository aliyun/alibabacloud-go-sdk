// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaTableBasicInfoResponseBodyData) *GetMetaTableBasicInfoResponseBody
	GetData() *GetMetaTableBasicInfoResponseBodyData
	SetErrorCode(v string) *GetMetaTableBasicInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableBasicInfoResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableBasicInfoResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableBasicInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableBasicInfoResponseBody
	GetSuccess() *bool
}

type GetMetaTableBasicInfoResponseBody struct {
	// The business data.
	Data *GetMetaTableBasicInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
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
	// 0bc1411515937
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableBasicInfoResponseBody) GetData() *GetMetaTableBasicInfoResponseBodyData {
	return s.Data
}

func (s *GetMetaTableBasicInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableBasicInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableBasicInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableBasicInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableBasicInfoResponseBody) SetData(v *GetMetaTableBasicInfoResponseBodyData) *GetMetaTableBasicInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaTableBasicInfoResponseBody) SetErrorCode(v string) *GetMetaTableBasicInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBody) SetErrorMessage(v string) *GetMetaTableBasicInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBody) SetHttpStatusCode(v int32) *GetMetaTableBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBody) SetRequestId(v string) *GetMetaTableBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBody) SetSuccess(v bool) *GetMetaTableBasicInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableBasicInfoResponseBodyData struct {
	// The display name of the metatable.
	//
	// example:
	//
	// test
	Caption *string `json:"Caption,omitempty" xml:"Caption,omitempty"`
	// The ID of the EMR cluster.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of fields.
	//
	// example:
	//
	// 3
	ColumnCount *int32 `json:"ColumnCount,omitempty" xml:"ColumnCount,omitempty"`
	// The comment of the metatable.
	//
	// example:
	//
	// test table
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the metatable was created.
	//
	// example:
	//
	// 1589870294000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The size of storage space that is occupied by the metatable. Unit: bytes.
	//
	// example:
	//
	// 10
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The name of the metadatabase.
	//
	// example:
	//
	// abc
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The type of the environment. Valid values:
	//
	// 	- 0: development environment
	//
	// 	- 1: production environment
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The number of times the metatable was added to a favorite list. This parameter is returned only if you set the Extension parameter to true and takes effect only if you set the DataSourceType parameter to odps.
	//
	// example:
	//
	// 6
	FavoriteCount *int64 `json:"FavoriteCount,omitempty" xml:"FavoriteCount,omitempty"`
	// Indicates whether the metatable is a partitioned table. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsPartitionTable *bool `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// Indicates whether the metatable is a view. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsView *bool `json:"IsView,omitempty" xml:"IsView,omitempty"`
	// The scope in which the metatable is visible. Valid values:
	//
	// 	- 0: The metatable is visible to workspace members.
	//
	// 	- 1: The metatable is visible to users within the tenant.
	//
	// 	- 2: The metatable is visible to all tenants.
	//
	// 	- 3: The metatable is visible only to the metatable owner.
	//
	// example:
	//
	// 1
	IsVisible *int32 `json:"IsVisible,omitempty" xml:"IsVisible,omitempty"`
	// The time when the metatable was last accessed.
	//
	// example:
	//
	// 1589870294000
	LastAccessTime *int64 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	// The time when the schema of the metatable was last changed.
	//
	// example:
	//
	// 1589870294000
	LastDdlTime *int64 `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// The time when the metatable was last updated.
	//
	// example:
	//
	// 1589870294000
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The lifecycle of the table. Unit: day.
	//
	// >  If the lifecycle is not set for a MaxCompute table, the return value is 0, indicating that the table is permanently valid.
	//
	// example:
	//
	// 5
	LifeCycle *int32 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// The storage path of the Hive metadatabase.
	//
	// example:
	//
	// hdfs://
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The ID of the metatable owner.
	//
	// example:
	//
	// 123
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The partition key of the Hive metatable.
	//
	// example:
	//
	// ab
	PartitionKeys *string `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 232
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The number of times the metatable was read. This parameter is returned only if you set the Extension parameter to true and takes effect only if you set the DataSourceType parameter to odps.
	//
	// example:
	//
	// 3
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The schema information of the metatable. This parameter is returned if the three-layer model of MaxCompute is enabled.
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The GUID of the metatable.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the metatable.
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
	// The number of times the metatable was viewed. This parameter is returned only if you set the Extension parameter to true and takes effect only if you set the DataSourceType parameter to odps.
	//
	// example:
	//
	// 2
	ViewCount *int64 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
}

func (s GetMetaTableBasicInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableBasicInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetCaption() *string {
	return s.Caption
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetColumnCount() *int32 {
	return s.ColumnCount
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetDataSize() *int64 {
	return s.DataSize
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetEnvType() *int32 {
	return s.EnvType
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetFavoriteCount() *int64 {
	return s.FavoriteCount
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetIsView() *bool {
	return s.IsView
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetIsVisible() *int32 {
	return s.IsVisible
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetLastDdlTime() *int64 {
	return s.LastDdlTime
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetLifeCycle() *int32 {
	return s.LifeCycle
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetPartitionKeys() *string {
	return s.PartitionKeys
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetReadCount() *int64 {
	return s.ReadCount
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetMetaTableBasicInfoResponseBodyData) GetViewCount() *int64 {
	return s.ViewCount
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetCaption(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.Caption = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetClusterId(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetColumnCount(v int32) *GetMetaTableBasicInfoResponseBodyData {
	s.ColumnCount = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetComment(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetCreateTime(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetDataSize(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.DataSize = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetDatabaseName(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.DatabaseName = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetEnvType(v int32) *GetMetaTableBasicInfoResponseBodyData {
	s.EnvType = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetFavoriteCount(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.FavoriteCount = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetIsPartitionTable(v bool) *GetMetaTableBasicInfoResponseBodyData {
	s.IsPartitionTable = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetIsView(v bool) *GetMetaTableBasicInfoResponseBodyData {
	s.IsView = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetIsVisible(v int32) *GetMetaTableBasicInfoResponseBodyData {
	s.IsVisible = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetLastAccessTime(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.LastAccessTime = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetLastDdlTime(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetLastModifyTime(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetLifeCycle(v int32) *GetMetaTableBasicInfoResponseBodyData {
	s.LifeCycle = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetLocation(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetOwnerId(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.OwnerId = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetPartitionKeys(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.PartitionKeys = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetProjectId(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetProjectName(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetReadCount(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.ReadCount = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetSchema(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetTableGuid(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetTableName(v string) *GetMetaTableBasicInfoResponseBodyData {
	s.TableName = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetTenantId(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) SetViewCount(v int64) *GetMetaTableBasicInfoResponseBodyData {
	s.ViewCount = &v
	return s
}

func (s *GetMetaTableBasicInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
