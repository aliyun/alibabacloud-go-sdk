// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticKnowledgeView interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *AgenticKnowledgeView
	GetCatalogUuid() *string
	SetColumnName(v string) *AgenticKnowledgeView
	GetColumnName() *string
	SetCreateTime(v int64) *AgenticKnowledgeView
	GetCreateTime() *int64
	SetDatabaseUuid(v string) *AgenticKnowledgeView
	GetDatabaseUuid() *string
	SetDescription(v string) *AgenticKnowledgeView
	GetDescription() *string
	SetEntityType(v string) *AgenticKnowledgeView
	GetEntityType() *string
	SetExtra(v map[string]interface{}) *AgenticKnowledgeView
	GetExtra() map[string]interface{}
	SetKnowledgeUuid(v string) *AgenticKnowledgeView
	GetKnowledgeUuid() *string
	SetLevel(v string) *AgenticKnowledgeView
	GetLevel() *string
	SetLocked(v bool) *AgenticKnowledgeView
	GetLocked() *bool
	SetLockedBy(v string) *AgenticKnowledgeView
	GetLockedBy() *string
	SetLockedTime(v int64) *AgenticKnowledgeView
	GetLockedTime() *int64
	SetModifyTime(v int64) *AgenticKnowledgeView
	GetModifyTime() *int64
	SetQualifiedName(v string) *AgenticKnowledgeView
	GetQualifiedName() *string
	SetSource(v string) *AgenticKnowledgeView
	GetSource() *string
	SetSummary(v string) *AgenticKnowledgeView
	GetSummary() *string
	SetTitle(v string) *AgenticKnowledgeView
	GetTitle() *string
	SetUnitCatalogUuid(v string) *AgenticKnowledgeView
	GetUnitCatalogUuid() *string
	SetUnitDatabaseUuid(v string) *AgenticKnowledgeView
	GetUnitDatabaseUuid() *string
	SetVersion(v string) *AgenticKnowledgeView
	GetVersion() *string
}

