// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumnKnowledgeVO interface {
	dara.Model
	String() string
	GoString() string
	SetAssetDescription(v string) *ColumnKnowledgeVO
	GetAssetDescription() *string
	SetAssetModifiedGmt(v string) *ColumnKnowledgeVO
	GetAssetModifiedGmt() *string
	SetAutoIncrement(v bool) *ColumnKnowledgeVO
	GetAutoIncrement() *bool
	SetColumnId(v int64) *ColumnKnowledgeVO
	GetColumnId() *int64
	SetColumnKeyType(v string) *ColumnKnowledgeVO
	GetColumnKeyType() *string
	SetColumnName(v string) *ColumnKnowledgeVO
	GetColumnName() *string
	SetColumnType(v string) *ColumnKnowledgeVO
	GetColumnType() *string
	SetDescription(v string) *ColumnKnowledgeVO
	GetDescription() *string
	SetHotLevel(v int32) *ColumnKnowledgeVO
	GetHotLevel() *int32
	SetLevel(v int32) *ColumnKnowledgeVO
	GetLevel() *int32
	SetLevelType(v string) *ColumnKnowledgeVO
	GetLevelType() *string
	SetNullable(v bool) *ColumnKnowledgeVO
	GetNullable() *bool
	SetPlain(v bool) *ColumnKnowledgeVO
	GetPlain() *bool
	SetPosition(v int32) *ColumnKnowledgeVO
	GetPosition() *int32
	SetSecurityLevel(v string) *ColumnKnowledgeVO
	GetSecurityLevel() *string
	SetTableId(v int64) *ColumnKnowledgeVO
	GetTableId() *int64
	SetTitle(v string) *ColumnKnowledgeVO
	GetTitle() *string
	SetUserSensitivityLevel(v string) *ColumnKnowledgeVO
	GetUserSensitivityLevel() *string
}

type ColumnKnowledgeVO struct {
	AssetDescription     *string `json:"AssetDescription,omitempty" xml:"AssetDescription,omitempty"`
	AssetModifiedGmt     *string `json:"AssetModifiedGmt,omitempty" xml:"AssetModifiedGmt,omitempty"`
	AutoIncrement        *bool   `json:"AutoIncrement,omitempty" xml:"AutoIncrement,omitempty"`
	ColumnId             *int64  `json:"ColumnId,omitempty" xml:"ColumnId,omitempty"`
	ColumnKeyType        *string `json:"ColumnKeyType,omitempty" xml:"ColumnKeyType,omitempty"`
	ColumnName           *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnType           *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HotLevel             *int32  `json:"HotLevel,omitempty" xml:"HotLevel,omitempty"`
	Level                *int32  `json:"Level,omitempty" xml:"Level,omitempty"`
	LevelType            *string `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	Nullable             *bool   `json:"Nullable,omitempty" xml:"Nullable,omitempty"`
	Plain                *bool   `json:"Plain,omitempty" xml:"Plain,omitempty"`
	Position             *int32  `json:"Position,omitempty" xml:"Position,omitempty"`
	SecurityLevel        *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	TableId              *int64  `json:"TableId,omitempty" xml:"TableId,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserSensitivityLevel *string `json:"UserSensitivityLevel,omitempty" xml:"UserSensitivityLevel,omitempty"`
}

func (s ColumnKnowledgeVO) String() string {
	return dara.Prettify(s)
}

func (s ColumnKnowledgeVO) GoString() string {
	return s.String()
}

func (s *ColumnKnowledgeVO) GetAssetDescription() *string {
	return s.AssetDescription
}

func (s *ColumnKnowledgeVO) GetAssetModifiedGmt() *string {
	return s.AssetModifiedGmt
}

func (s *ColumnKnowledgeVO) GetAutoIncrement() *bool {
	return s.AutoIncrement
}

func (s *ColumnKnowledgeVO) GetColumnId() *int64 {
	return s.ColumnId
}

func (s *ColumnKnowledgeVO) GetColumnKeyType() *string {
	return s.ColumnKeyType
}

