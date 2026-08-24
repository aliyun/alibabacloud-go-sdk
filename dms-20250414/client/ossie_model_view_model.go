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
	// The UUID of the associated instance.
	//
	// example:
	//
	// mc-SH-cd3ns***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the associated database.
	//
	// example:
	//
	// md-SH-q8XzcK***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The semantic description.
	//
	// example:
	//
	// Order summary
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The semantic document format type.
	//
	// example:
	//
	// JSON
	DocFormat *string `json:"DocFormat,omitempty" xml:"DocFormat,omitempty"`
	// The domain topic.
	//
	// example:
	//
	// Order
	DomainTopic *string `json:"DomainTopic,omitempty" xml:"DomainTopic,omitempty"`
	// The expression content.
	//
	// example:
	//
	// select 1
	Expr *string `json:"Expr,omitempty" xml:"Expr,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1663809374000
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 1780539699000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The UUID of the semantic knowledge.
	//
	// example:
	//
	// dfb58bd***
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	// The semantic document content.
	//
	// example:
	//
	// {
	//
	//   "version": "0.2.0.dev0",
	//
	//   "semantic_model": [
	//
	//     {
	//
	//       "name": "sales",
	//
	//       "datasets": [
	//
	//         {
	//
	//           "name": "orders",
	//
	//           "source": "analytics.public.orders"
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	//   ]
	//
	// }
	RawDoc *string `json:"RawDoc,omitempty" xml:"RawDoc,omitempty"`
	// The semantic type.
	//
	// example:
	//
	// Ossie
	SemanticType *string `json:"SemanticType,omitempty" xml:"SemanticType,omitempty"`
	// The source.
	//
	// example:
	//
	// USER_EDIT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The document summary.
	//
	// example:
	//
	// knowledge summary
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The knowledge base tag.
	//
	// example:
	//
	// 1dq7qod8hxtt1***
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The semantic title.
	//
	// example:
	//
	// Order total
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The version information.
	//
	// example:
	//
	// 0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
