// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableKnowledgeVO interface {
	dara.Model
	String() string
	GoString() string
	SetAssetCreatedGmt(v string) *TableKnowledgeVO
	GetAssetCreatedGmt() *string
	SetAssetDescription(v string) *TableKnowledgeVO
	GetAssetDescription() *string
	SetAssetModifiedGmt(v string) *TableKnowledgeVO
	GetAssetModifiedGmt() *string
	SetDbId(v int32) *TableKnowledgeVO
	GetDbId() *int32
	SetDbName(v string) *TableKnowledgeVO
	GetDbName() *string
	SetDbType(v string) *TableKnowledgeVO
	GetDbType() *string
	SetDescription(v string) *TableKnowledgeVO
	GetDescription() *string
	SetEnvType(v string) *TableKnowledgeVO
	GetEnvType() *string
	SetHotLevel(v int32) *TableKnowledgeVO
	GetHotLevel() *int32
	SetInstanceId(v int32) *TableKnowledgeVO
	GetInstanceId() *int32
	SetInstanceName(v string) *TableKnowledgeVO
	GetInstanceName() *string
	SetLevel(v int32) *TableKnowledgeVO
	GetLevel() *int32
	SetLevelType(v string) *TableKnowledgeVO
	GetLevelType() *string
	SetLogic(v bool) *TableKnowledgeVO
	GetLogic() *bool
	SetNumRows(v int64) *TableKnowledgeVO
	GetNumRows() *int64
	SetSchemaName(v string) *TableKnowledgeVO
	GetSchemaName() *string
	SetSize(v int64) *TableKnowledgeVO
	GetSize() *int64
	SetSummary(v string) *TableKnowledgeVO
	GetSummary() *string
	SetTableId(v int64) *TableKnowledgeVO
	GetTableId() *int64
	SetTableName(v string) *TableKnowledgeVO
	GetTableName() *string
	SetTitle(v string) *TableKnowledgeVO
	GetTitle() *string
}

type TableKnowledgeVO struct {
	AssetCreatedGmt  *string `json:"AssetCreatedGmt,omitempty" xml:"AssetCreatedGmt,omitempty"`
	AssetDescription *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	AssetModifiedGmt *string `json:"AssetModifiedGmt,omitempty" xml:"AssetModifiedGmt,omitempty"`
	DbId             *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	DbName           *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DbType           *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvType          *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	HotLevel         *int32  `json:"HotLevel,omitempty" xml:"HotLevel,omitempty"`
	InstanceId       *int32  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Level            *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	LevelType        *string `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	Logic            *bool   `json:"Logic,omitempty" xml:"Logic,omitempty"`
	NumRows          *int64  `json:"NumRows,omitempty" xml:"NumRows,omitempty"`
	SchemaName       *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Summary          *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TableId          *int64  `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName        *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s TableKnowledgeVO) String() string {
	return dara.Prettify(s)
}

func (s TableKnowledgeVO) GoString() string {
	return s.String()
}

func (s *TableKnowledgeVO) GetAssetCreatedGmt() *string {
	return s.AssetCreatedGmt
}

func (s *TableKnowledgeVO) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *TableKnowledgeVO) GetAssetModifiedGmt() *string {
	return s.AssetModifiedGmt
}

func (s *TableKnowledgeVO) GetDbId() *int32 {
	return s.DbId
}

func (s *TableKnowledgeVO) GetDbName() *string {
	return s.DbName
}

func (s *TableKnowledgeVO) GetDbType() *string {
	return s.DbType
}

func (s *TableKnowledgeVO) GetDescription() *string {
	return s.Description
}

func (s *TableKnowledgeVO) GetEnvType() *string {
	return s.EnvType
}

func (s *TableKnowledgeVO) GetHotLevel() *int32 {
	return s.HotLevel
}

func (s *TableKnowledgeVO) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *TableKnowledgeVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *TableKnowledgeVO) GetLevel() *int32 {
	return s.Level
}

