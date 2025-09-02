// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGuid(v string) *UpdateTableRequest
	GetAppGuid() *string
	SetCategoryId(v int64) *UpdateTableRequest
	GetCategoryId() *int64
	SetColumns(v []*UpdateTableRequestColumns) *UpdateTableRequest
	GetColumns() []*UpdateTableRequestColumns
	SetComment(v string) *UpdateTableRequest
	GetComment() *string
	SetCreateIfNotExists(v bool) *UpdateTableRequest
	GetCreateIfNotExists() *bool
	SetEndpoint(v string) *UpdateTableRequest
	GetEndpoint() *string
	SetEnvType(v int32) *UpdateTableRequest
	GetEnvType() *int32
	SetExternalTableType(v string) *UpdateTableRequest
	GetExternalTableType() *string
	SetHasPart(v int32) *UpdateTableRequest
	GetHasPart() *int32
	SetIsView(v int32) *UpdateTableRequest
	GetIsView() *int32
	SetLifeCycle(v int32) *UpdateTableRequest
	GetLifeCycle() *int32
	SetLocation(v string) *UpdateTableRequest
	GetLocation() *string
	SetLogicalLevelId(v int64) *UpdateTableRequest
	GetLogicalLevelId() *int64
	SetOwnerId(v string) *UpdateTableRequest
	GetOwnerId() *string
	SetPhysicsLevelId(v int64) *UpdateTableRequest
	GetPhysicsLevelId() *int64
	SetProjectId(v int64) *UpdateTableRequest
	GetProjectId() *int64
	SetSchema(v string) *UpdateTableRequest
	GetSchema() *string
	SetTableName(v string) *UpdateTableRequest
	GetTableName() *string
	SetThemes(v []*UpdateTableRequestThemes) *UpdateTableRequest
	GetThemes() []*UpdateTableRequestThemes
	SetVisibility(v int32) *UpdateTableRequest
	GetVisibility() *int32
}

