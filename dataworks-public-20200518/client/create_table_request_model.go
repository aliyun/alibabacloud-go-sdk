// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGuid(v string) *CreateTableRequest
	GetAppGuid() *string
	SetCategoryId(v int64) *CreateTableRequest
	GetCategoryId() *int64
	SetClientToken(v string) *CreateTableRequest
	GetClientToken() *string
	SetColumns(v []*CreateTableRequestColumns) *CreateTableRequest
	GetColumns() []*CreateTableRequestColumns
	SetComment(v string) *CreateTableRequest
	GetComment() *string
	SetEndpoint(v string) *CreateTableRequest
	GetEndpoint() *string
	SetEnvType(v int32) *CreateTableRequest
	GetEnvType() *int32
	SetExternalTableType(v string) *CreateTableRequest
	GetExternalTableType() *string
	SetHasPart(v int32) *CreateTableRequest
	GetHasPart() *int32
	SetIsView(v int32) *CreateTableRequest
	GetIsView() *int32
	SetLifeCycle(v int32) *CreateTableRequest
	GetLifeCycle() *int32
	SetLocation(v string) *CreateTableRequest
	GetLocation() *string
	SetLogicalLevelId(v int64) *CreateTableRequest
	GetLogicalLevelId() *int64
	SetOwnerId(v string) *CreateTableRequest
	GetOwnerId() *string
	SetPhysicsLevelId(v int64) *CreateTableRequest
	GetPhysicsLevelId() *int64
	SetProjectId(v int64) *CreateTableRequest
	GetProjectId() *int64
	SetSchema(v string) *CreateTableRequest
	GetSchema() *string
	SetTableName(v string) *CreateTableRequest
	GetTableName() *string
	SetThemes(v []*CreateTableRequestThemes) *CreateTableRequest
	GetThemes() []*CreateTableRequestThemes
	SetVisibility(v int32) *CreateTableRequest
	GetVisibility() *int32
}