func (s *TableKnowledgeVO) GetLevelType() *string {
	return s.LevelType
}

func (s *TableKnowledgeVO) GetLogic() *bool {
	return s.Logic
}

func (s *TableKnowledgeVO) GetNumRows() *int64 {
	return s.NumRows
}

func (s *TableKnowledgeVO) GetSchemaName() *string {
	return s.SchemaName
}

func (s *TableKnowledgeVO) GetSize() *int64 {
	return s.Size
}

func (s *TableKnowledgeVO) GetSummary() *string {
	return s.Summary
}

func (s *TableKnowledgeVO) GetTableId() *int64 {
	return s.TableId
}

func (s *TableKnowledgeVO) GetTableName() *string {
	return s.TableName
}

func (s *TableKnowledgeVO) GetTitle() *string {
	return s.Title
}

func (s *TableKnowledgeVO) SetAssetCreatedGmt(v string) *TableKnowledgeVO {
	s.AssetCreatedGmt = &v
	return s
}

func (s *TableKnowledgeVO) SetAssetDescription(v string) *TableKnowledgeVO {
	s.AssetDescription = &v
	return s
}

func (s *TableKnowledgeVO) SetAssetModifiedGmt(v string) *TableKnowledgeVO {
	s.AssetModifiedGmt = &v
	return s
}

func (s *TableKnowledgeVO) SetDbId(v int32) *TableKnowledgeVO {
	s.DbId = &v
	return s
}

func (s *TableKnowledgeVO) SetDbName(v string) *TableKnowledgeVO {
	s.DbName = &v
	return s
}

func (s *TableKnowledgeVO) SetDbType(v string) *TableKnowledgeVO {
	s.DbType = &v
	return s
}

func (s *TableKnowledgeVO) SetDescription(v string) *TableKnowledgeVO {
	s.Description = &v
	return s
}

func (s *TableKnowledgeVO) SetEnvType(v string) *TableKnowledgeVO {
	s.EnvType = &v
	return s
}

func (s *TableKnowledgeVO) SetHotLevel(v int32) *TableKnowledgeVO {
	s.HotLevel = &v
	return s
}

func (s *TableKnowledgeVO) SetInstanceId(v int32) *TableKnowledgeVO {
	s.InstanceId = &v
	return s
}

func (s *TableKnowledgeVO) SetInstanceName(v string) *TableKnowledgeVO {
	s.InstanceName = &v
	return s
}

func (s *TableKnowledgeVO) SetLevel(v int32) *TableKnowledgeVO {
	s.Level = &v
	return s
}

func (s *TableKnowledgeVO) SetLevelType(v string) *TableKnowledgeVO {
	s.LevelType = &v
	return s
}

func (s *TableKnowledgeVO) SetLogic(v bool) *TableKnowledgeVO {
	s.Logic = &v
	return s
}

func (s *TableKnowledgeVO) SetNumRows(v int64) *TableKnowledgeVO {
	s.NumRows = &v
	return s
}

func (s *TableKnowledgeVO) SetSchemaName(v string) *TableKnowledgeVO {
	s.SchemaName = &v
	return s
}

func (s *TableKnowledgeVO) SetSize(v int64) *TableKnowledgeVO {
	s.Size = &v
	return s
}

func (s *TableKnowledgeVO) SetSummary(v string) *TableKnowledgeVO {
	s.Summary = &v
	return s
}

func (s *TableKnowledgeVO) SetTableId(v int64) *TableKnowledgeVO {
	s.TableId = &v
	return s
}

func (s *TableKnowledgeVO) SetTableName(v string) *TableKnowledgeVO {
	s.TableName = &v
	return s
}

func (s *TableKnowledgeVO) SetTitle(v string) *TableKnowledgeVO {
	s.Title = &v
	return s
}

func (s *TableKnowledgeVO) Validate() error {
	return dara.Validate(s)
}