func (s *ColumnKnowledgeVO) GetColumnName() *string {
	return s.ColumnName
}

func (s *ColumnKnowledgeVO) GetColumnType() *string {
	return s.ColumnType
}

func (s *ColumnKnowledgeVO) GetDescription() *string {
	return s.Description
}

func (s *ColumnKnowledgeVO) GetHotLevel() *int32 {
	return s.HotLevel
}

func (s *ColumnKnowledgeVO) GetLevel() *int32 {
	return s.Level
}

func (s *ColumnKnowledgeVO) GetLevelType() *string {
	return s.LevelType
}

func (s *ColumnKnowledgeVO) GetNullable() *bool {
	return s.Nullable
}

func (s *ColumnKnowledgeVO) GetPlain() *bool {
	return s.Plain
}

func (s *ColumnKnowledgeVO) GetPosition() *int32 {
	return s.Position
}

func (s *ColumnKnowledgeVO) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *ColumnKnowledgeVO) GetTableId() *int64 {
	return s.TableId
}

func (s *ColumnKnowledgeVO) GetTitle() *string {
	return s.Title
}

func (s *ColumnKnowledgeVO) GetUserSensitivityLevel() *string {
	return s.UserSensitivityLevel
}

func (s *ColumnKnowledgeVO) SetAssetDescription(v string) *ColumnKnowledgeVO {
	s.AssetDescription = &v
	return s
}

func (s *ColumnKnowledgeVO) SetAssetModifiedGmt(v string) *ColumnKnowledgeVO {
	s.AssetModifiedGmt = &v
	return s
}

func (s *ColumnKnowledgeVO) SetAutoIncrement(v bool) *ColumnKnowledgeVO {
	s.AutoIncrement = &v
	return s
}

func (s *ColumnKnowledgeVO) SetColumnId(v int64) *ColumnKnowledgeVO {
	s.ColumnId = &v
	return s
}

func (s *ColumnKnowledgeVO) SetColumnKeyType(v string) *ColumnKnowledgeVO {
	s.ColumnKeyType = &v
	return s
}

func (s *ColumnKnowledgeVO) SetColumnName(v string) *ColumnKnowledgeVO {
	s.ColumnName = &v
	return s
}

func (s *ColumnKnowledgeVO) SetColumnType(v string) *ColumnKnowledgeVO {
	s.ColumnType = &v
	return s
}

func (s *ColumnKnowledgeVO) SetDescription(v string) *ColumnKnowledgeVO {
	s.Description = &v
	return s
}

func (s *ColumnKnowledgeVO) SetHotLevel(v int32) *ColumnKnowledgeVO {
	s.HotLevel = &v
	return s
}

func (s *ColumnKnowledgeVO) SetLevel(v int32) *ColumnKnowledgeVO {
	s.Level = &v
	return s
}

func (s *ColumnKnowledgeVO) SetLevelType(v string) *ColumnKnowledgeVO {
	s.LevelType = &v
	return s
}

func (s *ColumnKnowledgeVO) SetNullable(v bool) *ColumnKnowledgeVO {
	s.Nullable = &v
	return s
}

func (s *ColumnKnowledgeVO) SetPlain(v bool) *ColumnKnowledgeVO {
	s.Plain = &v
	return s
}

func (s *ColumnKnowledgeVO) SetPosition(v int32) *ColumnKnowledgeVO {
	s.Position = &v
	return s
}

func (s *ColumnKnowledgeVO) SetSecurityLevel(v string) *ColumnKnowledgeVO {
	s.SecurityLevel = &v
	return s
}

func (s *ColumnKnowledgeVO) SetTableId(v int64) *ColumnKnowledgeVO {
	s.TableId = &v
	return s
}

func (s *ColumnKnowledgeVO) SetTitle(v string) *ColumnKnowledgeVO {
	s.Title = &v
	return s
}

func (s *ColumnKnowledgeVO) SetUserSensitivityLevel(v string) *ColumnKnowledgeVO {
	s.UserSensitivityLevel = &v
	return s
}

func (s *ColumnKnowledgeVO) Validate() error {
	return dara.Validate(s)
}
