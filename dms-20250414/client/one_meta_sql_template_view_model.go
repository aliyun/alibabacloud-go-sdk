// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaSqlTemplateView interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *OneMetaSqlTemplateView
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *OneMetaSqlTemplateView
	GetDatabaseUuid() *string
	SetDescription(v string) *OneMetaSqlTemplateView
	GetDescription() *string
	SetExpr(v string) *OneMetaSqlTemplateView
	GetExpr() *string
	SetGmtCreated(v int64) *OneMetaSqlTemplateView
	GetGmtCreated() *int64
	SetGmtModified(v int64) *OneMetaSqlTemplateView
	GetGmtModified() *int64
	SetKnowledgeUuid(v string) *OneMetaSqlTemplateView
	GetKnowledgeUuid() *string
	SetSource(v string) *OneMetaSqlTemplateView
	GetSource() *string
	SetSqlParams(v string) *OneMetaSqlTemplateView
	GetSqlParams() *string
	SetSummary(v string) *OneMetaSqlTemplateView
	GetSummary() *string
	SetTag(v string) *OneMetaSqlTemplateView
	GetTag() *string
	SetTitle(v string) *OneMetaSqlTemplateView
	GetTitle() *string
	SetVersion(v string) *OneMetaSqlTemplateView
	GetVersion() *string
}

type OneMetaSqlTemplateView struct {
	CatalogUuid   *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	DatabaseUuid  *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Expr          *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	GmtCreated    *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SqlParams     *string `json:"SqlParams,omitempty" xml:"SqlParams,omitempty"`
	Summary       *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OneMetaSqlTemplateView) String() string {
	return dara.Prettify(s)
}

func (s OneMetaSqlTemplateView) GoString() string {
	return s.String()
}

func (s *OneMetaSqlTemplateView) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *OneMetaSqlTemplateView) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *OneMetaSqlTemplateView) GetDescription() *string {
	return s.Description
}

func (s *OneMetaSqlTemplateView) GetExpr() *string {
	return s.Expr
}

func (s *OneMetaSqlTemplateView) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *OneMetaSqlTemplateView) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *OneMetaSqlTemplateView) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *OneMetaSqlTemplateView) GetSource() *string {
	return s.Source
}

func (s *OneMetaSqlTemplateView) GetSqlParams() *string {
	return s.SqlParams
}

func (s *OneMetaSqlTemplateView) GetSummary() *string {
	return s.Summary
}

func (s *OneMetaSqlTemplateView) GetTag() *string {
	return s.Tag
}

func (s *OneMetaSqlTemplateView) GetTitle() *string {
	return s.Title
}

func (s *OneMetaSqlTemplateView) GetVersion() *string {
	return s.Version
}

func (s *OneMetaSqlTemplateView) SetCatalogUuid(v string) *OneMetaSqlTemplateView {
	s.CatalogUuid = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetDatabaseUuid(v string) *OneMetaSqlTemplateView {
	s.DatabaseUuid = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetDescription(v string) *OneMetaSqlTemplateView {
	s.Description = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetExpr(v string) *OneMetaSqlTemplateView {
	s.Expr = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetGmtCreated(v int64) *OneMetaSqlTemplateView {
	s.GmtCreated = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetGmtModified(v int64) *OneMetaSqlTemplateView {
	s.GmtModified = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetKnowledgeUuid(v string) *OneMetaSqlTemplateView {
	s.KnowledgeUuid = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetSource(v string) *OneMetaSqlTemplateView {
	s.Source = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetSqlParams(v string) *OneMetaSqlTemplateView {
	s.SqlParams = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetSummary(v string) *OneMetaSqlTemplateView {
	s.Summary = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetTag(v string) *OneMetaSqlTemplateView {
	s.Tag = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetTitle(v string) *OneMetaSqlTemplateView {
	s.Title = &v
	return s
}

func (s *OneMetaSqlTemplateView) SetVersion(v string) *OneMetaSqlTemplateView {
	s.Version = &v
	return s
}

func (s *OneMetaSqlTemplateView) Validate() error {
	return dara.Validate(s)
}