type UpdateTableRequest struct {
	// The unique identifier of the MaxCompute project. Specify the GUID in the odps.{projectName} format.
	//
	// example:
	//
	// odps.test
	AppGuid *string `json:"AppGuid,omitempty" xml:"AppGuid,omitempty"`
	// The ID of the associated category.
	//
	// example:
	//
	// 101
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The list of fields.
	Columns []*UpdateTableRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The comment.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether the table exists. Valid values:
	//
	// 	- true: The table exists.
	//
	// 	- false: The table does not exist.
	//
	// This parameter is deprecated. Do not use this parameter.
	//
	// example:
	//
	// true
	CreateIfNotExists *bool `json:"CreateIfNotExists,omitempty" xml:"CreateIfNotExists,omitempty"`
	// The endpoint of MaxCompute. If you do not specify this parameter, the endpoint of the MaxCompute project is used.
	//
	// example:
	//
	// odps://
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The environment of the DataWorks workspace. Valid values: 0 and 1. The value 0 indicates the development environment. The value 1 indicates the production environment.
	//
	// example:
	//
	// 1
	EnvType *int32 `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The type of the external table. Valid values: 0, 1, 2, and 3. The value 0 indicates that the external table is an OSS external table. The value 1 indicates that the external table is a Tablestore external table. The value 2 indicates that the external table is a volume external table. The value 3 indicates that the external table is a MySQL external table. This parameter is deprecated. Do not use this parameter.
	//
	// example:
	//
	// 1
	ExternalTableType *string `json:"ExternalTableType,omitempty" xml:"ExternalTableType,omitempty"`
	// Specifies whether the table that you want to update is a partitioned table. Valid values: 0 and 1. The value 0 indicates that the table is not a partitioned table. The value 1 indicates that the table is a partitioned table. This parameter is deprecated. Do not use this parameter. The Column.N.isPartitionCol parameter is used instead of the HasPart parameter to specify whether the MaxCompute table is a partitioned table. If the Column.N.isPartitionCol parameter is set to 1, the MaxCompute table is a partitioned table.
	//
	// example:
	//
	// 0
	HasPart *int32 `json:"HasPart,omitempty" xml:"HasPart,omitempty"`
	// Specifies whether the table is a view. Valid values: 0 and 1. The value 0 indicates that the table is not a view. The value 1 indicates that the table is a view. This parameter is deprecated. Do not use this parameter.
	//
	// example:
	//
	// 0
	IsView *int32 `json:"IsView,omitempty" xml:"IsView,omitempty"`
	// The lifecycle of the table. Unit: days. If this parameter is left empty, the table is permanently stored.
	//
	// example:
	//
	// 10
	LifeCycle *int32 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// The storage location of the external table. This parameter is deprecated. Do not use this parameter.
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
	// The ID of the physical layer.
	//
	// example:
	//
	// 101
	PhysicsLevelId *int64 `json:"PhysicsLevelId,omitempty" xml:"PhysicsLevelId,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console to obtain the ID of the DataWorks workspace.
	//
	// example:
	//
	// 101
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
	// The name of the MaxCompute table.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The list of fields.
	Themes []*UpdateTableRequestThemes `json:"Themes,omitempty" xml:"Themes,omitempty" type:"Repeated"`
	// The scope in which the table is visible. Valid values: 0, 1, and 2. The value 0 indicates that the table is invisible to all workspace members. The value 1 indicates that the table is visible to all workspace members. The value 2 indicates that the table is visible to workspace members.
	//
	// example:
	//
	// 1
	Visibility *int32 `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableRequest) GetAppGuid() *string {
	return s.AppGuid
}

func (s *UpdateTableRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *UpdateTableRequest) GetColumns() []*UpdateTableRequestColumns {
	return s.Columns
}

func (s *UpdateTableRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateTableRequest) GetCreateIfNotExists() *bool {
	return s.CreateIfNotExists
}

func (s *UpdateTableRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateTableRequest) GetEnvType() *int32 {
	return s.EnvType
}

func (s *UpdateTableRequest) GetExternalTableType() *string {
	return s.ExternalTableType
}

func (s *UpdateTableRequest) GetHasPart() *int32 {
	return s.HasPart
}

func (s *UpdateTableRequest) GetIsView() *int32 {
	return s.IsView
}

func (s *UpdateTableRequest) GetLifeCycle() *int32 {
	return s.LifeCycle
}

func (s *UpdateTableRequest) GetLocation() *string {
	return s.Location
}

func (s *UpdateTableRequest) GetLogicalLevelId() *int64 {
	return s.LogicalLevelId
}

func (s *UpdateTableRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *UpdateTableRequest) GetPhysicsLevelId() *int64 {
	return s.PhysicsLevelId
}

func (s *UpdateTableRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateTableRequest) GetSchema() *string {
	return s.Schema
}

func (s *UpdateTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *UpdateTableRequest) GetThemes() []*UpdateTableRequestThemes {
	return s.Themes
}

func (s *UpdateTableRequest) GetVisibility() *int32 {
	return s.Visibility
}

func (s *UpdateTableRequest) SetAppGuid(v string) *UpdateTableRequest {
	s.AppGuid = &v
	return s
}

func (s *UpdateTableRequest) SetCategoryId(v int64) *UpdateTableRequest {
	s.CategoryId = &v
	return s
}

func (s *UpdateTableRequest) SetColumns(v []*UpdateTableRequestColumns) *UpdateTableRequest {
	s.Columns = v
	return s
}

func (s *UpdateTableRequest) SetComment(v string) *UpdateTableRequest {
	s.Comment = &v
	return s
}

func (s *UpdateTableRequest) SetCreateIfNotExists(v bool) *UpdateTableRequest {
	s.CreateIfNotExists = &v
	return s
}

func (s *UpdateTableRequest) SetEndpoint(v string) *UpdateTableRequest {
	s.Endpoint = &v
	return s
}

func (s *UpdateTableRequest) SetEnvType(v int32) *UpdateTableRequest {
	s.EnvType = &v
	return s
}

func (s *UpdateTableRequest) SetExternalTableType(v string) *UpdateTableRequest {
	s.ExternalTableType = &v
	return s
}

func (s *UpdateTableRequest) SetHasPart(v int32) *UpdateTableRequest {
	s.HasPart = &v
	return s
}

func (s *UpdateTableRequest) SetIsView(v int32) *UpdateTableRequest {
	s.IsView = &v
	return s
}

func (s *UpdateTableRequest) SetLifeCycle(v int32) *UpdateTableRequest {
	s.LifeCycle = &v
	return s
}

func (s *UpdateTableRequest) SetLocation(v string) *UpdateTableRequest {
	s.Location = &v
	return s
}

func (s *UpdateTableRequest) SetLogicalLevelId(v int64) *UpdateTableRequest {
	s.LogicalLevelId = &v
	return s
}

func (s *UpdateTableRequest) SetOwnerId(v string) *UpdateTableRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTableRequest) SetPhysicsLevelId(v int64) *UpdateTableRequest {
	s.PhysicsLevelId = &v
	return s
}

func (s *UpdateTableRequest) SetProjectId(v int64) *UpdateTableRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTableRequest) SetSchema(v string) *UpdateTableRequest {
	s.Schema = &v
	return s
}

func (s *UpdateTableRequest) SetTableName(v string) *UpdateTableRequest {
	s.TableName = &v
	return s
}

func (s *UpdateTableRequest) SetThemes(v []*UpdateTableRequestThemes) *UpdateTableRequest {
	s.Themes = v
	return s
}

func (s *UpdateTableRequest) SetVisibility(v int32) *UpdateTableRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateTableRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestColumns struct {
	// The name of the field.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// 名称
	ColumnNameCn *string `json:"ColumnNameCn,omitempty" xml:"ColumnNameCn,omitempty"`
	// The type of the field. For more information, see MaxCompute field types.
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
	// Specifies whether the field is a partition field. Valid values: 0 and 1. The value 0 indicates that the field is not a partition field. The value 1 indicates that the field is a partition field.
	//
	// example:
	//
	// 0
	IsPartitionCol *bool `json:"IsPartitionCol,omitempty" xml:"IsPartitionCol,omitempty"`
	// The length of the field.
	//
	// example:
	//
	// 10
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The sequence number of the field. If the field is a partition field, this parameter is not supported.
	//
	// example:
	//
	// 1
	SeqNumber *int32 `json:"SeqNumber,omitempty" xml:"SeqNumber,omitempty"`
}

func (s UpdateTableRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestColumns) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *UpdateTableRequestColumns) GetColumnNameCn() *string {
	return s.ColumnNameCn
}

func (s *UpdateTableRequestColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *UpdateTableRequestColumns) GetComment() *string {
	return s.Comment
}

func (s *UpdateTableRequestColumns) GetIsPartitionCol() *bool {
	return s.IsPartitionCol
}

func (s *UpdateTableRequestColumns) GetLength() *int32 {
	return s.Length
}

func (s *UpdateTableRequestColumns) GetSeqNumber() *int32 {
	return s.SeqNumber
}

func (s *UpdateTableRequestColumns) SetColumnName(v string) *UpdateTableRequestColumns {
	s.ColumnName = &v
	return s
}

func (s *UpdateTableRequestColumns) SetColumnNameCn(v string) *UpdateTableRequestColumns {
	s.ColumnNameCn = &v
	return s
}

func (s *UpdateTableRequestColumns) SetColumnType(v string) *UpdateTableRequestColumns {
	s.ColumnType = &v
	return s
}

func (s *UpdateTableRequestColumns) SetComment(v string) *UpdateTableRequestColumns {
	s.Comment = &v
	return s
}

func (s *UpdateTableRequestColumns) SetIsPartitionCol(v bool) *UpdateTableRequestColumns {
	s.IsPartitionCol = &v
	return s
}

func (s *UpdateTableRequestColumns) SetLength(v int32) *UpdateTableRequestColumns {
	s.Length = &v
	return s
}

func (s *UpdateTableRequestColumns) SetSeqNumber(v int32) *UpdateTableRequestColumns {
	s.SeqNumber = &v
	return s
}

func (s *UpdateTableRequestColumns) Validate() error {
	return dara.Validate(s)
}

type UpdateTableRequestThemes struct {
	// The ID of the associated topic.
	//
	// example:
	//
	// 101
	ThemeId *int64 `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
	// The level that corresponds to the topic ID.
	//
	// example:
	//
	// 101
	ThemeLevel *int32 `json:"ThemeLevel,omitempty" xml:"ThemeLevel,omitempty"`
}

func (s UpdateTableRequestThemes) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableRequestThemes) GoString() string {
	return s.String()
}

func (s *UpdateTableRequestThemes) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *UpdateTableRequestThemes) GetThemeLevel() *int32 {
	return s.ThemeLevel
}

func (s *UpdateTableRequestThemes) SetThemeId(v int64) *UpdateTableRequestThemes {
	s.ThemeId = &v
	return s
}

func (s *UpdateTableRequestThemes) SetThemeLevel(v int32) *UpdateTableRequestThemes {
	s.ThemeLevel = &v
	return s
}

func (s *UpdateTableRequestThemes) Validate() error {
	return dara.Validate(s)
}