type AgenticKnowledgeView struct {
	CatalogUuid      *string                `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	ColumnName       *string                `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	CreateTime       *int64                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseUuid     *string                `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description      *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	EntityType       *string                `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Extra            map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	KnowledgeUuid    *string                `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	Level            *string                `json:"Level,omitempty" xml:"Level,omitempty"`
	Locked           *bool                  `json:"Locked,omitempty" xml:"Locked,omitempty"`
	LockedBy         *string                `json:"LockedBy,omitempty" xml:"LockedBy,omitempty"`
	LockedTime       *int64                 `json:"LockedTime,omitempty" xml:"LockedTime,omitempty"`
	ModifyTime       *int64                 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	QualifiedName    *string                `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	Source           *string                `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary          *string                `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title            *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	UnitCatalogUuid  *string                `json:"UnitCatalogUuid,omitempty" xml:"UnitCatalogUuid,omitempty"`
	UnitDatabaseUuid *string                `json:"UnitDatabaseUuid,omitempty" xml:"UnitDatabaseUuid,omitempty"`
	Version          *string                `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s AgenticKnowledgeView) String() string {
	return dara.Prettify(s)
}

func (s AgenticKnowledgeView) GoString() string {
	return s.String()
}

func (s *AgenticKnowledgeView) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *AgenticKnowledgeView) GetColumnName() *string {
	return s.ColumnName
}

func (s *AgenticKnowledgeView) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *AgenticKnowledgeView) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *AgenticKnowledgeView) GetDescription() *string {
	return s.Description
}

func (s *AgenticKnowledgeView) GetEntityType() *string {
	return s.EntityType
}

func (s *AgenticKnowledgeView) GetExtra() map[string]interface{} {
	return s.Extra
}

func (s *AgenticKnowledgeView) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *AgenticKnowledgeView) GetLevel() *string {
	return s.Level
}

func (s *AgenticKnowledgeView) GetLocked() *bool {
	return s.Locked
}

func (s *AgenticKnowledgeView) GetLockedBy() *string {
	return s.LockedBy
}

func (s *AgenticKnowledgeView) GetLockedTime() *int64 {
	return s.LockedTime
}

func (s *AgenticKnowledgeView) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *AgenticKnowledgeView) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *AgenticKnowledgeView) GetSource() *string {
	return s.Source
}

func (s *AgenticKnowledgeView) GetSummary() *string {
	return s.Summary
}

func (s *AgenticKnowledgeView) GetTitle() *string {
	return s.Title
}

func (s *AgenticKnowledgeView) GetUnitCatalogUuid() *string {
	return s.UnitCatalogUuid
}

func (s *AgenticKnowledgeView) GetUnitDatabaseUuid() *string {
	return s.UnitDatabaseUuid
}

func (s *AgenticKnowledgeView) GetVersion() *string {
	return s.Version
}

func (s *AgenticKnowledgeView) SetCatalogUuid(v string) *AgenticKnowledgeView {
	s.CatalogUuid = &v
	return s
}

func (s *AgenticKnowledgeView) SetColumnName(v string) *AgenticKnowledgeView {
	s.ColumnName = &v
	return s
}

func (s *AgenticKnowledgeView) SetCreateTime(v int64) *AgenticKnowledgeView {
	s.CreateTime = &v
	return s
}

func (s *AgenticKnowledgeView) SetDatabaseUuid(v string) *AgenticKnowledgeView {
	s.DatabaseUuid = &v
	return s
}

func (s *AgenticKnowledgeView) SetDescription(v string) *AgenticKnowledgeView {
	s.Description = &v
	return s
}

func (s *AgenticKnowledgeView) SetEntityType(v string) *AgenticKnowledgeView {
	s.EntityType = &v
	return s
}

func (s *AgenticKnowledgeView) SetExtra(v map[string]interface{}) *AgenticKnowledgeView {
	s.Extra = v
	return s
}

func (s *AgenticKnowledgeView) SetKnowledgeUuid(v string) *AgenticKnowledgeView {
	s.KnowledgeUuid = &v
	return s
}

func (s *AgenticKnowledgeView) SetLevel(v string) *AgenticKnowledgeView {
	s.Level = &v
	return s
}

func (s *AgenticKnowledgeView) SetLocked(v bool) *AgenticKnowledgeView {
	s.Locked = &v
	return s
}

func (s *AgenticKnowledgeView) SetLockedBy(v string) *AgenticKnowledgeView {
	s.LockedBy = &v
	return s
}

func (s *AgenticKnowledgeView) SetLockedTime(v int64) *AgenticKnowledgeView {
	s.LockedTime = &v
	return s
}

func (s *AgenticKnowledgeView) SetModifyTime(v int64) *AgenticKnowledgeView {
	s.ModifyTime = &v
	return s
}

func (s *AgenticKnowledgeView) SetQualifiedName(v string) *AgenticKnowledgeView {
	s.QualifiedName = &v
	return s
}

func (s *AgenticKnowledgeView) SetSource(v string) *AgenticKnowledgeView {
	s.Source = &v
	return s
}

func (s *AgenticKnowledgeView) SetSummary(v string) *AgenticKnowledgeView {
	s.Summary = &v
	return s
}

func (s *AgenticKnowledgeView) SetTitle(v string) *AgenticKnowledgeView {
	s.Title = &v
	return s
}

func (s *AgenticKnowledgeView) SetUnitCatalogUuid(v string) *AgenticKnowledgeView {
	s.UnitCatalogUuid = &v
	return s
}

func (s *AgenticKnowledgeView) SetUnitDatabaseUuid(v string) *AgenticKnowledgeView {
	s.UnitDatabaseUuid = &v
	return s
}

func (s *AgenticKnowledgeView) SetVersion(v string) *AgenticKnowledgeView {
	s.Version = &v
	return s
}

func (s *AgenticKnowledgeView) Validate() error {
	return dara.Validate(s)
}
