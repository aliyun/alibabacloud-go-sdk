// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOssieModelView interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *OssieModelView
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *OssieModelView
	GetDatabaseUuid() *string
	SetDescription(v string) *OssieModelView
	GetDescription() *string
	SetDocFormat(v string) *OssieModelView
	GetDocFormat() *string
	SetDomainTopic(v string) *OssieModelView
	GetDomainTopic() *string
	SetExpr(v string) *OssieModelView
	GetExpr() *string
	SetGmtCreated(v int64) *OssieModelView
	GetGmtCreated() *int64
	SetGmtModified(v int64) *OssieModelView
	GetGmtModified() *int64
	SetKnowledgeUuid(v string) *OssieModelView
	GetKnowledgeUuid() *string
	SetRawDoc(v string) *OssieModelView
	GetRawDoc() *string
	SetSemanticType(v string) *OssieModelView
	GetSemanticType() *string
	SetSource(v string) *OssieModelView
	GetSource() *string
	SetSummary(v string) *OssieModelView
	GetSummary() *string
	SetTag(v string) *OssieModelView
	GetTag() *string
	SetTitle(v string) *OssieModelView
	GetTitle() *string
	SetVersion(v string) *OssieModelView
	GetVersion() *string
}

type OssieModelView struct {
	CatalogUuid   *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	DatabaseUuid  *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocFormat     *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	DomainTopic   *string `json:"DomainTopic,omitempty" xml:"DomainTopic,omitempty"`
	Expr          *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	GmtCreated    *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	RawDoc        *string `json:"RawDoc,omitempty" xml:"RawDoc,omitempty"`
	SemanticType  *string `json:"SemanticType,omitempty" xml:"SemanticType,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary       *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag           *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OssieModelView) String() string {
	return dara.Prettify(s)
}

func (s OssieModelView) GoString() string {
	return s.String()
}

func (s *OssieModelView) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *OssieModelView) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *OssieModelView) GetDescription() *string {
	return s.Description
}

func (s *OssieModelView) GetDocFormat() *string {
	return s.DocFormat
}

func (s *OssieModelView) GetDomainTopic() *string {
	return s.DomainTopic
}

func (s *OssieModelView) GetExpr() *string {
	return s.Expr
}

func (s *OssieModelView) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *OssieModelView) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *OssieModelView) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *OssieModelView) GetRawDoc() *string {
	return s.RawDoc
}

func (s *OssieModelView) GetSemanticType() *string {
	return s.SemanticType
}

func (s *OssieModelView) GetSource() *string {
	return s.Source
}

func (s *OssieModelView) GetSummary() *string {
	return s.Summary
}

func (s *OssieModelView) GetTag() *string {
	return s.Tag
}

func (s *OssieModelView) GetTitle() *string {
	return s.Title
}

func (s *OssieModelView) GetVersion() *string {
	return s.Version
}

func (s *OssieModelView) SetCatalogUuid(v string) *OssieModelView {
	s.CatalogUuid = &v
	return s
}

func (s *OssieModelView) SetDatabaseUuid(v string) *OssieModelView {
	s.DatabaseUuid = &v
	return s
}

func (s *OssieModelView) SetDescription(v string) *OssieModelView {
	s.Description = &v
	return s
}

func (s *OssieModelView) SetDocFormat(v string) *OssieModelView {
	s.DocFormat = &v
	return s
}

func (s *OssieModelView) SetDomainTopic(v string) *OssieModelView {
	s.DomainTopic = &v
	return s
}

func (s *OssieModelView) SetExpr(v string) *OssieModelView {
	s.Expr = &v
	return s
}

func (s *OssieModelView) SetGmtCreated(v int64) *OssieModelView {
	s.GmtCreated = &v
	return s
}

func (s *OssieModelView) SetGmtModified(v int64) *OssieModelView {
	s.GmtModified = &v
	return s
}

func (s *OssieModelView) SetKnowledgeUuid(v string) *OssieModelView {
	s.KnowledgeUuid = &v
	return s
}

func (s *OssieModelView) SetRawDoc(v string) *OssieModelView {
	s.RawDoc = &v
	return s
}

func (s *OssieModelView) SetSemanticType(v string) *OssieModelView {
	s.SemanticType = &v
	return s
}

func (s *OssieModelView) SetSource(v string) *OssieModelView {
	s.Source = &v
	return s
}

func (s *OssieModelView) SetSummary(v string) *OssieModelView {
	s.Summary = &v
	return s
}

func (s *OssieModelView) SetTag(v string) *OssieModelView {
	s.Tag = &v
	return s
}

func (s *OssieModelView) SetTitle(v string) *OssieModelView {
	s.Title = &v
	return s
}

func (s *OssieModelView) SetVersion(v string) *OssieModelView {
	s.Version = &v
	return s
}

func (s *OssieModelView) Validate() error {
	return dara.Validate(s)
}