type CreateTableRequest struct {
	// The ID of the MaxCompute project. Specify the ID in the odps.{projectName} format.
	//
	// example:
	//
	// odps.test
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
	// The ID of the associated category. You can call the [GetMetaCategory](https://help.aliyun.com/document_detail/173932.html) operation to query the IDs of all categories that can be associated.
	//
	// example:
	//
	// 101
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// reserved
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of fields. A maximum of 1,000 fields are supported.
	//
	// This parameter is required.
	Columns []*CreateTableRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The endpoint of MaxCompute.
	//
	// example:
	//
	// odps://abc
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The environment type of the DataWorks workspace. Valid values:
	//
	// 	- 0: development environment
	//
	// 	- 1: production environment
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The storage type of the external table. Valid values:
	//
	// 	- 0: Object Storage Service (OSS)
	//
	// 	- 1: Tablestore
	//
	// 	- 2: Volume
	//
	// 	- 3: MySQL
	//
	// example:
	//
	// 0
	ExternalTableType *string `json:"ExternalTableType,omitempty" xml:"ExternalTableType,omitempty"`
	// Specifies whether to create a MaxCompute partitioned table. Valid values: 1 and 0. The value 1 indicates a partitioned table. The value 0 indicates a non-partitioned table. This parameter is deprecated. Do not use this parameter. The Column.N.isPartitionCol parameter is used to specify whether to create a MaxCompute partitioned table. If the Column.N.isPartitionCol parameter is set to true, a MaxCompute partitioned table is created.
	//
	// example:
	//
	// 0
	HasPart *int32 `json:"HasPart,omitempty" xml:"HasPart,omitempty"`
	// Specifies whether to create a view or table. Valid values:
	//
	// 	- 0: Create a table.
	//
	// 	- 1: Create a view.
	//
	// example:
	//
	// 0
	IsView *int32 `json:"IsView,omitempty" xml:"IsView,omitempty"`
	// The lifecycle of the table. Unit: days. By default, this parameter is left empty, which indicates that the table is permanently stored.
	//
	// example:
	//
	// 10
	LifeCycle *int32 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// The storage location of the external table.
	//
	// example:
	//
	// location
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The ID of the logical level.
	//
	// example:
	//
	// 101
	LogicalLevelId *int64  `json:"LogicalLevelId,omitempty" xml:"LogicalLevelId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the physical level.
	//
	// example:
	//
	// 101
	PhysicsLevelId *int64 `json:"PhysicsLevelId,omitempty" xml:"PhysicsLevelId,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 23
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The schema information of the table. You need to enter the schema information of the table if you enable the table schema in MaxCompute.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// default
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// tableName1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The list of themes.
	Themes []*CreateTableRequestThemes `json:"Themes,omitempty" xml:"Themes,omitempty" type:"Repeated"`
	// Specifies whether the table or workspace is visible. Valid values:
	//
	// 	- 0: Both the table and workspace are invisible.
	//
	// 	- 1: Both the table and workspace are visible.
	//
	// 	- 2: Only the workspace is visible.
	//
	// example:
	//
	// 1
	Visibility *int32 `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTableRequest) GetAppGuid() *string {
	return s.AppGuid
}

func (s *CreateTableRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTableRequest) GetColumns() []*CreateTableRequestColumns {
	return s.Columns
}

func (s *CreateTableRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateTableRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateTableRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *CreateTableRequest) GetExternalTableType() *string {
	return s.ExternalTableType
}

func (s *CreateTableRequest) GetHasPart() *int32 {
	return s.HasPart
}

func (s *CreateTableRequest) GetIsView() *int32 {
	return s.IsView
}

func (s *CreateTableRequest) GetLifeCycle() *int32 {
	return s.LifeCycle
}

func (s *CreateTableRequest) GetLocation() *string {
	return s.Location
}

func (s *CreateTableRequest) GetLogicalLevelId() *int64 {
	return s.LogicalLevelId
}

func (s *CreateTableRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateTableRequest) GetPhysicsLevelId() *int64 {
	return s.PhysicsLevelId
}

func (s *CreateTableRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateTableRequest) GetSchema() *string {
	return s.Schema
}

func (s *CreateTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateTableRequest) GetThemes() []*CreateTableRequestThemes {
	return s.Themes
}

func (s *CreateTableRequest) GetVisibility() *int32 {
	return s.Visibility
}

func (s *CreateTableRequest) SetAppGuid(v string) *CreateTableRequest {
	s.AppGuid = &v
	return s
}

func (s *CreateTableRequest) SetCategoryId(v int64) *CreateTableRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTableRequest) SetClientToken(v string) *CreateTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTableRequest) SetColumns(v []*CreateTableRequestColumns) *CreateTableRequest {
	s.Columns = v
	return s
}

func (s *CreateTableRequest) SetComment(v string) *CreateTableRequest {
	s.Comment = &v
	return s
}

func (s *CreateTableRequest) SetEndpoint(v string) *CreateTableRequest {
	s.Endpoint = &v
	return s
}

func (s *CreateTableRequest) SetEnvType(v int32) *CreateTableRequest {
	s.EnvType = &v
	return s
}

func (s *CreateTableRequest) SetExternalTableType(v string) *CreateTableRequest {
	s.ExternalTableType = &v
	return s
}

func (s *CreateTableRequest) SetHasPart(v int32) *CreateTableRequest {
	s.HasPart = &v
	return s
}

func (s *CreateTableRequest) SetIsView(v int32) *CreateTableRequest {
	s.IsView = &v
	return s
}

func (s *CreateTableRequest) SetLifeCycle(v int32) *CreateTableRequest {
	s.LifeCycle = &v
	return s
}

func (s *CreateTableRequest) SetLocation(v string) *CreateTableRequest {
	s.Location = &v
	return s
}

func (s *CreateTableRequest) SetLogicalLevelId(v int64) *CreateTableRequest {
	s.LogicalLevelId = &v
	return s
}

func (s *CreateTableRequest) SetOwnerId(v string) *CreateTableRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTableRequest) SetPhysicsLevelId(v int64) *CreateTableRequest {
	s.PhysicsLevelId = &v
	return s
}

func (s *CreateTableRequest) SetProjectId(v int64) *CreateTableRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateTableRequest) SetSchema(v string) *CreateTableRequest {
	s.Schema = &v
	return s
}

func (s *CreateTableRequest) SetTableName(v string) *CreateTableRequest {
	s.TableName = &v
	return s
}

func (s *CreateTableRequest) SetThemes(v []*CreateTableRequestThemes) *CreateTableRequest {
	s.Themes = v
	return s
}

func (s *CreateTableRequest) SetVisibility(v int32) *CreateTableRequest {
	s.Visibility = &v
	return s
}

func (s *CreateTableRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestColumns struct {
	// The name of the field. You can configure a maximum of 1,000 fields when you call the CreateTable operation to create a table.
	//
	// This parameter is required.
	//
	// example:
	//
	// columnName1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// columnName in chinese
	ColumnNameCn *string `json:"ColumnNameCn,omitempty" xml:"ColumnNameCn,omitempty"`
	// The data type of the field. For information about supported data types, see [Data type editions](https://help.aliyun.com/document_detail/27821.html) in MaxCompute documentation.
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// The comment of the field.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether the field is a partition field.
	//
	// example:
	//
	// true
	IsPartitionCol *bool `json:"IsPartitionCol,omitempty" xml:"IsPartitionCol,omitempty"`
	// The length of the field. For more information, see [MaxCompute data type editions](https://help.aliyun.com/document_detail/159541.html).
	//
	// example:
	//
	// 10
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The sequence number of the field. You can use this parameter to specify how fields are sorted in a table. By default, fields are sorted based on the order in which requests are created. If the field is a partition field, this parameter is not supported.
	//
	// example:
	//
	// 1
	SeqNumber *int32 `json:"SeqNumber,omitempty" xml:"SeqNumber,omitempty"`
}

func (s CreateTableRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateTableRequestColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *CreateTableRequestColumns) GetColumnNameCn() *string {
	return s.ColumnNameCn
}

func (s *CreateTableRequestColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *CreateTableRequestColumns) GetComment() *string {
	return s.Comment
}

func (s *CreateTableRequestColumns) GetIsPartitionCol() *bool {
	return s.IsPartitionCol
}

func (s *CreateTableRequestColumns) GetLength() *int32 {
	return s.Length
}

func (s *CreateTableRequestColumns) GetSeqNumber() *int32 {
	return s.SeqNumber
}

func (s *CreateTableRequestColumns) SetColumnName(v string) *CreateTableRequestColumns {
	s.ColumnName = &v
	return s
}

func (s *CreateTableRequestColumns) SetColumnNameCn(v string) *CreateTableRequestColumns {
	s.ColumnNameCn = &v
	return s
}

func (s *CreateTableRequestColumns) SetColumnType(v string) *CreateTableRequestColumns {
	s.ColumnType = &v
	return s
}

func (s *CreateTableRequestColumns) SetComment(v string) *CreateTableRequestColumns {
	s.Comment = &v
	return s
}

func (s *CreateTableRequestColumns) SetIsPartitionCol(v bool) *CreateTableRequestColumns {
	s.IsPartitionCol = &v
	return s
}

func (s *CreateTableRequestColumns) SetLength(v int32) *CreateTableRequestColumns {
	s.Length = &v
	return s
}

func (s *CreateTableRequestColumns) SetSeqNumber(v int32) *CreateTableRequestColumns {
	s.SeqNumber = &v
	return s
}

func (s *CreateTableRequestColumns) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestThemes struct {
	// The theme ID.
	//
	// example:
	//
	// 101
	ThemeId *int64 `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The level that corresponds to the theme ID.
	//
	// example:
	//
	// 101
	ThemeLevel *int32 `json:"ThemeLevel,omitempty" xml:"ThemeLevel,omitempty"`
}

func (s CreateTableRequestThemes) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestThemes) GoString() string {
	return s.String()
}

func (s *CreateTableRequestThemes) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *CreateTableRequestThemes) GetThemeLevel() *int32 {
	return s.ThemeLevel
}

func (s *CreateTableRequestThemes) SetThemeId(v int64) *CreateTableRequestThemes {
	s.ThemeId = &v
	return s
}

func (s *CreateTableRequestThemes) SetThemeLevel(v int32) *CreateTableRequestThemes {
	s.ThemeLevel = &v
	return s
}

func (s *CreateTableRequestThemes) Validate() error {
	return dara.Validate(s)
}
